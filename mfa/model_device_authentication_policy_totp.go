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

// DeviceAuthenticationPolicyTotp TOTP device authentication policy settings.
type DeviceAuthenticationPolicyTotp struct {
	// Enabled or disabled in the policy.
	Enabled *bool `json:"enabled,omitempty"`
	Otp *DeviceAuthenticationPolicyMobileOtpWindow `json:"otp,omitempty"`
}

// NewDeviceAuthenticationPolicyTotp instantiates a new DeviceAuthenticationPolicyTotp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyTotp() *DeviceAuthenticationPolicyTotp {
	this := DeviceAuthenticationPolicyTotp{}
	return &this
}

// NewDeviceAuthenticationPolicyTotpWithDefaults instantiates a new DeviceAuthenticationPolicyTotp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyTotpWithDefaults() *DeviceAuthenticationPolicyTotp {
	this := DeviceAuthenticationPolicyTotp{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyTotp) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyTotp) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DeviceAuthenticationPolicyTotp) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyTotp) GetOtp() DeviceAuthenticationPolicyMobileOtpWindow {
	if o == nil || o.Otp == nil {
		var ret DeviceAuthenticationPolicyMobileOtpWindow
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetOtpOk() (*DeviceAuthenticationPolicyMobileOtpWindow, bool) {
	if o == nil || o.Otp == nil {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyTotp) HasOtp() bool {
	if o != nil && o.Otp != nil {
		return true
	}

	return false
}

// SetOtp gets a reference to the given DeviceAuthenticationPolicyMobileOtpWindow and assigns it to the Otp field.
func (o *DeviceAuthenticationPolicyTotp) SetOtp(v DeviceAuthenticationPolicyMobileOtpWindow) {
	o.Otp = &v
}

func (o DeviceAuthenticationPolicyTotp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Otp != nil {
		toSerialize["otp"] = o.Otp
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyTotp struct {
	value *DeviceAuthenticationPolicyTotp
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyTotp) Get() *DeviceAuthenticationPolicyTotp {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyTotp) Set(val *DeviceAuthenticationPolicyTotp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyTotp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyTotp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyTotp(val *DeviceAuthenticationPolicyTotp) *NullableDeviceAuthenticationPolicyTotp {
	return &NullableDeviceAuthenticationPolicyTotp{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyTotp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyTotp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


