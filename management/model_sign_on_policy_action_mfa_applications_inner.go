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

// SignOnPolicyActionMFAApplicationsInner struct for SignOnPolicyActionMFAApplicationsInner
type SignOnPolicyActionMFAApplicationsInner struct {
	Id string `json:"id"`
	AutoEnrollment *SignOnPolicyActionMFAApplicationsInnerAutoEnrollment `json:"autoEnrollment,omitempty"`
	DeviceAuthorization *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization `json:"deviceAuthorization,omitempty"`
}

// NewSignOnPolicyActionMFAApplicationsInner instantiates a new SignOnPolicyActionMFAApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAApplicationsInner(id string) *SignOnPolicyActionMFAApplicationsInner {
	this := SignOnPolicyActionMFAApplicationsInner{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionMFAApplicationsInnerWithDefaults instantiates a new SignOnPolicyActionMFAApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAApplicationsInnerWithDefaults() *SignOnPolicyActionMFAApplicationsInner {
	this := SignOnPolicyActionMFAApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionMFAApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionMFAApplicationsInner) SetId(v string) {
	o.Id = v
}

// GetAutoEnrollment returns the AutoEnrollment field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAApplicationsInner) GetAutoEnrollment() SignOnPolicyActionMFAApplicationsInnerAutoEnrollment {
	if o == nil || o.AutoEnrollment == nil {
		var ret SignOnPolicyActionMFAApplicationsInnerAutoEnrollment
		return ret
	}
	return *o.AutoEnrollment
}

// GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInner) GetAutoEnrollmentOk() (*SignOnPolicyActionMFAApplicationsInnerAutoEnrollment, bool) {
	if o == nil || o.AutoEnrollment == nil {
		return nil, false
	}
	return o.AutoEnrollment, true
}

// HasAutoEnrollment returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAApplicationsInner) HasAutoEnrollment() bool {
	if o != nil && o.AutoEnrollment != nil {
		return true
	}

	return false
}

// SetAutoEnrollment gets a reference to the given SignOnPolicyActionMFAApplicationsInnerAutoEnrollment and assigns it to the AutoEnrollment field.
func (o *SignOnPolicyActionMFAApplicationsInner) SetAutoEnrollment(v SignOnPolicyActionMFAApplicationsInnerAutoEnrollment) {
	o.AutoEnrollment = &v
}

// GetDeviceAuthorization returns the DeviceAuthorization field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAApplicationsInner) GetDeviceAuthorization() SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization {
	if o == nil || o.DeviceAuthorization == nil {
		var ret SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization
		return ret
	}
	return *o.DeviceAuthorization
}

// GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAApplicationsInner) GetDeviceAuthorizationOk() (*SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization, bool) {
	if o == nil || o.DeviceAuthorization == nil {
		return nil, false
	}
	return o.DeviceAuthorization, true
}

// HasDeviceAuthorization returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAApplicationsInner) HasDeviceAuthorization() bool {
	if o != nil && o.DeviceAuthorization != nil {
		return true
	}

	return false
}

// SetDeviceAuthorization gets a reference to the given SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization and assigns it to the DeviceAuthorization field.
func (o *SignOnPolicyActionMFAApplicationsInner) SetDeviceAuthorization(v SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) {
	o.DeviceAuthorization = &v
}

func (o SignOnPolicyActionMFAApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.AutoEnrollment != nil {
		toSerialize["autoEnrollment"] = o.AutoEnrollment
	}
	if o.DeviceAuthorization != nil {
		toSerialize["deviceAuthorization"] = o.DeviceAuthorization
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAApplicationsInner struct {
	value *SignOnPolicyActionMFAApplicationsInner
	isSet bool
}

func (v NullableSignOnPolicyActionMFAApplicationsInner) Get() *SignOnPolicyActionMFAApplicationsInner {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAApplicationsInner) Set(val *SignOnPolicyActionMFAApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAApplicationsInner(val *SignOnPolicyActionMFAApplicationsInner) *NullableSignOnPolicyActionMFAApplicationsInner {
	return &NullableSignOnPolicyActionMFAApplicationsInner{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


