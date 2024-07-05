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

// checks if the GovernmentIdConfigurationProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernmentIdConfigurationProvider{}

// GovernmentIdConfigurationProvider struct for GovernmentIdConfigurationProvider
type GovernmentIdConfigurationProvider struct {
	Auto *EnumProviderAuto `json:"auto,omitempty"`
	Manual *EnumProviderManual `json:"manual,omitempty"`
}

// NewGovernmentIdConfigurationProvider instantiates a new GovernmentIdConfigurationProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernmentIdConfigurationProvider() *GovernmentIdConfigurationProvider {
	this := GovernmentIdConfigurationProvider{}
	return &this
}

// NewGovernmentIdConfigurationProviderWithDefaults instantiates a new GovernmentIdConfigurationProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernmentIdConfigurationProviderWithDefaults() *GovernmentIdConfigurationProvider {
	this := GovernmentIdConfigurationProvider{}
	return &this
}

// GetAuto returns the Auto field value if set, zero value otherwise.
func (o *GovernmentIdConfigurationProvider) GetAuto() EnumProviderAuto {
	if o == nil || IsNil(o.Auto) {
		var ret EnumProviderAuto
		return ret
	}
	return *o.Auto
}

// GetAutoOk returns a tuple with the Auto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfigurationProvider) GetAutoOk() (*EnumProviderAuto, bool) {
	if o == nil || IsNil(o.Auto) {
		return nil, false
	}
	return o.Auto, true
}

// HasAuto returns a boolean if a field has been set.
func (o *GovernmentIdConfigurationProvider) HasAuto() bool {
	if o != nil && !IsNil(o.Auto) {
		return true
	}

	return false
}

// SetAuto gets a reference to the given EnumProviderAuto and assigns it to the Auto field.
func (o *GovernmentIdConfigurationProvider) SetAuto(v EnumProviderAuto) {
	o.Auto = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *GovernmentIdConfigurationProvider) GetManual() EnumProviderManual {
	if o == nil || IsNil(o.Manual) {
		var ret EnumProviderManual
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfigurationProvider) GetManualOk() (*EnumProviderManual, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *GovernmentIdConfigurationProvider) HasManual() bool {
	if o != nil && !IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given EnumProviderManual and assigns it to the Manual field.
func (o *GovernmentIdConfigurationProvider) SetManual(v EnumProviderManual) {
	o.Manual = &v
}

func (o GovernmentIdConfigurationProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernmentIdConfigurationProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Auto) {
		toSerialize["auto"] = o.Auto
	}
	if !IsNil(o.Manual) {
		toSerialize["manual"] = o.Manual
	}
	return toSerialize, nil
}

type NullableGovernmentIdConfigurationProvider struct {
	value *GovernmentIdConfigurationProvider
	isSet bool
}

func (v NullableGovernmentIdConfigurationProvider) Get() *GovernmentIdConfigurationProvider {
	return v.value
}

func (v *NullableGovernmentIdConfigurationProvider) Set(val *GovernmentIdConfigurationProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernmentIdConfigurationProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernmentIdConfigurationProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernmentIdConfigurationProvider(val *GovernmentIdConfigurationProvider) *NullableGovernmentIdConfigurationProvider {
	return &NullableGovernmentIdConfigurationProvider{value: val, isSet: true}
}

func (v NullableGovernmentIdConfigurationProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernmentIdConfigurationProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


