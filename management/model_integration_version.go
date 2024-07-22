/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// IntegrationVersion - struct for IntegrationVersion
type IntegrationVersion struct {
	IntegrationVersionIntegrationKit *IntegrationVersionIntegrationKit
	IntegrationVersionSAML *IntegrationVersionSAML
}

// IntegrationVersionIntegrationKitAsIntegrationVersion is a convenience function that returns IntegrationVersionIntegrationKit wrapped in IntegrationVersion
func IntegrationVersionIntegrationKitAsIntegrationVersion(v *IntegrationVersionIntegrationKit) IntegrationVersion {
	return IntegrationVersion{
		IntegrationVersionIntegrationKit: v,
	}
}

// IntegrationVersionSAMLAsIntegrationVersion is a convenience function that returns IntegrationVersionSAML wrapped in IntegrationVersion
func IntegrationVersionSAMLAsIntegrationVersion(v *IntegrationVersionSAML) IntegrationVersion {
	return IntegrationVersion{
		IntegrationVersionSAML: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IntegrationVersion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntegrationVersionIntegrationKit
	err = newStrictDecoder(data).Decode(&dst.IntegrationVersionIntegrationKit)
	if err == nil {
		jsonIntegrationVersionIntegrationKit, _ := json.Marshal(dst.IntegrationVersionIntegrationKit)
		if string(jsonIntegrationVersionIntegrationKit) == "{}" { // empty struct
			dst.IntegrationVersionIntegrationKit = nil
		} else {
			match++
		}
	} else {
		dst.IntegrationVersionIntegrationKit = nil
	}

	// try to unmarshal data into IntegrationVersionSAML
	err = newStrictDecoder(data).Decode(&dst.IntegrationVersionSAML)
	if err == nil {
		jsonIntegrationVersionSAML, _ := json.Marshal(dst.IntegrationVersionSAML)
		if string(jsonIntegrationVersionSAML) == "{}" { // empty struct
			dst.IntegrationVersionSAML = nil
		} else {
			match++
		}
	} else {
		dst.IntegrationVersionSAML = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IntegrationVersionIntegrationKit = nil
		dst.IntegrationVersionSAML = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IntegrationVersion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IntegrationVersion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IntegrationVersion) MarshalJSON() ([]byte, error) {
	if src.IntegrationVersionIntegrationKit != nil {
		return json.Marshal(&src.IntegrationVersionIntegrationKit)
	}

	if src.IntegrationVersionSAML != nil {
		return json.Marshal(&src.IntegrationVersionSAML)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IntegrationVersion) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IntegrationVersionIntegrationKit != nil {
		return obj.IntegrationVersionIntegrationKit
	}

	if obj.IntegrationVersionSAML != nil {
		return obj.IntegrationVersionSAML
	}

	// all schemas are nil
	return nil
}

type NullableIntegrationVersion struct {
	value *IntegrationVersion
	isSet bool
}

func (v NullableIntegrationVersion) Get() *IntegrationVersion {
	return v.value
}

func (v *NullableIntegrationVersion) Set(val *IntegrationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersion(val *IntegrationVersion) *NullableIntegrationVersion {
	return &NullableIntegrationVersion{value: val, isSet: true}
}

func (v NullableIntegrationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


