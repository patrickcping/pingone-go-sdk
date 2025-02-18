# AuthorizeEditorDataChangeDetailsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedResource** | Pointer to [**TypedRelationship**](TypedRelationship.md) |  | [optional] 
**ChangeType** | Pointer to [**EnumAuthorizeEditorDataChangeDetailsDTOChangeType**](EnumAuthorizeEditorDataChangeDetailsDTOChangeType.md) |  | [optional] 
**ChangedEntityName** | Pointer to **string** |  | [optional] 
**ChangedEntityType** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**AuthorizeEditorDataUserDTO**](AuthorizeEditorDataUserDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataChangeDetailsDTO

`func NewAuthorizeEditorDataChangeDetailsDTO() *AuthorizeEditorDataChangeDetailsDTO`

NewAuthorizeEditorDataChangeDetailsDTO instantiates a new AuthorizeEditorDataChangeDetailsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataChangeDetailsDTOWithDefaults

`func NewAuthorizeEditorDataChangeDetailsDTOWithDefaults() *AuthorizeEditorDataChangeDetailsDTO`

NewAuthorizeEditorDataChangeDetailsDTOWithDefaults instantiates a new AuthorizeEditorDataChangeDetailsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedResource

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedResource() TypedRelationship`

GetChangedResource returns the ChangedResource field if non-nil, zero value otherwise.

### GetChangedResourceOk

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedResourceOk() (*TypedRelationship, bool)`

GetChangedResourceOk returns a tuple with the ChangedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedResource

`func (o *AuthorizeEditorDataChangeDetailsDTO) SetChangedResource(v TypedRelationship)`

SetChangedResource sets ChangedResource field to given value.

### HasChangedResource

`func (o *AuthorizeEditorDataChangeDetailsDTO) HasChangedResource() bool`

HasChangedResource returns a boolean if a field has been set.

### GetChangeType

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangeType() EnumAuthorizeEditorDataChangeDetailsDTOChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangeTypeOk() (*EnumAuthorizeEditorDataChangeDetailsDTOChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *AuthorizeEditorDataChangeDetailsDTO) SetChangeType(v EnumAuthorizeEditorDataChangeDetailsDTOChangeType)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *AuthorizeEditorDataChangeDetailsDTO) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetChangedEntityName

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedEntityName() string`

GetChangedEntityName returns the ChangedEntityName field if non-nil, zero value otherwise.

### GetChangedEntityNameOk

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedEntityNameOk() (*string, bool)`

GetChangedEntityNameOk returns a tuple with the ChangedEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedEntityName

`func (o *AuthorizeEditorDataChangeDetailsDTO) SetChangedEntityName(v string)`

SetChangedEntityName sets ChangedEntityName field to given value.

### HasChangedEntityName

`func (o *AuthorizeEditorDataChangeDetailsDTO) HasChangedEntityName() bool`

HasChangedEntityName returns a boolean if a field has been set.

### GetChangedEntityType

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedEntityType() string`

GetChangedEntityType returns the ChangedEntityType field if non-nil, zero value otherwise.

### GetChangedEntityTypeOk

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetChangedEntityTypeOk() (*string, bool)`

GetChangedEntityTypeOk returns a tuple with the ChangedEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedEntityType

`func (o *AuthorizeEditorDataChangeDetailsDTO) SetChangedEntityType(v string)`

SetChangedEntityType sets ChangedEntityType field to given value.

### HasChangedEntityType

`func (o *AuthorizeEditorDataChangeDetailsDTO) HasChangedEntityType() bool`

HasChangedEntityType returns a boolean if a field has been set.

### GetUser

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetUser() AuthorizeEditorDataUserDTO`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthorizeEditorDataChangeDetailsDTO) GetUserOk() (*AuthorizeEditorDataUserDTO, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthorizeEditorDataChangeDetailsDTO) SetUser(v AuthorizeEditorDataUserDTO)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthorizeEditorDataChangeDetailsDTO) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


