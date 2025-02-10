# AuthorizeEditorDataProcessorsReferenceProcessorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**EnumAuthorizeEditorDataProcessorDTOType**](EnumAuthorizeEditorDataProcessorDTOType.md) |  | 
**Processor** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataProcessorsReferenceProcessorDTO

`func NewAuthorizeEditorDataProcessorsReferenceProcessorDTO(name string, type_ EnumAuthorizeEditorDataProcessorDTOType, processor AuthorizeEditorDataReferenceObjectDTO, ) *AuthorizeEditorDataProcessorsReferenceProcessorDTO`

NewAuthorizeEditorDataProcessorsReferenceProcessorDTO instantiates a new AuthorizeEditorDataProcessorsReferenceProcessorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataProcessorsReferenceProcessorDTOWithDefaults

`func NewAuthorizeEditorDataProcessorsReferenceProcessorDTOWithDefaults() *AuthorizeEditorDataProcessorsReferenceProcessorDTO`

NewAuthorizeEditorDataProcessorsReferenceProcessorDTOWithDefaults instantiates a new AuthorizeEditorDataProcessorsReferenceProcessorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetType() EnumAuthorizeEditorDataProcessorDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetTypeOk() (*EnumAuthorizeEditorDataProcessorDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) SetType(v EnumAuthorizeEditorDataProcessorDTOType)`

SetType sets Type field to given value.


### GetProcessor

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetProcessor() AuthorizeEditorDataReferenceObjectDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) GetProcessorOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataProcessorsReferenceProcessorDTO) SetProcessor(v AuthorizeEditorDataReferenceObjectDTO)`

SetProcessor sets Processor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


