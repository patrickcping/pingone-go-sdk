/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// GroupPopulation The population to assign the group to
type GroupPopulation struct {
	// The population ID
	Id string `json:"id"`
}

// NewGroupPopulation instantiates a new GroupPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPopulation(id string) *GroupPopulation {
	this := GroupPopulation{}
	this.Id = id
	return &this
}

// NewGroupPopulationWithDefaults instantiates a new GroupPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPopulationWithDefaults() *GroupPopulation {
	this := GroupPopulation{}
	return &this
}

// GetId returns the Id field value
func (o *GroupPopulation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupPopulation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupPopulation) SetId(v string) {
	o.Id = v
}

func (o GroupPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGroupPopulation struct {
	value *GroupPopulation
	isSet bool
}

func (v NullableGroupPopulation) Get() *GroupPopulation {
	return v.value
}

func (v *NullableGroupPopulation) Set(val *GroupPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPopulation(val *GroupPopulation) *NullableGroupPopulation {
	return &NullableGroupPopulation{value: val, isSet: true}
}

func (v NullableGroupPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


