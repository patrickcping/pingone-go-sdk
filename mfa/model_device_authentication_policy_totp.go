/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the DeviceAuthenticationPolicyTotp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyTotp{}

// DeviceAuthenticationPolicyTotp TOTP device authentication policy settings.
type DeviceAuthenticationPolicyTotp struct {
	// Enabled or disabled in the policy.
	Enabled bool `json:"enabled"`
	// You can set `pairingDisabled` to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices.
	PairingDisabled *bool `json:"pairingDisabled,omitempty"`
	Otp DeviceAuthenticationPolicyTotpOtp `json:"otp"`
	// Set to true if you want to allow users to provide nicknames for devices during pairing.
	PromptForNicknameOnPairing *bool `json:"promptForNicknameOnPairing,omitempty"`
}

// NewDeviceAuthenticationPolicyTotp instantiates a new DeviceAuthenticationPolicyTotp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyTotp(enabled bool, otp DeviceAuthenticationPolicyTotpOtp) *DeviceAuthenticationPolicyTotp {
	this := DeviceAuthenticationPolicyTotp{}
	this.Enabled = enabled
	this.Otp = otp
	return &this
}

// NewDeviceAuthenticationPolicyTotpWithDefaults instantiates a new DeviceAuthenticationPolicyTotp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyTotpWithDefaults() *DeviceAuthenticationPolicyTotp {
	this := DeviceAuthenticationPolicyTotp{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyTotp) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyTotp) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPairingDisabled returns the PairingDisabled field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyTotp) GetPairingDisabled() bool {
	if o == nil || IsNil(o.PairingDisabled) {
		var ret bool
		return ret
	}
	return *o.PairingDisabled
}

// GetPairingDisabledOk returns a tuple with the PairingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetPairingDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PairingDisabled) {
		return nil, false
	}
	return o.PairingDisabled, true
}

// HasPairingDisabled returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyTotp) HasPairingDisabled() bool {
	if o != nil && !IsNil(o.PairingDisabled) {
		return true
	}

	return false
}

// SetPairingDisabled gets a reference to the given bool and assigns it to the PairingDisabled field.
func (o *DeviceAuthenticationPolicyTotp) SetPairingDisabled(v bool) {
	o.PairingDisabled = &v
}

// GetOtp returns the Otp field value
func (o *DeviceAuthenticationPolicyTotp) GetOtp() DeviceAuthenticationPolicyTotpOtp {
	if o == nil {
		var ret DeviceAuthenticationPolicyTotpOtp
		return ret
	}

	return o.Otp
}

// GetOtpOk returns a tuple with the Otp field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetOtpOk() (*DeviceAuthenticationPolicyTotpOtp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Otp, true
}

// SetOtp sets field value
func (o *DeviceAuthenticationPolicyTotp) SetOtp(v DeviceAuthenticationPolicyTotpOtp) {
	o.Otp = v
}

// GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyTotp) GetPromptForNicknameOnPairing() bool {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		var ret bool
		return ret
	}
	return *o.PromptForNicknameOnPairing
}

// GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyTotp) GetPromptForNicknameOnPairingOk() (*bool, bool) {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		return nil, false
	}
	return o.PromptForNicknameOnPairing, true
}

// HasPromptForNicknameOnPairing returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyTotp) HasPromptForNicknameOnPairing() bool {
	if o != nil && !IsNil(o.PromptForNicknameOnPairing) {
		return true
	}

	return false
}

// SetPromptForNicknameOnPairing gets a reference to the given bool and assigns it to the PromptForNicknameOnPairing field.
func (o *DeviceAuthenticationPolicyTotp) SetPromptForNicknameOnPairing(v bool) {
	o.PromptForNicknameOnPairing = &v
}

func (o DeviceAuthenticationPolicyTotp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyTotp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.PairingDisabled) {
		toSerialize["pairingDisabled"] = o.PairingDisabled
	}
	toSerialize["otp"] = o.Otp
	if !IsNil(o.PromptForNicknameOnPairing) {
		toSerialize["promptForNicknameOnPairing"] = o.PromptForNicknameOnPairing
	}
	return toSerialize, nil
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


