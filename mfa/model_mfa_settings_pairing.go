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

// checks if the MFASettingsPairing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFASettingsPairing{}

// MFASettingsPairing An object that contains pairing settings.
type MFASettingsPairing struct {
	// An integer that defines the maximum number of MFA devices each user can have. This can be any number up to 15. The default value is 5.
	MaxAllowedDevices int32 `json:"maxAllowedDevices"`
	PairingKeyFormat EnumMFASettingsPairingKeyFormat `json:"pairingKeyFormat"`
}

// NewMFASettingsPairing instantiates a new MFASettingsPairing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFASettingsPairing(maxAllowedDevices int32, pairingKeyFormat EnumMFASettingsPairingKeyFormat) *MFASettingsPairing {
	this := MFASettingsPairing{}
	this.MaxAllowedDevices = maxAllowedDevices
	this.PairingKeyFormat = pairingKeyFormat
	return &this
}

// NewMFASettingsPairingWithDefaults instantiates a new MFASettingsPairing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFASettingsPairingWithDefaults() *MFASettingsPairing {
	this := MFASettingsPairing{}
	var maxAllowedDevices int32 = 5
	this.MaxAllowedDevices = maxAllowedDevices
	return &this
}

// GetMaxAllowedDevices returns the MaxAllowedDevices field value
func (o *MFASettingsPairing) GetMaxAllowedDevices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxAllowedDevices
}

// GetMaxAllowedDevicesOk returns a tuple with the MaxAllowedDevices field value
// and a boolean to check if the value has been set.
func (o *MFASettingsPairing) GetMaxAllowedDevicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAllowedDevices, true
}

// SetMaxAllowedDevices sets field value
func (o *MFASettingsPairing) SetMaxAllowedDevices(v int32) {
	o.MaxAllowedDevices = v
}

// GetPairingKeyFormat returns the PairingKeyFormat field value
func (o *MFASettingsPairing) GetPairingKeyFormat() EnumMFASettingsPairingKeyFormat {
	if o == nil {
		var ret EnumMFASettingsPairingKeyFormat
		return ret
	}

	return o.PairingKeyFormat
}

// GetPairingKeyFormatOk returns a tuple with the PairingKeyFormat field value
// and a boolean to check if the value has been set.
func (o *MFASettingsPairing) GetPairingKeyFormatOk() (*EnumMFASettingsPairingKeyFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PairingKeyFormat, true
}

// SetPairingKeyFormat sets field value
func (o *MFASettingsPairing) SetPairingKeyFormat(v EnumMFASettingsPairingKeyFormat) {
	o.PairingKeyFormat = v
}

func (o MFASettingsPairing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFASettingsPairing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxAllowedDevices"] = o.MaxAllowedDevices
	toSerialize["pairingKeyFormat"] = o.PairingKeyFormat
	return toSerialize, nil
}

type NullableMFASettingsPairing struct {
	value *MFASettingsPairing
	isSet bool
}

func (v NullableMFASettingsPairing) Get() *MFASettingsPairing {
	return v.value
}

func (v *NullableMFASettingsPairing) Set(val *MFASettingsPairing) {
	v.value = val
	v.isSet = true
}

func (v NullableMFASettingsPairing) IsSet() bool {
	return v.isSet
}

func (v *NullableMFASettingsPairing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFASettingsPairing(val *MFASettingsPairing) *NullableMFASettingsPairing {
	return &NullableMFASettingsPairing{value: val, isSet: true}
}

func (v NullableMFASettingsPairing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFASettingsPairing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


