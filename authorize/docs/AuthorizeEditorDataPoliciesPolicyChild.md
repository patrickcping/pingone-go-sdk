# AuthorizeEditorDataPoliciesPolicyChild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataPoliciesPolicyChildCommonType**](EnumAuthorizeEditorDataPoliciesPolicyChildCommonType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Statements** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**Value** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**CombiningAlgorithm** | Pointer to [**AuthorizeEditorDataPoliciesCombiningAlgorithmDTO**](AuthorizeEditorDataPoliciesCombiningAlgorithmDTO.md) |  | [optional] 
**Children** | Pointer to [**[]AuthorizeEditorDataPoliciesPolicyChild**](AuthorizeEditorDataPoliciesPolicyChild.md) |  | [optional] 
**RepetitionSettings** | Pointer to [**AuthorizeEditorDataPoliciesRepetitionSettingsDTO**](AuthorizeEditorDataPoliciesRepetitionSettingsDTO.md) |  | [optional] 
**EffectSettings** | Pointer to [**AuthorizeEditorDataRulesEffectSettingsDTO**](AuthorizeEditorDataRulesEffectSettingsDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataPoliciesPolicyChild

`func NewAuthorizeEditorDataPoliciesPolicyChild(type_ EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, ) *AuthorizeEditorDataPoliciesPolicyChild`

NewAuthorizeEditorDataPoliciesPolicyChild instantiates a new AuthorizeEditorDataPoliciesPolicyChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesPolicyChildWithDefaults

`func NewAuthorizeEditorDataPoliciesPolicyChildWithDefaults() *AuthorizeEditorDataPoliciesPolicyChild`

NewAuthorizeEditorDataPoliciesPolicyChildWithDefaults instantiates a new AuthorizeEditorDataPoliciesPolicyChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetType() EnumAuthorizeEditorDataPoliciesPolicyChildCommonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetTypeOk() (*EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetType(v EnumAuthorizeEditorDataPoliciesPolicyChildCommonType)`

SetType sets Type field to given value.


### GetName

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetValue() AuthorizeEditorDataReferenceObjectDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetValue(v AuthorizeEditorDataReferenceObjectDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetCombiningAlgorithm() AuthorizeEditorDataPoliciesCombiningAlgorithmDTO`

GetCombiningAlgorithm returns the CombiningAlgorithm field if non-nil, zero value otherwise.

### GetCombiningAlgorithmOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetCombiningAlgorithmOk() (*AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, bool)`

GetCombiningAlgorithmOk returns a tuple with the CombiningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetCombiningAlgorithm(v AuthorizeEditorDataPoliciesCombiningAlgorithmDTO)`

SetCombiningAlgorithm sets CombiningAlgorithm field to given value.

### HasCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasCombiningAlgorithm() bool`

HasCombiningAlgorithm returns a boolean if a field has been set.

### GetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetChildren() []AuthorizeEditorDataPoliciesPolicyChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetChildrenOk() (*[]AuthorizeEditorDataPoliciesPolicyChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetChildren(v []AuthorizeEditorDataPoliciesPolicyChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetRepetitionSettings() AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

GetRepetitionSettings returns the RepetitionSettings field if non-nil, zero value otherwise.

### GetRepetitionSettingsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetRepetitionSettingsOk() (*AuthorizeEditorDataPoliciesRepetitionSettingsDTO, bool)`

GetRepetitionSettingsOk returns a tuple with the RepetitionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetRepetitionSettings(v AuthorizeEditorDataPoliciesRepetitionSettingsDTO)`

SetRepetitionSettings sets RepetitionSettings field to given value.

### HasRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasRepetitionSettings() bool`

HasRepetitionSettings returns a boolean if a field has been set.

### GetEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO`

GetEffectSettings returns the EffectSettings field if non-nil, zero value otherwise.

### GetEffectSettingsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChild) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool)`

GetEffectSettingsOk returns a tuple with the EffectSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO)`

SetEffectSettings sets EffectSettings field to given value.

### HasEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChild) HasEffectSettings() bool`

HasEffectSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


