/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RiskPolicySetTriggersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicySetTriggersInner{}

// RiskPolicySetTriggersInner struct for RiskPolicySetTriggersInner
type RiskPolicySetTriggersInner struct {
	Type      EnumRiskPolicySetTriggerType          `json:"type"`
	PolicySet RiskPolicySetEvaluatedPredictorsInner `json:"policySet"`
	// The time the trigger expires (format ISO-8061).
	ExpiresAt time.Time `json:"expiresAt"`
}

type _RiskPolicySetTriggersInner RiskPolicySetTriggersInner

// NewRiskPolicySetTriggersInner instantiates a new RiskPolicySetTriggersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySetTriggersInner(type_ EnumRiskPolicySetTriggerType, policySet RiskPolicySetEvaluatedPredictorsInner, expiresAt time.Time) *RiskPolicySetTriggersInner {
	this := RiskPolicySetTriggersInner{}
	this.Type = type_
	this.PolicySet = policySet
	this.ExpiresAt = expiresAt
	return &this
}

// NewRiskPolicySetTriggersInnerWithDefaults instantiates a new RiskPolicySetTriggersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetTriggersInnerWithDefaults() *RiskPolicySetTriggersInner {
	this := RiskPolicySetTriggersInner{}
	return &this
}

// GetType returns the Type field value
func (o *RiskPolicySetTriggersInner) GetType() EnumRiskPolicySetTriggerType {
	if o == nil {
		var ret EnumRiskPolicySetTriggerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetTriggersInner) GetTypeOk() (*EnumRiskPolicySetTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPolicySetTriggersInner) SetType(v EnumRiskPolicySetTriggerType) {
	o.Type = v
}

// GetPolicySet returns the PolicySet field value
func (o *RiskPolicySetTriggersInner) GetPolicySet() RiskPolicySetEvaluatedPredictorsInner {
	if o == nil {
		var ret RiskPolicySetEvaluatedPredictorsInner
		return ret
	}

	return o.PolicySet
}

// GetPolicySetOk returns a tuple with the PolicySet field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetTriggersInner) GetPolicySetOk() (*RiskPolicySetEvaluatedPredictorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicySet, true
}

// SetPolicySet sets field value
func (o *RiskPolicySetTriggersInner) SetPolicySet(v RiskPolicySetEvaluatedPredictorsInner) {
	o.PolicySet = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *RiskPolicySetTriggersInner) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetTriggersInner) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *RiskPolicySetTriggersInner) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o RiskPolicySetTriggersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicySetTriggersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["policySet"] = o.PolicySet
	toSerialize["expiresAt"] = o.ExpiresAt
	return toSerialize, nil
}

func (o *RiskPolicySetTriggersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"policySet",
		"expiresAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskPolicySetTriggersInner := _RiskPolicySetTriggersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPolicySetTriggersInner)

	if err != nil {
		return err
	}

	*o = RiskPolicySetTriggersInner(varRiskPolicySetTriggersInner)

	return err
}

type NullableRiskPolicySetTriggersInner struct {
	value *RiskPolicySetTriggersInner
	isSet bool
}

func (v NullableRiskPolicySetTriggersInner) Get() *RiskPolicySetTriggersInner {
	return v.value
}

func (v *NullableRiskPolicySetTriggersInner) Set(val *RiskPolicySetTriggersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetTriggersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetTriggersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetTriggersInner(val *RiskPolicySetTriggersInner) *NullableRiskPolicySetTriggersInner {
	return &NullableRiskPolicySetTriggersInner{value: val, isSet: true}
}

func (v NullableRiskPolicySetTriggersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetTriggersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
