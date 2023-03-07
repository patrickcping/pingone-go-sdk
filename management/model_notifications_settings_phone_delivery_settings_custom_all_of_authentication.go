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

// checks if the NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication{}

// NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication struct for NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication
type NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication struct {
	Method EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod `json:"method"`
	// The username for the custom provider account. Required when `authentication.method=BASIC`
	Username *string `json:"username,omitempty"`
	// The password for the custom provider account. Required when `authentication.method=BASIC`
	Password *string `json:"password,omitempty"`
	// The authentication token for the custom provider account.  Required when `authentication.method=BEARER`
	Token *string `json:"token,omitempty"`
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication(method EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod) *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication {
	this := NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication{}
	this.Method = method
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthenticationWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthenticationWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication {
	this := NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication{}
	return &this
}

// GetMethod returns the Method field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod) {
	o.Method = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetPassword(v string) {
	o.Password = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetToken(v string) {
	o.Token = &v
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication struct {
	value *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) Get() *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) Set(val *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication(val *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication {
	return &NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


