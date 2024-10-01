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

// checks if the AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO{}

// AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO struct for AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO
type AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO struct {
	AuthorizeEditorDataRulesEffectSettingsDTO
	Condition AuthorizeEditorDataConditionDTO `json:"condition"`
}

// NewAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO instantiates a new AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO(condition AuthorizeEditorDataConditionDTO, type_ EnumAuthorizeEditorDataRulesEffectSettingsDTOType) *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO {
	this := AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO{}
	this.Type = type_
	this.Condition = condition
	return &this
}

// NewAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTOWithDefaults instantiates a new AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTOWithDefaults() *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO {
	this := AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO{}
	return &this
}

// GetCondition returns the Condition field value
func (o *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) GetCondition() AuthorizeEditorDataConditionDTO {
	if o == nil {
		var ret AuthorizeEditorDataConditionDTO
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) SetCondition(v AuthorizeEditorDataConditionDTO) {
	o.Condition = v
}

func (o AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataRulesEffectSettingsDTO, errAuthorizeEditorDataRulesEffectSettingsDTO := json.Marshal(o.AuthorizeEditorDataRulesEffectSettingsDTO)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	errAuthorizeEditorDataRulesEffectSettingsDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataRulesEffectSettingsDTO), &toSerialize)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	toSerialize["condition"] = o.Condition
	return toSerialize, nil
}

type NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO struct {
	value *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) Get() *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) Set(val *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO(val *AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) *NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO {
	return &NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


