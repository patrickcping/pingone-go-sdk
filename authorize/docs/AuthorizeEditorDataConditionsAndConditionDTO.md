# AuthorizeEditorDataConditionsAndConditionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataConditionDTOType**](EnumAuthorizeEditorDataConditionDTOType.md) |  | 
**Conditions** | [**[]AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataConditionsAndConditionDTO

`func NewAuthorizeEditorDataConditionsAndConditionDTO(type_ EnumAuthorizeEditorDataConditionDTOType, conditions []AuthorizeEditorDataConditionDTO, ) *AuthorizeEditorDataConditionsAndConditionDTO`

NewAuthorizeEditorDataConditionsAndConditionDTO instantiates a new AuthorizeEditorDataConditionsAndConditionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataConditionsAndConditionDTOWithDefaults

`func NewAuthorizeEditorDataConditionsAndConditionDTOWithDefaults() *AuthorizeEditorDataConditionsAndConditionDTO`

NewAuthorizeEditorDataConditionsAndConditionDTOWithDefaults instantiates a new AuthorizeEditorDataConditionsAndConditionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) GetType() EnumAuthorizeEditorDataConditionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) GetTypeOk() (*EnumAuthorizeEditorDataConditionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) SetType(v EnumAuthorizeEditorDataConditionDTOType)`

SetType sets Type field to given value.


### GetConditions

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) GetConditions() []AuthorizeEditorDataConditionDTO`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) GetConditionsOk() (*[]AuthorizeEditorDataConditionDTO, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizeEditorDataConditionsAndConditionDTO) SetConditions(v []AuthorizeEditorDataConditionDTO)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


