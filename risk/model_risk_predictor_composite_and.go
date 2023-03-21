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

// checks if the RiskPredictorCompositeAnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCompositeAnd{}

// RiskPredictorCompositeAnd struct for RiskPredictorCompositeAnd
type RiskPredictorCompositeAnd struct {
	And []RiskPredictorCompositeCondition `json:"and"`
}

// NewRiskPredictorCompositeAnd instantiates a new RiskPredictorCompositeAnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCompositeAnd(and []RiskPredictorCompositeCondition) *RiskPredictorCompositeAnd {
	this := RiskPredictorCompositeAnd{}
	this.And = and
	return &this
}

// NewRiskPredictorCompositeAndWithDefaults instantiates a new RiskPredictorCompositeAnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCompositeAndWithDefaults() *RiskPredictorCompositeAnd {
	this := RiskPredictorCompositeAnd{}
	return &this
}

// GetAnd returns the And field value
func (o *RiskPredictorCompositeAnd) GetAnd() []RiskPredictorCompositeCondition {
	if o == nil {
		var ret []RiskPredictorCompositeCondition
		return ret
	}

	return o.And
}

// GetAndOk returns a tuple with the And field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeAnd) GetAndOk() ([]RiskPredictorCompositeCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.And, true
}

// SetAnd sets field value
func (o *RiskPredictorCompositeAnd) SetAnd(v []RiskPredictorCompositeCondition) {
	o.And = v
}

func (o RiskPredictorCompositeAnd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCompositeAnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["and"] = o.And
	return toSerialize, nil
}

type NullableRiskPredictorCompositeAnd struct {
	value *RiskPredictorCompositeAnd
	isSet bool
}

func (v NullableRiskPredictorCompositeAnd) Get() *RiskPredictorCompositeAnd {
	return v.value
}

func (v *NullableRiskPredictorCompositeAnd) Set(val *RiskPredictorCompositeAnd) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeAnd) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeAnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeAnd(val *RiskPredictorCompositeAnd) *NullableRiskPredictorCompositeAnd {
	return &NullableRiskPredictorCompositeAnd{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeAnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeAnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


