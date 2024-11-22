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

// checks if the ReadAllCredentialTypes200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllCredentialTypes200ResponseEmbedded{}

// ReadAllCredentialTypes200ResponseEmbedded struct for ReadAllCredentialTypes200ResponseEmbedded
type ReadAllCredentialTypes200ResponseEmbedded struct {
	Items []CredentialType `json:"items,omitempty"`
}

// NewReadAllCredentialTypes200ResponseEmbedded instantiates a new ReadAllCredentialTypes200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllCredentialTypes200ResponseEmbedded() *ReadAllCredentialTypes200ResponseEmbedded {
	this := ReadAllCredentialTypes200ResponseEmbedded{}
	return &this
}

// NewReadAllCredentialTypes200ResponseEmbeddedWithDefaults instantiates a new ReadAllCredentialTypes200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllCredentialTypes200ResponseEmbeddedWithDefaults() *ReadAllCredentialTypes200ResponseEmbedded {
	this := ReadAllCredentialTypes200ResponseEmbedded{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ReadAllCredentialTypes200ResponseEmbedded) GetItems() []CredentialType {
	if o == nil || IsNil(o.Items) {
		var ret []CredentialType
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllCredentialTypes200ResponseEmbedded) GetItemsOk() ([]CredentialType, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ReadAllCredentialTypes200ResponseEmbedded) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CredentialType and assigns it to the Items field.
func (o *ReadAllCredentialTypes200ResponseEmbedded) SetItems(v []CredentialType) {
	o.Items = v
}

func (o ReadAllCredentialTypes200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllCredentialTypes200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableReadAllCredentialTypes200ResponseEmbedded struct {
	value *ReadAllCredentialTypes200ResponseEmbedded
	isSet bool
}

func (v NullableReadAllCredentialTypes200ResponseEmbedded) Get() *ReadAllCredentialTypes200ResponseEmbedded {
	return v.value
}

func (v *NullableReadAllCredentialTypes200ResponseEmbedded) Set(val *ReadAllCredentialTypes200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllCredentialTypes200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllCredentialTypes200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllCredentialTypes200ResponseEmbedded(val *ReadAllCredentialTypes200ResponseEmbedded) *NullableReadAllCredentialTypes200ResponseEmbedded {
	return &NullableReadAllCredentialTypes200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadAllCredentialTypes200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllCredentialTypes200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
