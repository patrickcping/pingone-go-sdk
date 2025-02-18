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

// EnumAuthorizeEditorDataEntityTestingOverrideDTOType the model 'EnumAuthorizeEditorDataEntityTestingOverrideDTOType'
type EnumAuthorizeEditorDataEntityTestingOverrideDTOType string

// List of EnumAuthorizeEditorDataEntityTestingOverrideDTOType
const (
	ENUMAUTHORIZEEDITORDATAENTITYTESTINGOVERRIDEDTOTYPE_ATTRIBUTE EnumAuthorizeEditorDataEntityTestingOverrideDTOType = "ATTRIBUTE"
	ENUMAUTHORIZEEDITORDATAENTITYTESTINGOVERRIDEDTOTYPE_SERVICE EnumAuthorizeEditorDataEntityTestingOverrideDTOType = "SERVICE"
)

// All allowed values of EnumAuthorizeEditorDataEntityTestingOverrideDTOType enum
var AllowedEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeEnumValues = []EnumAuthorizeEditorDataEntityTestingOverrideDTOType{
	"ATTRIBUTE",
	"SERVICE",
}

func (v *EnumAuthorizeEditorDataEntityTestingOverrideDTOType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataEntityTestingOverrideDTOType(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataEntityTestingOverrideDTOType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeFromValue returns a pointer to a valid EnumAuthorizeEditorDataEntityTestingOverrideDTOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeFromValue(v string) (*EnumAuthorizeEditorDataEntityTestingOverrideDTOType, error) {
	ev := EnumAuthorizeEditorDataEntityTestingOverrideDTOType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataEntityTestingOverrideDTOType: valid values are %v", v, AllowedEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataEntityTestingOverrideDTOType) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataEntityTestingOverrideDTOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataEntityTestingOverrideDTOType value
func (v EnumAuthorizeEditorDataEntityTestingOverrideDTOType) Ptr() *EnumAuthorizeEditorDataEntityTestingOverrideDTOType {
	return &v
}

type NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType struct {
	value *EnumAuthorizeEditorDataEntityTestingOverrideDTOType
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) Get() *EnumAuthorizeEditorDataEntityTestingOverrideDTOType {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) Set(val *EnumAuthorizeEditorDataEntityTestingOverrideDTOType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType(val *EnumAuthorizeEditorDataEntityTestingOverrideDTOType) *NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType {
	return &NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataEntityTestingOverrideDTOType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

