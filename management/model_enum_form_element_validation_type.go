/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumFormElementValidationType A string that specifies the validation type. This is a required property when the field type is `TEXT`.
type EnumFormElementValidationType string

// List of EnumFormElementValidationType
const (
	ENUMFORMELEMENTVALIDATIONTYPE_NONE EnumFormElementValidationType = "NONE"
	ENUMFORMELEMENTVALIDATIONTYPE_CUSTOM EnumFormElementValidationType = "CUSTOM"
)

// All allowed values of EnumFormElementValidationType enum
var AllowedEnumFormElementValidationTypeEnumValues = []EnumFormElementValidationType{
	"NONE",
	"CUSTOM",
}

func (v *EnumFormElementValidationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFormElementValidationType(value)
	for _, existing := range AllowedEnumFormElementValidationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFormElementValidationType", value)
}

// NewEnumFormElementValidationTypeFromValue returns a pointer to a valid EnumFormElementValidationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFormElementValidationTypeFromValue(v string) (*EnumFormElementValidationType, error) {
	ev := EnumFormElementValidationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFormElementValidationType: valid values are %v", v, AllowedEnumFormElementValidationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFormElementValidationType) IsValid() bool {
	for _, existing := range AllowedEnumFormElementValidationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFormElementValidationType value
func (v EnumFormElementValidationType) Ptr() *EnumFormElementValidationType {
	return &v
}

type NullableEnumFormElementValidationType struct {
	value *EnumFormElementValidationType
	isSet bool
}

func (v NullableEnumFormElementValidationType) Get() *EnumFormElementValidationType {
	return v.value
}

func (v *NullableEnumFormElementValidationType) Set(val *EnumFormElementValidationType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFormElementValidationType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFormElementValidationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFormElementValidationType(val *EnumFormElementValidationType) *NullableEnumFormElementValidationType {
	return &NullableEnumFormElementValidationType{value: val, isSet: true}
}

func (v NullableEnumFormElementValidationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFormElementValidationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

