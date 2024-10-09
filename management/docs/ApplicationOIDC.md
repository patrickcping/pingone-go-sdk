# ApplicationOIDC

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
**AdditionalRefreshTokenReplayProtectionEnabled** | Pointer to **bool** | When set to &#x60;true&#x60; (the default), if you attempt to reuse the refresh token, the authorization server immediately revokes the reused refresh token, as well as all descendant tokens. Setting this to null equates to a &#x60;false&#x60; setting. | [optional] [default to true]
**AllowWildcardInRedirectUris** | Pointer to **bool** | A boolean to specify whether wildcards are allowed in redirect URIs. For more information, see [Wildcards in Redirect URIs](https://docs.pingidentity.com/csh?context&#x3D;p1_c_wildcard_redirect_uri). | [optional] 
**AssignActorRoles** | Pointer to **bool** | A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request. | [optional] 
**ClientId** | Pointer to **string** | (Required when &#x60;clientSecret&#x60; is specified.) Supported only for &#x60;POST&#x60; operations. This is the UUID of an external application that is being migrated to PingOne. The UUID must be a minimum of 8 alpha-numeric characters, and must be globally unique in PingOne. | [optional] 
**ClientSecret** | Pointer to **string** | (Required when &#x60;clientId&#x60; is specified.) Supported only for &#x60;POST&#x60; operations. This is the client secret associated with &#x60;clientId&#x60; for an external application that is being migrated to PingOne. This must be a minimum of 8 alpha-numeric characters. | [optional] 
**CorsSettings** | Pointer to [**ApplicationCorsSettings**](ApplicationCorsSettings.md) |  | [optional] 
**DevicePathId** | Pointer to **string** | A string that specifies a unique identifier within an environment for a device authorization grant flow to provide a short identifier to the application. This property is ignored when the &#x60;deviceCustomVerificationUri&#x60; property is configured. The string can contain any letters, numbers, and some special characters (regex &#x60;a-zA-Z0-9_-&#x60;). It can have a length of no more than 50 characters (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;50&#x60;). | [optional] 
**DeviceCustomVerificationUri** | Pointer to **string** | A string that specifies an optional custom verification URI that is returned for the &#x60;/device_authorization&#x60; endpoint. | [optional] 
**DeviceTimeout** | Pointer to **int32** | An integer that specifies the length of time (in seconds) that the &#x60;userCode&#x60; and &#x60;deviceCode&#x60; returned by the &#x60;/device_authorization&#x60; endpoint are valid. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;600&#x60; seconds. It can have a value of no more than &#x60;3600&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;3600&#x60;). | [optional] [default to 600]
**DevicePollingInterval** | Pointer to **int32** | An integer that specifies the frequency (in seconds) for the client to poll the &#x60;/as/token&#x60; endpoint. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;5&#x60; seconds. It can have a value of no more than &#x60;60&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;60&#x60;). | [optional] [default to 5]
**Jwks** | Pointer to **string** | A JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwksUrl&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwksUrl&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**JwksUrl** | Pointer to **string** | A URL (supports &#x60;https://&#x60; only) that provides access to a JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwks&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwks&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**Mobile** | Pointer to [**ApplicationOIDCAllOfMobile**](ApplicationOIDCAllOfMobile.md) |  | [optional] 
**BundleId** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.bundleId&#x60; instead.  A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  | [optional] 
**PackageName** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.packageName&#x60; instead.  A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  | [optional] 
**Kerberos** | Pointer to [**ApplicationOIDCAllOfKerberos**](ApplicationOIDCAllOfKerberos.md) |  | [optional] 
**GrantTypes** | [**[]EnumApplicationOIDCGrantType**](EnumApplicationOIDCGrantType.md) | A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS. | 
**HomePageUrl** | Pointer to **string** | A string that specifies the custom home page URL for the application. | [optional] 
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
**Template** | Pointer to [**ApplicationTemplate**](ApplicationTemplate.md) |  | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 
**ParRequirement** | Pointer to [**EnumApplicationOIDCPARRequirement**](EnumApplicationOIDCPARRequirement.md) |  | [optional] [default to ENUMAPPLICATIONOIDCPARREQUIREMENT_OPTIONAL]
**ParTimeout** | Pointer to **int32** | PAR timeout in seconds. Must be between &#x60;1&#x60; and &#x60;600&#x60;. The default value is &#x60;60&#x60;. | [optional] [default to 60]
**Signing** | Pointer to [**ApplicationOIDCAllOfSigning**](ApplicationOIDCAllOfSigning.md) |  | [optional] 

## Methods

### NewApplicationOIDC

`func NewApplicationOIDC(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, ) *ApplicationOIDC`

NewApplicationOIDC instantiates a new ApplicationOIDC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCWithDefaults

`func NewApplicationOIDCWithDefaults() *ApplicationOIDC`

NewApplicationOIDCWithDefaults instantiates a new ApplicationOIDC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationOIDC) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationOIDC) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationOIDC) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationOIDC) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessControl

