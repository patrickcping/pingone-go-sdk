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

// checks if the AuthorizeEditorDataInputMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataInputMappingDTO{}

// AuthorizeEditorDataInputMappingDTO struct for AuthorizeEditorDataInputMappingDTO
type AuthorizeEditorDataInputMappingDTO struct {
	Property string `json:"property"`
	Type EnumAuthorizeEditorDataInputMappingDTOType `json:"type"`
}

// NewAuthorizeEditorDataInputMappingDTO instantiates a new AuthorizeEditorDataInputMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataInputMappingDTO(property string, type_ EnumAuthorizeEditorDataInputMappingDTOType) *AuthorizeEditorDataInputMappingDTO {
	this := AuthorizeEditorDataInputMappingDTO{}
	this.Property = property
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataInputMappingDTOWithDefaults instantiates a new AuthorizeEditorDataInputMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataInputMappingDTOWithDefaults() *AuthorizeEditorDataInputMappingDTO {
	this := AuthorizeEditorDataInputMappingDTO{}
	return &this
}

// GetProperty returns the Property field value
func (o *AuthorizeEditorDataInputMappingDTO) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataInputMappingDTO) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *AuthorizeEditorDataInputMappingDTO) SetProperty(v string) {
	o.Property = v
}

// GetType returns the Type field value
func (o *AuthorizeEditorDataInputMappingDTO) GetType() EnumAuthorizeEditorDataInputMappingDTOType {
	if o == nil {
		var ret EnumAuthorizeEditorDataInputMappingDTOType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataInputMappingDTO) GetTypeOk() (*EnumAuthorizeEditorDataInputMappingDTOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizeEditorDataInputMappingDTO) SetType(v EnumAuthorizeEditorDataInputMappingDTOType) {
	o.Type = v
}

func (o AuthorizeEditorDataInputMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataInputMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property"] = o.Property
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthorizeEditorDataInputMappingDTO struct {
	value *AuthorizeEditorDataInputMappingDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataInputMappingDTO) Get() *AuthorizeEditorDataInputMappingDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataInputMappingDTO) Set(val *AuthorizeEditorDataInputMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataInputMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataInputMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataInputMappingDTO(val *AuthorizeEditorDataInputMappingDTO) *NullableAuthorizeEditorDataInputMappingDTO {
	return &NullableAuthorizeEditorDataInputMappingDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataInputMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataInputMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


