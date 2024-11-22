/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ReadAllUserCredentials200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllUserCredentials200ResponseEmbedded{}

// ReadAllUserCredentials200ResponseEmbedded struct for ReadAllUserCredentials200ResponseEmbedded
type ReadAllUserCredentials200ResponseEmbedded struct {
	Items []UserCredential `json:"items,omitempty"`
}

// NewReadAllUserCredentials200ResponseEmbedded instantiates a new ReadAllUserCredentials200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllUserCredentials200ResponseEmbedded() *ReadAllUserCredentials200ResponseEmbedded {
	this := ReadAllUserCredentials200ResponseEmbedded{}
	return &this
}

// NewReadAllUserCredentials200ResponseEmbeddedWithDefaults instantiates a new ReadAllUserCredentials200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllUserCredentials200ResponseEmbeddedWithDefaults() *ReadAllUserCredentials200ResponseEmbedded {
	this := ReadAllUserCredentials200ResponseEmbedded{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ReadAllUserCredentials200ResponseEmbedded) GetItems() []UserCredential {
	if o == nil || IsNil(o.Items) {
		var ret []UserCredential
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllUserCredentials200ResponseEmbedded) GetItemsOk() ([]UserCredential, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ReadAllUserCredentials200ResponseEmbedded) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserCredential and assigns it to the Items field.
func (o *ReadAllUserCredentials200ResponseEmbedded) SetItems(v []UserCredential) {
	o.Items = v
}

func (o ReadAllUserCredentials200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllUserCredentials200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableReadAllUserCredentials200ResponseEmbedded struct {
	value *ReadAllUserCredentials200ResponseEmbedded
	isSet bool
}

func (v NullableReadAllUserCredentials200ResponseEmbedded) Get() *ReadAllUserCredentials200ResponseEmbedded {
	return v.value
}

func (v *NullableReadAllUserCredentials200ResponseEmbedded) Set(val *ReadAllUserCredentials200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllUserCredentials200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllUserCredentials200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllUserCredentials200ResponseEmbedded(val *ReadAllUserCredentials200ResponseEmbedded) *NullableReadAllUserCredentials200ResponseEmbedded {
	return &NullableReadAllUserCredentials200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadAllUserCredentials200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllUserCredentials200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
