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

// checks if the SignOnPolicyActionIDPAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionIDPAllOf{}

// SignOnPolicyActionIDPAllOf struct for SignOnPolicyActionIDPAllOf
type SignOnPolicyActionIDPAllOf struct {
	// A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type `SAML` or `OPENID_CONNECT`.
	AcrValues *string `json:"acrValues,omitempty"`
	IdentityProvider SignOnPolicyActionIDPAllOfIdentityProvider `json:"identityProvider"`
	// A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider.
	PassUserContext *bool `json:"passUserContext,omitempty"`
	Registration *SignOnPolicyActionIDPAllOfRegistration `json:"registration,omitempty"`
}

// NewSignOnPolicyActionIDPAllOf instantiates a new SignOnPolicyActionIDPAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDPAllOf(identityProvider SignOnPolicyActionIDPAllOfIdentityProvider) *SignOnPolicyActionIDPAllOf {
	this := SignOnPolicyActionIDPAllOf{}
	this.IdentityProvider = identityProvider
	return &this
}

// NewSignOnPolicyActionIDPAllOfWithDefaults instantiates a new SignOnPolicyActionIDPAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDPAllOfWithDefaults() *SignOnPolicyActionIDPAllOf {
	this := SignOnPolicyActionIDPAllOf{}
	return &this
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDPAllOf) GetAcrValues() string {
	if o == nil || IsNil(o.AcrValues) {
		var ret string
		return ret
	}
	return *o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOf) GetAcrValuesOk() (*string, bool) {
	if o == nil || IsNil(o.AcrValues) {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDPAllOf) HasAcrValues() bool {
	if o != nil && !IsNil(o.AcrValues) {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given string and assigns it to the AcrValues field.
func (o *SignOnPolicyActionIDPAllOf) SetAcrValues(v string) {
	o.AcrValues = &v
}

// GetIdentityProvider returns the IdentityProvider field value
func (o *SignOnPolicyActionIDPAllOf) GetIdentityProvider() SignOnPolicyActionIDPAllOfIdentityProvider {
	if o == nil {
		var ret SignOnPolicyActionIDPAllOfIdentityProvider
		return ret
	}

	return o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOf) GetIdentityProviderOk() (*SignOnPolicyActionIDPAllOfIdentityProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityProvider, true
}

// SetIdentityProvider sets field value
func (o *SignOnPolicyActionIDPAllOf) SetIdentityProvider(v SignOnPolicyActionIDPAllOfIdentityProvider) {
	o.IdentityProvider = v
}

// GetPassUserContext returns the PassUserContext field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDPAllOf) GetPassUserContext() bool {
	if o == nil || IsNil(o.PassUserContext) {
		var ret bool
		return ret
	}
	return *o.PassUserContext
}

// GetPassUserContextOk returns a tuple with the PassUserContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOf) GetPassUserContextOk() (*bool, bool) {
	if o == nil || IsNil(o.PassUserContext) {
		return nil, false
	}
	return o.PassUserContext, true
}

// HasPassUserContext returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDPAllOf) HasPassUserContext() bool {
	if o != nil && !IsNil(o.PassUserContext) {
		return true
	}

	return false
}

// SetPassUserContext gets a reference to the given bool and assigns it to the PassUserContext field.
func (o *SignOnPolicyActionIDPAllOf) SetPassUserContext(v bool) {
	o.PassUserContext = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionIDPAllOf) GetRegistration() SignOnPolicyActionIDPAllOfRegistration {
	if o == nil || IsNil(o.Registration) {
		var ret SignOnPolicyActionIDPAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOf) GetRegistrationOk() (*SignOnPolicyActionIDPAllOfRegistration, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionIDPAllOf) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionIDPAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionIDPAllOf) SetRegistration(v SignOnPolicyActionIDPAllOfRegistration) {
	o.Registration = &v
}

func (o SignOnPolicyActionIDPAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionIDPAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcrValues) {
		toSerialize["acrValues"] = o.AcrValues
	}
	toSerialize["identityProvider"] = o.IdentityProvider
	if !IsNil(o.PassUserContext) {
		toSerialize["passUserContext"] = o.PassUserContext
	}
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionIDPAllOf struct {
	value *SignOnPolicyActionIDPAllOf
	isSet bool
}

func (v NullableSignOnPolicyActionIDPAllOf) Get() *SignOnPolicyActionIDPAllOf {
	return v.value
}

func (v *NullableSignOnPolicyActionIDPAllOf) Set(val *SignOnPolicyActionIDPAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDPAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDPAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDPAllOf(val *SignOnPolicyActionIDPAllOf) *NullableSignOnPolicyActionIDPAllOf {
	return &NullableSignOnPolicyActionIDPAllOf{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDPAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDPAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


