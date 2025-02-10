# AuthorizeEditorDataConnectorsConnectorTemplateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Code** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConnectorCapability** | Pointer to [**[]AuthorizeEditorDataConnectorsConnectorCapability**](AuthorizeEditorDataConnectorsConnectorCapability.md) |  | [optional] 
**Capabilities** | Pointer to [**[]AuthorizeEditorDataConnectorsConnectorCapability**](AuthorizeEditorDataConnectorsConnectorCapability.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataConnectorsConnectorTemplateDTO

`func NewAuthorizeEditorDataConnectorsConnectorTemplateDTO() *AuthorizeEditorDataConnectorsConnectorTemplateDTO`

NewAuthorizeEditorDataConnectorsConnectorTemplateDTO instantiates a new AuthorizeEditorDataConnectorsConnectorTemplateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataConnectorsConnectorTemplateDTOWithDefaults

`func NewAuthorizeEditorDataConnectorsConnectorTemplateDTOWithDefaults() *AuthorizeEditorDataConnectorsConnectorTemplateDTO`

NewAuthorizeEditorDataConnectorsConnectorTemplateDTOWithDefaults instantiates a new AuthorizeEditorDataConnectorsConnectorTemplateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetCode

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetChannel

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetName

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectorCapability

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetConnectorCapability() []AuthorizeEditorDataConnectorsConnectorCapability`

GetConnectorCapability returns the ConnectorCapability field if non-nil, zero value otherwise.

### GetConnectorCapabilityOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetConnectorCapabilityOk() (*[]AuthorizeEditorDataConnectorsConnectorCapability, bool)`

GetConnectorCapabilityOk returns a tuple with the ConnectorCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCapability

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetConnectorCapability(v []AuthorizeEditorDataConnectorsConnectorCapability)`

SetConnectorCapability sets ConnectorCapability field to given value.

### HasConnectorCapability

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasConnectorCapability() bool`

HasConnectorCapability returns a boolean if a field has been set.

### GetCapabilities

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCapabilities() []AuthorizeEditorDataConnectorsConnectorCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCapabilitiesOk() (*[]AuthorizeEditorDataConnectorsConnectorCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetCapabilities(v []AuthorizeEditorDataConnectorsConnectorCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


