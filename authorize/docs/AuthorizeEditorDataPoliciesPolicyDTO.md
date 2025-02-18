# AuthorizeEditorDataPoliciesPolicyDTO

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
**CombiningAlgorithm** | [**AuthorizeEditorDataPoliciesCombiningAlgorithmDTO**](AuthorizeEditorDataPoliciesCombiningAlgorithmDTO.md) |  | 
**Children** | Pointer to [**[]AuthorizeEditorDataPoliciesPolicyChild**](AuthorizeEditorDataPoliciesPolicyChild.md) |  | [optional] 
**RepetitionSettings** | Pointer to [**AuthorizeEditorDataPoliciesRepetitionSettingsDTO**](AuthorizeEditorDataPoliciesRepetitionSettingsDTO.md) |  | [optional] 
**ManagedEntity** | Pointer to [**AuthorizeEditorDataManagedEntityDTO**](AuthorizeEditorDataManagedEntityDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataPoliciesPolicyDTO

`func NewAuthorizeEditorDataPoliciesPolicyDTO(name string, combiningAlgorithm AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, ) *AuthorizeEditorDataPoliciesPolicyDTO`

NewAuthorizeEditorDataPoliciesPolicyDTO instantiates a new AuthorizeEditorDataPoliciesPolicyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesPolicyDTOWithDefaults

`func NewAuthorizeEditorDataPoliciesPolicyDTOWithDefaults() *AuthorizeEditorDataPoliciesPolicyDTO`

NewAuthorizeEditorDataPoliciesPolicyDTOWithDefaults instantiates a new AuthorizeEditorDataPoliciesPolicyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetCombiningAlgorithm() AuthorizeEditorDataPoliciesCombiningAlgorithmDTO`

GetCombiningAlgorithm returns the CombiningAlgorithm field if non-nil, zero value otherwise.

### GetCombiningAlgorithmOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetCombiningAlgorithmOk() (*AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, bool)`

GetCombiningAlgorithmOk returns a tuple with the CombiningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombiningAlgorithm

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetCombiningAlgorithm(v AuthorizeEditorDataPoliciesCombiningAlgorithmDTO)`

SetCombiningAlgorithm sets CombiningAlgorithm field to given value.


### GetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetChildren() []AuthorizeEditorDataPoliciesPolicyChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetChildrenOk() (*[]AuthorizeEditorDataPoliciesPolicyChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetChildren(v []AuthorizeEditorDataPoliciesPolicyChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetRepetitionSettings() AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

GetRepetitionSettings returns the RepetitionSettings field if non-nil, zero value otherwise.

### GetRepetitionSettingsOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetRepetitionSettingsOk() (*AuthorizeEditorDataPoliciesRepetitionSettingsDTO, bool)`

GetRepetitionSettingsOk returns a tuple with the RepetitionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetRepetitionSettings(v AuthorizeEditorDataPoliciesRepetitionSettingsDTO)`

SetRepetitionSettings sets RepetitionSettings field to given value.

### HasRepetitionSettings

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasRepetitionSettings() bool`

HasRepetitionSettings returns a boolean if a field has been set.

### GetManagedEntity

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetManagedEntity() AuthorizeEditorDataManagedEntityDTO`

GetManagedEntity returns the ManagedEntity field if non-nil, zero value otherwise.

### GetManagedEntityOk

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) GetManagedEntityOk() (*AuthorizeEditorDataManagedEntityDTO, bool)`

GetManagedEntityOk returns a tuple with the ManagedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedEntity

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) SetManagedEntity(v AuthorizeEditorDataManagedEntityDTO)`

SetManagedEntity sets ManagedEntity field to given value.

### HasManagedEntity

`func (o *AuthorizeEditorDataPoliciesPolicyDTO) HasManagedEntity() bool`

HasManagedEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


