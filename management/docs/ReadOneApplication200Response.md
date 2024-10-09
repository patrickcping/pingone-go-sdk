# ReadOneApplication200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessControl** | Pointer to [**ApplicationAccessControl**](ApplicationAccessControl.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the application. | [optional] 
**Enabled** | **bool** | A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**HiddenFromAppPortal** | Pointer to **bool** | A boolean to specify whether the application is hidden in the application portal despite the configured group access policy. | [optional] 
**Icon** | Pointer to [**ApplicationIcon**](ApplicationIcon.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**LoginPageUrl** | Pointer to **string** | A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains. | [optional] 
**Name** | **string** | A string that specifies the name of the application. This is a required property. | 
**Protocol** | [**EnumApplicationProtocol**](EnumApplicationProtocol.md) |  | 
**Type** | [**EnumApplicationType**](EnumApplicationType.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**HomePageUrl** | **string** | A string that specifies the custom home page URL for the application. | 
**AcsUrls** | **[]string** | A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property. | 
**AssertionDuration** | **int32** | An integer that specifies the assertion validity duration in seconds. This is a required property. | 
**AssertionSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion itself should be signed. The default value is &#x60;true&#x60;. | [optional] [default to true]
**CorsSettings** | Pointer to [**ApplicationCorsSettings**](ApplicationCorsSettings.md) |  | [optional] 
**DefaultTargetUrl** | Pointer to **string** | This is used as the RelayState parameter by the IdP to deep link into the application after authentication. This value can be overridden by the applicationUrl query parameter for GET Identity Provider Initiated SSO. Although both of these parameters are generally URLs, because they are used as deep links, this is not enforced. If neither defaultTargetUrl nor applicationUrl is specified during a SAML authentication flow, no RelayState value is supplied to the application. The defaultTargetUrl (or the applicationUrl) value is passed to the SAML application&#39;s ACS URL as a separate RelayState key value (not within the SAMLResponse key value). | [optional] 
**EnableRequestedAuthnContext** | Pointer to **bool** | Indicates whether &#x60;requestedAuthnContext&#x60; is taken into account in policy decision-making during authentication. | [optional] 
**IdpSigning** | [**ApplicationWSFEDAllOfIdpSigning**](ApplicationWSFEDAllOfIdpSigning.md) |  | 
**NameIdFormat** | Pointer to **string** | A string that specifies the format of the Subject NameID attibute in the SAML assertion | [optional] 
**ResponseSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion response itself should be signed. The default value is &#x60;false&#x60;. | [optional] [default to false]
**SloBinding** | Pointer to [**EnumApplicationSAMLSloBinding**](EnumApplicationSAMLSloBinding.md) |  | [optional] [default to ENUMAPPLICATIONSAMLSLOBINDING_POST]
**SloEndpoint** | Pointer to **string** | The single logout endpoint URL. | [optional] 
**SloResponseEndpoint** | Pointer to **string** | A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response. | [optional] 
**SloWindow** | Pointer to **int32** | Defines how long PingOne can exchange logout messages with the application, specifically a &#x60;LogoutRequest&#x60; from the application, since the initial request. PingOne can also send a &#x60;LogoutRequest&#x60; to the application when a single logout is initiated by the user from other session participants, such as an application or identity provider. This setting is per application. The SLO logout is separate from the user session logout that revokes all tokens. | [optional] 
**SpEncryption** | Pointer to [**ApplicationSAMLAllOfSpEncryption**](ApplicationSAMLAllOfSpEncryption.md) |  | [optional] 
**SpEntityId** | **string** | A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment. | 
**SpVerification** | Pointer to [**ApplicationSAMLAllOfSpVerification**](ApplicationSAMLAllOfSpVerification.md) |  | [optional] 
**Template** | Pointer to [**ApplicationTemplate**](ApplicationTemplate.md) |  | [optional] 
**AdditionalRefreshTokenReplayProtectionEnabled** | Pointer to **bool** | When set to &#x60;true&#x60; (the default), if you attempt to reuse the refresh token, the authorization server immediately revokes the reused refresh token, as well as all descendant tokens. Setting this to null equates to a &#x60;false&#x60; setting. | [optional] [default to true]
**AllowWildcardInRedirectUris** | Pointer to **bool** | A boolean to specify whether wildcards are allowed in redirect URIs. For more information, see [Wildcards in Redirect URIs](https://docs.pingidentity.com/csh?context&#x3D;p1_c_wildcard_redirect_uri). | [optional] 
**AssignActorRoles** | Pointer to **bool** | A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request. | [optional] 
**ClientId** | Pointer to **string** | (Required when &#x60;clientSecret&#x60; is specified.) Supported only for &#x60;POST&#x60; operations. This is the UUID of an external application that is being migrated to PingOne. The UUID must be a minimum of 8 alpha-numeric characters, and must be globally unique in PingOne. | [optional] 
**ClientSecret** | Pointer to **string** | (Required when &#x60;clientId&#x60; is specified.) Supported only for &#x60;POST&#x60; operations. This is the client secret associated with &#x60;clientId&#x60; for an external application that is being migrated to PingOne. This must be a minimum of 8 alpha-numeric characters. | [optional] 
**DevicePathId** | Pointer to **string** | A string that specifies a unique identifier within an environment for a device authorization grant flow to provide a short identifier to the application. This property is ignored when the &#x60;deviceCustomVerificationUri&#x60; property is configured. The string can contain any letters, numbers, and some special characters (regex &#x60;a-zA-Z0-9_-&#x60;). It can have a length of no more than 50 characters (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;50&#x60;). | [optional] 
**DeviceCustomVerificationUri** | Pointer to **string** | A string that specifies an optional custom verification URI that is returned for the &#x60;/device_authorization&#x60; endpoint. | [optional] 
**DeviceTimeout** | Pointer to **int32** | An integer that specifies the length of time (in seconds) that the &#x60;userCode&#x60; and &#x60;deviceCode&#x60; returned by the &#x60;/device_authorization&#x60; endpoint are valid. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;600&#x60; seconds. It can have a value of no more than &#x60;3600&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;3600&#x60;). | [optional] [default to 600]
**DevicePollingInterval** | Pointer to **int32** | An integer that specifies the frequency (in seconds) for the client to poll the &#x60;/as/token&#x60; endpoint. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;5&#x60; seconds. It can have a value of no more than &#x60;60&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;60&#x60;). | [optional] [default to 5]
**Jwks** | Pointer to **string** | A JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwksUrl&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwksUrl&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**JwksUrl** | Pointer to **string** | A URL (supports &#x60;https://&#x60; only) that provides access to a JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwks&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwks&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**Mobile** | Pointer to [**ApplicationOIDCAllOfMobile**](ApplicationOIDCAllOfMobile.md) |  | [optional] 
**BundleId** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.bundleId&#x60; instead.  A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  | [optional] 
**PackageName** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.packageName&#x60; instead.  A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  | [optional] 
**Kerberos** | Pointer to [**ApplicationWSFEDAllOfKerberos**](ApplicationWSFEDAllOfKerberos.md) |  | [optional] 
**GrantTypes** | [**[]EnumApplicationOIDCGrantType**](EnumApplicationOIDCGrantType.md) | A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS. | 
**InitiateLoginUri** | Pointer to **string** | A string that specifies the URI to use for third-parties to begin the sign-on process for the application. If specified, PingOne redirects users to this URI to initiate SSO to PingOne. The application is responsible for implementing the relevant OIDC flow when the initiate login URI is requested. This property is required if you want the application to appear in the PingOne Application Portal. See the OIDC specification section of [Initiating Login from a Third Party](https://openid.net/specs/openid-connect-core-1_0.html#ThirdPartyInitiatedLogin) for more information. | [optional] 
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | A string that specifies the URLs that the browser can be redirected to after logout. | [optional] 
**RedirectUris** | Pointer to **[]string** | A string that specifies the callback URI for the authentication response. | [optional] 
**RefreshTokenDuration** | Pointer to **int32** | An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the &#x60;refreshTokenRollingDuration&#x60; property is specified for the application, then this property must be less than or equal to the value of &#x60;refreshTokenRollingDuration&#x60;. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] [default to 2592000]
**RefreshTokenRollingDuration** | Pointer to **int32** | An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] 
**RefreshTokenRollingGracePeriodDuration** | Pointer to **int32** | The number of seconds that a refresh token may be reused after having been exchanged for a new set of tokens. This is useful in the case of network errors on the client. Valid values are between 0 and 86400 seconds. Null is treated the same as 0. | [optional] 
**ResponseTypes** | Pointer to [**[]EnumApplicationOIDCResponseType**](EnumApplicationOIDCResponseType.md) | The code or token type returned by an authorization request. Options are &#x60;TOKEN&#x60;, &#x60;ID_TOKEN&#x60;, and &#x60;CODE&#x60;. For hybrid flows that specify &#x60;CODE&#x60; with &#x60;TOKEN&#x60; or &#x60;ID_TOKEN&#x60;, see [Hybrid grant type](https://apidocs.pingidentity.com/pingone/main/v1/api/#hybrid-grant-type). | [optional] 
**RequireSignedRequestObject** | Pointer to **bool** | Indicates that the Java Web Token (JWT) for the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter is required to be signed. If &#x60;false&#x60; or null (default), a signed request object is not required. Both &#x60;supportUnsignedRequestObject&#x60; and this property cannot be set to &#x60;true&#x60;. | [optional] 
**SupportUnsignedRequestObject** | Pointer to **bool** | A boolean that specifies whether the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed. | [optional] 
**Tags** | Pointer to [**[]EnumApplicationTags**](EnumApplicationTags.md) | An array that specifies the list of labels associated with the application. Options are &#x60;PING_FED_CONNECTION_INTEGRATION&#x60;.  Only applicable for creating worker applications. | [optional] 
**TargetLinkUri** | Pointer to **string** | The URI for the application. If specified, PingOne will redirect application users to this URI after a user is authenticated. In the PingOne admin console, this becomes the value of the &#x60;target_link_uri&#x60; parameter used for the Initiate Single Sign-On URL field. | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 
**ParRequirement** | Pointer to [**EnumApplicationOIDCPARRequirement**](EnumApplicationOIDCPARRequirement.md) |  | [optional] [default to ENUMAPPLICATIONOIDCPARREQUIREMENT_OPTIONAL]
**ParTimeout** | Pointer to **int32** | PAR timeout in seconds. Must be between &#x60;1&#x60; and &#x60;600&#x60;. The default value is &#x60;60&#x60;. | [optional] [default to 60]
**Signing** | Pointer to [**ApplicationOIDCAllOfSigning**](ApplicationOIDCAllOfSigning.md) |  | [optional] 
**AudienceRestriction** | Pointer to **string** | The service provider ID. Defaults to &#x60;urn:federation:MicrosoftOnline&#x60;. | [optional] [default to "urn:federation:MicrosoftOnline"]
**DomainName** | **string** | The federated domain name (for example, the Azure custom domain). | 
**ReplyUrl** | **string** | The URL that the replying party (such as, Office365) uses to accept submissions of RequestSecurityTokenResponse messages that are a result of SSO requests. | 
**ApplyDefaultTheme** | **bool** | If &#x60;true&#x60;, applies the default theme to the self service application. | 
**EnableDefaultThemeFooter** | Pointer to **bool** | If &#x60;true&#x60;, shows the default theme footer on the self service application. Applies only if &#x60;applyDefaultTheme&#x60; is also &#x60;true&#x60;. | [optional] 

## Methods

### NewReadOneApplication200Response

`func NewReadOneApplication200Response(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, homePageUrl string, acsUrls []string, assertionDuration int32, idpSigning ApplicationWSFEDAllOfIdpSigning, spEntityId string, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, domainName string, replyUrl string, applyDefaultTheme bool, ) *ReadOneApplication200Response`

NewReadOneApplication200Response instantiates a new ReadOneApplication200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOneApplication200ResponseWithDefaults

`func NewReadOneApplication200ResponseWithDefaults() *ReadOneApplication200Response`

NewReadOneApplication200ResponseWithDefaults instantiates a new ReadOneApplication200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ReadOneApplication200Response) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReadOneApplication200Response) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReadOneApplication200Response) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReadOneApplication200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessControl

`func (o *ReadOneApplication200Response) GetAccessControl() ApplicationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ReadOneApplication200Response) GetAccessControlOk() (*ApplicationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ReadOneApplication200Response) SetAccessControl(v ApplicationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ReadOneApplication200Response) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReadOneApplication200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReadOneApplication200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReadOneApplication200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReadOneApplication200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ReadOneApplication200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReadOneApplication200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReadOneApplication200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReadOneApplication200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReadOneApplication200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReadOneApplication200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReadOneApplication200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *ReadOneApplication200Response) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReadOneApplication200Response) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReadOneApplication200Response) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ReadOneApplication200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHiddenFromAppPortal

`func (o *ReadOneApplication200Response) GetHiddenFromAppPortal() bool`

GetHiddenFromAppPortal returns the HiddenFromAppPortal field if non-nil, zero value otherwise.

### GetHiddenFromAppPortalOk

`func (o *ReadOneApplication200Response) GetHiddenFromAppPortalOk() (*bool, bool)`

GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromAppPortal

`func (o *ReadOneApplication200Response) SetHiddenFromAppPortal(v bool)`

SetHiddenFromAppPortal sets HiddenFromAppPortal field to given value.

### HasHiddenFromAppPortal

`func (o *ReadOneApplication200Response) HasHiddenFromAppPortal() bool`

HasHiddenFromAppPortal returns a boolean if a field has been set.

### GetIcon

`func (o *ReadOneApplication200Response) GetIcon() ApplicationIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ReadOneApplication200Response) GetIconOk() (*ApplicationIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ReadOneApplication200Response) SetIcon(v ApplicationIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ReadOneApplication200Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *ReadOneApplication200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReadOneApplication200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReadOneApplication200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReadOneApplication200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginPageUrl

`func (o *ReadOneApplication200Response) GetLoginPageUrl() string`

GetLoginPageUrl returns the LoginPageUrl field if non-nil, zero value otherwise.

### GetLoginPageUrlOk

`func (o *ReadOneApplication200Response) GetLoginPageUrlOk() (*string, bool)`

GetLoginPageUrlOk returns a tuple with the LoginPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageUrl

`func (o *ReadOneApplication200Response) SetLoginPageUrl(v string)`

SetLoginPageUrl sets LoginPageUrl field to given value.

### HasLoginPageUrl

`func (o *ReadOneApplication200Response) HasLoginPageUrl() bool`

HasLoginPageUrl returns a boolean if a field has been set.

### GetName

`func (o *ReadOneApplication200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReadOneApplication200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReadOneApplication200Response) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *ReadOneApplication200Response) GetProtocol() EnumApplicationProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ReadOneApplication200Response) GetProtocolOk() (*EnumApplicationProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ReadOneApplication200Response) SetProtocol(v EnumApplicationProtocol)`

SetProtocol sets Protocol field to given value.


### GetType

`func (o *ReadOneApplication200Response) GetType() EnumApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReadOneApplication200Response) GetTypeOk() (*EnumApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReadOneApplication200Response) SetType(v EnumApplicationType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ReadOneApplication200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReadOneApplication200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReadOneApplication200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReadOneApplication200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHomePageUrl

`func (o *ReadOneApplication200Response) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ReadOneApplication200Response) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ReadOneApplication200Response) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.


### GetAcsUrls

`func (o *ReadOneApplication200Response) GetAcsUrls() []string`

GetAcsUrls returns the AcsUrls field if non-nil, zero value otherwise.

### GetAcsUrlsOk

`func (o *ReadOneApplication200Response) GetAcsUrlsOk() (*[]string, bool)`

GetAcsUrlsOk returns a tuple with the AcsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrls

`func (o *ReadOneApplication200Response) SetAcsUrls(v []string)`

SetAcsUrls sets AcsUrls field to given value.


### GetAssertionDuration

`func (o *ReadOneApplication200Response) GetAssertionDuration() int32`

GetAssertionDuration returns the AssertionDuration field if non-nil, zero value otherwise.

### GetAssertionDurationOk

`func (o *ReadOneApplication200Response) GetAssertionDurationOk() (*int32, bool)`

GetAssertionDurationOk returns a tuple with the AssertionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionDuration

`func (o *ReadOneApplication200Response) SetAssertionDuration(v int32)`

SetAssertionDuration sets AssertionDuration field to given value.


### GetAssertionSigned

`func (o *ReadOneApplication200Response) GetAssertionSigned() bool`

GetAssertionSigned returns the AssertionSigned field if non-nil, zero value otherwise.

### GetAssertionSignedOk

`func (o *ReadOneApplication200Response) GetAssertionSignedOk() (*bool, bool)`

GetAssertionSignedOk returns a tuple with the AssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSigned

`func (o *ReadOneApplication200Response) SetAssertionSigned(v bool)`

SetAssertionSigned sets AssertionSigned field to given value.

### HasAssertionSigned

`func (o *ReadOneApplication200Response) HasAssertionSigned() bool`

HasAssertionSigned returns a boolean if a field has been set.

### GetCorsSettings

`func (o *ReadOneApplication200Response) GetCorsSettings() ApplicationCorsSettings`

GetCorsSettings returns the CorsSettings field if non-nil, zero value otherwise.

### GetCorsSettingsOk

`func (o *ReadOneApplication200Response) GetCorsSettingsOk() (*ApplicationCorsSettings, bool)`

GetCorsSettingsOk returns a tuple with the CorsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsSettings

`func (o *ReadOneApplication200Response) SetCorsSettings(v ApplicationCorsSettings)`

SetCorsSettings sets CorsSettings field to given value.

### HasCorsSettings

`func (o *ReadOneApplication200Response) HasCorsSettings() bool`

HasCorsSettings returns a boolean if a field has been set.

### GetDefaultTargetUrl

`func (o *ReadOneApplication200Response) GetDefaultTargetUrl() string`

GetDefaultTargetUrl returns the DefaultTargetUrl field if non-nil, zero value otherwise.

### GetDefaultTargetUrlOk

`func (o *ReadOneApplication200Response) GetDefaultTargetUrlOk() (*string, bool)`

GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetUrl

`func (o *ReadOneApplication200Response) SetDefaultTargetUrl(v string)`

SetDefaultTargetUrl sets DefaultTargetUrl field to given value.

### HasDefaultTargetUrl

`func (o *ReadOneApplication200Response) HasDefaultTargetUrl() bool`

HasDefaultTargetUrl returns a boolean if a field has been set.

### GetEnableRequestedAuthnContext

`func (o *ReadOneApplication200Response) GetEnableRequestedAuthnContext() bool`

GetEnableRequestedAuthnContext returns the EnableRequestedAuthnContext field if non-nil, zero value otherwise.

### GetEnableRequestedAuthnContextOk

`func (o *ReadOneApplication200Response) GetEnableRequestedAuthnContextOk() (*bool, bool)`

GetEnableRequestedAuthnContextOk returns a tuple with the EnableRequestedAuthnContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRequestedAuthnContext

`func (o *ReadOneApplication200Response) SetEnableRequestedAuthnContext(v bool)`

SetEnableRequestedAuthnContext sets EnableRequestedAuthnContext field to given value.

### HasEnableRequestedAuthnContext

`func (o *ReadOneApplication200Response) HasEnableRequestedAuthnContext() bool`

HasEnableRequestedAuthnContext returns a boolean if a field has been set.

### GetIdpSigning

`func (o *ReadOneApplication200Response) GetIdpSigning() ApplicationWSFEDAllOfIdpSigning`

GetIdpSigning returns the IdpSigning field if non-nil, zero value otherwise.

### GetIdpSigningOk

`func (o *ReadOneApplication200Response) GetIdpSigningOk() (*ApplicationWSFEDAllOfIdpSigning, bool)`

GetIdpSigningOk returns a tuple with the IdpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSigning

`func (o *ReadOneApplication200Response) SetIdpSigning(v ApplicationWSFEDAllOfIdpSigning)`

SetIdpSigning sets IdpSigning field to given value.


### GetNameIdFormat

`func (o *ReadOneApplication200Response) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *ReadOneApplication200Response) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *ReadOneApplication200Response) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *ReadOneApplication200Response) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetResponseSigned

`func (o *ReadOneApplication200Response) GetResponseSigned() bool`

GetResponseSigned returns the ResponseSigned field if non-nil, zero value otherwise.

### GetResponseSignedOk

`func (o *ReadOneApplication200Response) GetResponseSignedOk() (*bool, bool)`

GetResponseSignedOk returns a tuple with the ResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSigned

`func (o *ReadOneApplication200Response) SetResponseSigned(v bool)`

SetResponseSigned sets ResponseSigned field to given value.

### HasResponseSigned

`func (o *ReadOneApplication200Response) HasResponseSigned() bool`

HasResponseSigned returns a boolean if a field has been set.

### GetSloBinding

`func (o *ReadOneApplication200Response) GetSloBinding() EnumApplicationSAMLSloBinding`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *ReadOneApplication200Response) GetSloBindingOk() (*EnumApplicationSAMLSloBinding, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *ReadOneApplication200Response) SetSloBinding(v EnumApplicationSAMLSloBinding)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *ReadOneApplication200Response) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### GetSloEndpoint

`func (o *ReadOneApplication200Response) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *ReadOneApplication200Response) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *ReadOneApplication200Response) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *ReadOneApplication200Response) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetSloResponseEndpoint

