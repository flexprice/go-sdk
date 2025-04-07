# DtoGetEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]DtoEvent**](DtoEvent.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**IterFirstKey** | Pointer to **string** |  | [optional] 
**IterLastKey** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewDtoGetEventsResponse

`func NewDtoGetEventsResponse() *DtoGetEventsResponse`

NewDtoGetEventsResponse instantiates a new DtoGetEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetEventsResponseWithDefaults

`func NewDtoGetEventsResponseWithDefaults() *DtoGetEventsResponse`

NewDtoGetEventsResponseWithDefaults instantiates a new DtoGetEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *DtoGetEventsResponse) GetEvents() []DtoEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DtoGetEventsResponse) GetEventsOk() (*[]DtoEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DtoGetEventsResponse) SetEvents(v []DtoEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *DtoGetEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHasMore

`func (o *DtoGetEventsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DtoGetEventsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DtoGetEventsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *DtoGetEventsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetIterFirstKey

`func (o *DtoGetEventsResponse) GetIterFirstKey() string`

GetIterFirstKey returns the IterFirstKey field if non-nil, zero value otherwise.

### GetIterFirstKeyOk

`func (o *DtoGetEventsResponse) GetIterFirstKeyOk() (*string, bool)`

GetIterFirstKeyOk returns a tuple with the IterFirstKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterFirstKey

`func (o *DtoGetEventsResponse) SetIterFirstKey(v string)`

SetIterFirstKey sets IterFirstKey field to given value.

### HasIterFirstKey

`func (o *DtoGetEventsResponse) HasIterFirstKey() bool`

HasIterFirstKey returns a boolean if a field has been set.

### GetIterLastKey

`func (o *DtoGetEventsResponse) GetIterLastKey() string`

GetIterLastKey returns the IterLastKey field if non-nil, zero value otherwise.

### GetIterLastKeyOk

`func (o *DtoGetEventsResponse) GetIterLastKeyOk() (*string, bool)`

GetIterLastKeyOk returns a tuple with the IterLastKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterLastKey

`func (o *DtoGetEventsResponse) SetIterLastKey(v string)`

SetIterLastKey sets IterLastKey field to given value.

### HasIterLastKey

`func (o *DtoGetEventsResponse) HasIterLastKey() bool`

HasIterLastKey returns a boolean if a field has been set.

### GetOffset

`func (o *DtoGetEventsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DtoGetEventsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DtoGetEventsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DtoGetEventsResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotalCount

`func (o *DtoGetEventsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DtoGetEventsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DtoGetEventsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DtoGetEventsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


