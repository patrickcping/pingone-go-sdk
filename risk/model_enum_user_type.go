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

// EnumUserType A string that specifies the type of user associated with the event. Options are `EXTERNAL`.
type EnumUserType string

// List of EnumUserType
const (
	ENUMUSERTYPE_EXTERNAL EnumUserType = "EXTERNAL"
)

// All allowed values of EnumUserType enum
var AllowedEnumUserTypeEnumValues = []EnumUserType{
	"EXTERNAL",
}

func (v *EnumUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserType(value)
	for _, existing := range AllowedEnumUserTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserTypeFromValue returns a pointer to a valid EnumUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserTypeFromValue(v string) (*EnumUserType, error) {
	ev := EnumUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserType: valid values are %v", v, AllowedEnumUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserType) IsValid() bool {
	for _, existing := range AllowedEnumUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserType value
func (v EnumUserType) Ptr() *EnumUserType {
	return &v
}

type NullableEnumUserType struct {
	value *EnumUserType
	isSet bool
}

func (v NullableEnumUserType) Get() *EnumUserType {
	return v.value
}

func (v *NullableEnumUserType) Set(val *EnumUserType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserType(val *EnumUserType) *NullableEnumUserType {
	return &NullableEnumUserType{value: val, isSet: true}
}

func (v NullableEnumUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
