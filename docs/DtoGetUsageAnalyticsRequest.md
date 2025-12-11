# DtoGetUsageAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **[]string** | allowed values: \&quot;price\&quot;, \&quot;meter\&quot;, \&quot;feature\&quot;, \&quot;subscription_line_item\&quot;,\&quot;plan\&quot;,\&quot;addon\&quot; | [optional] 
**ExternalCustomerId** | **string** |  | 
**FeatureIds** | Pointer to **[]string** |  | [optional] 
**GroupBy** | Pointer to **[]string** | allowed values: \&quot;source\&quot;, \&quot;feature_id\&quot;, \&quot;properties.&lt;field_name&gt;\&quot; | [optional] 
**PropertyFilters** | Pointer to **map[string][]string** | Property filters to filter the events by the keys in &#x60;properties&#x60; field of the event | [optional] 
**Sources** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**WindowSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 

## Methods

### NewDtoGetUsageAnalyticsRequest

`func NewDtoGetUsageAnalyticsRequest(externalCustomerId string, ) *DtoGetUsageAnalyticsRequest`

NewDtoGetUsageAnalyticsRequest instantiates a new DtoGetUsageAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageAnalyticsRequestWithDefaults

`func NewDtoGetUsageAnalyticsRequestWithDefaults() *DtoGetUsageAnalyticsRequest`

NewDtoGetUsageAnalyticsRequestWithDefaults instantiates a new DtoGetUsageAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DtoGetUsageAnalyticsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetUsageAnalyticsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetUsageAnalyticsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetUsageAnalyticsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *DtoGetUsageAnalyticsRequest) GetExpand() []string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *DtoGetUsageAnalyticsRequest) GetExpandOk() (*[]string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *DtoGetUsageAnalyticsRequest) SetExpand(v []string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *DtoGetUsageAnalyticsRequest) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *DtoGetUsageAnalyticsRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoGetUsageAnalyticsRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoGetUsageAnalyticsRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetFeatureIds

`func (o *DtoGetUsageAnalyticsRequest) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *DtoGetUsageAnalyticsRequest) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *DtoGetUsageAnalyticsRequest) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *DtoGetUsageAnalyticsRequest) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetGroupBy

`func (o *DtoGetUsageAnalyticsRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *DtoGetUsageAnalyticsRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *DtoGetUsageAnalyticsRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *DtoGetUsageAnalyticsRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetPropertyFilters

`func (o *DtoGetUsageAnalyticsRequest) GetPropertyFilters() map[string][]string`

GetPropertyFilters returns the PropertyFilters field if non-nil, zero value otherwise.

### GetPropertyFiltersOk

`func (o *DtoGetUsageAnalyticsRequest) GetPropertyFiltersOk() (*map[string][]string, bool)`

GetPropertyFiltersOk returns a tuple with the PropertyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyFilters

`func (o *DtoGetUsageAnalyticsRequest) SetPropertyFilters(v map[string][]string)`

SetPropertyFilters sets PropertyFilters field to given value.

### HasPropertyFilters

`func (o *DtoGetUsageAnalyticsRequest) HasPropertyFilters() bool`

HasPropertyFilters returns a boolean if a field has been set.

### GetSources

`func (o *DtoGetUsageAnalyticsRequest) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtoGetUsageAnalyticsRequest) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtoGetUsageAnalyticsRequest) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtoGetUsageAnalyticsRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetUsageAnalyticsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetUsageAnalyticsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetUsageAnalyticsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetUsageAnalyticsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWindowSize

`func (o *DtoGetUsageAnalyticsRequest) GetWindowSize() TypesWindowSize`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *DtoGetUsageAnalyticsRequest) GetWindowSizeOk() (*TypesWindowSize, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *DtoGetUsageAnalyticsRequest) SetWindowSize(v TypesWindowSize)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *DtoGetUsageAnalyticsRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


