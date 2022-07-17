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

// ApplicationAccessControlGroupGroupsInner struct for ApplicationAccessControlGroupGroupsInner
type ApplicationAccessControlGroupGroupsInner struct {
	Id string `json:"id"`
}

// NewApplicationAccessControlGroupGroupsInner instantiates a new ApplicationAccessControlGroupGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAccessControlGroupGroupsInner(id string) *ApplicationAccessControlGroupGroupsInner {
	this := ApplicationAccessControlGroupGroupsInner{}
	this.Id = id
	return &this
}

// NewApplicationAccessControlGroupGroupsInnerWithDefaults instantiates a new ApplicationAccessControlGroupGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAccessControlGroupGroupsInnerWithDefaults() *ApplicationAccessControlGroupGroupsInner {
	this := ApplicationAccessControlGroupGroupsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationAccessControlGroupGroupsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationAccessControlGroupGroupsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationAccessControlGroupGroupsInner) SetId(v string) {
	o.Id = v
}

func (o ApplicationAccessControlGroupGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAccessControlGroupGroupsInner struct {
	value *ApplicationAccessControlGroupGroupsInner
	isSet bool
}

func (v NullableApplicationAccessControlGroupGroupsInner) Get() *ApplicationAccessControlGroupGroupsInner {
	return v.value
}

func (v *NullableApplicationAccessControlGroupGroupsInner) Set(val *ApplicationAccessControlGroupGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAccessControlGroupGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAccessControlGroupGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAccessControlGroupGroupsInner(val *ApplicationAccessControlGroupGroupsInner) *NullableApplicationAccessControlGroupGroupsInner {
	return &NullableApplicationAccessControlGroupGroupsInner{value: val, isSet: true}
}

func (v NullableApplicationAccessControlGroupGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAccessControlGroupGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


