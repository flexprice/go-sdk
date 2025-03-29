# DtoIngestEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 
**ExternalCustomerId** | **string** |  | 
**Properties** | Pointer to **map[string]string** | Handled separately for dynamic columns | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** | Handled separately due to parsing | [optional] 

## Methods

### NewDtoIngestEventRequest

`func NewDtoIngestEventRequest(eventName string, externalCustomerId string, ) *DtoIngestEventRequest`

NewDtoIngestEventRequest instantiates a new DtoIngestEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoIngestEventRequestWithDefaults

`func NewDtoIngestEventRequestWithDefaults() *DtoIngestEventRequest`

NewDtoIngestEventRequestWithDefaults instantiates a new DtoIngestEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DtoIngestEventRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoIngestEventRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoIngestEventRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoIngestEventRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEventId

`func (o *DtoIngestEventRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *DtoIngestEventRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *DtoIngestEventRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *DtoIngestEventRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *DtoIngestEventRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoIngestEventRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoIngestEventRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetExternalCustomerId

`func (o *DtoIngestEventRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoIngestEventRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoIngestEventRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.


### GetProperties

`func (o *DtoIngestEventRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DtoIngestEventRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DtoIngestEventRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DtoIngestEventRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSource

`func (o *DtoIngestEventRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DtoIngestEventRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DtoIngestEventRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DtoIngestEventRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimestamp

`func (o *DtoIngestEventRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DtoIngestEventRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DtoIngestEventRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DtoIngestEventRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


