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

// checks if the ApplicationSAML type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSAML{}

// ApplicationSAML struct for ApplicationSAML
type ApplicationSAML struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
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
	HomePageUrl *string `json:"homePageUrl,omitempty"`
	// A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property.
	AcsUrls []string `json:"acsUrls"`
	// An integer that specifies the assertion validity duration in seconds. This is a required property.
	AssertionDuration int32 `json:"assertionDuration"`
	// A boolean that specifies whether the SAML assertion itself should be signed. The default value is `true`.
	AssertionSigned *bool `json:"assertionSigned,omitempty"`
	CorsSettings *ApplicationCorsSettings `json:"corsSettings,omitempty"`
	// This is used as the RelayState parameter by the IdP to deep link into the application after authentication. This value can be overridden by the applicationUrl query parameter for GET Identity Provider Initiated SSO. Although both of these parameters are generally URLs, because they are used as deep links, this is not enforced. If neither defaultTargetUrl nor applicationUrl is specified during a SAML authentication flow, no RelayState value is supplied to the application. The defaultTargetUrl (or the applicationUrl) value is passed to the SAML application's ACS URL as a separate RelayState key value (not within the SAMLResponse key value).
	DefaultTargetUrl *string `json:"defaultTargetUrl,omitempty"`
	// Indicates whether `requestedAuthnContext` is taken into account in policy decision-making during authentication.
	EnableRequestedAuthnContext *bool `json:"enableRequestedAuthnContext,omitempty"`
	IdpSigning *ApplicationSAMLAllOfIdpSigning `json:"idpSigning,omitempty"`
	// A string that specifies the format of the Subject NameID attibute in the SAML assertion
	NameIdFormat *string `json:"nameIdFormat,omitempty"`
	// A boolean that specifies whether the SAML assertion response itself should be signed. The default value is `false`.
	ResponseSigned *bool `json:"responseSigned,omitempty"`
	SloBinding *EnumApplicationSAMLSloBinding `json:"sloBinding,omitempty"`
	// A string that specifies the logout endpoint URL. This is an optional property. However, if a sloEndpoint logout endpoint URL is not defined, logout actions result in an error.
	SloEndpoint *string `json:"sloEndpoint,omitempty"`
	// A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response.
	SloResponseEndpoint *string `json:"sloResponseEndpoint,omitempty"`
	// Defines how long PingOne can exchange logout messages with the application, specifically a `LogoutRequest` from the application, since the initial request. PingOne can also send a `LogoutRequest` to the application when a single logout is initiated by the user from other session participants, such as an application or identity provider. This setting is per application. The SLO logout is separate from the user session logout that revokes all tokens.
	SloWindow *int32 `json:"sloWindow,omitempty"`
	SpEncryption *ApplicationSAMLAllOfSpEncryption `json:"spEncryption,omitempty"`
	// A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment.
	SpEntityId string `json:"spEntityId"`
	SpVerification *ApplicationSAMLAllOfSpVerification `json:"spVerification,omitempty"`
}

// NewApplicationSAML instantiates a new ApplicationSAML object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAML(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, acsUrls []string, assertionDuration int32, spEntityId string) *ApplicationSAML {
	this := ApplicationSAML{}
	this.Enabled = enabled
	this.Name = name
	this.Protocol = protocol
	this.Type = type_
	this.AcsUrls = acsUrls
	this.AssertionDuration = assertionDuration
	var assertionSigned bool = true
	this.AssertionSigned = &assertionSigned
	var responseSigned bool = false
	this.ResponseSigned = &responseSigned
	var sloBinding EnumApplicationSAMLSloBinding = ENUMAPPLICATIONSAMLSLOBINDING_POST
	this.SloBinding = &sloBinding
	this.SpEntityId = spEntityId
	return &this
}

// NewApplicationSAMLWithDefaults instantiates a new ApplicationSAML object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLWithDefaults() *ApplicationSAML {
	this := ApplicationSAML{}
	var assertionSigned bool = true
	this.AssertionSigned = &assertionSigned
	var responseSigned bool = false
	this.ResponseSigned = &responseSigned
	var sloBinding EnumApplicationSAMLSloBinding = ENUMAPPLICATIONSAMLSLOBINDING_POST
	this.SloBinding = &sloBinding
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationSAML) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationSAML) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *ApplicationSAML) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAccessControl() ApplicationAccessControl {
	if o == nil || IsNil(o.AccessControl) {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *ApplicationSAML) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationSAML) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationSAML) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApplicationSAML) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationSAML) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationSAML) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationSAML) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationSAML) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationSAML) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationSAML) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationSAML) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationSAML) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetHiddenFromAppPortal returns the HiddenFromAppPortal field value if set, zero value otherwise.
