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

// checks if the APIServerDeploymentPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerDeploymentPolicy{}

// APIServerDeploymentPolicy struct for APIServerDeploymentPolicy
type APIServerDeploymentPolicy struct {
	// The ID of the root policy.
	Id *string `json:"id,omitempty"`
}

// NewAPIServerDeploymentPolicy instantiates a new APIServerDeploymentPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerDeploymentPolicy() *APIServerDeploymentPolicy {
	this := APIServerDeploymentPolicy{}
	return &this
}

// NewAPIServerDeploymentPolicyWithDefaults instantiates a new APIServerDeploymentPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerDeploymentPolicyWithDefaults() *APIServerDeploymentPolicy {
	this := APIServerDeploymentPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIServerDeploymentPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIServerDeploymentPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIServerDeploymentPolicy) SetId(v string) {
	o.Id = &v
}

func (o APIServerDeploymentPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerDeploymentPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAPIServerDeploymentPolicy struct {
	value *APIServerDeploymentPolicy
	isSet bool
}

func (v NullableAPIServerDeploymentPolicy) Get() *APIServerDeploymentPolicy {
	return v.value
}

func (v *NullableAPIServerDeploymentPolicy) Set(val *APIServerDeploymentPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerDeploymentPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerDeploymentPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerDeploymentPolicy(val *APIServerDeploymentPolicy) *NullableAPIServerDeploymentPolicy {
	return &NullableAPIServerDeploymentPolicy{value: val, isSet: true}
}

func (v NullableAPIServerDeploymentPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerDeploymentPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
