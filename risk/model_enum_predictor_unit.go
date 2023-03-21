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

// EnumPredictorUnit the model 'EnumPredictorUnit'
type EnumPredictorUnit string

// List of EnumPredictorUnit
const (
	ENUMPREDICTORUNIT_DAY EnumPredictorUnit = "DAY"
	ENUMPREDICTORUNIT_HOUR EnumPredictorUnit = "HOUR"
)

// All allowed values of EnumPredictorUnit enum
var AllowedEnumPredictorUnitEnumValues = []EnumPredictorUnit{
	"DAY",
	"HOUR",
}

func (v *EnumPredictorUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPredictorUnit(value)
	for _, existing := range AllowedEnumPredictorUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumPredictorUnit", value)
}

// NewEnumPredictorUnitFromValue returns a pointer to a valid EnumPredictorUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPredictorUnitFromValue(v string) (*EnumPredictorUnit, error) {
	ev := EnumPredictorUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPredictorUnit: valid values are %v", v, AllowedEnumPredictorUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPredictorUnit) IsValid() bool {
	for _, existing := range AllowedEnumPredictorUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPredictorUnit value
func (v EnumPredictorUnit) Ptr() *EnumPredictorUnit {
	return &v
}

type NullableEnumPredictorUnit struct {
	value *EnumPredictorUnit
	isSet bool
}

func (v NullableEnumPredictorUnit) Get() *EnumPredictorUnit {
	return v.value
}

func (v *NullableEnumPredictorUnit) Set(val *EnumPredictorUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPredictorUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPredictorUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPredictorUnit(val *EnumPredictorUnit) *NullableEnumPredictorUnit {
	return &NullableEnumPredictorUnit{value: val, isSet: true}
}

func (v NullableEnumPredictorUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPredictorUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