func (o *ApplicationSAML) GetHiddenFromAppPortal() bool {
	if o == nil || IsNil(o.HiddenFromAppPortal) {
		var ret bool
		return ret
	}
	return *o.HiddenFromAppPortal
}

// GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetHiddenFromAppPortalOk() (*bool, bool) {
	if o == nil || IsNil(o.HiddenFromAppPortal) {
		return nil, false
	}
	return o.HiddenFromAppPortal, true
}

// HasHiddenFromAppPortal returns a boolean if a field has been set.
func (o *ApplicationSAML) HasHiddenFromAppPortal() bool {
	if o != nil && !IsNil(o.HiddenFromAppPortal) {
		return true
	}

	return false
}

// SetHiddenFromAppPortal gets a reference to the given bool and assigns it to the HiddenFromAppPortal field.
func (o *ApplicationSAML) SetHiddenFromAppPortal(v bool) {
	o.HiddenFromAppPortal = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationSAML) GetIcon() ApplicationIcon {
	if o == nil || IsNil(o.Icon) {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationSAML) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *ApplicationSAML) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationSAML) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationSAML) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationSAML) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetLoginPageUrl() string {
	if o == nil || IsNil(o.LoginPageUrl) {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginPageUrl) {
		return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasLoginPageUrl() bool {
	if o != nil && !IsNil(o.LoginPageUrl) {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *ApplicationSAML) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value
func (o *ApplicationSAML) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationSAML) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *ApplicationSAML) GetProtocol() EnumApplicationProtocol {
	if o == nil {
		var ret EnumApplicationProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetProtocolOk() (*EnumApplicationProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *ApplicationSAML) SetProtocol(v EnumApplicationProtocol) {
	o.Protocol = v
}

// GetType returns the Type field value
func (o *ApplicationSAML) GetType() EnumApplicationType {
	if o == nil {
		var ret EnumApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetTypeOk() (*EnumApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationSAML) SetType(v EnumApplicationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationSAML) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationSAML) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApplicationSAML) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetHomePageUrl() string {
	if o == nil || IsNil(o.HomePageUrl) {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetHomePageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomePageUrl) {
		return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasHomePageUrl() bool {
	if o != nil && !IsNil(o.HomePageUrl) {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *ApplicationSAML) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetAcsUrls returns the AcsUrls field value
func (o *ApplicationSAML) GetAcsUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcsUrls
}

// GetAcsUrlsOk returns a tuple with the AcsUrls field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAcsUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcsUrls, true
}

// SetAcsUrls sets field value
func (o *ApplicationSAML) SetAcsUrls(v []string) {
	o.AcsUrls = v
}

// GetAssertionDuration returns the AssertionDuration field value
func (o *ApplicationSAML) GetAssertionDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssertionDuration
}

// GetAssertionDurationOk returns a tuple with the AssertionDuration field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAssertionDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionDuration, true
}

// SetAssertionDuration sets field value
func (o *ApplicationSAML) SetAssertionDuration(v int32) {
	o.AssertionDuration = v
}

// GetAssertionSigned returns the AssertionSigned field value if set, zero value otherwise.
func (o *ApplicationSAML) GetAssertionSigned() bool {
	if o == nil || IsNil(o.AssertionSigned) {
		var ret bool
		return ret
	}
	return *o.AssertionSigned
}

// GetAssertionSignedOk returns a tuple with the AssertionSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetAssertionSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.AssertionSigned) {
		return nil, false
	}
	return o.AssertionSigned, true
}

// HasAssertionSigned returns a boolean if a field has been set.
func (o *ApplicationSAML) HasAssertionSigned() bool {
	if o != nil && !IsNil(o.AssertionSigned) {
		return true
	}

	return false
}

// SetAssertionSigned gets a reference to the given bool and assigns it to the AssertionSigned field.
func (o *ApplicationSAML) SetAssertionSigned(v bool) {
	o.AssertionSigned = &v
}

// GetCorsSettings returns the CorsSettings field value if set, zero value otherwise.
func (o *ApplicationSAML) GetCorsSettings() ApplicationCorsSettings {
	if o == nil || IsNil(o.CorsSettings) {
		var ret ApplicationCorsSettings
		return ret
	}
	return *o.CorsSettings
}

