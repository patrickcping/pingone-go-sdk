/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// AuthorizeEditorDataPoliciesPolicyChild - struct for AuthorizeEditorDataPoliciesPolicyChild
type AuthorizeEditorDataPoliciesPolicyChild struct {
	AuthorizeEditorDataPoliciesPolicyChildPolicy *AuthorizeEditorDataPoliciesPolicyChildPolicy
	AuthorizeEditorDataPoliciesPolicyChildRule *AuthorizeEditorDataPoliciesPolicyChildRule
}

// AuthorizeEditorDataPoliciesPolicyChildPolicyAsAuthorizeEditorDataPoliciesPolicyChild is a convenience function that returns AuthorizeEditorDataPoliciesPolicyChildPolicy wrapped in AuthorizeEditorDataPoliciesPolicyChild
func AuthorizeEditorDataPoliciesPolicyChildPolicyAsAuthorizeEditorDataPoliciesPolicyChild(v *AuthorizeEditorDataPoliciesPolicyChildPolicy) AuthorizeEditorDataPoliciesPolicyChild {
	return AuthorizeEditorDataPoliciesPolicyChild{
		AuthorizeEditorDataPoliciesPolicyChildPolicy: v,
	}
}

// AuthorizeEditorDataPoliciesPolicyChildRuleAsAuthorizeEditorDataPoliciesPolicyChild is a convenience function that returns AuthorizeEditorDataPoliciesPolicyChildRule wrapped in AuthorizeEditorDataPoliciesPolicyChild
func AuthorizeEditorDataPoliciesPolicyChildRuleAsAuthorizeEditorDataPoliciesPolicyChild(v *AuthorizeEditorDataPoliciesPolicyChildRule) AuthorizeEditorDataPoliciesPolicyChild {
	return AuthorizeEditorDataPoliciesPolicyChild{
		AuthorizeEditorDataPoliciesPolicyChildRule: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthorizeEditorDataPoliciesPolicyChild) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataPoliciesPolicyChildCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataPoliciesPolicyChildPolicy = nil
	dst.AuthorizeEditorDataPoliciesPolicyChildRule = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAPOLICIESPOLICYCHILDCOMMONTYPE_POLICY:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataPoliciesPolicyChildPolicy); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPOLICIESPOLICYCHILDCOMMONTYPE_RULE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataPoliciesPolicyChildRule); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataPoliciesPolicyChild)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthorizeEditorDataPoliciesPolicyChild) MarshalJSON() ([]byte, error) {
	if src.AuthorizeEditorDataPoliciesPolicyChildPolicy != nil {
		return json.Marshal(&src.AuthorizeEditorDataPoliciesPolicyChildPolicy)
	}

	if src.AuthorizeEditorDataPoliciesPolicyChildRule != nil {
		return json.Marshal(&src.AuthorizeEditorDataPoliciesPolicyChildRule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AuthorizeEditorDataPoliciesPolicyChild) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthorizeEditorDataPoliciesPolicyChildPolicy != nil {
		return obj.AuthorizeEditorDataPoliciesPolicyChildPolicy
	}

	if obj.AuthorizeEditorDataPoliciesPolicyChildRule != nil {
		return obj.AuthorizeEditorDataPoliciesPolicyChildRule
	}

	// all schemas are nil
	return nil
}

type NullableAuthorizeEditorDataPoliciesPolicyChild struct {
	value *AuthorizeEditorDataPoliciesPolicyChild
	isSet bool
}

func (v NullableAuthorizeEditorDataPoliciesPolicyChild) Get() *AuthorizeEditorDataPoliciesPolicyChild {
	return v.value
}

func (v *NullableAuthorizeEditorDataPoliciesPolicyChild) Set(val *AuthorizeEditorDataPoliciesPolicyChild) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataPoliciesPolicyChild) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataPoliciesPolicyChild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataPoliciesPolicyChild(val *AuthorizeEditorDataPoliciesPolicyChild) *NullableAuthorizeEditorDataPoliciesPolicyChild {
	return &NullableAuthorizeEditorDataPoliciesPolicyChild{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataPoliciesPolicyChild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataPoliciesPolicyChild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