`func (o *ReadOneApplication200Response) GetSloResponseEndpoint() string`

GetSloResponseEndpoint returns the SloResponseEndpoint field if non-nil, zero value otherwise.

### GetSloResponseEndpointOk

`func (o *ReadOneApplication200Response) GetSloResponseEndpointOk() (*string, bool)`

GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloResponseEndpoint

`func (o *ReadOneApplication200Response) SetSloResponseEndpoint(v string)`

SetSloResponseEndpoint sets SloResponseEndpoint field to given value.

### HasSloResponseEndpoint

`func (o *ReadOneApplication200Response) HasSloResponseEndpoint() bool`

HasSloResponseEndpoint returns a boolean if a field has been set.

### GetSloWindow

`func (o *ReadOneApplication200Response) GetSloWindow() int32`

GetSloWindow returns the SloWindow field if non-nil, zero value otherwise.

### GetSloWindowOk

`func (o *ReadOneApplication200Response) GetSloWindowOk() (*int32, bool)`

GetSloWindowOk returns a tuple with the SloWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloWindow

`func (o *ReadOneApplication200Response) SetSloWindow(v int32)`

SetSloWindow sets SloWindow field to given value.

### HasSloWindow

`func (o *ReadOneApplication200Response) HasSloWindow() bool`

