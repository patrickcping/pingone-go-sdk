/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner{}

// FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner struct for FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner
type FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner struct {
	// The name of the sub-attribute to use for the user display name.
	Name string `json:"name"`
}

type _FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner

// NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner(name string) *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner {
	this := FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner{}
	this.Name = name
	return &this
}

// NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInnerWithDefaults instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInnerWithDefaults() *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner {
	this := FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner{}
	return &this
}

// GetName returns the Name field value
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) SetName(v string) {
	o.Name = v
}

func (o FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner := _FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner)

	if err != nil {
		return err
	}

	*o = FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner(varFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner)

	return err
}

type NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner struct {
	value *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner
	isSet bool
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) Get() *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner {
	return v.value
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) Set(val *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner(val *FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner {
	return &NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner{value: val, isSet: true}
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
