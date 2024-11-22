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

// checks if the ApplicationResourcePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResourcePermission{}

// ApplicationResourcePermission struct for ApplicationResourcePermission
type ApplicationResourcePermission struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The action associated with this permission.
	Action string `json:"action"`
	// The resource's description.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	Resource *ApplicationResourcePermissionResource `json:"resource,omitempty"`
}

type _ApplicationResourcePermission ApplicationResourcePermission

// NewApplicationResourcePermission instantiates a new ApplicationResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourcePermission(action string) *ApplicationResourcePermission {
	this := ApplicationResourcePermission{}
	this.Action = action
	return &this
}

// NewApplicationResourcePermissionWithDefaults instantiates a new ApplicationResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourcePermissionWithDefaults() *ApplicationResourcePermission {
	this := ApplicationResourcePermission{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationResourcePermission) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationResourcePermission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ApplicationResourcePermission) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetAction returns the Action field value
func (o *ApplicationResourcePermission) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ApplicationResourcePermission) SetAction(v string) {
	o.Action = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationResourcePermission) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationResourcePermission) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationResourcePermission) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ApplicationResourcePermission) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ApplicationResourcePermission) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ApplicationResourcePermission) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationResourcePermission) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationResourcePermission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationResourcePermission) SetId(v string) {
	o.Id = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ApplicationResourcePermission) GetResource() ApplicationResourcePermissionResource {
	if o == nil || IsNil(o.Resource) {
		var ret ApplicationResourcePermissionResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourcePermission) GetResourceOk() (*ApplicationResourcePermissionResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ApplicationResourcePermission) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ApplicationResourcePermissionResource and assigns it to the Resource field.
func (o *ApplicationResourcePermission) SetResource(v ApplicationResourcePermissionResource) {
	o.Resource = &v
}

func (o ApplicationResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResourcePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

func (o *ApplicationResourcePermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varApplicationResourcePermission := _ApplicationResourcePermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationResourcePermission)

	if err != nil {
		return err
	}

	*o = ApplicationResourcePermission(varApplicationResourcePermission)

	return err
}

type NullableApplicationResourcePermission struct {
	value *ApplicationResourcePermission
	isSet bool
}

func (v NullableApplicationResourcePermission) Get() *ApplicationResourcePermission {
	return v.value
}

func (v *NullableApplicationResourcePermission) Set(val *ApplicationResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourcePermission(val *ApplicationResourcePermission) *NullableApplicationResourcePermission {
	return &NullableApplicationResourcePermission{value: val, isSet: true}
}

func (v NullableApplicationResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


