/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationRolePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationRolePermission{}

// ApplicationRolePermission struct for ApplicationRolePermission
type ApplicationRolePermission struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The ID of the application resource permission to associate with this role.
	Id string `json:"id"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Key *string `json:"key,omitempty"`
	Description *string `json:"description,omitempty"`
	// The action associated with this permission.
	Action *string `json:"action,omitempty"`
	Resource *ApplicationRolePermissionResource `json:"resource,omitempty"`
}

type _ApplicationRolePermission ApplicationRolePermission

// NewApplicationRolePermission instantiates a new ApplicationRolePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRolePermission(id string) *ApplicationRolePermission {
	this := ApplicationRolePermission{}
	this.Id = id
	return &this
}

// NewApplicationRolePermissionWithDefaults instantiates a new ApplicationRolePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRolePermissionWithDefaults() *ApplicationRolePermission {
	this := ApplicationRolePermission{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ApplicationRolePermission) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetId returns the Id field value
func (o *ApplicationRolePermission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationRolePermission) SetId(v string) {
	o.Id = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationRolePermission) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApplicationRolePermission) SetKey(v string) {
	o.Key = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationRolePermission) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ApplicationRolePermission) SetAction(v string) {
	o.Action = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetResource() ApplicationRolePermissionResource {
	if o == nil || IsNil(o.Resource) {
		var ret ApplicationRolePermissionResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetResourceOk() (*ApplicationRolePermissionResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ApplicationRolePermissionResource and assigns it to the Resource field.
func (o *ApplicationRolePermission) SetResource(v ApplicationRolePermissionResource) {
	o.Resource = &v
}

func (o ApplicationRolePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationRolePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

func (o *ApplicationRolePermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationRolePermission := _ApplicationRolePermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationRolePermission)

	if err != nil {
		return err
	}

	*o = ApplicationRolePermission(varApplicationRolePermission)

	return err
}

type NullableApplicationRolePermission struct {
	value *ApplicationRolePermission
	isSet bool
}

func (v NullableApplicationRolePermission) Get() *ApplicationRolePermission {
	return v.value
}

func (v *NullableApplicationRolePermission) Set(val *ApplicationRolePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRolePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRolePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRolePermission(val *ApplicationRolePermission) *NullableApplicationRolePermission {
	return &NullableApplicationRolePermission{value: val, isSet: true}
}

func (v NullableApplicationRolePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRolePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


