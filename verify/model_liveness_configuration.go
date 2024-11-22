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

// checks if the LivenessConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LivenessConfiguration{}

// LivenessConfiguration struct for LivenessConfiguration
type LivenessConfiguration struct {
	Retry *ObjectRetry `json:"retry,omitempty"`
	// Threshold for successful facial comparison; can be LOW, MEDIUM, or HIGH (for which PingOne Verify uses industry and vendor recommended definitions).
	Threshold EnumThreshold `json:"threshold"`
	// Controls if liveness check is REQUIRED, OPTIONAL, or DISABLED.
	Verify EnumVerify `json:"verify"`
}

type _LivenessConfiguration LivenessConfiguration

// NewLivenessConfiguration instantiates a new LivenessConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLivenessConfiguration(threshold EnumThreshold, verify EnumVerify) *LivenessConfiguration {
	this := LivenessConfiguration{}
	this.Threshold = threshold
	this.Verify = verify
	return &this
}

// NewLivenessConfigurationWithDefaults instantiates a new LivenessConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLivenessConfigurationWithDefaults() *LivenessConfiguration {
	this := LivenessConfiguration{}
	return &this
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *LivenessConfiguration) GetRetry() ObjectRetry {
	if o == nil || IsNil(o.Retry) {
		var ret ObjectRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LivenessConfiguration) GetRetryOk() (*ObjectRetry, bool) {
	if o == nil || IsNil(o.Retry) {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *LivenessConfiguration) HasRetry() bool {
	if o != nil && !IsNil(o.Retry) {
		return true
	}

	return false
}

// SetRetry gets a reference to the given ObjectRetry and assigns it to the Retry field.
func (o *LivenessConfiguration) SetRetry(v ObjectRetry) {
	o.Retry = &v
}

// GetThreshold returns the Threshold field value
func (o *LivenessConfiguration) GetThreshold() EnumThreshold {
	if o == nil {
		var ret EnumThreshold
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *LivenessConfiguration) GetThresholdOk() (*EnumThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *LivenessConfiguration) SetThreshold(v EnumThreshold) {
	o.Threshold = v
}

// GetVerify returns the Verify field value
func (o *LivenessConfiguration) GetVerify() EnumVerify {
	if o == nil {
		var ret EnumVerify
		return ret
	}

	return o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value
// and a boolean to check if the value has been set.
func (o *LivenessConfiguration) GetVerifyOk() (*EnumVerify, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verify, true
}

// SetVerify sets field value
func (o *LivenessConfiguration) SetVerify(v EnumVerify) {
	o.Verify = v
}

func (o LivenessConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LivenessConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Retry) {
		toSerialize["retry"] = o.Retry
	}
	toSerialize["threshold"] = o.Threshold
	toSerialize["verify"] = o.Verify
	return toSerialize, nil
}

func (o *LivenessConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threshold",
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

	varLivenessConfiguration := _LivenessConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLivenessConfiguration)

	if err != nil {
		return err
	}

	*o = LivenessConfiguration(varLivenessConfiguration)

	return err
}

type NullableLivenessConfiguration struct {
	value *LivenessConfiguration
	isSet bool
}

func (v NullableLivenessConfiguration) Get() *LivenessConfiguration {
	return v.value
}

func (v *NullableLivenessConfiguration) Set(val *LivenessConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLivenessConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLivenessConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLivenessConfiguration(val *LivenessConfiguration) *NullableLivenessConfiguration {
	return &NullableLivenessConfiguration{value: val, isSet: true}
}

func (v NullableLivenessConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLivenessConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
