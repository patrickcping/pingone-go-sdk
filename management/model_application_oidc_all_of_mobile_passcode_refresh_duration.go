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

// checks if the ApplicationOIDCAllOfMobilePasscodeRefreshDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOIDCAllOfMobilePasscodeRefreshDuration{}

// ApplicationOIDCAllOfMobilePasscodeRefreshDuration struct for ApplicationOIDCAllOfMobilePasscodeRefreshDuration
type ApplicationOIDCAllOfMobilePasscodeRefreshDuration struct {
	// The amount of time a passcode should be displayed before being replaced with a new passcode - must be between 30 and 60.
	Duration int32 `json:"duration"`
	TimeUnit EnumPasscodeRefreshTimeUnit `json:"timeUnit"`
}

// NewApplicationOIDCAllOfMobilePasscodeRefreshDuration instantiates a new ApplicationOIDCAllOfMobilePasscodeRefreshDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOfMobilePasscodeRefreshDuration(duration int32, timeUnit EnumPasscodeRefreshTimeUnit) *ApplicationOIDCAllOfMobilePasscodeRefreshDuration {
	this := ApplicationOIDCAllOfMobilePasscodeRefreshDuration{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewApplicationOIDCAllOfMobilePasscodeRefreshDurationWithDefaults instantiates a new ApplicationOIDCAllOfMobilePasscodeRefreshDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfMobilePasscodeRefreshDurationWithDefaults() *ApplicationOIDCAllOfMobilePasscodeRefreshDuration {
	this := ApplicationOIDCAllOfMobilePasscodeRefreshDuration{}
	var duration int32 = 30
	this.Duration = duration
	var timeUnit EnumPasscodeRefreshTimeUnit = ENUMPASSCODEREFRESHTIMEUNIT_SECONDS
	this.TimeUnit = timeUnit
	return &this
}

// GetDuration returns the Duration field value
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetTimeUnit() EnumPasscodeRefreshTimeUnit {
	if o == nil {
		var ret EnumPasscodeRefreshTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) GetTimeUnitOk() (*EnumPasscodeRefreshTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) SetTimeUnit(v EnumPasscodeRefreshTimeUnit) {
	o.TimeUnit = v
}

func (o ApplicationOIDCAllOfMobilePasscodeRefreshDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOIDCAllOfMobilePasscodeRefreshDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

type NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration struct {
	value *ApplicationOIDCAllOfMobilePasscodeRefreshDuration
	isSet bool
}

func (v NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) Get() *ApplicationOIDCAllOfMobilePasscodeRefreshDuration {
	return v.value
}

func (v *NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) Set(val *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOfMobilePasscodeRefreshDuration(val *ApplicationOIDCAllOfMobilePasscodeRefreshDuration) *NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration {
	return &NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOfMobilePasscodeRefreshDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