HasSloWindow returns a boolean if a field has been set.

### GetSpEncryption

`func (o *ReadOneApplication200Response) GetSpEncryption() ApplicationSAMLAllOfSpEncryption`

GetSpEncryption returns the SpEncryption field if non-nil, zero value otherwise.

### GetSpEncryptionOk

`func (o *ReadOneApplication200Response) GetSpEncryptionOk() (*ApplicationSAMLAllOfSpEncryption, bool)`

GetSpEncryptionOk returns a tuple with the SpEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEncryption

`func (o *ReadOneApplication200Response) SetSpEncryption(v ApplicationSAMLAllOfSpEncryption)`

SetSpEncryption sets SpEncryption field to given value.

### HasSpEncryption

`func (o *ReadOneApplication200Response) HasSpEncryption() bool`

HasSpEncryption returns a boolean if a field has been set.

### GetSpEntityId

`func (o *ReadOneApplication200Response) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *ReadOneApplication200Response) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *ReadOneApplication200Response) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.


### GetSpVerification

`func (o *ReadOneApplication200Response) GetSpVerification() ApplicationSAMLAllOfSpVerification`

GetSpVerification returns the SpVerification field if non-nil, zero value otherwise.

### GetSpVerificationOk

`func (o *ReadOneApplication200Response) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool)`

GetSpVerificationOk returns a tuple with the SpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVerification

`func (o *ReadOneApplication200Response) SetSpVerification(v ApplicationSAMLAllOfSpVerification)`

SetSpVerification sets SpVerification field to given value.

### HasSpVerification

`func (o *ReadOneApplication200Response) HasSpVerification() bool`

HasSpVerification returns a boolean if a field has been set.

### GetTemplate

`func (o *ReadOneApplication200Response) GetTemplate() ApplicationTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ReadOneApplication200Response) GetTemplateOk() (*ApplicationTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ReadOneApplication200Response) SetTemplate(v ApplicationTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ReadOneApplication200Response) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ReadOneApplication200Response) GetAdditionalRefreshTokenReplayProtectionEnabled() bool`

GetAdditionalRefreshTokenReplayProtectionEnabled returns the AdditionalRefreshTokenReplayProtectionEnabled field if non-nil, zero value otherwise.

### GetAdditionalRefreshTokenReplayProtectionEnabledOk

`func (o *ReadOneApplication200Response) GetAdditionalRefreshTokenReplayProtectionEnabledOk() (*bool, bool)`

GetAdditionalRefreshTokenReplayProtectionEnabledOk returns a tuple with the AdditionalRefreshTokenReplayProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ReadOneApplication200Response) SetAdditionalRefreshTokenReplayProtectionEnabled(v bool)`

