/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LinkedAccountsApiService LinkedAccountsApi service
type LinkedAccountsApiService service

type ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest struct {
	ctx context.Context
	ApiService *LinkedAccountsApiService
	environmentID string
	userID string
	xPingExternalTransactionID *string
	xPingExternalSessionID *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet READ Linked Accounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param userID
 @return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest
*/
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet(ctx context.Context, environmentID string, userID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		userID: userID,
	}
}

// Execute executes the request
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *LinkedAccountsApiService) internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkedAccountsApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/linkedAccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest struct {
	ctx context.Context
	ApiService *LinkedAccountsApiService
	environmentID string
	userID string
	linkedAccountID string
	xPingExternalTransactionID *string
	xPingExternalSessionID *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete DELETE Linked Account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param userID
 @param linkedAccountID
 @return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest
*/
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(ctx context.Context, environmentID string, userID string, linkedAccountID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		userID: userID,
		linkedAccountID: linkedAccountID,
	}
}

// Execute executes the request
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *LinkedAccountsApiService) internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkedAccountsApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkedAccountID"+"}", url.PathEscape(parameterValueToString(r.linkedAccountID, "linkedAccountID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest struct {
	ctx context.Context
	ApiService *LinkedAccountsApiService
	environmentID string
	userID string
	linkedAccountID string
	xPingExternalTransactionID *string
	xPingExternalSessionID *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60; 
func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r)
}

/*
EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet READ One Linked Account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param userID
 @param linkedAccountID
 @return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest
*/
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet(ctx context.Context, environmentID string, userID string, linkedAccountID string) ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest {
	return ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		userID: userID,
		linkedAccountID: linkedAccountID,
	}
}

// Execute executes the request
func (a *LinkedAccountsApiService) EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *LinkedAccountsApiService) internalEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetExecute(r ApiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinkedAccountsApiService.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkedAccountID"+"}", url.PathEscape(parameterValueToString(r.linkedAccountID, "linkedAccountID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
