# v0.9.0 (Unreleased)

* **Feature** Add support for Authorize Editor Attributes API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Conditions API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Connector Templates API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Policies API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Processors API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Rules API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Services API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)
* **Feature** Add support for Authorize Editor Statements API. [#372](https://github.com/patrickcping/pingone-go-sdk-v2/pull/372)

# v0.8.0 (2025-02-10)

* **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
* **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)

# v0.7.0 (2024-11-15)

* **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
* **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
* **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)

# v0.6.0 (2024-07-04)

* **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)

# v0.5.0 (2024-06-07)

* **Note** Removed unused `Operations` field from the `APIServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)
* **Feature** Add support for Application Resource Permissions API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for Application Resources API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for Application Roles API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for Application Role Assignments API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for Application Role Permissions API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for Application Roles API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
* **Feature** Add support for API Server Operations. [#354](https://github.com/patrickcping/pingone-go-sdk-v2/pull/354)
* **Feature** Add support for API Server Deployment.
* **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
* **Enhancement** Added `AccessControl` and `Directory` fields to the `APIServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)
* **Enhancement** Added `Type` field to the `APIServerAuthorizationServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)

# v0.4.1 (2024-01-12)

* **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
* **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)

# v0.4.0 (2023-11-10)

* **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)

# v0.3.1 (2023-10-16)

* **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)

# v0.3.0 (2023-08-08)

* **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)

# v0.2.1 (2023-07-12)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)

# v0.2.0 (2023-07-04)

* **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
* **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# v0.1.7 (2023-05-19)

* **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)

# v0.1.6 (2023-04-28)

* **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)

# v0.1.5 (2023-04-24)

* **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)

# v0.1.4 (2023-03-20)

* **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)

# v0.1.3 (2023-03-10)

* **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)

# v0.1.2 (2023-02-22)

* **Note** bump `golang.org/x/net` v0.2.0 => v0.7.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
* **Note** bump `golang.org/x/oauth2` v0.2.0 => v0.5.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)

# v0.1.1 (2023-01-09)

* **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)

# v0.1.0 (2022-10-09)

* **Feature** Initial release [#55](https://github.com/patrickcping/pingone-go-sdk-v2/pull/55)