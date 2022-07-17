/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumRegionCode A string that specifies the region in which this environment will be used. The value is set when the environment is created and cannot be updated.
type EnumRegionCode string

// List of EnumRegionCode
const (
	AP EnumRegionCode = "AP"
	CA EnumRegionCode = "CA"
	EU EnumRegionCode = "EU"
	NA EnumRegionCode = "NA"
)

// All allowed values of EnumRegionCode enum
var AllowedEnumRegionCodeEnumValues = []EnumRegionCode{
	"AP",
	"CA",
	"EU",
	"NA",
}

func (v *EnumRegionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRegionCode(value)
	for _, existing := range AllowedEnumRegionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumRegionCode", value)
}

// NewEnumRegionCodeFromValue returns a pointer to a valid EnumRegionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRegionCodeFromValue(v string) (*EnumRegionCode, error) {
	ev := EnumRegionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRegionCode: valid values are %v", v, AllowedEnumRegionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRegionCode) IsValid() bool {
	for _, existing := range AllowedEnumRegionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRegionCode value
func (v EnumRegionCode) Ptr() *EnumRegionCode {
	return &v
}

type NullableEnumRegionCode struct {
	value *EnumRegionCode
	isSet bool
}

func (v NullableEnumRegionCode) Get() *EnumRegionCode {
	return v.value
}

func (v *NullableEnumRegionCode) Set(val *EnumRegionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRegionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRegionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRegionCode(val *EnumRegionCode) *NullableEnumRegionCode {
	return &NullableEnumRegionCode{value: val, isSet: true}
}

func (v NullableEnumRegionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRegionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

