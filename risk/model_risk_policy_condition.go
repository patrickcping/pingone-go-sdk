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

// checks if the RiskPolicyCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicyCondition{}

// RiskPolicyCondition Contains the condition logic that determines when a policy is evaluated to true and when it is evaluated to false.
type RiskPolicyCondition struct {
	Type     *EnumRiskPolicyConditionType `json:"type,omitempty"`
	Contains *string                      `json:"contains,omitempty"`
	IpRange  []string                     `json:"ipRange,omitempty"`
	Value    *string                      `json:"value,omitempty"`
	Equals   *RiskPolicyConditionEquals   `json:"equals,omitempty"`
	// Required for weights-based policies. The elements in the array are `value`-`weight` pairs, representing a weighting for the weighted average calculation that should be assigned to a specific predictor when it is determined that there is a high risk for the predictor.
	AggregatedWeights []RiskPolicyConditionAggregatedWeightsInner `json:"aggregatedWeights,omitempty"`
	// Required for score-based policies. The elements in the array are `value`-`score` pairs, representing the score that should be assigned to a specific predictor when it is determined that there is a high risk for the predictor.
	AggregatedScores []RiskPolicyConditionAggregatedScoresInner `json:"aggregatedScores,omitempty"`
	Between          *RiskPolicyConditionBetween                `json:"between,omitempty"`
}

// NewRiskPolicyCondition instantiates a new RiskPolicyCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicyCondition() *RiskPolicyCondition {
	this := RiskPolicyCondition{}
	return &this
}