SetAdditionalRefreshTokenReplayProtectionEnabled sets AdditionalRefreshTokenReplayProtectionEnabled field to given value.

### HasAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ReadOneApplication200Response) HasAdditionalRefreshTokenReplayProtectionEnabled() bool`

HasAdditionalRefreshTokenReplayProtectionEnabled returns a boolean if a field has been set.

### GetAllowWildcardInRedirectUris

`func (o *ReadOneApplication200Response) GetAllowWildcardInRedirectUris() bool`

GetAllowWildcardInRedirectUris returns the AllowWildcardInRedirectUris field if non-nil, zero value otherwise.

### GetAllowWildcardInRedirectUrisOk

`func (o *ReadOneApplication200Response) GetAllowWildcardInRedirectUrisOk() (*bool, bool)`

GetAllowWildcardInRedirectUrisOk returns a tuple with the AllowWildcardInRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardInRedirectUris

`func (o *ReadOneApplication200Response) SetAllowWildcardInRedirectUris(v bool)`

SetAllowWildcardInRedirectUris sets AllowWildcardInRedirectUris field to given value.

### HasAllowWildcardInRedirectUris

`func (o *ReadOneApplication200Response) HasAllowWildcardInRedirectUris() bool`

HasAllowWildcardInRedirectUris returns a boolean if a field has been set.

