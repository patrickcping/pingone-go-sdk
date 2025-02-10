# AuthorizeEditorDataEntityTestResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedMicroseconds** | Pointer to **int32** |  | [optional] 
**Request** | [**AuthorizeEditorDataEntityTestRequestDTO**](AuthorizeEditorDataEntityTestRequestDTO.md) |  | 
**Authorized** | Pointer to **bool** |  | [optional] 
**Status** | [**AuthorizeEditorDataEntityTestingDecisionStatusDTO**](AuthorizeEditorDataEntityTestingDecisionStatusDTO.md) |  | 
**Result** | Pointer to [**AuthorizeEditorDataEntityTestingDecisionResultDTO**](AuthorizeEditorDataEntityTestingDecisionResultDTO.md) |  | [optional] 

## Methods

### NewAuthorizeEditorDataEntityTestResponseDTO

`func NewAuthorizeEditorDataEntityTestResponseDTO(request AuthorizeEditorDataEntityTestRequestDTO, status AuthorizeEditorDataEntityTestingDecisionStatusDTO, ) *AuthorizeEditorDataEntityTestResponseDTO`

NewAuthorizeEditorDataEntityTestResponseDTO instantiates a new AuthorizeEditorDataEntityTestResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataEntityTestResponseDTOWithDefaults

`func NewAuthorizeEditorDataEntityTestResponseDTOWithDefaults() *AuthorizeEditorDataEntityTestResponseDTO`

NewAuthorizeEditorDataEntityTestResponseDTOWithDefaults instantiates a new AuthorizeEditorDataEntityTestResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedMicroseconds

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetElapsedMicroseconds() int32`

GetElapsedMicroseconds returns the ElapsedMicroseconds field if non-nil, zero value otherwise.

### GetElapsedMicrosecondsOk

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetElapsedMicrosecondsOk() (*int32, bool)`

GetElapsedMicrosecondsOk returns a tuple with the ElapsedMicroseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedMicroseconds

`func (o *AuthorizeEditorDataEntityTestResponseDTO) SetElapsedMicroseconds(v int32)`

SetElapsedMicroseconds sets ElapsedMicroseconds field to given value.

### HasElapsedMicroseconds

`func (o *AuthorizeEditorDataEntityTestResponseDTO) HasElapsedMicroseconds() bool`

HasElapsedMicroseconds returns a boolean if a field has been set.

### GetRequest

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetRequest() AuthorizeEditorDataEntityTestRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetRequestOk() (*AuthorizeEditorDataEntityTestRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AuthorizeEditorDataEntityTestResponseDTO) SetRequest(v AuthorizeEditorDataEntityTestRequestDTO)`

SetRequest sets Request field to given value.


### GetAuthorized

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *AuthorizeEditorDataEntityTestResponseDTO) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *AuthorizeEditorDataEntityTestResponseDTO) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetStatus() AuthorizeEditorDataEntityTestingDecisionStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetStatusOk() (*AuthorizeEditorDataEntityTestingDecisionStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizeEditorDataEntityTestResponseDTO) SetStatus(v AuthorizeEditorDataEntityTestingDecisionStatusDTO)`

SetStatus sets Status field to given value.


### GetResult

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetResult() AuthorizeEditorDataEntityTestingDecisionResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthorizeEditorDataEntityTestResponseDTO) GetResultOk() (*AuthorizeEditorDataEntityTestingDecisionResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthorizeEditorDataEntityTestResponseDTO) SetResult(v AuthorizeEditorDataEntityTestingDecisionResultDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuthorizeEditorDataEntityTestResponseDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


