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

// SignOnPolicyActionMFAApplicationsInnerAutoEnrollment struct for SignOnPolicyActionMFAApplicationsInnerAutoEnrollment
type SignOnPolicyActionMFAApplicationsInnerAutoEnrollment struct {
	// A boolean that specifies the enabled/disabled state of automatic MFA enrollment for the application.  When enabled, it allows automatic enrollment of the native application to MFA during the authentication flow.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFAApplicationsInnerAutoEnrollment instantiates a new SignOnPolicyActionMFAApplicationsInnerAutoEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAApplicationsInnerAutoEnrollment() *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment {
	this := SignOnPolicyActionMFAApplicationsInnerAutoEnrollment{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFAApplicationsInnerAutoEnrollmentWithDefaults instantiates a new SignOnPolicyActionMFAApplicationsInnerAutoEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAApplicationsInnerAutoEnrollmentWithDefaults() *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment {
	this := SignOnPolicyActionMFAApplicationsInnerAutoEnrollment{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment struct {
	value *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment
	isSet bool
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) Get() *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) Set(val *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment(val *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) *NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment {
	return &NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAApplicationsInnerAutoEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


