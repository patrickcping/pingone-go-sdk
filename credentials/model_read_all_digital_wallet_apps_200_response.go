/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ReadAllDigitalWalletApps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllDigitalWalletApps200Response{}

// ReadAllDigitalWalletApps200Response struct for ReadAllDigitalWalletApps200Response
type ReadAllDigitalWalletApps200Response struct {
	Links    map[string]LinksHATEOASValue                 `json:"_links,omitempty"`
	Embedded *ReadAllDigitalWalletApps200ResponseEmbedded `json:"_embedded,omitempty"`
	Count    *float32                                     `json:"count,omitempty"`
	Size     *float32                                     `json:"size,omitempty"`
}

// NewReadAllDigitalWalletApps200Response instantiates a new ReadAllDigitalWalletApps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllDigitalWalletApps200Response() *ReadAllDigitalWalletApps200Response {
	this := ReadAllDigitalWalletApps200Response{}
	return &this
}

// NewReadAllDigitalWalletApps200ResponseWithDefaults instantiates a new ReadAllDigitalWalletApps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllDigitalWalletApps200ResponseWithDefaults() *ReadAllDigitalWalletApps200Response {
	this := ReadAllDigitalWalletApps200Response{}
	return &this
}

func (o ReadAllDigitalWalletApps200Response) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o ReadAllDigitalWalletApps200Response) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o ReadAllDigitalWalletApps200Response) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o ReadAllDigitalWalletApps200Response) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllDigitalWalletApps200Response) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadAllDigitalWalletApps200Response) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadAllDigitalWalletApps200Response) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadAllDigitalWalletApps200Response) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReadAllDigitalWalletApps200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllDigitalWalletApps200Response) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReadAllDigitalWalletApps200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ReadAllDigitalWalletApps200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ReadAllDigitalWalletApps200Response) GetEmbedded() ReadAllDigitalWalletApps200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ReadAllDigitalWalletApps200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllDigitalWalletApps200Response) GetEmbeddedOk() (*ReadAllDigitalWalletApps200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ReadAllDigitalWalletApps200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ReadAllDigitalWalletApps200ResponseEmbedded and assigns it to the Embedded field.
func (o *ReadAllDigitalWalletApps200Response) SetEmbedded(v ReadAllDigitalWalletApps200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReadAllDigitalWalletApps200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllDigitalWalletApps200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReadAllDigitalWalletApps200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ReadAllDigitalWalletApps200Response) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReadAllDigitalWalletApps200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllDigitalWalletApps200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReadAllDigitalWalletApps200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReadAllDigitalWalletApps200Response) SetSize(v float32) {
	o.Size = &v
}

func (o ReadAllDigitalWalletApps200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllDigitalWalletApps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableReadAllDigitalWalletApps200Response struct {
	value *ReadAllDigitalWalletApps200Response
	isSet bool
}

func (v NullableReadAllDigitalWalletApps200Response) Get() *ReadAllDigitalWalletApps200Response {
	return v.value
}

func (v *NullableReadAllDigitalWalletApps200Response) Set(val *ReadAllDigitalWalletApps200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllDigitalWalletApps200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllDigitalWalletApps200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllDigitalWalletApps200Response(val *ReadAllDigitalWalletApps200Response) *NullableReadAllDigitalWalletApps200Response {
	return &NullableReadAllDigitalWalletApps200Response{value: val, isSet: true}
}

func (v NullableReadAllDigitalWalletApps200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllDigitalWalletApps200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
