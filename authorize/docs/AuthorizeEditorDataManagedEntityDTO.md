# AuthorizeEditorDataManagedEntityDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [**AuthorizeEditorDataManagedEntityOwnerDTO**](AuthorizeEditorDataManagedEntityOwnerDTO.md) |  | 
**Restrictions** | Pointer to [**AuthorizeEditorDataManagedEntityRestrictionsDTO**](AuthorizeEditorDataManagedEntityRestrictionsDTO.md) |  | [optional] 
**Reference** | Pointer to [**AuthorizeEditorDataManagedEntityManagedEntityReferenceDTO**](AuthorizeEditorDataManagedEntityManagedEntityReferenceDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataManagedEntityDTO

`func NewAuthorizeEditorDataManagedEntityDTO(owner AuthorizeEditorDataManagedEntityOwnerDTO, ) *AuthorizeEditorDataManagedEntityDTO`

NewAuthorizeEditorDataManagedEntityDTO instantiates a new AuthorizeEditorDataManagedEntityDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataManagedEntityDTOWithDefaults

`func NewAuthorizeEditorDataManagedEntityDTOWithDefaults() *AuthorizeEditorDataManagedEntityDTO`

NewAuthorizeEditorDataManagedEntityDTOWithDefaults instantiates a new AuthorizeEditorDataManagedEntityDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AuthorizeEditorDataManagedEntityDTO) GetOwner() AuthorizeEditorDataManagedEntityOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AuthorizeEditorDataManagedEntityDTO) GetOwnerOk() (*AuthorizeEditorDataManagedEntityOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AuthorizeEditorDataManagedEntityDTO) SetOwner(v AuthorizeEditorDataManagedEntityOwnerDTO)`

SetOwner sets Owner field to given value.


### GetRestrictions

`func (o *AuthorizeEditorDataManagedEntityDTO) GetRestrictions() AuthorizeEditorDataManagedEntityRestrictionsDTO`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *AuthorizeEditorDataManagedEntityDTO) GetRestrictionsOk() (*AuthorizeEditorDataManagedEntityRestrictionsDTO, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *AuthorizeEditorDataManagedEntityDTO) SetRestrictions(v AuthorizeEditorDataManagedEntityRestrictionsDTO)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *AuthorizeEditorDataManagedEntityDTO) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetReference

`func (o *AuthorizeEditorDataManagedEntityDTO) GetReference() AuthorizeEditorDataManagedEntityManagedEntityReferenceDTO`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AuthorizeEditorDataManagedEntityDTO) GetReferenceOk() (*AuthorizeEditorDataManagedEntityManagedEntityReferenceDTO, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AuthorizeEditorDataManagedEntityDTO) SetReference(v AuthorizeEditorDataManagedEntityManagedEntityReferenceDTO)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AuthorizeEditorDataManagedEntityDTO) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


