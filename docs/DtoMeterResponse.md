# DtoMeterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**MeterAggregation**](MeterAggregation.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]MeterFilter**](MeterFilter.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResetUsage** | Pointer to [**TypesResetUsage**](TypesResetUsage.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoMeterResponse

`func NewDtoMeterResponse() *DtoMeterResponse`

NewDtoMeterResponse instantiates a new DtoMeterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoMeterResponseWithDefaults

`func NewDtoMeterResponseWithDefaults() *DtoMeterResponse`

NewDtoMeterResponseWithDefaults instantiates a new DtoMeterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *DtoMeterResponse) GetAggregation() MeterAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *DtoMeterResponse) GetAggregationOk() (*MeterAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *DtoMeterResponse) SetAggregation(v MeterAggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *DtoMeterResponse) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoMeterResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoMeterResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoMeterResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoMeterResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEventName

`func (o *DtoMeterResponse) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoMeterResponse) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoMeterResponse) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DtoMeterResponse) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetFilters

`func (o *DtoMeterResponse) GetFilters() []MeterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DtoMeterResponse) GetFiltersOk() (*[]MeterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DtoMeterResponse) SetFilters(v []MeterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DtoMeterResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *DtoMeterResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoMeterResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoMeterResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoMeterResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DtoMeterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoMeterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoMeterResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoMeterResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResetUsage

`func (o *DtoMeterResponse) GetResetUsage() TypesResetUsage`

GetResetUsage returns the ResetUsage field if non-nil, zero value otherwise.

### GetResetUsageOk

`func (o *DtoMeterResponse) GetResetUsageOk() (*TypesResetUsage, bool)`

GetResetUsageOk returns a tuple with the ResetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUsage

`func (o *DtoMeterResponse) SetResetUsage(v TypesResetUsage)`

SetResetUsage sets ResetUsage field to given value.

### HasResetUsage

`func (o *DtoMeterResponse) HasResetUsage() bool`

HasResetUsage returns a boolean if a field has been set.

### GetStatus

`func (o *DtoMeterResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoMeterResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoMeterResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoMeterResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoMeterResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoMeterResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoMeterResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoMeterResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoMeterResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoMeterResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoMeterResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoMeterResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


