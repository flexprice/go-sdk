# ErrorsErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**InternalError** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorsErrorDetail

`func NewErrorsErrorDetail() *ErrorsErrorDetail`

NewErrorsErrorDetail instantiates a new ErrorsErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsErrorDetailWithDefaults

`func NewErrorsErrorDetailWithDefaults() *ErrorsErrorDetail`

NewErrorsErrorDetailWithDefaults instantiates a new ErrorsErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ErrorsErrorDetail) GetDetails() map[string]map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorsErrorDetail) GetDetailsOk() (*map[string]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorsErrorDetail) SetDetails(v map[string]map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorsErrorDetail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInternalError

`func (o *ErrorsErrorDetail) GetInternalError() string`

GetInternalError returns the InternalError field if non-nil, zero value otherwise.

### GetInternalErrorOk

`func (o *ErrorsErrorDetail) GetInternalErrorOk() (*string, bool)`

GetInternalErrorOk returns a tuple with the InternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalError

`func (o *ErrorsErrorDetail) SetInternalError(v string)`

SetInternalError sets InternalError field to given value.

### HasInternalError

`func (o *ErrorsErrorDetail) HasInternalError() bool`

HasInternalError returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorsErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorsErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorsErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorsErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


