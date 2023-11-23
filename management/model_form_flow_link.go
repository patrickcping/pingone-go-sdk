/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the FormFlowLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFlowLink{}

// FormFlowLink struct for FormFlowLink
type FormFlowLink struct {
	// A string that specifies an identifier for the field component.
	Key string `json:"key"`
	// A string that specifies the link label.
	Label string `json:"label"`
	Styles *FormFlowLinkStyles `json:"styles,omitempty"`
}

// NewFormFlowLink instantiates a new FormFlowLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFlowLink(key string, label string) *FormFlowLink {
	this := FormFlowLink{}
	this.Key = key
	this.Label = label
	return &this
}

// NewFormFlowLinkWithDefaults instantiates a new FormFlowLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFlowLinkWithDefaults() *FormFlowLink {
	this := FormFlowLink{}
	return &this
}

// GetKey returns the Key field value
func (o *FormFlowLink) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FormFlowLink) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FormFlowLink) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *FormFlowLink) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FormFlowLink) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FormFlowLink) SetLabel(v string) {
	o.Label = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *FormFlowLink) GetStyles() FormFlowLinkStyles {
	if o == nil || IsNil(o.Styles) {
		var ret FormFlowLinkStyles
		return ret
	}
	return *o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFlowLink) GetStylesOk() (*FormFlowLinkStyles, bool) {
	if o == nil || IsNil(o.Styles) {
		return nil, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *FormFlowLink) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given FormFlowLinkStyles and assigns it to the Styles field.
func (o *FormFlowLink) SetStyles(v FormFlowLinkStyles) {
	o.Styles = &v
}

func (o FormFlowLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFlowLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	return toSerialize, nil
}

type NullableFormFlowLink struct {
	value *FormFlowLink
	isSet bool
}

func (v NullableFormFlowLink) Get() *FormFlowLink {
	return v.value
}

func (v *NullableFormFlowLink) Set(val *FormFlowLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFlowLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFlowLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFlowLink(val *FormFlowLink) *NullableFormFlowLink {
	return &NullableFormFlowLink{value: val, isSet: true}
}

func (v NullableFormFlowLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFlowLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