// GetCorsSettingsOk returns a tuple with the CorsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetCorsSettingsOk() (*ApplicationCorsSettings, bool) {
	if o == nil || IsNil(o.CorsSettings) {
		return nil, false
	}
	return o.CorsSettings, true
}

// HasCorsSettings returns a boolean if a field has been set.
func (o *ApplicationSAML) HasCorsSettings() bool {
	if o != nil && !IsNil(o.CorsSettings) {
		return true
	}

	return false
}

// SetCorsSettings gets a reference to the given ApplicationCorsSettings and assigns it to the CorsSettings field.
func (o *ApplicationSAML) SetCorsSettings(v ApplicationCorsSettings) {
	o.CorsSettings = &v
}

// GetDefaultTargetUrl returns the DefaultTargetUrl field value if set, zero value otherwise.
func (o *ApplicationSAML) GetDefaultTargetUrl() string {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		var ret string
		return ret
	}
	return *o.DefaultTargetUrl
}

// GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetDefaultTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		return nil, false
	}
	return o.DefaultTargetUrl, true
}

// HasDefaultTargetUrl returns a boolean if a field has been set.
func (o *ApplicationSAML) HasDefaultTargetUrl() bool {
	if o != nil && !IsNil(o.DefaultTargetUrl) {
		return true
	}

	return false
}

// SetDefaultTargetUrl gets a reference to the given string and assigns it to the DefaultTargetUrl field.
func (o *ApplicationSAML) SetDefaultTargetUrl(v string) {
	o.DefaultTargetUrl = &v
}

// GetEnableRequestedAuthnContext returns the EnableRequestedAuthnContext field value if set, zero value otherwise.
func (o *ApplicationSAML) GetEnableRequestedAuthnContext() bool {
	if o == nil || IsNil(o.EnableRequestedAuthnContext) {
		var ret bool
		return ret
	}
	return *o.EnableRequestedAuthnContext
}

// GetEnableRequestedAuthnContextOk returns a tuple with the EnableRequestedAuthnContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetEnableRequestedAuthnContextOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRequestedAuthnContext) {
		return nil, false
	}
	return o.EnableRequestedAuthnContext, true
}

// HasEnableRequestedAuthnContext returns a boolean if a field has been set.
func (o *ApplicationSAML) HasEnableRequestedAuthnContext() bool {
	if o != nil && !IsNil(o.EnableRequestedAuthnContext) {
		return true
	}

	return false
}

// SetEnableRequestedAuthnContext gets a reference to the given bool and assigns it to the EnableRequestedAuthnContext field.
func (o *ApplicationSAML) SetEnableRequestedAuthnContext(v bool) {
	o.EnableRequestedAuthnContext = &v
}

// GetIdpSigning returns the IdpSigning field value if set, zero value otherwise.
func (o *ApplicationSAML) GetIdpSigning() ApplicationSAMLAllOfIdpSigning {
	if o == nil || IsNil(o.IdpSigning) {
		var ret ApplicationSAMLAllOfIdpSigning
		return ret
	}
	return *o.IdpSigning
}

// GetIdpSigningOk returns a tuple with the IdpSigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetIdpSigningOk() (*ApplicationSAMLAllOfIdpSigning, bool) {
	if o == nil || IsNil(o.IdpSigning) {
		return nil, false
	}
	return o.IdpSigning, true
}

// HasIdpSigning returns a boolean if a field has been set.
func (o *ApplicationSAML) HasIdpSigning() bool {
	if o != nil && !IsNil(o.IdpSigning) {
		return true
	}

	return false
}

// SetIdpSigning gets a reference to the given ApplicationSAMLAllOfIdpSigning and assigns it to the IdpSigning field.
func (o *ApplicationSAML) SetIdpSigning(v ApplicationSAMLAllOfIdpSigning) {
	o.IdpSigning = &v
}

// GetNameIdFormat returns the NameIdFormat field value if set, zero value otherwise.
func (o *ApplicationSAML) GetNameIdFormat() string {
	if o == nil || IsNil(o.NameIdFormat) {
		var ret string
		return ret
	}
	return *o.NameIdFormat
}

// GetNameIdFormatOk returns a tuple with the NameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetNameIdFormatOk() (*string, bool) {
	if o == nil || IsNil(o.NameIdFormat) {
		return nil, false
	}
	return o.NameIdFormat, true
}