// NewRiskPolicyConditionWithDefaults instantiates a new RiskPolicyCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicyConditionWithDefaults() *RiskPolicyCondition {
	this := RiskPolicyCondition{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetType() EnumRiskPolicyConditionType {
	if o == nil || IsNil(o.Type) {
		var ret EnumRiskPolicyConditionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetTypeOk() (*EnumRiskPolicyConditionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumRiskPolicyConditionType and assigns it to the Type field.
func (o *RiskPolicyCondition) SetType(v EnumRiskPolicyConditionType) {
	o.Type = &v
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetContains() string {
	if o == nil || IsNil(o.Contains) {
		var ret string
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetContainsOk() (*string, bool) {
	if o == nil || IsNil(o.Contains) {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasContains() bool {
	if o != nil && !IsNil(o.Contains) {
		return true
	}

	return false
}

// SetContains gets a reference to the given string and assigns it to the Contains field.
func (o *RiskPolicyCondition) SetContains(v string) {
	o.Contains = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetIpRange() []string {
	if o == nil || IsNil(o.IpRange) {
		var ret []string
		return ret
	}
	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetIpRangeOk() ([]string, bool) {
	if o == nil || IsNil(o.IpRange) {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasIpRange() bool {
	if o != nil && !IsNil(o.IpRange) {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given []string and assigns it to the IpRange field.
func (o *RiskPolicyCondition) SetIpRange(v []string) {
	o.IpRange = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskPolicyCondition) SetValue(v string) {
	o.Value = &v
}

// GetEquals returns the Equals field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetEquals() RiskPolicyConditionEquals {
	if o == nil || IsNil(o.Equals) {
		var ret RiskPolicyConditionEquals
		return ret
	}
	return *o.Equals
}

// GetEqualsOk returns a tuple with the Equals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetEqualsOk() (*RiskPolicyConditionEquals, bool) {
	if o == nil || IsNil(o.Equals) {
		return nil, false
	}
	return o.Equals, true
}

// HasEquals returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasEquals() bool {
	if o != nil && !IsNil(o.Equals) {
		return true
	}

	return false
}

// SetEquals gets a reference to the given RiskPolicyConditionEquals and assigns it to the Equals field.
func (o *RiskPolicyCondition) SetEquals(v RiskPolicyConditionEquals) {
	o.Equals = &v
}

// GetAggregatedWeights returns the AggregatedWeights field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetAggregatedWeights() []RiskPolicyConditionAggregatedWeightsInner {
	if o == nil || IsNil(o.AggregatedWeights) {
		var ret []RiskPolicyConditionAggregatedWeightsInner
		return ret
	}
	return o.AggregatedWeights
}

// GetAggregatedWeightsOk returns a tuple with the AggregatedWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetAggregatedWeightsOk() ([]RiskPolicyConditionAggregatedWeightsInner, bool) {
	if o == nil || IsNil(o.AggregatedWeights) {
		return nil, false
	}
	return o.AggregatedWeights, true
}

// HasAggregatedWeights returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasAggregatedWeights() bool {
	if o != nil && !IsNil(o.AggregatedWeights) {
		return true
	}

	return false
}

// SetAggregatedWeights gets a reference to the given []RiskPolicyConditionAggregatedWeightsInner and assigns it to the AggregatedWeights field.
func (o *RiskPolicyCondition) SetAggregatedWeights(v []RiskPolicyConditionAggregatedWeightsInner) {
	o.AggregatedWeights = v
}

// GetAggregatedScores returns the AggregatedScores field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetAggregatedScores() []RiskPolicyConditionAggregatedScoresInner {
	if o == nil || IsNil(o.AggregatedScores) {
		var ret []RiskPolicyConditionAggregatedScoresInner
		return ret
	}
	return o.AggregatedScores
}

// GetAggregatedScoresOk returns a tuple with the AggregatedScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetAggregatedScoresOk() ([]RiskPolicyConditionAggregatedScoresInner, bool) {
	if o == nil || IsNil(o.AggregatedScores) {
		return nil, false
	}
	return o.AggregatedScores, true
}

// HasAggregatedScores returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasAggregatedScores() bool {
	if o != nil && !IsNil(o.AggregatedScores) {
		return true
	}

	return false
}

// SetAggregatedScores gets a reference to the given []RiskPolicyConditionAggregatedScoresInner and assigns it to the AggregatedScores field.
func (o *RiskPolicyCondition) SetAggregatedScores(v []RiskPolicyConditionAggregatedScoresInner) {
	o.AggregatedScores = v
}

// GetBetween returns the Between field value if set, zero value otherwise.
func (o *RiskPolicyCondition) GetBetween() RiskPolicyConditionBetween {
	if o == nil || IsNil(o.Between) {
		var ret RiskPolicyConditionBetween
		return ret
	}
	return *o.Between
}

// GetBetweenOk returns a tuple with the Between field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicyCondition) GetBetweenOk() (*RiskPolicyConditionBetween, bool) {
	if o == nil || IsNil(o.Between) {
		return nil, false
	}
	return o.Between, true
}

// HasBetween returns a boolean if a field has been set.
func (o *RiskPolicyCondition) HasBetween() bool {
	if o != nil && !IsNil(o.Between) {
		return true
	}

	return false
}

// SetBetween gets a reference to the given RiskPolicyConditionBetween and assigns it to the Between field.
func (o *RiskPolicyCondition) SetBetween(v RiskPolicyConditionBetween) {
	o.Between = &v
}

func (o RiskPolicyCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicyCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Contains) {
		toSerialize["contains"] = o.Contains
	}
	if !IsNil(o.IpRange) {
		toSerialize["ipRange"] = o.IpRange
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Equals) {
		toSerialize["equals"] = o.Equals
	}
	if !IsNil(o.AggregatedWeights) {
		toSerialize["aggregatedWeights"] = o.AggregatedWeights
	}
	if !IsNil(o.AggregatedScores) {
		toSerialize["aggregatedScores"] = o.AggregatedScores
	}
	if !IsNil(o.Between) {
		toSerialize["between"] = o.Between
	}
	return toSerialize, nil
}

type NullableRiskPolicyCondition struct {
	value *RiskPolicyCondition
	isSet bool
}

func (v NullableRiskPolicyCondition) Get() *RiskPolicyCondition {
	return v.value
}

func (v *NullableRiskPolicyCondition) Set(val *RiskPolicyCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicyCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicyCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicyCondition(val *RiskPolicyCondition) *NullableRiskPolicyCondition {
	return &NullableRiskPolicyCondition{value: val, isSet: true}
}

func (v NullableRiskPolicyCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicyCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
