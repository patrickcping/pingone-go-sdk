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

// DeviceAuthenticationPolicyTotpOtp struct for DeviceAuthenticationPolicyTotpOtp
type DeviceAuthenticationPolicyTotpOtp struct {
	Failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure `json:"failure"`
}

// NewDeviceAuthenticationPolicyTotpOtp instantiates a new DeviceAuthenticationPolicyTotpOtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyTotpOtp(failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure) *DeviceAuthenticationPolicyTotpOtp {
	this := DeviceAuthenticationPolicyTotpOtp{}
	this.Failure = failure
	return &this
}

// NewDeviceAuthenticationPolicyTotpOtpWithDefaults instantiates a new DeviceAuthenticationPolicyTotpOtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyTotpOtpWithDefaults() *DeviceAuthenticationPolicyTotpOtp {
	this := DeviceAuthenticationPolicyTotpOtp{}
	return &this
}

// GetFailure returns the Failure field value
func (o *DeviceAuthenticationPolicyTotpOtp) GetFailure() DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtpFailure
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotpOtp) GetFailureOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *DeviceAuthenticationPolicyTotpOtp) SetFailure(v DeviceAuthenticationPolicyOfflineDeviceOtpFailure) {
	o.Failure = v
}

func (o DeviceAuthenticationPolicyTotpOtp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["failure"] = o.Failure
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyTotpOtp struct {
	value *DeviceAuthenticationPolicyTotpOtp
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyTotpOtp) Get() *DeviceAuthenticationPolicyTotpOtp {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyTotpOtp) Set(val *DeviceAuthenticationPolicyTotpOtp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyTotpOtp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyTotpOtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyTotpOtp(val *DeviceAuthenticationPolicyTotpOtp) *NullableDeviceAuthenticationPolicyTotpOtp {
	return &NullableDeviceAuthenticationPolicyTotpOtp{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyTotpOtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyTotpOtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


