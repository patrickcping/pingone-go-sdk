/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumRiskPolicyResultLevel Contains the default result level returned if none of the conditions in the policy are evaluated to true. At this time, the value must be `LOW`.
type EnumRiskPolicyResultLevel string

// List of EnumRiskPolicyResultLevel
const (
	ENUMRISKPOLICYRESULTLEVEL_LOW EnumRiskPolicyResultLevel = "LOW"
)

// All allowed values of EnumRiskPolicyResultLevel enum
var AllowedEnumRiskPolicyResultLevelEnumValues = []EnumRiskPolicyResultLevel{
	"LOW",
}

func (v *EnumRiskPolicyResultLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRiskPolicyResultLevel(value)
	for _, existing := range AllowedEnumRiskPolicyResultLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumRiskPolicyResultLevel", value)
}

// NewEnumRiskPolicyResultLevelFromValue returns a pointer to a valid EnumRiskPolicyResultLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRiskPolicyResultLevelFromValue(v string) (*EnumRiskPolicyResultLevel, error) {
	ev := EnumRiskPolicyResultLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRiskPolicyResultLevel: valid values are %v", v, AllowedEnumRiskPolicyResultLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRiskPolicyResultLevel) IsValid() bool {
	for _, existing := range AllowedEnumRiskPolicyResultLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRiskPolicyResultLevel value
func (v EnumRiskPolicyResultLevel) Ptr() *EnumRiskPolicyResultLevel {
	return &v
}

type NullableEnumRiskPolicyResultLevel struct {
	value *EnumRiskPolicyResultLevel
	isSet bool
}

func (v NullableEnumRiskPolicyResultLevel) Get() *EnumRiskPolicyResultLevel {
	return v.value
}

func (v *NullableEnumRiskPolicyResultLevel) Set(val *EnumRiskPolicyResultLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRiskPolicyResultLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRiskPolicyResultLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRiskPolicyResultLevel(val *EnumRiskPolicyResultLevel) *NullableEnumRiskPolicyResultLevel {
	return &NullableEnumRiskPolicyResultLevel{value: val, isSet: true}
}

func (v NullableEnumRiskPolicyResultLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRiskPolicyResultLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

