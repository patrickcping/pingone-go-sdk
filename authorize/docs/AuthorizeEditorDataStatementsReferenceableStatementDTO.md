# AuthorizeEditorDataStatementsReferenceableStatementDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | **string** | The resource&#39;s unique identifier | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**AppliesTo** | [**EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo**](EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo.md) |  | 
**AppliesIf** | [**EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf**](EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf.md) |  | 
**Payload** | **string** |  | 
**Obligatory** | Pointer to **bool** |  | [optional] 
**Attributes** | [**[]AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Version** | **string** |  | 

## Methods

### NewAuthorizeEditorDataStatementsReferenceableStatementDTO

`func NewAuthorizeEditorDataStatementsReferenceableStatementDTO(id string, name string, code string, appliesTo EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo, appliesIf EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf, payload string, attributes []AuthorizeEditorDataReferenceObjectDTO, version string, ) *AuthorizeEditorDataStatementsReferenceableStatementDTO`

NewAuthorizeEditorDataStatementsReferenceableStatementDTO instantiates a new AuthorizeEditorDataStatementsReferenceableStatementDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataStatementsReferenceableStatementDTOWithDefaults

`func NewAuthorizeEditorDataStatementsReferenceableStatementDTOWithDefaults() *AuthorizeEditorDataStatementsReferenceableStatementDTO`

NewAuthorizeEditorDataStatementsReferenceableStatementDTOWithDefaults instantiates a new AuthorizeEditorDataStatementsReferenceableStatementDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetCode(v string)`

SetCode sets Code field to given value.


### GetAppliesTo

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAppliesTo() EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAppliesToOk() (*EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetAppliesTo(v EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesTo)`

SetAppliesTo sets AppliesTo field to given value.


### GetAppliesIf

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAppliesIf() EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf`

GetAppliesIf returns the AppliesIf field if non-nil, zero value otherwise.

### GetAppliesIfOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAppliesIfOk() (*EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf, bool)`

GetAppliesIfOk returns a tuple with the AppliesIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesIf

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetAppliesIf(v EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf)`

SetAppliesIf sets AppliesIf field to given value.


### GetPayload

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetObligatory

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetObligatory() bool`

GetObligatory returns the Obligatory field if non-nil, zero value otherwise.

### GetObligatoryOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetObligatoryOk() (*bool, bool)`

GetObligatoryOk returns a tuple with the Obligatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligatory

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetObligatory(v bool)`

SetObligatory sets Obligatory field to given value.

### HasObligatory

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) HasObligatory() bool`

HasObligatory returns a boolean if a field has been set.

### GetAttributes

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAttributes() []AuthorizeEditorDataReferenceObjectDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetAttributesOk() (*[]AuthorizeEditorDataReferenceObjectDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetAttributes(v []AuthorizeEditorDataReferenceObjectDTO)`

SetAttributes sets Attributes field to given value.


### GetVersion

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthorizeEditorDataStatementsReferenceableStatementDTO) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


