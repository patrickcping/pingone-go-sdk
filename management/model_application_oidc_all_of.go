/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the ApplicationOIDCAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOIDCAllOf{}

// ApplicationOIDCAllOf struct for ApplicationOIDCAllOf
type ApplicationOIDCAllOf struct {
	// A boolean to specify whether wildcards are allowed in redirect URIs. For more information, see [Wildcards in Redirect URIs](https://docs.pingidentity.com/csh?context=p1_c_wildcard_redirect_uri).
	AllowWildcardInRedirectUris *bool `json:"allowWildcardInRedirectUris,omitempty"`
	// A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request.
	AssignActorRoles *bool `json:"assignActorRoles,omitempty"`
	Mobile *ApplicationOIDCAllOfMobile `json:"mobile,omitempty"`
	// **Deprecation Notice** This field is deprecated and will be removed in a future release. Use `mobile.bundleId` instead.  A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable. 
	// Deprecated
	BundleId *string `json:"bundleId,omitempty"`
	// **Deprecation Notice** This field is deprecated and will be removed in a future release. Use `mobile.packageName` instead.  A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable. 
	// Deprecated
	PackageName *string `json:"packageName,omitempty"`
	Kerberos *ApplicationOIDCAllOfKerberos `json:"kerberos,omitempty"`
	// A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS.
	GrantTypes []EnumApplicationOIDCGrantType `json:"grantTypes"`
	// A string that specifies the custom home page URL for the application.
	HomePageUrl *string `json:"homePageUrl,omitempty"`
	// A string that specifies the URI to use for third-parties to begin the sign-on process for the application. If specified, PingOne redirects users to this URI to initiate SSO to PingOne. The application is responsible for implementing the relevant OIDC flow when the initiate login URI is requested. This property is required if you want the application to appear in the PingOne Application Portal. See the OIDC specification section of [Initiating Login from a Third Party](https://openid.net/specs/openid-connect-core-1_0.html#ThirdPartyInitiatedLogin) for more information.
	InitiateLoginUri *string `json:"initiateLoginUri,omitempty"`
	PkceEnforcement *EnumApplicationOIDCPKCEOption `json:"pkceEnforcement,omitempty"`
	// A string that specifies the URLs that the browser can be redirected to after logout.
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris,omitempty"`
	// A string that specifies the callback URI for the authentication response.
	RedirectUris []string `json:"redirectUris,omitempty"`
	// An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the `refreshTokenRollingDuration` property is specified for the application, then this property must be less than or equal to the value of `refreshTokenRollingDuration`. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenDuration *int32 `json:"refreshTokenDuration,omitempty"`
	// An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenRollingDuration *int32 `json:"refreshTokenRollingDuration,omitempty"`
	// The number of seconds that a refresh token may be reused after having been exchanged for a new set of tokens. This is useful in the case of network errors on the client. Valid values are between 0 and 86400 seconds. Null is treated the same as 0.
	RefreshTokenRollingGracePeriodDuration *int32 `json:"refreshTokenRollingGracePeriodDuration,omitempty"`
	// A string that specifies the code or token type returned by an authorization request. Options are TOKEN, ID_TOKEN, and CODE. Note that CODE cannot be used in an authorization request with TOKEN or ID_TOKEN because PingOne does not currently support OIDC hybrid flows.
	ResponseTypes []EnumApplicationOIDCResponseType `json:"responseTypes,omitempty"`
	// A boolean that specifies whether the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed.
	SupportUnsignedRequestObject *bool `json:"supportUnsignedRequestObject,omitempty"`
	// An array that specifies the list of labels associated with the application. Options are `PING_FED_CONNECTION_INTEGRATION`.  Only applicable for creating worker applications.
	Tags []EnumApplicationTags `json:"tags,omitempty"`
	// The URI for the application. If specified, PingOne will redirect application users to this URI after a user is authenticated. In the PingOne admin console, this becomes the value of the `target_link_uri` parameter used for the Initiate Single Sign-On URL field.
	TargetLinkUri *string `json:"targetLinkUri,omitempty"`
	TokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod `json:"tokenEndpointAuthMethod"`
}

// NewApplicationOIDCAllOf instantiates a new ApplicationOIDCAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOf(grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod) *ApplicationOIDCAllOf {
	this := ApplicationOIDCAllOf{}
	this.GrantTypes = grantTypes
	var refreshTokenDuration int32 = 2592000
	this.RefreshTokenDuration = &refreshTokenDuration
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	return &this
}

// NewApplicationOIDCAllOfWithDefaults instantiates a new ApplicationOIDCAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfWithDefaults() *ApplicationOIDCAllOf {
	this := ApplicationOIDCAllOf{}
	var refreshTokenDuration int32 = 2592000
	this.RefreshTokenDuration = &refreshTokenDuration
	return &this
}

