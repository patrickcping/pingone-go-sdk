# AuthorizeEditorDataServicesNoneServiceDefinitionDTO

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

## Methods

### NewAuthorizeEditorDataServicesNoneServiceDefinitionDTO

`func NewAuthorizeEditorDataServicesNoneServiceDefinitionDTO(name string, serviceType EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, ) *AuthorizeEditorDataServicesNoneServiceDefinitionDTO`

NewAuthorizeEditorDataServicesNoneServiceDefinitionDTO instantiates a new AuthorizeEditorDataServicesNoneServiceDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServicesNoneServiceDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataServicesNoneServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataServicesNoneServiceDefinitionDTO`

NewAuthorizeEditorDataServicesNoneServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataServicesNoneServiceDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetServiceType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) GetServiceTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AuthorizeEditorDataServicesNoneServiceDefinitionDTO) SetServiceType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