`func (o *ApplicationOIDC) GetAccessControl() ApplicationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ApplicationOIDC) GetAccessControlOk() (*ApplicationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ApplicationOIDC) SetAccessControl(v ApplicationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ApplicationOIDC) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationOIDC) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationOIDC) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationOIDC) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationOIDC) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationOIDC) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationOIDC) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationOIDC) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationOIDC) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationOIDC) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationOIDC) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationOIDC) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *ApplicationOIDC) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationOIDC) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationOIDC) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationOIDC) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHiddenFromAppPortal

`func (o *ApplicationOIDC) GetHiddenFromAppPortal() bool`

GetHiddenFromAppPortal returns the HiddenFromAppPortal field if non-nil, zero value otherwise.

### GetHiddenFromAppPortalOk

`func (o *ApplicationOIDC) GetHiddenFromAppPortalOk() (*bool, bool)`

GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromAppPortal

`func (o *ApplicationOIDC) SetHiddenFromAppPortal(v bool)`

SetHiddenFromAppPortal sets HiddenFromAppPortal field to given value.

### HasHiddenFromAppPortal

`func (o *ApplicationOIDC) HasHiddenFromAppPortal() bool`

HasHiddenFromAppPortal returns a boolean if a field has been set.

### GetIcon

`func (o *ApplicationOIDC) GetIcon() ApplicationIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ApplicationOIDC) GetIconOk() (*ApplicationIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ApplicationOIDC) SetIcon(v ApplicationIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ApplicationOIDC) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *ApplicationOIDC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationOIDC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationOIDC) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationOIDC) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginPageUrl

`func (o *ApplicationOIDC) GetLoginPageUrl() string`

GetLoginPageUrl returns the LoginPageUrl field if non-nil, zero value otherwise.

### GetLoginPageUrlOk

`func (o *ApplicationOIDC) GetLoginPageUrlOk() (*string, bool)`

GetLoginPageUrlOk returns a tuple with the LoginPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageUrl

`func (o *ApplicationOIDC) SetLoginPageUrl(v string)`

SetLoginPageUrl sets LoginPageUrl field to given value.

### HasLoginPageUrl

`func (o *ApplicationOIDC) HasLoginPageUrl() bool`

HasLoginPageUrl returns a boolean if a field has been set.

### GetName

`func (o *ApplicationOIDC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationOIDC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationOIDC) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *ApplicationOIDC) GetProtocol() EnumApplicationProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationOIDC) GetProtocolOk() (*EnumApplicationProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationOIDC) SetProtocol(v EnumApplicationProtocol)`

SetProtocol sets Protocol field to given value.


### GetType

`func (o *ApplicationOIDC) GetType() EnumApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationOIDC) GetTypeOk() (*EnumApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationOIDC) SetType(v EnumApplicationType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ApplicationOIDC) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationOIDC) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationOIDC) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationOIDC) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDC) GetAdditionalRefreshTokenReplayProtectionEnabled() bool`

GetAdditionalRefreshTokenReplayProtectionEnabled returns the AdditionalRefreshTokenReplayProtectionEnabled field if non-nil, zero value otherwise.

### GetAdditionalRefreshTokenReplayProtectionEnabledOk

`func (o *ApplicationOIDC) GetAdditionalRefreshTokenReplayProtectionEnabledOk() (*bool, bool)`

GetAdditionalRefreshTokenReplayProtectionEnabledOk returns a tuple with the AdditionalRefreshTokenReplayProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDC) SetAdditionalRefreshTokenReplayProtectionEnabled(v bool)`

SetAdditionalRefreshTokenReplayProtectionEnabled sets AdditionalRefreshTokenReplayProtectionEnabled field to given value.

### HasAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDC) HasAdditionalRefreshTokenReplayProtectionEnabled() bool`

HasAdditionalRefreshTokenReplayProtectionEnabled returns a boolean if a field has been set.

### GetAllowWildcardInRedirectUris

`func (o *ApplicationOIDC) GetAllowWildcardInRedirectUris() bool`

GetAllowWildcardInRedirectUris returns the AllowWildcardInRedirectUris field if non-nil, zero value otherwise.

### GetAllowWildcardInRedirectUrisOk