### GetAssignActorRoles

`func (o *ReadOneApplication200Response) GetAssignActorRoles() bool`

GetAssignActorRoles returns the AssignActorRoles field if non-nil, zero value otherwise.

### GetAssignActorRolesOk

`func (o *ReadOneApplication200Response) GetAssignActorRolesOk() (*bool, bool)`

GetAssignActorRolesOk returns a tuple with the AssignActorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignActorRoles

`func (o *ReadOneApplication200Response) SetAssignActorRoles(v bool)`

SetAssignActorRoles sets AssignActorRoles field to given value.

### HasAssignActorRoles

`func (o *ReadOneApplication200Response) HasAssignActorRoles() bool`

HasAssignActorRoles returns a boolean if a field has been set.

### GetClientId

`func (o *ReadOneApplication200Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ReadOneApplication200Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ReadOneApplication200Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ReadOneApplication200Response) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ReadOneApplication200Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ReadOneApplication200Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ReadOneApplication200Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ReadOneApplication200Response) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetDevicePathId

`func (o *ReadOneApplication200Response) GetDevicePathId() string`

GetDevicePathId returns the DevicePathId field if non-nil, zero value otherwise.

### GetDevicePathIdOk

`func (o *ReadOneApplication200Response) GetDevicePathIdOk() (*string, bool)`

GetDevicePathIdOk returns a tuple with the DevicePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePathId

`func (o *ReadOneApplication200Response) SetDevicePathId(v string)`

SetDevicePathId sets DevicePathId field to given value.

### HasDevicePathId

`func (o *ReadOneApplication200Response) HasDevicePathId() bool`

HasDevicePathId returns a boolean if a field has been set.

### GetDeviceCustomVerificationUri

`func (o *ReadOneApplication200Response) GetDeviceCustomVerificationUri() string`

GetDeviceCustomVerificationUri returns the DeviceCustomVerificationUri field if non-nil, zero value otherwise.

### GetDeviceCustomVerificationUriOk

`func (o *ReadOneApplication200Response) GetDeviceCustomVerificationUriOk() (*string, bool)`

GetDeviceCustomVerificationUriOk returns a tuple with the DeviceCustomVerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCustomVerificationUri

`func (o *ReadOneApplication200Response) SetDeviceCustomVerificationUri(v string)`

SetDeviceCustomVerificationUri sets DeviceCustomVerificationUri field to given value.

### HasDeviceCustomVerificationUri

`func (o *ReadOneApplication200Response) HasDeviceCustomVerificationUri() bool`

HasDeviceCustomVerificationUri returns a boolean if a field has been set.

### GetDeviceTimeout

`func (o *ReadOneApplication200Response) GetDeviceTimeout() int32`

GetDeviceTimeout returns the DeviceTimeout field if non-nil, zero value otherwise.

### GetDeviceTimeoutOk

`func (o *ReadOneApplication200Response) GetDeviceTimeoutOk() (*int32, bool)`

GetDeviceTimeoutOk returns a tuple with the DeviceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTimeout

`func (o *ReadOneApplication200Response) SetDeviceTimeout(v int32)`

SetDeviceTimeout sets DeviceTimeout field to given value.

### HasDeviceTimeout

`func (o *ReadOneApplication200Response) HasDeviceTimeout() bool`

HasDeviceTimeout returns a boolean if a field has been set.

### GetDevicePollingInterval

`func (o *ReadOneApplication200Response) GetDevicePollingInterval() int32`

GetDevicePollingInterval returns the DevicePollingInterval field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOk

`func (o *ReadOneApplication200Response) GetDevicePollingIntervalOk() (*int32, bool)`

GetDevicePollingIntervalOk returns a tuple with the DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingInterval

`func (o *ReadOneApplication200Response) SetDevicePollingInterval(v int32)`

SetDevicePollingInterval sets DevicePollingInterval field to given value.

### HasDevicePollingInterval

`func (o *ReadOneApplication200Response) HasDevicePollingInterval() bool`

HasDevicePollingInterval returns a boolean if a field has been set.

### GetJwks

`func (o *ReadOneApplication200Response) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *ReadOneApplication200Response) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *ReadOneApplication200Response) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *ReadOneApplication200Response) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetJwksUrl

