/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskPolicySetRiskPoliciesInner struct for RiskPolicySetRiskPoliciesInner
type RiskPolicySetRiskPoliciesInner struct {
	Condition RiskPolicySetRiskPoliciesInnerCondition `json:"condition"`
	// The time the resource was first created (format ISO-8061).
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies a description for this risk policy. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies a name for this risk policy. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters.
	Name string `json:"name"`
	// An integer that specifies priority of the policy inside a risk policy set, designating which policy should run first. This is a read-only value. The priority is determined by the order in which policies are listed in the policy set. The first policy in the list is assigned priority 1 and is evaluated first. The next policy in the list is assigned priority 2 and so on.
	Priority *int32 `json:"priority,omitempty"`
	Result RiskPolicyResult `json:"result"`
	// The time the resource was last updated (format ISO-8061).
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewRiskPolicySetRiskPoliciesInner instantiates a new RiskPolicySetRiskPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySetRiskPoliciesInner(condition RiskPolicySetRiskPoliciesInnerCondition, name string, result RiskPolicyResult) *RiskPolicySetRiskPoliciesInner {
	this := RiskPolicySetRiskPoliciesInner{}
	this.Condition = condition
	this.Name = name
	this.Result = result
	return &this
}

// NewRiskPolicySetRiskPoliciesInnerWithDefaults instantiates a new RiskPolicySetRiskPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetRiskPoliciesInnerWithDefaults() *RiskPolicySetRiskPoliciesInner {
	this := RiskPolicySetRiskPoliciesInner{}
	return &this
}

// GetCondition returns the Condition field value
func (o *RiskPolicySetRiskPoliciesInner) GetCondition() RiskPolicySetRiskPoliciesInnerCondition {
	if o == nil {
		var ret RiskPolicySetRiskPoliciesInnerCondition
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetConditionOk() (*RiskPolicySetRiskPoliciesInnerCondition, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *RiskPolicySetRiskPoliciesInner) SetCondition(v RiskPolicySetRiskPoliciesInnerCondition) {
	o.Condition = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskPolicySetRiskPoliciesInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPolicySetRiskPoliciesInner) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *RiskPolicySetRiskPoliciesInner) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPolicySetRiskPoliciesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPolicySetRiskPoliciesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPolicySetRiskPoliciesInner) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RiskPolicySetRiskPoliciesInner) SetPriority(v int32) {
	o.Priority = &v
}

// GetResult returns the Result field value
func (o *RiskPolicySetRiskPoliciesInner) GetResult() RiskPolicyResult {
	if o == nil {
		var ret RiskPolicyResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetResultOk() (*RiskPolicyResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *RiskPolicySetRiskPoliciesInner) SetResult(v RiskPolicyResult) {
	o.Result = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInner) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInner) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskPolicySetRiskPoliciesInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o RiskPolicySetRiskPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["result"] = o.Result
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPolicySetRiskPoliciesInner struct {
	value *RiskPolicySetRiskPoliciesInner
	isSet bool
}

func (v NullableRiskPolicySetRiskPoliciesInner) Get() *RiskPolicySetRiskPoliciesInner {
	return v.value
}

func (v *NullableRiskPolicySetRiskPoliciesInner) Set(val *RiskPolicySetRiskPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetRiskPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetRiskPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetRiskPoliciesInner(val *RiskPolicySetRiskPoliciesInner) *NullableRiskPolicySetRiskPoliciesInner {
	return &NullableRiskPolicySetRiskPoliciesInner{value: val, isSet: true}
}

func (v NullableRiskPolicySetRiskPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetRiskPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


