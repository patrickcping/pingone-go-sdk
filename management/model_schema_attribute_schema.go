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

// SchemaAttributeSchema struct for SchemaAttributeSchema
type SchemaAttributeSchema struct {
	// A string that specifies the identifier of the resource referenced by this relationship.
	Id *string `json:"id,omitempty"`
}

// NewSchemaAttributeSchema instantiates a new SchemaAttributeSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAttributeSchema() *SchemaAttributeSchema {
	this := SchemaAttributeSchema{}
	return &this
}

// NewSchemaAttributeSchemaWithDefaults instantiates a new SchemaAttributeSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAttributeSchemaWithDefaults() *SchemaAttributeSchema {
	this := SchemaAttributeSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SchemaAttributeSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttributeSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SchemaAttributeSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SchemaAttributeSchema) SetId(v string) {
	o.Id = &v
}

func (o SchemaAttributeSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaAttributeSchema struct {
	value *SchemaAttributeSchema
	isSet bool
}

func (v NullableSchemaAttributeSchema) Get() *SchemaAttributeSchema {
	return v.value
}

func (v *NullableSchemaAttributeSchema) Set(val *SchemaAttributeSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAttributeSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAttributeSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAttributeSchema(val *SchemaAttributeSchema) *NullableSchemaAttributeSchema {
	return &NullableSchemaAttributeSchema{value: val, isSet: true}
}

func (v NullableSchemaAttributeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAttributeSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


