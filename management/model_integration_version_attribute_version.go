/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the IntegrationVersionAttributeVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionAttributeVersion{}

// IntegrationVersionAttributeVersion The relationship to the parent integration version.
type IntegrationVersionAttributeVersion struct {
	// The platform-generated ID of the parent integration version.
	Id string `json:"id"`
}

// NewIntegrationVersionAttributeVersion instantiates a new IntegrationVersionAttributeVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionAttributeVersion(id string) *IntegrationVersionAttributeVersion {
	this := IntegrationVersionAttributeVersion{}
	this.Id = id
	return &this
}

// NewIntegrationVersionAttributeVersionWithDefaults instantiates a new IntegrationVersionAttributeVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionAttributeVersionWithDefaults() *IntegrationVersionAttributeVersion {
	this := IntegrationVersionAttributeVersion{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationVersionAttributeVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionAttributeVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationVersionAttributeVersion) SetId(v string) {
	o.Id = v
}

func (o IntegrationVersionAttributeVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionAttributeVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableIntegrationVersionAttributeVersion struct {
	value *IntegrationVersionAttributeVersion
	isSet bool
}

func (v NullableIntegrationVersionAttributeVersion) Get() *IntegrationVersionAttributeVersion {
	return v.value
}

func (v *NullableIntegrationVersionAttributeVersion) Set(val *IntegrationVersionAttributeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionAttributeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionAttributeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionAttributeVersion(val *IntegrationVersionAttributeVersion) *NullableIntegrationVersionAttributeVersion {
	return &NullableIntegrationVersionAttributeVersion{value: val, isSet: true}
}

func (v NullableIntegrationVersionAttributeVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionAttributeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


