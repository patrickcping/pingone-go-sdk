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

// checks if the SignOnPolicyActionLoginAllOfNewUserProvisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionLoginAllOfNewUserProvisioning{}

// SignOnPolicyActionLoginAllOfNewUserProvisioning Enables user entries existing outside of PingOne to be provisioned in PingDirectory during login, using an external integration solution (such as a Gateway).
type SignOnPolicyActionLoginAllOfNewUserProvisioning struct {
	// Allows a set of preconfigured gateways or `userType` pairs that are specified in the [Gateway Management](https://apidocs.pingidentity.com/pingone/platform/v1/api/#gateway-management) schema to determine how to find and migrate user entries existing in an external directory.
	Gateways []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways `json:"gateways"`
}

// NewSignOnPolicyActionLoginAllOfNewUserProvisioning instantiates a new SignOnPolicyActionLoginAllOfNewUserProvisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOfNewUserProvisioning(gateways []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways) *SignOnPolicyActionLoginAllOfNewUserProvisioning {
	this := SignOnPolicyActionLoginAllOfNewUserProvisioning{}
	this.Gateways = gateways
	return &this
}

// NewSignOnPolicyActionLoginAllOfNewUserProvisioningWithDefaults instantiates a new SignOnPolicyActionLoginAllOfNewUserProvisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfNewUserProvisioningWithDefaults() *SignOnPolicyActionLoginAllOfNewUserProvisioning {
	this := SignOnPolicyActionLoginAllOfNewUserProvisioning{}
	return &this
}

// GetGateways returns the Gateways field value
func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) GetGateways() []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways {
	if o == nil {
		var ret []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways
		return ret
	}

	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) GetGatewaysOk() ([]SignOnPolicyActionLoginAllOfNewUserProvisioningGateways, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateways, true
}

// SetGateways sets field value
func (o *SignOnPolicyActionLoginAllOfNewUserProvisioning) SetGateways(v []SignOnPolicyActionLoginAllOfNewUserProvisioningGateways) {
	o.Gateways = v
}

func (o SignOnPolicyActionLoginAllOfNewUserProvisioning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionLoginAllOfNewUserProvisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gateways"] = o.Gateways
	return toSerialize, nil
}

type NullableSignOnPolicyActionLoginAllOfNewUserProvisioning struct {
	value *SignOnPolicyActionLoginAllOfNewUserProvisioning
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) Get() *SignOnPolicyActionLoginAllOfNewUserProvisioning {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) Set(val *SignOnPolicyActionLoginAllOfNewUserProvisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOfNewUserProvisioning(val *SignOnPolicyActionLoginAllOfNewUserProvisioning) *NullableSignOnPolicyActionLoginAllOfNewUserProvisioning {
	return &NullableSignOnPolicyActionLoginAllOfNewUserProvisioning{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOfNewUserProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


