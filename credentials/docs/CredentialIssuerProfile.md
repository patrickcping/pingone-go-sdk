# CredentialIssuerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationInstance** | Pointer to [**CredentialIssuerProfileApplicationInstance**](CredentialIssuerProfileApplicationInstance.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | A string that specifies the date and time the issuer profile was created. | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) of the credential issuer. | [optional] 
**Name** | **string** | The name of the credential issuer. This will be included in credentials issued. | 
**UpdatedAt** | Pointer to **string** | A string that specifies the date and time the credential issuer profile was last updated; can be null. | [optional] 
**SiteUrl** | Pointer to **string** | A string that specifies the base URL associated with the credential issuer. | [optional] 
**CustomEmailTemplate** | Pointer to **string** | A string that specifies the default notification template used in credential issuance notifications. Deprecated. | [optional] 

## Methods

### NewCredentialIssuerProfile

`func NewCredentialIssuerProfile(name string, ) *CredentialIssuerProfile`

NewCredentialIssuerProfile instantiates a new CredentialIssuerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuerProfileWithDefaults

`func NewCredentialIssuerProfileWithDefaults() *CredentialIssuerProfile`

NewCredentialIssuerProfileWithDefaults instantiates a new CredentialIssuerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationInstance

`func (o *CredentialIssuerProfile) GetApplicationInstance() CredentialIssuerProfileApplicationInstance`

GetApplicationInstance returns the ApplicationInstance field if non-nil, zero value otherwise.

### GetApplicationInstanceOk

`func (o *CredentialIssuerProfile) GetApplicationInstanceOk() (*CredentialIssuerProfileApplicationInstance, bool)`

GetApplicationInstanceOk returns a tuple with the ApplicationInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInstance

`func (o *CredentialIssuerProfile) SetApplicationInstance(v CredentialIssuerProfileApplicationInstance)`

SetApplicationInstance sets ApplicationInstance field to given value.

### HasApplicationInstance

`func (o *CredentialIssuerProfile) HasApplicationInstance() bool`

HasApplicationInstance returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialIssuerProfile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialIssuerProfile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialIssuerProfile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialIssuerProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CredentialIssuerProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialIssuerProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialIssuerProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialIssuerProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CredentialIssuerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialIssuerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialIssuerProfile) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *CredentialIssuerProfile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialIssuerProfile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialIssuerProfile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialIssuerProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSiteUrl

`func (o *CredentialIssuerProfile) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *CredentialIssuerProfile) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *CredentialIssuerProfile) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.

### HasSiteUrl

`func (o *CredentialIssuerProfile) HasSiteUrl() bool`

HasSiteUrl returns a boolean if a field has been set.

### GetCustomEmailTemplate

`func (o *CredentialIssuerProfile) GetCustomEmailTemplate() string`

GetCustomEmailTemplate returns the CustomEmailTemplate field if non-nil, zero value otherwise.

### GetCustomEmailTemplateOk

`func (o *CredentialIssuerProfile) GetCustomEmailTemplateOk() (*string, bool)`

GetCustomEmailTemplateOk returns a tuple with the CustomEmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmailTemplate

`func (o *CredentialIssuerProfile) SetCustomEmailTemplate(v string)`

SetCustomEmailTemplate sets CustomEmailTemplate field to given value.

### HasCustomEmailTemplate

`func (o *CredentialIssuerProfile) HasCustomEmailTemplate() bool`

HasCustomEmailTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