// HasNameIdFormat returns a boolean if a field has been set.
func (o *ApplicationSAML) HasNameIdFormat() bool {
	if o != nil && !IsNil(o.NameIdFormat) {
		return true
	}

	return false
}

// SetNameIdFormat gets a reference to the given string and assigns it to the NameIdFormat field.
func (o *ApplicationSAML) SetNameIdFormat(v string) {
	o.NameIdFormat = &v
}

// GetResponseSigned returns the ResponseSigned field value if set, zero value otherwise.
func (o *ApplicationSAML) GetResponseSigned() bool {
	if o == nil || IsNil(o.ResponseSigned) {
		var ret bool
		return ret
	}
	return *o.ResponseSigned
}

// GetResponseSignedOk returns a tuple with the ResponseSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetResponseSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseSigned) {
		return nil, false
	}
	return o.ResponseSigned, true
}

// HasResponseSigned returns a boolean if a field has been set.
func (o *ApplicationSAML) HasResponseSigned() bool {
	if o != nil && !IsNil(o.ResponseSigned) {
		return true
	}

	return false
}

// SetResponseSigned gets a reference to the given bool and assigns it to the ResponseSigned field.
func (o *ApplicationSAML) SetResponseSigned(v bool) {
	o.ResponseSigned = &v
}

// GetSloBinding returns the SloBinding field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloBinding() EnumApplicationSAMLSloBinding {
	if o == nil || IsNil(o.SloBinding) {
		var ret EnumApplicationSAMLSloBinding
		return ret
	}
	return *o.SloBinding
}

// GetSloBindingOk returns a tuple with the SloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloBindingOk() (*EnumApplicationSAMLSloBinding, bool) {
	if o == nil || IsNil(o.SloBinding) {
		return nil, false
	}
	return o.SloBinding, true
}

// HasSloBinding returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloBinding() bool {
	if o != nil && !IsNil(o.SloBinding) {
		return true
	}

	return false
}

// SetSloBinding gets a reference to the given EnumApplicationSAMLSloBinding and assigns it to the SloBinding field.
func (o *ApplicationSAML) SetSloBinding(v EnumApplicationSAMLSloBinding) {
	o.SloBinding = &v
}

// GetSloEndpoint returns the SloEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloEndpoint() string {
	if o == nil || IsNil(o.SloEndpoint) {
		var ret string
		return ret
	}
	return *o.SloEndpoint
}

// GetSloEndpointOk returns a tuple with the SloEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.SloEndpoint) {
		return nil, false
	}
	return o.SloEndpoint, true
}

// HasSloEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloEndpoint() bool {
	if o != nil && !IsNil(o.SloEndpoint) {
		return true
	}

	return false
}

// SetSloEndpoint gets a reference to the given string and assigns it to the SloEndpoint field.
func (o *ApplicationSAML) SetSloEndpoint(v string) {
	o.SloEndpoint = &v
}

// GetSloResponseEndpoint returns the SloResponseEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloResponseEndpoint() string {
	if o == nil || IsNil(o.SloResponseEndpoint) {
		var ret string
		return ret
	}
	return *o.SloResponseEndpoint
}

// GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloResponseEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.SloResponseEndpoint) {
		return nil, false
	}
	return o.SloResponseEndpoint, true
}

// HasSloResponseEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloResponseEndpoint() bool {
	if o != nil && !IsNil(o.SloResponseEndpoint) {
		return true
	}

	return false
}

// SetSloResponseEndpoint gets a reference to the given string and assigns it to the SloResponseEndpoint field.
func (o *ApplicationSAML) SetSloResponseEndpoint(v string) {
	o.SloResponseEndpoint = &v
}

// GetSloWindow returns the SloWindow field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSloWindow() int32 {
	if o == nil || IsNil(o.SloWindow) {
		var ret int32
		return ret
	}
	return *o.SloWindow
}

// GetSloWindowOk returns a tuple with the SloWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSloWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.SloWindow) {
		return nil, false
	}
	return o.SloWindow, true
}

// HasSloWindow returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSloWindow() bool {
	if o != nil && !IsNil(o.SloWindow) {
		return true
	}

	return false
}

// SetSloWindow gets a reference to the given int32 and assigns it to the SloWindow field.
func (o *ApplicationSAML) SetSloWindow(v int32) {
	o.SloWindow = &v
}

