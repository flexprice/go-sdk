# DtoGetUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]DtoUsageResult**](DtoUsageResult.md) |  | [optional] 
**Type** | Pointer to [**TypesAggregationType**](TypesAggregationType.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewDtoGetUsageResponse

`func NewDtoGetUsageResponse() *DtoGetUsageResponse`

NewDtoGetUsageResponse instantiates a new DtoGetUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageResponseWithDefaults

`func NewDtoGetUsageResponseWithDefaults() *DtoGetUsageResponse`

NewDtoGetUsageResponseWithDefaults instantiates a new DtoGetUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *DtoGetUsageResponse) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoGetUsageResponse) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoGetUsageResponse) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DtoGetUsageResponse) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetResults

`func (o *DtoGetUsageResponse) GetResults() []DtoUsageResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DtoGetUsageResponse) GetResultsOk() (*[]DtoUsageResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DtoGetUsageResponse) SetResults(v []DtoUsageResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *DtoGetUsageResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetType

`func (o *DtoGetUsageResponse) GetType() TypesAggregationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoGetUsageResponse) GetTypeOk() (*TypesAggregationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoGetUsageResponse) SetType(v TypesAggregationType)`

SetType sets Type field to given value.

### HasType

`func (o *DtoGetUsageResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DtoGetUsageResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DtoGetUsageResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DtoGetUsageResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *DtoGetUsageResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


