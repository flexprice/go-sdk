# DtoCancelScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** | schedule_id is the ID of the schedule to cancel (optional if subscription_id and schedule_type are provided) | [optional] 
**ScheduleType** | Pointer to [**TypesSubscriptionScheduleChangeType**](TypesSubscriptionScheduleChangeType.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the ID of the subscription (required if schedule_id is not provided) | [optional] 

## Methods

### NewDtoCancelScheduleRequest

`func NewDtoCancelScheduleRequest() *DtoCancelScheduleRequest`

NewDtoCancelScheduleRequest instantiates a new DtoCancelScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCancelScheduleRequestWithDefaults

`func NewDtoCancelScheduleRequestWithDefaults() *DtoCancelScheduleRequest`

NewDtoCancelScheduleRequestWithDefaults instantiates a new DtoCancelScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleId

`func (o *DtoCancelScheduleRequest) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DtoCancelScheduleRequest) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DtoCancelScheduleRequest) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DtoCancelScheduleRequest) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetScheduleType

`func (o *DtoCancelScheduleRequest) GetScheduleType() TypesSubscriptionScheduleChangeType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *DtoCancelScheduleRequest) GetScheduleTypeOk() (*TypesSubscriptionScheduleChangeType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *DtoCancelScheduleRequest) SetScheduleType(v TypesSubscriptionScheduleChangeType)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *DtoCancelScheduleRequest) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCancelScheduleRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCancelScheduleRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCancelScheduleRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCancelScheduleRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


