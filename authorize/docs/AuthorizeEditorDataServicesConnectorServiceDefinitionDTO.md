# AuthorizeEditorDataServicesConnectorServiceDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**Type** | Pointer to [**EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType**](EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType.md) |  | [optional] 
**CacheSettings** | Pointer to [**AuthorizeEditorDataCacheSettingsDTO**](AuthorizeEditorDataCacheSettingsDTO.md) |  | [optional] 
**ServiceType** | [**EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType**](EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType.md) |  | 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**ServiceSettings** | [**AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO**](AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO

`func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO(name string, serviceType EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, ) *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO`

NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO`

NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetServiceType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType)`

SetServiceType sets ServiceType field to given value.


### GetProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettings() AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetServiceSettings(v AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


