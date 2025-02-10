# TypedRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the resource referenced by this relationship | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the resource referenced by this relationship | [optional] [readonly] 

## Methods

### NewTypedRelationship

`func NewTypedRelationship() *TypedRelationship`

NewTypedRelationship instantiates a new TypedRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedRelationshipWithDefaults

`func NewTypedRelationshipWithDefaults() *TypedRelationship`

NewTypedRelationshipWithDefaults instantiates a new TypedRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypedRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypedRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypedRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TypedRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TypedRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypedRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypedRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypedRelationship) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


