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

// ApplicationAttributeMapping struct for ApplicationAttributeMapping
type ApplicationAttributeMapping struct {
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	MappingType *EnumAttributeMappingType `json:"mappingType,omitempty"`
	// A string that specifies the name of attribute and must be unique within an application. For SAML applications, the samlAssertion.subject name is a reserved case-insensitive name which indicates the mapping to be used for the subject in an assertion. For OpenID Connect applications, the following names are reserved and cannot be used acr, amr, at_hash, aud, auth_time, azp, client_id, exp, iat, iss, jti, nbf, nonce, org, scope, sid, sub  This is a required property.
	Name string `json:"name"`
	// A boolean to specify whether a mapping value is required for this attribute. If true, a value must be set and a non-empty value must be available in the SAML assertion or ID token.
	Required bool `json:"required"`
	// The time the resource was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// A string that specifies the string constants or expression for mapping the attribute path against a specific source. The expression format is ${<source>.<attribute_path>}. The only supported source is user (for example, ${user.id}). This is a required property.
	Value string `json:"value"`
}

// NewApplicationAttributeMapping instantiates a new ApplicationAttributeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAttributeMapping(name string, required bool, value string) *ApplicationAttributeMapping {
	this := ApplicationAttributeMapping{}
	this.Name = name
	this.Required = required
	this.Value = value
	return &this
}

// NewApplicationAttributeMappingWithDefaults instantiates a new ApplicationAttributeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAttributeMappingWithDefaults() *ApplicationAttributeMapping {
	this := ApplicationAttributeMapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationAttributeMapping) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationAttributeMapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationAttributeMapping) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationAttributeMapping) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationAttributeMapping) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApplicationAttributeMapping) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetMappingType returns the MappingType field value if set, zero value otherwise.
func (o *ApplicationAttributeMapping) GetMappingType() EnumAttributeMappingType {
	if o == nil || o.MappingType == nil {
		var ret EnumAttributeMappingType
		return ret
	}
	return *o.MappingType
}

// GetMappingTypeOk returns a tuple with the MappingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetMappingTypeOk() (*EnumAttributeMappingType, bool) {
	if o == nil || o.MappingType == nil {
		return nil, false
	}
	return o.MappingType, true
}

// HasMappingType returns a boolean if a field has been set.
func (o *ApplicationAttributeMapping) HasMappingType() bool {
	if o != nil && o.MappingType != nil {
		return true
	}

	return false
}

// SetMappingType gets a reference to the given EnumAttributeMappingType and assigns it to the MappingType field.
func (o *ApplicationAttributeMapping) SetMappingType(v EnumAttributeMappingType) {
	o.MappingType = &v
}

// GetName returns the Name field value
func (o *ApplicationAttributeMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationAttributeMapping) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value
func (o *ApplicationAttributeMapping) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *ApplicationAttributeMapping) SetRequired(v bool) {
	o.Required = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationAttributeMapping) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationAttributeMapping) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApplicationAttributeMapping) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value
func (o *ApplicationAttributeMapping) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApplicationAttributeMapping) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ApplicationAttributeMapping) SetValue(v string) {
	o.Value = v
}

func (o ApplicationAttributeMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.MappingType != nil {
		toSerialize["mappingType"] = o.MappingType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["required"] = o.Required
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAttributeMapping struct {
	value *ApplicationAttributeMapping
	isSet bool
}

func (v NullableApplicationAttributeMapping) Get() *ApplicationAttributeMapping {
	return v.value
}

func (v *NullableApplicationAttributeMapping) Set(val *ApplicationAttributeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAttributeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAttributeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAttributeMapping(val *ApplicationAttributeMapping) *NullableApplicationAttributeMapping {
	return &NullableApplicationAttributeMapping{value: val, isSet: true}
}

func (v NullableApplicationAttributeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAttributeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


