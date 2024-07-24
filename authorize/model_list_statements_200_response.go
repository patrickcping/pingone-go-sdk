/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the ListStatements200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStatements200Response{}

// ListStatements200Response struct for ListStatements200Response
type ListStatements200Response struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Embedded *ListStatements200ResponseEmbedded `json:"_embedded,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewListStatements200Response instantiates a new ListStatements200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStatements200Response() *ListStatements200Response {
	this := ListStatements200Response{}
	return &this
}

// NewListStatements200ResponseWithDefaults instantiates a new ListStatements200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStatements200ResponseWithDefaults() *ListStatements200Response {
	this := ListStatements200Response{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListStatements200Response) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatements200Response) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListStatements200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ListStatements200Response) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ListStatements200Response) GetEmbedded() ListStatements200ResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ListStatements200ResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatements200Response) GetEmbeddedOk() (*ListStatements200ResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ListStatements200Response) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ListStatements200ResponseEmbedded and assigns it to the Embedded field.
func (o *ListStatements200Response) SetEmbedded(v ListStatements200ResponseEmbedded) {
	o.Embedded = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListStatements200Response) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatements200Response) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListStatements200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ListStatements200Response) SetCount(v int32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ListStatements200Response) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStatements200Response) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ListStatements200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ListStatements200Response) SetSize(v int32) {
	o.Size = &v
}

func (o ListStatements200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStatements200Response) ToMap() (map[string]interface{}, error) {
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

type NullableListStatements200Response struct {
	value *ListStatements200Response
	isSet bool
}

func (v NullableListStatements200Response) Get() *ListStatements200Response {
	return v.value
}

func (v *NullableListStatements200Response) Set(val *ListStatements200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListStatements200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListStatements200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStatements200Response(val *ListStatements200Response) *NullableListStatements200Response {
	return &NullableListStatements200Response{value: val, isSet: true}
}

func (v NullableListStatements200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStatements200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


