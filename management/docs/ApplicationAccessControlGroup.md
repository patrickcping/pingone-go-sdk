# ApplicationAccessControlGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumApplicationAccessControlGroupType**](EnumApplicationAccessControlGroupType.md) |  | 
**Groups** | [**[]ApplicationAccessControlGroupGroupsInner**](ApplicationAccessControlGroupGroupsInner.md) | A set that specifies the group IDs for the groups the actor must belong to for access to the application. | 

## Methods

### NewApplicationAccessControlGroup

`func NewApplicationAccessControlGroup(type_ EnumApplicationAccessControlGroupType, groups []ApplicationAccessControlGroupGroupsInner, ) *ApplicationAccessControlGroup`

NewApplicationAccessControlGroup instantiates a new ApplicationAccessControlGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAccessControlGroupWithDefaults

`func NewApplicationAccessControlGroupWithDefaults() *ApplicationAccessControlGroup`

NewApplicationAccessControlGroupWithDefaults instantiates a new ApplicationAccessControlGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationAccessControlGroup) GetType() EnumApplicationAccessControlGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationAccessControlGroup) GetTypeOk() (*EnumApplicationAccessControlGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationAccessControlGroup) SetType(v EnumApplicationAccessControlGroupType)`

SetType sets Type field to given value.


### GetGroups

`func (o *ApplicationAccessControlGroup) GetGroups() []ApplicationAccessControlGroupGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ApplicationAccessControlGroup) GetGroupsOk() (*[]ApplicationAccessControlGroupGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ApplicationAccessControlGroup) SetGroups(v []ApplicationAccessControlGroupGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