// GetAllowWildcardInRedirectUris returns the AllowWildcardInRedirectUris field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetAllowWildcardInRedirectUris() bool {
	if o == nil || IsNil(o.AllowWildcardInRedirectUris) {
		var ret bool
		return ret
	}
	return *o.AllowWildcardInRedirectUris
}

// GetAllowWildcardInRedirectUrisOk returns a tuple with the AllowWildcardInRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetAllowWildcardInRedirectUrisOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowWildcardInRedirectUris) {
		return nil, false
	}
	return o.AllowWildcardInRedirectUris, true
}

// HasAllowWildcardInRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasAllowWildcardInRedirectUris() bool {
	if o != nil && !IsNil(o.AllowWildcardInRedirectUris) {
		return true
	}

	return false
}

// SetAllowWildcardInRedirectUris gets a reference to the given bool and assigns it to the AllowWildcardInRedirectUris field.
func (o *ApplicationOIDCAllOf) SetAllowWildcardInRedirectUris(v bool) {
	o.AllowWildcardInRedirectUris = &v
}

// GetAssignActorRoles returns the AssignActorRoles field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetAssignActorRoles() bool {
	if o == nil || IsNil(o.AssignActorRoles) {
		var ret bool
		return ret
	}
	return *o.AssignActorRoles
}

// GetAssignActorRolesOk returns a tuple with the AssignActorRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetAssignActorRolesOk() (*bool, bool) {
	if o == nil || IsNil(o.AssignActorRoles) {
		return nil, false
	}
	return o.AssignActorRoles, true
}

// HasAssignActorRoles returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasAssignActorRoles() bool {
	if o != nil && !IsNil(o.AssignActorRoles) {
		return true
	}

	return false
}

// SetAssignActorRoles gets a reference to the given bool and assigns it to the AssignActorRoles field.
func (o *ApplicationOIDCAllOf) SetAssignActorRoles(v bool) {
	o.AssignActorRoles = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetMobile() ApplicationOIDCAllOfMobile {
	if o == nil || IsNil(o.Mobile) {
		var ret ApplicationOIDCAllOfMobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given ApplicationOIDCAllOfMobile and assigns it to the Mobile field.
func (o *ApplicationOIDCAllOf) SetMobile(v ApplicationOIDCAllOfMobile) {
	o.Mobile = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
// Deprecated
func (o *ApplicationOIDCAllOf) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationOIDCAllOf) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
// Deprecated
func (o *ApplicationOIDCAllOf) SetBundleId(v string) {
	o.BundleId = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
// Deprecated
func (o *ApplicationOIDCAllOf) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationOIDCAllOf) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
// Deprecated
func (o *ApplicationOIDCAllOf) SetPackageName(v string) {
	o.PackageName = &v
}

// GetKerberos returns the Kerberos field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetKerberos() ApplicationOIDCAllOfKerberos {
	if o == nil || IsNil(o.Kerberos) {
		var ret ApplicationOIDCAllOfKerberos
		return ret
	}
	return *o.Kerberos
}

// GetKerberosOk returns a tuple with the Kerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetKerberosOk() (*ApplicationOIDCAllOfKerberos, bool) {
	if o == nil || IsNil(o.Kerberos) {
		return nil, false
	}
	return o.Kerberos, true
}

// HasKerberos returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasKerberos() bool {
	if o != nil && !IsNil(o.Kerberos) {
		return true
	}

	return false
}

// SetKerberos gets a reference to the given ApplicationOIDCAllOfKerberos and assigns it to the Kerberos field.
func (o *ApplicationOIDCAllOf) SetKerberos(v ApplicationOIDCAllOfKerberos) {
	o.Kerberos = &v
}

