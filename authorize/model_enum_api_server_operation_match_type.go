/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// EnumAPIServerOperationMatchType An enumeration defining the match type of the scope rule. `ALL` means the client must be authorized with all scopes configured in the `scopes` array to obtain access. `ANY` means the client must be authorized with one or more of the scopes.
type EnumAPIServerOperationMatchType string

// List of EnumAPIServerOperationMatchType
const (
	ENUMAPISERVEROPERATIONMATCHTYPE_ANY EnumAPIServerOperationMatchType = "ANY"
	ENUMAPISERVEROPERATIONMATCHTYPE_ALL EnumAPIServerOperationMatchType = "ALL"
)

// All allowed values of EnumAPIServerOperationMatchType enum
var AllowedEnumAPIServerOperationMatchTypeEnumValues = []EnumAPIServerOperationMatchType{
	"ANY",
	"ALL",
}

func (v *EnumAPIServerOperationMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAPIServerOperationMatchType(value)
	for _, existing := range AllowedEnumAPIServerOperationMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAPIServerOperationMatchType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAPIServerOperationMatchTypeFromValue returns a pointer to a valid EnumAPIServerOperationMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAPIServerOperationMatchTypeFromValue(v string) (*EnumAPIServerOperationMatchType, error) {
	ev := EnumAPIServerOperationMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAPIServerOperationMatchType: valid values are %v", v, AllowedEnumAPIServerOperationMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAPIServerOperationMatchType) IsValid() bool {
	for _, existing := range AllowedEnumAPIServerOperationMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAPIServerOperationMatchType value
func (v EnumAPIServerOperationMatchType) Ptr() *EnumAPIServerOperationMatchType {
	return &v
}

type NullableEnumAPIServerOperationMatchType struct {
	value *EnumAPIServerOperationMatchType
	isSet bool
}

func (v NullableEnumAPIServerOperationMatchType) Get() *EnumAPIServerOperationMatchType {
	return v.value
}

func (v *NullableEnumAPIServerOperationMatchType) Set(val *EnumAPIServerOperationMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAPIServerOperationMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAPIServerOperationMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAPIServerOperationMatchType(val *EnumAPIServerOperationMatchType) *NullableEnumAPIServerOperationMatchType {
	return &NullableEnumAPIServerOperationMatchType{value: val, isSet: true}
}

func (v NullableEnumAPIServerOperationMatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAPIServerOperationMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
