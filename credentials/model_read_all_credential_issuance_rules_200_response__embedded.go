/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ReadAllCredentialIssuanceRules200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllCredentialIssuanceRules200ResponseEmbedded{}

// ReadAllCredentialIssuanceRules200ResponseEmbedded struct for ReadAllCredentialIssuanceRules200ResponseEmbedded
type ReadAllCredentialIssuanceRules200ResponseEmbedded struct {
	IssuanceRules []CredentialIssuanceRule `json:"issuanceRules,omitempty"`
}

// NewReadAllCredentialIssuanceRules200ResponseEmbedded instantiates a new ReadAllCredentialIssuanceRules200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllCredentialIssuanceRules200ResponseEmbedded() *ReadAllCredentialIssuanceRules200ResponseEmbedded {
	this := ReadAllCredentialIssuanceRules200ResponseEmbedded{}
	return &this
}

// NewReadAllCredentialIssuanceRules200ResponseEmbeddedWithDefaults instantiates a new ReadAllCredentialIssuanceRules200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllCredentialIssuanceRules200ResponseEmbeddedWithDefaults() *ReadAllCredentialIssuanceRules200ResponseEmbedded {
	this := ReadAllCredentialIssuanceRules200ResponseEmbedded{}
	return &this
}

// GetIssuanceRules returns the IssuanceRules field value if set, zero value otherwise.
func (o *ReadAllCredentialIssuanceRules200ResponseEmbedded) GetIssuanceRules() []CredentialIssuanceRule {
	if o == nil || IsNil(o.IssuanceRules) {
		var ret []CredentialIssuanceRule
		return ret
	}
	return o.IssuanceRules
}

// GetIssuanceRulesOk returns a tuple with the IssuanceRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllCredentialIssuanceRules200ResponseEmbedded) GetIssuanceRulesOk() ([]CredentialIssuanceRule, bool) {
	if o == nil || IsNil(o.IssuanceRules) {
		return nil, false
	}
	return o.IssuanceRules, true
}

// HasIssuanceRules returns a boolean if a field has been set.
func (o *ReadAllCredentialIssuanceRules200ResponseEmbedded) HasIssuanceRules() bool {
	if o != nil && !IsNil(o.IssuanceRules) {
		return true
	}

	return false
}

// SetIssuanceRules gets a reference to the given []CredentialIssuanceRule and assigns it to the IssuanceRules field.
func (o *ReadAllCredentialIssuanceRules200ResponseEmbedded) SetIssuanceRules(v []CredentialIssuanceRule) {
	o.IssuanceRules = v
}

func (o ReadAllCredentialIssuanceRules200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllCredentialIssuanceRules200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IssuanceRules) {
		toSerialize["issuanceRules"] = o.IssuanceRules
	}
	return toSerialize, nil
}

type NullableReadAllCredentialIssuanceRules200ResponseEmbedded struct {
	value *ReadAllCredentialIssuanceRules200ResponseEmbedded
	isSet bool
}

func (v NullableReadAllCredentialIssuanceRules200ResponseEmbedded) Get() *ReadAllCredentialIssuanceRules200ResponseEmbedded {
	return v.value
}

func (v *NullableReadAllCredentialIssuanceRules200ResponseEmbedded) Set(val *ReadAllCredentialIssuanceRules200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllCredentialIssuanceRules200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllCredentialIssuanceRules200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllCredentialIssuanceRules200ResponseEmbedded(val *ReadAllCredentialIssuanceRules200ResponseEmbedded) *NullableReadAllCredentialIssuanceRules200ResponseEmbedded {
	return &NullableReadAllCredentialIssuanceRules200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadAllCredentialIssuanceRules200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllCredentialIssuanceRules200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
