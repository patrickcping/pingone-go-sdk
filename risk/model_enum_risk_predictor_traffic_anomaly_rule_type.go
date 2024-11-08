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

// EnumRiskPredictorTrafficAnomalyRuleType The type of rule. Currently the only valid value is `UNIQUES_USERS_PER_DEVICE` - for tracking the number of unique users for a device over a defined timeframe.
type EnumRiskPredictorTrafficAnomalyRuleType string

// List of EnumRiskPredictorTrafficAnomalyRuleType
const (
	ENUMRISKPREDICTORTRAFFICANOMALYRULETYPE_UNIQUES_USERS_PER_DEVICE EnumRiskPredictorTrafficAnomalyRuleType = "UNIQUES_USERS_PER_DEVICE"
)

// All allowed values of EnumRiskPredictorTrafficAnomalyRuleType enum
var AllowedEnumRiskPredictorTrafficAnomalyRuleTypeEnumValues = []EnumRiskPredictorTrafficAnomalyRuleType{
	"UNIQUES_USERS_PER_DEVICE",
}

func (v *EnumRiskPredictorTrafficAnomalyRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRiskPredictorTrafficAnomalyRuleType(value)
	for _, existing := range AllowedEnumRiskPredictorTrafficAnomalyRuleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumRiskPredictorTrafficAnomalyRuleType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumRiskPredictorTrafficAnomalyRuleTypeFromValue returns a pointer to a valid EnumRiskPredictorTrafficAnomalyRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRiskPredictorTrafficAnomalyRuleTypeFromValue(v string) (*EnumRiskPredictorTrafficAnomalyRuleType, error) {
	ev := EnumRiskPredictorTrafficAnomalyRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRiskPredictorTrafficAnomalyRuleType: valid values are %v", v, AllowedEnumRiskPredictorTrafficAnomalyRuleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRiskPredictorTrafficAnomalyRuleType) IsValid() bool {
	for _, existing := range AllowedEnumRiskPredictorTrafficAnomalyRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRiskPredictorTrafficAnomalyRuleType value
func (v EnumRiskPredictorTrafficAnomalyRuleType) Ptr() *EnumRiskPredictorTrafficAnomalyRuleType {
	return &v
}

type NullableEnumRiskPredictorTrafficAnomalyRuleType struct {
	value *EnumRiskPredictorTrafficAnomalyRuleType
	isSet bool
}

func (v NullableEnumRiskPredictorTrafficAnomalyRuleType) Get() *EnumRiskPredictorTrafficAnomalyRuleType {
	return v.value
}

func (v *NullableEnumRiskPredictorTrafficAnomalyRuleType) Set(val *EnumRiskPredictorTrafficAnomalyRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRiskPredictorTrafficAnomalyRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRiskPredictorTrafficAnomalyRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRiskPredictorTrafficAnomalyRuleType(val *EnumRiskPredictorTrafficAnomalyRuleType) *NullableEnumRiskPredictorTrafficAnomalyRuleType {
	return &NullableEnumRiskPredictorTrafficAnomalyRuleType{value: val, isSet: true}
}

func (v NullableEnumRiskPredictorTrafficAnomalyRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRiskPredictorTrafficAnomalyRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

