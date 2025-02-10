# AuthorizeEditorDataAuthorizationVersionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ChangeDetails** | Pointer to [**AuthorizeEditorDataChangeDetailsDTO**](AuthorizeEditorDataChangeDetailsDTO.md) |  | [optional] 
**Tag** | Pointer to [**AuthorizeEditorDataTagResponseDTO**](AuthorizeEditorDataTagResponseDTO.md) |  | [optional] 
**Environment** | Pointer to [**AuthorizeEditorDataEnvironmentDTO**](AuthorizeEditorDataEnvironmentDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataAuthorizationVersionDTO

`func NewAuthorizeEditorDataAuthorizationVersionDTO() *AuthorizeEditorDataAuthorizationVersionDTO`

NewAuthorizeEditorDataAuthorizationVersionDTO instantiates a new AuthorizeEditorDataAuthorizationVersionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataAuthorizationVersionDTOWithDefaults

`func NewAuthorizeEditorDataAuthorizationVersionDTOWithDefaults() *AuthorizeEditorDataAuthorizationVersionDTO`

NewAuthorizeEditorDataAuthorizationVersionDTOWithDefaults instantiates a new AuthorizeEditorDataAuthorizationVersionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetChangeDetails

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetChangeDetails() AuthorizeEditorDataChangeDetailsDTO`

GetChangeDetails returns the ChangeDetails field if non-nil, zero value otherwise.

### GetChangeDetailsOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetChangeDetailsOk() (*AuthorizeEditorDataChangeDetailsDTO, bool)`

GetChangeDetailsOk returns a tuple with the ChangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDetails

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetChangeDetails(v AuthorizeEditorDataChangeDetailsDTO)`

SetChangeDetails sets ChangeDetails field to given value.

### HasChangeDetails

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasChangeDetails() bool`

HasChangeDetails returns a boolean if a field has been set.

### GetTag

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetTag() AuthorizeEditorDataTagResponseDTO`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetTagOk() (*AuthorizeEditorDataTagResponseDTO, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetTag(v AuthorizeEditorDataTagResponseDTO)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetEnvironment() AuthorizeEditorDataEnvironmentDTO`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) GetEnvironmentOk() (*AuthorizeEditorDataEnvironmentDTO, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) SetEnvironment(v AuthorizeEditorDataEnvironmentDTO)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataAuthorizationVersionDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