// GetGrantTypes returns the GrantTypes field value
func (o *ApplicationOIDCAllOf) GetGrantTypes() []EnumApplicationOIDCGrantType {
	if o == nil {
		var ret []EnumApplicationOIDCGrantType
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetGrantTypesOk() ([]EnumApplicationOIDCGrantType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *ApplicationOIDCAllOf) SetGrantTypes(v []EnumApplicationOIDCGrantType) {
	o.GrantTypes = v
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetHomePageUrl() string {
	if o == nil || IsNil(o.HomePageUrl) {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetHomePageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomePageUrl) {
		return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasHomePageUrl() bool {
	if o != nil && !IsNil(o.HomePageUrl) {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *ApplicationOIDCAllOf) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetInitiateLoginUri returns the InitiateLoginUri field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetInitiateLoginUri() string {
	if o == nil || IsNil(o.InitiateLoginUri) {
		var ret string
		return ret
	}
	return *o.InitiateLoginUri
}

// GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetInitiateLoginUriOk() (*string, bool) {
	if o == nil || IsNil(o.InitiateLoginUri) {
		return nil, false
	}
	return o.InitiateLoginUri, true
}

// HasInitiateLoginUri returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasInitiateLoginUri() bool {
	if o != nil && !IsNil(o.InitiateLoginUri) {
		return true
	}

	return false
}

// SetInitiateLoginUri gets a reference to the given string and assigns it to the InitiateLoginUri field.
func (o *ApplicationOIDCAllOf) SetInitiateLoginUri(v string) {
	o.InitiateLoginUri = &v
}

// GetPkceEnforcement returns the PkceEnforcement field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetPkceEnforcement() EnumApplicationOIDCPKCEOption {
	if o == nil || IsNil(o.PkceEnforcement) {
		var ret EnumApplicationOIDCPKCEOption
		return ret
	}
	return *o.PkceEnforcement
}

// GetPkceEnforcementOk returns a tuple with the PkceEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool) {
	if o == nil || IsNil(o.PkceEnforcement) {
		return nil, false
	}
	return o.PkceEnforcement, true
}

// HasPkceEnforcement returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPkceEnforcement() bool {
	if o != nil && !IsNil(o.PkceEnforcement) {
		return true
	}

	return false
}

// SetPkceEnforcement gets a reference to the given EnumApplicationOIDCPKCEOption and assigns it to the PkceEnforcement field.
func (o *ApplicationOIDCAllOf) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption) {
	o.PkceEnforcement = &v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetPostLogoutRedirectUris() []string {
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPostLogoutRedirectUris() bool {
	if o != nil && !IsNil(o.PostLogoutRedirectUris) {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *ApplicationOIDCAllOf) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ApplicationOIDCAllOf) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetRefreshTokenDuration() int32 {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		var ret int32
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRefreshTokenDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRefreshTokenDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int32 and assigns it to the RefreshTokenDuration field.
func (o *ApplicationOIDCAllOf) SetRefreshTokenDuration(v int32) {
	o.RefreshTokenDuration = &v
}

// GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingDuration() int32 {
	if o == nil || IsNil(o.RefreshTokenRollingDuration) {
		var ret int32
		return ret
	}
	return *o.RefreshTokenRollingDuration
}

// GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshTokenRollingDuration) {
		return nil, false
	}
	return o.RefreshTokenRollingDuration, true
}

// HasRefreshTokenRollingDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRefreshTokenRollingDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenRollingDuration) {
		return true
	}

	return false
}

// SetRefreshTokenRollingDuration gets a reference to the given int32 and assigns it to the RefreshTokenRollingDuration field.
func (o *ApplicationOIDCAllOf) SetRefreshTokenRollingDuration(v int32) {
	o.RefreshTokenRollingDuration = &v
}

// GetRefreshTokenRollingGracePeriodDuration returns the RefreshTokenRollingGracePeriodDuration field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingGracePeriodDuration() int32 {
	if o == nil || IsNil(o.RefreshTokenRollingGracePeriodDuration) {
		var ret int32
		return ret
	}
	return *o.RefreshTokenRollingGracePeriodDuration
}

// GetRefreshTokenRollingGracePeriodDurationOk returns a tuple with the RefreshTokenRollingGracePeriodDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingGracePeriodDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshTokenRollingGracePeriodDuration) {
		return nil, false
	}
	return o.RefreshTokenRollingGracePeriodDuration, true
}

// HasRefreshTokenRollingGracePeriodDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRefreshTokenRollingGracePeriodDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenRollingGracePeriodDuration) {
		return true
	}

	return false
}

