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

// DeviceAuthenticationPolicyMobileOtpWindow struct for DeviceAuthenticationPolicyMobileOtpWindow
type DeviceAuthenticationPolicyMobileOtpWindow struct {
	StepSize *DeviceAuthenticationPolicyMobileOtpWindowStepSize `json:"stepSize,omitempty"`
}

// NewDeviceAuthenticationPolicyMobileOtpWindow instantiates a new DeviceAuthenticationPolicyMobileOtpWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileOtpWindow() *DeviceAuthenticationPolicyMobileOtpWindow {
	this := DeviceAuthenticationPolicyMobileOtpWindow{}
	return &this
}

// NewDeviceAuthenticationPolicyMobileOtpWindowWithDefaults instantiates a new DeviceAuthenticationPolicyMobileOtpWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileOtpWindowWithDefaults() *DeviceAuthenticationPolicyMobileOtpWindow {
	this := DeviceAuthenticationPolicyMobileOtpWindow{}
	return &this
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileOtpWindow) GetStepSize() DeviceAuthenticationPolicyMobileOtpWindowStepSize {
	if o == nil || o.StepSize == nil {
		var ret DeviceAuthenticationPolicyMobileOtpWindowStepSize
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtpWindow) GetStepSizeOk() (*DeviceAuthenticationPolicyMobileOtpWindowStepSize, bool) {
	if o == nil || o.StepSize == nil {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileOtpWindow) HasStepSize() bool {
	if o != nil && o.StepSize != nil {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given DeviceAuthenticationPolicyMobileOtpWindowStepSize and assigns it to the StepSize field.
func (o *DeviceAuthenticationPolicyMobileOtpWindow) SetStepSize(v DeviceAuthenticationPolicyMobileOtpWindowStepSize) {
	o.StepSize = &v
}

func (o DeviceAuthenticationPolicyMobileOtpWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StepSize != nil {
		toSerialize["stepSize"] = o.StepSize
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyMobileOtpWindow struct {
	value *DeviceAuthenticationPolicyMobileOtpWindow
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) Get() *DeviceAuthenticationPolicyMobileOtpWindow {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) Set(val *DeviceAuthenticationPolicyMobileOtpWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileOtpWindow(val *DeviceAuthenticationPolicyMobileOtpWindow) *NullableDeviceAuthenticationPolicyMobileOtpWindow {
	return &NullableDeviceAuthenticationPolicyMobileOtpWindow{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


