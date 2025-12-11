# MeterMeter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**MeterAggregation**](MeterAggregation.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** | EnvironmentID is the environment identifier for the meter | [optional] 
**EventName** | Pointer to **string** | EventName is the unique identifier for the event that this meter is tracking It is a mandatory field in the events table and hence being used as the primary matching field We can have multiple meters tracking the same event but with different filters and aggregation | [optional] 
**Filters** | Pointer to [**[]MeterFilter**](MeterFilter.md) | Filters define the criteria for the meter to be applied on the events before aggregation It also defines the possible values on which later the charges will be applied | [optional] 
**Id** | Pointer to **string** | ID is the unique identifier for the meter | [optional] 
**Name** | Pointer to **string** | Name is the display name of the meter | [optional] 
**ResetUsage** | Pointer to [**TypesResetUsage**](TypesResetUsage.md) |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewMeterMeter

`func NewMeterMeter() *MeterMeter`

NewMeterMeter instantiates a new MeterMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterMeterWithDefaults

`func NewMeterMeterWithDefaults() *MeterMeter`

NewMeterMeterWithDefaults instantiates a new MeterMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MeterMeter) GetAggregation() MeterAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MeterMeter) GetAggregationOk() (*MeterAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MeterMeter) SetAggregation(v MeterAggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MeterMeter) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MeterMeter) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeterMeter) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeterMeter) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MeterMeter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MeterMeter) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MeterMeter) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MeterMeter) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MeterMeter) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *MeterMeter) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *MeterMeter) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *MeterMeter) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *MeterMeter) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEventName

`func (o *MeterMeter) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MeterMeter) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MeterMeter) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *MeterMeter) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetFilters

`func (o *MeterMeter) GetFilters() []MeterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MeterMeter) GetFiltersOk() (*[]MeterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MeterMeter) SetFilters(v []MeterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MeterMeter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *MeterMeter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeterMeter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeterMeter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeterMeter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MeterMeter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeterMeter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeterMeter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeterMeter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResetUsage

`func (o *MeterMeter) GetResetUsage() TypesResetUsage`

GetResetUsage returns the ResetUsage field if non-nil, zero value otherwise.

### GetResetUsageOk

`func (o *MeterMeter) GetResetUsageOk() (*TypesResetUsage, bool)`

GetResetUsageOk returns a tuple with the ResetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUsage

`func (o *MeterMeter) SetResetUsage(v TypesResetUsage)`

SetResetUsage sets ResetUsage field to given value.

### HasResetUsage

`func (o *MeterMeter) HasResetUsage() bool`

HasResetUsage returns a boolean if a field has been set.

### GetStatus

`func (o *MeterMeter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeterMeter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeterMeter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeterMeter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *MeterMeter) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MeterMeter) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MeterMeter) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MeterMeter) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MeterMeter) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MeterMeter) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MeterMeter) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MeterMeter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *MeterMeter) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *MeterMeter) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *MeterMeter) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *MeterMeter) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


