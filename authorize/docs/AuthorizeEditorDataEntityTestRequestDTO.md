# AuthorizeEditorDataEntityTestRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**[]AuthorizeEditorDataEntityTestingParameterDTO**](AuthorizeEditorDataEntityTestingParameterDTO.md) |  | 
**UserContext** | Pointer to [**AuthorizeEditorDataEntityTestingUserContextDTO**](AuthorizeEditorDataEntityTestingUserContextDTO.md) |  | [optional] 
**Overrides** | Pointer to [**[]AuthorizeEditorDataEntityTestingOverrideDTO**](AuthorizeEditorDataEntityTestingOverrideDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataEntityTestRequestDTO

`func NewAuthorizeEditorDataEntityTestRequestDTO(parameters []AuthorizeEditorDataEntityTestingParameterDTO, ) *AuthorizeEditorDataEntityTestRequestDTO`

NewAuthorizeEditorDataEntityTestRequestDTO instantiates a new AuthorizeEditorDataEntityTestRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataEntityTestRequestDTOWithDefaults

`func NewAuthorizeEditorDataEntityTestRequestDTOWithDefaults() *AuthorizeEditorDataEntityTestRequestDTO`

NewAuthorizeEditorDataEntityTestRequestDTOWithDefaults instantiates a new AuthorizeEditorDataEntityTestRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetParameters() []AuthorizeEditorDataEntityTestingParameterDTO`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetParametersOk() (*[]AuthorizeEditorDataEntityTestingParameterDTO, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AuthorizeEditorDataEntityTestRequestDTO) SetParameters(v []AuthorizeEditorDataEntityTestingParameterDTO)`

SetParameters sets Parameters field to given value.


### GetUserContext

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetUserContext() AuthorizeEditorDataEntityTestingUserContextDTO`

GetUserContext returns the UserContext field if non-nil, zero value otherwise.

### GetUserContextOk

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetUserContextOk() (*AuthorizeEditorDataEntityTestingUserContextDTO, bool)`

GetUserContextOk returns a tuple with the UserContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserContext

`func (o *AuthorizeEditorDataEntityTestRequestDTO) SetUserContext(v AuthorizeEditorDataEntityTestingUserContextDTO)`

SetUserContext sets UserContext field to given value.

### HasUserContext

`func (o *AuthorizeEditorDataEntityTestRequestDTO) HasUserContext() bool`

HasUserContext returns a boolean if a field has been set.

### GetOverrides

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetOverrides() []AuthorizeEditorDataEntityTestingOverrideDTO`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *AuthorizeEditorDataEntityTestRequestDTO) GetOverridesOk() (*[]AuthorizeEditorDataEntityTestingOverrideDTO, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *AuthorizeEditorDataEntityTestRequestDTO) SetOverrides(v []AuthorizeEditorDataEntityTestingOverrideDTO)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *AuthorizeEditorDataEntityTestRequestDTO) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


