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

// checks if the RiskEvaluationDetailsUserRiskBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationDetailsUserRiskBehavior{}

// RiskEvaluationDetailsUserRiskBehavior struct for RiskEvaluationDetailsUserRiskBehavior
type RiskEvaluationDetailsUserRiskBehavior struct {
	Level *EnumRiskLevel `json:"level,omitempty"`
	// A string that describes the reason (or reasons) provided for the user behavior risk score classification within the organization (for example, the operating system or browser type used by the device, and country in which the accessing device is located). Each reason is classified as Unusual or Very Unusual, to indicate how much it deviates from normal user behavior, and its effect in calculating the overall user behavior risk score.
	Reason *string `json:"reason,omitempty"`
}

// NewRiskEvaluationDetailsUserRiskBehavior instantiates a new RiskEvaluationDetailsUserRiskBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsUserRiskBehavior() *RiskEvaluationDetailsUserRiskBehavior {
	this := RiskEvaluationDetailsUserRiskBehavior{}
	return &this
}

// NewRiskEvaluationDetailsUserRiskBehaviorWithDefaults instantiates a new RiskEvaluationDetailsUserRiskBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsUserRiskBehaviorWithDefaults() *RiskEvaluationDetailsUserRiskBehavior {
	this := RiskEvaluationDetailsUserRiskBehavior{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserRiskBehavior) GetLevel() EnumRiskLevel {
	if o == nil || IsNil(o.Level) {
		var ret EnumRiskLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserRiskBehavior) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserRiskBehavior) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given EnumRiskLevel and assigns it to the Level field.
func (o *RiskEvaluationDetailsUserRiskBehavior) SetLevel(v EnumRiskLevel) {
	o.Level = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserRiskBehavior) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserRiskBehavior) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserRiskBehavior) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RiskEvaluationDetailsUserRiskBehavior) SetReason(v string) {
	o.Reason = &v
}

func (o RiskEvaluationDetailsUserRiskBehavior) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationDetailsUserRiskBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableRiskEvaluationDetailsUserRiskBehavior struct {
	value *RiskEvaluationDetailsUserRiskBehavior
	isSet bool
}

func (v NullableRiskEvaluationDetailsUserRiskBehavior) Get() *RiskEvaluationDetailsUserRiskBehavior {
	return v.value
}

func (v *NullableRiskEvaluationDetailsUserRiskBehavior) Set(val *RiskEvaluationDetailsUserRiskBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsUserRiskBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsUserRiskBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsUserRiskBehavior(val *RiskEvaluationDetailsUserRiskBehavior) *NullableRiskEvaluationDetailsUserRiskBehavior {
	return &NullableRiskEvaluationDetailsUserRiskBehavior{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsUserRiskBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsUserRiskBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


