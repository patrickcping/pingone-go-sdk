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

// checks if the ApplicationRolePermissionPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationRolePermissionPermission{}

// ApplicationRolePermissionPermission struct for ApplicationRolePermissionPermission
type ApplicationRolePermissionPermission struct {
	// The ID of the permission resource associated with a specified role.
	Id *string `json:"id,omitempty"`
	// The action associated with this permission.
	Action *string `json:"action,omitempty"`
	Resource *ApplicationRolePermissionPermissionResource `json:"resource,omitempty"`
}

// NewApplicationRolePermissionPermission instantiates a new ApplicationRolePermissionPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRolePermissionPermission() *ApplicationRolePermissionPermission {
	this := ApplicationRolePermissionPermission{}
	return &this
}

// NewApplicationRolePermissionPermissionWithDefaults instantiates a new ApplicationRolePermissionPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRolePermissionPermissionWithDefaults() *ApplicationRolePermissionPermission {
	this := ApplicationRolePermissionPermission{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationRolePermissionPermission) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermissionPermission) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationRolePermissionPermission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationRolePermissionPermission) SetId(v string) {
	o.Id = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApplicationRolePermissionPermission) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermissionPermission) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApplicationRolePermissionPermission) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ApplicationRolePermissionPermission) SetAction(v string) {
	o.Action = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ApplicationRolePermissionPermission) GetResource() ApplicationRolePermissionPermissionResource {
	if o == nil || IsNil(o.Resource) {
		var ret ApplicationRolePermissionPermissionResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermissionPermission) GetResourceOk() (*ApplicationRolePermissionPermissionResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ApplicationRolePermissionPermission) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ApplicationRolePermissionPermissionResource and assigns it to the Resource field.
func (o *ApplicationRolePermissionPermission) SetResource(v ApplicationRolePermissionPermissionResource) {
	o.Resource = &v
}

func (o ApplicationRolePermissionPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationRolePermissionPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableApplicationRolePermissionPermission struct {
	value *ApplicationRolePermissionPermission
	isSet bool
}

func (v NullableApplicationRolePermissionPermission) Get() *ApplicationRolePermissionPermission {
	return v.value
}

func (v *NullableApplicationRolePermissionPermission) Set(val *ApplicationRolePermissionPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRolePermissionPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRolePermissionPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRolePermissionPermission(val *ApplicationRolePermissionPermission) *NullableApplicationRolePermissionPermission {
	return &NullableApplicationRolePermissionPermission{value: val, isSet: true}
}

func (v NullableApplicationRolePermissionPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRolePermissionPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


