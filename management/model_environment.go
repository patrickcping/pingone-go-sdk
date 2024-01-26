/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the Environment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Environment{}

// Environment struct for Environment
type Environment struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	BillOfMaterials *BillOfMaterials `json:"billOfMaterials,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the description of the population.
	Description *string `json:"description,omitempty"`
	// The URL referencing the image to use for the environment icon. The supported image types are JPEG/JPG, PNG, and GIF.
	Icon *string `json:"icon,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	License EnvironmentLicense `json:"license"`
	// A string that specifies the environment name, which must be provided and must be unique within an organization.
	Name string `json:"name"`
	Organization *EnvironmentOrganization `json:"organization,omitempty"`
	Region EnvironmentRegion `json:"region"`
	Type EnumEnvironmentType `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment(license EnvironmentLicense, name string, region EnvironmentRegion, type_ EnumEnvironmentType) *Environment {
	this := Environment{}
	this.License = license
	this.Name = name
	this.Region = region
	this.Type = type_
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Environment) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Environment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *Environment) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetBillOfMaterials returns the BillOfMaterials field value if set, zero value otherwise.
func (o *Environment) GetBillOfMaterials() BillOfMaterials {
	if o == nil || IsNil(o.BillOfMaterials) {
		var ret BillOfMaterials
		return ret
	}
	return *o.BillOfMaterials
}

// GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetBillOfMaterialsOk() (*BillOfMaterials, bool) {
	if o == nil || IsNil(o.BillOfMaterials) {
		return nil, false
	}
	return o.BillOfMaterials, true
}

// HasBillOfMaterials returns a boolean if a field has been set.
func (o *Environment) HasBillOfMaterials() bool {
	if o != nil && !IsNil(o.BillOfMaterials) {
		return true
	}

	return false
}

// SetBillOfMaterials gets a reference to the given BillOfMaterials and assigns it to the BillOfMaterials field.
func (o *Environment) SetBillOfMaterials(v BillOfMaterials) {
	o.BillOfMaterials = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Environment) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Environment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Environment) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Environment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Environment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Environment) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Environment) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Environment) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Environment) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Environment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Environment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Environment) SetId(v string) {
	o.Id = &v
}

// GetLicense returns the License field value
func (o *Environment) GetLicense() EnvironmentLicense {
	if o == nil {
		var ret EnvironmentLicense
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *Environment) GetLicenseOk() (*EnvironmentLicense, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *Environment) SetLicense(v EnvironmentLicense) {
	o.License = v
}

// GetName returns the Name field value
func (o *Environment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Environment) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Environment) GetOrganization() EnvironmentOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret EnvironmentOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetOrganizationOk() (*EnvironmentOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Environment) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given EnvironmentOrganization and assigns it to the Organization field.
func (o *Environment) SetOrganization(v EnvironmentOrganization) {
	o.Organization = &v
}

// GetRegion returns the Region field value
func (o *Environment) GetRegion() EnvironmentRegion {
	if o == nil {
		var ret EnvironmentRegion
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Environment) GetRegionOk() (*EnvironmentRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Environment) SetRegion(v EnvironmentRegion) {
	o.Region = v
}

// GetType returns the Type field value
func (o *Environment) GetType() EnumEnvironmentType {
	if o == nil {
		var ret EnumEnvironmentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Environment) GetTypeOk() (*EnumEnvironmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Environment) SetType(v EnumEnvironmentType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Environment) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Environment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Environment) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Environment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.BillOfMaterials) {
		toSerialize["billOfMaterials"] = o.BillOfMaterials
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["license"] = o.License
	toSerialize["name"] = o.Name
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	toSerialize["region"] = o.Region
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


