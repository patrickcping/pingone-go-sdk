# AuthorizeEditorDataPoliciesPolicyChildPolicy

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

## Methods

### NewAuthorizeEditorDataPoliciesPolicyChildPolicy

`func NewAuthorizeEditorDataPoliciesPolicyChildPolicy(type_ EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, ) *AuthorizeEditorDataPoliciesPolicyChildPolicy`

NewAuthorizeEditorDataPoliciesPolicyChildPolicy instantiates a new AuthorizeEditorDataPoliciesPolicyChildPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesPolicyChildPolicyWithDefaults

`func NewAuthorizeEditorDataPoliciesPolicyChildPolicyWithDefaults() *AuthorizeEditorDataPoliciesPolicyChildPolicy`

NewAuthorizeEditorDataPoliciesPolicyChildPolicyWithDefaults instantiates a new AuthorizeEditorDataPoliciesPolicyChildPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetType() EnumAuthorizeEditorDataPoliciesPolicyChildCommonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetTypeOk() (*EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetType(v EnumAuthorizeEditorDataPoliciesPolicyChildCommonType)`

SetType sets Type field to given value.


### GetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetValue() AuthorizeEditorDataReferenceObjectDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetValue(v AuthorizeEditorDataReferenceObjectDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetCombiningAlgorithm() AuthorizeEditorDataPoliciesCombiningAlgorithmDTO`

GetCombiningAlgorithm returns the CombiningAlgorithm field if non-nil, zero value otherwise.

### GetCombiningAlgorithmOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetCombiningAlgorithmOk() (*AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, bool)`

GetCombiningAlgorithmOk returns a tuple with the CombiningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetCombiningAlgorithm(v AuthorizeEditorDataPoliciesCombiningAlgorithmDTO)`

SetCombiningAlgorithm sets CombiningAlgorithm field to given value.

### HasCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasCombiningAlgorithm() bool`

HasCombiningAlgorithm returns a boolean if a field has been set.

### GetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetChildren() []AuthorizeEditorDataPoliciesPolicyChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetChildrenOk() (*[]AuthorizeEditorDataPoliciesPolicyChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetChildren(v []AuthorizeEditorDataPoliciesPolicyChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetRepetitionSettings() AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

GetRepetitionSettings returns the RepetitionSettings field if non-nil, zero value otherwise.

### GetRepetitionSettingsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) GetRepetitionSettingsOk() (*AuthorizeEditorDataPoliciesRepetitionSettingsDTO, bool)`

GetRepetitionSettingsOk returns a tuple with the RepetitionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) SetRepetitionSettings(v AuthorizeEditorDataPoliciesRepetitionSettingsDTO)`

SetRepetitionSettings sets RepetitionSettings field to given value.

### HasRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildPolicy) HasRepetitionSettings() bool`

HasRepetitionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


