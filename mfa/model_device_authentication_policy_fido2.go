/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the DeviceAuthenticationPolicyFido2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyFido2{}

// DeviceAuthenticationPolicyFido2 struct for DeviceAuthenticationPolicyFido2
type DeviceAuthenticationPolicyFido2 struct {
	// A boolean that specifies whether the method is enabled or disabled in the policy.
	Enabled bool `json:"enabled"`
	// A boolean to specify whether pairing is disabled for the method.
	PairingDisabled *bool `json:"pairingDisabled,omitempty"`
	// Specifies the UUID that represents the FIDO2 policy in PingOne. This property can be null. When null, the environment's default FIDO2 Policy is used.
	Fido2PolicyId *string `json:"fido2PolicyId,omitempty"`
}

// NewDeviceAuthenticationPolicyFido2 instantiates a new DeviceAuthenticationPolicyFido2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyFido2(enabled bool) *DeviceAuthenticationPolicyFido2 {
	this := DeviceAuthenticationPolicyFido2{}
	this.Enabled = enabled
	return &this
}

// NewDeviceAuthenticationPolicyFido2WithDefaults instantiates a new DeviceAuthenticationPolicyFido2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyFido2WithDefaults() *DeviceAuthenticationPolicyFido2 {
	this := DeviceAuthenticationPolicyFido2{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyFido2) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyFido2) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyFido2) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPairingDisabled returns the PairingDisabled field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyFido2) GetPairingDisabled() bool {
	if o == nil || IsNil(o.PairingDisabled) {
		var ret bool
		return ret
	}
	return *o.PairingDisabled
}

// GetPairingDisabledOk returns a tuple with the PairingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyFido2) GetPairingDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PairingDisabled) {
		return nil, false
	}
	return o.PairingDisabled, true
}

// HasPairingDisabled returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyFido2) HasPairingDisabled() bool {
	if o != nil && !IsNil(o.PairingDisabled) {
		return true
	}

	return false
}

// SetPairingDisabled gets a reference to the given bool and assigns it to the PairingDisabled field.
func (o *DeviceAuthenticationPolicyFido2) SetPairingDisabled(v bool) {
	o.PairingDisabled = &v
}

// GetFido2PolicyId returns the Fido2PolicyId field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyFido2) GetFido2PolicyId() string {
	if o == nil || IsNil(o.Fido2PolicyId) {
		var ret string
		return ret
	}
	return *o.Fido2PolicyId
}

// GetFido2PolicyIdOk returns a tuple with the Fido2PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyFido2) GetFido2PolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.Fido2PolicyId) {
		return nil, false
	}
	return o.Fido2PolicyId, true
}

// HasFido2PolicyId returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyFido2) HasFido2PolicyId() bool {
	if o != nil && !IsNil(o.Fido2PolicyId) {
		return true
	}

	return false
}

// SetFido2PolicyId gets a reference to the given string and assigns it to the Fido2PolicyId field.
func (o *DeviceAuthenticationPolicyFido2) SetFido2PolicyId(v string) {
	o.Fido2PolicyId = &v
}

func (o DeviceAuthenticationPolicyFido2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyFido2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.PairingDisabled) {
		toSerialize["pairingDisabled"] = o.PairingDisabled
	}
	if !IsNil(o.Fido2PolicyId) {
		toSerialize["fido2PolicyId"] = o.Fido2PolicyId
	}
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyFido2 struct {
	value *DeviceAuthenticationPolicyFido2
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyFido2) Get() *DeviceAuthenticationPolicyFido2 {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyFido2) Set(val *DeviceAuthenticationPolicyFido2) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyFido2) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyFido2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyFido2(val *DeviceAuthenticationPolicyFido2) *NullableDeviceAuthenticationPolicyFido2 {
	return &NullableDeviceAuthenticationPolicyFido2{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyFido2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyFido2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


