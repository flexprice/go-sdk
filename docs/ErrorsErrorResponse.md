# ErrorsErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorsErrorDetail**](ErrorsErrorDetail.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewErrorsErrorResponse

`func NewErrorsErrorResponse() *ErrorsErrorResponse`

NewErrorsErrorResponse instantiates a new ErrorsErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsErrorResponseWithDefaults

`func NewErrorsErrorResponseWithDefaults() *ErrorsErrorResponse`

NewErrorsErrorResponseWithDefaults instantiates a new ErrorsErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorsErrorResponse) GetError() ErrorsErrorDetail`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorsErrorResponse) GetErrorOk() (*ErrorsErrorDetail, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorsErrorResponse) SetError(v ErrorsErrorDetail)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorsErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *ErrorsErrorResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ErrorsErrorResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ErrorsErrorResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ErrorsErrorResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


