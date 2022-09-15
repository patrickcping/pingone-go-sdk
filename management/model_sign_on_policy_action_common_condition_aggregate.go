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

// SignOnPolicyActionCommonConditionAggregate - struct for SignOnPolicyActionCommonConditionAggregate
type SignOnPolicyActionCommonConditionAggregate struct {
	SignOnPolicyActionCommonConditionAnonymousNetwork *SignOnPolicyActionCommonConditionAnonymousNetwork
	SignOnPolicyActionCommonConditionEquals           *SignOnPolicyActionCommonConditionEquals
	SignOnPolicyActionCommonConditionGeovelocity      *SignOnPolicyActionCommonConditionGeovelocity
	SignOnPolicyActionCommonConditionGreater          *SignOnPolicyActionCommonConditionGreater
	SignOnPolicyActionCommonConditionIPRange          *SignOnPolicyActionCommonConditionIPRange
	SignOnPolicyActionCommonConditionIPRisk           *SignOnPolicyActionCommonConditionIPRisk
}

// SignOnPolicyActionCommonConditionAnonymousNetworkAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionAnonymousNetwork wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionAnonymousNetworkAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionAnonymousNetwork) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionAnonymousNetwork: v,
	}
}

// SignOnPolicyActionCommonConditionEqualsAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionEquals wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionEqualsAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionEquals) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionEquals: v,
	}
}

// SignOnPolicyActionCommonConditionGeovelocityAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionGeovelocity wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionGeovelocityAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionGeovelocity) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionGeovelocity: v,
	}
}

// SignOnPolicyActionCommonConditionGreaterAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionGreater wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionGreaterAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionGreater) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionGreater: v,
	}
}

// SignOnPolicyActionCommonConditionIPRangeAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionIPRange wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionIPRangeAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionIPRange) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionIPRange: v,
	}
}

// SignOnPolicyActionCommonConditionIPRiskAsSignOnPolicyActionCommonConditionAggregate is a convenience function that returns SignOnPolicyActionCommonConditionIPRisk wrapped in SignOnPolicyActionCommonConditionAggregate
func SignOnPolicyActionCommonConditionIPRiskAsSignOnPolicyActionCommonConditionAggregate(v *SignOnPolicyActionCommonConditionIPRisk) SignOnPolicyActionCommonConditionAggregate {
	return SignOnPolicyActionCommonConditionAggregate{
		SignOnPolicyActionCommonConditionIPRisk: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SignOnPolicyActionCommonConditionAggregate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SignOnPolicyActionCommonConditionAnonymousNetwork
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionAnonymousNetwork)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionAnonymousNetwork, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionAnonymousNetwork)
		if string(jsonSignOnPolicyActionCommonConditionAnonymousNetwork) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionAnonymousNetwork = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionAnonymousNetwork.GetAnonymousNetworkOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionAnonymousNetwork = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionAnonymousNetwork = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionEquals
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionEquals)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionEquals, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionEquals)
		if string(jsonSignOnPolicyActionCommonConditionEquals) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionEquals = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionEquals.GetEqualsOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionEquals = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionEquals = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionGeovelocity
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionGeovelocity)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionGeovelocity, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionGeovelocity)
		if string(jsonSignOnPolicyActionCommonConditionGeovelocity) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionGeovelocity = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionGeovelocity.GetGeoVelocityOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionGeovelocity = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionGeovelocity = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionGreater
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionGreater)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionGreater, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionGreater)
		if string(jsonSignOnPolicyActionCommonConditionGreater) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionGreater = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionGreater.GetSecondsSinceOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionGreater = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionGreater = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionIPRange
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionIPRange)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionIPRange, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionIPRange)
		if string(jsonSignOnPolicyActionCommonConditionIPRange) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionIPRange = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionIPRange.GetIpRangeOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionIPRange = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionIPRange = nil
	}

	// try to unmarshal data into SignOnPolicyActionCommonConditionIPRisk
	err = json.Unmarshal(data, &dst.SignOnPolicyActionCommonConditionIPRisk)
	if err == nil {
		jsonSignOnPolicyActionCommonConditionIPRisk, _ := json.Marshal(dst.SignOnPolicyActionCommonConditionIPRisk)
		if string(jsonSignOnPolicyActionCommonConditionIPRisk) == "{}" { // empty struct
			dst.SignOnPolicyActionCommonConditionIPRisk = nil
		} else {
			if _, ok := dst.SignOnPolicyActionCommonConditionIPRisk.GetIpRiskOk(); ok {
				match++
			} else {
				dst.SignOnPolicyActionCommonConditionIPRisk = nil
			}
		}
	} else {
		dst.SignOnPolicyActionCommonConditionIPRisk = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SignOnPolicyActionCommonConditionAnonymousNetwork = nil
		dst.SignOnPolicyActionCommonConditionEquals = nil
		dst.SignOnPolicyActionCommonConditionGeovelocity = nil
		dst.SignOnPolicyActionCommonConditionGreater = nil
		dst.SignOnPolicyActionCommonConditionIPRange = nil
		dst.SignOnPolicyActionCommonConditionIPRisk = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SignOnPolicyActionCommonConditionAggregate)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SignOnPolicyActionCommonConditionAggregate)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SignOnPolicyActionCommonConditionAggregate) MarshalJSON() ([]byte, error) {
	if src.SignOnPolicyActionCommonConditionAnonymousNetwork != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionAnonymousNetwork)
	}

	if src.SignOnPolicyActionCommonConditionEquals != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionEquals)
	}

	if src.SignOnPolicyActionCommonConditionGeovelocity != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionGeovelocity)
	}

	if src.SignOnPolicyActionCommonConditionGreater != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionGreater)
	}

	if src.SignOnPolicyActionCommonConditionIPRange != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionIPRange)
	}

	if src.SignOnPolicyActionCommonConditionIPRisk != nil {
		return json.Marshal(&src.SignOnPolicyActionCommonConditionIPRisk)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SignOnPolicyActionCommonConditionAggregate) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SignOnPolicyActionCommonConditionAnonymousNetwork != nil {
		return obj.SignOnPolicyActionCommonConditionAnonymousNetwork
	}

	if obj.SignOnPolicyActionCommonConditionEquals != nil {
		return obj.SignOnPolicyActionCommonConditionEquals
	}

	if obj.SignOnPolicyActionCommonConditionGeovelocity != nil {
		return obj.SignOnPolicyActionCommonConditionGeovelocity
	}

	if obj.SignOnPolicyActionCommonConditionGreater != nil {
		return obj.SignOnPolicyActionCommonConditionGreater
	}

	if obj.SignOnPolicyActionCommonConditionIPRange != nil {
		return obj.SignOnPolicyActionCommonConditionIPRange
	}

	if obj.SignOnPolicyActionCommonConditionIPRisk != nil {
		return obj.SignOnPolicyActionCommonConditionIPRisk
	}

	// all schemas are nil
	return nil
}

type NullableSignOnPolicyActionCommonConditionAggregate struct {
	value *SignOnPolicyActionCommonConditionAggregate
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionAggregate) Get() *SignOnPolicyActionCommonConditionAggregate {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionAggregate) Set(val *SignOnPolicyActionCommonConditionAggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionAggregate(val *SignOnPolicyActionCommonConditionAggregate) *NullableSignOnPolicyActionCommonConditionAggregate {
	return &NullableSignOnPolicyActionCommonConditionAggregate{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
