/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionLogin struct for SignOnPolicyActionLogin
type SignOnPolicyActionLogin struct {
	// A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user's profile during account creation. This is an optional property. If omitted, the default value is set to false.
	ConfirmIdentityProviderAttributes *bool `json:"confirmIdentityProviderAttributes,omitempty"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	Recovery *SignOnPolicyActionLoginRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionLoginRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginSocialProvidersInner `json:"socialProviders,omitempty"`
}

// NewSignOnPolicyActionLogin instantiates a new SignOnPolicyActionLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLogin() *SignOnPolicyActionLogin {
	this := SignOnPolicyActionLogin{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// NewSignOnPolicyActionLoginWithDefaults instantiates a new SignOnPolicyActionLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginWithDefaults() *SignOnPolicyActionLogin {
	this := SignOnPolicyActionLogin{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetConfirmIdentityProviderAttributes() bool {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmIdentityProviderAttributes
}

// GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetConfirmIdentityProviderAttributesOk() (*bool, bool) {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		return nil, false
	}
	return o.ConfirmIdentityProviderAttributes, true
}

// HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasConfirmIdentityProviderAttributes() bool {
	if o != nil && o.ConfirmIdentityProviderAttributes != nil {
		return true
	}

	return false
}

// SetConfirmIdentityProviderAttributes gets a reference to the given bool and assigns it to the ConfirmIdentityProviderAttributes field.
func (o *SignOnPolicyActionLogin) SetConfirmIdentityProviderAttributes(v bool) {
	o.ConfirmIdentityProviderAttributes = &v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && o.EnforceLockoutForIdentityProviders != nil {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyActionLogin) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetRecovery() SignOnPolicyActionLoginRecovery {
	if o == nil || o.Recovery == nil {
		var ret SignOnPolicyActionLoginRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetRecoveryOk() (*SignOnPolicyActionLoginRecovery, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyActionLogin) SetRecovery(v SignOnPolicyActionLoginRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetRegistration() SignOnPolicyActionLoginRegistration {
	if o == nil || o.Registration == nil {
		var ret SignOnPolicyActionLoginRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetRegistrationOk() (*SignOnPolicyActionLoginRegistration, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionLoginRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionLogin) SetRegistration(v SignOnPolicyActionLoginRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetSocialProviders() []SignOnPolicyActionLoginSocialProvidersInner {
	if o == nil || o.SocialProviders == nil {
		var ret []SignOnPolicyActionLoginSocialProvidersInner
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetSocialProvidersOk() ([]SignOnPolicyActionLoginSocialProvidersInner, bool) {
	if o == nil || o.SocialProviders == nil {
		return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasSocialProviders() bool {
	if o != nil && o.SocialProviders != nil {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginSocialProvidersInner and assigns it to the SocialProviders field.
func (o *SignOnPolicyActionLogin) SetSocialProviders(v []SignOnPolicyActionLoginSocialProvidersInner) {
	o.SocialProviders = v
}

func (o SignOnPolicyActionLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfirmIdentityProviderAttributes != nil {
		toSerialize["confirmIdentityProviderAttributes"] = o.ConfirmIdentityProviderAttributes
	}
	if o.EnforceLockoutForIdentityProviders != nil {
		toSerialize["enforceLockoutForIdentityProviders"] = o.EnforceLockoutForIdentityProviders
	}
	if o.Recovery != nil {
		toSerialize["recovery"] = o.Recovery
	}
	if o.Registration != nil {
		toSerialize["registration"] = o.Registration
	}
	if o.SocialProviders != nil {
		toSerialize["socialProviders"] = o.SocialProviders
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionLogin struct {
	value *SignOnPolicyActionLogin
	isSet bool
}

func (v NullableSignOnPolicyActionLogin) Get() *SignOnPolicyActionLogin {
	return v.value
}

func (v *NullableSignOnPolicyActionLogin) Set(val *SignOnPolicyActionLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLogin(val *SignOnPolicyActionLogin) *NullableSignOnPolicyActionLogin {
	return &NullableSignOnPolicyActionLogin{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


