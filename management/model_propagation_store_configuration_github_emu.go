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

// checks if the PropagationStoreConfigurationGithubEMU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropagationStoreConfigurationGithubEMU{}

// PropagationStoreConfigurationGithubEMU struct for PropagationStoreConfigurationGithubEMU
type PropagationStoreConfigurationGithubEMU struct {
	// Base URL of the target propagation store.
	BASE_URL string `json:"BASE_URL"`
	// Whether or not users are allowed to be created.
	CREATE_USERS *bool `json:"CREATE_USERS,omitempty"`
	// Whether or not users are allowed to be deprovisioned (removed) following action specified in `REMOVE_ACTION`.
	DEPROVISION_USERS *bool `json:"DEPROVISION_USERS,omitempty"`
	// OAuth 2 access token.
	OAUTH_ACCESS_TOKEN string `json:"OAUTH_ACCESS_TOKEN"`
	REMOVE_ACTION *EnumPropagationStoreTypeRemoveActionDisableDelete `json:"REMOVE_ACTION,omitempty"`
	// Whether or not users are allowed to be updated.
	UPDATE_USERS *bool `json:"UPDATE_USERS,omitempty"`
}

// NewPropagationStoreConfigurationGithubEMU instantiates a new PropagationStoreConfigurationGithubEMU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropagationStoreConfigurationGithubEMU(bASEURL string, oAUTHACCESSTOKEN string) *PropagationStoreConfigurationGithubEMU {
	this := PropagationStoreConfigurationGithubEMU{}
	this.BASE_URL = bASEURL
	this.OAUTH_ACCESS_TOKEN = oAUTHACCESSTOKEN
	return &this
}

// NewPropagationStoreConfigurationGithubEMUWithDefaults instantiates a new PropagationStoreConfigurationGithubEMU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropagationStoreConfigurationGithubEMUWithDefaults() *PropagationStoreConfigurationGithubEMU {
	this := PropagationStoreConfigurationGithubEMU{}
	return &this
}

// GetBASE_URL returns the BASE_URL field value
func (o *PropagationStoreConfigurationGithubEMU) GetBASE_URL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BASE_URL
}

// GetBASE_URLOk returns a tuple with the BASE_URL field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetBASE_URLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BASE_URL, true
}

// SetBASE_URL sets field value
func (o *PropagationStoreConfigurationGithubEMU) SetBASE_URL(v string) {
	o.BASE_URL = v
}

// GetCREATE_USERS returns the CREATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationGithubEMU) GetCREATE_USERS() bool {
	if o == nil || IsNil(o.CREATE_USERS) {
		var ret bool
		return ret
	}
	return *o.CREATE_USERS
}

// GetCREATE_USERSOk returns a tuple with the CREATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetCREATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.CREATE_USERS) {
		return nil, false
	}
	return o.CREATE_USERS, true
}

// HasCREATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationGithubEMU) HasCREATE_USERS() bool {
	if o != nil && !IsNil(o.CREATE_USERS) {
		return true
	}

	return false
}

// SetCREATE_USERS gets a reference to the given bool and assigns it to the CREATE_USERS field.
func (o *PropagationStoreConfigurationGithubEMU) SetCREATE_USERS(v bool) {
	o.CREATE_USERS = &v
}

// GetDEPROVISION_USERS returns the DEPROVISION_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationGithubEMU) GetDEPROVISION_USERS() bool {
	if o == nil || IsNil(o.DEPROVISION_USERS) {
		var ret bool
		return ret
	}
	return *o.DEPROVISION_USERS
}

// GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetDEPROVISION_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.DEPROVISION_USERS) {
		return nil, false
	}
	return o.DEPROVISION_USERS, true
}

// HasDEPROVISION_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationGithubEMU) HasDEPROVISION_USERS() bool {
	if o != nil && !IsNil(o.DEPROVISION_USERS) {
		return true
	}

	return false
}

// SetDEPROVISION_USERS gets a reference to the given bool and assigns it to the DEPROVISION_USERS field.
func (o *PropagationStoreConfigurationGithubEMU) SetDEPROVISION_USERS(v bool) {
	o.DEPROVISION_USERS = &v
}

// GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field value
func (o *PropagationStoreConfigurationGithubEMU) GetOAUTH_ACCESS_TOKEN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OAUTH_ACCESS_TOKEN
}

// GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetOAUTH_ACCESS_TOKENOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OAUTH_ACCESS_TOKEN, true
}

// SetOAUTH_ACCESS_TOKEN sets field value
func (o *PropagationStoreConfigurationGithubEMU) SetOAUTH_ACCESS_TOKEN(v string) {
	o.OAUTH_ACCESS_TOKEN = v
}

// GetREMOVE_ACTION returns the REMOVE_ACTION field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationGithubEMU) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		var ret EnumPropagationStoreTypeRemoveActionDisableDelete
		return ret
	}
	return *o.REMOVE_ACTION
}

// GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool) {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		return nil, false
	}
	return o.REMOVE_ACTION, true
}

// HasREMOVE_ACTION returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationGithubEMU) HasREMOVE_ACTION() bool {
	if o != nil && !IsNil(o.REMOVE_ACTION) {
		return true
	}

	return false
}

// SetREMOVE_ACTION gets a reference to the given EnumPropagationStoreTypeRemoveActionDisableDelete and assigns it to the REMOVE_ACTION field.
func (o *PropagationStoreConfigurationGithubEMU) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete) {
	o.REMOVE_ACTION = &v
}

// GetUPDATE_USERS returns the UPDATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationGithubEMU) GetUPDATE_USERS() bool {
	if o == nil || IsNil(o.UPDATE_USERS) {
		var ret bool
		return ret
	}
	return *o.UPDATE_USERS
}

// GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationGithubEMU) GetUPDATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.UPDATE_USERS) {
		return nil, false
	}
	return o.UPDATE_USERS, true
}

// HasUPDATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationGithubEMU) HasUPDATE_USERS() bool {
	if o != nil && !IsNil(o.UPDATE_USERS) {
		return true
	}

	return false
}

// SetUPDATE_USERS gets a reference to the given bool and assigns it to the UPDATE_USERS field.
func (o *PropagationStoreConfigurationGithubEMU) SetUPDATE_USERS(v bool) {
	o.UPDATE_USERS = &v
}

func (o PropagationStoreConfigurationGithubEMU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropagationStoreConfigurationGithubEMU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BASE_URL"] = o.BASE_URL
	if !IsNil(o.CREATE_USERS) {
		toSerialize["CREATE_USERS"] = o.CREATE_USERS
	}
	if !IsNil(o.DEPROVISION_USERS) {
		toSerialize["DEPROVISION_USERS"] = o.DEPROVISION_USERS
	}
	toSerialize["OAUTH_ACCESS_TOKEN"] = o.OAUTH_ACCESS_TOKEN
	if !IsNil(o.REMOVE_ACTION) {
		toSerialize["REMOVE_ACTION"] = o.REMOVE_ACTION
	}
	if !IsNil(o.UPDATE_USERS) {
		toSerialize["UPDATE_USERS"] = o.UPDATE_USERS
	}
	return toSerialize, nil
}

type NullablePropagationStoreConfigurationGithubEMU struct {
	value *PropagationStoreConfigurationGithubEMU
	isSet bool
}

func (v NullablePropagationStoreConfigurationGithubEMU) Get() *PropagationStoreConfigurationGithubEMU {
	return v.value
}

func (v *NullablePropagationStoreConfigurationGithubEMU) Set(val *PropagationStoreConfigurationGithubEMU) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationStoreConfigurationGithubEMU) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationStoreConfigurationGithubEMU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationStoreConfigurationGithubEMU(val *PropagationStoreConfigurationGithubEMU) *NullablePropagationStoreConfigurationGithubEMU {
	return &NullablePropagationStoreConfigurationGithubEMU{value: val, isSet: true}
}

func (v NullablePropagationStoreConfigurationGithubEMU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationStoreConfigurationGithubEMU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


