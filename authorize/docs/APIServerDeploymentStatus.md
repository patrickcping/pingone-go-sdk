# APIServerDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**EnumAPIServerDeploymentStatusCode**](EnumAPIServerDeploymentStatusCode.md) |  | [optional] 
**Error** | Pointer to [**APIServerDeploymentStatusError**](APIServerDeploymentStatusError.md) |  | [optional] 

## Methods

### NewAPIServerDeploymentStatus

`func NewAPIServerDeploymentStatus() *APIServerDeploymentStatus`

NewAPIServerDeploymentStatus instantiates a new APIServerDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerDeploymentStatusWithDefaults

`func NewAPIServerDeploymentStatusWithDefaults() *APIServerDeploymentStatus`

NewAPIServerDeploymentStatusWithDefaults instantiates a new APIServerDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIServerDeploymentStatus) GetCode() EnumAPIServerDeploymentStatusCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIServerDeploymentStatus) GetCodeOk() (*EnumAPIServerDeploymentStatusCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIServerDeploymentStatus) SetCode(v EnumAPIServerDeploymentStatusCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIServerDeploymentStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *APIServerDeploymentStatus) GetError() APIServerDeploymentStatusError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *APIServerDeploymentStatus) GetErrorOk() (*APIServerDeploymentStatusError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *APIServerDeploymentStatus) SetError(v APIServerDeploymentStatusError)`

SetError sets Error field to given value.

### HasError

`func (o *APIServerDeploymentStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