`func (o *ReadOneApplication200Response) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *ReadOneApplication200Response) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *ReadOneApplication200Response) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *ReadOneApplication200Response) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetMobile

`func (o *ReadOneApplication200Response) GetMobile() ApplicationOIDCAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ReadOneApplication200Response) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ReadOneApplication200Response) SetMobile(v ApplicationOIDCAllOfMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ReadOneApplication200Response) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetBundleId

`func (o *ReadOneApplication200Response) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ReadOneApplication200Response) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ReadOneApplication200Response) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ReadOneApplication200Response) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *ReadOneApplication200Response) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ReadOneApplication200Response) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ReadOneApplication200Response) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ReadOneApplication200Response) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetKerberos

`func (o *ReadOneApplication200Response) GetKerberos() ApplicationWSFEDAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *ReadOneApplication200Response) GetKerberosOk() (*ApplicationWSFEDAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *ReadOneApplication200Response) SetKerberos(v ApplicationWSFEDAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *ReadOneApplication200Response) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ReadOneApplication200Response) GetGrantTypes() []EnumApplicationOIDCGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ReadOneApplication200Response) GetGrantTypesOk() (*[]EnumApplicationOIDCGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ReadOneApplication200Response) SetGrantTypes(v []EnumApplicationOIDCGrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetInitiateLoginUri

`func (o *ReadOneApplication200Response) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *ReadOneApplication200Response) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *ReadOneApplication200Response) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *ReadOneApplication200Response) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetPkceEnforcement

`func (o *ReadOneApplication200Response) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *ReadOneApplication200Response) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *ReadOneApplication200Response) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *ReadOneApplication200Response) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *ReadOneApplication200Response) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *ReadOneApplication200Response) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *ReadOneApplication200Response) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *ReadOneApplication200Response) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ReadOneApplication200Response) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ReadOneApplication200Response) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ReadOneApplication200Response) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ReadOneApplication200Response) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *ReadOneApplication200Response) GetRefreshTokenDuration() int32`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *ReadOneApplication200Response) GetRefreshTokenDurationOk() (*int32, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *ReadOneApplication200Response) SetRefreshTokenDuration(v int32)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *ReadOneApplication200Response) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingDuration

`func (o *ReadOneApplication200Response) GetRefreshTokenRollingDuration() int32`

GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingDurationOk

`func (o *ReadOneApplication200Response) GetRefreshTokenRollingDurationOk() (*int32, bool)`

GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingDuration

`func (o *ReadOneApplication200Response) SetRefreshTokenRollingDuration(v int32)`

SetRefreshTokenRollingDuration sets RefreshTokenRollingDuration field to given value.

### HasRefreshTokenRollingDuration

`func (o *ReadOneApplication200Response) HasRefreshTokenRollingDuration() bool`

HasRefreshTokenRollingDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodDuration

`func (o *ReadOneApplication200Response) GetRefreshTokenRollingGracePeriodDuration() int32`

GetRefreshTokenRollingGracePeriodDuration returns the RefreshTokenRollingGracePeriodDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodDurationOk

`func (o *ReadOneApplication200Response) GetRefreshTokenRollingGracePeriodDurationOk() (*int32, bool)`

GetRefreshTokenRollingGracePeriodDurationOk returns a tuple with the RefreshTokenRollingGracePeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodDuration

`func (o *ReadOneApplication200Response) SetRefreshTokenRollingGracePeriodDuration(v int32)`

SetRefreshTokenRollingGracePeriodDuration sets RefreshTokenRollingGracePeriodDuration field to given value.

### HasRefreshTokenRollingGracePeriodDuration

`func (o *ReadOneApplication200Response) HasRefreshTokenRollingGracePeriodDuration() bool`

HasRefreshTokenRollingGracePeriodDuration returns a boolean if a field has been set.

### GetResponseTypes

`func (o *ReadOneApplication200Response) GetResponseTypes() []EnumApplicationOIDCResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *ReadOneApplication200Response) GetResponseTypesOk() (*[]EnumApplicationOIDCResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *ReadOneApplication200Response) SetResponseTypes(v []EnumApplicationOIDCResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *ReadOneApplication200Response) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRequireSignedRequestObject

