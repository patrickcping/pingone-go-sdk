# AuthorizeEditorDataConditionsOrConditionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataConditionDTOType**](EnumAuthorizeEditorDataConditionDTOType.md) |  | 
**Conditions** | [**[]AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataConditionsOrConditionDTO

`func NewAuthorizeEditorDataConditionsOrConditionDTO(type_ EnumAuthorizeEditorDataConditionDTOType, conditions []AuthorizeEditorDataConditionDTO, ) *AuthorizeEditorDataConditionsOrConditionDTO`

NewAuthorizeEditorDataConditionsOrConditionDTO instantiates a new AuthorizeEditorDataConditionsOrConditionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataConditionsOrConditionDTOWithDefaults

`func NewAuthorizeEditorDataConditionsOrConditionDTOWithDefaults() *AuthorizeEditorDataConditionsOrConditionDTO`

NewAuthorizeEditorDataConditionsOrConditionDTOWithDefaults instantiates a new AuthorizeEditorDataConditionsOrConditionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetType() EnumAuthorizeEditorDataConditionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetTypeOk() (*EnumAuthorizeEditorDataConditionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) SetType(v EnumAuthorizeEditorDataConditionDTOType)`

SetType sets Type field to given value.


### GetConditions

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetConditions() []AuthorizeEditorDataConditionDTO`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetConditionsOk() (*[]AuthorizeEditorDataConditionDTO, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizeEditorDataConditionsOrConditionDTO) SetConditions(v []AuthorizeEditorDataConditionDTO)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


