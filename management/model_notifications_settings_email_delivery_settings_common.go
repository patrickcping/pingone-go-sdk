/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the NotificationsSettingsEmailDeliverySettingsCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsEmailDeliverySettingsCommon{}

// NotificationsSettingsEmailDeliverySettingsCommon struct for NotificationsSettingsEmailDeliverySettingsCommon
type NotificationsSettingsEmailDeliverySettingsCommon struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewNotificationsSettingsEmailDeliverySettingsCommon instantiates a new NotificationsSettingsEmailDeliverySettingsCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsEmailDeliverySettingsCommon() *NotificationsSettingsEmailDeliverySettingsCommon {
	this := NotificationsSettingsEmailDeliverySettingsCommon{}
	return &this
}

// NewNotificationsSettingsEmailDeliverySettingsCommonWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsEmailDeliverySettingsCommonWithDefaults() *NotificationsSettingsEmailDeliverySettingsCommon {
	this := NotificationsSettingsEmailDeliverySettingsCommon{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NotificationsSettingsEmailDeliverySettingsCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsEmailDeliverySettingsCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsEmailDeliverySettingsCommon struct {
	value *NotificationsSettingsEmailDeliverySettingsCommon
	isSet bool
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCommon) Get() *NotificationsSettingsEmailDeliverySettingsCommon {
	return v.value
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCommon) Set(val *NotificationsSettingsEmailDeliverySettingsCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsEmailDeliverySettingsCommon(val *NotificationsSettingsEmailDeliverySettingsCommon) *NullableNotificationsSettingsEmailDeliverySettingsCommon {
	return &NullableNotificationsSettingsEmailDeliverySettingsCommon{value: val, isSet: true}
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


