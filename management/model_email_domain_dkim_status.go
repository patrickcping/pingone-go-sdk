/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// EmailDomainDKIMStatus struct for EmailDomainDKIMStatus
type EmailDomainDKIMStatus struct {
	// A string that specifies the type of DNS record, with the value \"CNAME\".
	Type *string `json:"type,omitempty"`
	// The regions collection specifies the properties for the 4 AWS SES regions that are used for sending email for the environment. The regions are determined by the geography where this environment was provisioned (North America, Canada, Europe & Asia-Pacific).
	Regions []EmailDomainDKIMStatusRegionsInner `json:"regions,omitempty"`
}

// NewEmailDomainDKIMStatus instantiates a new EmailDomainDKIMStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainDKIMStatus() *EmailDomainDKIMStatus {
	this := EmailDomainDKIMStatus{}
	return &this
}

// NewEmailDomainDKIMStatusWithDefaults instantiates a new EmailDomainDKIMStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainDKIMStatusWithDefaults() *EmailDomainDKIMStatus {
	this := EmailDomainDKIMStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EmailDomainDKIMStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDKIMStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EmailDomainDKIMStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EmailDomainDKIMStatus) SetType(v string) {
	o.Type = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *EmailDomainDKIMStatus) GetRegions() []EmailDomainDKIMStatusRegionsInner {
	if o == nil || o.Regions == nil {
		var ret []EmailDomainDKIMStatusRegionsInner
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDKIMStatus) GetRegionsOk() ([]EmailDomainDKIMStatusRegionsInner, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *EmailDomainDKIMStatus) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []EmailDomainDKIMStatusRegionsInner and assigns it to the Regions field.
func (o *EmailDomainDKIMStatus) SetRegions(v []EmailDomainDKIMStatusRegionsInner) {
	o.Regions = v
}

func (o EmailDomainDKIMStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDomainDKIMStatus struct {
	value *EmailDomainDKIMStatus
	isSet bool
}

func (v NullableEmailDomainDKIMStatus) Get() *EmailDomainDKIMStatus {
	return v.value
}

func (v *NullableEmailDomainDKIMStatus) Set(val *EmailDomainDKIMStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainDKIMStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainDKIMStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainDKIMStatus(val *EmailDomainDKIMStatus) *NullableEmailDomainDKIMStatus {
	return &NullableEmailDomainDKIMStatus{value: val, isSet: true}
}

func (v NullableEmailDomainDKIMStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainDKIMStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


