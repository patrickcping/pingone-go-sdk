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

// ApplicationPingOneSelfService struct for ApplicationPingOneSelfService
type ApplicationPingOneSelfService struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	AccessControl *ApplicationAccessControl `json:"accessControl,omitempty"`
	// A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request.
	AssignActorRoles *bool `json:"assignActorRoles,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A string that specifies the description of the application.
	Description *string `json:"description,omitempty"`
	// A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *ApplicationIcon `json:"icon,omitempty"`
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains.
	LoginPageUrl *string `json:"loginPageUrl,omitempty"`
	// A string that specifies the name of the application. This is a required property.
	Name string `json:"name"`
	Protocol EnumApplicationProtocol `json:"protocol"`
	// An array that specifies the list of labels associated with the application. Options are PING_FED_CONNECTION_INTEGRATION.
	Tags []EnumApplicationTags `json:"tags,omitempty"`
	Type EnumApplicationType `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A boolean that specifies whether the request query parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed.
	SupportUnsignedRequestObject *bool `json:"supportUnsignedRequestObject,omitempty"`
	PkceEnforcement *EnumApplicationOIDCPKCEOption `json:"pkceEnforcement,omitempty"`
	TokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod `json:"tokenEndpointAuthMethod"`
	// If `true`, shows the default theme footer on the self service application. Applies only if `applyDefaultTheme` is also `true`.
	EnableDefaultThemeFooter *bool `json:"enableDefaultThemeFooter,omitempty"`
	// If `true`, applies the default theme to the self service application.
	ApplyDefaultTheme bool `json:"applyDefaultTheme"`
}

// NewApplicationPingOneSelfService instantiates a new ApplicationPingOneSelfService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPingOneSelfService(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, applyDefaultTheme bool) *ApplicationPingOneSelfService {
	this := ApplicationPingOneSelfService{}
	this.Enabled = enabled
	this.Name = name
	this.Protocol = protocol
	this.Type = type_
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	this.ApplyDefaultTheme = applyDefaultTheme
	return &this
}

