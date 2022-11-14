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

// LicenseOrchestrate struct for LicenseOrchestrate
type LicenseOrchestrate struct {
	AllowOrchestration *bool `json:"allowOrchestration,omitempty"`
}

// NewLicenseOrchestrate instantiates a new LicenseOrchestrate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseOrchestrate() *LicenseOrchestrate {
	this := LicenseOrchestrate{}
	return &this
}

// NewLicenseOrchestrateWithDefaults instantiates a new LicenseOrchestrate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseOrchestrateWithDefaults() *LicenseOrchestrate {
	this := LicenseOrchestrate{}
	return &this
}

// GetAllowOrchestration returns the AllowOrchestration field value if set, zero value otherwise.
func (o *LicenseOrchestrate) GetAllowOrchestration() bool {
	if o == nil || isNil(o.AllowOrchestration) {
		var ret bool
		return ret
	}
	return *o.AllowOrchestration
}

// GetAllowOrchestrationOk returns a tuple with the AllowOrchestration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseOrchestrate) GetAllowOrchestrationOk() (*bool, bool) {
	if o == nil || isNil(o.AllowOrchestration) {
    return nil, false
	}
	return o.AllowOrchestration, true
}

// HasAllowOrchestration returns a boolean if a field has been set.
func (o *LicenseOrchestrate) HasAllowOrchestration() bool {
	if o != nil && !isNil(o.AllowOrchestration) {
		return true
	}

	return false
}

// SetAllowOrchestration gets a reference to the given bool and assigns it to the AllowOrchestration field.
func (o *LicenseOrchestrate) SetAllowOrchestration(v bool) {
	o.AllowOrchestration = &v
}

func (o LicenseOrchestrate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowOrchestration) {
		toSerialize["allowOrchestration"] = o.AllowOrchestration
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseOrchestrate struct {
	value *LicenseOrchestrate
	isSet bool
}

func (v NullableLicenseOrchestrate) Get() *LicenseOrchestrate {
	return v.value
}

func (v *NullableLicenseOrchestrate) Set(val *LicenseOrchestrate) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseOrchestrate) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseOrchestrate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseOrchestrate(val *LicenseOrchestrate) *NullableLicenseOrchestrate {
	return &NullableLicenseOrchestrate{value: val, isSet: true}
}

func (v NullableLicenseOrchestrate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseOrchestrate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


