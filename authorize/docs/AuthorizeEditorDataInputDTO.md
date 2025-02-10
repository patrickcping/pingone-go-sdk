# AuthorizeEditorDataInputDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataInputDTOType**](EnumAuthorizeEditorDataInputDTOType.md) |  | 
**Attribute** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Value** | **string** |  | 

## Methods

### NewAuthorizeEditorDataInputDTO

`func NewAuthorizeEditorDataInputDTO(type_ EnumAuthorizeEditorDataInputDTOType, attribute AuthorizeEditorDataReferenceObjectDTO, value string, ) *AuthorizeEditorDataInputDTO`

NewAuthorizeEditorDataInputDTO instantiates a new AuthorizeEditorDataInputDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataInputDTOWithDefaults

`func NewAuthorizeEditorDataInputDTOWithDefaults() *AuthorizeEditorDataInputDTO`

NewAuthorizeEditorDataInputDTOWithDefaults instantiates a new AuthorizeEditorDataInputDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataInputDTO) GetType() EnumAuthorizeEditorDataInputDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataInputDTO) GetTypeOk() (*EnumAuthorizeEditorDataInputDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataInputDTO) SetType(v EnumAuthorizeEditorDataInputDTOType)`

SetType sets Type field to given value.


### GetAttribute

`func (o *AuthorizeEditorDataInputDTO) GetAttribute() AuthorizeEditorDataReferenceObjectDTO`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AuthorizeEditorDataInputDTO) GetAttributeOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AuthorizeEditorDataInputDTO) SetAttribute(v AuthorizeEditorDataReferenceObjectDTO)`

SetAttribute sets Attribute field to given value.


### GetValue

`func (o *AuthorizeEditorDataInputDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataInputDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataInputDTO) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


