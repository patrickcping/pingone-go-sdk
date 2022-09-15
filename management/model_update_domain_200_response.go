/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package management

import (
	"encoding/json"
	"fmt"
)

// UpdateDomain200Response - struct for UpdateDomain200Response
type UpdateDomain200Response struct {
	CustomDomain            *CustomDomain
	CustomDomainCertificate *CustomDomainCertificate
}

// CustomDomainAsUpdateDomain200Response is a convenience function that returns CustomDomain wrapped in UpdateDomain200Response
func CustomDomainAsUpdateDomain200Response(v *CustomDomain) UpdateDomain200Response {
	return UpdateDomain200Response{
		CustomDomain: v,
	}
}

// CustomDomainCertificateAsUpdateDomain200Response is a convenience function that returns CustomDomainCertificate wrapped in UpdateDomain200Response
func CustomDomainCertificateAsUpdateDomain200Response(v *CustomDomainCertificate) UpdateDomain200Response {
	return UpdateDomain200Response{
		CustomDomainCertificate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateDomain200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDomain
	err = json.Unmarshal(data, &dst.CustomDomain)
	if err == nil {
		jsonCustomDomain, _ := json.Marshal(dst.CustomDomain)
		if string(jsonCustomDomain) == "{}" { // empty struct
			dst.CustomDomain = nil
		} else {
			if _, ok := dst.CustomDomain.GetIdOk(); ok {
				match++
			} else {
				dst.CustomDomain = nil
			}
		}
	} else {
		dst.CustomDomain = nil
	}

	// try to unmarshal data into CustomDomainCertificate
	err = json.Unmarshal(data, &dst.CustomDomainCertificate)
	if err == nil {
		jsonCustomDomainCertificate, _ := json.Marshal(dst.CustomDomainCertificate)
		if string(jsonCustomDomainCertificate) == "{}" { // empty struct
			dst.CustomDomainCertificate = nil
		} else {
			if _, ok := dst.CustomDomainCertificate.GetExpiresAtOk(); ok {
				match++
			} else {
				dst.CustomDomainCertificate = nil
			}
		}
	} else {
		dst.CustomDomainCertificate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomDomain = nil
		dst.CustomDomainCertificate = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UpdateDomain200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateDomain200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDomain200Response) MarshalJSON() ([]byte, error) {
	if src.CustomDomain != nil {
		return json.Marshal(&src.CustomDomain)
	}

	if src.CustomDomainCertificate != nil {
		return json.Marshal(&src.CustomDomainCertificate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDomain200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomDomain != nil {
		return obj.CustomDomain
	}

	if obj.CustomDomainCertificate != nil {
		return obj.CustomDomainCertificate
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDomain200Response struct {
	value *UpdateDomain200Response
	isSet bool
}

func (v NullableUpdateDomain200Response) Get() *UpdateDomain200Response {
	return v.value
}

func (v *NullableUpdateDomain200Response) Set(val *UpdateDomain200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDomain200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDomain200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDomain200Response(val *UpdateDomain200Response) *NullableUpdateDomain200Response {
	return &NullableUpdateDomain200Response{value: val, isSet: true}
}

func (v NullableUpdateDomain200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDomain200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
