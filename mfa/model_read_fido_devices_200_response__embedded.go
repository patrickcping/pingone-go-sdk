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

// checks if the ReadFidoDevices200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadFidoDevices200ResponseEmbedded{}

// ReadFidoDevices200ResponseEmbedded struct for ReadFidoDevices200ResponseEmbedded
type ReadFidoDevices200ResponseEmbedded struct {
	FidoDevicesMetadata []map[string]interface{} `json:"fidoDevicesMetadata,omitempty"`
}

// NewReadFidoDevices200ResponseEmbedded instantiates a new ReadFidoDevices200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFidoDevices200ResponseEmbedded() *ReadFidoDevices200ResponseEmbedded {
	this := ReadFidoDevices200ResponseEmbedded{}
	return &this
}

// NewReadFidoDevices200ResponseEmbeddedWithDefaults instantiates a new ReadFidoDevices200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFidoDevices200ResponseEmbeddedWithDefaults() *ReadFidoDevices200ResponseEmbedded {
	this := ReadFidoDevices200ResponseEmbedded{}
	return &this
}

// GetFidoDevicesMetadata returns the FidoDevicesMetadata field value if set, zero value otherwise.
func (o *ReadFidoDevices200ResponseEmbedded) GetFidoDevicesMetadata() []map[string]interface{} {
	if o == nil || IsNil(o.FidoDevicesMetadata) {
		var ret []map[string]interface{}
		return ret
	}
	return o.FidoDevicesMetadata
}

// GetFidoDevicesMetadataOk returns a tuple with the FidoDevicesMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFidoDevices200ResponseEmbedded) GetFidoDevicesMetadataOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.FidoDevicesMetadata) {
		return nil, false
	}
	return o.FidoDevicesMetadata, true
}

// HasFidoDevicesMetadata returns a boolean if a field has been set.
func (o *ReadFidoDevices200ResponseEmbedded) HasFidoDevicesMetadata() bool {
	if o != nil && !IsNil(o.FidoDevicesMetadata) {
		return true
	}

	return false
}

// SetFidoDevicesMetadata gets a reference to the given []map[string]interface{} and assigns it to the FidoDevicesMetadata field.
func (o *ReadFidoDevices200ResponseEmbedded) SetFidoDevicesMetadata(v []map[string]interface{}) {
	o.FidoDevicesMetadata = v
}

func (o ReadFidoDevices200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadFidoDevices200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FidoDevicesMetadata) {
		toSerialize["fidoDevicesMetadata"] = o.FidoDevicesMetadata
	}
	return toSerialize, nil
}

type NullableReadFidoDevices200ResponseEmbedded struct {
	value *ReadFidoDevices200ResponseEmbedded
	isSet bool
}

func (v NullableReadFidoDevices200ResponseEmbedded) Get() *ReadFidoDevices200ResponseEmbedded {
	return v.value
}

func (v *NullableReadFidoDevices200ResponseEmbedded) Set(val *ReadFidoDevices200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFidoDevices200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFidoDevices200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFidoDevices200ResponseEmbedded(val *ReadFidoDevices200ResponseEmbedded) *NullableReadFidoDevices200ResponseEmbedded {
	return &NullableReadFidoDevices200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadFidoDevices200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFidoDevices200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
