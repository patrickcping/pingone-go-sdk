# AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumConcurrentRequests** | Pointer to **int32** |  | [optional] 
**MaximumRequestsPerSecond** | Pointer to **float64** |  | [optional] 
**TimeoutMilliseconds** | Pointer to **int32** |  | [optional] 
**Channel** | [**EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel**](EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel.md) |  | 
**Code** | [**EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode**](EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode.md) |  | 
**Capability** | **string** |  | 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**InputMappings** | [**[]AuthorizeEditorDataInputMappingDTO**](AuthorizeEditorDataInputMappingDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO

`func NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO(channel EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel, code EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode, capability string, inputMappings []AuthorizeEditorDataInputMappingDTO, ) *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO`

NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO instantiates a new AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOWithDefaults

`func NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOWithDefaults() *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO`

NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOWithDefaults instantiates a new AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetMaximumConcurrentRequests() int32`

GetMaximumConcurrentRequests returns the MaximumConcurrentRequests field if non-nil, zero value otherwise.

### GetMaximumConcurrentRequestsOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetMaximumConcurrentRequestsOk() (*int32, bool)`

GetMaximumConcurrentRequestsOk returns a tuple with the MaximumConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetMaximumConcurrentRequests(v int32)`

SetMaximumConcurrentRequests sets MaximumConcurrentRequests field to given value.

### HasMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) HasMaximumConcurrentRequests() bool`

HasMaximumConcurrentRequests returns a boolean if a field has been set.

### GetMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetMaximumRequestsPerSecond() float64`

GetMaximumRequestsPerSecond returns the MaximumRequestsPerSecond field if non-nil, zero value otherwise.

### GetMaximumRequestsPerSecondOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetMaximumRequestsPerSecondOk() (*float64, bool)`

GetMaximumRequestsPerSecondOk returns a tuple with the MaximumRequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetMaximumRequestsPerSecond(v float64)`

SetMaximumRequestsPerSecond sets MaximumRequestsPerSecond field to given value.

### HasMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) HasMaximumRequestsPerSecond() bool`

HasMaximumRequestsPerSecond returns a boolean if a field has been set.

### GetTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetTimeoutMilliseconds() int32`

GetTimeoutMilliseconds returns the TimeoutMilliseconds field if non-nil, zero value otherwise.

### GetTimeoutMillisecondsOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetTimeoutMillisecondsOk() (*int32, bool)`

GetTimeoutMillisecondsOk returns a tuple with the TimeoutMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetTimeoutMilliseconds(v int32)`

SetTimeoutMilliseconds sets TimeoutMilliseconds field to given value.

### HasTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) HasTimeoutMilliseconds() bool`

HasTimeoutMilliseconds returns a boolean if a field has been set.

### GetChannel

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetChannel() EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetChannelOk() (*EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetChannel(v EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel)`

SetChannel sets Channel field to given value.


### GetCode

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetCode() EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetCodeOk() (*EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetCode(v EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode)`

SetCode sets Code field to given value.


### GetCapability

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSchemaVersion

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetInputMappings

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetInputMappings() []AuthorizeEditorDataInputMappingDTO`

GetInputMappings returns the InputMappings field if non-nil, zero value otherwise.

### GetInputMappingsOk

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) GetInputMappingsOk() (*[]AuthorizeEditorDataInputMappingDTO, bool)`

GetInputMappingsOk returns a tuple with the InputMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMappings

`func (o *AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO) SetInputMappings(v []AuthorizeEditorDataInputMappingDTO)`

SetInputMappings sets InputMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