`func (o *ApplicationOIDC) GetAllowWildcardInRedirectUrisOk() (*bool, bool)`

GetAllowWildcardInRedirectUrisOk returns a tuple with the AllowWildcardInRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardInRedirectUris

`func (o *ApplicationOIDC) SetAllowWildcardInRedirectUris(v bool)`

SetAllowWildcardInRedirectUris sets AllowWildcardInRedirectUris field to given value.

### HasAllowWildcardInRedirectUris

`func (o *ApplicationOIDC) HasAllowWildcardInRedirectUris() bool`

HasAllowWildcardInRedirectUris returns a boolean if a field has been set.

### GetAssignActorRoles

`func (o *ApplicationOIDC) GetAssignActorRoles() bool`

GetAssignActorRoles returns the AssignActorRoles field if non-nil, zero value otherwise.

### GetAssignActorRolesOk

`func (o *ApplicationOIDC) GetAssignActorRolesOk() (*bool, bool)`

GetAssignActorRolesOk returns a tuple with the AssignActorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignActorRoles

`func (o *ApplicationOIDC) SetAssignActorRoles(v bool)`

SetAssignActorRoles sets AssignActorRoles field to given value.

### HasAssignActorRoles

`func (o *ApplicationOIDC) HasAssignActorRoles() bool`

HasAssignActorRoles returns a boolean if a field has been set.

### GetClientId

`func (o *ApplicationOIDC) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApplicationOIDC) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApplicationOIDC) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApplicationOIDC) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ApplicationOIDC) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApplicationOIDC) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApplicationOIDC) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ApplicationOIDC) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCorsSettings

`func (o *ApplicationOIDC) GetCorsSettings() ApplicationCorsSettings`

GetCorsSettings returns the CorsSettings field if non-nil, zero value otherwise.

### GetCorsSettingsOk

`func (o *ApplicationOIDC) GetCorsSettingsOk() (*ApplicationCorsSettings, bool)`

GetCorsSettingsOk returns a tuple with the CorsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsSettings

`func (o *ApplicationOIDC) SetCorsSettings(v ApplicationCorsSettings)`

SetCorsSettings sets CorsSettings field to given value.

### HasCorsSettings

`func (o *ApplicationOIDC) HasCorsSettings() bool`

HasCorsSettings returns a boolean if a field has been set.

### GetDevicePathId

`func (o *ApplicationOIDC) GetDevicePathId() string`

GetDevicePathId returns the DevicePathId field if non-nil, zero value otherwise.

### GetDevicePathIdOk

`func (o *ApplicationOIDC) GetDevicePathIdOk() (*string, bool)`

GetDevicePathIdOk returns a tuple with the DevicePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePathId

`func (o *ApplicationOIDC) SetDevicePathId(v string)`

SetDevicePathId sets DevicePathId field to given value.

### HasDevicePathId

`func (o *ApplicationOIDC) HasDevicePathId() bool`

HasDevicePathId returns a boolean if a field has been set.

### GetDeviceCustomVerificationUri

`func (o *ApplicationOIDC) GetDeviceCustomVerificationUri() string`

GetDeviceCustomVerificationUri returns the DeviceCustomVerificationUri field if non-nil, zero value otherwise.

### GetDeviceCustomVerificationUriOk

`func (o *ApplicationOIDC) GetDeviceCustomVerificationUriOk() (*string, bool)`

GetDeviceCustomVerificationUriOk returns a tuple with the DeviceCustomVerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCustomVerificationUri

`func (o *ApplicationOIDC) SetDeviceCustomVerificationUri(v string)`

SetDeviceCustomVerificationUri sets DeviceCustomVerificationUri field to given value.

### HasDeviceCustomVerificationUri

`func (o *ApplicationOIDC) HasDeviceCustomVerificationUri() bool`

HasDeviceCustomVerificationUri returns a boolean if a field has been set.

### GetDeviceTimeout

`func (o *ApplicationOIDC) GetDeviceTimeout() int32`

GetDeviceTimeout returns the DeviceTimeout field if non-nil, zero value otherwise.

### GetDeviceTimeoutOk

`func (o *ApplicationOIDC) GetDeviceTimeoutOk() (*int32, bool)`

GetDeviceTimeoutOk returns a tuple with the DeviceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTimeout

`func (o *ApplicationOIDC) SetDeviceTimeout(v int32)`

SetDeviceTimeout sets DeviceTimeout field to given value.

### HasDeviceTimeout

`func (o *ApplicationOIDC) HasDeviceTimeout() bool`

HasDeviceTimeout returns a boolean if a field has been set.

### GetDevicePollingInterval

`func (o *ApplicationOIDC) GetDevicePollingInterval() int32`

GetDevicePollingInterval returns the DevicePollingInterval field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOk

`func (o *ApplicationOIDC) GetDevicePollingIntervalOk() (*int32, bool)`

GetDevicePollingIntervalOk returns a tuple with the DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingInterval

`func (o *ApplicationOIDC) SetDevicePollingInterval(v int32)`

SetDevicePollingInterval sets DevicePollingInterval field to given value.

### HasDevicePollingInterval

`func (o *ApplicationOIDC) HasDevicePollingInterval() bool`

HasDevicePollingInterval returns a boolean if a field has been set.

### GetJwks

`func (o *ApplicationOIDC) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *ApplicationOIDC) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *ApplicationOIDC) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *ApplicationOIDC) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetJwksUrl

