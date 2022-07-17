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

// ApplicationResourceGrantApplication struct for ApplicationResourceGrantApplication
type ApplicationResourceGrantApplication struct {
	// A string that specifies the application resource ID.
	Id *string `json:"id,omitempty"`
}

// NewApplicationResourceGrantApplication instantiates a new ApplicationResourceGrantApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceGrantApplication() *ApplicationResourceGrantApplication {
	this := ApplicationResourceGrantApplication{}
	return &this
}

// NewApplicationResourceGrantApplicationWithDefaults instantiates a new ApplicationResourceGrantApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceGrantApplicationWithDefaults() *ApplicationResourceGrantApplication {
	this := ApplicationResourceGrantApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationResourceGrantApplication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrantApplication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationResourceGrantApplication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationResourceGrantApplication) SetId(v string) {
	o.Id = &v
}

func (o ApplicationResourceGrantApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResourceGrantApplication struct {
	value *ApplicationResourceGrantApplication
	isSet bool
}

func (v NullableApplicationResourceGrantApplication) Get() *ApplicationResourceGrantApplication {
	return v.value
}

func (v *NullableApplicationResourceGrantApplication) Set(val *ApplicationResourceGrantApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceGrantApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceGrantApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceGrantApplication(val *ApplicationResourceGrantApplication) *NullableApplicationResourceGrantApplication {
	return &NullableApplicationResourceGrantApplication{value: val, isSet: true}
}

func (v NullableApplicationResourceGrantApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceGrantApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


