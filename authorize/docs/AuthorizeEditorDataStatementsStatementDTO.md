# AuthorizeEditorDataStatementsStatementDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**AppliesTo** | [**EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo**](EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo.md) |  | 
**AppliesIf** | [**EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf**](EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf.md) |  | 
**Payload** | **string** |  | 
**Obligatory** | Pointer to **bool** |  | [optional] 
**Attributes** | [**[]AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataStatementsStatementDTO

`func NewAuthorizeEditorDataStatementsStatementDTO(name string, code string, appliesTo EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo, appliesIf EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf, payload string, attributes []AuthorizeEditorDataReferenceObjectDTO, ) *AuthorizeEditorDataStatementsStatementDTO`

NewAuthorizeEditorDataStatementsStatementDTO instantiates a new AuthorizeEditorDataStatementsStatementDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataStatementsStatementDTOWithDefaults

`func NewAuthorizeEditorDataStatementsStatementDTOWithDefaults() *AuthorizeEditorDataStatementsStatementDTO`

NewAuthorizeEditorDataStatementsStatementDTOWithDefaults instantiates a new AuthorizeEditorDataStatementsStatementDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetCode(v string)`

SetCode sets Code field to given value.


### GetAppliesTo

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAppliesTo() EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAppliesToOk() (*EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetAppliesTo(v EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo)`

SetAppliesTo sets AppliesTo field to given value.


### GetAppliesIf

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAppliesIf() EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf`

GetAppliesIf returns the AppliesIf field if non-nil, zero value otherwise.

### GetAppliesIfOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAppliesIfOk() (*EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf, bool)`

GetAppliesIfOk returns a tuple with the AppliesIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesIf

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetAppliesIf(v EnumAuthorizeEditorDataStatementsStatementDTOAppliesIf)`

SetAppliesIf sets AppliesIf field to given value.


### GetPayload

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetObligatory

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetObligatory() bool`

GetObligatory returns the Obligatory field if non-nil, zero value otherwise.

### GetObligatoryOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetObligatoryOk() (*bool, bool)`

GetObligatoryOk returns a tuple with the Obligatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligatory

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetObligatory(v bool)`

SetObligatory sets Obligatory field to given value.

### HasObligatory

`func (o *AuthorizeEditorDataStatementsStatementDTO) HasObligatory() bool`

HasObligatory returns a boolean if a field has been set.

### GetAttributes

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAttributes() []AuthorizeEditorDataReferenceObjectDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthorizeEditorDataStatementsStatementDTO) GetAttributesOk() (*[]AuthorizeEditorDataReferenceObjectDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthorizeEditorDataStatementsStatementDTO) SetAttributes(v []AuthorizeEditorDataReferenceObjectDTO)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


