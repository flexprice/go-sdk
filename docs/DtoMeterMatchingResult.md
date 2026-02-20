# DtoMeterMatchingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorsErrorResponse**](ErrorsErrorResponse.md) |  | [optional] 
**MatchedMeters** | Pointer to [**[]DtoMatchedMeter**](DtoMatchedMeter.md) |  | [optional] 
**Status** | Pointer to [**TypesDebugTrackerStatus**](TypesDebugTrackerStatus.md) |  | [optional] 

## Methods

### NewDtoMeterMatchingResult

`func NewDtoMeterMatchingResult() *DtoMeterMatchingResult`

NewDtoMeterMatchingResult instantiates a new DtoMeterMatchingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoMeterMatchingResultWithDefaults

`func NewDtoMeterMatchingResultWithDefaults() *DtoMeterMatchingResult`

NewDtoMeterMatchingResultWithDefaults instantiates a new DtoMeterMatchingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DtoMeterMatchingResult) GetError() ErrorsErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DtoMeterMatchingResult) GetErrorOk() (*ErrorsErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DtoMeterMatchingResult) SetError(v ErrorsErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *DtoMeterMatchingResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMatchedMeters

`func (o *DtoMeterMatchingResult) GetMatchedMeters() []DtoMatchedMeter`

GetMatchedMeters returns the MatchedMeters field if non-nil, zero value otherwise.

### GetMatchedMetersOk

`func (o *DtoMeterMatchingResult) GetMatchedMetersOk() (*[]DtoMatchedMeter, bool)`

GetMatchedMetersOk returns a tuple with the MatchedMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedMeters

`func (o *DtoMeterMatchingResult) SetMatchedMeters(v []DtoMatchedMeter)`

SetMatchedMeters sets MatchedMeters field to given value.

### HasMatchedMeters

`func (o *DtoMeterMatchingResult) HasMatchedMeters() bool`

HasMatchedMeters returns a boolean if a field has been set.

### GetStatus

`func (o *DtoMeterMatchingResult) GetStatus() TypesDebugTrackerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoMeterMatchingResult) GetStatusOk() (*TypesDebugTrackerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoMeterMatchingResult) SetStatus(v TypesDebugTrackerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoMeterMatchingResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


