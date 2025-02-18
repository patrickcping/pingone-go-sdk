# AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataAuthenticationDTOType**](EnumAuthorizeEditorDataAuthenticationDTOType.md) |  | 
**TokenEndpoint** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Scope** | **string** |  | 

## Methods

### NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO

`func NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO(type_ EnumAuthorizeEditorDataAuthenticationDTOType, tokenEndpoint string, clientId string, clientSecret AuthorizeEditorDataReferenceObjectDTO, scope string, ) *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO`

NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO instantiates a new AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTOWithDefaults

`func NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTOWithDefaults() *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO`

NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTOWithDefaults instantiates a new AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetType() EnumAuthorizeEditorDataAuthenticationDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetTypeOk() (*EnumAuthorizeEditorDataAuthenticationDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetType(v EnumAuthorizeEditorDataAuthenticationDTOType)`

SetType sets Type field to given value.


### GetTokenEndpoint

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetClientId

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientSecret() AuthorizeEditorDataReferenceObjectDTO`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientSecretOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetClientSecret(v AuthorizeEditorDataReferenceObjectDTO)`

SetClientSecret sets ClientSecret field to given value.


### GetScope

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


