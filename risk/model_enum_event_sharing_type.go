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

// EnumEventSharingType A string that specifies the device sharing type. Options are UNSPECIFIED, SHARED, and PRIVATE.
type EnumEventSharingType string

// List of EnumEventSharingType
const (
	ENUMEVENTSHARINGTYPE_UNSPECIFIED EnumEventSharingType = "UNSPECIFIED"
	ENUMEVENTSHARINGTYPE_SHARED EnumEventSharingType = "SHARED"
	ENUMEVENTSHARINGTYPE_PRIVATE EnumEventSharingType = "PRIVATE"
)

// All allowed values of EnumEventSharingType enum
var AllowedEnumEventSharingTypeEnumValues = []EnumEventSharingType{
	"UNSPECIFIED",
	"SHARED",
	"PRIVATE",
}

func (v *EnumEventSharingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumEventSharingType(value)
	for _, existing := range AllowedEnumEventSharingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumEventSharingType", value)
}

// NewEnumEventSharingTypeFromValue returns a pointer to a valid EnumEventSharingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumEventSharingTypeFromValue(v string) (*EnumEventSharingType, error) {
	ev := EnumEventSharingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumEventSharingType: valid values are %v", v, AllowedEnumEventSharingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumEventSharingType) IsValid() bool {
	for _, existing := range AllowedEnumEventSharingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumEventSharingType value
func (v EnumEventSharingType) Ptr() *EnumEventSharingType {
	return &v
}

type NullableEnumEventSharingType struct {
	value *EnumEventSharingType
	isSet bool
}

func (v NullableEnumEventSharingType) Get() *EnumEventSharingType {
	return v.value
}

func (v *NullableEnumEventSharingType) Set(val *EnumEventSharingType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumEventSharingType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumEventSharingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumEventSharingType(val *EnumEventSharingType) *NullableEnumEventSharingType {
	return &NullableEnumEventSharingType{value: val, isSet: true}
}

func (v NullableEnumEventSharingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumEventSharingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

