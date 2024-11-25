/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumAlertChannelAlertType Alert channel alerting types.
type EnumAlertChannelAlertType string

// List of EnumAlertChannelAlertType
const (
	ENUMALERTCHANNELALERTTYPE_CERTIFICATE_EXPIRED EnumAlertChannelAlertType = "CERTIFICATE_EXPIRED"
	ENUMALERTCHANNELALERTTYPE_CERTIFICATE_EXPIRING EnumAlertChannelAlertType = "CERTIFICATE_EXPIRING"
	ENUMALERTCHANNELALERTTYPE_KEY_PAIR_EXPIRED EnumAlertChannelAlertType = "KEY_PAIR_EXPIRED"
	ENUMALERTCHANNELALERTTYPE_KEY_PAIR_EXPIRING EnumAlertChannelAlertType = "KEY_PAIR_EXPIRING"
	ENUMALERTCHANNELALERTTYPE_GATEWAY_VERSION_DEPRECATED EnumAlertChannelAlertType = "GATEWAY_VERSION_DEPRECATED"
	ENUMALERTCHANNELALERTTYPE_GATEWAY_VERSION_DEPRECATING EnumAlertChannelAlertType = "GATEWAY_VERSION_DEPRECATING"
	ENUMALERTCHANNELALERTTYPE_LICENSE_90_PERCENT_USER_SOFT_LIMIT EnumAlertChannelAlertType = "LICENSE_90_PERCENT_USER_SOFT_LIMIT"
	ENUMALERTCHANNELALERTTYPE_LICENSE_EXPIRED EnumAlertChannelAlertType = "LICENSE_EXPIRED"
	ENUMALERTCHANNELALERTTYPE_LICENSE_EXPIRING EnumAlertChannelAlertType = "LICENSE_EXPIRING"
	ENUMALERTCHANNELALERTTYPE_LICENSE_ROTATED EnumAlertChannelAlertType = "LICENSE_ROTATED"
	ENUMALERTCHANNELALERTTYPE_LICENSE_USER_HARD_LIMIT_EXCEEDED EnumAlertChannelAlertType = "LICENSE_USER_HARD_LIMIT_EXCEEDED"
	ENUMALERTCHANNELALERTTYPE_LICENSE_USER_SOFT_LIMIT_EXCEEDED EnumAlertChannelAlertType = "LICENSE_USER_SOFT_LIMIT_EXCEEDED"
	ENUMALERTCHANNELALERTTYPE_RISK_CONFIGURATION EnumAlertChannelAlertType = "RISK_CONFIGURATION"
	ENUMALERTCHANNELALERTTYPE_SUSPICIOUS_TRAFFIC EnumAlertChannelAlertType = "SUSPICIOUS_TRAFFIC"
)

// All allowed values of EnumAlertChannelAlertType enum
var AllowedEnumAlertChannelAlertTypeEnumValues = []EnumAlertChannelAlertType{
	"CERTIFICATE_EXPIRED",
	"CERTIFICATE_EXPIRING",
	"KEY_PAIR_EXPIRED",
	"KEY_PAIR_EXPIRING",
	"GATEWAY_VERSION_DEPRECATED",
	"GATEWAY_VERSION_DEPRECATING",
	"LICENSE_90_PERCENT_USER_SOFT_LIMIT",
	"LICENSE_EXPIRED",
	"LICENSE_EXPIRING",
	"LICENSE_ROTATED",
	"LICENSE_USER_HARD_LIMIT_EXCEEDED",
	"LICENSE_USER_SOFT_LIMIT_EXCEEDED",
	"RISK_CONFIGURATION",
	"SUSPICIOUS_TRAFFIC",
}

func (v *EnumAlertChannelAlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAlertChannelAlertType(value)
	for _, existing := range AllowedEnumAlertChannelAlertTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAlertChannelAlertType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAlertChannelAlertTypeFromValue returns a pointer to a valid EnumAlertChannelAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAlertChannelAlertTypeFromValue(v string) (*EnumAlertChannelAlertType, error) {
	ev := EnumAlertChannelAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAlertChannelAlertType: valid values are %v", v, AllowedEnumAlertChannelAlertTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAlertChannelAlertType) IsValid() bool {
	for _, existing := range AllowedEnumAlertChannelAlertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAlertChannelAlertType value
func (v EnumAlertChannelAlertType) Ptr() *EnumAlertChannelAlertType {
	return &v
}

type NullableEnumAlertChannelAlertType struct {
	value *EnumAlertChannelAlertType
	isSet bool
}

func (v NullableEnumAlertChannelAlertType) Get() *EnumAlertChannelAlertType {
	return v.value
}

func (v *NullableEnumAlertChannelAlertType) Set(val *EnumAlertChannelAlertType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAlertChannelAlertType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAlertChannelAlertType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAlertChannelAlertType(val *EnumAlertChannelAlertType) *NullableEnumAlertChannelAlertType {
	return &NullableEnumAlertChannelAlertType{value: val, isSet: true}
}

func (v NullableEnumAlertChannelAlertType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAlertChannelAlertType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

