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

// SignOnPolicyActionCommonConditionEquals struct for SignOnPolicyActionCommonConditionEquals
type SignOnPolicyActionCommonConditionEquals struct {
	Value string `json:"value"`
	Equals SignOnPolicyActionCommonConditionEqualsEquals `json:"equals"`
}

// NewSignOnPolicyActionCommonConditionEquals instantiates a new SignOnPolicyActionCommonConditionEquals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonConditionEquals(value string, equals SignOnPolicyActionCommonConditionEqualsEquals) *SignOnPolicyActionCommonConditionEquals {
	this := SignOnPolicyActionCommonConditionEquals{}
	this.Value = value
	this.Equals = equals
	return &this
}

// NewSignOnPolicyActionCommonConditionEqualsWithDefaults instantiates a new SignOnPolicyActionCommonConditionEquals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonConditionEqualsWithDefaults() *SignOnPolicyActionCommonConditionEquals {
	this := SignOnPolicyActionCommonConditionEquals{}
	return &this
}

// GetValue returns the Value field value
func (o *SignOnPolicyActionCommonConditionEquals) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionEquals) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SignOnPolicyActionCommonConditionEquals) SetValue(v string) {
	o.Value = v
}

// GetEquals returns the Equals field value
func (o *SignOnPolicyActionCommonConditionEquals) GetEquals() SignOnPolicyActionCommonConditionEqualsEquals {
	if o == nil {
		var ret SignOnPolicyActionCommonConditionEqualsEquals
		return ret
	}

	return o.Equals
}

// GetEqualsOk returns a tuple with the Equals field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionEquals) GetEqualsOk() (*SignOnPolicyActionCommonConditionEqualsEquals, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Equals, true
}

// SetEquals sets field value
func (o *SignOnPolicyActionCommonConditionEquals) SetEquals(v SignOnPolicyActionCommonConditionEqualsEquals) {
	o.Equals = v
}

func (o SignOnPolicyActionCommonConditionEquals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["equals"] = o.Equals
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionCommonConditionEquals struct {
	value *SignOnPolicyActionCommonConditionEquals
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionEquals) Get() *SignOnPolicyActionCommonConditionEquals {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionEquals) Set(val *SignOnPolicyActionCommonConditionEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionEquals(val *SignOnPolicyActionCommonConditionEquals) *NullableSignOnPolicyActionCommonConditionEquals {
	return &NullableSignOnPolicyActionCommonConditionEquals{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


