# DtoUpdateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAt** | Pointer to **string** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**ParentSubscriptionId** | Pointer to **string** | ParentSubscriptionID sets or clears the parent subscription. Omit to leave unchanged; send \&quot;\&quot; to clear. | [optional] 
**Status** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 

## Methods

### NewDtoUpdateSubscriptionRequest

`func NewDtoUpdateSubscriptionRequest() *DtoUpdateSubscriptionRequest`

NewDtoUpdateSubscriptionRequest instantiates a new DtoUpdateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateSubscriptionRequestWithDefaults

`func NewDtoUpdateSubscriptionRequestWithDefaults() *DtoUpdateSubscriptionRequest`

NewDtoUpdateSubscriptionRequestWithDefaults instantiates a new DtoUpdateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAt

`func (o *DtoUpdateSubscriptionRequest) GetCancelAt() string`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *DtoUpdateSubscriptionRequest) GetCancelAtOk() (*string, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *DtoUpdateSubscriptionRequest) SetCancelAt(v string)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *DtoUpdateSubscriptionRequest) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *DtoUpdateSubscriptionRequest) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *DtoUpdateSubscriptionRequest) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *DtoUpdateSubscriptionRequest) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *DtoUpdateSubscriptionRequest) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetParentSubscriptionId

`func (o *DtoUpdateSubscriptionRequest) GetParentSubscriptionId() string`

GetParentSubscriptionId returns the ParentSubscriptionId field if non-nil, zero value otherwise.

### GetParentSubscriptionIdOk

`func (o *DtoUpdateSubscriptionRequest) GetParentSubscriptionIdOk() (*string, bool)`

GetParentSubscriptionIdOk returns a tuple with the ParentSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubscriptionId

`func (o *DtoUpdateSubscriptionRequest) SetParentSubscriptionId(v string)`

SetParentSubscriptionId sets ParentSubscriptionId field to given value.

### HasParentSubscriptionId

`func (o *DtoUpdateSubscriptionRequest) HasParentSubscriptionId() bool`

HasParentSubscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *DtoUpdateSubscriptionRequest) GetStatus() TypesSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoUpdateSubscriptionRequest) GetStatusOk() (*TypesSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoUpdateSubscriptionRequest) SetStatus(v TypesSubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoUpdateSubscriptionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


