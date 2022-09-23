/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2022-09-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// APIServerOperations A map from the operation name to the operation object. Each key must be valid ObjectName, and each value must be a valid operation. Each key must be unique within the operations object, which means the operation key is unique within an API server. No duplicate operation values are allowed; operations with the same paths and methods members are not allowed. The operations object is limited to 25 keys (25 individual operations).
type APIServerOperations struct {
	// A string that specifies the name of the operation.
	Key *string `json:"key,omitempty"`
	Value *APIServerOperationsValue `json:"value,omitempty"`
}

// NewAPIServerOperations instantiates a new APIServerOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerOperations() *APIServerOperations {
	this := APIServerOperations{}
	return &this
}

// NewAPIServerOperationsWithDefaults instantiates a new APIServerOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerOperationsWithDefaults() *APIServerOperations {
	this := APIServerOperations{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *APIServerOperations) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerOperations) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *APIServerOperations) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *APIServerOperations) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *APIServerOperations) GetValue() APIServerOperationsValue {
	if o == nil || o.Value == nil {
		var ret APIServerOperationsValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerOperations) GetValueOk() (*APIServerOperationsValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *APIServerOperations) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given APIServerOperationsValue and assigns it to the Value field.
func (o *APIServerOperations) SetValue(v APIServerOperationsValue) {
	o.Value = &v
}

func (o APIServerOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAPIServerOperations struct {
	value *APIServerOperations
	isSet bool
}

func (v NullableAPIServerOperations) Get() *APIServerOperations {
	return v.value
}

func (v *NullableAPIServerOperations) Set(val *APIServerOperations) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerOperations(val *APIServerOperations) *NullableAPIServerOperations {
	return &NullableAPIServerOperations{value: val, isSet: true}
}

func (v NullableAPIServerOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


