/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider struct for SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider
type SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider struct {
	// A string that specifies the identity provider that will be used to authenticate the user if the condition is matched.
	Id string `json:"id"`
}

// NewSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider instantiates a new SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider(id string) *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider {
	this := SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProviderWithDefaults instantiates a new SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProviderWithDefaults() *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider {
	this := SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider struct {
	value *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider
	isSet bool
}

func (v NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) Get() *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider {
	return v.value
}

func (v *NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) Set(val *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider(val *SignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) *NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider {
	return &NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDFirstDiscoveryRulesInnerIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