// SetRefreshTokenRollingGracePeriodDuration gets a reference to the given int32 and assigns it to the RefreshTokenRollingGracePeriodDuration field.
func (o *ApplicationOIDCAllOf) SetRefreshTokenRollingGracePeriodDuration(v int32) {
	o.RefreshTokenRollingGracePeriodDuration = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetResponseTypes() []EnumApplicationOIDCResponseType {
	if o == nil || IsNil(o.ResponseTypes) {
		var ret []EnumApplicationOIDCResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetResponseTypesOk() ([]EnumApplicationOIDCResponseType, bool) {
	if o == nil || IsNil(o.ResponseTypes) {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasResponseTypes() bool {
	if o != nil && !IsNil(o.ResponseTypes) {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []EnumApplicationOIDCResponseType and assigns it to the ResponseTypes field.
func (o *ApplicationOIDCAllOf) SetResponseTypes(v []EnumApplicationOIDCResponseType) {
	o.ResponseTypes = v
}

// GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetSupportUnsignedRequestObject() bool {
	if o == nil || IsNil(o.SupportUnsignedRequestObject) {
		var ret bool
		return ret
	}
	return *o.SupportUnsignedRequestObject
}

// GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetSupportUnsignedRequestObjectOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportUnsignedRequestObject) {
		return nil, false
	}
	return o.SupportUnsignedRequestObject, true
}

// HasSupportUnsignedRequestObject returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasSupportUnsignedRequestObject() bool {
	if o != nil && !IsNil(o.SupportUnsignedRequestObject) {
		return true
	}

	return false
}

// SetSupportUnsignedRequestObject gets a reference to the given bool and assigns it to the SupportUnsignedRequestObject field.
func (o *ApplicationOIDCAllOf) SetSupportUnsignedRequestObject(v bool) {
	o.SupportUnsignedRequestObject = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetTags() []EnumApplicationTags {
	if o == nil || IsNil(o.Tags) {
		var ret []EnumApplicationTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetTagsOk() ([]EnumApplicationTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []EnumApplicationTags and assigns it to the Tags field.
func (o *ApplicationOIDCAllOf) SetTags(v []EnumApplicationTags) {
	o.Tags = v
}

// GetTargetLinkUri returns the TargetLinkUri field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetTargetLinkUri() string {
	if o == nil || IsNil(o.TargetLinkUri) {
		var ret string
		return ret
	}
	return *o.TargetLinkUri
}

// GetTargetLinkUriOk returns a tuple with the TargetLinkUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetTargetLinkUriOk() (*string, bool) {
	if o == nil || IsNil(o.TargetLinkUri) {
		return nil, false
	}
	return o.TargetLinkUri, true
}

// HasTargetLinkUri returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasTargetLinkUri() bool {
	if o != nil && !IsNil(o.TargetLinkUri) {
		return true
	}

	return false
}

// SetTargetLinkUri gets a reference to the given string and assigns it to the TargetLinkUri field.
func (o *ApplicationOIDCAllOf) SetTargetLinkUri(v string) {
	o.TargetLinkUri = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value
func (o *ApplicationOIDCAllOf) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod {
	if o == nil {
		var ret EnumApplicationOIDCTokenAuthMethod
		return ret
	}

	return o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpointAuthMethod, true
}

// SetTokenEndpointAuthMethod sets field value
func (o *ApplicationOIDCAllOf) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod) {
	o.TokenEndpointAuthMethod = v
}

func (o ApplicationOIDCAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOIDCAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowWildcardInRedirectUris) {
		toSerialize["allowWildcardInRedirectUris"] = o.AllowWildcardInRedirectUris
	}
	if !IsNil(o.AssignActorRoles) {
		toSerialize["assignActorRoles"] = o.AssignActorRoles
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.PackageName) {
		toSerialize["packageName"] = o.PackageName
	}
	if !IsNil(o.Kerberos) {
		toSerialize["kerberos"] = o.Kerberos
	}
	toSerialize["grantTypes"] = o.GrantTypes
	if !IsNil(o.HomePageUrl) {
		toSerialize["homePageUrl"] = o.HomePageUrl
	}
	if !IsNil(o.InitiateLoginUri) {
		toSerialize["initiateLoginUri"] = o.InitiateLoginUri
	}
	if !IsNil(o.PkceEnforcement) {
		toSerialize["pkceEnforcement"] = o.PkceEnforcement
	}
	if !IsNil(o.PostLogoutRedirectUris) {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if !IsNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if !IsNil(o.RefreshTokenRollingDuration) {
		toSerialize["refreshTokenRollingDuration"] = o.RefreshTokenRollingDuration
	}
	if !IsNil(o.RefreshTokenRollingGracePeriodDuration) {
		toSerialize["refreshTokenRollingGracePeriodDuration"] = o.RefreshTokenRollingGracePeriodDuration
	}
	if !IsNil(o.ResponseTypes) {
		toSerialize["responseTypes"] = o.ResponseTypes
	}
	if !IsNil(o.SupportUnsignedRequestObject) {
		toSerialize["supportUnsignedRequestObject"] = o.SupportUnsignedRequestObject
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TargetLinkUri) {
		toSerialize["targetLinkUri"] = o.TargetLinkUri
	}
	toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod
	return toSerialize, nil
}

type NullableApplicationOIDCAllOf struct {
	value *ApplicationOIDCAllOf
	isSet bool
}

func (v NullableApplicationOIDCAllOf) Get() *ApplicationOIDCAllOf {
	return v.value
}

func (v *NullableApplicationOIDCAllOf) Set(val *ApplicationOIDCAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOf(val *ApplicationOIDCAllOf) *NullableApplicationOIDCAllOf {
	return &NullableApplicationOIDCAllOf{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


