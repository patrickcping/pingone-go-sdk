# AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon

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

### NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon

`func NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon(name string, serviceType EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, ) *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon`

NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon instantiates a new AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommonWithDefaults

`func NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommonWithDefaults() *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon`

NewAuthorizeEditorDataDefinitionsServiceDefinitionDTOCommonWithDefaults instantiates a new AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetServiceType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) GetServiceTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon) SetServiceType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


