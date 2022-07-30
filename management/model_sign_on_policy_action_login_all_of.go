/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionLoginAllOf struct for SignOnPolicyActionLoginAllOf
type SignOnPolicyActionLoginAllOf struct {
	// A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user's profile during account creation. This is an optional property. If omitted, the default value is set to false.
	ConfirmIdentityProviderAttributes *bool `json:"confirmIdentityProviderAttributes,omitempty"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	Recovery *SignOnPolicyActionLoginAllOfRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionLoginAllOfRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginAllOfSocialProviders `json:"socialProviders,omitempty"`
}

// NewSignOnPolicyActionLoginAllOf instantiates a new SignOnPolicyActionLoginAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOf() *SignOnPolicyActionLoginAllOf {
	this := SignOnPolicyActionLoginAllOf{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// NewSignOnPolicyActionLoginAllOfWithDefaults instantiates a new SignOnPolicyActionLoginAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfWithDefaults() *SignOnPolicyActionLoginAllOf {
	this := SignOnPolicyActionLoginAllOf{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOf) GetConfirmIdentityProviderAttributes() bool {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmIdentityProviderAttributes
}

// GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOf) GetConfirmIdentityProviderAttributesOk() (*bool, bool) {
	if o == nil || o.ConfirmIdentityProviderAttributes == nil {
		return nil, false
	}
	return o.ConfirmIdentityProviderAttributes, true
}

// HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOf) HasConfirmIdentityProviderAttributes() bool {
	if o != nil && o.ConfirmIdentityProviderAttributes != nil {
		return true
	}

	return false
}

// SetConfirmIdentityProviderAttributes gets a reference to the given bool and assigns it to the ConfirmIdentityProviderAttributes field.
func (o *SignOnPolicyActionLoginAllOf) SetConfirmIdentityProviderAttributes(v bool) {
	o.ConfirmIdentityProviderAttributes = &v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOf) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOf) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || o.EnforceLockoutForIdentityProviders == nil {
		return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOf) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && o.EnforceLockoutForIdentityProviders != nil {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyActionLoginAllOf) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOf) GetRecovery() SignOnPolicyActionLoginAllOfRecovery {
	if o == nil || o.Recovery == nil {
		var ret SignOnPolicyActionLoginAllOfRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOf) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOf) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginAllOfRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyActionLoginAllOf) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOf) GetRegistration() SignOnPolicyActionLoginAllOfRegistration {
	if o == nil || o.Registration == nil {
		var ret SignOnPolicyActionLoginAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOf) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOf) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionLoginAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionLoginAllOf) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOf) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders {
	if o == nil || o.SocialProviders == nil {
		var ret []SignOnPolicyActionLoginAllOfSocialProviders
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOf) GetSocialProvidersOk() ([]SignOnPolicyActionLoginAllOfSocialProviders, bool) {
	if o == nil || o.SocialProviders == nil {
		return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOf) HasSocialProviders() bool {
	if o != nil && o.SocialProviders != nil {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginAllOfSocialProviders and assigns it to the SocialProviders field.
func (o *SignOnPolicyActionLoginAllOf) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders) {
	o.SocialProviders = v
}

func (o SignOnPolicyActionLoginAllOf) MarshalJSON() ([]byte, error) {
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

type NullableSignOnPolicyActionLoginAllOf struct {
	value *SignOnPolicyActionLoginAllOf
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOf) Get() *SignOnPolicyActionLoginAllOf {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOf) Set(val *SignOnPolicyActionLoginAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOf(val *SignOnPolicyActionLoginAllOf) *NullableSignOnPolicyActionLoginAllOf {
	return &NullableSignOnPolicyActionLoginAllOf{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