// NewApplicationPingOneSelfServiceWithDefaults instantiates a new ApplicationPingOneSelfService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPingOneSelfServiceWithDefaults() *ApplicationPingOneSelfService {
	this := ApplicationPingOneSelfService{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetLinks() map[string]interface{} {
	if o == nil || isNil(o.Links) {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Links) {
    return map[string]interface{}{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *ApplicationPingOneSelfService) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetAccessControl() ApplicationAccessControl {
	if o == nil || isNil(o.AccessControl) {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || isNil(o.AccessControl) {
    return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasAccessControl() bool {
	if o != nil && !isNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *ApplicationPingOneSelfService) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetAssignActorRoles returns the AssignActorRoles field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetAssignActorRoles() bool {
	if o == nil || isNil(o.AssignActorRoles) {
		var ret bool
		return ret
	}
	return *o.AssignActorRoles
}

// GetAssignActorRolesOk returns a tuple with the AssignActorRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetAssignActorRolesOk() (*bool, bool) {
	if o == nil || isNil(o.AssignActorRoles) {
    return nil, false
	}
	return o.AssignActorRoles, true
}

// HasAssignActorRoles returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasAssignActorRoles() bool {
	if o != nil && !isNil(o.AssignActorRoles) {
		return true
	}

	return false
}

// SetAssignActorRoles gets a reference to the given bool and assigns it to the AssignActorRoles field.
func (o *ApplicationPingOneSelfService) SetAssignActorRoles(v bool) {
	o.AssignActorRoles = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationPingOneSelfService) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationPingOneSelfService) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationPingOneSelfService) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationPingOneSelfService) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetEnvironment() ObjectEnvironment {
	if o == nil || isNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationPingOneSelfService) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetIcon() ApplicationIcon {
	if o == nil || isNil(o.Icon) {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *ApplicationPingOneSelfService) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationPingOneSelfService) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetLoginPageUrl() string {
	if o == nil || isNil(o.LoginPageUrl) {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || isNil(o.LoginPageUrl) {
    return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasLoginPageUrl() bool {
	if o != nil && !isNil(o.LoginPageUrl) {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *ApplicationPingOneSelfService) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value
func (o *ApplicationPingOneSelfService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationPingOneSelfService) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *ApplicationPingOneSelfService) GetProtocol() EnumApplicationProtocol {
	if o == nil {
		var ret EnumApplicationProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetProtocolOk() (*EnumApplicationProtocol, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ApplicationPingOneSelfService) SetProtocol(v EnumApplicationProtocol) {
	o.Protocol = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetTags() []EnumApplicationTags {
	if o == nil || isNil(o.Tags) {
		var ret []EnumApplicationTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetTagsOk() ([]EnumApplicationTags, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []EnumApplicationTags and assigns it to the Tags field.
func (o *ApplicationPingOneSelfService) SetTags(v []EnumApplicationTags) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ApplicationPingOneSelfService) GetType() EnumApplicationType {
	if o == nil {
		var ret EnumApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetTypeOk() (*EnumApplicationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationPingOneSelfService) SetType(v EnumApplicationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationPingOneSelfService) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetSupportUnsignedRequestObject() bool {
	if o == nil || isNil(o.SupportUnsignedRequestObject) {
		var ret bool
		return ret
	}
	return *o.SupportUnsignedRequestObject
}

// GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetSupportUnsignedRequestObjectOk() (*bool, bool) {
	if o == nil || isNil(o.SupportUnsignedRequestObject) {
    return nil, false
	}
	return o.SupportUnsignedRequestObject, true
}

// HasSupportUnsignedRequestObject returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasSupportUnsignedRequestObject() bool {
	if o != nil && !isNil(o.SupportUnsignedRequestObject) {
		return true
	}

	return false
}

// SetSupportUnsignedRequestObject gets a reference to the given bool and assigns it to the SupportUnsignedRequestObject field.
func (o *ApplicationPingOneSelfService) SetSupportUnsignedRequestObject(v bool) {
	o.SupportUnsignedRequestObject = &v
}

// GetPkceEnforcement returns the PkceEnforcement field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetPkceEnforcement() EnumApplicationOIDCPKCEOption {
	if o == nil || isNil(o.PkceEnforcement) {
		var ret EnumApplicationOIDCPKCEOption
		return ret
	}
	return *o.PkceEnforcement
}

// GetPkceEnforcementOk returns a tuple with the PkceEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool) {
	if o == nil || isNil(o.PkceEnforcement) {
    return nil, false
	}
	return o.PkceEnforcement, true
}

// HasPkceEnforcement returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasPkceEnforcement() bool {
	if o != nil && !isNil(o.PkceEnforcement) {
		return true
	}

	return false
}

// SetPkceEnforcement gets a reference to the given EnumApplicationOIDCPKCEOption and assigns it to the PkceEnforcement field.
func (o *ApplicationPingOneSelfService) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption) {
	o.PkceEnforcement = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value
func (o *ApplicationPingOneSelfService) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod {
	if o == nil {
		var ret EnumApplicationOIDCTokenAuthMethod
		return ret
	}

	return o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TokenEndpointAuthMethod, true
}

// SetTokenEndpointAuthMethod sets field value
func (o *ApplicationPingOneSelfService) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod) {
	o.TokenEndpointAuthMethod = v
}

// GetEnableDefaultThemeFooter returns the EnableDefaultThemeFooter field value if set, zero value otherwise.
func (o *ApplicationPingOneSelfService) GetEnableDefaultThemeFooter() bool {
	if o == nil || isNil(o.EnableDefaultThemeFooter) {
		var ret bool
		return ret
	}
	return *o.EnableDefaultThemeFooter
}

// GetEnableDefaultThemeFooterOk returns a tuple with the EnableDefaultThemeFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetEnableDefaultThemeFooterOk() (*bool, bool) {
	if o == nil || isNil(o.EnableDefaultThemeFooter) {
    return nil, false
	}
	return o.EnableDefaultThemeFooter, true
}

// HasEnableDefaultThemeFooter returns a boolean if a field has been set.
func (o *ApplicationPingOneSelfService) HasEnableDefaultThemeFooter() bool {
	if o != nil && !isNil(o.EnableDefaultThemeFooter) {
		return true
	}

	return false
}

// SetEnableDefaultThemeFooter gets a reference to the given bool and assigns it to the EnableDefaultThemeFooter field.
func (o *ApplicationPingOneSelfService) SetEnableDefaultThemeFooter(v bool) {
	o.EnableDefaultThemeFooter = &v
}

// GetApplyDefaultTheme returns the ApplyDefaultTheme field value
func (o *ApplicationPingOneSelfService) GetApplyDefaultTheme() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ApplyDefaultTheme
}

// GetApplyDefaultThemeOk returns a tuple with the ApplyDefaultTheme field value
// and a boolean to check if the value has been set.
func (o *ApplicationPingOneSelfService) GetApplyDefaultThemeOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApplyDefaultTheme, true
}

// SetApplyDefaultTheme sets field value
func (o *ApplicationPingOneSelfService) SetApplyDefaultTheme(v bool) {
	o.ApplyDefaultTheme = v
}

func (o ApplicationPingOneSelfService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !isNil(o.AssignActorRoles) {
		toSerialize["assignActorRoles"] = o.AssignActorRoles
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !isNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LoginPageUrl) {
		toSerialize["loginPageUrl"] = o.LoginPageUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.SupportUnsignedRequestObject) {
		toSerialize["supportUnsignedRequestObject"] = o.SupportUnsignedRequestObject
	}
	if !isNil(o.PkceEnforcement) {
		toSerialize["pkceEnforcement"] = o.PkceEnforcement
	}
	if true {
		toSerialize["tokenEndpointAuthMethod"] = o.TokenEndpointAuthMethod
	}
	if !isNil(o.EnableDefaultThemeFooter) {
		toSerialize["enableDefaultThemeFooter"] = o.EnableDefaultThemeFooter
	}
	if true {
		toSerialize["applyDefaultTheme"] = o.ApplyDefaultTheme
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPingOneSelfService struct {
	value *ApplicationPingOneSelfService
	isSet bool
}

func (v NullableApplicationPingOneSelfService) Get() *ApplicationPingOneSelfService {
	return v.value
}

func (v *NullableApplicationPingOneSelfService) Set(val *ApplicationPingOneSelfService) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPingOneSelfService) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPingOneSelfService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPingOneSelfService(val *ApplicationPingOneSelfService) *NullableApplicationPingOneSelfService {
	return &NullableApplicationPingOneSelfService{value: val, isSet: true}
}

func (v NullableApplicationPingOneSelfService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPingOneSelfService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


