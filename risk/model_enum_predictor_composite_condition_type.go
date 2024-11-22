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

// EnumPredictorCompositeConditionType the model 'EnumPredictorCompositeConditionType'
type EnumPredictorCompositeConditionType string

// List of EnumPredictorCompositeConditionType
const (
	ENUMPREDICTORCOMPOSITECONDITIONTYPE_VALUE_COMPARISON EnumPredictorCompositeConditionType = "VALUE_COMPARISON"
	ENUMPREDICTORCOMPOSITECONDITIONTYPE_STRING_LIST      EnumPredictorCompositeConditionType = "STRING_LIST"
	ENUMPREDICTORCOMPOSITECONDITIONTYPE_AND              EnumPredictorCompositeConditionType = "AND"
	ENUMPREDICTORCOMPOSITECONDITIONTYPE_NOT              EnumPredictorCompositeConditionType = "NOT"
	ENUMPREDICTORCOMPOSITECONDITIONTYPE_OR               EnumPredictorCompositeConditionType = "OR"
)

// All allowed values of EnumPredictorCompositeConditionType enum
var AllowedEnumPredictorCompositeConditionTypeEnumValues = []EnumPredictorCompositeConditionType{
	"VALUE_COMPARISON",
	"STRING_LIST",
	"AND",
	"NOT",
	"OR",
}

func (v *EnumPredictorCompositeConditionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPredictorCompositeConditionType(value)
	for _, existing := range AllowedEnumPredictorCompositeConditionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumPredictorCompositeConditionType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumPredictorCompositeConditionTypeFromValue returns a pointer to a valid EnumPredictorCompositeConditionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPredictorCompositeConditionTypeFromValue(v string) (*EnumPredictorCompositeConditionType, error) {
	ev := EnumPredictorCompositeConditionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPredictorCompositeConditionType: valid values are %v", v, AllowedEnumPredictorCompositeConditionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPredictorCompositeConditionType) IsValid() bool {
	for _, existing := range AllowedEnumPredictorCompositeConditionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPredictorCompositeConditionType value
func (v EnumPredictorCompositeConditionType) Ptr() *EnumPredictorCompositeConditionType {
	return &v
}

type NullableEnumPredictorCompositeConditionType struct {
	value *EnumPredictorCompositeConditionType
	isSet bool
}

func (v NullableEnumPredictorCompositeConditionType) Get() *EnumPredictorCompositeConditionType {
	return v.value
}

func (v *NullableEnumPredictorCompositeConditionType) Set(val *EnumPredictorCompositeConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPredictorCompositeConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPredictorCompositeConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPredictorCompositeConditionType(val *EnumPredictorCompositeConditionType) *NullableEnumPredictorCompositeConditionType {
	return &NullableEnumPredictorCompositeConditionType{value: val, isSet: true}
}

func (v NullableEnumPredictorCompositeConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPredictorCompositeConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
