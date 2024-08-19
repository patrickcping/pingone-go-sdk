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

// checks if the AuthorizeEditorDataEnvironmentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataEnvironmentDTO{}

// AuthorizeEditorDataEnvironmentDTO struct for AuthorizeEditorDataEnvironmentDTO
type AuthorizeEditorDataEnvironmentDTO struct {
	Id *string `json:"id,omitempty"`
}

// NewAuthorizeEditorDataEnvironmentDTO instantiates a new AuthorizeEditorDataEnvironmentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataEnvironmentDTO() *AuthorizeEditorDataEnvironmentDTO {
	this := AuthorizeEditorDataEnvironmentDTO{}
	return &this
}

// NewAuthorizeEditorDataEnvironmentDTOWithDefaults instantiates a new AuthorizeEditorDataEnvironmentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataEnvironmentDTOWithDefaults() *AuthorizeEditorDataEnvironmentDTO {
	this := AuthorizeEditorDataEnvironmentDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizeEditorDataEnvironmentDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataEnvironmentDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizeEditorDataEnvironmentDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizeEditorDataEnvironmentDTO) SetId(v string) {
	o.Id = &v
}

func (o AuthorizeEditorDataEnvironmentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataEnvironmentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataEnvironmentDTO struct {
	value *AuthorizeEditorDataEnvironmentDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataEnvironmentDTO) Get() *AuthorizeEditorDataEnvironmentDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataEnvironmentDTO) Set(val *AuthorizeEditorDataEnvironmentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataEnvironmentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataEnvironmentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataEnvironmentDTO(val *AuthorizeEditorDataEnvironmentDTO) *NullableAuthorizeEditorDataEnvironmentDTO {
	return &NullableAuthorizeEditorDataEnvironmentDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataEnvironmentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataEnvironmentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


