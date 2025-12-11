# DtoSubscriptionChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCadence** | [**TypesBillingCadence**](TypesBillingCadence.md) |  | 
**BillingCycle** | [**TypesBillingCycle**](TypesBillingCycle.md) |  | 
**BillingPeriod** | [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | 
**BillingPeriodCount** | Pointer to **int32** | billing_period_count is the billing period count for the new subscription | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata contains additional key-value pairs for storing extra information | [optional] 
**ProrationBehavior** | [**TypesProrationBehavior**](TypesProrationBehavior.md) |  | 
**TargetPlanId** | **string** | target_plan_id is the ID of the new plan to change to (required) | 

## Methods

### NewDtoSubscriptionChangeRequest

`func NewDtoSubscriptionChangeRequest(billingCadence TypesBillingCadence, billingCycle TypesBillingCycle, billingPeriod TypesBillingPeriod, prorationBehavior TypesProrationBehavior, targetPlanId string, ) *DtoSubscriptionChangeRequest`

NewDtoSubscriptionChangeRequest instantiates a new DtoSubscriptionChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionChangeRequestWithDefaults

`func NewDtoSubscriptionChangeRequestWithDefaults() *DtoSubscriptionChangeRequest`

NewDtoSubscriptionChangeRequestWithDefaults instantiates a new DtoSubscriptionChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCadence

`func (o *DtoSubscriptionChangeRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoSubscriptionChangeRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoSubscriptionChangeRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingCycle

`func (o *DtoSubscriptionChangeRequest) GetBillingCycle() TypesBillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DtoSubscriptionChangeRequest) GetBillingCycleOk() (*TypesBillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DtoSubscriptionChangeRequest) SetBillingCycle(v TypesBillingCycle)`

SetBillingCycle sets BillingCycle field to given value.


### GetBillingPeriod

`func (o *DtoSubscriptionChangeRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoSubscriptionChangeRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoSubscriptionChangeRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoSubscriptionChangeRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoSubscriptionChangeRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoSubscriptionChangeRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *DtoSubscriptionChangeRequest) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionChangeRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionChangeRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionChangeRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionChangeRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *DtoSubscriptionChangeRequest) GetProrationBehavior() TypesProrationBehavior`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DtoSubscriptionChangeRequest) GetProrationBehaviorOk() (*TypesProrationBehavior, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DtoSubscriptionChangeRequest) SetProrationBehavior(v TypesProrationBehavior)`

SetProrationBehavior sets ProrationBehavior field to given value.


### GetTargetPlanId

`func (o *DtoSubscriptionChangeRequest) GetTargetPlanId() string`

GetTargetPlanId returns the TargetPlanId field if non-nil, zero value otherwise.

### GetTargetPlanIdOk

`func (o *DtoSubscriptionChangeRequest) GetTargetPlanIdOk() (*string, bool)`

GetTargetPlanIdOk returns a tuple with the TargetPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlanId

`func (o *DtoSubscriptionChangeRequest) SetTargetPlanId(v string)`

SetTargetPlanId sets TargetPlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


