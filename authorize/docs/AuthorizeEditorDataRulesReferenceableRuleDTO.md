# AuthorizeEditorDataRulesReferenceableRuleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Id** | **string** | The resource&#39;s unique identifier | [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Statements** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**EffectSettings** | [**AuthorizeEditorDataRulesEffectSettingsDTO**](AuthorizeEditorDataRulesEffectSettingsDTO.md) |  | 
**Version** | **string** |  | 

## Methods

### NewAuthorizeEditorDataRulesReferenceableRuleDTO

`func NewAuthorizeEditorDataRulesReferenceableRuleDTO(id string, name string, effectSettings AuthorizeEditorDataRulesEffectSettingsDTO, version string, ) *AuthorizeEditorDataRulesReferenceableRuleDTO`

NewAuthorizeEditorDataRulesReferenceableRuleDTO instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults

`func NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults() *AuthorizeEditorDataRulesReferenceableRuleDTO`

NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEffectSettings

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO`

GetEffectSettings returns the EffectSettings field if non-nil, zero value otherwise.

### GetEffectSettingsOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool)`

GetEffectSettingsOk returns a tuple with the EffectSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectSettings

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO)`

SetEffectSettings sets EffectSettings field to given value.


### GetVersion

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


