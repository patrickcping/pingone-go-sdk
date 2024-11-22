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

// EnumUserRiskBehaviorRiskModel Enum string that specifies the risk model.  Options are `points` (individual user model) and `login_anomaly_statistic` (organization wide model).
type EnumUserRiskBehaviorRiskModel string

// List of EnumUserRiskBehaviorRiskModel
const (
	ENUMUSERRISKBEHAVIORRISKMODEL_POINTS                  EnumUserRiskBehaviorRiskModel = "points"
	ENUMUSERRISKBEHAVIORRISKMODEL_LOGIN_ANOMALY_STATISTIC EnumUserRiskBehaviorRiskModel = "login_anomaly_statistic"
)

// All allowed values of EnumUserRiskBehaviorRiskModel enum
var AllowedEnumUserRiskBehaviorRiskModelEnumValues = []EnumUserRiskBehaviorRiskModel{
	"points",
	"login_anomaly_statistic",
}

func (v *EnumUserRiskBehaviorRiskModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserRiskBehaviorRiskModel(value)
	for _, existing := range AllowedEnumUserRiskBehaviorRiskModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserRiskBehaviorRiskModel(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserRiskBehaviorRiskModelFromValue returns a pointer to a valid EnumUserRiskBehaviorRiskModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserRiskBehaviorRiskModelFromValue(v string) (*EnumUserRiskBehaviorRiskModel, error) {
	ev := EnumUserRiskBehaviorRiskModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserRiskBehaviorRiskModel: valid values are %v", v, AllowedEnumUserRiskBehaviorRiskModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserRiskBehaviorRiskModel) IsValid() bool {
	for _, existing := range AllowedEnumUserRiskBehaviorRiskModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserRiskBehaviorRiskModel value
func (v EnumUserRiskBehaviorRiskModel) Ptr() *EnumUserRiskBehaviorRiskModel {
	return &v
}

type NullableEnumUserRiskBehaviorRiskModel struct {
	value *EnumUserRiskBehaviorRiskModel
	isSet bool
}

func (v NullableEnumUserRiskBehaviorRiskModel) Get() *EnumUserRiskBehaviorRiskModel {
	return v.value
}

func (v *NullableEnumUserRiskBehaviorRiskModel) Set(val *EnumUserRiskBehaviorRiskModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserRiskBehaviorRiskModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserRiskBehaviorRiskModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserRiskBehaviorRiskModel(val *EnumUserRiskBehaviorRiskModel) *NullableEnumUserRiskBehaviorRiskModel {
	return &NullableEnumUserRiskBehaviorRiskModel{value: val, isSet: true}
}

func (v NullableEnumUserRiskBehaviorRiskModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserRiskBehaviorRiskModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
