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

// checks if the FormQrCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormQrCode{}

// FormQrCode struct for FormQrCode
type FormQrCode struct {
	QrCodeType EnumFormQrCodeType `json:"QrCodeType"`
	Alignment EnumFormQrCodeAlignment `json:"alignment"`
	// A boolean that specifies the border visibility.
	ShowBorder bool `json:"showBorder"`
}

// NewFormQrCode instantiates a new FormQrCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormQrCode(qrCodeType EnumFormQrCodeType, alignment EnumFormQrCodeAlignment, showBorder bool) *FormQrCode {
	this := FormQrCode{}
	this.QrCodeType = qrCodeType
	this.Alignment = alignment
	this.ShowBorder = showBorder
	return &this
}

// NewFormQrCodeWithDefaults instantiates a new FormQrCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormQrCodeWithDefaults() *FormQrCode {
	this := FormQrCode{}
	return &this
}

// GetQrCodeType returns the QrCodeType field value
func (o *FormQrCode) GetQrCodeType() EnumFormQrCodeType {
	if o == nil {
		var ret EnumFormQrCodeType
		return ret
	}

	return o.QrCodeType
}

// GetQrCodeTypeOk returns a tuple with the QrCodeType field value
// and a boolean to check if the value has been set.
func (o *FormQrCode) GetQrCodeTypeOk() (*EnumFormQrCodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QrCodeType, true
}

// SetQrCodeType sets field value
func (o *FormQrCode) SetQrCodeType(v EnumFormQrCodeType) {
	o.QrCodeType = v
}

// GetAlignment returns the Alignment field value
func (o *FormQrCode) GetAlignment() EnumFormQrCodeAlignment {
	if o == nil {
		var ret EnumFormQrCodeAlignment
		return ret
	}

	return o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value
// and a boolean to check if the value has been set.
func (o *FormQrCode) GetAlignmentOk() (*EnumFormQrCodeAlignment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alignment, true
}

// SetAlignment sets field value
func (o *FormQrCode) SetAlignment(v EnumFormQrCodeAlignment) {
	o.Alignment = v
}

// GetShowBorder returns the ShowBorder field value
func (o *FormQrCode) GetShowBorder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowBorder
}

// GetShowBorderOk returns a tuple with the ShowBorder field value
// and a boolean to check if the value has been set.
func (o *FormQrCode) GetShowBorderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowBorder, true
}

// SetShowBorder sets field value
func (o *FormQrCode) SetShowBorder(v bool) {
	o.ShowBorder = v
}

func (o FormQrCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormQrCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["QrCodeType"] = o.QrCodeType
	toSerialize["alignment"] = o.Alignment
	toSerialize["showBorder"] = o.ShowBorder
	return toSerialize, nil
}

type NullableFormQrCode struct {
	value *FormQrCode
	isSet bool
}

func (v NullableFormQrCode) Get() *FormQrCode {
	return v.value
}

func (v *NullableFormQrCode) Set(val *FormQrCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFormQrCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFormQrCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormQrCode(val *FormQrCode) *NullableFormQrCode {
	return &NullableFormQrCode{value: val, isSet: true}
}

func (v NullableFormQrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormQrCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


