/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RiskPredictorComposite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorComposite{}

// RiskPredictorComposite struct for RiskPredictorComposite
type RiskPredictorComposite struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights.
	Name string `json:"name"`
	// A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details).
	CompactName string            `json:"compactName"`
	Type        EnumPredictorType `json:"type"`
	// A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Indicates whether PingOne Risk is licensed for the environment.
	Licensed *bool `json:"licensed,omitempty"`
	// A boolean to indicate whether the predictor is deletable in the environment.
	Deletable *bool                         `json:"deletable,omitempty"`
	Default   *RiskPredictorCommonDefault   `json:"default,omitempty"`
	Condition *RiskPredictorCommonCondition `json:"condition,omitempty"`
	// Deprecated
	Composition *RiskPredictorCompositeAllOfComposition `json:"composition,omitempty"`
	// Contains the objects that specify the conditions to test and the risk level that should be assigned if the conditions are met. The array can contain a maximum of three elements.
	Compositions []RiskPredictorCompositeAllOfCompositions `json:"compositions"`
}

type _RiskPredictorComposite RiskPredictorComposite

// NewRiskPredictorComposite instantiates a new RiskPredictorComposite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorComposite(name string, compactName string, type_ EnumPredictorType, compositions []RiskPredictorCompositeAllOfCompositions) *RiskPredictorComposite {
	this := RiskPredictorComposite{}
	this.Name = name
	this.CompactName = compactName
	this.Type = type_
	this.Compositions = compositions
	return &this
}

// NewRiskPredictorCompositeWithDefaults instantiates a new RiskPredictorComposite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCompositeWithDefaults() *RiskPredictorComposite {
	this := RiskPredictorComposite{}
	return &this
}

func (o RiskPredictorComposite) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o RiskPredictorComposite) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o RiskPredictorComposite) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o RiskPredictorComposite) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorComposite) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorComposite) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorComposite) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorComposite) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorComposite) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorComposite) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorComposite) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorComposite) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorComposite) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *RiskPredictorComposite) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictorComposite) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPredictorComposite) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorComposite) SetName(v string) {
	o.Name = v
}

// GetCompactName returns the CompactName field value
func (o *RiskPredictorComposite) GetCompactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetCompactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompactName, true
}

// SetCompactName sets field value
func (o *RiskPredictorComposite) SetCompactName(v string) {
	o.CompactName = v
}

// GetType returns the Type field value
func (o *RiskPredictorComposite) GetType() EnumPredictorType {
	if o == nil {
		var ret EnumPredictorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetTypeOk() (*EnumPredictorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPredictorComposite) SetType(v EnumPredictorType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictorComposite) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPredictorComposite) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPredictorComposite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RiskPredictorComposite) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetDeletable() bool {
	if o == nil || IsNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetDeletableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deletable) {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasDeletable() bool {
	if o != nil && !IsNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *RiskPredictorComposite) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetDefault() RiskPredictorCommonDefault {
	if o == nil || IsNil(o.Default) {
		var ret RiskPredictorCommonDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetDefaultOk() (*RiskPredictorCommonDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorCommonDefault and assigns it to the Default field.
func (o *RiskPredictorComposite) SetDefault(v RiskPredictorCommonDefault) {
	o.Default = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RiskPredictorComposite) GetCondition() RiskPredictorCommonCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RiskPredictorCommonCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetConditionOk() (*RiskPredictorCommonCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RiskPredictorCommonCondition and assigns it to the Condition field.
func (o *RiskPredictorComposite) SetCondition(v RiskPredictorCommonCondition) {
	o.Condition = &v
}

// GetComposition returns the Composition field value if set, zero value otherwise.
// Deprecated
func (o *RiskPredictorComposite) GetComposition() RiskPredictorCompositeAllOfComposition {
	if o == nil || IsNil(o.Composition) {
		var ret RiskPredictorCompositeAllOfComposition
		return ret
	}
	return *o.Composition
}

// GetCompositionOk returns a tuple with the Composition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RiskPredictorComposite) GetCompositionOk() (*RiskPredictorCompositeAllOfComposition, bool) {
	if o == nil || IsNil(o.Composition) {
		return nil, false
	}
	return o.Composition, true
}

// HasComposition returns a boolean if a field has been set.
func (o *RiskPredictorComposite) HasComposition() bool {
	if o != nil && !IsNil(o.Composition) {
		return true
	}

	return false
}

// SetComposition gets a reference to the given RiskPredictorCompositeAllOfComposition and assigns it to the Composition field.
// Deprecated
func (o *RiskPredictorComposite) SetComposition(v RiskPredictorCompositeAllOfComposition) {
	o.Composition = &v
}

// GetCompositions returns the Compositions field value
func (o *RiskPredictorComposite) GetCompositions() []RiskPredictorCompositeAllOfCompositions {
	if o == nil {
		var ret []RiskPredictorCompositeAllOfCompositions
		return ret
	}

	return o.Compositions
}

// GetCompositionsOk returns a tuple with the Compositions field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorComposite) GetCompositionsOk() ([]RiskPredictorCompositeAllOfCompositions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Compositions, true
}

// SetCompositions sets field value
func (o *RiskPredictorComposite) SetCompositions(v []RiskPredictorCompositeAllOfCompositions) {
	o.Compositions = v
}

func (o RiskPredictorComposite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorComposite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["compactName"] = o.CompactName
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Licensed) {
		toSerialize["licensed"] = o.Licensed
	}
	if !IsNil(o.Deletable) {
		toSerialize["deletable"] = o.Deletable
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Composition) {
		toSerialize["composition"] = o.Composition
	}
	toSerialize["compositions"] = o.Compositions
	return toSerialize, nil
}

func (o *RiskPredictorComposite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"compactName",
		"type",
		"compositions",
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

	varRiskPredictorComposite := _RiskPredictorComposite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPredictorComposite)

	if err != nil {
		return err
	}

	*o = RiskPredictorComposite(varRiskPredictorComposite)

	return err
}

type NullableRiskPredictorComposite struct {
	value *RiskPredictorComposite
	isSet bool
}

func (v NullableRiskPredictorComposite) Get() *RiskPredictorComposite {
	return v.value
}

func (v *NullableRiskPredictorComposite) Set(val *RiskPredictorComposite) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorComposite) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorComposite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorComposite(val *RiskPredictorComposite) *NullableRiskPredictorComposite {
	return &NullableRiskPredictorComposite{value: val, isSet: true}
}

func (v NullableRiskPredictorComposite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorComposite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
