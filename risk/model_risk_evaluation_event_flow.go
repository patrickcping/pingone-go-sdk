/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// checks if the RiskEvaluationEventFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationEventFlow{}

// RiskEvaluationEventFlow struct for RiskEvaluationEventFlow
type RiskEvaluationEventFlow struct {
	Type *EnumFlowType `json:"type,omitempty"`
}

// NewRiskEvaluationEventFlow instantiates a new RiskEvaluationEventFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEventFlow() *RiskEvaluationEventFlow {
	this := RiskEvaluationEventFlow{}
	var type_ EnumFlowType = ENUMFLOWTYPE_AUTHENTICATION
	this.Type = &type_
	return &this
}

// NewRiskEvaluationEventFlowWithDefaults instantiates a new RiskEvaluationEventFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventFlowWithDefaults() *RiskEvaluationEventFlow {
	this := RiskEvaluationEventFlow{}
	var type_ EnumFlowType = ENUMFLOWTYPE_AUTHENTICATION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskEvaluationEventFlow) GetType() EnumFlowType {
	if o == nil || IsNil(o.Type) {
		var ret EnumFlowType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventFlow) GetTypeOk() (*EnumFlowType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskEvaluationEventFlow) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumFlowType and assigns it to the Type field.
func (o *RiskEvaluationEventFlow) SetType(v EnumFlowType) {
	o.Type = &v
}

func (o RiskEvaluationEventFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationEventFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRiskEvaluationEventFlow struct {
	value *RiskEvaluationEventFlow
	isSet bool
}

func (v NullableRiskEvaluationEventFlow) Get() *RiskEvaluationEventFlow {
	return v.value
}

func (v *NullableRiskEvaluationEventFlow) Set(val *RiskEvaluationEventFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEventFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEventFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEventFlow(val *RiskEvaluationEventFlow) *NullableRiskEvaluationEventFlow {
	return &NullableRiskEvaluationEventFlow{value: val, isSet: true}
}

func (v NullableRiskEvaluationEventFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEventFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


