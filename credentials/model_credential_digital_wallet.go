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

// checks if the CredentialDigitalWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialDigitalWallet{}

// CredentialDigitalWallet struct for CredentialDigitalWallet
type CredentialDigitalWallet struct {
	Links       map[string]LinksHATEOASValue        `json:"_links,omitempty"`
	Application *CredentialDigitalWalletApplication `json:"application,omitempty"`
	// A string that specifies the date and time the credential digital wallet was created.
	CreatedAt                *time.Time                                       `json:"createdAt,omitempty"`
	DigitalWalletApplication *CredentialDigitalWalletDigitalWalletApplication `json:"digitalWalletApplication,omitempty"`
	Environment              *ObjectEnvironment                               `json:"environment,omitempty"`
	// A string that specifies the identifier (UUID) associated with the credential digital wallet app.
	Id           *string                              `json:"id,omitempty"`
	Notification *CredentialDigitalWalletNotification `json:"notification,omitempty"`
	// A string that specifies the date and time the credential digital wallet was last updated; can be null.
	UpdatedAt *time.Time                   `json:"updatedAt,omitempty"`
	User      *CredentialDigitalWalletUser `json:"user,omitempty"`
	// A string that specifies the status of the wallet.
	Status         *string                                `json:"status,omitempty"`
	PairingSession *CredentialDigitalWalletPairingSession `json:"pairingSession,omitempty"`
}

// NewCredentialDigitalWallet instantiates a new CredentialDigitalWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialDigitalWallet() *CredentialDigitalWallet {
	this := CredentialDigitalWallet{}
	return &this
}

// NewCredentialDigitalWalletWithDefaults instantiates a new CredentialDigitalWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialDigitalWalletWithDefaults() *CredentialDigitalWallet {
	this := CredentialDigitalWallet{}
	return &this
}

func (o CredentialDigitalWallet) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o CredentialDigitalWallet) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o CredentialDigitalWallet) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o CredentialDigitalWallet) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o CredentialDigitalWallet) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o CredentialDigitalWallet) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o CredentialDigitalWallet) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o CredentialDigitalWallet) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o CredentialDigitalWallet) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o CredentialDigitalWallet) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o CredentialDigitalWallet) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o CredentialDigitalWallet) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o CredentialDigitalWallet) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *CredentialDigitalWallet) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetApplication() CredentialDigitalWalletApplication {
	if o == nil || IsNil(o.Application) {
		var ret CredentialDigitalWalletApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetApplicationOk() (*CredentialDigitalWalletApplication, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given CredentialDigitalWalletApplication and assigns it to the Application field.
func (o *CredentialDigitalWallet) SetApplication(v CredentialDigitalWalletApplication) {
	o.Application = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialDigitalWallet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDigitalWalletApplication returns the DigitalWalletApplication field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetDigitalWalletApplication() CredentialDigitalWalletDigitalWalletApplication {
	if o == nil || IsNil(o.DigitalWalletApplication) {
		var ret CredentialDigitalWalletDigitalWalletApplication
		return ret
	}
	return *o.DigitalWalletApplication
}

// GetDigitalWalletApplicationOk returns a tuple with the DigitalWalletApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetDigitalWalletApplicationOk() (*CredentialDigitalWalletDigitalWalletApplication, bool) {
	if o == nil || IsNil(o.DigitalWalletApplication) {
		return nil, false
	}
	return o.DigitalWalletApplication, true
}

// HasDigitalWalletApplication returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasDigitalWalletApplication() bool {
	if o != nil && !IsNil(o.DigitalWalletApplication) {
		return true
	}

	return false
}

// SetDigitalWalletApplication gets a reference to the given CredentialDigitalWalletDigitalWalletApplication and assigns it to the DigitalWalletApplication field.
func (o *CredentialDigitalWallet) SetDigitalWalletApplication(v CredentialDigitalWalletDigitalWalletApplication) {
	o.DigitalWalletApplication = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CredentialDigitalWallet) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialDigitalWallet) SetId(v string) {
	o.Id = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetNotification() CredentialDigitalWalletNotification {
	if o == nil || IsNil(o.Notification) {
		var ret CredentialDigitalWalletNotification
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetNotificationOk() (*CredentialDigitalWalletNotification, bool) {
	if o == nil || IsNil(o.Notification) {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasNotification() bool {
	if o != nil && !IsNil(o.Notification) {
		return true
	}

	return false
}

// SetNotification gets a reference to the given CredentialDigitalWalletNotification and assigns it to the Notification field.
func (o *CredentialDigitalWallet) SetNotification(v CredentialDigitalWalletNotification) {
	o.Notification = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CredentialDigitalWallet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetUser() CredentialDigitalWalletUser {
	if o == nil || IsNil(o.User) {
		var ret CredentialDigitalWalletUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetUserOk() (*CredentialDigitalWalletUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CredentialDigitalWalletUser and assigns it to the User field.
func (o *CredentialDigitalWallet) SetUser(v CredentialDigitalWalletUser) {
	o.User = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CredentialDigitalWallet) SetStatus(v string) {
	o.Status = &v
}

// GetPairingSession returns the PairingSession field value if set, zero value otherwise.
func (o *CredentialDigitalWallet) GetPairingSession() CredentialDigitalWalletPairingSession {
	if o == nil || IsNil(o.PairingSession) {
		var ret CredentialDigitalWalletPairingSession
		return ret
	}
	return *o.PairingSession
}

// GetPairingSessionOk returns a tuple with the PairingSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWallet) GetPairingSessionOk() (*CredentialDigitalWalletPairingSession, bool) {
	if o == nil || IsNil(o.PairingSession) {
		return nil, false
	}
	return o.PairingSession, true
}

// HasPairingSession returns a boolean if a field has been set.
func (o *CredentialDigitalWallet) HasPairingSession() bool {
	if o != nil && !IsNil(o.PairingSession) {
		return true
	}

	return false
}

// SetPairingSession gets a reference to the given CredentialDigitalWalletPairingSession and assigns it to the PairingSession field.
func (o *CredentialDigitalWallet) SetPairingSession(v CredentialDigitalWalletPairingSession) {
	o.PairingSession = &v
}

func (o CredentialDigitalWallet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialDigitalWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DigitalWalletApplication) {
		toSerialize["digitalWalletApplication"] = o.DigitalWalletApplication
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Notification) {
		toSerialize["notification"] = o.Notification
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PairingSession) {
		toSerialize["pairingSession"] = o.PairingSession
	}
	return toSerialize, nil
}

type NullableCredentialDigitalWallet struct {
	value *CredentialDigitalWallet
	isSet bool
}

func (v NullableCredentialDigitalWallet) Get() *CredentialDigitalWallet {
	return v.value
}

func (v *NullableCredentialDigitalWallet) Set(val *CredentialDigitalWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialDigitalWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialDigitalWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialDigitalWallet(val *CredentialDigitalWallet) *NullableCredentialDigitalWallet {
	return &NullableCredentialDigitalWallet{value: val, isSet: true}
}

func (v NullableCredentialDigitalWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialDigitalWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
