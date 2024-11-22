/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the ReadFidoPolicies200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadFidoPolicies200ResponseEmbedded{}

// ReadFidoPolicies200ResponseEmbedded struct for ReadFidoPolicies200ResponseEmbedded
type ReadFidoPolicies200ResponseEmbedded struct {
	FidoPolicies []FIDOPolicy `json:"fidoPolicies,omitempty"`
}

// NewReadFidoPolicies200ResponseEmbedded instantiates a new ReadFidoPolicies200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFidoPolicies200ResponseEmbedded() *ReadFidoPolicies200ResponseEmbedded {
	this := ReadFidoPolicies200ResponseEmbedded{}
	return &this
}

// NewReadFidoPolicies200ResponseEmbeddedWithDefaults instantiates a new ReadFidoPolicies200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFidoPolicies200ResponseEmbeddedWithDefaults() *ReadFidoPolicies200ResponseEmbedded {
	this := ReadFidoPolicies200ResponseEmbedded{}
	return &this
}

// GetFidoPolicies returns the FidoPolicies field value if set, zero value otherwise.
func (o *ReadFidoPolicies200ResponseEmbedded) GetFidoPolicies() []FIDOPolicy {
	if o == nil || IsNil(o.FidoPolicies) {
		var ret []FIDOPolicy
		return ret
	}
	return o.FidoPolicies
}

// GetFidoPoliciesOk returns a tuple with the FidoPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFidoPolicies200ResponseEmbedded) GetFidoPoliciesOk() ([]FIDOPolicy, bool) {
	if o == nil || IsNil(o.FidoPolicies) {
		return nil, false
	}
	return o.FidoPolicies, true
}

// HasFidoPolicies returns a boolean if a field has been set.
func (o *ReadFidoPolicies200ResponseEmbedded) HasFidoPolicies() bool {
	if o != nil && !IsNil(o.FidoPolicies) {
		return true
	}

	return false
}

// SetFidoPolicies gets a reference to the given []FIDOPolicy and assigns it to the FidoPolicies field.
func (o *ReadFidoPolicies200ResponseEmbedded) SetFidoPolicies(v []FIDOPolicy) {
	o.FidoPolicies = v
}

func (o ReadFidoPolicies200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadFidoPolicies200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FidoPolicies) {
		toSerialize["fidoPolicies"] = o.FidoPolicies
	}
	return toSerialize, nil
}

type NullableReadFidoPolicies200ResponseEmbedded struct {
	value *ReadFidoPolicies200ResponseEmbedded
	isSet bool
}

func (v NullableReadFidoPolicies200ResponseEmbedded) Get() *ReadFidoPolicies200ResponseEmbedded {
	return v.value
}

func (v *NullableReadFidoPolicies200ResponseEmbedded) Set(val *ReadFidoPolicies200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFidoPolicies200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFidoPolicies200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFidoPolicies200ResponseEmbedded(val *ReadFidoPolicies200ResponseEmbedded) *NullableReadFidoPolicies200ResponseEmbedded {
	return &NullableReadFidoPolicies200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadFidoPolicies200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFidoPolicies200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
