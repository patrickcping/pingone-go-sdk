/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the ReadFIDO2Policies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadFIDO2Policies200Response{}

// ReadFIDO2Policies200Response struct for ReadFIDO2Policies200Response
type ReadFIDO2Policies200Response struct {
	Links    map[string]LinksHATEOASValue          `json:"_links,omitempty"`
	Embedded *ReadFIDO2Policies200ResponseEmbedded `json:"_embedded,omitempty"`
	Count    *float32                              `json:"count,omitempty"`
	Size     *float32                              `json:"size,omitempty"`
}

// NewReadFIDO2Policies200Response instantiates a new ReadFIDO2Policies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFIDO2Policies200Response() *ReadFIDO2Policies200Response {
	this := ReadFIDO2Policies200Response{}
	return &this
}

// NewReadFIDO2Policies200ResponseWithDefaults instantiates a new ReadFIDO2Policies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFIDO2Policies200ResponseWithDefaults() *ReadFIDO2Policies200Response {
	this := ReadFIDO2Policies200Response{}
	return &this
}

func (o ReadFIDO2Policies200Response) hasHalLink(linkIndex string) bool {
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

func (o ReadFIDO2Policies200Response) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o ReadFIDO2Policies200Response) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o ReadFIDO2Policies200Response) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadFIDO2Policies200Response) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadFIDO2Policies200Response) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadFIDO2Policies200Response) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o ReadFIDO2Policies200Response) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadFIDO2Policies200Response) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadFIDO2Policies200Response) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o ReadFIDO2Policies200Response) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadFIDO2Policies200Response) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o ReadFIDO2Policies200Response) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReadFIDO2Policies200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFIDO2Policies200Response) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReadFIDO2Policies200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ReadFIDO2Policies200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ReadFIDO2Policies200Response) GetEmbedded() ReadFIDO2Policies200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ReadFIDO2Policies200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFIDO2Policies200Response) GetEmbeddedOk() (*ReadFIDO2Policies200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ReadFIDO2Policies200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ReadFIDO2Policies200ResponseEmbedded and assigns it to the Embedded field.
func (o *ReadFIDO2Policies200Response) SetEmbedded(v ReadFIDO2Policies200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReadFIDO2Policies200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFIDO2Policies200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReadFIDO2Policies200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ReadFIDO2Policies200Response) SetCount(v float32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReadFIDO2Policies200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFIDO2Policies200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReadFIDO2Policies200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReadFIDO2Policies200Response) SetSize(v float32) {
	o.Size = &v
}

func (o ReadFIDO2Policies200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadFIDO2Policies200Response) ToMap() (map[string]interface{}, error) {
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

type NullableReadFIDO2Policies200Response struct {
	value *ReadFIDO2Policies200Response
	isSet bool
}

func (v NullableReadFIDO2Policies200Response) Get() *ReadFIDO2Policies200Response {
	return v.value
}

func (v *NullableReadFIDO2Policies200Response) Set(val *ReadFIDO2Policies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFIDO2Policies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFIDO2Policies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFIDO2Policies200Response(val *ReadFIDO2Policies200Response) *NullableReadFIDO2Policies200Response {
	return &NullableReadFIDO2Policies200Response{value: val, isSet: true}
}

func (v NullableReadFIDO2Policies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFIDO2Policies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
