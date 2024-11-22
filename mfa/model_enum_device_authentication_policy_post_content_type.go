/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// EnumDeviceAuthenticationPolicyPostContentType the model 'EnumDeviceAuthenticationPolicyPostContentType'
type EnumDeviceAuthenticationPolicyPostContentType string

// List of EnumDeviceAuthenticationPolicyPostContentType
const (
	ENUMDEVICEAUTHENTICATIONPOLICYPOSTCONTENTTYPE_JSON                                                            EnumDeviceAuthenticationPolicyPostContentType = "application/json"
	ENUMDEVICEAUTHENTICATIONPOLICYPOSTCONTENTTYPE_VND_PINGIDENTITY_DEVICE_AUTHENTICATION_POLICY_FIDO2_MIGRATEJSON EnumDeviceAuthenticationPolicyPostContentType = "application/vnd.pingidentity.deviceAuthenticationPolicy.fido2.migrate+json"
)

// All allowed values of EnumDeviceAuthenticationPolicyPostContentType enum
var AllowedEnumDeviceAuthenticationPolicyPostContentTypeEnumValues = []EnumDeviceAuthenticationPolicyPostContentType{
	"application/json",
	"application/vnd.pingidentity.deviceAuthenticationPolicy.fido2.migrate+json",
}

func (v *EnumDeviceAuthenticationPolicyPostContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumDeviceAuthenticationPolicyPostContentType(value)
	for _, existing := range AllowedEnumDeviceAuthenticationPolicyPostContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumDeviceAuthenticationPolicyPostContentType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumDeviceAuthenticationPolicyPostContentTypeFromValue returns a pointer to a valid EnumDeviceAuthenticationPolicyPostContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumDeviceAuthenticationPolicyPostContentTypeFromValue(v string) (*EnumDeviceAuthenticationPolicyPostContentType, error) {
	ev := EnumDeviceAuthenticationPolicyPostContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumDeviceAuthenticationPolicyPostContentType: valid values are %v", v, AllowedEnumDeviceAuthenticationPolicyPostContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumDeviceAuthenticationPolicyPostContentType) IsValid() bool {
	for _, existing := range AllowedEnumDeviceAuthenticationPolicyPostContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumDeviceAuthenticationPolicyPostContentType value
func (v EnumDeviceAuthenticationPolicyPostContentType) Ptr() *EnumDeviceAuthenticationPolicyPostContentType {
	return &v
}

type NullableEnumDeviceAuthenticationPolicyPostContentType struct {
	value *EnumDeviceAuthenticationPolicyPostContentType
	isSet bool
}

func (v NullableEnumDeviceAuthenticationPolicyPostContentType) Get() *EnumDeviceAuthenticationPolicyPostContentType {
	return v.value
}

func (v *NullableEnumDeviceAuthenticationPolicyPostContentType) Set(val *EnumDeviceAuthenticationPolicyPostContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumDeviceAuthenticationPolicyPostContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumDeviceAuthenticationPolicyPostContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumDeviceAuthenticationPolicyPostContentType(val *EnumDeviceAuthenticationPolicyPostContentType) *NullableEnumDeviceAuthenticationPolicyPostContentType {
	return &NullableEnumDeviceAuthenticationPolicyPostContentType{value: val, isSet: true}
}

func (v NullableEnumDeviceAuthenticationPolicyPostContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumDeviceAuthenticationPolicyPostContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
