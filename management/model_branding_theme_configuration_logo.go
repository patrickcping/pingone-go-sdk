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

// BrandingThemeConfigurationLogo struct for BrandingThemeConfigurationLogo
type BrandingThemeConfigurationLogo struct {
	// The URL or fully qualified path to the logo file used for branding. This is a required property when configuration.logoType is set to IMAGE.
	Href string `json:"href"`
	// The ID of the logo image.
	Id string `json:"id"`
}

// NewBrandingThemeConfigurationLogo instantiates a new BrandingThemeConfigurationLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingThemeConfigurationLogo(href string, id string) *BrandingThemeConfigurationLogo {
	this := BrandingThemeConfigurationLogo{}
	this.Href = href
	this.Id = id
	return &this
}

// NewBrandingThemeConfigurationLogoWithDefaults instantiates a new BrandingThemeConfigurationLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingThemeConfigurationLogoWithDefaults() *BrandingThemeConfigurationLogo {
	this := BrandingThemeConfigurationLogo{}
	return &this
}

// GetHref returns the Href field value
func (o *BrandingThemeConfigurationLogo) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *BrandingThemeConfigurationLogo) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *BrandingThemeConfigurationLogo) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *BrandingThemeConfigurationLogo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BrandingThemeConfigurationLogo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BrandingThemeConfigurationLogo) SetId(v string) {
	o.Id = v
}

func (o BrandingThemeConfigurationLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBrandingThemeConfigurationLogo struct {
	value *BrandingThemeConfigurationLogo
	isSet bool
}

func (v NullableBrandingThemeConfigurationLogo) Get() *BrandingThemeConfigurationLogo {
	return v.value
}

func (v *NullableBrandingThemeConfigurationLogo) Set(val *BrandingThemeConfigurationLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingThemeConfigurationLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingThemeConfigurationLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingThemeConfigurationLogo(val *BrandingThemeConfigurationLogo) *NullableBrandingThemeConfigurationLogo {
	return &NullableBrandingThemeConfigurationLogo{value: val, isSet: true}
}

func (v NullableBrandingThemeConfigurationLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingThemeConfigurationLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


