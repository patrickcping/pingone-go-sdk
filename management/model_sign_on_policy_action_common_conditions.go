/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionCommonConditions struct for SignOnPolicyActionCommonConditions
type SignOnPolicyActionCommonConditions struct {
	// A string that specifies the supported network IP addresses expressed as classless inter-domain routing (CIDR) strings.
	IpRange *string `json:"ipRange,omitempty"`
	// An integer that specifies the maximum number of minutes to wait since the last sign on before prompting for a new sign-on action.
	SecondsSince *int32 `json:"secondsSince,omitempty"`
}

// NewSignOnPolicyActionCommonConditions instantiates a new SignOnPolicyActionCommonConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonConditions() *SignOnPolicyActionCommonConditions {
	this := SignOnPolicyActionCommonConditions{}
	return &this
}

// NewSignOnPolicyActionCommonConditionsWithDefaults instantiates a new SignOnPolicyActionCommonConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonConditionsWithDefaults() *SignOnPolicyActionCommonConditions {
	this := SignOnPolicyActionCommonConditions{}
	return &this
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *SignOnPolicyActionCommonConditions) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditions) GetIpRangeOk() (*string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *SignOnPolicyActionCommonConditions) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *SignOnPolicyActionCommonConditions) SetIpRange(v string) {
	o.IpRange = &v
}

// GetSecondsSince returns the SecondsSince field value if set, zero value otherwise.
func (o *SignOnPolicyActionCommonConditions) GetSecondsSince() int32 {
	if o == nil || o.SecondsSince == nil {
		var ret int32
		return ret
	}
	return *o.SecondsSince
}

// GetSecondsSinceOk returns a tuple with the SecondsSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditions) GetSecondsSinceOk() (*int32, bool) {
	if o == nil || o.SecondsSince == nil {
		return nil, false
	}
	return o.SecondsSince, true
}

// HasSecondsSince returns a boolean if a field has been set.
func (o *SignOnPolicyActionCommonConditions) HasSecondsSince() bool {
	if o != nil && o.SecondsSince != nil {
		return true
	}

	return false
}

// SetSecondsSince gets a reference to the given int32 and assigns it to the SecondsSince field.
func (o *SignOnPolicyActionCommonConditions) SetSecondsSince(v int32) {
	o.SecondsSince = &v
}

func (o SignOnPolicyActionCommonConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpRange != nil {
		toSerialize["ipRange"] = o.IpRange
	}
	if o.SecondsSince != nil {
		toSerialize["secondsSince"] = o.SecondsSince
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionCommonConditions struct {
	value *SignOnPolicyActionCommonConditions
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditions) Get() *SignOnPolicyActionCommonConditions {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditions) Set(val *SignOnPolicyActionCommonConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditions(val *SignOnPolicyActionCommonConditions) *NullableSignOnPolicyActionCommonConditions {
	return &NullableSignOnPolicyActionCommonConditions{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


