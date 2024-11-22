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

// EnumProviderManual A string that defines the automatic verification subprocessor; can be `MITEK`.
type EnumProviderManual string

// List of EnumProviderManual
const (
	ENUMPROVIDERMANUAL_MITEK EnumProviderManual = "MITEK"
)

// All allowed values of EnumProviderManual enum
var AllowedEnumProviderManualEnumValues = []EnumProviderManual{
	"MITEK",
}

func (v *EnumProviderManual) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumProviderManual(value)
	for _, existing := range AllowedEnumProviderManualEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumProviderManual(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumProviderManualFromValue returns a pointer to a valid EnumProviderManual
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumProviderManualFromValue(v string) (*EnumProviderManual, error) {
	ev := EnumProviderManual(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumProviderManual: valid values are %v", v, AllowedEnumProviderManualEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumProviderManual) IsValid() bool {
	for _, existing := range AllowedEnumProviderManualEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumProviderManual value
func (v EnumProviderManual) Ptr() *EnumProviderManual {
	return &v
}

type NullableEnumProviderManual struct {
	value *EnumProviderManual
	isSet bool
}

func (v NullableEnumProviderManual) Get() *EnumProviderManual {
	return v.value
}

func (v *NullableEnumProviderManual) Set(val *EnumProviderManual) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumProviderManual) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumProviderManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumProviderManual(val *EnumProviderManual) *NullableEnumProviderManual {
	return &NullableEnumProviderManual{value: val, isSet: true}
}

func (v NullableEnumProviderManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumProviderManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
