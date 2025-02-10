# AuthorizeEditorDataConditionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataConditionDTOType**](EnumAuthorizeEditorDataConditionDTOType.md) |  | 
**Conditions** | [**[]AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | 
**Left** | [**AuthorizeEditorDataConditionsComparandLeftDTO**](AuthorizeEditorDataConditionsComparandLeftDTO.md) |  | 
**Right** | [**AuthorizeEditorDataConditionsComparandRightDTO**](AuthorizeEditorDataConditionsComparandRightDTO.md) |  | 
**Comparator** | [**EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator**](EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator.md) |  | 
**Condition** | [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | 
**Reference** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataConditionDTO

`func NewAuthorizeEditorDataConditionDTO(type_ EnumAuthorizeEditorDataConditionDTOType, conditions []AuthorizeEditorDataConditionDTO, left AuthorizeEditorDataConditionsComparandLeftDTO, right AuthorizeEditorDataConditionsComparandRightDTO, comparator EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator, condition AuthorizeEditorDataConditionDTO, reference AuthorizeEditorDataReferenceObjectDTO, ) *AuthorizeEditorDataConditionDTO`

NewAuthorizeEditorDataConditionDTO instantiates a new AuthorizeEditorDataConditionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataConditionDTOWithDefaults

`func NewAuthorizeEditorDataConditionDTOWithDefaults() *AuthorizeEditorDataConditionDTO`

NewAuthorizeEditorDataConditionDTOWithDefaults instantiates a new AuthorizeEditorDataConditionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataConditionDTO) GetType() EnumAuthorizeEditorDataConditionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataConditionDTO) GetTypeOk() (*EnumAuthorizeEditorDataConditionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataConditionDTO) SetType(v EnumAuthorizeEditorDataConditionDTOType)`

SetType sets Type field to given value.


### GetConditions

`func (o *AuthorizeEditorDataConditionDTO) GetConditions() []AuthorizeEditorDataConditionDTO`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizeEditorDataConditionDTO) GetConditionsOk() (*[]AuthorizeEditorDataConditionDTO, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizeEditorDataConditionDTO) SetConditions(v []AuthorizeEditorDataConditionDTO)`

SetConditions sets Conditions field to given value.


### GetLeft

`func (o *AuthorizeEditorDataConditionDTO) GetLeft() AuthorizeEditorDataConditionsComparandLeftDTO`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *AuthorizeEditorDataConditionDTO) GetLeftOk() (*AuthorizeEditorDataConditionsComparandLeftDTO, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *AuthorizeEditorDataConditionDTO) SetLeft(v AuthorizeEditorDataConditionsComparandLeftDTO)`

SetLeft sets Left field to given value.


### GetRight

`func (o *AuthorizeEditorDataConditionDTO) GetRight() AuthorizeEditorDataConditionsComparandRightDTO`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *AuthorizeEditorDataConditionDTO) GetRightOk() (*AuthorizeEditorDataConditionsComparandRightDTO, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *AuthorizeEditorDataConditionDTO) SetRight(v AuthorizeEditorDataConditionsComparandRightDTO)`

SetRight sets Right field to given value.


### GetComparator

`func (o *AuthorizeEditorDataConditionDTO) GetComparator() EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *AuthorizeEditorDataConditionDTO) GetComparatorOk() (*EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *AuthorizeEditorDataConditionDTO) SetComparator(v EnumAuthorizeEditorDataConditionsComparisonConditionDTOComparator)`

SetComparator sets Comparator field to given value.


### GetCondition

`func (o *AuthorizeEditorDataConditionDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataConditionDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataConditionDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.


### GetReference

`func (o *AuthorizeEditorDataConditionDTO) GetReference() AuthorizeEditorDataReferenceObjectDTO`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AuthorizeEditorDataConditionDTO) GetReferenceOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AuthorizeEditorDataConditionDTO) SetReference(v AuthorizeEditorDataReferenceObjectDTO)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


