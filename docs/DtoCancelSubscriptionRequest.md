# DtoCancelSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationType** | [**TypesCancellationType**](TypesCancellationType.md) |  | 
**ProrationBehavior** | Pointer to [**TypesProrationBehavior**](TypesProrationBehavior.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason for cancellation (for audit and business intelligence) | [optional] 

## Methods

### NewDtoCancelSubscriptionRequest

`func NewDtoCancelSubscriptionRequest(cancellationType TypesCancellationType, ) *DtoCancelSubscriptionRequest`

NewDtoCancelSubscriptionRequest instantiates a new DtoCancelSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCancelSubscriptionRequestWithDefaults

`func NewDtoCancelSubscriptionRequestWithDefaults() *DtoCancelSubscriptionRequest`

NewDtoCancelSubscriptionRequestWithDefaults instantiates a new DtoCancelSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationType

`func (o *DtoCancelSubscriptionRequest) GetCancellationType() TypesCancellationType`

GetCancellationType returns the CancellationType field if non-nil, zero value otherwise.

### GetCancellationTypeOk

`func (o *DtoCancelSubscriptionRequest) GetCancellationTypeOk() (*TypesCancellationType, bool)`

GetCancellationTypeOk returns a tuple with the CancellationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationType

`func (o *DtoCancelSubscriptionRequest) SetCancellationType(v TypesCancellationType)`

SetCancellationType sets CancellationType field to given value.


### GetProrationBehavior

`func (o *DtoCancelSubscriptionRequest) GetProrationBehavior() TypesProrationBehavior`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DtoCancelSubscriptionRequest) GetProrationBehaviorOk() (*TypesProrationBehavior, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DtoCancelSubscriptionRequest) SetProrationBehavior(v TypesProrationBehavior)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *DtoCancelSubscriptionRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetReason

`func (o *DtoCancelSubscriptionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoCancelSubscriptionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoCancelSubscriptionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoCancelSubscriptionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


