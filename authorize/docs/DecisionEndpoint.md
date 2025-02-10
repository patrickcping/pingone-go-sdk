# DecisionEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AlternateId** | Pointer to **string** | A string that specifies alternative unique identifier for the endpoint, which provides a method for locating the resource by a known, fixed identifier. | [optional] 
**AuthorizationVersion** | Pointer to [**DecisionEndpointAuthorizationVersion**](DecisionEndpointAuthorizationVersion.md) |  | [optional] 
**Description** | **string** | A string that specifies the description of the policy decision resource. | 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the policy decision resource name. | 
**Owned** | Pointer to **bool** | A boolean that when true restricts modifications of the endpoint to PingOne-owned clients. | [optional] 
**PolicyID** | Pointer to **string** | A string that specifies the ID of the root policy configured for this endpoint. If omitted, the policy editor service decides on a suitable default. | [optional] 
**Provenance** | Pointer to **string** | A string that specifies a machine-readable identifier indicating the provenance of the current configuration. It has no meaning to the Policy Decision Service itself but exists to support integration with other Services | [optional] 
**RecentDecisionsEnabled** | Pointer to **bool** | A boolean that specifies whether to show recent decisions. | [optional] 
**RecentDecisions** | Pointer to [**DecisionEndpointRecentDecisions**](DecisionEndpointRecentDecisions.md) |  | [optional] 
**RecordRecentRequests** | **bool** | A boolean that specifies whether to record a limited history of recent decision requests and responses, which can be queried through a separate API. | 

## Methods

### NewDecisionEndpoint

`func NewDecisionEndpoint(description string, name string, recordRecentRequests bool, ) *DecisionEndpoint`

NewDecisionEndpoint instantiates a new DecisionEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionEndpointWithDefaults

`func NewDecisionEndpointWithDefaults() *DecisionEndpoint`

NewDecisionEndpointWithDefaults instantiates a new DecisionEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DecisionEndpoint) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DecisionEndpoint) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DecisionEndpoint) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DecisionEndpoint) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAlternateId

`func (o *DecisionEndpoint) GetAlternateId() string`

GetAlternateId returns the AlternateId field if non-nil, zero value otherwise.

### GetAlternateIdOk

`func (o *DecisionEndpoint) GetAlternateIdOk() (*string, bool)`

GetAlternateIdOk returns a tuple with the AlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateId

`func (o *DecisionEndpoint) SetAlternateId(v string)`

SetAlternateId sets AlternateId field to given value.

### HasAlternateId

`func (o *DecisionEndpoint) HasAlternateId() bool`

HasAlternateId returns a boolean if a field has been set.

### GetAuthorizationVersion

`func (o *DecisionEndpoint) GetAuthorizationVersion() DecisionEndpointAuthorizationVersion`

GetAuthorizationVersion returns the AuthorizationVersion field if non-nil, zero value otherwise.

### GetAuthorizationVersionOk

`func (o *DecisionEndpoint) GetAuthorizationVersionOk() (*DecisionEndpointAuthorizationVersion, bool)`

GetAuthorizationVersionOk returns a tuple with the AuthorizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationVersion

`func (o *DecisionEndpoint) SetAuthorizationVersion(v DecisionEndpointAuthorizationVersion)`

SetAuthorizationVersion sets AuthorizationVersion field to given value.

### HasAuthorizationVersion

`func (o *DecisionEndpoint) HasAuthorizationVersion() bool`

HasAuthorizationVersion returns a boolean if a field has been set.

### GetDescription

`func (o *DecisionEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DecisionEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DecisionEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DecisionEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DecisionEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetOwned

`func (o *DecisionEndpoint) GetOwned() bool`

GetOwned returns the Owned field if non-nil, zero value otherwise.

### GetOwnedOk

`func (o *DecisionEndpoint) GetOwnedOk() (*bool, bool)`

GetOwnedOk returns a tuple with the Owned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwned

`func (o *DecisionEndpoint) SetOwned(v bool)`

SetOwned sets Owned field to given value.

### HasOwned

`func (o *DecisionEndpoint) HasOwned() bool`

HasOwned returns a boolean if a field has been set.

### GetPolicyID

`func (o *DecisionEndpoint) GetPolicyID() string`

GetPolicyID returns the PolicyID field if non-nil, zero value otherwise.

### GetPolicyIDOk

`func (o *DecisionEndpoint) GetPolicyIDOk() (*string, bool)`

GetPolicyIDOk returns a tuple with the PolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyID

`func (o *DecisionEndpoint) SetPolicyID(v string)`

SetPolicyID sets PolicyID field to given value.

### HasPolicyID

`func (o *DecisionEndpoint) HasPolicyID() bool`

HasPolicyID returns a boolean if a field has been set.

### GetProvenance

`func (o *DecisionEndpoint) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DecisionEndpoint) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DecisionEndpoint) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DecisionEndpoint) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRecentDecisionsEnabled

`func (o *DecisionEndpoint) GetRecentDecisionsEnabled() bool`

GetRecentDecisionsEnabled returns the RecentDecisionsEnabled field if non-nil, zero value otherwise.

### GetRecentDecisionsEnabledOk

`func (o *DecisionEndpoint) GetRecentDecisionsEnabledOk() (*bool, bool)`

GetRecentDecisionsEnabledOk returns a tuple with the RecentDecisionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDecisionsEnabled

`func (o *DecisionEndpoint) SetRecentDecisionsEnabled(v bool)`

SetRecentDecisionsEnabled sets RecentDecisionsEnabled field to given value.

### HasRecentDecisionsEnabled

`func (o *DecisionEndpoint) HasRecentDecisionsEnabled() bool`

HasRecentDecisionsEnabled returns a boolean if a field has been set.

### GetRecentDecisions

`func (o *DecisionEndpoint) GetRecentDecisions() DecisionEndpointRecentDecisions`

GetRecentDecisions returns the RecentDecisions field if non-nil, zero value otherwise.

### GetRecentDecisionsOk

`func (o *DecisionEndpoint) GetRecentDecisionsOk() (*DecisionEndpointRecentDecisions, bool)`

GetRecentDecisionsOk returns a tuple with the RecentDecisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDecisions

`func (o *DecisionEndpoint) SetRecentDecisions(v DecisionEndpointRecentDecisions)`

SetRecentDecisions sets RecentDecisions field to given value.

### HasRecentDecisions

`func (o *DecisionEndpoint) HasRecentDecisions() bool`

HasRecentDecisions returns a boolean if a field has been set.

### GetRecordRecentRequests

`func (o *DecisionEndpoint) GetRecordRecentRequests() bool`

GetRecordRecentRequests returns the RecordRecentRequests field if non-nil, zero value otherwise.

### GetRecordRecentRequestsOk

`func (o *DecisionEndpoint) GetRecordRecentRequestsOk() (*bool, bool)`

GetRecordRecentRequestsOk returns a tuple with the RecordRecentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordRecentRequests

`func (o *DecisionEndpoint) SetRecordRecentRequests(v bool)`

SetRecordRecentRequests sets RecordRecentRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


