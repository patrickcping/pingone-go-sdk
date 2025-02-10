# AuthorizeEditorDataDefinitionsProcessorDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | Pointer to [**EnumAuthorizeEditorDataDefinitionsProcessorDefinitionDTOType**](EnumAuthorizeEditorDataDefinitionsProcessorDefinitionDTOType.md) |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**Processor** | [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO

`func NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO(name string, processor AuthorizeEditorDataProcessorDTO, ) *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO`

NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO instantiates a new AuthorizeEditorDataDefinitionsProcessorDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTOWithDefaults() *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO`

NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataDefinitionsProcessorDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetType() EnumAuthorizeEditorDataDefinitionsProcessorDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsProcessorDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetType(v EnumAuthorizeEditorDataDefinitionsProcessorDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFullName

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProcessor

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


