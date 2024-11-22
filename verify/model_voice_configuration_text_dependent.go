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

// checks if the VoiceConfigurationTextDependent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceConfigurationTextDependent{}

// VoiceConfigurationTextDependent struct for VoiceConfigurationTextDependent
type VoiceConfigurationTextDependent struct {
	Samples int32                                 `json:"samples"`
	Phrase  VoiceConfigurationTextDependentPhrase `json:"phrase"`
}

type _VoiceConfigurationTextDependent VoiceConfigurationTextDependent

// NewVoiceConfigurationTextDependent instantiates a new VoiceConfigurationTextDependent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceConfigurationTextDependent(samples int32, phrase VoiceConfigurationTextDependentPhrase) *VoiceConfigurationTextDependent {
	this := VoiceConfigurationTextDependent{}
	this.Samples = samples
	this.Phrase = phrase
	return &this
}

// NewVoiceConfigurationTextDependentWithDefaults instantiates a new VoiceConfigurationTextDependent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceConfigurationTextDependentWithDefaults() *VoiceConfigurationTextDependent {
	this := VoiceConfigurationTextDependent{}
	return &this
}

// GetSamples returns the Samples field value
func (o *VoiceConfigurationTextDependent) GetSamples() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *VoiceConfigurationTextDependent) GetSamplesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Samples, true
}

// SetSamples sets field value
func (o *VoiceConfigurationTextDependent) SetSamples(v int32) {
	o.Samples = v
}

// GetPhrase returns the Phrase field value
func (o *VoiceConfigurationTextDependent) GetPhrase() VoiceConfigurationTextDependentPhrase {
	if o == nil {
		var ret VoiceConfigurationTextDependentPhrase
		return ret
	}

	return o.Phrase
}

// GetPhraseOk returns a tuple with the Phrase field value
// and a boolean to check if the value has been set.
func (o *VoiceConfigurationTextDependent) GetPhraseOk() (*VoiceConfigurationTextDependentPhrase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phrase, true
}

// SetPhrase sets field value
func (o *VoiceConfigurationTextDependent) SetPhrase(v VoiceConfigurationTextDependentPhrase) {
	o.Phrase = v
}

func (o VoiceConfigurationTextDependent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceConfigurationTextDependent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["samples"] = o.Samples
	toSerialize["phrase"] = o.Phrase
	return toSerialize, nil
}

func (o *VoiceConfigurationTextDependent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"samples",
		"phrase",
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

	varVoiceConfigurationTextDependent := _VoiceConfigurationTextDependent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVoiceConfigurationTextDependent)

	if err != nil {
		return err
	}

	*o = VoiceConfigurationTextDependent(varVoiceConfigurationTextDependent)

	return err
}

type NullableVoiceConfigurationTextDependent struct {
	value *VoiceConfigurationTextDependent
	isSet bool
}

func (v NullableVoiceConfigurationTextDependent) Get() *VoiceConfigurationTextDependent {
	return v.value
}

func (v *NullableVoiceConfigurationTextDependent) Set(val *VoiceConfigurationTextDependent) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceConfigurationTextDependent) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceConfigurationTextDependent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceConfigurationTextDependent(val *VoiceConfigurationTextDependent) *NullableVoiceConfigurationTextDependent {
	return &NullableVoiceConfigurationTextDependent{value: val, isSet: true}
}

func (v NullableVoiceConfigurationTextDependent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceConfigurationTextDependent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
