/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization struct for SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization
type SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization struct {
	// A boolean that specifies the enabled/disabled state of automatic MFA for native devices paired with the user for the specified application.
	Enabled *bool `json:"enabled,omitempty"`
	ExtraVerification *EnumSignOnPolicyExtraVerification `json:"extraVerification,omitempty"`
}

// NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization instantiates a new SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization() *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization {
	this := SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization{}
	var enabled bool = false
	this.Enabled = &enabled
	var extraVerification EnumSignOnPolicyExtraVerification = DISABLED
	this.ExtraVerification = &extraVerification
	return &this
}

// NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorizationWithDefaults instantiates a new SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorizationWithDefaults() *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization {
	this := SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization{}
	var enabled bool = false
	this.Enabled = &enabled
	var extraVerification EnumSignOnPolicyExtraVerification = DISABLED
	this.ExtraVerification = &extraVerification
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExtraVerification returns the ExtraVerification field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerification() EnumSignOnPolicyExtraVerification {
	if o == nil || o.ExtraVerification == nil {
		var ret EnumSignOnPolicyExtraVerification
		return ret
	}
	return *o.ExtraVerification
}

// GetExtraVerificationOk returns a tuple with the ExtraVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerificationOk() (*EnumSignOnPolicyExtraVerification, bool) {
	if o == nil || o.ExtraVerification == nil {
		return nil, false
	}
	return o.ExtraVerification, true
}

// HasExtraVerification returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) HasExtraVerification() bool {
	if o != nil && o.ExtraVerification != nil {
		return true
	}

	return false
}

// SetExtraVerification gets a reference to the given EnumSignOnPolicyExtraVerification and assigns it to the ExtraVerification field.
func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) SetExtraVerification(v EnumSignOnPolicyExtraVerification) {
	o.ExtraVerification = &v
}

func (o SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExtraVerification != nil {
		toSerialize["extraVerification"] = o.ExtraVerification
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization struct {
	value *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization
	isSet bool
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) Get() *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) Set(val *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization(val *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) *NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization {
	return &NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


