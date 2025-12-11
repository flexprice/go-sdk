# DtoSourceUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | cost is the cost attributed to this source for the line item | [optional] 
**EventCount** | Pointer to **int32** | event_count is the number of events from this source (optional) | [optional] 
**Percentage** | Pointer to **string** | percentage is the percentage of total line item cost from this source (optional) | [optional] 
**Source** | Pointer to **string** | source is the name of the event source | [optional] 
**Usage** | Pointer to **string** | usage is the total usage amount from this source (optional, for additional context) | [optional] 

## Methods

### NewDtoSourceUsageItem

`func NewDtoSourceUsageItem() *DtoSourceUsageItem`

NewDtoSourceUsageItem instantiates a new DtoSourceUsageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSourceUsageItemWithDefaults

`func NewDtoSourceUsageItemWithDefaults() *DtoSourceUsageItem`

NewDtoSourceUsageItemWithDefaults instantiates a new DtoSourceUsageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *DtoSourceUsageItem) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DtoSourceUsageItem) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DtoSourceUsageItem) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DtoSourceUsageItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetEventCount

`func (o *DtoSourceUsageItem) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *DtoSourceUsageItem) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *DtoSourceUsageItem) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *DtoSourceUsageItem) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetPercentage

`func (o *DtoSourceUsageItem) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *DtoSourceUsageItem) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *DtoSourceUsageItem) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *DtoSourceUsageItem) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetSource

`func (o *DtoSourceUsageItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DtoSourceUsageItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DtoSourceUsageItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DtoSourceUsageItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUsage

`func (o *DtoSourceUsageItem) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DtoSourceUsageItem) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DtoSourceUsageItem) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DtoSourceUsageItem) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


