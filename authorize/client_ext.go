package authorize

import (
	"crypto/rand"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SDKInterfaceFunc func() (any, *http.Response, error)

var (
	maxRetries                       = 10
	maximumRetryAfterBackoff         = 30
	maximumRetryAfterBackoffDuration = time.Duration(maximumRetryAfterBackoff) * time.Second
	requestMutex                     sync.Mutex
)

func processResponse(f SDKInterfaceFunc, targetObject any) (*http.Response, error) {

	obj, response, error := exponentialBackOffRetry(f)

	if targetObject != nil {
		v := reflect.ValueOf(targetObject)
		if v.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("Target object must be a pointer.  This is always a problem with the SDK, please raise an issue with the SDK maintainers.")
		}
		if !v.Elem().IsValid() {
			return nil, fmt.Errorf("Target object is not valid.  This is always a problem with the SDK, please raise an issue with the SDK maintainers.")
		}

		if obj != nil {
			v.Elem().Set(reflect.ValueOf(obj))
		}
	}

	return response, reformError(error)
}

func reformError(err error) error {
	return err
}

func exponentialBackOffRetry(f SDKInterfaceFunc) (interface{}, *http.Response, error) {
	// This API action isn't thread safe - the environment may be switched by another thread.  We need to lock it
	requestMutex.Lock()
	defer requestMutex.Unlock()

	var obj interface{}
	var resp *http.Response
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxRetries; i++ {
		obj, resp, err = f()
		retryAttempt := i + 1

		backOffTime, isRetryable = testForRetryable(resp, err, retryAttempt)

		if isRetryable {
			slog.Info("Attempt failed, backing off by calculated duration.", "retry attempt", retryAttempt, "error", err, "backoff duration", backOffTime.String())
			time.Sleep(backOffTime)
			continue
		}

		return obj, resp, err
	}

	slog.Warn("Request failed after maximum number of attempts", "max retries", maxRetries)

	return obj, resp, err // output the final error
}

func testForRetryable(r *http.Response, err error, retryAttempt int) (time.Duration, bool) {

	baseDelay := time.Second
	requestDelayDuration, ebErr := calculateExponentialBackoff(retryAttempt, baseDelay)
	if ebErr != nil {
		slog.Error("Invalid backoff delay duration", "error", ebErr, "baseDelay", baseDelay, "retry", false)
		return maximumRetryAfterBackoffDuration, false
	}

	if r != nil {
		if r.StatusCode == http.StatusNotImplemented || r.StatusCode == http.StatusServiceUnavailable || r.StatusCode == http.StatusTooManyRequests {
			retryAfter, err := parseRetryAfterHeader(r)
			if err != nil {
				slog.Warn("Cannot parse the expected \"Retry-After\" header", "error", err)
			}

			if err == nil {
				if retryAfter >= maximumRetryAfterBackoffDuration {
					return maximumRetryAfterBackoffDuration, true // optimistically set to the maximum if beyond
				} else {
					return retryAfter, true
				}
			}
		}

		retryAbleCodes := []int{
			http.StatusRequestTimeout,
			http.StatusTooManyRequests,
			http.StatusInternalServerError,
			http.StatusNotImplemented,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		}

		if slices.Contains(retryAbleCodes, r.StatusCode) {
			slog.Info("HTTP status code detected, available for retry", "status code", r.StatusCode)
			return requestDelayDuration, true
		}
	}

	if err != nil {
		if genericOAError, ok := err.(*GenericOpenAPIError); ok && genericOAError.Model() != nil {
			// We have an application level error

			if modelError, ok := genericOAError.Model().(P1Error); ok {

				// Test for unexpected errors
				if strings.EqualFold(modelError.GetCode(), "UNEXPECTED_ERROR") {
					slog.Info("API reports unexpected error, available for retry")
					return requestDelayDuration, true
				}

				// Test for inconsistent role state
				m, _ := regexp.MatchString(`^Role assignment [a-z0-9\-]* cannot be deleted as it is read only`, modelError.GetMessage())

				if m {
					slog.Info("API reports inconsistent role assignment, available for retry")
					return requestDelayDuration, true
				}
			}
		}
	}

	return requestDelayDuration, false
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, error) {
	retryAfterHeader := resp.Header.Get("Retry-After")

	if retryAfterHeader == "" {
		return 0, fmt.Errorf("Retry-After header not found")
	}

	retryAfterSeconds, err := strconv.Atoi(retryAfterHeader)

	if err == nil {
		return time.Duration(retryAfterSeconds) * time.Second, nil
	}

	retryAfterTime, err := http.ParseTime(retryAfterHeader)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse Retry-After header value: %v", err)
	}

	return time.Until(retryAfterTime), nil
}

func calculateExponentialBackoff(attempt int, baseDelay time.Duration) (time.Duration, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(101))
	if err != nil {
		return 0, err
	}

	if !n.IsInt64() {
		return 0, fmt.Errorf("Generated random jitter value is too large. This is always a problem with the SDK. Please raise an issue with the SDK maintainers.")
	}

	jitter := time.Duration(n.Int64()) * time.Millisecond // Add random jitter
	calculatedBackOff := baseDelay*time.Duration(math.Pow(2, float64(attempt))) + jitter

	slog.Debug("Calculated backoff duration", "maximumRetryAfterBackoffDuration", maximumRetryAfterBackoffDuration.String(), "attempt", attempt, "baseDelay", baseDelay.String(), "jitter", jitter.String(), "calculatedBackOff", calculatedBackOff.String())

	return time.Duration(math.Min(float64(calculatedBackOff), float64(maximumRetryAfterBackoffDuration))), nil
}
