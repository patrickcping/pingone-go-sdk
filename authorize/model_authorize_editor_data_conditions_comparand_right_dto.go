/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// AuthorizeEditorDataConditionsComparandRightDTO - struct for AuthorizeEditorDataConditionsComparandRightDTO
type AuthorizeEditorDataConditionsComparandRightDTO struct {
	AuthorizeEditorDataConditionsComparandsAttributeComparandDTO *AuthorizeEditorDataConditionsComparandsAttributeComparandDTO
	AuthorizeEditorDataConditionsComparandsConstantComparandDTO *AuthorizeEditorDataConditionsComparandsConstantComparandDTO
}

// AuthorizeEditorDataConditionsComparandsAttributeComparandDTOAsAuthorizeEditorDataConditionsComparandRightDTO is a convenience function that returns AuthorizeEditorDataConditionsComparandsAttributeComparandDTO wrapped in AuthorizeEditorDataConditionsComparandRightDTO
func AuthorizeEditorDataConditionsComparandsAttributeComparandDTOAsAuthorizeEditorDataConditionsComparandRightDTO(v *AuthorizeEditorDataConditionsComparandsAttributeComparandDTO) AuthorizeEditorDataConditionsComparandRightDTO {
	return AuthorizeEditorDataConditionsComparandRightDTO{
		AuthorizeEditorDataConditionsComparandsAttributeComparandDTO: v,
	}
}

// AuthorizeEditorDataConditionsComparandsConstantComparandDTOAsAuthorizeEditorDataConditionsComparandRightDTO is a convenience function that returns AuthorizeEditorDataConditionsComparandsConstantComparandDTO wrapped in AuthorizeEditorDataConditionsComparandRightDTO
func AuthorizeEditorDataConditionsComparandsConstantComparandDTOAsAuthorizeEditorDataConditionsComparandRightDTO(v *AuthorizeEditorDataConditionsComparandsConstantComparandDTO) AuthorizeEditorDataConditionsComparandRightDTO {
	return AuthorizeEditorDataConditionsComparandRightDTO{
		AuthorizeEditorDataConditionsComparandsConstantComparandDTO: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthorizeEditorDataConditionsComparandRightDTO) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AuthorizeEditorDataConditionsComparandsAttributeComparandDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsComparandsAttributeComparandDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO)
		if string(jsonAuthorizeEditorDataConditionsComparandsAttributeComparandDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsComparandsConstantComparandDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsComparandsConstantComparandDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO)
		if string(jsonAuthorizeEditorDataConditionsComparandsConstantComparandDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO = nil
		dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AuthorizeEditorDataConditionsComparandRightDTO)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AuthorizeEditorDataConditionsComparandRightDTO)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthorizeEditorDataConditionsComparandRightDTO) MarshalJSON() ([]byte, error) {
	if src.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO)
	}

	if src.AuthorizeEditorDataConditionsComparandsConstantComparandDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsComparandsConstantComparandDTO)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AuthorizeEditorDataConditionsComparandRightDTO) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO != nil {
		return obj.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO
	}

	if obj.AuthorizeEditorDataConditionsComparandsConstantComparandDTO != nil {
		return obj.AuthorizeEditorDataConditionsComparandsConstantComparandDTO
	}

	// all schemas are nil
	return nil
}

type NullableAuthorizeEditorDataConditionsComparandRightDTO struct {
	value *AuthorizeEditorDataConditionsComparandRightDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataConditionsComparandRightDTO) Get() *AuthorizeEditorDataConditionsComparandRightDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataConditionsComparandRightDTO) Set(val *AuthorizeEditorDataConditionsComparandRightDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataConditionsComparandRightDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataConditionsComparandRightDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataConditionsComparandRightDTO(val *AuthorizeEditorDataConditionsComparandRightDTO) *NullableAuthorizeEditorDataConditionsComparandRightDTO {
	return &NullableAuthorizeEditorDataConditionsComparandRightDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataConditionsComparandRightDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataConditionsComparandRightDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


