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

// AuthorizeEditorDataRulesRuleDTOCondition - struct for AuthorizeEditorDataRulesRuleDTOCondition
type AuthorizeEditorDataRulesRuleDTOCondition struct {
	AuthorizeEditorDataConditionsAndConditionDTO *AuthorizeEditorDataConditionsAndConditionDTO
	AuthorizeEditorDataConditionsComparisonConditionDTO *AuthorizeEditorDataConditionsComparisonConditionDTO
	AuthorizeEditorDataConditionsEmptyConditionDTO *AuthorizeEditorDataConditionsEmptyConditionDTO
	AuthorizeEditorDataConditionsNotConditionDTO *AuthorizeEditorDataConditionsNotConditionDTO
	AuthorizeEditorDataConditionsOrConditionDTO *AuthorizeEditorDataConditionsOrConditionDTO
	AuthorizeEditorDataConditionsReferenceConditionDTO *AuthorizeEditorDataConditionsReferenceConditionDTO
}

// AuthorizeEditorDataConditionsAndConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsAndConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsAndConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsAndConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsAndConditionDTO: v,
	}
}

// AuthorizeEditorDataConditionsComparisonConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsComparisonConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsComparisonConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsComparisonConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsComparisonConditionDTO: v,
	}
}

// AuthorizeEditorDataConditionsEmptyConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsEmptyConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsEmptyConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsEmptyConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsEmptyConditionDTO: v,
	}
}

// AuthorizeEditorDataConditionsNotConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsNotConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsNotConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsNotConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsNotConditionDTO: v,
	}
}

// AuthorizeEditorDataConditionsOrConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsOrConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsOrConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsOrConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsOrConditionDTO: v,
	}
}

// AuthorizeEditorDataConditionsReferenceConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition is a convenience function that returns AuthorizeEditorDataConditionsReferenceConditionDTO wrapped in AuthorizeEditorDataRulesRuleDTOCondition
func AuthorizeEditorDataConditionsReferenceConditionDTOAsAuthorizeEditorDataRulesRuleDTOCondition(v *AuthorizeEditorDataConditionsReferenceConditionDTO) AuthorizeEditorDataRulesRuleDTOCondition {
	return AuthorizeEditorDataRulesRuleDTOCondition{
		AuthorizeEditorDataConditionsReferenceConditionDTO: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthorizeEditorDataRulesRuleDTOCondition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AuthorizeEditorDataConditionsAndConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsAndConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsAndConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsAndConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsAndConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsAndConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsAndConditionDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsComparisonConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsComparisonConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsComparisonConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsComparisonConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsComparisonConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsComparisonConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsComparisonConditionDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsEmptyConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsEmptyConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsEmptyConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsEmptyConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsEmptyConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsEmptyConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsEmptyConditionDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsNotConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsNotConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsNotConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsNotConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsNotConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsNotConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsNotConditionDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsOrConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsOrConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsOrConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsOrConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsOrConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsOrConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsOrConditionDTO = nil
	}

	// try to unmarshal data into AuthorizeEditorDataConditionsReferenceConditionDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataConditionsReferenceConditionDTO)
	if err == nil {
		jsonAuthorizeEditorDataConditionsReferenceConditionDTO, _ := json.Marshal(dst.AuthorizeEditorDataConditionsReferenceConditionDTO)
		if string(jsonAuthorizeEditorDataConditionsReferenceConditionDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataConditionsReferenceConditionDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataConditionsReferenceConditionDTO = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AuthorizeEditorDataConditionsAndConditionDTO = nil
		dst.AuthorizeEditorDataConditionsComparisonConditionDTO = nil
		dst.AuthorizeEditorDataConditionsEmptyConditionDTO = nil
		dst.AuthorizeEditorDataConditionsNotConditionDTO = nil
		dst.AuthorizeEditorDataConditionsOrConditionDTO = nil
		dst.AuthorizeEditorDataConditionsReferenceConditionDTO = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AuthorizeEditorDataRulesRuleDTOCondition)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AuthorizeEditorDataRulesRuleDTOCondition)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthorizeEditorDataRulesRuleDTOCondition) MarshalJSON() ([]byte, error) {
	if src.AuthorizeEditorDataConditionsAndConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsAndConditionDTO)
	}

	if src.AuthorizeEditorDataConditionsComparisonConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsComparisonConditionDTO)
	}

	if src.AuthorizeEditorDataConditionsEmptyConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsEmptyConditionDTO)
	}

	if src.AuthorizeEditorDataConditionsNotConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsNotConditionDTO)
	}

	if src.AuthorizeEditorDataConditionsOrConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsOrConditionDTO)
	}

	if src.AuthorizeEditorDataConditionsReferenceConditionDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataConditionsReferenceConditionDTO)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AuthorizeEditorDataRulesRuleDTOCondition) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthorizeEditorDataConditionsAndConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsAndConditionDTO
	}

	if obj.AuthorizeEditorDataConditionsComparisonConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsComparisonConditionDTO
	}

	if obj.AuthorizeEditorDataConditionsEmptyConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsEmptyConditionDTO
	}

	if obj.AuthorizeEditorDataConditionsNotConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsNotConditionDTO
	}

	if obj.AuthorizeEditorDataConditionsOrConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsOrConditionDTO
	}

	if obj.AuthorizeEditorDataConditionsReferenceConditionDTO != nil {
		return obj.AuthorizeEditorDataConditionsReferenceConditionDTO
	}

	// all schemas are nil
	return nil
}

type NullableAuthorizeEditorDataRulesRuleDTOCondition struct {
	value *AuthorizeEditorDataRulesRuleDTOCondition
	isSet bool
}

func (v NullableAuthorizeEditorDataRulesRuleDTOCondition) Get() *AuthorizeEditorDataRulesRuleDTOCondition {
	return v.value
}

func (v *NullableAuthorizeEditorDataRulesRuleDTOCondition) Set(val *AuthorizeEditorDataRulesRuleDTOCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRulesRuleDTOCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRulesRuleDTOCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRulesRuleDTOCondition(val *AuthorizeEditorDataRulesRuleDTOCondition) *NullableAuthorizeEditorDataRulesRuleDTOCondition {
	return &NullableAuthorizeEditorDataRulesRuleDTOCondition{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRulesRuleDTOCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRulesRuleDTOCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


