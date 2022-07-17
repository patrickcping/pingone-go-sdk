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

// EnumPredictorType An enum type. This can be either VELOCITY, USER_RISK_BEHAVIOR, or MAP
type EnumPredictorType string

// List of EnumPredictorType
const (
	ENUMPREDICTORTYPE_VELOCITY EnumPredictorType = "VELOCITY"
	ENUMPREDICTORTYPE_USER_RISK_BEHAVIOR EnumPredictorType = "USER_RISK_BEHAVIOR"
	ENUMPREDICTORTYPE_MAP EnumPredictorType = "MAP"
)

// All allowed values of EnumPredictorType enum
var AllowedEnumPredictorTypeEnumValues = []EnumPredictorType{
	"VELOCITY",
	"USER_RISK_BEHAVIOR",
	"MAP",
}

func (v *EnumPredictorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPredictorType(value)
	for _, existing := range AllowedEnumPredictorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumPredictorType", value)
}

// NewEnumPredictorTypeFromValue returns a pointer to a valid EnumPredictorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPredictorTypeFromValue(v string) (*EnumPredictorType, error) {
	ev := EnumPredictorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPredictorType: valid values are %v", v, AllowedEnumPredictorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPredictorType) IsValid() bool {
	for _, existing := range AllowedEnumPredictorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPredictorType value
func (v EnumPredictorType) Ptr() *EnumPredictorType {
	return &v
}

type NullableEnumPredictorType struct {
	value *EnumPredictorType
	isSet bool
}

func (v NullableEnumPredictorType) Get() *EnumPredictorType {
	return v.value
}

func (v *NullableEnumPredictorType) Set(val *EnumPredictorType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPredictorType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPredictorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPredictorType(val *EnumPredictorType) *NullableEnumPredictorType {
	return &NullableEnumPredictorType{value: val, isSet: true}
}

func (v NullableEnumPredictorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPredictorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

