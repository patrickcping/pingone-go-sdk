# AuthorizeEditorDataRulesEffectSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataRulesEffectSettingsDTOType**](EnumAuthorizeEditorDataRulesEffectSettingsDTOType.md) |  | 
**Condition** | [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataRulesEffectSettingsDTO

`func NewAuthorizeEditorDataRulesEffectSettingsDTO(type_ EnumAuthorizeEditorDataRulesEffectSettingsDTOType, condition AuthorizeEditorDataConditionDTO, ) *AuthorizeEditorDataRulesEffectSettingsDTO`

NewAuthorizeEditorDataRulesEffectSettingsDTO instantiates a new AuthorizeEditorDataRulesEffectSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataRulesEffectSettingsDTOWithDefaults

`func NewAuthorizeEditorDataRulesEffectSettingsDTOWithDefaults() *AuthorizeEditorDataRulesEffectSettingsDTO`

NewAuthorizeEditorDataRulesEffectSettingsDTOWithDefaults instantiates a new AuthorizeEditorDataRulesEffectSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) GetType() EnumAuthorizeEditorDataRulesEffectSettingsDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) GetTypeOk() (*EnumAuthorizeEditorDataRulesEffectSettingsDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) SetType(v EnumAuthorizeEditorDataRulesEffectSettingsDTOType)`

SetType sets Type field to given value.


### GetCondition

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataRulesEffectSettingsDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