`func (o *ApplicationOIDC) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *ApplicationOIDC) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *ApplicationOIDC) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *ApplicationOIDC) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetMobile

`func (o *ApplicationOIDC) GetMobile() ApplicationOIDCAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ApplicationOIDC) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ApplicationOIDC) SetMobile(v ApplicationOIDCAllOfMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ApplicationOIDC) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetBundleId

`func (o *ApplicationOIDC) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ApplicationOIDC) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ApplicationOIDC) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ApplicationOIDC) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *ApplicationOIDC) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ApplicationOIDC) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ApplicationOIDC) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ApplicationOIDC) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetKerberos

`func (o *ApplicationOIDC) GetKerberos() ApplicationOIDCAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *ApplicationOIDC) GetKerberosOk() (*ApplicationOIDCAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *ApplicationOIDC) SetKerberos(v ApplicationOIDCAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *ApplicationOIDC) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ApplicationOIDC) GetGrantTypes() []EnumApplicationOIDCGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ApplicationOIDC) GetGrantTypesOk() (*[]EnumApplicationOIDCGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ApplicationOIDC) SetGrantTypes(v []EnumApplicationOIDCGrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetHomePageUrl

`func (o *ApplicationOIDC) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ApplicationOIDC) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ApplicationOIDC) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *ApplicationOIDC) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### GetInitiateLoginUri

`func (o *ApplicationOIDC) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *ApplicationOIDC) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *ApplicationOIDC) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *ApplicationOIDC) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetPkceEnforcement

`func (o *ApplicationOIDC) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *ApplicationOIDC) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *ApplicationOIDC) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *ApplicationOIDC) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *ApplicationOIDC) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *ApplicationOIDC) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *ApplicationOIDC) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *ApplicationOIDC) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ApplicationOIDC) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ApplicationOIDC) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ApplicationOIDC) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ApplicationOIDC) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *ApplicationOIDC) GetRefreshTokenDuration() int32`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *ApplicationOIDC) GetRefreshTokenDurationOk() (*int32, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *ApplicationOIDC) SetRefreshTokenDuration(v int32)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *ApplicationOIDC) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingDuration

`func (o *ApplicationOIDC) GetRefreshTokenRollingDuration() int32`

GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingDurationOk

`func (o *ApplicationOIDC) GetRefreshTokenRollingDurationOk() (*int32, bool)`

GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingDuration

`func (o *ApplicationOIDC) SetRefreshTokenRollingDuration(v int32)`

SetRefreshTokenRollingDuration sets RefreshTokenRollingDuration field to given value.

### HasRefreshTokenRollingDuration

`func (o *ApplicationOIDC) HasRefreshTokenRollingDuration() bool`

HasRefreshTokenRollingDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDC) GetRefreshTokenRollingGracePeriodDuration() int32`

GetRefreshTokenRollingGracePeriodDuration returns the RefreshTokenRollingGracePeriodDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodDurationOk

`func (o *ApplicationOIDC) GetRefreshTokenRollingGracePeriodDurationOk() (*int32, bool)`

GetRefreshTokenRollingGracePeriodDurationOk returns a tuple with the RefreshTokenRollingGracePeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDC) SetRefreshTokenRollingGracePeriodDuration(v int32)`

SetRefreshTokenRollingGracePeriodDuration sets RefreshTokenRollingGracePeriodDuration field to given value.

### HasRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDC) HasRefreshTokenRollingGracePeriodDuration() bool`

HasRefreshTokenRollingGracePeriodDuration returns a boolean if a field has been set.

### GetResponseTypes

