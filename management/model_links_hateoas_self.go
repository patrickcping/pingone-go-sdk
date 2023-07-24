/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the LinksHATEOASSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksHATEOASSelf{}

// LinksHATEOASSelf An object that describes the current resource.
type LinksHATEOASSelf struct {
	// The URI of the resource.
	Href *string `json:"href,omitempty"`
}

// NewLinksHATEOASSelf instantiates a new LinksHATEOASSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksHATEOASSelf() *LinksHATEOASSelf {
	this := LinksHATEOASSelf{}
	return &this
}

// NewLinksHATEOASSelfWithDefaults instantiates a new LinksHATEOASSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksHATEOASSelfWithDefaults() *LinksHATEOASSelf {
	this := LinksHATEOASSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinksHATEOASSelf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksHATEOASSelf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinksHATEOASSelf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinksHATEOASSelf) SetHref(v string) {
	o.Href = &v
}

func (o LinksHATEOASSelf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksHATEOASSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableLinksHATEOASSelf struct {
	value *LinksHATEOASSelf
	isSet bool
}

func (v NullableLinksHATEOASSelf) Get() *LinksHATEOASSelf {
	return v.value
}

func (v *NullableLinksHATEOASSelf) Set(val *LinksHATEOASSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksHATEOASSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksHATEOASSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksHATEOASSelf(val *LinksHATEOASSelf) *NullableLinksHATEOASSelf {
	return &NullableLinksHATEOASSelf{value: val, isSet: true}
}

func (v NullableLinksHATEOASSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksHATEOASSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


