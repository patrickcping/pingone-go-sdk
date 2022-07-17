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

// EnumFlowType A string that specifies the flow type for the event. The only option (and default) is AUTHENTICATION.
type EnumFlowType string

// List of EnumFlowType
const (
	AUTHENTICATION EnumFlowType = "AUTHENTICATION"
)

// All allowed values of EnumFlowType enum
var AllowedEnumFlowTypeEnumValues = []EnumFlowType{
	"AUTHENTICATION",
}

func (v *EnumFlowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFlowType(value)
	for _, existing := range AllowedEnumFlowTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFlowType", value)
}

// NewEnumFlowTypeFromValue returns a pointer to a valid EnumFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFlowTypeFromValue(v string) (*EnumFlowType, error) {
	ev := EnumFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFlowType: valid values are %v", v, AllowedEnumFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFlowType) IsValid() bool {
	for _, existing := range AllowedEnumFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFlowType value
func (v EnumFlowType) Ptr() *EnumFlowType {
	return &v
}

type NullableEnumFlowType struct {
	value *EnumFlowType
	isSet bool
}

func (v NullableEnumFlowType) Get() *EnumFlowType {
	return v.value
}

func (v *NullableEnumFlowType) Set(val *EnumFlowType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFlowType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFlowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFlowType(val *EnumFlowType) *NullableEnumFlowType {
	return &NullableEnumFlowType{value: val, isSet: true}
}

func (v NullableEnumFlowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFlowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

