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

// checks if the Population type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Population{}

// Population struct for Population
type Population struct {
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The population to use as the default population for the environment. Only one population per environment can be the default. New users are assigned to the default population if it exists, and the Population ID is not provided in the [Create User](https://apidocs.pingidentity.com/pingone/platform/v1/api/#post-create-user) request.
	Default *bool `json:"default,omitempty"`
	// A string that specifies the description of the population.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the population name, which must be provided and must be unique within an environment.
	Name string `json:"name"`
	PasswordPolicy *PopulationPasswordPolicy `json:"passwordPolicy,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The number of users that belong to the population
	UserCount *int32 `json:"userCount,omitempty"`
}

// NewPopulation instantiates a new Population object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopulation(name string) *Population {
	this := Population{}
	this.Name = name
	return &this
}

// NewPopulationWithDefaults instantiates a new Population object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopulationWithDefaults() *Population {
	this := Population{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Population) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Population) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Population) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Population) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Population) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Population) SetDefault(v bool) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Population) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Population) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Population) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Population) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Population) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Population) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Population) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Population) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Population) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Population) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Population) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Population) SetName(v string) {
	o.Name = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *Population) GetPasswordPolicy() PopulationPasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy) {
		var ret PopulationPasswordPolicy
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetPasswordPolicyOk() (*PopulationPasswordPolicy, bool) {
	if o == nil || IsNil(o.PasswordPolicy) {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *Population) HasPasswordPolicy() bool {
	if o != nil && !IsNil(o.PasswordPolicy) {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given PopulationPasswordPolicy and assigns it to the PasswordPolicy field.
func (o *Population) SetPasswordPolicy(v PopulationPasswordPolicy) {
	o.PasswordPolicy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Population) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Population) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Population) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *Population) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount) {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Population) GetUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCount) {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *Population) HasUserCount() bool {
	if o != nil && !IsNil(o.UserCount) {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *Population) SetUserCount(v int32) {
	o.UserCount = &v
}

func (o Population) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Population) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: createdAt is readOnly
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.PasswordPolicy) {
		toSerialize["passwordPolicy"] = o.PasswordPolicy
	}
	// skip: updatedAt is readOnly
	// skip: userCount is readOnly
	return toSerialize, nil
}

type NullablePopulation struct {
	value *Population
	isSet bool
}

func (v NullablePopulation) Get() *Population {
	return v.value
}

func (v *NullablePopulation) Set(val *Population) {
	v.value = val
	v.isSet = true
}

func (v NullablePopulation) IsSet() bool {
	return v.isSet
}

func (v *NullablePopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopulation(val *Population) *NullablePopulation {
	return &NullablePopulation{value: val, isSet: true}
}

func (v NullablePopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


