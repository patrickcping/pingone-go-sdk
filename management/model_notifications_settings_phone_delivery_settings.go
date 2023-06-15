/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// NotificationsSettingsPhoneDeliverySettings - struct for NotificationsSettingsPhoneDeliverySettings
type NotificationsSettingsPhoneDeliverySettings struct {
	NotificationsSettingsPhoneDeliverySettingsCustom *NotificationsSettingsPhoneDeliverySettingsCustom
	NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse
}

// NotificationsSettingsPhoneDeliverySettingsCustomAsNotificationsSettingsPhoneDeliverySettings is a convenience function that returns NotificationsSettingsPhoneDeliverySettingsCustom wrapped in NotificationsSettingsPhoneDeliverySettings
func NotificationsSettingsPhoneDeliverySettingsCustomAsNotificationsSettingsPhoneDeliverySettings(v *NotificationsSettingsPhoneDeliverySettingsCustom) NotificationsSettingsPhoneDeliverySettings {
	return NotificationsSettingsPhoneDeliverySettings{
		NotificationsSettingsPhoneDeliverySettingsCustom: v,
	}
}

// NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAsNotificationsSettingsPhoneDeliverySettings is a convenience function that returns NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse wrapped in NotificationsSettingsPhoneDeliverySettings
func NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAsNotificationsSettingsPhoneDeliverySettings(v *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) NotificationsSettingsPhoneDeliverySettings {
	return NotificationsSettingsPhoneDeliverySettings{
		NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationsSettingsPhoneDeliverySettings) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotificationsSettingsPhoneDeliverySettingsCustom
	err = newStrictDecoder(data).Decode(&dst.NotificationsSettingsPhoneDeliverySettingsCustom)
	if err == nil {
		jsonNotificationsSettingsPhoneDeliverySettingsCustom, _ := json.Marshal(dst.NotificationsSettingsPhoneDeliverySettingsCustom)
		if string(jsonNotificationsSettingsPhoneDeliverySettingsCustom) == "{}" { // empty struct
			dst.NotificationsSettingsPhoneDeliverySettingsCustom = nil
		} else {
			match++
		}
	} else {
		dst.NotificationsSettingsPhoneDeliverySettingsCustom = nil
	}

	// try to unmarshal data into NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse
	err = newStrictDecoder(data).Decode(&dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse)
	if err == nil {
		jsonNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse, _ := json.Marshal(dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse)
		if string(jsonNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) == "{}" { // empty struct
			dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse = nil
		} else {
			match++
		}
	} else {
		dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NotificationsSettingsPhoneDeliverySettingsCustom = nil
		dst.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotificationsSettingsPhoneDeliverySettings)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotificationsSettingsPhoneDeliverySettings)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationsSettingsPhoneDeliverySettings) MarshalJSON() ([]byte, error) {
	if src.NotificationsSettingsPhoneDeliverySettingsCustom != nil {
		return json.Marshal(&src.NotificationsSettingsPhoneDeliverySettingsCustom)
	}

	if src.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse != nil {
		return json.Marshal(&src.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationsSettingsPhoneDeliverySettings) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NotificationsSettingsPhoneDeliverySettingsCustom != nil {
		return obj.NotificationsSettingsPhoneDeliverySettingsCustom
	}

	if obj.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse != nil {
		return obj.NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse
	}

	// all schemas are nil
	return nil
}

type NullableNotificationsSettingsPhoneDeliverySettings struct {
	value *NotificationsSettingsPhoneDeliverySettings
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) Get() *NotificationsSettingsPhoneDeliverySettings {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) Set(val *NotificationsSettingsPhoneDeliverySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettings(val *NotificationsSettingsPhoneDeliverySettings) *NullableNotificationsSettingsPhoneDeliverySettings {
	return &NullableNotificationsSettingsPhoneDeliverySettings{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


