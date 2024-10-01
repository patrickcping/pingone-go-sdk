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

// checks if the AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO{}

// AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO struct for AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO
type AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO struct {
	Type EnumAuthorizeEditorDataRulesEffectSettingsDTOType `json:"type"`
}

// NewAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO instantiates a new AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO(type_ EnumAuthorizeEditorDataRulesEffectSettingsDTOType) *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO {
	this := AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTOWithDefaults instantiates a new AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTOWithDefaults() *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO {
	this := AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO{}
	return &this
}

// GetType returns the Type field value
func (o *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) GetType() EnumAuthorizeEditorDataRulesEffectSettingsDTOType {
	if o == nil {
		var ret EnumAuthorizeEditorDataRulesEffectSettingsDTOType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) GetTypeOk() (*EnumAuthorizeEditorDataRulesEffectSettingsDTOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) SetType(v EnumAuthorizeEditorDataRulesEffectSettingsDTOType) {
	o.Type = v
}

func (o AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO struct {
	value *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) Get() *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) Set(val *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO(val *AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) *NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO {
	return &NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


