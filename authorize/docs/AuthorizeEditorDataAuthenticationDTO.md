# AuthorizeEditorDataAuthenticationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumAuthorizeEditorDataAuthenticationDTOType**](EnumAuthorizeEditorDataAuthenticationDTOType.md) |  | 
**Name** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Password** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**TokenEndpoint** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 
**Scope** | **string** |  | 
**Token** | [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataAuthenticationDTO

`func NewAuthorizeEditorDataAuthenticationDTO(type_ EnumAuthorizeEditorDataAuthenticationDTOType, name AuthorizeEditorDataReferenceObjectDTO, password AuthorizeEditorDataReferenceObjectDTO, tokenEndpoint string, clientId string, clientSecret AuthorizeEditorDataReferenceObjectDTO, scope string, token AuthorizeEditorDataReferenceObjectDTO, ) *AuthorizeEditorDataAuthenticationDTO`

NewAuthorizeEditorDataAuthenticationDTO instantiates a new AuthorizeEditorDataAuthenticationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataAuthenticationDTOWithDefaults

`func NewAuthorizeEditorDataAuthenticationDTOWithDefaults() *AuthorizeEditorDataAuthenticationDTO`

NewAuthorizeEditorDataAuthenticationDTOWithDefaults instantiates a new AuthorizeEditorDataAuthenticationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizeEditorDataAuthenticationDTO) GetType() EnumAuthorizeEditorDataAuthenticationDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetTypeOk() (*EnumAuthorizeEditorDataAuthenticationDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataAuthenticationDTO) SetType(v EnumAuthorizeEditorDataAuthenticationDTOType)`

SetType sets Type field to given value.


### GetName

`func (o *AuthorizeEditorDataAuthenticationDTO) GetName() AuthorizeEditorDataReferenceObjectDTO`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetNameOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataAuthenticationDTO) SetName(v AuthorizeEditorDataReferenceObjectDTO)`

SetName sets Name field to given value.


### GetPassword

`func (o *AuthorizeEditorDataAuthenticationDTO) GetPassword() AuthorizeEditorDataReferenceObjectDTO`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetPasswordOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthorizeEditorDataAuthenticationDTO) SetPassword(v AuthorizeEditorDataReferenceObjectDTO)`

SetPassword sets Password field to given value.


### GetTokenEndpoint

`func (o *AuthorizeEditorDataAuthenticationDTO) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *AuthorizeEditorDataAuthenticationDTO) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetClientId

`func (o *AuthorizeEditorDataAuthenticationDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthorizeEditorDataAuthenticationDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AuthorizeEditorDataAuthenticationDTO) GetClientSecret() AuthorizeEditorDataReferenceObjectDTO`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetClientSecretOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthorizeEditorDataAuthenticationDTO) SetClientSecret(v AuthorizeEditorDataReferenceObjectDTO)`

SetClientSecret sets ClientSecret field to given value.


### GetScope

`func (o *AuthorizeEditorDataAuthenticationDTO) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthorizeEditorDataAuthenticationDTO) SetScope(v string)`

SetScope sets Scope field to given value.


### GetToken

`func (o *AuthorizeEditorDataAuthenticationDTO) GetToken() AuthorizeEditorDataReferenceObjectDTO`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthorizeEditorDataAuthenticationDTO) GetTokenOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthorizeEditorDataAuthenticationDTO) SetToken(v AuthorizeEditorDataReferenceObjectDTO)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


