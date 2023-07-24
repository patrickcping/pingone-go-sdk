/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the LinksHATEOAS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksHATEOAS{}

// LinksHATEOAS struct for LinksHATEOAS
type LinksHATEOAS struct {
	Self *LinksHATEOASSelf `json:"self,omitempty"`
	Next *LinksHATEOASNext `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksHATEOAS LinksHATEOAS

// NewLinksHATEOAS instantiates a new LinksHATEOAS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksHATEOAS() *LinksHATEOAS {
	this := LinksHATEOAS{}
	return &this
}

// NewLinksHATEOASWithDefaults instantiates a new LinksHATEOAS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksHATEOASWithDefaults() *LinksHATEOAS {
	this := LinksHATEOAS{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksHATEOAS) GetSelf() LinksHATEOASSelf {
	if o == nil || IsNil(o.Self) {
		var ret LinksHATEOASSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksHATEOAS) GetSelfOk() (*LinksHATEOASSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksHATEOAS) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given LinksHATEOASSelf and assigns it to the Self field.
func (o *LinksHATEOAS) SetSelf(v LinksHATEOASSelf) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *LinksHATEOAS) GetNext() LinksHATEOASNext {
	if o == nil || IsNil(o.Next) {
		var ret LinksHATEOASNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksHATEOAS) GetNextOk() (*LinksHATEOASNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *LinksHATEOAS) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given LinksHATEOASNext and assigns it to the Next field.
func (o *LinksHATEOAS) SetNext(v LinksHATEOASNext) {
	o.Next = &v
}

func (o LinksHATEOAS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksHATEOAS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksHATEOAS) UnmarshalJSON(bytes []byte) (err error) {
	varLinksHATEOAS := _LinksHATEOAS{}

	if err = json.Unmarshal(bytes, &varLinksHATEOAS); err == nil {
		*o = LinksHATEOAS(varLinksHATEOAS)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksHATEOAS struct {
	value *LinksHATEOAS
	isSet bool
}

func (v NullableLinksHATEOAS) Get() *LinksHATEOAS {
	return v.value
}

func (v *NullableLinksHATEOAS) Set(val *LinksHATEOAS) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksHATEOAS) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksHATEOAS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksHATEOAS(val *LinksHATEOAS) *NullableLinksHATEOAS {
	return &NullableLinksHATEOAS{value: val, isSet: true}
}

func (v NullableLinksHATEOAS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksHATEOAS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


