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

// UserIdentityProvider struct for UserIdentityProvider
type UserIdentityProvider struct {
	// A mutable string that identifies the external identity provider used to authenticate the user. If not provided, PingOne is the identity provider. This attribute is required if the identity provider is authoritative for just-in-time user provisioning.
	Id *string `json:"id,omitempty"`
	Type *EnumIdentityProvider `json:"type,omitempty"`
}

// NewUserIdentityProvider instantiates a new UserIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentityProvider() *UserIdentityProvider {
	this := UserIdentityProvider{}
	return &this
}

// NewUserIdentityProviderWithDefaults instantiates a new UserIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityProviderWithDefaults() *UserIdentityProvider {
	this := UserIdentityProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserIdentityProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserIdentityProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserIdentityProvider) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserIdentityProvider) GetType() EnumIdentityProvider {
	if o == nil || o.Type == nil {
		var ret EnumIdentityProvider
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityProvider) GetTypeOk() (*EnumIdentityProvider, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserIdentityProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumIdentityProvider and assigns it to the Type field.
func (o *UserIdentityProvider) SetType(v EnumIdentityProvider) {
	o.Type = &v
}

func (o UserIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUserIdentityProvider struct {
	value *UserIdentityProvider
	isSet bool
}

func (v NullableUserIdentityProvider) Get() *UserIdentityProvider {
	return v.value
}

func (v *NullableUserIdentityProvider) Set(val *UserIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentityProvider(val *UserIdentityProvider) *NullableUserIdentityProvider {
	return &NullableUserIdentityProvider{value: val, isSet: true}
}

func (v NullableUserIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


