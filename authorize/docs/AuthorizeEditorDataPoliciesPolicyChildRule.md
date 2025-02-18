# AuthorizeEditorDataPoliciesPolicyChildRule

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
**EffectSettings** | Pointer to [**AuthorizeEditorDataRulesEffectSettingsDTO**](AuthorizeEditorDataRulesEffectSettingsDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataPoliciesPolicyChildRule

`func NewAuthorizeEditorDataPoliciesPolicyChildRule(type_ EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, ) *AuthorizeEditorDataPoliciesPolicyChildRule`

NewAuthorizeEditorDataPoliciesPolicyChildRule instantiates a new AuthorizeEditorDataPoliciesPolicyChildRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesPolicyChildRuleWithDefaults

`func NewAuthorizeEditorDataPoliciesPolicyChildRuleWithDefaults() *AuthorizeEditorDataPoliciesPolicyChildRule`

NewAuthorizeEditorDataPoliciesPolicyChildRuleWithDefaults instantiates a new AuthorizeEditorDataPoliciesPolicyChildRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetType() EnumAuthorizeEditorDataPoliciesPolicyChildCommonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetTypeOk() (*EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetType(v EnumAuthorizeEditorDataPoliciesPolicyChildCommonType)`

SetType sets Type field to given value.


### GetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetValue() AuthorizeEditorDataReferenceObjectDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetValue(v AuthorizeEditorDataReferenceObjectDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO`

GetEffectSettings returns the EffectSettings field if non-nil, zero value otherwise.

### GetEffectSettingsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool)`

GetEffectSettingsOk returns a tuple with the EffectSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO)`

SetEffectSettings sets EffectSettings field to given value.

### HasEffectSettings

`func (o *AuthorizeEditorDataPoliciesPolicyChildRule) HasEffectSettings() bool`

HasEffectSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


