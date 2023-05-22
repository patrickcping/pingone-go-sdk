/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumFormQrCodeType A string that specifies the QR Code type.
type EnumFormQrCodeType string

// List of EnumFormQrCodeType
const (
	ENUMFORMQRCODETYPE_MFA_AUTH EnumFormQrCodeType = "MFA_AUTH"
)

// All allowed values of EnumFormQrCodeType enum
var AllowedEnumFormQrCodeTypeEnumValues = []EnumFormQrCodeType{
	"MFA_AUTH",
}

func (v *EnumFormQrCodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFormQrCodeType(value)
	for _, existing := range AllowedEnumFormQrCodeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFormQrCodeType", value)
}

// NewEnumFormQrCodeTypeFromValue returns a pointer to a valid EnumFormQrCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFormQrCodeTypeFromValue(v string) (*EnumFormQrCodeType, error) {
	ev := EnumFormQrCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFormQrCodeType: valid values are %v", v, AllowedEnumFormQrCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFormQrCodeType) IsValid() bool {
	for _, existing := range AllowedEnumFormQrCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFormQrCodeType value
func (v EnumFormQrCodeType) Ptr() *EnumFormQrCodeType {
	return &v
}

type NullableEnumFormQrCodeType struct {
	value *EnumFormQrCodeType
	isSet bool
}

func (v NullableEnumFormQrCodeType) Get() *EnumFormQrCodeType {
	return v.value
}

func (v *NullableEnumFormQrCodeType) Set(val *EnumFormQrCodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFormQrCodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFormQrCodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFormQrCodeType(val *EnumFormQrCodeType) *NullableEnumFormQrCodeType {
	return &NullableEnumFormQrCodeType{value: val, isSet: true}
}

func (v NullableEnumFormQrCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFormQrCodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

