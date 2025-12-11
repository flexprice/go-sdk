# DtoUsageBreakdownItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | cost is the cost attributed to this group for the line item | [optional] 
**EventCount** | Pointer to **int32** | event_count is the number of events from this group (optional) | [optional] 
**GroupedBy** | Pointer to **map[string]string** | grouped_by contains the grouping field values (e.g., {\&quot;source\&quot;: \&quot;api\&quot;, \&quot;org_id\&quot;: \&quot;org123\&quot;}) | [optional] 
**Percentage** | Pointer to **string** | percentage is the percentage of total line item cost from this group (optional) | [optional] 
**Usage** | Pointer to **string** | usage is the total usage amount from this group (optional, for additional context) | [optional] 

## Methods

### NewDtoUsageBreakdownItem

`func NewDtoUsageBreakdownItem() *DtoUsageBreakdownItem`

NewDtoUsageBreakdownItem instantiates a new DtoUsageBreakdownItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUsageBreakdownItemWithDefaults

`func NewDtoUsageBreakdownItemWithDefaults() *DtoUsageBreakdownItem`

NewDtoUsageBreakdownItemWithDefaults instantiates a new DtoUsageBreakdownItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *DtoUsageBreakdownItem) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DtoUsageBreakdownItem) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DtoUsageBreakdownItem) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DtoUsageBreakdownItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetEventCount

`func (o *DtoUsageBreakdownItem) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *DtoUsageBreakdownItem) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *DtoUsageBreakdownItem) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *DtoUsageBreakdownItem) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetGroupedBy

`func (o *DtoUsageBreakdownItem) GetGroupedBy() map[string]string`

GetGroupedBy returns the GroupedBy field if non-nil, zero value otherwise.

### GetGroupedByOk

`func (o *DtoUsageBreakdownItem) GetGroupedByOk() (*map[string]string, bool)`

GetGroupedByOk returns a tuple with the GroupedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedBy

`func (o *DtoUsageBreakdownItem) SetGroupedBy(v map[string]string)`

SetGroupedBy sets GroupedBy field to given value.

### HasGroupedBy

`func (o *DtoUsageBreakdownItem) HasGroupedBy() bool`

HasGroupedBy returns a boolean if a field has been set.

### GetPercentage

`func (o *DtoUsageBreakdownItem) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *DtoUsageBreakdownItem) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *DtoUsageBreakdownItem) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *DtoUsageBreakdownItem) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetUsage

`func (o *DtoUsageBreakdownItem) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DtoUsageBreakdownItem) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DtoUsageBreakdownItem) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DtoUsageBreakdownItem) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


