/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the ObjectRetry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectRetry{}

// ObjectRetry struct for ObjectRetry
type ObjectRetry struct {
	// The number of retries permitted.
	Attempts *int32 `json:"attempts,omitempty"`
}

// NewObjectRetry instantiates a new ObjectRetry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRetry() *ObjectRetry {
	this := ObjectRetry{}
	return &this
}

// NewObjectRetryWithDefaults instantiates a new ObjectRetry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRetryWithDefaults() *ObjectRetry {
	this := ObjectRetry{}
	return &this
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *ObjectRetry) GetAttempts() int32 {
	if o == nil || IsNil(o.Attempts) {
		var ret int32
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRetry) GetAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.Attempts) {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *ObjectRetry) HasAttempts() bool {
	if o != nil && !IsNil(o.Attempts) {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given int32 and assigns it to the Attempts field.
func (o *ObjectRetry) SetAttempts(v int32) {
	o.Attempts = &v
}

func (o ObjectRetry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectRetry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attempts) {
		toSerialize["attempts"] = o.Attempts
	}
	return toSerialize, nil
}

type NullableObjectRetry struct {
	value *ObjectRetry
	isSet bool
}

func (v NullableObjectRetry) Get() *ObjectRetry {
	return v.value
}

func (v *NullableObjectRetry) Set(val *ObjectRetry) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRetry) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRetry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRetry(val *ObjectRetry) *NullableObjectRetry {
	return &NullableObjectRetry{value: val, isSet: true}
}

func (v NullableObjectRetry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRetry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
