# DtoCancelScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | message is a confirmation message | [optional] 
**Status** | Pointer to [**TypesScheduleStatus**](TypesScheduleStatus.md) |  | [optional] 

## Methods

### NewDtoCancelScheduleResponse

`func NewDtoCancelScheduleResponse() *DtoCancelScheduleResponse`

NewDtoCancelScheduleResponse instantiates a new DtoCancelScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCancelScheduleResponseWithDefaults

`func NewDtoCancelScheduleResponseWithDefaults() *DtoCancelScheduleResponse`

NewDtoCancelScheduleResponseWithDefaults instantiates a new DtoCancelScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DtoCancelScheduleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DtoCancelScheduleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DtoCancelScheduleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DtoCancelScheduleResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCancelScheduleResponse) GetStatus() TypesScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCancelScheduleResponse) GetStatusOk() (*TypesScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCancelScheduleResponse) SetStatus(v TypesScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCancelScheduleResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


