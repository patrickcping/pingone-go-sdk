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

// checks if the IdentityProviderSAMLAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderSAMLAllOf{}

// IdentityProviderSAMLAllOf struct for IdentityProviderSAMLAllOf
type IdentityProviderSAMLAllOf struct {
	// A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs.
	AuthnRequestSigned *bool `json:"authnRequestSigned,omitempty"`
	// A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response.
	IdpEntityId string `json:"idpEntityId"`
	IdpVerification IdentityProviderSAMLAllOfIdpVerification `json:"idpVerification"`
	// A string that specifies the service provider's entity ID, used to look up the application.
	SpEntityId string `json:"spEntityId"`
	SpSigning *IdentityProviderSAMLAllOfSpSigning `json:"spSigning,omitempty"`
	SsoBinding EnumIdentityProviderSAMLSSOBinding `json:"ssoBinding"`
	// A string that specifies the SSO endpoint for the authentication request.
	SsoEndpoint string `json:"ssoEndpoint"`
}

// NewIdentityProviderSAMLAllOf instantiates a new IdentityProviderSAMLAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderSAMLAllOf(idpEntityId string, idpVerification IdentityProviderSAMLAllOfIdpVerification, spEntityId string, ssoBinding EnumIdentityProviderSAMLSSOBinding, ssoEndpoint string) *IdentityProviderSAMLAllOf {
	this := IdentityProviderSAMLAllOf{}
	this.IdpEntityId = idpEntityId
	this.IdpVerification = idpVerification
	this.SpEntityId = spEntityId
	this.SsoBinding = ssoBinding
	this.SsoEndpoint = ssoEndpoint
	return &this
}

// NewIdentityProviderSAMLAllOfWithDefaults instantiates a new IdentityProviderSAMLAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderSAMLAllOfWithDefaults() *IdentityProviderSAMLAllOf {
	this := IdentityProviderSAMLAllOf{}
	return &this
}

// GetAuthnRequestSigned returns the AuthnRequestSigned field value if set, zero value otherwise.
func (o *IdentityProviderSAMLAllOf) GetAuthnRequestSigned() bool {
	if o == nil || IsNil(o.AuthnRequestSigned) {
		var ret bool
		return ret
	}
	return *o.AuthnRequestSigned
}

// GetAuthnRequestSignedOk returns a tuple with the AuthnRequestSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetAuthnRequestSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthnRequestSigned) {
		return nil, false
	}
	return o.AuthnRequestSigned, true
}

// HasAuthnRequestSigned returns a boolean if a field has been set.
func (o *IdentityProviderSAMLAllOf) HasAuthnRequestSigned() bool {
	if o != nil && !IsNil(o.AuthnRequestSigned) {
		return true
	}

	return false
}

// SetAuthnRequestSigned gets a reference to the given bool and assigns it to the AuthnRequestSigned field.
func (o *IdentityProviderSAMLAllOf) SetAuthnRequestSigned(v bool) {
	o.AuthnRequestSigned = &v
}

// GetIdpEntityId returns the IdpEntityId field value
func (o *IdentityProviderSAMLAllOf) GetIdpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpEntityId
}

// GetIdpEntityIdOk returns a tuple with the IdpEntityId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetIdpEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpEntityId, true
}

// SetIdpEntityId sets field value
func (o *IdentityProviderSAMLAllOf) SetIdpEntityId(v string) {
	o.IdpEntityId = v
}

// GetIdpVerification returns the IdpVerification field value
func (o *IdentityProviderSAMLAllOf) GetIdpVerification() IdentityProviderSAMLAllOfIdpVerification {
	if o == nil {
		var ret IdentityProviderSAMLAllOfIdpVerification
		return ret
	}

	return o.IdpVerification
}

// GetIdpVerificationOk returns a tuple with the IdpVerification field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetIdpVerificationOk() (*IdentityProviderSAMLAllOfIdpVerification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpVerification, true
}

