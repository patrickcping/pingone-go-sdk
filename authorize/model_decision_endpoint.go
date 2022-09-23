/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2022-09-23
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// DecisionEndpoint struct for DecisionEndpoint
type DecisionEndpoint struct {
	// A string that specifies alternative unique identifier for the endpoint, which provides a method for locating the resource by a known, fixed identifier.
	AlternateId *string `json:"alternateId,omitempty"`
	AuthorizationVersion *DecisionEndpointAuthorizationVersion `json:"authorizationVersion,omitempty"`
	// A string that specifies the description of the policy decision resource.
	Description string `json:"description"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the policy decision resource name.
	Name string `json:"name"`
	// A boolean that when true restricts modifications of the endpoint to PingOne-owned clients.
	Owned *bool `json:"owned,omitempty"`
	// A string that specifies the ID of the root policy configured for this endpoint. If omitted, the policy editor service decides on a suitable default.
	PolicyId *string `json:"policyId,omitempty"`
	// A string that specifies a machine-readable identifier indicating the provenance of the current configuration. It has no meaning to the Policy Decision Service itself but exists to support integration with other services.
	Provenance *string `json:"provenance,omitempty"`
	// A boolean that specifies whether to show recent decisions.
	RecentDecisionsEnabled *bool `json:"recentDecisionsEnabled,omitempty"`
	RecentDecisions *DecisionEndpointRecentDecisions `json:"recentDecisions,omitempty"`
	// A boolean that specifies whether to record a limited history of recent decision requests and responses, which can be queried through a separate API.
	RecordRecentRequests bool `json:"recordRecentRequests"`
}

// NewDecisionEndpoint instantiates a new DecisionEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionEndpoint(description string, name string, recordRecentRequests bool) *DecisionEndpoint {
	this := DecisionEndpoint{}
	this.Description = description
	this.Name = name
	this.RecordRecentRequests = recordRecentRequests
	return &this
}

// NewDecisionEndpointWithDefaults instantiates a new DecisionEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionEndpointWithDefaults() *DecisionEndpoint {
	this := DecisionEndpoint{}
	return &this
}

// GetAlternateId returns the AlternateId field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetAlternateId() string {
	if o == nil || o.AlternateId == nil {
		var ret string
		return ret
	}
	return *o.AlternateId
}

// GetAlternateIdOk returns a tuple with the AlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetAlternateIdOk() (*string, bool) {
	if o == nil || o.AlternateId == nil {
		return nil, false
	}
	return o.AlternateId, true
}

// HasAlternateId returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasAlternateId() bool {
	if o != nil && o.AlternateId != nil {
		return true
	}

	return false
}

// SetAlternateId gets a reference to the given string and assigns it to the AlternateId field.
func (o *DecisionEndpoint) SetAlternateId(v string) {
	o.AlternateId = &v
}

// GetAuthorizationVersion returns the AuthorizationVersion field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetAuthorizationVersion() DecisionEndpointAuthorizationVersion {
	if o == nil || o.AuthorizationVersion == nil {
		var ret DecisionEndpointAuthorizationVersion
		return ret
	}
	return *o.AuthorizationVersion
}

// GetAuthorizationVersionOk returns a tuple with the AuthorizationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetAuthorizationVersionOk() (*DecisionEndpointAuthorizationVersion, bool) {
	if o == nil || o.AuthorizationVersion == nil {
		return nil, false
	}
	return o.AuthorizationVersion, true
}

// HasAuthorizationVersion returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasAuthorizationVersion() bool {
	if o != nil && o.AuthorizationVersion != nil {
		return true
	}

	return false
}

// SetAuthorizationVersion gets a reference to the given DecisionEndpointAuthorizationVersion and assigns it to the AuthorizationVersion field.
func (o *DecisionEndpoint) SetAuthorizationVersion(v DecisionEndpointAuthorizationVersion) {
	o.AuthorizationVersion = &v
}

// GetDescription returns the Description field value
func (o *DecisionEndpoint) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DecisionEndpoint) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DecisionEndpoint) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DecisionEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DecisionEndpoint) SetName(v string) {
	o.Name = v
}

// GetOwned returns the Owned field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetOwned() bool {
	if o == nil || o.Owned == nil {
		var ret bool
		return ret
	}
	return *o.Owned
}

// GetOwnedOk returns a tuple with the Owned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetOwnedOk() (*bool, bool) {
	if o == nil || o.Owned == nil {
		return nil, false
	}
	return o.Owned, true
}

// HasOwned returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasOwned() bool {
	if o != nil && o.Owned != nil {
		return true
	}

	return false
}

// SetOwned gets a reference to the given bool and assigns it to the Owned field.
func (o *DecisionEndpoint) SetOwned(v bool) {
	o.Owned = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *DecisionEndpoint) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetProvenance returns the Provenance field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetProvenance() string {
	if o == nil || o.Provenance == nil {
		var ret string
		return ret
	}
	return *o.Provenance
}

// GetProvenanceOk returns a tuple with the Provenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetProvenanceOk() (*string, bool) {
	if o == nil || o.Provenance == nil {
		return nil, false
	}
	return o.Provenance, true
}

// HasProvenance returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasProvenance() bool {
	if o != nil && o.Provenance != nil {
		return true
	}

	return false
}

// SetProvenance gets a reference to the given string and assigns it to the Provenance field.
func (o *DecisionEndpoint) SetProvenance(v string) {
	o.Provenance = &v
}

// GetRecentDecisionsEnabled returns the RecentDecisionsEnabled field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetRecentDecisionsEnabled() bool {
	if o == nil || o.RecentDecisionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RecentDecisionsEnabled
}

// GetRecentDecisionsEnabledOk returns a tuple with the RecentDecisionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetRecentDecisionsEnabledOk() (*bool, bool) {
	if o == nil || o.RecentDecisionsEnabled == nil {
		return nil, false
	}
	return o.RecentDecisionsEnabled, true
}

// HasRecentDecisionsEnabled returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasRecentDecisionsEnabled() bool {
	if o != nil && o.RecentDecisionsEnabled != nil {
		return true
	}

	return false
}

// SetRecentDecisionsEnabled gets a reference to the given bool and assigns it to the RecentDecisionsEnabled field.
func (o *DecisionEndpoint) SetRecentDecisionsEnabled(v bool) {
	o.RecentDecisionsEnabled = &v
}

// GetRecentDecisions returns the RecentDecisions field value if set, zero value otherwise.
func (o *DecisionEndpoint) GetRecentDecisions() DecisionEndpointRecentDecisions {
	if o == nil || o.RecentDecisions == nil {
		var ret DecisionEndpointRecentDecisions
		return ret
	}
	return *o.RecentDecisions
}

// GetRecentDecisionsOk returns a tuple with the RecentDecisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetRecentDecisionsOk() (*DecisionEndpointRecentDecisions, bool) {
	if o == nil || o.RecentDecisions == nil {
		return nil, false
	}
	return o.RecentDecisions, true
}

// HasRecentDecisions returns a boolean if a field has been set.
func (o *DecisionEndpoint) HasRecentDecisions() bool {
	if o != nil && o.RecentDecisions != nil {
		return true
	}

	return false
}

// SetRecentDecisions gets a reference to the given DecisionEndpointRecentDecisions and assigns it to the RecentDecisions field.
func (o *DecisionEndpoint) SetRecentDecisions(v DecisionEndpointRecentDecisions) {
	o.RecentDecisions = &v
}

// GetRecordRecentRequests returns the RecordRecentRequests field value
func (o *DecisionEndpoint) GetRecordRecentRequests() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RecordRecentRequests
}

// GetRecordRecentRequestsOk returns a tuple with the RecordRecentRequests field value
// and a boolean to check if the value has been set.
func (o *DecisionEndpoint) GetRecordRecentRequestsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordRecentRequests, true
}

// SetRecordRecentRequests sets field value
func (o *DecisionEndpoint) SetRecordRecentRequests(v bool) {
	o.RecordRecentRequests = v
}

func (o DecisionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlternateId != nil {
		toSerialize["alternateId"] = o.AlternateId
	}
	if o.AuthorizationVersion != nil {
		toSerialize["authorizationVersion"] = o.AuthorizationVersion
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owned != nil {
		toSerialize["owned"] = o.Owned
	}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.Provenance != nil {
		toSerialize["provenance"] = o.Provenance
	}
	if o.RecentDecisionsEnabled != nil {
		toSerialize["recentDecisionsEnabled"] = o.RecentDecisionsEnabled
	}
	if o.RecentDecisions != nil {
		toSerialize["recentDecisions"] = o.RecentDecisions
	}
	if true {
		toSerialize["recordRecentRequests"] = o.RecordRecentRequests
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionEndpoint struct {
	value *DecisionEndpoint
	isSet bool
}

func (v NullableDecisionEndpoint) Get() *DecisionEndpoint {
	return v.value
}

func (v *NullableDecisionEndpoint) Set(val *DecisionEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionEndpoint(val *DecisionEndpoint) *NullableDecisionEndpoint {
	return &NullableDecisionEndpoint{value: val, isSet: true}
}

func (v NullableDecisionEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


