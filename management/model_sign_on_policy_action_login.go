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

// checks if the SignOnPolicyActionLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionLogin{}

// SignOnPolicyActionLogin struct for SignOnPolicyActionLogin
type SignOnPolicyActionLogin struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	Condition *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy,omitempty"`
	Type EnumSignOnPolicyType `json:"type"`
	// A boolean that if set to true and if the user's account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented.
	EnforceLockoutForIdentityProviders *bool `json:"enforceLockoutForIdentityProviders,omitempty"`
	NewUserProvisioning *SignOnPolicyActionLoginAllOfNewUserProvisioning `json:"newUserProvisioning,omitempty"`
	Recovery *SignOnPolicyActionLoginAllOfRecovery `json:"recovery,omitempty"`
	Registration *SignOnPolicyActionLoginAllOfRegistration `json:"registration,omitempty"`
	// An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow.
	SocialProviders []SignOnPolicyActionLoginAllOfSocialProviders `json:"socialProviders,omitempty"`
}

// NewSignOnPolicyActionLogin instantiates a new SignOnPolicyActionLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLogin(priority int32, type_ EnumSignOnPolicyType) *SignOnPolicyActionLogin {
	this := SignOnPolicyActionLogin{}
	this.Priority = priority
	this.Type = type_
	return &this
}

// NewSignOnPolicyActionLoginWithDefaults instantiates a new SignOnPolicyActionLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginWithDefaults() *SignOnPolicyActionLogin {
	this := SignOnPolicyActionLogin{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *SignOnPolicyActionLogin) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || IsNil(o.Condition) {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionLogin) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionLogin) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionLogin) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionLogin) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionLogin) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || IsNil(o.SignOnPolicy) {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || IsNil(o.SignOnPolicy) {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasSignOnPolicy() bool {
	if o != nil && !IsNil(o.SignOnPolicy) {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionLogin) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionLogin) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionLogin) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetEnforceLockoutForIdentityProviders() bool {
	if o == nil || IsNil(o.EnforceLockoutForIdentityProviders) {
		var ret bool
		return ret
	}
	return *o.EnforceLockoutForIdentityProviders
}

// GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceLockoutForIdentityProviders) {
		return nil, false
	}
	return o.EnforceLockoutForIdentityProviders, true
}

// HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasEnforceLockoutForIdentityProviders() bool {
	if o != nil && !IsNil(o.EnforceLockoutForIdentityProviders) {
		return true
	}

	return false
}

// SetEnforceLockoutForIdentityProviders gets a reference to the given bool and assigns it to the EnforceLockoutForIdentityProviders field.
func (o *SignOnPolicyActionLogin) SetEnforceLockoutForIdentityProviders(v bool) {
	o.EnforceLockoutForIdentityProviders = &v
}

// GetNewUserProvisioning returns the NewUserProvisioning field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetNewUserProvisioning() SignOnPolicyActionLoginAllOfNewUserProvisioning {
	if o == nil || IsNil(o.NewUserProvisioning) {
		var ret SignOnPolicyActionLoginAllOfNewUserProvisioning
		return ret
	}
	return *o.NewUserProvisioning
}

// GetNewUserProvisioningOk returns a tuple with the NewUserProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetNewUserProvisioningOk() (*SignOnPolicyActionLoginAllOfNewUserProvisioning, bool) {
	if o == nil || IsNil(o.NewUserProvisioning) {
		return nil, false
	}
	return o.NewUserProvisioning, true
}

// HasNewUserProvisioning returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasNewUserProvisioning() bool {
	if o != nil && !IsNil(o.NewUserProvisioning) {
		return true
	}

	return false
}

// SetNewUserProvisioning gets a reference to the given SignOnPolicyActionLoginAllOfNewUserProvisioning and assigns it to the NewUserProvisioning field.
func (o *SignOnPolicyActionLogin) SetNewUserProvisioning(v SignOnPolicyActionLoginAllOfNewUserProvisioning) {
	o.NewUserProvisioning = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetRecovery() SignOnPolicyActionLoginAllOfRecovery {
	if o == nil || IsNil(o.Recovery) {
		var ret SignOnPolicyActionLoginAllOfRecovery
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool) {
	if o == nil || IsNil(o.Recovery) {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasRecovery() bool {
	if o != nil && !IsNil(o.Recovery) {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given SignOnPolicyActionLoginAllOfRecovery and assigns it to the Recovery field.
func (o *SignOnPolicyActionLogin) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetRegistration() SignOnPolicyActionLoginAllOfRegistration {
	if o == nil || IsNil(o.Registration) {
		var ret SignOnPolicyActionLoginAllOfRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given SignOnPolicyActionLoginAllOfRegistration and assigns it to the Registration field.
func (o *SignOnPolicyActionLogin) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration) {
	o.Registration = &v
}

// GetSocialProviders returns the SocialProviders field value if set, zero value otherwise.
func (o *SignOnPolicyActionLogin) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders {
	if o == nil || IsNil(o.SocialProviders) {
		var ret []SignOnPolicyActionLoginAllOfSocialProviders
		return ret
	}
	return o.SocialProviders
}

// GetSocialProvidersOk returns a tuple with the SocialProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLogin) GetSocialProvidersOk() ([]SignOnPolicyActionLoginAllOfSocialProviders, bool) {
	if o == nil || IsNil(o.SocialProviders) {
		return nil, false
	}
	return o.SocialProviders, true
}

// HasSocialProviders returns a boolean if a field has been set.
func (o *SignOnPolicyActionLogin) HasSocialProviders() bool {
	if o != nil && !IsNil(o.SocialProviders) {
		return true
	}

	return false
}

// SetSocialProviders gets a reference to the given []SignOnPolicyActionLoginAllOfSocialProviders and assigns it to the SocialProviders field.
func (o *SignOnPolicyActionLogin) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders) {
	o.SocialProviders = v
}

func (o SignOnPolicyActionLogin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["priority"] = o.Priority
	if !IsNil(o.SignOnPolicy) {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.EnforceLockoutForIdentityProviders) {
		toSerialize["enforceLockoutForIdentityProviders"] = o.EnforceLockoutForIdentityProviders
	}
	if !IsNil(o.NewUserProvisioning) {
		toSerialize["newUserProvisioning"] = o.NewUserProvisioning
	}
	if !IsNil(o.Recovery) {
		toSerialize["recovery"] = o.Recovery
	}
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if !IsNil(o.SocialProviders) {
		toSerialize["socialProviders"] = o.SocialProviders
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionLogin struct {
	value *SignOnPolicyActionLogin
	isSet bool
}

func (v NullableSignOnPolicyActionLogin) Get() *SignOnPolicyActionLogin {
	return v.value
}

func (v *NullableSignOnPolicyActionLogin) Set(val *SignOnPolicyActionLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLogin(val *SignOnPolicyActionLogin) *NullableSignOnPolicyActionLogin {
	return &NullableSignOnPolicyActionLogin{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


