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

// checks if the IdentityProviderSAMLAllOfSpSigningKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderSAMLAllOfSpSigningKey{}

// IdentityProviderSAMLAllOfSpSigningKey struct for IdentityProviderSAMLAllOfSpSigningKey
type IdentityProviderSAMLAllOfSpSigningKey struct {
	// A string that specifies the service provider's signing key ID.
	Id string `json:"id"`
}

// NewIdentityProviderSAMLAllOfSpSigningKey instantiates a new IdentityProviderSAMLAllOfSpSigningKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderSAMLAllOfSpSigningKey(id string) *IdentityProviderSAMLAllOfSpSigningKey {
	this := IdentityProviderSAMLAllOfSpSigningKey{}
	this.Id = id
	return &this
}

// NewIdentityProviderSAMLAllOfSpSigningKeyWithDefaults instantiates a new IdentityProviderSAMLAllOfSpSigningKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderSAMLAllOfSpSigningKeyWithDefaults() *IdentityProviderSAMLAllOfSpSigningKey {
	this := IdentityProviderSAMLAllOfSpSigningKey{}
	return &this
}

// GetId returns the Id field value
func (o *IdentityProviderSAMLAllOfSpSigningKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOfSpSigningKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentityProviderSAMLAllOfSpSigningKey) SetId(v string) {
	o.Id = v
}

func (o IdentityProviderSAMLAllOfSpSigningKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderSAMLAllOfSpSigningKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableIdentityProviderSAMLAllOfSpSigningKey struct {
	value *IdentityProviderSAMLAllOfSpSigningKey
	isSet bool
}

func (v NullableIdentityProviderSAMLAllOfSpSigningKey) Get() *IdentityProviderSAMLAllOfSpSigningKey {
	return v.value
}

func (v *NullableIdentityProviderSAMLAllOfSpSigningKey) Set(val *IdentityProviderSAMLAllOfSpSigningKey) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderSAMLAllOfSpSigningKey) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderSAMLAllOfSpSigningKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderSAMLAllOfSpSigningKey(val *IdentityProviderSAMLAllOfSpSigningKey) *NullableIdentityProviderSAMLAllOfSpSigningKey {
	return &NullableIdentityProviderSAMLAllOfSpSigningKey{value: val, isSet: true}
}

func (v NullableIdentityProviderSAMLAllOfSpSigningKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderSAMLAllOfSpSigningKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


