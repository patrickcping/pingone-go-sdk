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

// checks if the AuthorizeEditorDataServicesConnectorServiceDefinitionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataServicesConnectorServiceDefinitionDTO{}

// AuthorizeEditorDataServicesConnectorServiceDefinitionDTO struct for AuthorizeEditorDataServicesConnectorServiceDefinitionDTO
type AuthorizeEditorDataServicesConnectorServiceDefinitionDTO struct {
	AuthorizeEditorDataDefinitionsServiceDefinitionDTO
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Processor *AuthorizeEditorDataProcessorDTO `json:"processor,omitempty"`
	ValueType AuthorizeEditorDataValueTypeDTO `json:"valueType"`
	ServiceSettings AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO `json:"serviceSettings"`
}

// NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO(valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, name string, serviceType string) *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO {
	this := AuthorizeEditorDataServicesConnectorServiceDefinitionDTO{}
	this.Name = name
	this.ServiceType = serviceType
	this.ValueType = valueType
	this.ServiceSettings = serviceSettings
	return &this
}

// NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO {
	this := AuthorizeEditorDataServicesConnectorServiceDefinitionDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO {
	if o == nil || IsNil(o.Processor) {
		var ret AuthorizeEditorDataProcessorDTO
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool) {
	if o == nil || IsNil(o.Processor) {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasProcessor() bool {
	if o != nil && !IsNil(o.Processor) {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given AuthorizeEditorDataProcessorDTO and assigns it to the Processor field.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO) {
	o.Processor = &v
}

// GetValueType returns the ValueType field value
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueType() AuthorizeEditorDataValueTypeDTO {
	if o == nil {
		var ret AuthorizeEditorDataValueTypeDTO
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO) {
	o.ValueType = v
}

// GetServiceSettings returns the ServiceSettings field value
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettings() AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO {
	if o == nil {
		var ret AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO
		return ret
	}

	return o.ServiceSettings
}

// GetServiceSettingsOk returns a tuple with the ServiceSettings field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceSettings, true
}

// SetServiceSettings sets field value
func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetServiceSettings(v AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) {
	o.ServiceSettings = v
}

func (o AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataDefinitionsServiceDefinitionDTO, errAuthorizeEditorDataDefinitionsServiceDefinitionDTO := json.Marshal(o.AuthorizeEditorDataDefinitionsServiceDefinitionDTO)
	if errAuthorizeEditorDataDefinitionsServiceDefinitionDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataDefinitionsServiceDefinitionDTO
	}
	errAuthorizeEditorDataDefinitionsServiceDefinitionDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataDefinitionsServiceDefinitionDTO), &toSerialize)
	if errAuthorizeEditorDataDefinitionsServiceDefinitionDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataDefinitionsServiceDefinitionDTO
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Processor) {
		toSerialize["processor"] = o.Processor
	}
	toSerialize["valueType"] = o.ValueType
	toSerialize["serviceSettings"] = o.ServiceSettings
	return toSerialize, nil
}

type NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO struct {
	value *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) Get() *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) Set(val *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO(val *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) *NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO {
	return &NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataServicesConnectorServiceDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


