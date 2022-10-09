/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"time"
)

// MFAPushCredentialResponse struct for MFAPushCredentialResponse
type MFAPushCredentialResponse struct {
	// A string that specifies the push credential ID.
	Id *string `json:"id,omitempty"`
	Type *EnumMFAPushCredentialAttrType `json:"type,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewMFAPushCredentialResponse instantiates a new MFAPushCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialResponse() *MFAPushCredentialResponse {
	this := MFAPushCredentialResponse{}
	return &this
}

// NewMFAPushCredentialResponseWithDefaults instantiates a new MFAPushCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialResponseWithDefaults() *MFAPushCredentialResponse {
	this := MFAPushCredentialResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MFAPushCredentialResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetType() EnumMFAPushCredentialAttrType {
	if o == nil || o.Type == nil {
		var ret EnumMFAPushCredentialAttrType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumMFAPushCredentialAttrType and assigns it to the Type field.
func (o *MFAPushCredentialResponse) SetType(v EnumMFAPushCredentialAttrType) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MFAPushCredentialResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MFAPushCredentialResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MFAPushCredentialResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MFAPushCredentialResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o MFAPushCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMFAPushCredentialResponse struct {
	value *MFAPushCredentialResponse
	isSet bool
}

func (v NullableMFAPushCredentialResponse) Get() *MFAPushCredentialResponse {
	return v.value
}

func (v *NullableMFAPushCredentialResponse) Set(val *MFAPushCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialResponse(val *MFAPushCredentialResponse) *NullableMFAPushCredentialResponse {
	return &NullableMFAPushCredentialResponse{value: val, isSet: true}
}

func (v NullableMFAPushCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


