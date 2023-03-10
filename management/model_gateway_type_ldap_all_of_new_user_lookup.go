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

// checks if the GatewayTypeLDAPAllOfNewUserLookup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTypeLDAPAllOfNewUserLookup{}

// GatewayTypeLDAPAllOfNewUserLookup The configurations for initially authenticating new users who will be migrated to PingOne. Note If there are multiple users having the same user name, only the first user processed is provisioned.
type GatewayTypeLDAPAllOfNewUserLookup struct {
	// A list of objects supplying a mapping of PingOne attributes to external LDAP attributes. One of the entries must be a mapping for \"username`. This is required for the PingOne user schema.
	AttributeMappings []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings `json:"attributeMappings"`
	// The LDAP user search filter to use to match users against the entered user identifier at login. For example, (((uid=${identifier})(mail=${identifier})). Alternatively, this can be a search against the user directory.
	LdapFilterPattern string `json:"ldapFilterPattern"`
	Population GatewayTypeLDAPAllOfNewUserLookupPopulation `json:"population"`
}

// NewGatewayTypeLDAPAllOfNewUserLookup instantiates a new GatewayTypeLDAPAllOfNewUserLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTypeLDAPAllOfNewUserLookup(attributeMappings []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings, ldapFilterPattern string, population GatewayTypeLDAPAllOfNewUserLookupPopulation) *GatewayTypeLDAPAllOfNewUserLookup {
	this := GatewayTypeLDAPAllOfNewUserLookup{}
	this.AttributeMappings = attributeMappings
	this.LdapFilterPattern = ldapFilterPattern
	this.Population = population
	return &this
}

// NewGatewayTypeLDAPAllOfNewUserLookupWithDefaults instantiates a new GatewayTypeLDAPAllOfNewUserLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTypeLDAPAllOfNewUserLookupWithDefaults() *GatewayTypeLDAPAllOfNewUserLookup {
	this := GatewayTypeLDAPAllOfNewUserLookup{}
	return &this
}

// GetAttributeMappings returns the AttributeMappings field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetAttributeMappings() []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings {
	if o == nil {
		var ret []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings
		return ret
	}

	return o.AttributeMappings
}

// GetAttributeMappingsOk returns a tuple with the AttributeMappings field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetAttributeMappingsOk() ([]GatewayTypeLDAPAllOfNewUserLookupAttributeMappings, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeMappings, true
}

// SetAttributeMappings sets field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) SetAttributeMappings(v []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings) {
	o.AttributeMappings = v
}

// GetLdapFilterPattern returns the LdapFilterPattern field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetLdapFilterPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapFilterPattern
}

// GetLdapFilterPatternOk returns a tuple with the LdapFilterPattern field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetLdapFilterPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapFilterPattern, true
}

// SetLdapFilterPattern sets field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) SetLdapFilterPattern(v string) {
	o.LdapFilterPattern = v
}

// GetPopulation returns the Population field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetPopulation() GatewayTypeLDAPAllOfNewUserLookupPopulation {
	if o == nil {
		var ret GatewayTypeLDAPAllOfNewUserLookupPopulation
		return ret
	}

	return o.Population
}

// GetPopulationOk returns a tuple with the Population field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfNewUserLookup) GetPopulationOk() (*GatewayTypeLDAPAllOfNewUserLookupPopulation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Population, true
}

// SetPopulation sets field value
func (o *GatewayTypeLDAPAllOfNewUserLookup) SetPopulation(v GatewayTypeLDAPAllOfNewUserLookupPopulation) {
	o.Population = v
}

func (o GatewayTypeLDAPAllOfNewUserLookup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTypeLDAPAllOfNewUserLookup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeMappings"] = o.AttributeMappings
	toSerialize["ldapFilterPattern"] = o.LdapFilterPattern
	toSerialize["population"] = o.Population
	return toSerialize, nil
}

type NullableGatewayTypeLDAPAllOfNewUserLookup struct {
	value *GatewayTypeLDAPAllOfNewUserLookup
	isSet bool
}

func (v NullableGatewayTypeLDAPAllOfNewUserLookup) Get() *GatewayTypeLDAPAllOfNewUserLookup {
	return v.value
}

func (v *NullableGatewayTypeLDAPAllOfNewUserLookup) Set(val *GatewayTypeLDAPAllOfNewUserLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTypeLDAPAllOfNewUserLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTypeLDAPAllOfNewUserLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTypeLDAPAllOfNewUserLookup(val *GatewayTypeLDAPAllOfNewUserLookup) *NullableGatewayTypeLDAPAllOfNewUserLookup {
	return &NullableGatewayTypeLDAPAllOfNewUserLookup{value: val, isSet: true}
}

func (v NullableGatewayTypeLDAPAllOfNewUserLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTypeLDAPAllOfNewUserLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


