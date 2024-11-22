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

// checks if the MFAPushCredentialHMS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAPushCredentialHMS{}

// MFAPushCredentialHMS struct for MFAPushCredentialHMS
type MFAPushCredentialHMS struct {
	Type EnumMFAPushCredentialAttrType `json:"type"`
	// Used only if type is set to HMS. OAuth 2.0 Client ID from the Huawei Developers API console.
	ClientId string `json:"clientId"`
	// Used only if type is set to HMS. The client secret associated with the OAuth 2.0 Client ID.
	ClientSecret string `json:"clientSecret"`
}

type _MFAPushCredentialHMS MFAPushCredentialHMS

// NewMFAPushCredentialHMS instantiates a new MFAPushCredentialHMS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialHMS(type_ EnumMFAPushCredentialAttrType, clientId string, clientSecret string) *MFAPushCredentialHMS {
	this := MFAPushCredentialHMS{}
	this.Type = type_
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewMFAPushCredentialHMSWithDefaults instantiates a new MFAPushCredentialHMS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialHMSWithDefaults() *MFAPushCredentialHMS {
	this := MFAPushCredentialHMS{}
	return &this
}

// GetType returns the Type field value
func (o *MFAPushCredentialHMS) GetType() EnumMFAPushCredentialAttrType {
	if o == nil {
		var ret EnumMFAPushCredentialAttrType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialHMS) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MFAPushCredentialHMS) SetType(v EnumMFAPushCredentialAttrType) {
	o.Type = v
}

// GetClientId returns the ClientId field value
func (o *MFAPushCredentialHMS) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialHMS) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *MFAPushCredentialHMS) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *MFAPushCredentialHMS) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialHMS) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *MFAPushCredentialHMS) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o MFAPushCredentialHMS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAPushCredentialHMS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	return toSerialize, nil
}

func (o *MFAPushCredentialHMS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"clientId",
		"clientSecret",
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

	varMFAPushCredentialHMS := _MFAPushCredentialHMS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMFAPushCredentialHMS)

	if err != nil {
		return err
	}

	*o = MFAPushCredentialHMS(varMFAPushCredentialHMS)

	return err
}

type NullableMFAPushCredentialHMS struct {
	value *MFAPushCredentialHMS
	isSet bool
}

func (v NullableMFAPushCredentialHMS) Get() *MFAPushCredentialHMS {
	return v.value
}

func (v *NullableMFAPushCredentialHMS) Set(val *MFAPushCredentialHMS) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialHMS) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialHMS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialHMS(val *MFAPushCredentialHMS) *NullableMFAPushCredentialHMS {
	return &NullableMFAPushCredentialHMS{value: val, isSet: true}
}

func (v NullableMFAPushCredentialHMS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialHMS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
