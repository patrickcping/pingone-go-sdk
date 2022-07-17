/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumApplicationOIDCPKCEOption A string that specifies how PKCE request parameters are handled on the authorize request. Options are OPTIONAL PKCE code_challenge is optional and any code challenge method is acceptable. REQUIRED PKCE code_challenge is required and any code challenge method is acceptable. S256_REQUIRED PKCE code_challege is required and the code_challenge_method must be S256.
type EnumApplicationOIDCPKCEOption string

// List of EnumApplicationOIDCPKCEOption
const (
	OPTIONAL EnumApplicationOIDCPKCEOption = "OPTIONAL"
	REQUIRED EnumApplicationOIDCPKCEOption = "REQUIRED"
	S256_REQUIRED EnumApplicationOIDCPKCEOption = "S256_REQUIRED"
)

// All allowed values of EnumApplicationOIDCPKCEOption enum
var AllowedEnumApplicationOIDCPKCEOptionEnumValues = []EnumApplicationOIDCPKCEOption{
	"OPTIONAL",
	"REQUIRED",
	"S256_REQUIRED",
}

func (v *EnumApplicationOIDCPKCEOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationOIDCPKCEOption(value)
	for _, existing := range AllowedEnumApplicationOIDCPKCEOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumApplicationOIDCPKCEOption", value)
}

// NewEnumApplicationOIDCPKCEOptionFromValue returns a pointer to a valid EnumApplicationOIDCPKCEOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationOIDCPKCEOptionFromValue(v string) (*EnumApplicationOIDCPKCEOption, error) {
	ev := EnumApplicationOIDCPKCEOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationOIDCPKCEOption: valid values are %v", v, AllowedEnumApplicationOIDCPKCEOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationOIDCPKCEOption) IsValid() bool {
	for _, existing := range AllowedEnumApplicationOIDCPKCEOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationOIDCPKCEOption value
func (v EnumApplicationOIDCPKCEOption) Ptr() *EnumApplicationOIDCPKCEOption {
	return &v
}

type NullableEnumApplicationOIDCPKCEOption struct {
	value *EnumApplicationOIDCPKCEOption
	isSet bool
}

func (v NullableEnumApplicationOIDCPKCEOption) Get() *EnumApplicationOIDCPKCEOption {
	return v.value
}

func (v *NullableEnumApplicationOIDCPKCEOption) Set(val *EnumApplicationOIDCPKCEOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationOIDCPKCEOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationOIDCPKCEOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationOIDCPKCEOption(val *EnumApplicationOIDCPKCEOption) *NullableEnumApplicationOIDCPKCEOption {
	return &NullableEnumApplicationOIDCPKCEOption{value: val, isSet: true}
}

func (v NullableEnumApplicationOIDCPKCEOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationOIDCPKCEOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

