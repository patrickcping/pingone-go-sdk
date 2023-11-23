/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the SignOnPolicyActionMFAAllOfApplications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionMFAAllOfApplications{}

// SignOnPolicyActionMFAAllOfApplications struct for SignOnPolicyActionMFAAllOfApplications
type SignOnPolicyActionMFAAllOfApplications struct {
	Id string `json:"id"`
	AutoEnrollment *SignOnPolicyActionMFAAllOfAutoEnrollment `json:"autoEnrollment,omitempty"`
	DeviceAuthorization *SignOnPolicyActionMFAAllOfDeviceAuthorization `json:"deviceAuthorization,omitempty"`
}

// NewSignOnPolicyActionMFAAllOfApplications instantiates a new SignOnPolicyActionMFAAllOfApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAllOfApplications(id string) *SignOnPolicyActionMFAAllOfApplications {
	this := SignOnPolicyActionMFAAllOfApplications{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionMFAAllOfApplicationsWithDefaults instantiates a new SignOnPolicyActionMFAAllOfApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAllOfApplicationsWithDefaults() *SignOnPolicyActionMFAAllOfApplications {
	this := SignOnPolicyActionMFAAllOfApplications{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionMFAAllOfApplications) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfApplications) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionMFAAllOfApplications) SetId(v string) {
	o.Id = v
}

// GetAutoEnrollment returns the AutoEnrollment field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAllOfApplications) GetAutoEnrollment() SignOnPolicyActionMFAAllOfAutoEnrollment {
	if o == nil || IsNil(o.AutoEnrollment) {
		var ret SignOnPolicyActionMFAAllOfAutoEnrollment
		return ret
	}
	return *o.AutoEnrollment
}

// GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfApplications) GetAutoEnrollmentOk() (*SignOnPolicyActionMFAAllOfAutoEnrollment, bool) {
	if o == nil || IsNil(o.AutoEnrollment) {
		return nil, false
	}
	return o.AutoEnrollment, true
}

// HasAutoEnrollment returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAllOfApplications) HasAutoEnrollment() bool {
	if o != nil && !IsNil(o.AutoEnrollment) {
		return true
	}

	return false
}

// SetAutoEnrollment gets a reference to the given SignOnPolicyActionMFAAllOfAutoEnrollment and assigns it to the AutoEnrollment field.
func (o *SignOnPolicyActionMFAAllOfApplications) SetAutoEnrollment(v SignOnPolicyActionMFAAllOfAutoEnrollment) {
	o.AutoEnrollment = &v
}

// GetDeviceAuthorization returns the DeviceAuthorization field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAllOfApplications) GetDeviceAuthorization() SignOnPolicyActionMFAAllOfDeviceAuthorization {
	if o == nil || IsNil(o.DeviceAuthorization) {
		var ret SignOnPolicyActionMFAAllOfDeviceAuthorization
		return ret
	}
	return *o.DeviceAuthorization
}

// GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfApplications) GetDeviceAuthorizationOk() (*SignOnPolicyActionMFAAllOfDeviceAuthorization, bool) {
	if o == nil || IsNil(o.DeviceAuthorization) {
		return nil, false
	}
	return o.DeviceAuthorization, true
}

// HasDeviceAuthorization returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAllOfApplications) HasDeviceAuthorization() bool {
	if o != nil && !IsNil(o.DeviceAuthorization) {
		return true
	}

	return false
}

// SetDeviceAuthorization gets a reference to the given SignOnPolicyActionMFAAllOfDeviceAuthorization and assigns it to the DeviceAuthorization field.
func (o *SignOnPolicyActionMFAAllOfApplications) SetDeviceAuthorization(v SignOnPolicyActionMFAAllOfDeviceAuthorization) {
	o.DeviceAuthorization = &v
}

func (o SignOnPolicyActionMFAAllOfApplications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionMFAAllOfApplications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.AutoEnrollment) {
		toSerialize["autoEnrollment"] = o.AutoEnrollment
	}
	if !IsNil(o.DeviceAuthorization) {
		toSerialize["deviceAuthorization"] = o.DeviceAuthorization
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionMFAAllOfApplications struct {
	value *SignOnPolicyActionMFAAllOfApplications
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAllOfApplications) Get() *SignOnPolicyActionMFAAllOfApplications {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAllOfApplications) Set(val *SignOnPolicyActionMFAAllOfApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAllOfApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAllOfApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAllOfApplications(val *SignOnPolicyActionMFAAllOfApplications) *NullableSignOnPolicyActionMFAAllOfApplications {
	return &NullableSignOnPolicyActionMFAAllOfApplications{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAllOfApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAllOfApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


