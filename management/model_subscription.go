/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// Subscription struct for Subscription
type Subscription struct {
	// The time the key resource expires.The date and time at which the subscription resource was created (ISO 8601 format).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A boolean that specifies whether a created or updated subscription should be active or suspended. A suspended state (`\"enabled\":false`) accumulates all matched events, but these events are not delivered until the subscription becomes active again (`\"enabled\":true`). For suspended subscriptions, events accumulate for a maximum of two weeks. Events older than two weeks are deleted. Restarted subscriptions receive the saved events (up to two weeks from the restart date). This is a required property.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	FilterOptions SubscriptionFilterOptions `json:"filterOptions"`
	Format EnumSubscriptionFormat `json:"format"`
	// A string that specifies the user resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	HttpEndpoint SubscriptionHttpEndpoint `json:"httpEndpoint"`
	// A string that specifies the subscription name. This is a required property.
	Name string `json:"name"`
	// The date and time at which the subscription resource was last updated (ISO 8601 format).
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A boolean that specifies whether a certificates should be verified. If this property's value is set to false, then all certificates are trusted. (Setting this property's value to false introduces a security risk.) This is a required property.
	VerifyTlsCertificates bool `json:"verifyTlsCertificates"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(enabled bool, filterOptions SubscriptionFilterOptions, format EnumSubscriptionFormat, httpEndpoint SubscriptionHttpEndpoint, name string, verifyTlsCertificates bool) *Subscription {
	this := Subscription{}
	this.Enabled = enabled
	this.FilterOptions = filterOptions
	this.Format = format
	this.HttpEndpoint = httpEndpoint
	this.Name = name
	this.VerifyTlsCertificates = verifyTlsCertificates
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subscription) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subscription) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Subscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnabled returns the Enabled field value
func (o *Subscription) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Subscription) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Subscription) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Subscription) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Subscription) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetFilterOptions returns the FilterOptions field value
func (o *Subscription) GetFilterOptions() SubscriptionFilterOptions {
	if o == nil {
		var ret SubscriptionFilterOptions
		return ret
	}

	return o.FilterOptions
}

// GetFilterOptionsOk returns a tuple with the FilterOptions field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetFilterOptionsOk() (*SubscriptionFilterOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterOptions, true
}

// SetFilterOptions sets field value
func (o *Subscription) SetFilterOptions(v SubscriptionFilterOptions) {
	o.FilterOptions = v
}

// GetFormat returns the Format field value
func (o *Subscription) GetFormat() EnumSubscriptionFormat {
	if o == nil {
		var ret EnumSubscriptionFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetFormatOk() (*EnumSubscriptionFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *Subscription) SetFormat(v EnumSubscriptionFormat) {
	o.Format = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subscription) SetId(v string) {
	o.Id = &v
}

// GetHttpEndpoint returns the HttpEndpoint field value
func (o *Subscription) GetHttpEndpoint() SubscriptionHttpEndpoint {
	if o == nil {
		var ret SubscriptionHttpEndpoint
		return ret
	}

	return o.HttpEndpoint
}

// GetHttpEndpointOk returns a tuple with the HttpEndpoint field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetHttpEndpointOk() (*SubscriptionHttpEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpEndpoint, true
}

// SetHttpEndpoint sets field value
func (o *Subscription) SetHttpEndpoint(v SubscriptionHttpEndpoint) {
	o.HttpEndpoint = v
}

// GetName returns the Name field value
func (o *Subscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subscription) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Subscription) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Subscription) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Subscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerifyTlsCertificates returns the VerifyTlsCertificates field value
func (o *Subscription) GetVerifyTlsCertificates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerifyTlsCertificates
}

// GetVerifyTlsCertificatesOk returns a tuple with the VerifyTlsCertificates field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetVerifyTlsCertificatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyTlsCertificates, true
}

// SetVerifyTlsCertificates sets field value
func (o *Subscription) SetVerifyTlsCertificates(v bool) {
	o.VerifyTlsCertificates = v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["filterOptions"] = o.FilterOptions
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["httpEndpoint"] = o.HttpEndpoint
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["verifyTlsCertificates"] = o.VerifyTlsCertificates
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


