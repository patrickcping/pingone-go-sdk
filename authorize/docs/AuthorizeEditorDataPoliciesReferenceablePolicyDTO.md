# AuthorizeEditorDataPoliciesReferenceablePolicyDTO

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
**CombiningAlgorithm** | [**AuthorizeEditorDataPoliciesCombiningAlgorithmDTO**](AuthorizeEditorDataPoliciesCombiningAlgorithmDTO.md) |  | 
**Children** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RepetitionSettings** | Pointer to [**AuthorizeEditorDataPoliciesRepetitionSettingsDTO**](AuthorizeEditorDataPoliciesRepetitionSettingsDTO.md) |  | [optional] 
**ManagedEntity** | Pointer to [**AuthorizeEditorDataManagedEntityDTO**](AuthorizeEditorDataManagedEntityDTO.md) |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO

`func NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO(id string, name string, combiningAlgorithm AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, version string, ) *AuthorizeEditorDataPoliciesReferenceablePolicyDTO`

NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO instantiates a new AuthorizeEditorDataPoliciesReferenceablePolicyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesReferenceablePolicyDTOWithDefaults

`func NewAuthorizeEditorDataPoliciesReferenceablePolicyDTOWithDefaults() *AuthorizeEditorDataPoliciesReferenceablePolicyDTO`

NewAuthorizeEditorDataPoliciesReferenceablePolicyDTOWithDefaults instantiates a new AuthorizeEditorDataPoliciesReferenceablePolicyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCombiningAlgorithm() AuthorizeEditorDataPoliciesCombiningAlgorithmDTO`

GetCombiningAlgorithm returns the CombiningAlgorithm field if non-nil, zero value otherwise.

### GetCombiningAlgorithmOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCombiningAlgorithmOk() (*AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, bool)`

GetCombiningAlgorithmOk returns a tuple with the CombiningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetCombiningAlgorithm(v AuthorizeEditorDataPoliciesCombiningAlgorithmDTO)`

SetCombiningAlgorithm sets CombiningAlgorithm field to given value.


### GetChildren

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetChildren() []map[string]interface{}`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetChildrenOk() (*[]map[string]interface{}, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetChildren(v []map[string]interface{})`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetRepetitionSettings() AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

GetRepetitionSettings returns the RepetitionSettings field if non-nil, zero value otherwise.

### GetRepetitionSettingsOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetRepetitionSettingsOk() (*AuthorizeEditorDataPoliciesRepetitionSettingsDTO, bool)`

GetRepetitionSettingsOk returns a tuple with the RepetitionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetRepetitionSettings(v AuthorizeEditorDataPoliciesRepetitionSettingsDTO)`

SetRepetitionSettings sets RepetitionSettings field to given value.

### HasRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasRepetitionSettings() bool`

HasRepetitionSettings returns a boolean if a field has been set.

### GetManagedEntity

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetManagedEntity() AuthorizeEditorDataManagedEntityDTO`

GetManagedEntity returns the ManagedEntity field if non-nil, zero value otherwise.

### GetManagedEntityOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetManagedEntityOk() (*AuthorizeEditorDataManagedEntityDTO, bool)`

GetManagedEntityOk returns a tuple with the ManagedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedEntity

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetManagedEntity(v AuthorizeEditorDataManagedEntityDTO)`

SetManagedEntity sets ManagedEntity field to given value.

### HasManagedEntity

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasManagedEntity() bool`

HasManagedEntity returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