// GetSpEncryption returns the SpEncryption field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSpEncryption() ApplicationSAMLAllOfSpEncryption {
	if o == nil || IsNil(o.SpEncryption) {
		var ret ApplicationSAMLAllOfSpEncryption
		return ret
	}
	return *o.SpEncryption
}

// GetSpEncryptionOk returns a tuple with the SpEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSpEncryptionOk() (*ApplicationSAMLAllOfSpEncryption, bool) {
	if o == nil || IsNil(o.SpEncryption) {
		return nil, false
	}
	return o.SpEncryption, true
}

// HasSpEncryption returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSpEncryption() bool {
	if o != nil && !IsNil(o.SpEncryption) {
		return true
	}

	return false
}

// SetSpEncryption gets a reference to the given ApplicationSAMLAllOfSpEncryption and assigns it to the SpEncryption field.
func (o *ApplicationSAML) SetSpEncryption(v ApplicationSAMLAllOfSpEncryption) {
	o.SpEncryption = &v
}

// GetSpEntityId returns the SpEntityId field value
func (o *ApplicationSAML) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSpEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *ApplicationSAML) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetSpVerification returns the SpVerification field value if set, zero value otherwise.
func (o *ApplicationSAML) GetSpVerification() ApplicationSAMLAllOfSpVerification {
	if o == nil || IsNil(o.SpVerification) {
		var ret ApplicationSAMLAllOfSpVerification
		return ret
	}
	return *o.SpVerification
}

// GetSpVerificationOk returns a tuple with the SpVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAML) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool) {
	if o == nil || IsNil(o.SpVerification) {
		return nil, false
	}
	return o.SpVerification, true
}

// HasSpVerification returns a boolean if a field has been set.
func (o *ApplicationSAML) HasSpVerification() bool {
	if o != nil && !IsNil(o.SpVerification) {
		return true
	}

	return false
}

// SetSpVerification gets a reference to the given ApplicationSAMLAllOfSpVerification and assigns it to the SpVerification field.
func (o *ApplicationSAML) SetSpVerification(v ApplicationSAMLAllOfSpVerification) {
	o.SpVerification = &v
}

func (o ApplicationSAML) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSAML) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HomePageUrl) {
		toSerialize["homePageUrl"] = o.HomePageUrl
	}
	toSerialize["acsUrls"] = o.AcsUrls
	toSerialize["assertionDuration"] = o.AssertionDuration
	if !IsNil(o.AssertionSigned) {
		toSerialize["assertionSigned"] = o.AssertionSigned
	}
	if !IsNil(o.CorsSettings) {
		toSerialize["corsSettings"] = o.CorsSettings
	}
	if !IsNil(o.DefaultTargetUrl) {
		toSerialize["defaultTargetUrl"] = o.DefaultTargetUrl
	}
	if !IsNil(o.EnableRequestedAuthnContext) {
		toSerialize["enableRequestedAuthnContext"] = o.EnableRequestedAuthnContext
	}
	if !IsNil(o.IdpSigning) {
		toSerialize["idpSigning"] = o.IdpSigning
	}
	if !IsNil(o.NameIdFormat) {
		toSerialize["nameIdFormat"] = o.NameIdFormat
	}
	if !IsNil(o.ResponseSigned) {
		toSerialize["responseSigned"] = o.ResponseSigned
	}
	if !IsNil(o.SloBinding) {
		toSerialize["sloBinding"] = o.SloBinding
	}
	if !IsNil(o.SloEndpoint) {
		toSerialize["sloEndpoint"] = o.SloEndpoint
	}
	if !IsNil(o.SloResponseEndpoint) {
		toSerialize["sloResponseEndpoint"] = o.SloResponseEndpoint
	}
	if !IsNil(o.SloWindow) {
		toSerialize["sloWindow"] = o.SloWindow
	}
	if !IsNil(o.SpEncryption) {
		toSerialize["spEncryption"] = o.SpEncryption
	}
	toSerialize["spEntityId"] = o.SpEntityId
	if !IsNil(o.SpVerification) {
		toSerialize["spVerification"] = o.SpVerification
	}
	return toSerialize, nil
}

type NullableApplicationSAML struct {
	value *ApplicationSAML
	isSet bool
}

func (v NullableApplicationSAML) Get() *ApplicationSAML {
	return v.value
}

func (v *NullableApplicationSAML) Set(val *ApplicationSAML) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAML) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAML) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAML(val *ApplicationSAML) *NullableApplicationSAML {
	return &NullableApplicationSAML{value: val, isSet: true}
}

func (v NullableApplicationSAML) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAML) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


