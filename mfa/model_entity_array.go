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

// checks if the EntityArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityArray{}

// EntityArray struct for EntityArray
type EntityArray struct {
	Embedded *EntityArrayEmbedded `json:"_embedded,omitempty"`
	Count *float32 `json:"count,omitempty"`
	Size *float32 `json:"size,omitempty"`
}

// NewEntityArray instantiates a new EntityArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityArray() *EntityArray {
	this := EntityArray{}
	return &this
}

// NewEntityArrayWithDefaults instantiates a new EntityArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityArrayWithDefaults() *EntityArray {
	this := EntityArray{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *EntityArray) GetEmbedded() EntityArrayEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret EntityArrayEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArray) GetEmbeddedOk() (*EntityArrayEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *EntityArray) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EntityArrayEmbedded and assigns it to the Embedded field.
func (o *EntityArray) SetEmbedded(v EntityArrayEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EntityArray) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArray) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EntityArray) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *EntityArray) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EntityArray) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArray) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EntityArray) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *EntityArray) SetSize(v float32) {
	o.Size = &v
}

func (o EntityArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableEntityArray struct {
	value *EntityArray
	isSet bool
}

func (v NullableEntityArray) Get() *EntityArray {
	return v.value
}

func (v *NullableEntityArray) Set(val *EntityArray) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArray) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArray(val *EntityArray) *NullableEntityArray {
	return &NullableEntityArray{value: val, isSet: true}
}

func (v NullableEntityArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


