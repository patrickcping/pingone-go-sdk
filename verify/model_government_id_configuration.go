/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GovernmentIdConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernmentIdConfiguration{}

// GovernmentIdConfiguration struct for GovernmentIdConfiguration
type GovernmentIdConfiguration struct {
	// Indicates whether verification should fail if the ID is expired.
	FailExpiredId *bool `json:"failExpiredId,omitempty"`
	// Determines whether document authentication is automated, manual, or a combination of both where manual authentication is performed if automated inspection fails. Can be AUTOMATIC, MANUAL, or STEP_UP.
	InspectionType *EnumInspectionType                `json:"inspectionType,omitempty"`
	Provider       *GovernmentIdConfigurationProvider `json:"provider,omitempty"`
	Retry          *ObjectRetry                       `json:"retry,omitempty"`
	// Controls if Government ID verification is REQUIRED, OPTIONAL, or DISABLED.
	Verify EnumVerify `json:"verify"`
}

type _GovernmentIdConfiguration GovernmentIdConfiguration

// NewGovernmentIdConfiguration instantiates a new GovernmentIdConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernmentIdConfiguration(verify EnumVerify) *GovernmentIdConfiguration {
	this := GovernmentIdConfiguration{}
	this.Verify = verify
	return &this
}

// NewGovernmentIdConfigurationWithDefaults instantiates a new GovernmentIdConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernmentIdConfigurationWithDefaults() *GovernmentIdConfiguration {
	this := GovernmentIdConfiguration{}
	return &this
}

// GetFailExpiredId returns the FailExpiredId field value if set, zero value otherwise.
func (o *GovernmentIdConfiguration) GetFailExpiredId() bool {
	if o == nil || IsNil(o.FailExpiredId) {
		var ret bool
		return ret
	}
	return *o.FailExpiredId
}

// GetFailExpiredIdOk returns a tuple with the FailExpiredId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetFailExpiredIdOk() (*bool, bool) {
	if o == nil || IsNil(o.FailExpiredId) {
		return nil, false
	}
	return o.FailExpiredId, true
}

// HasFailExpiredId returns a boolean if a field has been set.
func (o *GovernmentIdConfiguration) HasFailExpiredId() bool {
	if o != nil && !IsNil(o.FailExpiredId) {
		return true
	}

	return false
}

// SetFailExpiredId gets a reference to the given bool and assigns it to the FailExpiredId field.
func (o *GovernmentIdConfiguration) SetFailExpiredId(v bool) {
	o.FailExpiredId = &v
}

// GetInspectionType returns the InspectionType field value if set, zero value otherwise.
func (o *GovernmentIdConfiguration) GetInspectionType() EnumInspectionType {
	if o == nil || IsNil(o.InspectionType) {
		var ret EnumInspectionType
		return ret
	}
	return *o.InspectionType
}

// GetInspectionTypeOk returns a tuple with the InspectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetInspectionTypeOk() (*EnumInspectionType, bool) {
	if o == nil || IsNil(o.InspectionType) {
		return nil, false
	}
	return o.InspectionType, true
}

// HasInspectionType returns a boolean if a field has been set.
func (o *GovernmentIdConfiguration) HasInspectionType() bool {
	if o != nil && !IsNil(o.InspectionType) {
		return true
	}

	return false
}

// SetInspectionType gets a reference to the given EnumInspectionType and assigns it to the InspectionType field.
func (o *GovernmentIdConfiguration) SetInspectionType(v EnumInspectionType) {
	o.InspectionType = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GovernmentIdConfiguration) GetProvider() GovernmentIdConfigurationProvider {
	if o == nil || IsNil(o.Provider) {
		var ret GovernmentIdConfigurationProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetProviderOk() (*GovernmentIdConfigurationProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GovernmentIdConfiguration) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given GovernmentIdConfigurationProvider and assigns it to the Provider field.
func (o *GovernmentIdConfiguration) SetProvider(v GovernmentIdConfigurationProvider) {
	o.Provider = &v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *GovernmentIdConfiguration) GetRetry() ObjectRetry {
	if o == nil || IsNil(o.Retry) {
		var ret ObjectRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetRetryOk() (*ObjectRetry, bool) {
	if o == nil || IsNil(o.Retry) {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *GovernmentIdConfiguration) HasRetry() bool {
	if o != nil && !IsNil(o.Retry) {
		return true
	}

	return false
}

// SetRetry gets a reference to the given ObjectRetry and assigns it to the Retry field.
func (o *GovernmentIdConfiguration) SetRetry(v ObjectRetry) {
	o.Retry = &v
}

// GetVerify returns the Verify field value
func (o *GovernmentIdConfiguration) GetVerify() EnumVerify {
	if o == nil {
		var ret EnumVerify
		return ret
	}

	return o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetVerifyOk() (*EnumVerify, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verify, true
}

// SetVerify sets field value
func (o *GovernmentIdConfiguration) SetVerify(v EnumVerify) {
	o.Verify = v
}

func (o GovernmentIdConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernmentIdConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailExpiredId) {
		toSerialize["failExpiredId"] = o.FailExpiredId
	}
	if !IsNil(o.InspectionType) {
		toSerialize["inspectionType"] = o.InspectionType
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Retry) {
		toSerialize["retry"] = o.Retry
	}
	toSerialize["verify"] = o.Verify
	return toSerialize, nil
}

func (o *GovernmentIdConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verify",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGovernmentIdConfiguration := _GovernmentIdConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGovernmentIdConfiguration)

	if err != nil {
		return err
	}

	*o = GovernmentIdConfiguration(varGovernmentIdConfiguration)

	return err
}

type NullableGovernmentIdConfiguration struct {
	value *GovernmentIdConfiguration
	isSet bool
}

func (v NullableGovernmentIdConfiguration) Get() *GovernmentIdConfiguration {
	return v.value
}

func (v *NullableGovernmentIdConfiguration) Set(val *GovernmentIdConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernmentIdConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernmentIdConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernmentIdConfiguration(val *GovernmentIdConfiguration) *NullableGovernmentIdConfiguration {
	return &NullableGovernmentIdConfiguration{value: val, isSet: true}
}

func (v NullableGovernmentIdConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernmentIdConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
