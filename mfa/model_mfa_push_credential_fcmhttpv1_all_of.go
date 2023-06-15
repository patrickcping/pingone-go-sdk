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

// checks if the MFAPushCredentialFCMHTTPV1AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAPushCredentialFCMHTTPV1AllOf{}

// MFAPushCredentialFCMHTTPV1AllOf struct for MFAPushCredentialFCMHTTPV1AllOf
type MFAPushCredentialFCMHTTPV1AllOf struct {
	// Used when `type` is set to `FCM_HTTP_V1`. The value should be the contents of the JSON file that represents your Service Account Credentials.
	GoogleServiceAccountCredentials string `json:"googleServiceAccountCredentials"`
}

// NewMFAPushCredentialFCMHTTPV1AllOf instantiates a new MFAPushCredentialFCMHTTPV1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialFCMHTTPV1AllOf(googleServiceAccountCredentials string) *MFAPushCredentialFCMHTTPV1AllOf {
	this := MFAPushCredentialFCMHTTPV1AllOf{}
	this.GoogleServiceAccountCredentials = googleServiceAccountCredentials
	return &this
}

// NewMFAPushCredentialFCMHTTPV1AllOfWithDefaults instantiates a new MFAPushCredentialFCMHTTPV1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialFCMHTTPV1AllOfWithDefaults() *MFAPushCredentialFCMHTTPV1AllOf {
	this := MFAPushCredentialFCMHTTPV1AllOf{}
	return &this
}

// GetGoogleServiceAccountCredentials returns the GoogleServiceAccountCredentials field value
func (o *MFAPushCredentialFCMHTTPV1AllOf) GetGoogleServiceAccountCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleServiceAccountCredentials
}

// GetGoogleServiceAccountCredentialsOk returns a tuple with the GoogleServiceAccountCredentials field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialFCMHTTPV1AllOf) GetGoogleServiceAccountCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleServiceAccountCredentials, true
}

// SetGoogleServiceAccountCredentials sets field value
func (o *MFAPushCredentialFCMHTTPV1AllOf) SetGoogleServiceAccountCredentials(v string) {
	o.GoogleServiceAccountCredentials = v
}

func (o MFAPushCredentialFCMHTTPV1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAPushCredentialFCMHTTPV1AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["googleServiceAccountCredentials"] = o.GoogleServiceAccountCredentials
	return toSerialize, nil
}

type NullableMFAPushCredentialFCMHTTPV1AllOf struct {
	value *MFAPushCredentialFCMHTTPV1AllOf
	isSet bool
}

func (v NullableMFAPushCredentialFCMHTTPV1AllOf) Get() *MFAPushCredentialFCMHTTPV1AllOf {
	return v.value
}

func (v *NullableMFAPushCredentialFCMHTTPV1AllOf) Set(val *MFAPushCredentialFCMHTTPV1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialFCMHTTPV1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialFCMHTTPV1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialFCMHTTPV1AllOf(val *MFAPushCredentialFCMHTTPV1AllOf) *NullableMFAPushCredentialFCMHTTPV1AllOf {
	return &NullableMFAPushCredentialFCMHTTPV1AllOf{value: val, isSet: true}
}

func (v NullableMFAPushCredentialFCMHTTPV1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialFCMHTTPV1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


