/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EntityArrayEmbeddedApplicationsInner struct for EntityArrayEmbeddedApplicationsInner
type EntityArrayEmbeddedApplicationsInner struct {
	ApplicationOIDC *ApplicationOIDC
	ApplicationSAML *ApplicationSAML
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntityArrayEmbeddedApplicationsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ApplicationOIDC
	err = json.Unmarshal(data, &dst.ApplicationOIDC);
	if err == nil {
		jsonApplicationOIDC, _ := json.Marshal(dst.ApplicationOIDC)
		if string(jsonApplicationOIDC) == "{}" { // empty struct
			dst.ApplicationOIDC = nil
		} else {
			return nil // data stored in dst.ApplicationOIDC, return on the first match
		}
	} else {
		dst.ApplicationOIDC = nil
	}

	// try to unmarshal JSON data into ApplicationSAML
	err = json.Unmarshal(data, &dst.ApplicationSAML);
	if err == nil {
		jsonApplicationSAML, _ := json.Marshal(dst.ApplicationSAML)
		if string(jsonApplicationSAML) == "{}" { // empty struct
			dst.ApplicationSAML = nil
		} else {
			return nil // data stored in dst.ApplicationSAML, return on the first match
		}
	} else {
		dst.ApplicationSAML = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedApplicationsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntityArrayEmbeddedApplicationsInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntityArrayEmbeddedApplicationsInner struct {
	value *EntityArrayEmbeddedApplicationsInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedApplicationsInner) Get() *EntityArrayEmbeddedApplicationsInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) Set(val *EntityArrayEmbeddedApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedApplicationsInner(val *EntityArrayEmbeddedApplicationsInner) *NullableEntityArrayEmbeddedApplicationsInner {
	return &NullableEntityArrayEmbeddedApplicationsInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


