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

// checks if the CertificateKeyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateKeyUpdate{}

// CertificateKeyUpdate struct for CertificateKeyUpdate
type CertificateKeyUpdate struct {
	// Specifies whether this is the default key for the specified environment.
	Default bool `json:"default"`
	UsageType EnumCertificateKeyUsageType `json:"usageType"`
	// Specifies the distinguished name of the certificate issuer.
	IssuerDN *string `json:"issuerDN,omitempty"`
}

// NewCertificateKeyUpdate instantiates a new CertificateKeyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateKeyUpdate(default_ bool, usageType EnumCertificateKeyUsageType) *CertificateKeyUpdate {
	this := CertificateKeyUpdate{}
	this.Default = default_
	this.UsageType = usageType
	return &this
}

// NewCertificateKeyUpdateWithDefaults instantiates a new CertificateKeyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateKeyUpdateWithDefaults() *CertificateKeyUpdate {
	this := CertificateKeyUpdate{}
	return &this
}

// GetDefault returns the Default field value
func (o *CertificateKeyUpdate) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyUpdate) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *CertificateKeyUpdate) SetDefault(v bool) {
	o.Default = v
}

// GetUsageType returns the UsageType field value
func (o *CertificateKeyUpdate) GetUsageType() EnumCertificateKeyUsageType {
	if o == nil {
		var ret EnumCertificateKeyUsageType
		return ret
	}

	return o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyUpdate) GetUsageTypeOk() (*EnumCertificateKeyUsageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageType, true
}

// SetUsageType sets field value
func (o *CertificateKeyUpdate) SetUsageType(v EnumCertificateKeyUsageType) {
	o.UsageType = v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *CertificateKeyUpdate) GetIssuerDN() string {
	if o == nil || IsNil(o.IssuerDN) {
		var ret string
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateKeyUpdate) GetIssuerDNOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerDN) {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *CertificateKeyUpdate) HasIssuerDN() bool {
	if o != nil && !IsNil(o.IssuerDN) {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given string and assigns it to the IssuerDN field.
func (o *CertificateKeyUpdate) SetIssuerDN(v string) {
	o.IssuerDN = &v
}

func (o CertificateKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateKeyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	toSerialize["usageType"] = o.UsageType
	if !IsNil(o.IssuerDN) {
		toSerialize["issuerDN"] = o.IssuerDN
	}
	return toSerialize, nil
}

type NullableCertificateKeyUpdate struct {
	value *CertificateKeyUpdate
	isSet bool
}

func (v NullableCertificateKeyUpdate) Get() *CertificateKeyUpdate {
	return v.value
}

func (v *NullableCertificateKeyUpdate) Set(val *CertificateKeyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateKeyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateKeyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateKeyUpdate(val *CertificateKeyUpdate) *NullableCertificateKeyUpdate {
	return &NullableCertificateKeyUpdate{value: val, isSet: true}
}

func (v NullableCertificateKeyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateKeyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


