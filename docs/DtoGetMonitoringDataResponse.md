# DtoGetMonitoringDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionLag** | Pointer to **int32** |  | [optional] 
**Points** | Pointer to [**[]DtoEventCountPoint**](DtoEventCountPoint.md) |  | [optional] 
**PostProcessingLag** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewDtoGetMonitoringDataResponse

`func NewDtoGetMonitoringDataResponse() *DtoGetMonitoringDataResponse`

NewDtoGetMonitoringDataResponse instantiates a new DtoGetMonitoringDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetMonitoringDataResponseWithDefaults

`func NewDtoGetMonitoringDataResponseWithDefaults() *DtoGetMonitoringDataResponse`

NewDtoGetMonitoringDataResponseWithDefaults instantiates a new DtoGetMonitoringDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumptionLag

`func (o *DtoGetMonitoringDataResponse) GetConsumptionLag() int32`

GetConsumptionLag returns the ConsumptionLag field if non-nil, zero value otherwise.

### GetConsumptionLagOk

`func (o *DtoGetMonitoringDataResponse) GetConsumptionLagOk() (*int32, bool)`

GetConsumptionLagOk returns a tuple with the ConsumptionLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionLag

`func (o *DtoGetMonitoringDataResponse) SetConsumptionLag(v int32)`

SetConsumptionLag sets ConsumptionLag field to given value.

### HasConsumptionLag

`func (o *DtoGetMonitoringDataResponse) HasConsumptionLag() bool`

HasConsumptionLag returns a boolean if a field has been set.

### GetPoints

`func (o *DtoGetMonitoringDataResponse) GetPoints() []DtoEventCountPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DtoGetMonitoringDataResponse) GetPointsOk() (*[]DtoEventCountPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *DtoGetMonitoringDataResponse) SetPoints(v []DtoEventCountPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *DtoGetMonitoringDataResponse) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPostProcessingLag

`func (o *DtoGetMonitoringDataResponse) GetPostProcessingLag() int32`

GetPostProcessingLag returns the PostProcessingLag field if non-nil, zero value otherwise.

### GetPostProcessingLagOk

`func (o *DtoGetMonitoringDataResponse) GetPostProcessingLagOk() (*int32, bool)`

GetPostProcessingLagOk returns a tuple with the PostProcessingLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostProcessingLag

`func (o *DtoGetMonitoringDataResponse) SetPostProcessingLag(v int32)`

SetPostProcessingLag sets PostProcessingLag field to given value.

### HasPostProcessingLag

`func (o *DtoGetMonitoringDataResponse) HasPostProcessingLag() bool`

HasPostProcessingLag returns a boolean if a field has been set.

### GetTotalCount

`func (o *DtoGetMonitoringDataResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DtoGetMonitoringDataResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DtoGetMonitoringDataResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DtoGetMonitoringDataResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