`func (o *ApplicationOIDC) GetResponseTypes() []EnumApplicationOIDCResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *ApplicationOIDC) GetResponseTypesOk() (*[]EnumApplicationOIDCResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *ApplicationOIDC) SetResponseTypes(v []EnumApplicationOIDCResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *ApplicationOIDC) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRequireSignedRequestObject

`func (o *ApplicationOIDC) GetRequireSignedRequestObject() bool`

GetRequireSignedRequestObject returns the RequireSignedRequestObject field if non-nil, zero value otherwise.

### GetRequireSignedRequestObjectOk

`func (o *ApplicationOIDC) GetRequireSignedRequestObjectOk() (*bool, bool)`

GetRequireSignedRequestObjectOk returns a tuple with the RequireSignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequestObject

`func (o *ApplicationOIDC) SetRequireSignedRequestObject(v bool)`

SetRequireSignedRequestObject sets RequireSignedRequestObject field to given value.

### HasRequireSignedRequestObject

`func (o *ApplicationOIDC) HasRequireSignedRequestObject() bool`

HasRequireSignedRequestObject returns a boolean if a field has been set.

### GetSupportUnsignedRequestObject

`func (o *ApplicationOIDC) GetSupportUnsignedRequestObject() bool`

GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field if non-nil, zero value otherwise.

### GetSupportUnsignedRequestObjectOk

`func (o *ApplicationOIDC) GetSupportUnsignedRequestObjectOk() (*bool, bool)`

GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUnsignedRequestObject

`func (o *ApplicationOIDC) SetSupportUnsignedRequestObject(v bool)`

SetSupportUnsignedRequestObject sets SupportUnsignedRequestObject field to given value.

### HasSupportUnsignedRequestObject

`func (o *ApplicationOIDC) HasSupportUnsignedRequestObject() bool`

HasSupportUnsignedRequestObject returns a boolean if a field has been set.

### GetTags

`func (o *ApplicationOIDC) GetTags() []EnumApplicationTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationOIDC) GetTagsOk() (*[]EnumApplicationTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationOIDC) SetTags(v []EnumApplicationTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationOIDC) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetLinkUri

`func (o *ApplicationOIDC) GetTargetLinkUri() string`

GetTargetLinkUri returns the TargetLinkUri field if non-nil, zero value otherwise.

### GetTargetLinkUriOk

`func (o *ApplicationOIDC) GetTargetLinkUriOk() (*string, bool)`

GetTargetLinkUriOk returns a tuple with the TargetLinkUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLinkUri

`func (o *ApplicationOIDC) SetTargetLinkUri(v string)`

SetTargetLinkUri sets TargetLinkUri field to given value.

### HasTargetLinkUri

`func (o *ApplicationOIDC) HasTargetLinkUri() bool`

HasTargetLinkUri returns a boolean if a field has been set.

### GetTemplate

`func (o *ApplicationOIDC) GetTemplate() ApplicationTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ApplicationOIDC) GetTemplateOk() (*ApplicationTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ApplicationOIDC) SetTemplate(v ApplicationTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ApplicationOIDC) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationOIDC) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationOIDC) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationOIDC) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetParRequirement

`func (o *ApplicationOIDC) GetParRequirement() EnumApplicationOIDCPARRequirement`

GetParRequirement returns the ParRequirement field if non-nil, zero value otherwise.

### GetParRequirementOk

`func (o *ApplicationOIDC) GetParRequirementOk() (*EnumApplicationOIDCPARRequirement, bool)`

GetParRequirementOk returns a tuple with the ParRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequirement

`func (o *ApplicationOIDC) SetParRequirement(v EnumApplicationOIDCPARRequirement)`

SetParRequirement sets ParRequirement field to given value.

### HasParRequirement

`func (o *ApplicationOIDC) HasParRequirement() bool`

HasParRequirement returns a boolean if a field has been set.

### GetParTimeout

`func (o *ApplicationOIDC) GetParTimeout() int32`

GetParTimeout returns the ParTimeout field if non-nil, zero value otherwise.

### GetParTimeoutOk

`func (o *ApplicationOIDC) GetParTimeoutOk() (*int32, bool)`

GetParTimeoutOk returns a tuple with the ParTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParTimeout

`func (o *ApplicationOIDC) SetParTimeout(v int32)`

SetParTimeout sets ParTimeout field to given value.

### HasParTimeout

`func (o *ApplicationOIDC) HasParTimeout() bool`

HasParTimeout returns a boolean if a field has been set.

### GetSigning

`func (o *ApplicationOIDC) GetSigning() ApplicationOIDCAllOfSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *ApplicationOIDC) GetSigningOk() (*ApplicationOIDCAllOfSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *ApplicationOIDC) SetSigning(v ApplicationOIDCAllOfSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *ApplicationOIDC) HasSigning() bool`

HasSigning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


