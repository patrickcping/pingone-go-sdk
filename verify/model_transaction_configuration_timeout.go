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

// checks if the TransactionConfigurationTimeout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionConfigurationTimeout{}

// TransactionConfigurationTimeout struct for TransactionConfigurationTimeout
type TransactionConfigurationTimeout struct {
	// Length of time before transaction timeout expires; range is 0-30 minutes or 0-1800 seconds.
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnit `json:"timeUnit"`
}

// NewTransactionConfigurationTimeout instantiates a new TransactionConfigurationTimeout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionConfigurationTimeout(duration int32, timeUnit EnumTimeUnit) *TransactionConfigurationTimeout {
	this := TransactionConfigurationTimeout{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewTransactionConfigurationTimeoutWithDefaults instantiates a new TransactionConfigurationTimeout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionConfigurationTimeoutWithDefaults() *TransactionConfigurationTimeout {
	this := TransactionConfigurationTimeout{}
	return &this
}

// GetDuration returns the Duration field value
func (o *TransactionConfigurationTimeout) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *TransactionConfigurationTimeout) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *TransactionConfigurationTimeout) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *TransactionConfigurationTimeout) GetTimeUnit() EnumTimeUnit {
	if o == nil {
		var ret EnumTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *TransactionConfigurationTimeout) GetTimeUnitOk() (*EnumTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *TransactionConfigurationTimeout) SetTimeUnit(v EnumTimeUnit) {
	o.TimeUnit = v
}

func (o TransactionConfigurationTimeout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionConfigurationTimeout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

type NullableTransactionConfigurationTimeout struct {
	value *TransactionConfigurationTimeout
	isSet bool
}

func (v NullableTransactionConfigurationTimeout) Get() *TransactionConfigurationTimeout {
	return v.value
}

func (v *NullableTransactionConfigurationTimeout) Set(val *TransactionConfigurationTimeout) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionConfigurationTimeout) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionConfigurationTimeout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionConfigurationTimeout(val *TransactionConfigurationTimeout) *NullableTransactionConfigurationTimeout {
	return &NullableTransactionConfigurationTimeout{value: val, isSet: true}
}

func (v NullableTransactionConfigurationTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionConfigurationTimeout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


