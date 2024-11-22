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

// checks if the ObjectApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectApplication{}

// ObjectApplication struct for ObjectApplication
type ObjectApplication struct {
	// A string that specifies the application resource ID associated with the object.
	Id *string `json:"id,omitempty"`
}

// NewObjectApplication instantiates a new ObjectApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectApplication() *ObjectApplication {
	this := ObjectApplication{}
	return &this
}

// NewObjectApplicationWithDefaults instantiates a new ObjectApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectApplicationWithDefaults() *ObjectApplication {
	this := ObjectApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectApplication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectApplication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectApplication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectApplication) SetId(v string) {
	o.Id = &v
}

func (o ObjectApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableObjectApplication struct {
	value *ObjectApplication
	isSet bool
}

func (v NullableObjectApplication) Get() *ObjectApplication {
	return v.value
}

func (v *NullableObjectApplication) Set(val *ObjectApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectApplication(val *ObjectApplication) *NullableObjectApplication {
	return &NullableObjectApplication{value: val, isSet: true}
}

func (v NullableObjectApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
