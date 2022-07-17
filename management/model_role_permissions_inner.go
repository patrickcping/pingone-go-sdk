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

// RolePermissionsInner struct for RolePermissionsInner
type RolePermissionsInner struct {
	// A string that specifies the resource for which the permission is applicable.
	Classifier *string `json:"classifier,omitempty"`
	// A string that specifies the description of what the permission enables for the role.
	Description *string `json:"description,omitempty"`
}

// NewRolePermissionsInner instantiates a new RolePermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionsInner() *RolePermissionsInner {
	this := RolePermissionsInner{}
	return &this
}

// NewRolePermissionsInnerWithDefaults instantiates a new RolePermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionsInnerWithDefaults() *RolePermissionsInner {
	this := RolePermissionsInner{}
	return &this
}

// GetClassifier returns the Classifier field value if set, zero value otherwise.
func (o *RolePermissionsInner) GetClassifier() string {
	if o == nil || o.Classifier == nil {
		var ret string
		return ret
	}
	return *o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsInner) GetClassifierOk() (*string, bool) {
	if o == nil || o.Classifier == nil {
		return nil, false
	}
	return o.Classifier, true
}

// HasClassifier returns a boolean if a field has been set.
func (o *RolePermissionsInner) HasClassifier() bool {
	if o != nil && o.Classifier != nil {
		return true
	}

	return false
}

// SetClassifier gets a reference to the given string and assigns it to the Classifier field.
func (o *RolePermissionsInner) SetClassifier(v string) {
	o.Classifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RolePermissionsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolePermissionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RolePermissionsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RolePermissionsInner) SetDescription(v string) {
	o.Description = &v
}

func (o RolePermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Classifier != nil {
		toSerialize["classifier"] = o.Classifier
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableRolePermissionsInner struct {
	value *RolePermissionsInner
	isSet bool
}

func (v NullableRolePermissionsInner) Get() *RolePermissionsInner {
	return v.value
}

func (v *NullableRolePermissionsInner) Set(val *RolePermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionsInner(val *RolePermissionsInner) *NullableRolePermissionsInner {
	return &NullableRolePermissionsInner{value: val, isSet: true}
}

func (v NullableRolePermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


