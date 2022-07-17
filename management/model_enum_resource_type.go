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

// EnumResourceType A string that specifies the type of resource. Options are OPENID_CONNECT, PING_ONE_API, and CUSTOM. Only the CUSTOM resource type can be created. OPENID_CONNECT specifies the built-in platform resource for OpenID Connect. PING_ONE_API specifies the built-in platform resource for PingOne.
type EnumResourceType string

// List of EnumResourceType
const (
	OPENID_CONNECT EnumResourceType = "OPENID_CONNECT"
	PING_ONE_API EnumResourceType = "PING_ONE_API"
	CUSTOM EnumResourceType = "CUSTOM"
)

// All allowed values of EnumResourceType enum
var AllowedEnumResourceTypeEnumValues = []EnumResourceType{
	"OPENID_CONNECT",
	"PING_ONE_API",
	"CUSTOM",
}

func (v *EnumResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumResourceType(value)
	for _, existing := range AllowedEnumResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumResourceType", value)
}

// NewEnumResourceTypeFromValue returns a pointer to a valid EnumResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumResourceTypeFromValue(v string) (*EnumResourceType, error) {
	ev := EnumResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumResourceType: valid values are %v", v, AllowedEnumResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumResourceType) IsValid() bool {
	for _, existing := range AllowedEnumResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumResourceType value
func (v EnumResourceType) Ptr() *EnumResourceType {
	return &v
}

type NullableEnumResourceType struct {
	value *EnumResourceType
	isSet bool
}

func (v NullableEnumResourceType) Get() *EnumResourceType {
	return v.value
}

func (v *NullableEnumResourceType) Set(val *EnumResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumResourceType(val *EnumResourceType) *NullableEnumResourceType {
	return &NullableEnumResourceType{value: val, isSet: true}
}

func (v NullableEnumResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

