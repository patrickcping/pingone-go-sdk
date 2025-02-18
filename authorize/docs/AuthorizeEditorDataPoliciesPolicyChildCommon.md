# AuthorizeEditorDataPoliciesPolicyChildCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataPoliciesPolicyChildCommonType**](EnumAuthorizeEditorDataPoliciesPolicyChildCommonType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Statements** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**Value** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataPoliciesPolicyChildCommon

`func NewAuthorizeEditorDataPoliciesPolicyChildCommon(type_ EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, ) *AuthorizeEditorDataPoliciesPolicyChildCommon`

NewAuthorizeEditorDataPoliciesPolicyChildCommon instantiates a new AuthorizeEditorDataPoliciesPolicyChildCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataPoliciesPolicyChildCommonWithDefaults

`func NewAuthorizeEditorDataPoliciesPolicyChildCommonWithDefaults() *AuthorizeEditorDataPoliciesPolicyChildCommon`

NewAuthorizeEditorDataPoliciesPolicyChildCommonWithDefaults instantiates a new AuthorizeEditorDataPoliciesPolicyChildCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetType() EnumAuthorizeEditorDataPoliciesPolicyChildCommonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetTypeOk() (*EnumAuthorizeEditorDataPoliciesPolicyChildCommonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetType(v EnumAuthorizeEditorDataPoliciesPolicyChildCommonType)`

SetType sets Type field to given value.


### GetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetStatements() []map[string]interface{}`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetStatementsOk() (*[]map[string]interface{}, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetStatements(v []map[string]interface{})`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetValue() AuthorizeEditorDataReferenceObjectDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) SetValue(v AuthorizeEditorDataReferenceObjectDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthorizeEditorDataPoliciesPolicyChildCommon) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


