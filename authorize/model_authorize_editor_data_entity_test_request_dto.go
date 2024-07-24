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

// checks if the AuthorizeEditorDataEntityTestRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataEntityTestRequestDTO{}

// AuthorizeEditorDataEntityTestRequestDTO struct for AuthorizeEditorDataEntityTestRequestDTO
type AuthorizeEditorDataEntityTestRequestDTO struct {
	Parameters []AuthorizeEditorDataEntityTestingParameterDTO `json:"parameters"`
	UserContext *AuthorizeEditorDataEntityTestingUserContextDTO `json:"userContext,omitempty"`
	Overrides []AuthorizeEditorDataEntityTestingOverrideDTO `json:"overrides,omitempty"`
}

// NewAuthorizeEditorDataEntityTestRequestDTO instantiates a new AuthorizeEditorDataEntityTestRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataEntityTestRequestDTO(parameters []AuthorizeEditorDataEntityTestingParameterDTO) *AuthorizeEditorDataEntityTestRequestDTO {
	this := AuthorizeEditorDataEntityTestRequestDTO{}
	this.Parameters = parameters
	return &this
}

// NewAuthorizeEditorDataEntityTestRequestDTOWithDefaults instantiates a new AuthorizeEditorDataEntityTestRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataEntityTestRequestDTOWithDefaults() *AuthorizeEditorDataEntityTestRequestDTO {
	this := AuthorizeEditorDataEntityTestRequestDTO{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetParameters() []AuthorizeEditorDataEntityTestingParameterDTO {
	if o == nil {
		var ret []AuthorizeEditorDataEntityTestingParameterDTO
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetParametersOk() ([]AuthorizeEditorDataEntityTestingParameterDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *AuthorizeEditorDataEntityTestRequestDTO) SetParameters(v []AuthorizeEditorDataEntityTestingParameterDTO) {
	o.Parameters = v
}

// GetUserContext returns the UserContext field value if set, zero value otherwise.
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetUserContext() AuthorizeEditorDataEntityTestingUserContextDTO {
	if o == nil || IsNil(o.UserContext) {
		var ret AuthorizeEditorDataEntityTestingUserContextDTO
		return ret
	}
	return *o.UserContext
}

// GetUserContextOk returns a tuple with the UserContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetUserContextOk() (*AuthorizeEditorDataEntityTestingUserContextDTO, bool) {
	if o == nil || IsNil(o.UserContext) {
		return nil, false
	}
	return o.UserContext, true
}

// HasUserContext returns a boolean if a field has been set.
func (o *AuthorizeEditorDataEntityTestRequestDTO) HasUserContext() bool {
	if o != nil && !IsNil(o.UserContext) {
		return true
	}

	return false
}

// SetUserContext gets a reference to the given AuthorizeEditorDataEntityTestingUserContextDTO and assigns it to the UserContext field.
func (o *AuthorizeEditorDataEntityTestRequestDTO) SetUserContext(v AuthorizeEditorDataEntityTestingUserContextDTO) {
	o.UserContext = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetOverrides() []AuthorizeEditorDataEntityTestingOverrideDTO {
	if o == nil || IsNil(o.Overrides) {
		var ret []AuthorizeEditorDataEntityTestingOverrideDTO
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataEntityTestRequestDTO) GetOverridesOk() ([]AuthorizeEditorDataEntityTestingOverrideDTO, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *AuthorizeEditorDataEntityTestRequestDTO) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []AuthorizeEditorDataEntityTestingOverrideDTO and assigns it to the Overrides field.
func (o *AuthorizeEditorDataEntityTestRequestDTO) SetOverrides(v []AuthorizeEditorDataEntityTestingOverrideDTO) {
	o.Overrides = v
}

func (o AuthorizeEditorDataEntityTestRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataEntityTestRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	if !IsNil(o.UserContext) {
		toSerialize["userContext"] = o.UserContext
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataEntityTestRequestDTO struct {
	value *AuthorizeEditorDataEntityTestRequestDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataEntityTestRequestDTO) Get() *AuthorizeEditorDataEntityTestRequestDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataEntityTestRequestDTO) Set(val *AuthorizeEditorDataEntityTestRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataEntityTestRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataEntityTestRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataEntityTestRequestDTO(val *AuthorizeEditorDataEntityTestRequestDTO) *NullableAuthorizeEditorDataEntityTestRequestDTO {
	return &NullableAuthorizeEditorDataEntityTestRequestDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataEntityTestRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataEntityTestRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


