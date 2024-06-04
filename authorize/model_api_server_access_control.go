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

// checks if the APIServerAccessControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerAccessControl{}

// APIServerAccessControl struct for APIServerAccessControl
type APIServerAccessControl struct {
	Custom *APIServerAccessControlCustom `json:"custom,omitempty"`
}

// NewAPIServerAccessControl instantiates a new APIServerAccessControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerAccessControl() *APIServerAccessControl {
	this := APIServerAccessControl{}
	return &this
}

// NewAPIServerAccessControlWithDefaults instantiates a new APIServerAccessControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerAccessControlWithDefaults() *APIServerAccessControl {
	this := APIServerAccessControl{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *APIServerAccessControl) GetCustom() APIServerAccessControlCustom {
	if o == nil || IsNil(o.Custom) {
		var ret APIServerAccessControlCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerAccessControl) GetCustomOk() (*APIServerAccessControlCustom, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *APIServerAccessControl) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given APIServerAccessControlCustom and assigns it to the Custom field.
func (o *APIServerAccessControl) SetCustom(v APIServerAccessControlCustom) {
	o.Custom = &v
}

func (o APIServerAccessControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerAccessControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	return toSerialize, nil
}

type NullableAPIServerAccessControl struct {
	value *APIServerAccessControl
	isSet bool
}

func (v NullableAPIServerAccessControl) Get() *APIServerAccessControl {
	return v.value
}

func (v *NullableAPIServerAccessControl) Set(val *APIServerAccessControl) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerAccessControl) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerAccessControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerAccessControl(val *APIServerAccessControl) *NullableAPIServerAccessControl {
	return &NullableAPIServerAccessControl{value: val, isSet: true}
}

func (v NullableAPIServerAccessControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerAccessControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


