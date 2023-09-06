# UserIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A mutable string that identifies the external identity provider used to authenticate the user. If not provided, PingOne is the identity provider. This attribute is required if the identity provider is authoritative for just-in-time user provisioning. | [optional] 
**Type** | Pointer to [**EnumIdentityProvider**](EnumIdentityProvider.md) |  | [optional] [default to ENUMIDENTITYPROVIDER_PING_ONE]

## Methods

### NewUserIdentityProvider

`func NewUserIdentityProvider() *UserIdentityProvider`

NewUserIdentityProvider instantiates a new UserIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityProviderWithDefaults

`func NewUserIdentityProviderWithDefaults() *UserIdentityProvider`

NewUserIdentityProviderWithDefaults instantiates a new UserIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserIdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserIdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserIdentityProvider) GetType() EnumIdentityProvider`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserIdentityProvider) GetTypeOk() (*EnumIdentityProvider, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserIdentityProvider) SetType(v EnumIdentityProvider)`

SetType sets Type field to given value.

### HasType

`func (o *UserIdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


