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

// SignOnPolicyActionMFAAllOfBoundBiometrics Specifies MFA through a FIDO2 biometrics platform device.
type SignOnPolicyActionMFAAllOfBoundBiometrics struct {
	// A boolean that specifies whether to permit users to authenticate with a FIDO2 biometrics platform device.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFAAllOfBoundBiometrics instantiates a new SignOnPolicyActionMFAAllOfBoundBiometrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAllOfBoundBiometrics() *SignOnPolicyActionMFAAllOfBoundBiometrics {
	this := SignOnPolicyActionMFAAllOfBoundBiometrics{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFAAllOfBoundBiometricsWithDefaults instantiates a new SignOnPolicyActionMFAAllOfBoundBiometrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAllOfBoundBiometricsWithDefaults() *SignOnPolicyActionMFAAllOfBoundBiometrics {
	this := SignOnPolicyActionMFAAllOfBoundBiometrics{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAllOfBoundBiometrics) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfBoundBiometrics) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAllOfBoundBiometrics) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAAllOfBoundBiometrics) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFAAllOfBoundBiometrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAAllOfBoundBiometrics struct {
	value *SignOnPolicyActionMFAAllOfBoundBiometrics
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAllOfBoundBiometrics) Get() *SignOnPolicyActionMFAAllOfBoundBiometrics {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAllOfBoundBiometrics) Set(val *SignOnPolicyActionMFAAllOfBoundBiometrics) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAllOfBoundBiometrics) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAllOfBoundBiometrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAllOfBoundBiometrics(val *SignOnPolicyActionMFAAllOfBoundBiometrics) *NullableSignOnPolicyActionMFAAllOfBoundBiometrics {
	return &NullableSignOnPolicyActionMFAAllOfBoundBiometrics{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAllOfBoundBiometrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAllOfBoundBiometrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


