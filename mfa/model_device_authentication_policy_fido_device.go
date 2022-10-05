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

// DeviceAuthenticationPolicyFIDODevice struct for DeviceAuthenticationPolicyFIDODevice
type DeviceAuthenticationPolicyFIDODevice struct {
	// Enabled or disabled in the policy.
	Enabled bool `json:"enabled"`
	// Specifies the FIDO policy UUID. This property can be null. When null, the environment's default FIDO Policy is used.
	FidoPolicyId *string `json:"fidoPolicyId,omitempty"`
}

// NewDeviceAuthenticationPolicyFIDODevice instantiates a new DeviceAuthenticationPolicyFIDODevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyFIDODevice(enabled bool) *DeviceAuthenticationPolicyFIDODevice {
	this := DeviceAuthenticationPolicyFIDODevice{}
	this.Enabled = enabled
	return &this
}

// NewDeviceAuthenticationPolicyFIDODeviceWithDefaults instantiates a new DeviceAuthenticationPolicyFIDODevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyFIDODeviceWithDefaults() *DeviceAuthenticationPolicyFIDODevice {
	this := DeviceAuthenticationPolicyFIDODevice{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyFIDODevice) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyFIDODevice) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyFIDODevice) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFidoPolicyId returns the FidoPolicyId field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyFIDODevice) GetFidoPolicyId() string {
	if o == nil || o.FidoPolicyId == nil {
		var ret string
		return ret
	}
	return *o.FidoPolicyId
}

// GetFidoPolicyIdOk returns a tuple with the FidoPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyFIDODevice) GetFidoPolicyIdOk() (*string, bool) {
	if o == nil || o.FidoPolicyId == nil {
		return nil, false
	}
	return o.FidoPolicyId, true
}

// HasFidoPolicyId returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyFIDODevice) HasFidoPolicyId() bool {
	if o != nil && o.FidoPolicyId != nil {
		return true
	}

	return false
}

// SetFidoPolicyId gets a reference to the given string and assigns it to the FidoPolicyId field.
func (o *DeviceAuthenticationPolicyFIDODevice) SetFidoPolicyId(v string) {
	o.FidoPolicyId = &v
}

func (o DeviceAuthenticationPolicyFIDODevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FidoPolicyId != nil {
		toSerialize["fidoPolicyId"] = o.FidoPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyFIDODevice struct {
	value *DeviceAuthenticationPolicyFIDODevice
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyFIDODevice) Get() *DeviceAuthenticationPolicyFIDODevice {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyFIDODevice) Set(val *DeviceAuthenticationPolicyFIDODevice) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyFIDODevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyFIDODevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyFIDODevice(val *DeviceAuthenticationPolicyFIDODevice) *NullableDeviceAuthenticationPolicyFIDODevice {
	return &NullableDeviceAuthenticationPolicyFIDODevice{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyFIDODevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyFIDODevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


