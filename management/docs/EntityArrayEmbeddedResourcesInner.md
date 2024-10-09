# EntityArrayEmbeddedResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessTokenValiditySeconds** | Pointer to **int32** | An integer that specifies the number of seconds that the access token is valid. If a value is not specified, the default is 3600. The minimum value is 300 seconds (5 minutes); the maximum value is 2592000 seconds (30 days). | [optional] 
**ApplicationPermissionsSettings** | Pointer to [**ResourceApplicationPermissionsSettings**](ResourceApplicationPermissionsSettings.md) |  | [optional] 
**Audience** | Pointer to **string** | A string that specifies a URL without a fragment or &#x60;@ObjectName&#x60; and must not contain &#x60;pingone&#x60; or &#x60;pingidentity&#x60; (for example, &#x60;https://api.bxretail.org&#x60;). If a URL is not specified, the resource name is used. | [optional] 
**ClientId** | Pointer to **string** | (Required when &#x60;clientSecret&#x60; is specified.) Supported only for the &#x60;POST&#x60; operation. This is the UUID of an external resource that is being migrated to PingOne. The UUID must be a minimum of 8 alpha-numeric characters, and must be globally unique in PingOne. | [optional] 
**ClientSecret** | Pointer to **string** | (Required when clientId is specified.) Supported only for the POST operation. This is the client secret associated with clientId for an external resource that is being migrated to PingOne. This must be a minimum of 8 alpha-numeric characters. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The application resource&#39;s description. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Resource** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Name** | **string** | The application resource name. The name value must be unique. | 
**IntrospectEndpointAuthMethod** | Pointer to [**EnumResourceIntrospectEndpointAuthMethod**](EnumResourceIntrospectEndpointAuthMethod.md) |  | [optional] 
**Type** | Pointer to [**EnumResourceType**](EnumResourceType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Parent** | Pointer to [**ResourceApplicationResourceParent**](ResourceApplicationResourceParent.md) |  | [optional] 

## Methods

### NewEntityArrayEmbeddedResourcesInner

`func NewEntityArrayEmbeddedResourcesInner(name string, ) *EntityArrayEmbeddedResourcesInner`

NewEntityArrayEmbeddedResourcesInner instantiates a new EntityArrayEmbeddedResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedResourcesInnerWithDefaults

`func NewEntityArrayEmbeddedResourcesInnerWithDefaults() *EntityArrayEmbeddedResourcesInner`

NewEntityArrayEmbeddedResourcesInnerWithDefaults instantiates a new EntityArrayEmbeddedResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntityArrayEmbeddedResourcesInner) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntityArrayEmbeddedResourcesInner) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntityArrayEmbeddedResourcesInner) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntityArrayEmbeddedResourcesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *EntityArrayEmbeddedResourcesInner) GetAccessTokenValiditySeconds() int32`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *EntityArrayEmbeddedResourcesInner) GetAccessTokenValiditySecondsOk() (*int32, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *EntityArrayEmbeddedResourcesInner) SetAccessTokenValiditySeconds(v int32)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *EntityArrayEmbeddedResourcesInner) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetApplicationPermissionsSettings

`func (o *EntityArrayEmbeddedResourcesInner) GetApplicationPermissionsSettings() ResourceApplicationPermissionsSettings`

GetApplicationPermissionsSettings returns the ApplicationPermissionsSettings field if non-nil, zero value otherwise.

### GetApplicationPermissionsSettingsOk

`func (o *EntityArrayEmbeddedResourcesInner) GetApplicationPermissionsSettingsOk() (*ResourceApplicationPermissionsSettings, bool)`

GetApplicationPermissionsSettingsOk returns a tuple with the ApplicationPermissionsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPermissionsSettings

`func (o *EntityArrayEmbeddedResourcesInner) SetApplicationPermissionsSettings(v ResourceApplicationPermissionsSettings)`

SetApplicationPermissionsSettings sets ApplicationPermissionsSettings field to given value.

### HasApplicationPermissionsSettings

`func (o *EntityArrayEmbeddedResourcesInner) HasApplicationPermissionsSettings() bool`

HasApplicationPermissionsSettings returns a boolean if a field has been set.

### GetAudience

`func (o *EntityArrayEmbeddedResourcesInner) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *EntityArrayEmbeddedResourcesInner) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *EntityArrayEmbeddedResourcesInner) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *EntityArrayEmbeddedResourcesInner) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClientId

`func (o *EntityArrayEmbeddedResourcesInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EntityArrayEmbeddedResourcesInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EntityArrayEmbeddedResourcesInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EntityArrayEmbeddedResourcesInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *EntityArrayEmbeddedResourcesInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *EntityArrayEmbeddedResourcesInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *EntityArrayEmbeddedResourcesInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *EntityArrayEmbeddedResourcesInner) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityArrayEmbeddedResourcesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityArrayEmbeddedResourcesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityArrayEmbeddedResourcesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityArrayEmbeddedResourcesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *EntityArrayEmbeddedResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *EntityArrayEmbeddedResourcesInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedResourcesInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedResourcesInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedResourcesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetResource

`func (o *EntityArrayEmbeddedResourcesInner) GetResource() IdentityProviderAttributeIdentityProvider`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EntityArrayEmbeddedResourcesInner) GetResourceOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EntityArrayEmbeddedResourcesInner) SetResource(v IdentityProviderAttributeIdentityProvider)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EntityArrayEmbeddedResourcesInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetId

`func (o *EntityArrayEmbeddedResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedResourcesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedResourcesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntityArrayEmbeddedResourcesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedResourcesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedResourcesInner) SetName(v string)`

SetName sets Name field to given value.


### GetIntrospectEndpointAuthMethod

`func (o *EntityArrayEmbeddedResourcesInner) GetIntrospectEndpointAuthMethod() EnumResourceIntrospectEndpointAuthMethod`

GetIntrospectEndpointAuthMethod returns the IntrospectEndpointAuthMethod field if non-nil, zero value otherwise.

### GetIntrospectEndpointAuthMethodOk

`func (o *EntityArrayEmbeddedResourcesInner) GetIntrospectEndpointAuthMethodOk() (*EnumResourceIntrospectEndpointAuthMethod, bool)`

GetIntrospectEndpointAuthMethodOk returns a tuple with the IntrospectEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectEndpointAuthMethod

`func (o *EntityArrayEmbeddedResourcesInner) SetIntrospectEndpointAuthMethod(v EnumResourceIntrospectEndpointAuthMethod)`

SetIntrospectEndpointAuthMethod sets IntrospectEndpointAuthMethod field to given value.

### HasIntrospectEndpointAuthMethod

`func (o *EntityArrayEmbeddedResourcesInner) HasIntrospectEndpointAuthMethod() bool`

HasIntrospectEndpointAuthMethod returns a boolean if a field has been set.

### GetType

`func (o *EntityArrayEmbeddedResourcesInner) GetType() EnumResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityArrayEmbeddedResourcesInner) GetTypeOk() (*EnumResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityArrayEmbeddedResourcesInner) SetType(v EnumResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *EntityArrayEmbeddedResourcesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EntityArrayEmbeddedResourcesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityArrayEmbeddedResourcesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityArrayEmbeddedResourcesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityArrayEmbeddedResourcesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetParent

`func (o *EntityArrayEmbeddedResourcesInner) GetParent() ResourceApplicationResourceParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EntityArrayEmbeddedResourcesInner) GetParentOk() (*ResourceApplicationResourceParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EntityArrayEmbeddedResourcesInner) SetParent(v ResourceApplicationResourceParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EntityArrayEmbeddedResourcesInner) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


