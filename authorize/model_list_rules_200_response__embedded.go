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

// checks if the ListRules200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRules200ResponseEmbedded{}

// ListRules200ResponseEmbedded struct for ListRules200ResponseEmbedded
type ListRules200ResponseEmbedded struct {
	AuthorizationRules []AuthorizeEditorDataRulesReferenceableRuleDTO `json:"authorizationRules,omitempty"`
}

// NewListRules200ResponseEmbedded instantiates a new ListRules200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRules200ResponseEmbedded() *ListRules200ResponseEmbedded {
	this := ListRules200ResponseEmbedded{}
	return &this
}

// NewListRules200ResponseEmbeddedWithDefaults instantiates a new ListRules200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRules200ResponseEmbeddedWithDefaults() *ListRules200ResponseEmbedded {
	this := ListRules200ResponseEmbedded{}
	return &this
}

// GetAuthorizationRules returns the AuthorizationRules field value if set, zero value otherwise.
func (o *ListRules200ResponseEmbedded) GetAuthorizationRules() []AuthorizeEditorDataRulesReferenceableRuleDTO {
	if o == nil || IsNil(o.AuthorizationRules) {
		var ret []AuthorizeEditorDataRulesReferenceableRuleDTO
		return ret
	}
	return o.AuthorizationRules
}

// GetAuthorizationRulesOk returns a tuple with the AuthorizationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRules200ResponseEmbedded) GetAuthorizationRulesOk() ([]AuthorizeEditorDataRulesReferenceableRuleDTO, bool) {
	if o == nil || IsNil(o.AuthorizationRules) {
		return nil, false
	}
	return o.AuthorizationRules, true
}

// HasAuthorizationRules returns a boolean if a field has been set.
func (o *ListRules200ResponseEmbedded) HasAuthorizationRules() bool {
	if o != nil && !IsNil(o.AuthorizationRules) {
		return true
	}

	return false
}

// SetAuthorizationRules gets a reference to the given []AuthorizeEditorDataRulesReferenceableRuleDTO and assigns it to the AuthorizationRules field.
func (o *ListRules200ResponseEmbedded) SetAuthorizationRules(v []AuthorizeEditorDataRulesReferenceableRuleDTO) {
	o.AuthorizationRules = v
}

func (o ListRules200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRules200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationRules) {
		toSerialize["authorizationRules"] = o.AuthorizationRules
	}
	return toSerialize, nil
}

type NullableListRules200ResponseEmbedded struct {
	value *ListRules200ResponseEmbedded
	isSet bool
}

func (v NullableListRules200ResponseEmbedded) Get() *ListRules200ResponseEmbedded {
	return v.value
}

func (v *NullableListRules200ResponseEmbedded) Set(val *ListRules200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListRules200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListRules200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRules200ResponseEmbedded(val *ListRules200ResponseEmbedded) *NullableListRules200ResponseEmbedded {
	return &NullableListRules200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListRules200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRules200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


