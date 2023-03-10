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

// EnumFlowPolicyTriggerType the model 'EnumFlowPolicyTriggerType'
type EnumFlowPolicyTriggerType string

// List of EnumFlowPolicyTriggerType
const (
	ENUMFLOWPOLICYTRIGGERTYPE_AUTHENTICATION EnumFlowPolicyTriggerType = "AUTHENTICATION"
)

// All allowed values of EnumFlowPolicyTriggerType enum
var AllowedEnumFlowPolicyTriggerTypeEnumValues = []EnumFlowPolicyTriggerType{
	"AUTHENTICATION",
}

func (v *EnumFlowPolicyTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFlowPolicyTriggerType(value)
	for _, existing := range AllowedEnumFlowPolicyTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFlowPolicyTriggerType", value)
}

// NewEnumFlowPolicyTriggerTypeFromValue returns a pointer to a valid EnumFlowPolicyTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFlowPolicyTriggerTypeFromValue(v string) (*EnumFlowPolicyTriggerType, error) {
	ev := EnumFlowPolicyTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFlowPolicyTriggerType: valid values are %v", v, AllowedEnumFlowPolicyTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFlowPolicyTriggerType) IsValid() bool {
	for _, existing := range AllowedEnumFlowPolicyTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFlowPolicyTriggerType value
func (v EnumFlowPolicyTriggerType) Ptr() *EnumFlowPolicyTriggerType {
	return &v
}

type NullableEnumFlowPolicyTriggerType struct {
	value *EnumFlowPolicyTriggerType
	isSet bool
}

func (v NullableEnumFlowPolicyTriggerType) Get() *EnumFlowPolicyTriggerType {
	return v.value
}

func (v *NullableEnumFlowPolicyTriggerType) Set(val *EnumFlowPolicyTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFlowPolicyTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFlowPolicyTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFlowPolicyTriggerType(val *EnumFlowPolicyTriggerType) *NullableEnumFlowPolicyTriggerType {
	return &NullableEnumFlowPolicyTriggerType{value: val, isSet: true}
}

func (v NullableEnumFlowPolicyTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFlowPolicyTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

