/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// MFAPushCredentialHMSAllOf struct for MFAPushCredentialHMSAllOf
type MFAPushCredentialHMSAllOf struct {
	// Used only if type is set to HMS. OAuth 2.0 Client ID from the Huawei Developers API console.
	ClientId string `json:"clientId"`
	// Used only if type is set to HMS. The client secret associated with the OAuth 2.0 Client ID.
	ClientSecret string `json:"clientSecret"`
}

// NewMFAPushCredentialHMSAllOf instantiates a new MFAPushCredentialHMSAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialHMSAllOf(clientId string, clientSecret string) *MFAPushCredentialHMSAllOf {
	this := MFAPushCredentialHMSAllOf{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewMFAPushCredentialHMSAllOfWithDefaults instantiates a new MFAPushCredentialHMSAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialHMSAllOfWithDefaults() *MFAPushCredentialHMSAllOf {
	this := MFAPushCredentialHMSAllOf{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *MFAPushCredentialHMSAllOf) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialHMSAllOf) GetClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *MFAPushCredentialHMSAllOf) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *MFAPushCredentialHMSAllOf) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialHMSAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *MFAPushCredentialHMSAllOf) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o MFAPushCredentialHMSAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableMFAPushCredentialHMSAllOf struct {
	value *MFAPushCredentialHMSAllOf
	isSet bool
}

func (v NullableMFAPushCredentialHMSAllOf) Get() *MFAPushCredentialHMSAllOf {
	return v.value
}

func (v *NullableMFAPushCredentialHMSAllOf) Set(val *MFAPushCredentialHMSAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialHMSAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialHMSAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialHMSAllOf(val *MFAPushCredentialHMSAllOf) *NullableMFAPushCredentialHMSAllOf {
	return &NullableMFAPushCredentialHMSAllOf{value: val, isSet: true}
}

func (v NullableMFAPushCredentialHMSAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialHMSAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


