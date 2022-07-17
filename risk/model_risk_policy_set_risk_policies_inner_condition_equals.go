/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// RiskPolicySetRiskPoliciesInnerConditionEquals - struct for RiskPolicySetRiskPoliciesInnerConditionEquals
type RiskPolicySetRiskPoliciesInnerConditionEquals struct {
	Bool *bool
	String *string
}

// boolAsRiskPolicySetRiskPoliciesInnerConditionEquals is a convenience function that returns bool wrapped in RiskPolicySetRiskPoliciesInnerConditionEquals
func BoolAsRiskPolicySetRiskPoliciesInnerConditionEquals(v *bool) RiskPolicySetRiskPoliciesInnerConditionEquals {
	return RiskPolicySetRiskPoliciesInnerConditionEquals{
		Bool: v,
	}
}

// stringAsRiskPolicySetRiskPoliciesInnerConditionEquals is a convenience function that returns string wrapped in RiskPolicySetRiskPoliciesInnerConditionEquals
func StringAsRiskPolicySetRiskPoliciesInnerConditionEquals(v *string) RiskPolicySetRiskPoliciesInnerConditionEquals {
	return RiskPolicySetRiskPoliciesInnerConditionEquals{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskPolicySetRiskPoliciesInnerConditionEquals) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(RiskPolicySetRiskPoliciesInnerConditionEquals)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(RiskPolicySetRiskPoliciesInnerConditionEquals)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskPolicySetRiskPoliciesInnerConditionEquals) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskPolicySetRiskPoliciesInnerConditionEquals) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRiskPolicySetRiskPoliciesInnerConditionEquals struct {
	value *RiskPolicySetRiskPoliciesInnerConditionEquals
	isSet bool
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionEquals) Get() *RiskPolicySetRiskPoliciesInnerConditionEquals {
	return v.value
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionEquals) Set(val *RiskPolicySetRiskPoliciesInnerConditionEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetRiskPoliciesInnerConditionEquals(val *RiskPolicySetRiskPoliciesInnerConditionEquals) *NullableRiskPolicySetRiskPoliciesInnerConditionEquals {
	return &NullableRiskPolicySetRiskPoliciesInnerConditionEquals{value: val, isSet: true}
}

func (v NullableRiskPolicySetRiskPoliciesInnerConditionEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetRiskPoliciesInnerConditionEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


