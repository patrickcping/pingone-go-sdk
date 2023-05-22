/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the FormFieldTextAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldTextAllOf{}

// FormFieldTextAllOf struct for FormFieldTextAllOf
type FormFieldTextAllOf struct {
	Layout *EnumFormElementLayout `json:"layout,omitempty"`
	// An array of strings that specifies the unique list of options. This is a required property when the type is `RADIO`, `CHECKBOX`, or `DROPDOWN`.
	Options []string `json:"options,omitempty"`
	Validation FormElementValidation `json:"validation"`
}

// NewFormFieldTextAllOf instantiates a new FormFieldTextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldTextAllOf(validation FormElementValidation) *FormFieldTextAllOf {
	this := FormFieldTextAllOf{}
	this.Validation = validation
	return &this
}

// NewFormFieldTextAllOfWithDefaults instantiates a new FormFieldTextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldTextAllOfWithDefaults() *FormFieldTextAllOf {
	this := FormFieldTextAllOf{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *FormFieldTextAllOf) GetLayout() EnumFormElementLayout {
	if o == nil || IsNil(o.Layout) {
		var ret EnumFormElementLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldTextAllOf) GetLayoutOk() (*EnumFormElementLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *FormFieldTextAllOf) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given EnumFormElementLayout and assigns it to the Layout field.
func (o *FormFieldTextAllOf) SetLayout(v EnumFormElementLayout) {
	o.Layout = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FormFieldTextAllOf) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldTextAllOf) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FormFieldTextAllOf) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *FormFieldTextAllOf) SetOptions(v []string) {
	o.Options = v
}

// GetValidation returns the Validation field value
func (o *FormFieldTextAllOf) GetValidation() FormElementValidation {
	if o == nil {
		var ret FormElementValidation
		return ret
	}

	return o.Validation
}

// GetValidationOk returns a tuple with the Validation field value
// and a boolean to check if the value has been set.
func (o *FormFieldTextAllOf) GetValidationOk() (*FormElementValidation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validation, true
}

// SetValidation sets field value
func (o *FormFieldTextAllOf) SetValidation(v FormElementValidation) {
	o.Validation = v
}

func (o FormFieldTextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldTextAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["validation"] = o.Validation
	return toSerialize, nil
}

type NullableFormFieldTextAllOf struct {
	value *FormFieldTextAllOf
	isSet bool
}

func (v NullableFormFieldTextAllOf) Get() *FormFieldTextAllOf {
	return v.value
}

func (v *NullableFormFieldTextAllOf) Set(val *FormFieldTextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldTextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldTextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldTextAllOf(val *FormFieldTextAllOf) *NullableFormFieldTextAllOf {
	return &NullableFormFieldTextAllOf{value: val, isSet: true}
}

func (v NullableFormFieldTextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldTextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


