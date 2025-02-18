# AuthorizeEditorDataAttributeResolversRequestResolverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**Type** | [**EnumAuthorizeEditorDataResolverDTOType**](EnumAuthorizeEditorDataResolverDTOType.md) |  | 

## Methods

### NewAuthorizeEditorDataAttributeResolversRequestResolverDTO

`func NewAuthorizeEditorDataAttributeResolversRequestResolverDTO(type_ EnumAuthorizeEditorDataResolverDTOType, ) *AuthorizeEditorDataAttributeResolversRequestResolverDTO`

NewAuthorizeEditorDataAttributeResolversRequestResolverDTO instantiates a new AuthorizeEditorDataAttributeResolversRequestResolverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataAttributeResolversRequestResolverDTOWithDefaults

`func NewAuthorizeEditorDataAttributeResolversRequestResolverDTOWithDefaults() *AuthorizeEditorDataAttributeResolversRequestResolverDTO`

NewAuthorizeEditorDataAttributeResolversRequestResolverDTOWithDefaults instantiates a new AuthorizeEditorDataAttributeResolversRequestResolverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetProcessor

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetType() EnumAuthorizeEditorDataResolverDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) GetTypeOk() (*EnumAuthorizeEditorDataResolverDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataAttributeResolversRequestResolverDTO) SetType(v EnumAuthorizeEditorDataResolverDTOType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


