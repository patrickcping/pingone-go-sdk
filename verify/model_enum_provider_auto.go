/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"fmt"
)

// EnumProviderAuto A string that defines the automatic verification subprocessor; can be `MITEK` or `VERIFF`.
type EnumProviderAuto string

// List of EnumProviderAuto
const (
	ENUMPROVIDERAUTO_MITEK  EnumProviderAuto = "MITEK"
	ENUMPROVIDERAUTO_VERIFF EnumProviderAuto = "VERIFF"
)

// All allowed values of EnumProviderAuto enum
var AllowedEnumProviderAutoEnumValues = []EnumProviderAuto{
	"MITEK",
	"VERIFF",
}

func (v *EnumProviderAuto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumProviderAuto(value)
	for _, existing := range AllowedEnumProviderAutoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumProviderAuto(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumProviderAutoFromValue returns a pointer to a valid EnumProviderAuto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumProviderAutoFromValue(v string) (*EnumProviderAuto, error) {
	ev := EnumProviderAuto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumProviderAuto: valid values are %v", v, AllowedEnumProviderAutoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumProviderAuto) IsValid() bool {
	for _, existing := range AllowedEnumProviderAutoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumProviderAuto value
func (v EnumProviderAuto) Ptr() *EnumProviderAuto {
	return &v
}

type NullableEnumProviderAuto struct {
	value *EnumProviderAuto
	isSet bool
}

func (v NullableEnumProviderAuto) Get() *EnumProviderAuto {
	return v.value
}

func (v *NullableEnumProviderAuto) Set(val *EnumProviderAuto) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumProviderAuto) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumProviderAuto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumProviderAuto(val *EnumProviderAuto) *NullableEnumProviderAuto {
	return &NullableEnumProviderAuto{value: val, isSet: true}
}

func (v NullableEnumProviderAuto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumProviderAuto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
