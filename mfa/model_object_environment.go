/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the ObjectEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectEnvironment{}

// ObjectEnvironment struct for ObjectEnvironment
type ObjectEnvironment struct {
	// A string that specifies the environment associated with the object.
	Id *string `json:"id,omitempty"`
}

// NewObjectEnvironment instantiates a new ObjectEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectEnvironment() *ObjectEnvironment {
	this := ObjectEnvironment{}
	return &this
}

// NewObjectEnvironmentWithDefaults instantiates a new ObjectEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectEnvironmentWithDefaults() *ObjectEnvironment {
	this := ObjectEnvironment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectEnvironment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectEnvironment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectEnvironment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectEnvironment) SetId(v string) {
	o.Id = &v
}

func (o ObjectEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableObjectEnvironment struct {
	value *ObjectEnvironment
	isSet bool
}

func (v NullableObjectEnvironment) Get() *ObjectEnvironment {
	return v.value
}

func (v *NullableObjectEnvironment) Set(val *ObjectEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectEnvironment(val *ObjectEnvironment) *NullableObjectEnvironment {
	return &NullableObjectEnvironment{value: val, isSet: true}
}

func (v NullableObjectEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


