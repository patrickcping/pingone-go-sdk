/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumPredictorVelocityUseType the model 'EnumPredictorVelocityUseType'
type EnumPredictorVelocityUseType string

// List of EnumPredictorVelocityUseType
const (
	ENUMPREDICTORVELOCITYUSETYPE_POISSON_WITH_MAX EnumPredictorVelocityUseType = "POISSON_WITH_MAX"
)

// All allowed values of EnumPredictorVelocityUseType enum
var AllowedEnumPredictorVelocityUseTypeEnumValues = []EnumPredictorVelocityUseType{
	"POISSON_WITH_MAX",
}

func (v *EnumPredictorVelocityUseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPredictorVelocityUseType(value)
	for _, existing := range AllowedEnumPredictorVelocityUseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumPredictorVelocityUseType", value)
}

// NewEnumPredictorVelocityUseTypeFromValue returns a pointer to a valid EnumPredictorVelocityUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPredictorVelocityUseTypeFromValue(v string) (*EnumPredictorVelocityUseType, error) {
	ev := EnumPredictorVelocityUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPredictorVelocityUseType: valid values are %v", v, AllowedEnumPredictorVelocityUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPredictorVelocityUseType) IsValid() bool {
	for _, existing := range AllowedEnumPredictorVelocityUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPredictorVelocityUseType value
func (v EnumPredictorVelocityUseType) Ptr() *EnumPredictorVelocityUseType {
	return &v
}

type NullableEnumPredictorVelocityUseType struct {
	value *EnumPredictorVelocityUseType
	isSet bool
}

func (v NullableEnumPredictorVelocityUseType) Get() *EnumPredictorVelocityUseType {
	return v.value
}

func (v *NullableEnumPredictorVelocityUseType) Set(val *EnumPredictorVelocityUseType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPredictorVelocityUseType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPredictorVelocityUseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPredictorVelocityUseType(val *EnumPredictorVelocityUseType) *NullableEnumPredictorVelocityUseType {
	return &NullableEnumPredictorVelocityUseType{value: val, isSet: true}
}

func (v NullableEnumPredictorVelocityUseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPredictorVelocityUseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

