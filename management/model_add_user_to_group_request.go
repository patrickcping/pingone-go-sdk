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

// AddUserToGroupRequest struct for AddUserToGroupRequest
type AddUserToGroupRequest struct {
	// The group ID to add the user to
	Id *string `json:"id,omitempty"`
}

// NewAddUserToGroupRequest instantiates a new AddUserToGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserToGroupRequest() *AddUserToGroupRequest {
	this := AddUserToGroupRequest{}
	return &this
}

// NewAddUserToGroupRequestWithDefaults instantiates a new AddUserToGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserToGroupRequestWithDefaults() *AddUserToGroupRequest {
	this := AddUserToGroupRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddUserToGroupRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserToGroupRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddUserToGroupRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddUserToGroupRequest) SetId(v string) {
	o.Id = &v
}

func (o AddUserToGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAddUserToGroupRequest struct {
	value *AddUserToGroupRequest
	isSet bool
}

func (v NullableAddUserToGroupRequest) Get() *AddUserToGroupRequest {
	return v.value
}

func (v *NullableAddUserToGroupRequest) Set(val *AddUserToGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserToGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserToGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserToGroupRequest(val *AddUserToGroupRequest) *NullableAddUserToGroupRequest {
	return &NullableAddUserToGroupRequest{value: val, isSet: true}
}

func (v NullableAddUserToGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserToGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


