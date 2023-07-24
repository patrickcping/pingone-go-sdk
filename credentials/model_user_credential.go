/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the UserCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCredential{}

// UserCredential struct for UserCredential
type UserCredential struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CredentialType *CredentialDigitalWalletNotificationResultsInnerNotification `json:"credentialType,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Id *string `json:"id,omitempty"`
	Notification *CredentialDigitalWalletNotification `json:"notification,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	User *CredentialDigitalWalletNotificationResultsInnerNotification `json:"user,omitempty"`
	Userdata *UserCredentialUserdata `json:"userdata,omitempty"`
}

// NewUserCredential instantiates a new UserCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredential() *UserCredential {
	this := UserCredential{}
	return &this
}

// NewUserCredentialWithDefaults instantiates a new UserCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialWithDefaults() *UserCredential {
	this := UserCredential{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserCredential) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserCredential) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *UserCredential) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserCredential) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserCredential) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *UserCredential) GetCredentialType() CredentialDigitalWalletNotificationResultsInnerNotification {
	if o == nil || IsNil(o.CredentialType) {
		var ret CredentialDigitalWalletNotificationResultsInnerNotification
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetCredentialTypeOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool) {
	if o == nil || IsNil(o.CredentialType) {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *UserCredential) HasCredentialType() bool {
	if o != nil && !IsNil(o.CredentialType) {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given CredentialDigitalWalletNotificationResultsInnerNotification and assigns it to the CredentialType field.
func (o *UserCredential) SetCredentialType(v CredentialDigitalWalletNotificationResultsInnerNotification) {
	o.CredentialType = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *UserCredential) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *UserCredential) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *UserCredential) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserCredential) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserCredential) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UserCredential) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserCredential) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserCredential) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserCredential) SetId(v string) {
	o.Id = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *UserCredential) GetNotification() CredentialDigitalWalletNotification {
	if o == nil || IsNil(o.Notification) {
		var ret CredentialDigitalWalletNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetNotificationOk() (*CredentialDigitalWalletNotification, bool) {
	if o == nil || IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *UserCredential) HasNotification() bool {
	if o != nil && !IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given CredentialDigitalWalletNotification and assigns it to the Notification field.
func (o *UserCredential) SetNotification(v CredentialDigitalWalletNotification) {
	o.Notification = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserCredential) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserCredential) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserCredential) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserCredential) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserCredential) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserCredential) GetUser() CredentialDigitalWalletNotificationResultsInnerNotification {
	if o == nil || IsNil(o.User) {
		var ret CredentialDigitalWalletNotificationResultsInnerNotification
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetUserOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserCredential) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CredentialDigitalWalletNotificationResultsInnerNotification and assigns it to the User field.
func (o *UserCredential) SetUser(v CredentialDigitalWalletNotificationResultsInnerNotification) {
	o.User = &v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *UserCredential) GetUserdata() UserCredentialUserdata {
	if o == nil || IsNil(o.Userdata) {
		var ret UserCredentialUserdata
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredential) GetUserdataOk() (*UserCredentialUserdata, bool) {
	if o == nil || IsNil(o.Userdata) {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *UserCredential) HasUserdata() bool {
	if o != nil && !IsNil(o.Userdata) {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given UserCredentialUserdata and assigns it to the Userdata field.
func (o *UserCredential) SetUserdata(v UserCredentialUserdata) {
	o.Userdata = &v
}

func (o UserCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	// skip: createdAt is readOnly
	if !IsNil(o.CredentialType) {
		toSerialize["credentialType"] = o.CredentialType
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: expiresAt is readOnly
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	// skip: updatedAt is readOnly
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Userdata) {
		toSerialize["userdata"] = o.Userdata
	}
	return toSerialize, nil
}

type NullableUserCredential struct {
	value *UserCredential
	isSet bool
}

func (v NullableUserCredential) Get() *UserCredential {
	return v.value
}

func (v *NullableUserCredential) Set(val *UserCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredential(val *UserCredential) *NullableUserCredential {
	return &NullableUserCredential{value: val, isSet: true}
}

func (v NullableUserCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


