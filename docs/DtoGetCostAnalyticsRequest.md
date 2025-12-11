# DtoGetCostAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **[]string** | Expand options - specify which entities to expand | [optional] 
**ExternalCustomerId** | Pointer to **string** | Optional - for specific customer | [optional] 
**FeatureIds** | Pointer to **[]string** | Additional filters | [optional] 
**Limit** | Pointer to **int32** | Pagination | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **string** | Time range fields (optional - defaults to last 7 days if not provided) | [optional] 

## Methods

### NewDtoGetCostAnalyticsRequest

`func NewDtoGetCostAnalyticsRequest() *DtoGetCostAnalyticsRequest`

NewDtoGetCostAnalyticsRequest instantiates a new DtoGetCostAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetCostAnalyticsRequestWithDefaults

`func NewDtoGetCostAnalyticsRequestWithDefaults() *DtoGetCostAnalyticsRequest`

NewDtoGetCostAnalyticsRequestWithDefaults instantiates a new DtoGetCostAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DtoGetCostAnalyticsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetCostAnalyticsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetCostAnalyticsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetCostAnalyticsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *DtoGetCostAnalyticsRequest) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *DtoGetCostAnalyticsRequest) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *DtoGetCostAnalyticsRequest) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *DtoGetCostAnalyticsRequest) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *DtoGetCostAnalyticsRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoGetCostAnalyticsRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoGetCostAnalyticsRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoGetCostAnalyticsRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetFeatureIds

`func (o *DtoGetCostAnalyticsRequest) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *DtoGetCostAnalyticsRequest) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *DtoGetCostAnalyticsRequest) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *DtoGetCostAnalyticsRequest) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetLimit

`func (o *DtoGetCostAnalyticsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DtoGetCostAnalyticsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DtoGetCostAnalyticsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DtoGetCostAnalyticsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DtoGetCostAnalyticsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DtoGetCostAnalyticsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DtoGetCostAnalyticsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DtoGetCostAnalyticsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetCostAnalyticsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetCostAnalyticsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetCostAnalyticsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetCostAnalyticsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


