/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the APIServerDeploymentAccessControlCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerDeploymentAccessControlCustom{}

// APIServerDeploymentAccessControlCustom Defines if the operation will use custom policy rather than the \"Group\" or \"Scope\" `accessControl` requirement.
type APIServerDeploymentAccessControlCustom struct {
	// If `TRUE`, custom policy will be used for the endpoint. Defaults to `FALSE`.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewAPIServerDeploymentAccessControlCustom instantiates a new APIServerDeploymentAccessControlCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerDeploymentAccessControlCustom() *APIServerDeploymentAccessControlCustom {
	this := APIServerDeploymentAccessControlCustom{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewAPIServerDeploymentAccessControlCustomWithDefaults instantiates a new APIServerDeploymentAccessControlCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerDeploymentAccessControlCustomWithDefaults() *APIServerDeploymentAccessControlCustom {
	this := APIServerDeploymentAccessControlCustom{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *APIServerDeploymentAccessControlCustom) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentAccessControlCustom) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *APIServerDeploymentAccessControlCustom) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *APIServerDeploymentAccessControlCustom) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o APIServerDeploymentAccessControlCustom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerDeploymentAccessControlCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableAPIServerDeploymentAccessControlCustom struct {
	value *APIServerDeploymentAccessControlCustom
	isSet bool
}

func (v NullableAPIServerDeploymentAccessControlCustom) Get() *APIServerDeploymentAccessControlCustom {
	return v.value
}

func (v *NullableAPIServerDeploymentAccessControlCustom) Set(val *APIServerDeploymentAccessControlCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerDeploymentAccessControlCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerDeploymentAccessControlCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerDeploymentAccessControlCustom(val *APIServerDeploymentAccessControlCustom) *NullableAPIServerDeploymentAccessControlCustom {
	return &NullableAPIServerDeploymentAccessControlCustom{value: val, isSet: true}
}

func (v NullableAPIServerDeploymentAccessControlCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerDeploymentAccessControlCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


