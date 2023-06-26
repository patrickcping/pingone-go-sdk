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

// checks if the NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf{}

// NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf struct for NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf
type NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf struct {
	// The public ID of the Twilio account. Relevant to Twilio only. 
	Sid string `json:"sid"`
	// The secret key of the Twilio or Syniverse account.
	AuthToken string `json:"authToken"`
	Numbers []NotificationsSettingsPhoneDeliverySettingsCustomNumbers `json:"numbers,omitempty"`
}

// NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf(sid string, authToken string) *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf {
	this := NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf{}
	this.Sid = sid
	this.AuthToken = authToken
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfWithDefaults() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf {
	this := NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf{}
	return &this
}

// GetSid returns the Sid field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sid
}

// GetSidOk returns a tuple with the Sid field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sid, true
}

// SetSid sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) SetSid(v string) {
	o.Sid = v
}

// GetAuthToken returns the AuthToken field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) SetAuthToken(v string) {
	o.AuthToken = v
}

// GetNumbers returns the Numbers field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsCustomNumbers {
	if o == nil || IsNil(o.Numbers) {
		var ret []NotificationsSettingsPhoneDeliverySettingsCustomNumbers
		return ret
	}
	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) GetNumbersOk() ([]NotificationsSettingsPhoneDeliverySettingsCustomNumbers, bool) {
	if o == nil || IsNil(o.Numbers) {
		return nil, false
	}
	return o.Numbers, true
}

// HasNumbers returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) HasNumbers() bool {
	if o != nil && !IsNil(o.Numbers) {
		return true
	}

	return false
}

// SetNumbers gets a reference to the given []NotificationsSettingsPhoneDeliverySettingsCustomNumbers and assigns it to the Numbers field.
func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsCustomNumbers) {
	o.Numbers = v
}

func (o NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sid"] = o.Sid
	toSerialize["authToken"] = o.AuthToken
	if !IsNil(o.Numbers) {
		toSerialize["numbers"] = o.Numbers
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf struct {
	value *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) Get() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) Set(val *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf(val *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf {
	return &NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


