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

// checks if the NotificationsSettingsEmailDeliverySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsEmailDeliverySettings{}

// NotificationsSettingsEmailDeliverySettings struct for NotificationsSettingsEmailDeliverySettings
type NotificationsSettingsEmailDeliverySettings struct {
	// A string that specifies the organization's SMTP server.
	Host *string `json:"host,omitempty"`
	// An integer that specifies the port used by the organization's SMTP server to send emails (default `465`). Note that the protocol used depends upon the port specified. If you specify port `25`, `587`, or `2525`, SMTP with `STARTTLS` is used. Otherwise, `SMTPS` is used.
	Port *int32 `json:"port,omitempty"`
	// A string that specifies the organization's SMTP server's protocol.
	Protocol *string `json:"protocol,omitempty"`
	// A string that specifies the organization's SMTP server's username.
	Username *string `json:"username,omitempty"`
	// A string that specifies the organization's SMTP server's password.
	Password *string `json:"password,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	From *NotificationsSettingsEmailDeliverySettingsFrom `json:"from,omitempty"`
	ReplyTo *NotificationsSettingsEmailDeliverySettingsReplyTo `json:"replyTo,omitempty"`
}

// NewNotificationsSettingsEmailDeliverySettings instantiates a new NotificationsSettingsEmailDeliverySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsEmailDeliverySettings() *NotificationsSettingsEmailDeliverySettings {
	this := NotificationsSettingsEmailDeliverySettings{}
	var port int32 = 465
	this.Port = &port
	return &this
}

// NewNotificationsSettingsEmailDeliverySettingsWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsEmailDeliverySettingsWithDefaults() *NotificationsSettingsEmailDeliverySettings {
	this := NotificationsSettingsEmailDeliverySettings{}
	var port int32 = 465
	this.Port = &port
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NotificationsSettingsEmailDeliverySettings) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NotificationsSettingsEmailDeliverySettings) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *NotificationsSettingsEmailDeliverySettings) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationsSettingsEmailDeliverySettings) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NotificationsSettingsEmailDeliverySettings) SetPassword(v string) {
	o.Password = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *NotificationsSettingsEmailDeliverySettings) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetFrom() NotificationsSettingsEmailDeliverySettingsFrom {
	if o == nil || IsNil(o.From) {
		var ret NotificationsSettingsEmailDeliverySettingsFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetFromOk() (*NotificationsSettingsEmailDeliverySettingsFrom, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NotificationsSettingsEmailDeliverySettingsFrom and assigns it to the From field.
func (o *NotificationsSettingsEmailDeliverySettings) SetFrom(v NotificationsSettingsEmailDeliverySettingsFrom) {
	o.From = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettings) GetReplyTo() NotificationsSettingsEmailDeliverySettingsReplyTo {
	if o == nil || IsNil(o.ReplyTo) {
		var ret NotificationsSettingsEmailDeliverySettingsReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettings) GetReplyToOk() (*NotificationsSettingsEmailDeliverySettingsReplyTo, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettings) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given NotificationsSettingsEmailDeliverySettingsReplyTo and assigns it to the ReplyTo field.
func (o *NotificationsSettingsEmailDeliverySettings) SetReplyTo(v NotificationsSettingsEmailDeliverySettingsReplyTo) {
	o.ReplyTo = &v
}

func (o NotificationsSettingsEmailDeliverySettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsEmailDeliverySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	// skip: protocol is readOnly
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsEmailDeliverySettings struct {
	value *NotificationsSettingsEmailDeliverySettings
	isSet bool
}

func (v NullableNotificationsSettingsEmailDeliverySettings) Get() *NotificationsSettingsEmailDeliverySettings {
	return v.value
}

func (v *NullableNotificationsSettingsEmailDeliverySettings) Set(val *NotificationsSettingsEmailDeliverySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsEmailDeliverySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsEmailDeliverySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsEmailDeliverySettings(val *NotificationsSettingsEmailDeliverySettings) *NullableNotificationsSettingsEmailDeliverySettings {
	return &NullableNotificationsSettingsEmailDeliverySettings{value: val, isSet: true}
}

func (v NullableNotificationsSettingsEmailDeliverySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsEmailDeliverySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


