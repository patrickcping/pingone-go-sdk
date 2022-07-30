/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionAgreement struct for SignOnPolicyActionAgreement
type SignOnPolicyActionAgreement struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Conditions *SignOnPolicyActionCommonConditions `json:"conditions,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy,omitempty"`
	Type EnumSignOnPolicyType `json:"type"`
	Agreement SignOnPolicyActionAgreementAllOfAgreement `json:"agreement"`
}

// NewSignOnPolicyActionAgreement instantiates a new SignOnPolicyActionAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionAgreement(priority int32, type_ EnumSignOnPolicyType, agreement SignOnPolicyActionAgreementAllOfAgreement) *SignOnPolicyActionAgreement {
	this := SignOnPolicyActionAgreement{}
	this.Priority = priority
	this.Type = type_
	this.Agreement = agreement
	return &this
}

// NewSignOnPolicyActionAgreementWithDefaults instantiates a new SignOnPolicyActionAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionAgreementWithDefaults() *SignOnPolicyActionAgreement {
	this := SignOnPolicyActionAgreement{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreement) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreement) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *SignOnPolicyActionAgreement) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreement) GetConditions() SignOnPolicyActionCommonConditions {
	if o == nil || o.Conditions == nil {
		var ret SignOnPolicyActionCommonConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetConditionsOk() (*SignOnPolicyActionCommonConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreement) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SignOnPolicyActionCommonConditions and assigns it to the Conditions field.
func (o *SignOnPolicyActionAgreement) SetConditions(v SignOnPolicyActionCommonConditions) {
	o.Conditions = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreement) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreement) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionAgreement) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreement) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionAgreement) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionAgreement) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionAgreement) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionAgreement) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || o.SignOnPolicy == nil {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || o.SignOnPolicy == nil {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionAgreement) HasSignOnPolicy() bool {
	if o != nil && o.SignOnPolicy != nil {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionAgreement) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionAgreement) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionAgreement) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetAgreement returns the Agreement field value
func (o *SignOnPolicyActionAgreement) GetAgreement() SignOnPolicyActionAgreementAllOfAgreement {
	if o == nil {
		var ret SignOnPolicyActionAgreementAllOfAgreement
		return ret
	}

	return o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreement) GetAgreementOk() (*SignOnPolicyActionAgreementAllOfAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agreement, true
}

// SetAgreement sets field value
func (o *SignOnPolicyActionAgreement) SetAgreement(v SignOnPolicyActionAgreementAllOfAgreement) {
	o.Agreement = v
}

func (o SignOnPolicyActionAgreement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if o.SignOnPolicy != nil {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["agreement"] = o.Agreement
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionAgreement struct {
	value *SignOnPolicyActionAgreement
	isSet bool
}

func (v NullableSignOnPolicyActionAgreement) Get() *SignOnPolicyActionAgreement {
	return v.value
}

func (v *NullableSignOnPolicyActionAgreement) Set(val *SignOnPolicyActionAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionAgreement(val *SignOnPolicyActionAgreement) *NullableSignOnPolicyActionAgreement {
	return &NullableSignOnPolicyActionAgreement{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


