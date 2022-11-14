/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskPolicySetRiskPoliciesInnerConditionBetween struct for RiskPolicySetRiskPoliciesInnerConditionBetween
type RiskPolicySetRiskPoliciesInnerConditionBetween struct {
	MinScore int32 `json:"minScore"`
	MaxScore int32 `json:"maxScore"`
}

// NewRiskPolicySetRiskPoliciesInnerConditionBetween instantiates a new RiskPolicySetRiskPoliciesInnerConditionBetween object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySetRiskPoliciesInnerConditionBetween(minScore int32, maxScore int32) *RiskPolicySetRiskPoliciesInnerConditionBetween {
	this := RiskPolicySetRiskPoliciesInnerConditionBetween{}
	this.MinScore = minScore
	this.MaxScore = maxScore
	return &this
}

// NewRiskPolicySetRiskPoliciesInnerConditionBetweenWithDefaults instantiates a new RiskPolicySetRiskPoliciesInnerConditionBetween object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetRiskPoliciesInnerConditionBetweenWithDefaults() *RiskPolicySetRiskPoliciesInnerConditionBetween {
	this := RiskPolicySetRiskPoliciesInnerConditionBetween{}
	return &this
}

// GetMinScore returns the MinScore field value
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) GetMinScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinScore
}

// GetMinScoreOk returns a tuple with the MinScore field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) GetMinScoreOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MinScore, true
}

// SetMinScore sets field value
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) SetMinScore(v int32) {
	o.MinScore = v
}

// GetMaxScore returns the MaxScore field value
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) GetMaxScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) GetMaxScoreOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxScore, true
}

// SetMaxScore sets field value
func (o *RiskPolicySetRiskPoliciesInnerConditionBetween) SetMaxScore(v int32) {
	o.MaxScore = v
}

func (o RiskPolicySetRiskPoliciesInnerConditionBetween) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["minScore"] = o.MinScore
	}
	if true {
		toSerialize["maxScore"] = o.MaxScore
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPolicySetRiskPoliciesInnerConditionBetween struct {
	value *RiskPolicySetRiskPoliciesInnerConditionBetween
	isSet bool
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionBetween) Get() *RiskPolicySetRiskPoliciesInnerConditionBetween {
	return v.value
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionBetween) Set(val *RiskPolicySetRiskPoliciesInnerConditionBetween) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionBetween) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionBetween) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetRiskPoliciesInnerConditionBetween(val *RiskPolicySetRiskPoliciesInnerConditionBetween) *NullableRiskPolicySetRiskPoliciesInnerConditionBetween {
	return &NullableRiskPolicySetRiskPoliciesInnerConditionBetween{value: val, isSet: true}
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionBetween) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionBetween) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


