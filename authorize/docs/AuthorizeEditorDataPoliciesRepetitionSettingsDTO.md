# AuthorizeEditorDataPoliciesRepetitionSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Decision** | [**EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision**](EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision.md) |  | 

## Methods

### NewAuthorizeEditorDataPoliciesRepetitionSettingsDTO

`func NewAuthorizeEditorDataPoliciesRepetitionSettingsDTO(source AuthorizeEditorDataReferenceObjectDTO, decision EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision, ) *AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

NewAuthorizeEditorDataPoliciesRepetitionSettingsDTO instantiates a new AuthorizeEditorDataPoliciesRepetitionSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesRepetitionSettingsDTOWithDefaults

`func NewAuthorizeEditorDataPoliciesRepetitionSettingsDTOWithDefaults() *AuthorizeEditorDataPoliciesRepetitionSettingsDTO`

NewAuthorizeEditorDataPoliciesRepetitionSettingsDTOWithDefaults instantiates a new AuthorizeEditorDataPoliciesRepetitionSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) GetSource() AuthorizeEditorDataReferenceObjectDTO`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) GetSourceOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) SetSource(v AuthorizeEditorDataReferenceObjectDTO)`

SetSource sets Source field to given value.


### GetDecision

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) GetDecision() EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) GetDecisionOk() (*EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AuthorizeEditorDataPoliciesRepetitionSettingsDTO) SetDecision(v EnumAuthorizeEditorDataPoliciesRepetitionSettingsDTODecision)`

SetDecision sets Decision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


