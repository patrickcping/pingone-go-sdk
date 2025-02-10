# AuthorizeEditorDataProcessorsChainProcessorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**EnumAuthorizeEditorDataProcessorDTOType**](EnumAuthorizeEditorDataProcessorDTOType.md) |  | 
**Processors** | [**[]AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataProcessorsChainProcessorDTO

`func NewAuthorizeEditorDataProcessorsChainProcessorDTO(name string, type_ EnumAuthorizeEditorDataProcessorDTOType, processors []AuthorizeEditorDataProcessorDTO, ) *AuthorizeEditorDataProcessorsChainProcessorDTO`

NewAuthorizeEditorDataProcessorsChainProcessorDTO instantiates a new AuthorizeEditorDataProcessorsChainProcessorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataProcessorsChainProcessorDTOWithDefaults

`func NewAuthorizeEditorDataProcessorsChainProcessorDTOWithDefaults() *AuthorizeEditorDataProcessorsChainProcessorDTO`

NewAuthorizeEditorDataProcessorsChainProcessorDTOWithDefaults instantiates a new AuthorizeEditorDataProcessorsChainProcessorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetType() EnumAuthorizeEditorDataProcessorDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetTypeOk() (*EnumAuthorizeEditorDataProcessorDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) SetType(v EnumAuthorizeEditorDataProcessorDTOType)`

SetType sets Type field to given value.


### GetProcessors

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetProcessors() []AuthorizeEditorDataProcessorDTO`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetProcessorsOk() (*[]AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) SetProcessors(v []AuthorizeEditorDataProcessorDTO)`

SetProcessors sets Processors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


