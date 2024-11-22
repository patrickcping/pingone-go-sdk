/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the ReadApplicationPermissions200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadApplicationPermissions200ResponseEmbedded{}

// ReadApplicationPermissions200ResponseEmbedded struct for ReadApplicationPermissions200ResponseEmbedded
type ReadApplicationPermissions200ResponseEmbedded struct {
	Permissions []ApplicationResourcePermission `json:"permissions,omitempty"`
}

// NewReadApplicationPermissions200ResponseEmbedded instantiates a new ReadApplicationPermissions200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApplicationPermissions200ResponseEmbedded() *ReadApplicationPermissions200ResponseEmbedded {
	this := ReadApplicationPermissions200ResponseEmbedded{}
	return &this
}

// NewReadApplicationPermissions200ResponseEmbeddedWithDefaults instantiates a new ReadApplicationPermissions200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApplicationPermissions200ResponseEmbeddedWithDefaults() *ReadApplicationPermissions200ResponseEmbedded {
	this := ReadApplicationPermissions200ResponseEmbedded{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ReadApplicationPermissions200ResponseEmbedded) GetPermissions() []ApplicationResourcePermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []ApplicationResourcePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationPermissions200ResponseEmbedded) GetPermissionsOk() ([]ApplicationResourcePermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ReadApplicationPermissions200ResponseEmbedded) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []ApplicationResourcePermission and assigns it to the Permissions field.
func (o *ReadApplicationPermissions200ResponseEmbedded) SetPermissions(v []ApplicationResourcePermission) {
	o.Permissions = v
}

func (o ReadApplicationPermissions200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadApplicationPermissions200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableReadApplicationPermissions200ResponseEmbedded struct {
	value *ReadApplicationPermissions200ResponseEmbedded
	isSet bool
}

func (v NullableReadApplicationPermissions200ResponseEmbedded) Get() *ReadApplicationPermissions200ResponseEmbedded {
	return v.value
}

func (v *NullableReadApplicationPermissions200ResponseEmbedded) Set(val *ReadApplicationPermissions200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApplicationPermissions200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApplicationPermissions200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApplicationPermissions200ResponseEmbedded(val *ReadApplicationPermissions200ResponseEmbedded) *NullableReadApplicationPermissions200ResponseEmbedded {
	return &NullableReadApplicationPermissions200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadApplicationPermissions200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApplicationPermissions200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
