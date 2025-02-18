# DecisionEndpointAuthorizationVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the ID of the Authorization Version deployed to this endpoint. Versioning allows independent development and deployment of Policies If omitted, the endpoint always uses the latest policy version available from the policy editor service. | [optional] 
**Href** | Pointer to **string** | A string that specifies the request URL for the authorization version endpoint. | [optional] 
**Title** | Pointer to **string** | A string that specifies the title for the authorization version response. | [optional] 
**Type** | Pointer to **string** | A string that specifies the content type for the authorization version response. | [optional] 

## Methods

### NewDecisionEndpointAuthorizationVersion

`func NewDecisionEndpointAuthorizationVersion() *DecisionEndpointAuthorizationVersion`

NewDecisionEndpointAuthorizationVersion instantiates a new DecisionEndpointAuthorizationVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionEndpointAuthorizationVersionWithDefaults

`func NewDecisionEndpointAuthorizationVersionWithDefaults() *DecisionEndpointAuthorizationVersion`

NewDecisionEndpointAuthorizationVersionWithDefaults instantiates a new DecisionEndpointAuthorizationVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionEndpointAuthorizationVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionEndpointAuthorizationVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionEndpointAuthorizationVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionEndpointAuthorizationVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *DecisionEndpointAuthorizationVersion) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DecisionEndpointAuthorizationVersion) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DecisionEndpointAuthorizationVersion) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DecisionEndpointAuthorizationVersion) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetTitle

`func (o *DecisionEndpointAuthorizationVersion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DecisionEndpointAuthorizationVersion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DecisionEndpointAuthorizationVersion) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DecisionEndpointAuthorizationVersion) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DecisionEndpointAuthorizationVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecisionEndpointAuthorizationVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecisionEndpointAuthorizationVersion) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DecisionEndpointAuthorizationVersion) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


