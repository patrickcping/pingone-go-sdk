# AuthorizeEditorDataProcessorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**EnumAuthorizeEditorDataProcessorDTOType**](EnumAuthorizeEditorDataProcessorDTOType.md) |  | 
**Processors** | [**[]AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | 
**Predicate** | [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | 
**Processor** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Expression** | **string** |  | 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataProcessorDTO

`func NewAuthorizeEditorDataProcessorDTO(name string, type_ EnumAuthorizeEditorDataProcessorDTOType, processors []AuthorizeEditorDataProcessorDTO, predicate AuthorizeEditorDataProcessorDTO, processor AuthorizeEditorDataReferenceObjectDTO, expression string, valueType AuthorizeEditorDataValueTypeDTO, ) *AuthorizeEditorDataProcessorDTO`

NewAuthorizeEditorDataProcessorDTO instantiates a new AuthorizeEditorDataProcessorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataProcessorDTOWithDefaults

`func NewAuthorizeEditorDataProcessorDTOWithDefaults() *AuthorizeEditorDataProcessorDTO`

NewAuthorizeEditorDataProcessorDTOWithDefaults instantiates a new AuthorizeEditorDataProcessorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorizeEditorDataProcessorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataProcessorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataProcessorDTO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AuthorizeEditorDataProcessorDTO) GetType() EnumAuthorizeEditorDataProcessorDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataProcessorDTO) GetTypeOk() (*EnumAuthorizeEditorDataProcessorDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataProcessorDTO) SetType(v EnumAuthorizeEditorDataProcessorDTOType)`

SetType sets Type field to given value.


### GetProcessors

`func (o *AuthorizeEditorDataProcessorDTO) GetProcessors() []AuthorizeEditorDataProcessorDTO`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *AuthorizeEditorDataProcessorDTO) GetProcessorsOk() (*[]AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *AuthorizeEditorDataProcessorDTO) SetProcessors(v []AuthorizeEditorDataProcessorDTO)`

SetProcessors sets Processors field to given value.


### GetPredicate

`func (o *AuthorizeEditorDataProcessorDTO) GetPredicate() AuthorizeEditorDataProcessorDTO`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *AuthorizeEditorDataProcessorDTO) GetPredicateOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *AuthorizeEditorDataProcessorDTO) SetPredicate(v AuthorizeEditorDataProcessorDTO)`

SetPredicate sets Predicate field to given value.


### GetProcessor

`func (o *AuthorizeEditorDataProcessorDTO) GetProcessor() AuthorizeEditorDataReferenceObjectDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataProcessorDTO) GetProcessorOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataProcessorDTO) SetProcessor(v AuthorizeEditorDataReferenceObjectDTO)`

SetProcessor sets Processor field to given value.


### GetExpression

`func (o *AuthorizeEditorDataProcessorDTO) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *AuthorizeEditorDataProcessorDTO) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *AuthorizeEditorDataProcessorDTO) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetValueType

`func (o *AuthorizeEditorDataProcessorDTO) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AuthorizeEditorDataProcessorDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AuthorizeEditorDataProcessorDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


