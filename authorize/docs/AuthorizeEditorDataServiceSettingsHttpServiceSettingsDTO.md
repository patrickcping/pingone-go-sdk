# AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumConcurrentRequests** | Pointer to **int32** |  | [optional] 
**MaximumRequestsPerSecond** | Pointer to **float64** |  | [optional] 
**TimeoutMilliseconds** | Pointer to **int32** |  | [optional] 
**Url** | **string** |  | 
**Verb** | [**EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb**](EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb.md) |  | 
**Body** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]AuthorizeEditorDataHttpRequestHeaderDTO**](AuthorizeEditorDataHttpRequestHeaderDTO.md) |  | [optional] 
**Authentication** | [**AuthorizeEditorDataAuthenticationDTO**](AuthorizeEditorDataAuthenticationDTO.md) |  | 
**TlsSettings** | [**AuthorizeEditorDataTlsSettingsDTO**](AuthorizeEditorDataTlsSettingsDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO

`func NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO(url string, verb EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb, authentication AuthorizeEditorDataAuthenticationDTO, tlsSettings AuthorizeEditorDataTlsSettingsDTO, ) *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO instantiates a new AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOWithDefaults

`func NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOWithDefaults() *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

NewAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOWithDefaults instantiates a new AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetMaximumConcurrentRequests() int32`

GetMaximumConcurrentRequests returns the MaximumConcurrentRequests field if non-nil, zero value otherwise.

### GetMaximumConcurrentRequestsOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetMaximumConcurrentRequestsOk() (*int32, bool)`

GetMaximumConcurrentRequestsOk returns a tuple with the MaximumConcurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetMaximumConcurrentRequests(v int32)`

SetMaximumConcurrentRequests sets MaximumConcurrentRequests field to given value.

### HasMaximumConcurrentRequests

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasMaximumConcurrentRequests() bool`

HasMaximumConcurrentRequests returns a boolean if a field has been set.

### GetMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetMaximumRequestsPerSecond() float64`

GetMaximumRequestsPerSecond returns the MaximumRequestsPerSecond field if non-nil, zero value otherwise.

### GetMaximumRequestsPerSecondOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetMaximumRequestsPerSecondOk() (*float64, bool)`

GetMaximumRequestsPerSecondOk returns a tuple with the MaximumRequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetMaximumRequestsPerSecond(v float64)`

SetMaximumRequestsPerSecond sets MaximumRequestsPerSecond field to given value.

### HasMaximumRequestsPerSecond

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasMaximumRequestsPerSecond() bool`

HasMaximumRequestsPerSecond returns a boolean if a field has been set.

### GetTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetTimeoutMilliseconds() int32`

GetTimeoutMilliseconds returns the TimeoutMilliseconds field if non-nil, zero value otherwise.

### GetTimeoutMillisecondsOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetTimeoutMillisecondsOk() (*int32, bool)`

GetTimeoutMillisecondsOk returns a tuple with the TimeoutMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetTimeoutMilliseconds(v int32)`

SetTimeoutMilliseconds sets TimeoutMilliseconds field to given value.

### HasTimeoutMilliseconds

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasTimeoutMilliseconds() bool`

HasTimeoutMilliseconds returns a boolean if a field has been set.

### GetUrl

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerb

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetVerb() EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetVerbOk() (*EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetVerb(v EnumAuthorizeEditorDataServiceSettingsHttpServiceSettingsDTOVerb)`

SetVerb sets Verb field to given value.


### GetBody

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContentType

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHeaders

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetHeaders() []AuthorizeEditorDataHttpRequestHeaderDTO`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetHeadersOk() (*[]AuthorizeEditorDataHttpRequestHeaderDTO, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetHeaders(v []AuthorizeEditorDataHttpRequestHeaderDTO)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAuthentication

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetAuthentication() AuthorizeEditorDataAuthenticationDTO`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetAuthenticationOk() (*AuthorizeEditorDataAuthenticationDTO, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetAuthentication(v AuthorizeEditorDataAuthenticationDTO)`

SetAuthentication sets Authentication field to given value.


### GetTlsSettings

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetTlsSettings() AuthorizeEditorDataTlsSettingsDTO`

GetTlsSettings returns the TlsSettings field if non-nil, zero value otherwise.

### GetTlsSettingsOk

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) GetTlsSettingsOk() (*AuthorizeEditorDataTlsSettingsDTO, bool)`

GetTlsSettingsOk returns a tuple with the TlsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSettings

`func (o *AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO) SetTlsSettings(v AuthorizeEditorDataTlsSettingsDTO)`

SetTlsSettings sets TlsSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


