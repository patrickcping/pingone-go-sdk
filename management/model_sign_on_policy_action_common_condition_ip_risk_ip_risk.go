/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the SignOnPolicyActionCommonConditionIPRiskIpRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionCommonConditionIPRiskIpRisk{}

// SignOnPolicyActionCommonConditionIPRiskIpRisk struct for SignOnPolicyActionCommonConditionIPRiskIpRisk
type SignOnPolicyActionCommonConditionIPRiskIpRisk struct {
	MinScore int32 `json:"minScore"`
	MaxScore int32 `json:"maxScore"`
}

// NewSignOnPolicyActionCommonConditionIPRiskIpRisk instantiates a new SignOnPolicyActionCommonConditionIPRiskIpRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonConditionIPRiskIpRisk(minScore int32, maxScore int32) *SignOnPolicyActionCommonConditionIPRiskIpRisk {
	this := SignOnPolicyActionCommonConditionIPRiskIpRisk{}
	this.MinScore = minScore
	this.MaxScore = maxScore
	return &this
}

// NewSignOnPolicyActionCommonConditionIPRiskIpRiskWithDefaults instantiates a new SignOnPolicyActionCommonConditionIPRiskIpRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonConditionIPRiskIpRiskWithDefaults() *SignOnPolicyActionCommonConditionIPRiskIpRisk {
	this := SignOnPolicyActionCommonConditionIPRiskIpRisk{}
	return &this
}

// GetMinScore returns the MinScore field value
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) GetMinScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinScore
}

// GetMinScoreOk returns a tuple with the MinScore field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) GetMinScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinScore, true
}

// SetMinScore sets field value
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) SetMinScore(v int32) {
	o.MinScore = v
}

// GetMaxScore returns the MaxScore field value
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) GetMaxScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) GetMaxScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxScore, true
}

// SetMaxScore sets field value
func (o *SignOnPolicyActionCommonConditionIPRiskIpRisk) SetMaxScore(v int32) {
	o.MaxScore = v
}

func (o SignOnPolicyActionCommonConditionIPRiskIpRisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionCommonConditionIPRiskIpRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minScore"] = o.MinScore
	toSerialize["maxScore"] = o.MaxScore
	return toSerialize, nil
}

type NullableSignOnPolicyActionCommonConditionIPRiskIpRisk struct {
	value *SignOnPolicyActionCommonConditionIPRiskIpRisk
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) Get() *SignOnPolicyActionCommonConditionIPRiskIpRisk {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) Set(val *SignOnPolicyActionCommonConditionIPRiskIpRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionIPRiskIpRisk(val *SignOnPolicyActionCommonConditionIPRiskIpRisk) *NullableSignOnPolicyActionCommonConditionIPRiskIpRisk {
	return &NullableSignOnPolicyActionCommonConditionIPRiskIpRisk{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionIPRiskIpRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


