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

// checks if the CredentialIssuanceRuleNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleNotification{}

// CredentialIssuanceRuleNotification struct for CredentialIssuanceRuleNotification
type CredentialIssuanceRuleNotification struct {
	Methods  []EnumCredentialIssuanceRuleNotificationMethod `json:"methods,omitempty"`
	Template *CredentialIssuanceRuleNotificationTemplate    `json:"template,omitempty"`
}

// NewCredentialIssuanceRuleNotification instantiates a new CredentialIssuanceRuleNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleNotification() *CredentialIssuanceRuleNotification {
	this := CredentialIssuanceRuleNotification{}
	return &this
}

// NewCredentialIssuanceRuleNotificationWithDefaults instantiates a new CredentialIssuanceRuleNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleNotificationWithDefaults() *CredentialIssuanceRuleNotification {
	this := CredentialIssuanceRuleNotification{}
	return &this
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleNotification) GetMethods() []EnumCredentialIssuanceRuleNotificationMethod {
	if o == nil || IsNil(o.Methods) {
		var ret []EnumCredentialIssuanceRuleNotificationMethod
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleNotification) GetMethodsOk() ([]EnumCredentialIssuanceRuleNotificationMethod, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleNotification) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []EnumCredentialIssuanceRuleNotificationMethod and assigns it to the Methods field.
func (o *CredentialIssuanceRuleNotification) SetMethods(v []EnumCredentialIssuanceRuleNotificationMethod) {
	o.Methods = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleNotification) GetTemplate() CredentialIssuanceRuleNotificationTemplate {
	if o == nil || IsNil(o.Template) {
		var ret CredentialIssuanceRuleNotificationTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleNotification) GetTemplateOk() (*CredentialIssuanceRuleNotificationTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleNotification) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given CredentialIssuanceRuleNotificationTemplate and assigns it to the Template field.
func (o *CredentialIssuanceRuleNotification) SetTemplate(v CredentialIssuanceRuleNotificationTemplate) {
	o.Template = &v
}

func (o CredentialIssuanceRuleNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleNotification struct {
	value *CredentialIssuanceRuleNotification
	isSet bool
}

func (v NullableCredentialIssuanceRuleNotification) Get() *CredentialIssuanceRuleNotification {
	return v.value
}

func (v *NullableCredentialIssuanceRuleNotification) Set(val *CredentialIssuanceRuleNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleNotification(val *CredentialIssuanceRuleNotification) *NullableCredentialIssuanceRuleNotification {
	return &NullableCredentialIssuanceRuleNotification{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
