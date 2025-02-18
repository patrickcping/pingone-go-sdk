# AuthorizeEditorDataRulesRuleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Statements** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**EffectSettings** | [**AuthorizeEditorDataRulesEffectSettingsDTO**](AuthorizeEditorDataRulesEffectSettingsDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataRulesRuleDTO

`func NewAuthorizeEditorDataRulesRuleDTO(name string, effectSettings AuthorizeEditorDataRulesEffectSettingsDTO, ) *AuthorizeEditorDataRulesRuleDTO`

NewAuthorizeEditorDataRulesRuleDTO instantiates a new AuthorizeEditorDataRulesRuleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataRulesRuleDTOWithDefaults

`func NewAuthorizeEditorDataRulesRuleDTOWithDefaults() *AuthorizeEditorDataRulesRuleDTO`

NewAuthorizeEditorDataRulesRuleDTOWithDefaults instantiates a new AuthorizeEditorDataRulesRuleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataRulesRuleDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataRulesRuleDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataRulesRuleDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataRulesRuleDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataRulesRuleDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataRulesRuleDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataRulesRuleDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataRulesRuleDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataRulesRuleDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataRulesRuleDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataRulesRuleDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataRulesRuleDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataRulesRuleDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataRulesRuleDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataRulesRuleDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataRulesRuleDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataRulesRuleDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataRulesRuleDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataRulesRuleDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataRulesRuleDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataRulesRuleDTO) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataRulesRuleDTO) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataRulesRuleDTO) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataRulesRuleDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataRulesRuleDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataRulesRuleDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEffectSettings

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO`

GetEffectSettings returns the EffectSettings field if non-nil, zero value otherwise.

### GetEffectSettingsOk

`func (o *AuthorizeEditorDataRulesRuleDTO) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool)`

GetEffectSettingsOk returns a tuple with the EffectSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectSettings

`func (o *AuthorizeEditorDataRulesRuleDTO) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO)`

SetEffectSettings sets EffectSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


