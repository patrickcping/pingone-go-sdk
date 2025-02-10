# AuthorizeEditorDataDefinitionsAttributeDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**EnumAuthorizeEditorDataDefinitionsAttributeDefinitionDTOType**](EnumAuthorizeEditorDataDefinitionsAttributeDefinitionDTOType.md) |  | [optional] 
**Parent** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**Resolvers** | Pointer to [**[]AuthorizeEditorDataResolverDTO**](AuthorizeEditorDataResolverDTO.md) |  | [optional] 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**RepetitionSource** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**ValueSchema** | Pointer to **string** |  | [optional] 
**ManagedEntity** | Pointer to [**AuthorizeEditorDataManagedEntityDTO**](AuthorizeEditorDataManagedEntityDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTO

`func NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTO(name string, valueType AuthorizeEditorDataValueTypeDTO, ) *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO`

NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTO instantiates a new AuthorizeEditorDataDefinitionsAttributeDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTOWithDefaults() *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO`

NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataDefinitionsAttributeDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetType() EnumAuthorizeEditorDataDefinitionsAttributeDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsAttributeDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetType(v EnumAuthorizeEditorDataDefinitionsAttributeDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParent

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetResolvers

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetResolvers() []AuthorizeEditorDataResolverDTO`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetResolversOk() (*[]AuthorizeEditorDataResolverDTO, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetResolvers(v []AuthorizeEditorDataResolverDTO)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.

### GetProcessor

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetDefaultValue

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRepetitionSource

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetRepetitionSource() AuthorizeEditorDataReferenceObjectDTO`

GetRepetitionSource returns the RepetitionSource field if non-nil, zero value otherwise.

### GetRepetitionSourceOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetRepetitionSourceOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetRepetitionSourceOk returns a tuple with the RepetitionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionSource

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetRepetitionSource(v AuthorizeEditorDataReferenceObjectDTO)`

SetRepetitionSource sets RepetitionSource field to given value.

### HasRepetitionSource

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasRepetitionSource() bool`

HasRepetitionSource returns a boolean if a field has been set.

### GetValueSchema

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetValueSchema() string`

GetValueSchema returns the ValueSchema field if non-nil, zero value otherwise.

### GetValueSchemaOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetValueSchemaOk() (*string, bool)`

GetValueSchemaOk returns a tuple with the ValueSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSchema

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetValueSchema(v string)`

SetValueSchema sets ValueSchema field to given value.

### HasValueSchema

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasValueSchema() bool`

HasValueSchema returns a boolean if a field has been set.

### GetManagedEntity

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetManagedEntity() AuthorizeEditorDataManagedEntityDTO`

GetManagedEntity returns the ManagedEntity field if non-nil, zero value otherwise.

### GetManagedEntityOk

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) GetManagedEntityOk() (*AuthorizeEditorDataManagedEntityDTO, bool)`

GetManagedEntityOk returns a tuple with the ManagedEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedEntity

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) SetManagedEntity(v AuthorizeEditorDataManagedEntityDTO)`

SetManagedEntity sets ManagedEntity field to given value.

### HasManagedEntity

`func (o *AuthorizeEditorDataDefinitionsAttributeDefinitionDTO) HasManagedEntity() bool`

HasManagedEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


