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

// checks if the FormElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormElement{}

// FormElement struct for FormElement
type FormElement struct {
	// A boolean that specifies whether the linked directory attribute is disabled.
	AttributeDisabled *bool `json:"attributeDisabled,omitempty"`
	// A string that specifies an identifier for the field component.
	Key string `json:"key"`
	// A string of escaped JSON that is designed to store a series of text and translatable keys.
	Label *string `json:"label,omitempty"`
	LabelMode *EnumFormElementLabelMode `json:"labelMode,omitempty"`
	// A boolean that specifies whether the field is required.
	Required bool `json:"required"`
	// A boolean that specifies whether the end user can type an entry that is not in a predefined list.
	OtherOptionEnabled *bool `json:"otherOptionEnabled,omitempty"`
	// A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list.
	OtherOptionKey *string `json:"otherOptionKey,omitempty"`
	// A string that specifies the label for a custom or \"other\" choice in a list.
	OtherOptionlabel *string `json:"otherOptionlabel,omitempty"`
	// A string that specifies the label for the other option in drop-down controls.
	OtherOptionInputlabel *string `json:"otherOptionInputlabel,omitempty"`
	// A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute.
	OtherOptionAttributeDisabled *bool `json:"otherOptionAttributeDisabled,omitempty"`
}

// NewFormElement instantiates a new FormElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormElement(key string, required bool) *FormElement {
	this := FormElement{}
	this.Key = key
	this.Required = required
	return &this
}

// NewFormElementWithDefaults instantiates a new FormElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormElementWithDefaults() *FormElement {
	this := FormElement{}
	return &this
}

// GetAttributeDisabled returns the AttributeDisabled field value if set, zero value otherwise.
func (o *FormElement) GetAttributeDisabled() bool {
	if o == nil || IsNil(o.AttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.AttributeDisabled
}

// GetAttributeDisabledOk returns a tuple with the AttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AttributeDisabled) {
		return nil, false
	}
	return o.AttributeDisabled, true
}

// HasAttributeDisabled returns a boolean if a field has been set.
func (o *FormElement) HasAttributeDisabled() bool {
	if o != nil && !IsNil(o.AttributeDisabled) {
		return true
	}

	return false
}

// SetAttributeDisabled gets a reference to the given bool and assigns it to the AttributeDisabled field.
func (o *FormElement) SetAttributeDisabled(v bool) {
	o.AttributeDisabled = &v
}

// GetKey returns the Key field value
func (o *FormElement) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FormElement) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FormElement) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FormElement) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FormElement) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FormElement) SetLabel(v string) {
	o.Label = &v
}

// GetLabelMode returns the LabelMode field value if set, zero value otherwise.
func (o *FormElement) GetLabelMode() EnumFormElementLabelMode {
	if o == nil || IsNil(o.LabelMode) {
		var ret EnumFormElementLabelMode
		return ret
	}
	return *o.LabelMode
}

// GetLabelModeOk returns a tuple with the LabelMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetLabelModeOk() (*EnumFormElementLabelMode, bool) {
	if o == nil || IsNil(o.LabelMode) {
		return nil, false
	}
	return o.LabelMode, true
}

// HasLabelMode returns a boolean if a field has been set.
func (o *FormElement) HasLabelMode() bool {
	if o != nil && !IsNil(o.LabelMode) {
		return true
	}

	return false
}

// SetLabelMode gets a reference to the given EnumFormElementLabelMode and assigns it to the LabelMode field.
func (o *FormElement) SetLabelMode(v EnumFormElementLabelMode) {
	o.LabelMode = &v
}

// GetRequired returns the Required field value
func (o *FormElement) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FormElement) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FormElement) SetRequired(v bool) {
	o.Required = v
}

// GetOtherOptionEnabled returns the OtherOptionEnabled field value if set, zero value otherwise.
func (o *FormElement) GetOtherOptionEnabled() bool {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionEnabled
}

// GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetOtherOptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		return nil, false
	}
	return o.OtherOptionEnabled, true
}

// HasOtherOptionEnabled returns a boolean if a field has been set.
func (o *FormElement) HasOtherOptionEnabled() bool {
	if o != nil && !IsNil(o.OtherOptionEnabled) {
		return true
	}

	return false
}

// SetOtherOptionEnabled gets a reference to the given bool and assigns it to the OtherOptionEnabled field.
func (o *FormElement) SetOtherOptionEnabled(v bool) {
	o.OtherOptionEnabled = &v
}

// GetOtherOptionKey returns the OtherOptionKey field value if set, zero value otherwise.
func (o *FormElement) GetOtherOptionKey() string {
	if o == nil || IsNil(o.OtherOptionKey) {
		var ret string
		return ret
	}
	return *o.OtherOptionKey
}

// GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetOtherOptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionKey) {
		return nil, false
	}
	return o.OtherOptionKey, true
}

// HasOtherOptionKey returns a boolean if a field has been set.
func (o *FormElement) HasOtherOptionKey() bool {
	if o != nil && !IsNil(o.OtherOptionKey) {
		return true
	}

	return false
}

// SetOtherOptionKey gets a reference to the given string and assigns it to the OtherOptionKey field.
func (o *FormElement) SetOtherOptionKey(v string) {
	o.OtherOptionKey = &v
}

// GetOtherOptionlabel returns the OtherOptionlabel field value if set, zero value otherwise.
func (o *FormElement) GetOtherOptionlabel() string {
	if o == nil || IsNil(o.OtherOptionlabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionlabel
}

// GetOtherOptionlabelOk returns a tuple with the OtherOptionlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetOtherOptionlabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionlabel) {
		return nil, false
	}
	return o.OtherOptionlabel, true
}

// HasOtherOptionlabel returns a boolean if a field has been set.
func (o *FormElement) HasOtherOptionlabel() bool {
	if o != nil && !IsNil(o.OtherOptionlabel) {
		return true
	}

	return false
}

// SetOtherOptionlabel gets a reference to the given string and assigns it to the OtherOptionlabel field.
func (o *FormElement) SetOtherOptionlabel(v string) {
	o.OtherOptionlabel = &v
}

// GetOtherOptionInputlabel returns the OtherOptionInputlabel field value if set, zero value otherwise.
func (o *FormElement) GetOtherOptionInputlabel() string {
	if o == nil || IsNil(o.OtherOptionInputlabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionInputlabel
}

// GetOtherOptionInputlabelOk returns a tuple with the OtherOptionInputlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetOtherOptionInputlabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionInputlabel) {
		return nil, false
	}
	return o.OtherOptionInputlabel, true
}

// HasOtherOptionInputlabel returns a boolean if a field has been set.
func (o *FormElement) HasOtherOptionInputlabel() bool {
	if o != nil && !IsNil(o.OtherOptionInputlabel) {
		return true
	}

	return false
}

// SetOtherOptionInputlabel gets a reference to the given string and assigns it to the OtherOptionInputlabel field.
func (o *FormElement) SetOtherOptionInputlabel(v string) {
	o.OtherOptionInputlabel = &v
}

// GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field value if set, zero value otherwise.
func (o *FormElement) GetOtherOptionAttributeDisabled() bool {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionAttributeDisabled
}

// GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElement) GetOtherOptionAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		return nil, false
	}
	return o.OtherOptionAttributeDisabled, true
}

// HasOtherOptionAttributeDisabled returns a boolean if a field has been set.
func (o *FormElement) HasOtherOptionAttributeDisabled() bool {
	if o != nil && !IsNil(o.OtherOptionAttributeDisabled) {
		return true
	}

	return false
}

// SetOtherOptionAttributeDisabled gets a reference to the given bool and assigns it to the OtherOptionAttributeDisabled field.
func (o *FormElement) SetOtherOptionAttributeDisabled(v bool) {
	o.OtherOptionAttributeDisabled = &v
}

func (o FormElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: attributeDisabled is readOnly
	toSerialize["key"] = o.Key
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LabelMode) {
		toSerialize["labelMode"] = o.LabelMode
	}
	toSerialize["required"] = o.Required
	if !IsNil(o.OtherOptionEnabled) {
		toSerialize["otherOptionEnabled"] = o.OtherOptionEnabled
	}
	if !IsNil(o.OtherOptionKey) {
		toSerialize["otherOptionKey"] = o.OtherOptionKey
	}
	if !IsNil(o.OtherOptionlabel) {
		toSerialize["otherOptionlabel"] = o.OtherOptionlabel
	}
	if !IsNil(o.OtherOptionInputlabel) {
		toSerialize["otherOptionInputlabel"] = o.OtherOptionInputlabel
	}
	if !IsNil(o.OtherOptionAttributeDisabled) {
		toSerialize["otherOptionAttributeDisabled"] = o.OtherOptionAttributeDisabled
	}
	return toSerialize, nil
}

type NullableFormElement struct {
	value *FormElement
	isSet bool
}

func (v NullableFormElement) Get() *FormElement {
	return v.value
}

func (v *NullableFormElement) Set(val *FormElement) {
	v.value = val
	v.isSet = true
}

func (v NullableFormElement) IsSet() bool {
	return v.isSet
}

func (v *NullableFormElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormElement(val *FormElement) *NullableFormElement {
	return &NullableFormElement{value: val, isSet: true}
}

func (v NullableFormElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


