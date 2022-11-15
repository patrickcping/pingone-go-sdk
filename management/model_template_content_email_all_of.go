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

// TemplateContentEmailAllOf struct for TemplateContentEmailAllOf
type TemplateContentEmailAllOf struct {
	From *TemplateContentEmailAllOfFrom `json:"from,omitempty"`
	// The email's subject line. Cannot exceed 256 characters. If supported, can include variables.
	Subject *string `json:"subject,omitempty"`
	ReplyTo *TemplateContentEmailAllOfReplyTo `json:"replyTo,omitempty"`
	// If not specified, `UTF-8` is the default value.
	Charset *string `json:"charset,omitempty"`
	// If not specified, `text/html` is the default value.
	EmailContentType *string `json:"emailContentType,omitempty"`
}

// NewTemplateContentEmailAllOf instantiates a new TemplateContentEmailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentEmailAllOf() *TemplateContentEmailAllOf {
	this := TemplateContentEmailAllOf{}
	var charset string = "UTF-8"
	this.Charset = &charset
	var emailContentType string = "text/html"
	this.EmailContentType = &emailContentType
	return &this
}

// NewTemplateContentEmailAllOfWithDefaults instantiates a new TemplateContentEmailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentEmailAllOfWithDefaults() *TemplateContentEmailAllOf {
	this := TemplateContentEmailAllOf{}
	var charset string = "UTF-8"
	this.Charset = &charset
	var emailContentType string = "text/html"
	this.EmailContentType = &emailContentType
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOf) GetFrom() TemplateContentEmailAllOfFrom {
	if o == nil || isNil(o.From) {
		var ret TemplateContentEmailAllOfFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOf) GetFromOk() (*TemplateContentEmailAllOfFrom, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOf) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given TemplateContentEmailAllOfFrom and assigns it to the From field.
func (o *TemplateContentEmailAllOf) SetFrom(v TemplateContentEmailAllOfFrom) {
	o.From = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOf) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOf) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOf) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TemplateContentEmailAllOf) SetSubject(v string) {
	o.Subject = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOf) GetReplyTo() TemplateContentEmailAllOfReplyTo {
	if o == nil || isNil(o.ReplyTo) {
		var ret TemplateContentEmailAllOfReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOf) GetReplyToOk() (*TemplateContentEmailAllOfReplyTo, bool) {
	if o == nil || isNil(o.ReplyTo) {
    return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOf) HasReplyTo() bool {
	if o != nil && !isNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given TemplateContentEmailAllOfReplyTo and assigns it to the ReplyTo field.
func (o *TemplateContentEmailAllOf) SetReplyTo(v TemplateContentEmailAllOfReplyTo) {
	o.ReplyTo = &v
}

// GetCharset returns the Charset field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOf) GetCharset() string {
	if o == nil || isNil(o.Charset) {
		var ret string
		return ret
	}
	return *o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOf) GetCharsetOk() (*string, bool) {
	if o == nil || isNil(o.Charset) {
    return nil, false
	}
	return o.Charset, true
}

// HasCharset returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOf) HasCharset() bool {
	if o != nil && !isNil(o.Charset) {
		return true
	}

	return false
}

// SetCharset gets a reference to the given string and assigns it to the Charset field.
func (o *TemplateContentEmailAllOf) SetCharset(v string) {
	o.Charset = &v
}

// GetEmailContentType returns the EmailContentType field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOf) GetEmailContentType() string {
	if o == nil || isNil(o.EmailContentType) {
		var ret string
		return ret
	}
	return *o.EmailContentType
}

// GetEmailContentTypeOk returns a tuple with the EmailContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOf) GetEmailContentTypeOk() (*string, bool) {
	if o == nil || isNil(o.EmailContentType) {
    return nil, false
	}
	return o.EmailContentType, true
}

// HasEmailContentType returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOf) HasEmailContentType() bool {
	if o != nil && !isNil(o.EmailContentType) {
		return true
	}

	return false
}

// SetEmailContentType gets a reference to the given string and assigns it to the EmailContentType field.
func (o *TemplateContentEmailAllOf) SetEmailContentType(v string) {
	o.EmailContentType = &v
}

func (o TemplateContentEmailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !isNil(o.Charset) {
		toSerialize["charset"] = o.Charset
	}
	if !isNil(o.EmailContentType) {
		toSerialize["emailContentType"] = o.EmailContentType
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateContentEmailAllOf struct {
	value *TemplateContentEmailAllOf
	isSet bool
}

func (v NullableTemplateContentEmailAllOf) Get() *TemplateContentEmailAllOf {
	return v.value
}

func (v *NullableTemplateContentEmailAllOf) Set(val *TemplateContentEmailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentEmailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentEmailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentEmailAllOf(val *TemplateContentEmailAllOf) *NullableTemplateContentEmailAllOf {
	return &NullableTemplateContentEmailAllOf{value: val, isSet: true}
}

func (v NullableTemplateContentEmailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentEmailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


