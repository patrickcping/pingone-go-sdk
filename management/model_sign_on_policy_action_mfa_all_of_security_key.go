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

// checks if the SignOnPolicyActionMFAAllOfSecurityKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionMFAAllOfSecurityKey{}

// SignOnPolicyActionMFAAllOfSecurityKey Specifies MFA using a FIDO2 or U2F security key device.
type SignOnPolicyActionMFAAllOfSecurityKey struct {
	// A boolean that specifies whether to permit users to authenticate with a security key device.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFAAllOfSecurityKey instantiates a new SignOnPolicyActionMFAAllOfSecurityKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAllOfSecurityKey() *SignOnPolicyActionMFAAllOfSecurityKey {
	this := SignOnPolicyActionMFAAllOfSecurityKey{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFAAllOfSecurityKeyWithDefaults instantiates a new SignOnPolicyActionMFAAllOfSecurityKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAllOfSecurityKeyWithDefaults() *SignOnPolicyActionMFAAllOfSecurityKey {
	this := SignOnPolicyActionMFAAllOfSecurityKey{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAllOfSecurityKey) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfSecurityKey) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAllOfSecurityKey) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAAllOfSecurityKey) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFAAllOfSecurityKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionMFAAllOfSecurityKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionMFAAllOfSecurityKey struct {
	value *SignOnPolicyActionMFAAllOfSecurityKey
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAllOfSecurityKey) Get() *SignOnPolicyActionMFAAllOfSecurityKey {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAllOfSecurityKey) Set(val *SignOnPolicyActionMFAAllOfSecurityKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAllOfSecurityKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAllOfSecurityKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAllOfSecurityKey(val *SignOnPolicyActionMFAAllOfSecurityKey) *NullableSignOnPolicyActionMFAAllOfSecurityKey {
	return &NullableSignOnPolicyActionMFAAllOfSecurityKey{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAllOfSecurityKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAllOfSecurityKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


