/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the LicenseAuthorize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseAuthorize{}

// LicenseAuthorize struct for LicenseAuthorize
type LicenseAuthorize struct {
	// A read-only boolean that specifies whether to enable the PingOne Authorize API access management feature.
	AllowApiAccessManagement *bool `json:"allowApiAccessManagement,omitempty"`
	// A read-only boolean that specifies whether to enable the PingOne Authorize dynamic authorization feature.
	AllowDynamicAuthorization *bool `json:"allowDynamicAuthorization,omitempty"`
}

// NewLicenseAuthorize instantiates a new LicenseAuthorize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAuthorize() *LicenseAuthorize {
	this := LicenseAuthorize{}
	return &this
}

// NewLicenseAuthorizeWithDefaults instantiates a new LicenseAuthorize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAuthorizeWithDefaults() *LicenseAuthorize {
	this := LicenseAuthorize{}
	return &this
}

// GetAllowApiAccessManagement returns the AllowApiAccessManagement field value if set, zero value otherwise.
func (o *LicenseAuthorize) GetAllowApiAccessManagement() bool {
	if o == nil || IsNil(o.AllowApiAccessManagement) {
		var ret bool
		return ret
	}
	return *o.AllowApiAccessManagement
}

// GetAllowApiAccessManagementOk returns a tuple with the AllowApiAccessManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAuthorize) GetAllowApiAccessManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowApiAccessManagement) {
		return nil, false
	}
	return o.AllowApiAccessManagement, true
}

// HasAllowApiAccessManagement returns a boolean if a field has been set.
func (o *LicenseAuthorize) HasAllowApiAccessManagement() bool {
	if o != nil && !IsNil(o.AllowApiAccessManagement) {
		return true
	}

	return false
}

// SetAllowApiAccessManagement gets a reference to the given bool and assigns it to the AllowApiAccessManagement field.
func (o *LicenseAuthorize) SetAllowApiAccessManagement(v bool) {
	o.AllowApiAccessManagement = &v
}

// GetAllowDynamicAuthorization returns the AllowDynamicAuthorization field value if set, zero value otherwise.
func (o *LicenseAuthorize) GetAllowDynamicAuthorization() bool {
	if o == nil || IsNil(o.AllowDynamicAuthorization) {
		var ret bool
		return ret
	}
	return *o.AllowDynamicAuthorization
}

// GetAllowDynamicAuthorizationOk returns a tuple with the AllowDynamicAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAuthorize) GetAllowDynamicAuthorizationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDynamicAuthorization) {
		return nil, false
	}
	return o.AllowDynamicAuthorization, true
}

// HasAllowDynamicAuthorization returns a boolean if a field has been set.
func (o *LicenseAuthorize) HasAllowDynamicAuthorization() bool {
	if o != nil && !IsNil(o.AllowDynamicAuthorization) {
		return true
	}

	return false
}

// SetAllowDynamicAuthorization gets a reference to the given bool and assigns it to the AllowDynamicAuthorization field.
func (o *LicenseAuthorize) SetAllowDynamicAuthorization(v bool) {
	o.AllowDynamicAuthorization = &v
}

func (o LicenseAuthorize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseAuthorize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowApiAccessManagement) {
		toSerialize["allowApiAccessManagement"] = o.AllowApiAccessManagement
	}
	if !IsNil(o.AllowDynamicAuthorization) {
		toSerialize["allowDynamicAuthorization"] = o.AllowDynamicAuthorization
	}
	return toSerialize, nil
}

type NullableLicenseAuthorize struct {
	value *LicenseAuthorize
	isSet bool
}

func (v NullableLicenseAuthorize) Get() *LicenseAuthorize {
	return v.value
}

func (v *NullableLicenseAuthorize) Set(val *LicenseAuthorize) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAuthorize) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAuthorize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAuthorize(val *LicenseAuthorize) *NullableLicenseAuthorize {
	return &NullableLicenseAuthorize{value: val, isSet: true}
}

func (v NullableLicenseAuthorize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAuthorize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


