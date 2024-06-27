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

// checks if the ApplicationExternalLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationExternalLink{}

// ApplicationExternalLink struct for ApplicationExternalLink
type ApplicationExternalLink struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	AccessControl *ApplicationAccessControl `json:"accessControl,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A string that specifies the description of the application.
	Description *string `json:"description,omitempty"`
	// A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A boolean to specify whether the application is hidden in the application portal despite the configured group access policy.
	HiddenFromAppPortal *bool `json:"hiddenFromAppPortal,omitempty"`
	Icon *ApplicationIcon `json:"icon,omitempty"`
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains.
	LoginPageUrl *string `json:"loginPageUrl,omitempty"`
	// A string that specifies the name of the application. This is a required property.
	Name string `json:"name"`
	Protocol EnumApplicationProtocol `json:"protocol"`
	Type EnumApplicationType `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A string that specifies the custom home page URL for the application.
	HomePageUrl string `json:"homePageUrl"`
}

// NewApplicationExternalLink instantiates a new ApplicationExternalLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationExternalLink(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, homePageUrl string) *ApplicationExternalLink {
	this := ApplicationExternalLink{}
	this.Enabled = enabled
	this.Name = name
	this.Protocol = protocol
	this.Type = type_
	this.HomePageUrl = homePageUrl
	return &this
}

// NewApplicationExternalLinkWithDefaults instantiates a new ApplicationExternalLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationExternalLinkWithDefaults() *ApplicationExternalLink {
	this := ApplicationExternalLink{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ApplicationExternalLink) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetAccessControl() ApplicationAccessControl {
	if o == nil || IsNil(o.AccessControl) {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *ApplicationExternalLink) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationExternalLink) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationExternalLink) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationExternalLink) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationExternalLink) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationExternalLink) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetHiddenFromAppPortal returns the HiddenFromAppPortal field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetHiddenFromAppPortal() bool {
	if o == nil || IsNil(o.HiddenFromAppPortal) {
		var ret bool
		return ret
	}
	return *o.HiddenFromAppPortal
}

// GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetHiddenFromAppPortalOk() (*bool, bool) {
	if o == nil || IsNil(o.HiddenFromAppPortal) {
		return nil, false
	}
	return o.HiddenFromAppPortal, true
}

// HasHiddenFromAppPortal returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasHiddenFromAppPortal() bool {
	if o != nil && !IsNil(o.HiddenFromAppPortal) {
		return true
	}

	return false
}

// SetHiddenFromAppPortal gets a reference to the given bool and assigns it to the HiddenFromAppPortal field.
func (o *ApplicationExternalLink) SetHiddenFromAppPortal(v bool) {
	o.HiddenFromAppPortal = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetIcon() ApplicationIcon {
	if o == nil || IsNil(o.Icon) {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *ApplicationExternalLink) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationExternalLink) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetLoginPageUrl() string {
	if o == nil || IsNil(o.LoginPageUrl) {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginPageUrl) {
		return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasLoginPageUrl() bool {
	if o != nil && !IsNil(o.LoginPageUrl) {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *ApplicationExternalLink) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value
func (o *ApplicationExternalLink) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationExternalLink) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *ApplicationExternalLink) GetProtocol() EnumApplicationProtocol {
	if o == nil {
		var ret EnumApplicationProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetProtocolOk() (*EnumApplicationProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ApplicationExternalLink) SetProtocol(v EnumApplicationProtocol) {
	o.Protocol = v
}

// GetType returns the Type field value
func (o *ApplicationExternalLink) GetType() EnumApplicationType {
	if o == nil {
		var ret EnumApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetTypeOk() (*EnumApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationExternalLink) SetType(v EnumApplicationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationExternalLink) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationExternalLink) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationExternalLink) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetHomePageUrl returns the HomePageUrl field value
func (o *ApplicationExternalLink) GetHomePageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLink) GetHomePageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomePageUrl, true
}

// SetHomePageUrl sets field value
func (o *ApplicationExternalLink) SetHomePageUrl(v string) {
	o.HomePageUrl = v
}

func (o ApplicationExternalLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationExternalLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.HiddenFromAppPortal) {
		toSerialize["hiddenFromAppPortal"] = o.HiddenFromAppPortal
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LoginPageUrl) {
		toSerialize["loginPageUrl"] = o.LoginPageUrl
	}
	toSerialize["name"] = o.Name
	toSerialize["protocol"] = o.Protocol
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["homePageUrl"] = o.HomePageUrl
	return toSerialize, nil
}

type NullableApplicationExternalLink struct {
	value *ApplicationExternalLink
	isSet bool
}

func (v NullableApplicationExternalLink) Get() *ApplicationExternalLink {
	return v.value
}

func (v *NullableApplicationExternalLink) Set(val *ApplicationExternalLink) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationExternalLink) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationExternalLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationExternalLink(val *ApplicationExternalLink) *NullableApplicationExternalLink {
	return &NullableApplicationExternalLink{value: val, isSet: true}
}

func (v NullableApplicationExternalLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationExternalLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


