/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the AuthorizeEditorDataDefinitionsServiceDefinitionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataDefinitionsServiceDefinitionDTO{}

// AuthorizeEditorDataDefinitionsServiceDefinitionDTO struct for AuthorizeEditorDataDefinitionsServiceDefinitionDTO
type AuthorizeEditorDataDefinitionsServiceDefinitionDTO struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// HAL embedded resources
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	Version *string `json:"version,omitempty"`
	Name string `json:"name"`
	FullName *string `json:"fullName,omitempty"`
	Description *string `json:"description,omitempty"`
	Parent *AuthorizeEditorDataReferenceObjectDTO `json:"parent,omitempty"`
	Type *string `json:"type,omitempty"`
	CacheSettings *AuthorizeEditorDataCacheSettingsDTO `json:"cacheSettings,omitempty"`
	ServiceType string `json:"serviceType"`
}

// NewAuthorizeEditorDataDefinitionsServiceDefinitionDTO instantiates a new AuthorizeEditorDataDefinitionsServiceDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataDefinitionsServiceDefinitionDTO(name string, serviceType string) *AuthorizeEditorDataDefinitionsServiceDefinitionDTO {
	this := AuthorizeEditorDataDefinitionsServiceDefinitionDTO{}
	this.Name = name
	this.ServiceType = serviceType
	return &this
}

// NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataDefinitionsServiceDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataDefinitionsServiceDefinitionDTO {
	this := AuthorizeEditorDataDefinitionsServiceDefinitionDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetName(v string) {
	o.Name = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetFullName(v string) {
	o.FullName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO {
	if o == nil || IsNil(o.Parent) {
		var ret AuthorizeEditorDataReferenceObjectDTO
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given AuthorizeEditorDataReferenceObjectDTO and assigns it to the Parent field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO) {
	o.Parent = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetType(v string) {
	o.Type = &v
}

// GetCacheSettings returns the CacheSettings field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO {
	if o == nil || IsNil(o.CacheSettings) {
		var ret AuthorizeEditorDataCacheSettingsDTO
		return ret
	}
	return *o.CacheSettings
}

// GetCacheSettingsOk returns a tuple with the CacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool) {
	if o == nil || IsNil(o.CacheSettings) {
		return nil, false
	}
	return o.CacheSettings, true
}

// HasCacheSettings returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) HasCacheSettings() bool {
	if o != nil && !IsNil(o.CacheSettings) {
		return true
	}

	return false
}

// SetCacheSettings gets a reference to the given AuthorizeEditorDataCacheSettingsDTO and assigns it to the CacheSettings field.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO) {
	o.CacheSettings = &v
}

// GetServiceType returns the ServiceType field value
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) SetServiceType(v string) {
	o.ServiceType = v
}

func (o AuthorizeEditorDataDefinitionsServiceDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataDefinitionsServiceDefinitionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CacheSettings) {
		toSerialize["cacheSettings"] = o.CacheSettings
	}
	toSerialize["serviceType"] = o.ServiceType
	return toSerialize, nil
}

type NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO struct {
	value *AuthorizeEditorDataDefinitionsServiceDefinitionDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) Get() *AuthorizeEditorDataDefinitionsServiceDefinitionDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) Set(val *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO(val *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) *NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO {
	return &NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataDefinitionsServiceDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