`func (o *ReadOneApplication200Response) GetRequireSignedRequestObject() bool`

GetRequireSignedRequestObject returns the RequireSignedRequestObject field if non-nil, zero value otherwise.

### GetRequireSignedRequestObjectOk

`func (o *ReadOneApplication200Response) GetRequireSignedRequestObjectOk() (*bool, bool)`

GetRequireSignedRequestObjectOk returns a tuple with the RequireSignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequestObject

`func (o *ReadOneApplication200Response) SetRequireSignedRequestObject(v bool)`

SetRequireSignedRequestObject sets RequireSignedRequestObject field to given value.

### HasRequireSignedRequestObject

`func (o *ReadOneApplication200Response) HasRequireSignedRequestObject() bool`

HasRequireSignedRequestObject returns a boolean if a field has been set.

### GetSupportUnsignedRequestObject

`func (o *ReadOneApplication200Response) GetSupportUnsignedRequestObject() bool`

GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field if non-nil, zero value otherwise.

### GetSupportUnsignedRequestObjectOk

`func (o *ReadOneApplication200Response) GetSupportUnsignedRequestObjectOk() (*bool, bool)`

GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUnsignedRequestObject

`func (o *ReadOneApplication200Response) SetSupportUnsignedRequestObject(v bool)`

SetSupportUnsignedRequestObject sets SupportUnsignedRequestObject field to given value.

### HasSupportUnsignedRequestObject

`func (o *ReadOneApplication200Response) HasSupportUnsignedRequestObject() bool`

