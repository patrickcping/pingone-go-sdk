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

// EnumFIDOAttestationRequirements Determines whether attestation is requested from the authenticator, and whether this information is used to restrict authenticator usage. Can take one of these values: `NONE` - attestation is not requested `AUDIT_ONLY` - Attestation is requested and the information is used for logging purposes, but the information is not used for filtering authenticators `GLOBAL` - all entries in the MDS table can be used for authentication `CERTIFIED` - only FIDO-certified authenticators can be used `SPECIFIC` - only specific authenticators can be used. Used in conjunction with `allowedAuthenticators`.
type EnumFIDOAttestationRequirements string

// List of EnumFIDOAttestationRequirements
const (
	ENUMFIDOATTESTATIONREQUIREMENTS_NONE       EnumFIDOAttestationRequirements = "NONE"
	ENUMFIDOATTESTATIONREQUIREMENTS_AUDIT_ONLY EnumFIDOAttestationRequirements = "AUDIT_ONLY"
	ENUMFIDOATTESTATIONREQUIREMENTS_GLOBAL     EnumFIDOAttestationRequirements = "GLOBAL"
	ENUMFIDOATTESTATIONREQUIREMENTS_CERTIFIED  EnumFIDOAttestationRequirements = "CERTIFIED"
	ENUMFIDOATTESTATIONREQUIREMENTS_SPECIFIC   EnumFIDOAttestationRequirements = "SPECIFIC"
)

// All allowed values of EnumFIDOAttestationRequirements enum
var AllowedEnumFIDOAttestationRequirementsEnumValues = []EnumFIDOAttestationRequirements{
	"NONE",
	"AUDIT_ONLY",
	"GLOBAL",
	"CERTIFIED",
	"SPECIFIC",
}

func (v *EnumFIDOAttestationRequirements) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFIDOAttestationRequirements(value)
	for _, existing := range AllowedEnumFIDOAttestationRequirementsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumFIDOAttestationRequirements(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumFIDOAttestationRequirementsFromValue returns a pointer to a valid EnumFIDOAttestationRequirements
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFIDOAttestationRequirementsFromValue(v string) (*EnumFIDOAttestationRequirements, error) {
	ev := EnumFIDOAttestationRequirements(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFIDOAttestationRequirements: valid values are %v", v, AllowedEnumFIDOAttestationRequirementsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFIDOAttestationRequirements) IsValid() bool {
	for _, existing := range AllowedEnumFIDOAttestationRequirementsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFIDOAttestationRequirements value
func (v EnumFIDOAttestationRequirements) Ptr() *EnumFIDOAttestationRequirements {
	return &v
}

type NullableEnumFIDOAttestationRequirements struct {
	value *EnumFIDOAttestationRequirements
	isSet bool
}

func (v NullableEnumFIDOAttestationRequirements) Get() *EnumFIDOAttestationRequirements {
	return v.value
}

func (v *NullableEnumFIDOAttestationRequirements) Set(val *EnumFIDOAttestationRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFIDOAttestationRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFIDOAttestationRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFIDOAttestationRequirements(val *EnumFIDOAttestationRequirements) *NullableEnumFIDOAttestationRequirements {
	return &NullableEnumFIDOAttestationRequirements{value: val, isSet: true}
}

func (v NullableEnumFIDOAttestationRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFIDOAttestationRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