// SetIdpVerification sets field value
func (o *IdentityProviderSAMLAllOf) SetIdpVerification(v IdentityProviderSAMLAllOfIdpVerification) {
	o.IdpVerification = v
}

// GetSpEntityId returns the SpEntityId field value
func (o *IdentityProviderSAMLAllOf) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetSpEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *IdentityProviderSAMLAllOf) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetSpSigning returns the SpSigning field value if set, zero value otherwise.
func (o *IdentityProviderSAMLAllOf) GetSpSigning() IdentityProviderSAMLAllOfSpSigning {
	if o == nil || IsNil(o.SpSigning) {
		var ret IdentityProviderSAMLAllOfSpSigning
		return ret
	}
	return *o.SpSigning
}

// GetSpSigningOk returns a tuple with the SpSigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetSpSigningOk() (*IdentityProviderSAMLAllOfSpSigning, bool) {
	if o == nil || IsNil(o.SpSigning) {
		return nil, false
	}
	return o.SpSigning, true
}

// HasSpSigning returns a boolean if a field has been set.
func (o *IdentityProviderSAMLAllOf) HasSpSigning() bool {
	if o != nil && !IsNil(o.SpSigning) {
		return true
	}

	return false
}

// SetSpSigning gets a reference to the given IdentityProviderSAMLAllOfSpSigning and assigns it to the SpSigning field.
func (o *IdentityProviderSAMLAllOf) SetSpSigning(v IdentityProviderSAMLAllOfSpSigning) {
	o.SpSigning = &v
}

// GetSsoBinding returns the SsoBinding field value
func (o *IdentityProviderSAMLAllOf) GetSsoBinding() EnumIdentityProviderSAMLSSOBinding {
	if o == nil {
		var ret EnumIdentityProviderSAMLSSOBinding
		return ret
	}

	return o.SsoBinding
}

// GetSsoBindingOk returns a tuple with the SsoBinding field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetSsoBindingOk() (*EnumIdentityProviderSAMLSSOBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoBinding, true
}

// SetSsoBinding sets field value
func (o *IdentityProviderSAMLAllOf) SetSsoBinding(v EnumIdentityProviderSAMLSSOBinding) {
	o.SsoBinding = v
}

// GetSsoEndpoint returns the SsoEndpoint field value
func (o *IdentityProviderSAMLAllOf) GetSsoEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoEndpoint
}

// GetSsoEndpointOk returns a tuple with the SsoEndpoint field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderSAMLAllOf) GetSsoEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoEndpoint, true
}

// SetSsoEndpoint sets field value
func (o *IdentityProviderSAMLAllOf) SetSsoEndpoint(v string) {
	o.SsoEndpoint = v
}

func (o IdentityProviderSAMLAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderSAMLAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthnRequestSigned) {
		toSerialize["authnRequestSigned"] = o.AuthnRequestSigned
	}
	toSerialize["idpEntityId"] = o.IdpEntityId
	toSerialize["idpVerification"] = o.IdpVerification
	toSerialize["spEntityId"] = o.SpEntityId
	if !IsNil(o.SpSigning) {
		toSerialize["spSigning"] = o.SpSigning
	}
	toSerialize["ssoBinding"] = o.SsoBinding
	toSerialize["ssoEndpoint"] = o.SsoEndpoint
	return toSerialize, nil
}

type NullableIdentityProviderSAMLAllOf struct {
	value *IdentityProviderSAMLAllOf
	isSet bool
}

func (v NullableIdentityProviderSAMLAllOf) Get() *IdentityProviderSAMLAllOf {
	return v.value
}

func (v *NullableIdentityProviderSAMLAllOf) Set(val *IdentityProviderSAMLAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderSAMLAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderSAMLAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderSAMLAllOf(val *IdentityProviderSAMLAllOf) *NullableIdentityProviderSAMLAllOf {
	return &NullableIdentityProviderSAMLAllOf{value: val, isSet: true}
}

func (v NullableIdentityProviderSAMLAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderSAMLAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