HasSupportUnsignedRequestObject returns a boolean if a field has been set.

### GetTags

`func (o *ReadOneApplication200Response) GetTags() []EnumApplicationTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReadOneApplication200Response) GetTagsOk() (*[]EnumApplicationTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReadOneApplication200Response) SetTags(v []EnumApplicationTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReadOneApplication200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetLinkUri

`func (o *ReadOneApplication200Response) GetTargetLinkUri() string`

GetTargetLinkUri returns the TargetLinkUri field if non-nil, zero value otherwise.

### GetTargetLinkUriOk

`func (o *ReadOneApplication200Response) GetTargetLinkUriOk() (*string, bool)`

GetTargetLinkUriOk returns a tuple with the TargetLinkUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLinkUri

`func (o *ReadOneApplication200Response) SetTargetLinkUri(v string)`

SetTargetLinkUri sets TargetLinkUri field to given value.

### HasTargetLinkUri

`func (o *ReadOneApplication200Response) HasTargetLinkUri() bool`

HasTargetLinkUri returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ReadOneApplication200Response) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ReadOneApplication200Response) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ReadOneApplication200Response) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetParRequirement

`func (o *ReadOneApplication200Response) GetParRequirement() EnumApplicationOIDCPARRequirement`

GetParRequirement returns the ParRequirement field if non-nil, zero value otherwise.

### GetParRequirementOk

`func (o *ReadOneApplication200Response) GetParRequirementOk() (*EnumApplicationOIDCPARRequirement, bool)`

GetParRequirementOk returns a tuple with the ParRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequirement

`func (o *ReadOneApplication200Response) SetParRequirement(v EnumApplicationOIDCPARRequirement)`

SetParRequirement sets ParRequirement field to given value.

### HasParRequirement

`func (o *ReadOneApplication200Response) HasParRequirement() bool`

HasParRequirement returns a boolean if a field has been set.

### GetParTimeout

`func (o *ReadOneApplication200Response) GetParTimeout() int32`

GetParTimeout returns the ParTimeout field if non-nil, zero value otherwise.

### GetParTimeoutOk

`func (o *ReadOneApplication200Response) GetParTimeoutOk() (*int32, bool)`

GetParTimeoutOk returns a tuple with the ParTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParTimeout

`func (o *ReadOneApplication200Response) SetParTimeout(v int32)`

SetParTimeout sets ParTimeout field to given value.

### HasParTimeout

`func (o *ReadOneApplication200Response) HasParTimeout() bool`

HasParTimeout returns a boolean if a field has been set.

### GetSigning

`func (o *ReadOneApplication200Response) GetSigning() ApplicationOIDCAllOfSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *ReadOneApplication200Response) GetSigningOk() (*ApplicationOIDCAllOfSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *ReadOneApplication200Response) SetSigning(v ApplicationOIDCAllOfSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *ReadOneApplication200Response) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetAudienceRestriction

`func (o *ReadOneApplication200Response) GetAudienceRestriction() string`

GetAudienceRestriction returns the AudienceRestriction field if non-nil, zero value otherwise.

### GetAudienceRestrictionOk

`func (o *ReadOneApplication200Response) GetAudienceRestrictionOk() (*string, bool)`

GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceRestriction

`func (o *ReadOneApplication200Response) SetAudienceRestriction(v string)`

SetAudienceRestriction sets AudienceRestriction field to given value.

### HasAudienceRestriction

`func (o *ReadOneApplication200Response) HasAudienceRestriction() bool`

HasAudienceRestriction returns a boolean if a field has been set.

### GetDomainName

`func (o *ReadOneApplication200Response) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ReadOneApplication200Response) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ReadOneApplication200Response) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetReplyUrl

`func (o *ReadOneApplication200Response) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *ReadOneApplication200Response) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *ReadOneApplication200Response) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.


### GetApplyDefaultTheme

`func (o *ReadOneApplication200Response) GetApplyDefaultTheme() bool`

GetApplyDefaultTheme returns the ApplyDefaultTheme field if non-nil, zero value otherwise.

### GetApplyDefaultThemeOk

`func (o *ReadOneApplication200Response) GetApplyDefaultThemeOk() (*bool, bool)`

GetApplyDefaultThemeOk returns a tuple with the ApplyDefaultTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyDefaultTheme

`func (o *ReadOneApplication200Response) SetApplyDefaultTheme(v bool)`

SetApplyDefaultTheme sets ApplyDefaultTheme field to given value.


### GetEnableDefaultThemeFooter

`func (o *ReadOneApplication200Response) GetEnableDefaultThemeFooter() bool`

GetEnableDefaultThemeFooter returns the EnableDefaultThemeFooter field if non-nil, zero value otherwise.

### GetEnableDefaultThemeFooterOk

`func (o *ReadOneApplication200Response) GetEnableDefaultThemeFooterOk() (*bool, bool)`

GetEnableDefaultThemeFooterOk returns a tuple with the EnableDefaultThemeFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDefaultThemeFooter

`func (o *ReadOneApplication200Response) SetEnableDefaultThemeFooter(v bool)`

SetEnableDefaultThemeFooter sets EnableDefaultThemeFooter field to given value.

### HasEnableDefaultThemeFooter

`func (o *ReadOneApplication200Response) HasEnableDefaultThemeFooter() bool`

HasEnableDefaultThemeFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


