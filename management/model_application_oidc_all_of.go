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

// ApplicationOIDCAllOf struct for ApplicationOIDCAllOf
type ApplicationOIDCAllOf struct {
	Mobile *ApplicationOIDCAllOfMobile `json:"mobile,omitempty"`
	// A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.
	BundleId *string `json:"bundleId,omitempty"`
	// A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.
	PackageName *string `json:"packageName,omitempty"`
	Kerberos *ApplicationOIDCAllOfKerberos `json:"kerberos,omitempty"`
	// A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS.
	GrantTypes []EnumApplicationOIDCGrantType `json:"grantTypes"`
	// A string that specifies the custom home page URL for the application.
	HomePageUrl *string `json:"homePageUrl,omitempty"`
	PkceEnforcement *EnumApplicationOIDCPKCEOption `json:"pkceEnforcement,omitempty"`
	// A string that specifies the URLs that the browser can be redirected to after logout.
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris,omitempty"`
	// A string that specifies the callback URI for the authentication response.
	RedirectUris []string `json:"redirectUris,omitempty"`
	// An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the refreshTokenRollingDuration property is specified for the application, then this property must be less than or equal to the value of refreshTokenRollingDuration. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenDuration *int32 `json:"refreshTokenDuration,omitempty"`
	// An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token.
	RefreshTokenRollingDuration *int32 `json:"refreshTokenRollingDuration,omitempty"`
	// A string that specifies the code or token type returned by an authorization request. Options are TOKEN, ID_TOKEN, and CODE. Note that CODE cannot be used in an authorization request with TOKEN or ID_TOKEN because PingOne does not currently support OIDC hybrid flows.
	ResponseTypes []EnumApplicationOIDCResponseType `json:"responseTypes,omitempty"`
	TokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod `json:"tokenEndpointAuthMethod"`
}

// NewApplicationOIDCAllOf instantiates a new ApplicationOIDCAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOf(grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod) *ApplicationOIDCAllOf {
	this := ApplicationOIDCAllOf{}
	this.GrantTypes = grantTypes
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	return &this
}

// NewApplicationOIDCAllOfWithDefaults instantiates a new ApplicationOIDCAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfWithDefaults() *ApplicationOIDCAllOf {
	this := ApplicationOIDCAllOf{}
	return &this
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetMobile() ApplicationOIDCAllOfMobile {
	if o == nil || o.Mobile == nil {
		var ret ApplicationOIDCAllOfMobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given ApplicationOIDCAllOfMobile and assigns it to the Mobile field.
func (o *ApplicationOIDCAllOf) SetMobile(v ApplicationOIDCAllOfMobile) {
	o.Mobile = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetBundleIdOk() (*string, bool) {
	if o == nil || o.BundleId == nil {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *ApplicationOIDCAllOf) SetBundleId(v string) {
	o.BundleId = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *ApplicationOIDCAllOf) SetPackageName(v string) {
	o.PackageName = &v
}

// GetKerberos returns the Kerberos field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetKerberos() ApplicationOIDCAllOfKerberos {
	if o == nil || o.Kerberos == nil {
		var ret ApplicationOIDCAllOfKerberos
		return ret
	}
	return *o.Kerberos
}

// GetKerberosOk returns a tuple with the Kerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetKerberosOk() (*ApplicationOIDCAllOfKerberos, bool) {
	if o == nil || o.Kerberos == nil {
		return nil, false
	}
	return o.Kerberos, true
}

// HasKerberos returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasKerberos() bool {
	if o != nil && o.Kerberos != nil {
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
	if o == nil || o.HomePageUrl == nil {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetHomePageUrlOk() (*string, bool) {
	if o == nil || o.HomePageUrl == nil {
		return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasHomePageUrl() bool {
	if o != nil && o.HomePageUrl != nil {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *ApplicationOIDCAllOf) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetPkceEnforcement returns the PkceEnforcement field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetPkceEnforcement() EnumApplicationOIDCPKCEOption {
	if o == nil || o.PkceEnforcement == nil {
		var ret EnumApplicationOIDCPKCEOption
		return ret
	}
	return *o.PkceEnforcement
}

// GetPkceEnforcementOk returns a tuple with the PkceEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool) {
	if o == nil || o.PkceEnforcement == nil {
		return nil, false
	}
	return o.PkceEnforcement, true
}

// HasPkceEnforcement returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPkceEnforcement() bool {
	if o != nil && o.PkceEnforcement != nil {
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
	if o == nil || o.PostLogoutRedirectUris == nil {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || o.PostLogoutRedirectUris == nil {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasPostLogoutRedirectUris() bool {
	if o != nil && o.PostLogoutRedirectUris != nil {
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
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
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
	if o == nil || o.RefreshTokenDuration == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRefreshTokenDurationOk() (*int32, bool) {
	if o == nil || o.RefreshTokenDuration == nil {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRefreshTokenDuration() bool {
	if o != nil && o.RefreshTokenDuration != nil {
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
	if o == nil || o.RefreshTokenRollingDuration == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenRollingDuration
}

// GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingDurationOk() (*int32, bool) {
	if o == nil || o.RefreshTokenRollingDuration == nil {
		return nil, false
	}
	return o.RefreshTokenRollingDuration, true
}

// HasRefreshTokenRollingDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasRefreshTokenRollingDuration() bool {
	if o != nil && o.RefreshTokenRollingDuration != nil {
		return true
	}

	return false
}

// SetRefreshTokenRollingDuration gets a reference to the given int32 and assigns it to the RefreshTokenRollingDuration field.
func (o *ApplicationOIDCAllOf) SetRefreshTokenRollingDuration(v int32) {
	o.RefreshTokenRollingDuration = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOf) GetResponseTypes() []EnumApplicationOIDCResponseType {
	if o == nil || o.ResponseTypes == nil {
		var ret []EnumApplicationOIDCResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOf) GetResponseTypesOk() ([]EnumApplicationOIDCResponseType, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOf) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []EnumApplicationOIDCResponseType and assigns it to the ResponseTypes field.
func (o *ApplicationOIDCAllOf) SetResponseTypes(v []EnumApplicationOIDCResponseType) {
	o.ResponseTypes = v
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
	toSerialize := map[string]interface{}{}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	if o.BundleId != nil {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	if o.Kerberos != nil {
		toSerialize["kerberos"] = o.Kerberos
	}
	if true {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if o.HomePageUrl != nil {
		toSerialize["homePageUrl"] = o.HomePageUrl
	}
	if o.PkceEnforcement != nil {
		toSerialize["pkceEnforcement"] = o.PkceEnforcement
	}
	if o.PostLogoutRedirectUris != nil {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris
	}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if o.RefreshTokenDuration != nil {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if o.RefreshTokenRollingDuration != nil {
		toSerialize["refreshTokenRollingDuration"] = o.RefreshTokenRollingDuration
	}
	if o.ResponseTypes != nil {
		toSerialize["responseTypes"] = o.ResponseTypes
	}
	if true {
		toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod
	}
	return json.Marshal(toSerialize)
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


