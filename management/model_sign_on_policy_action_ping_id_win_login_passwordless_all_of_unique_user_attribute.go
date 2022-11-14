/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute struct for SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute
type SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute struct {
	// Schema attributes to match against the provided identifier when searching for a user in the directory. Only unique attributes may be included.
	Name string `json:"name"`
}

// NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute instantiates a new SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute(name string) *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute {
	this := SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute{}
	this.Name = name
	return &this
}

// NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttributeWithDefaults instantiates a new SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttributeWithDefaults() *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute {
	this := SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) SetName(v string) {
	o.Name = v
}

func (o SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute struct {
	value *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute
	isSet bool
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) Get() *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute {
	return v.value
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) Set(val *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute(val *SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute {
	return &NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


