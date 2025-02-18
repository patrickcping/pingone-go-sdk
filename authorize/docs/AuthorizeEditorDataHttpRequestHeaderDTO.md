# AuthorizeEditorDataHttpRequestHeaderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to [**AuthorizeEditorDataInputDTO**](AuthorizeEditorDataInputDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataHttpRequestHeaderDTO

`func NewAuthorizeEditorDataHttpRequestHeaderDTO(key string, ) *AuthorizeEditorDataHttpRequestHeaderDTO`

NewAuthorizeEditorDataHttpRequestHeaderDTO instantiates a new AuthorizeEditorDataHttpRequestHeaderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataHttpRequestHeaderDTOWithDefaults

`func NewAuthorizeEditorDataHttpRequestHeaderDTOWithDefaults() *AuthorizeEditorDataHttpRequestHeaderDTO`

NewAuthorizeEditorDataHttpRequestHeaderDTOWithDefaults instantiates a new AuthorizeEditorDataHttpRequestHeaderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) GetValue() AuthorizeEditorDataInputDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) GetValueOk() (*AuthorizeEditorDataInputDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) SetValue(v AuthorizeEditorDataInputDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthorizeEditorDataHttpRequestHeaderDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


