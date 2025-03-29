# DtoCreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCadence** | [**TypesBillingCadence**](TypesBillingCadence.md) |  | 
**BillingPeriod** | [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | 
**BillingPeriodCount** | **int32** |  | 
**Currency** | **string** |  | 
**CustomerId** | **string** |  | 
**EndDate** | Pointer to **string** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PlanId** | **string** |  | 
**StartDate** | **string** |  | 
**TrialEnd** | Pointer to **string** |  | [optional] 
**TrialStart** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateSubscriptionRequest

`func NewDtoCreateSubscriptionRequest(billingCadence TypesBillingCadence, billingPeriod TypesBillingPeriod, billingPeriodCount int32, currency string, customerId string, planId string, startDate string, ) *DtoCreateSubscriptionRequest`

NewDtoCreateSubscriptionRequest instantiates a new DtoCreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateSubscriptionRequestWithDefaults

`func NewDtoCreateSubscriptionRequestWithDefaults() *DtoCreateSubscriptionRequest`

NewDtoCreateSubscriptionRequestWithDefaults instantiates a new DtoCreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCadence

`func (o *DtoCreateSubscriptionRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoCreateSubscriptionRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoCreateSubscriptionRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingPeriod

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoCreateSubscriptionRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoCreateSubscriptionRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.


### GetCurrency

`func (o *DtoCreateSubscriptionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateSubscriptionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateSubscriptionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *DtoCreateSubscriptionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCreateSubscriptionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCreateSubscriptionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEndDate

`func (o *DtoCreateSubscriptionRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoCreateSubscriptionRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoCreateSubscriptionRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoCreateSubscriptionRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoCreateSubscriptionRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreateSubscriptionRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreateSubscriptionRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreateSubscriptionRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateSubscriptionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateSubscriptionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateSubscriptionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreateSubscriptionRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreateSubscriptionRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreateSubscriptionRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetStartDate

`func (o *DtoCreateSubscriptionRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoCreateSubscriptionRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoCreateSubscriptionRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetTrialEnd

`func (o *DtoCreateSubscriptionRequest) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *DtoCreateSubscriptionRequest) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *DtoCreateSubscriptionRequest) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *DtoCreateSubscriptionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *DtoCreateSubscriptionRequest) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *DtoCreateSubscriptionRequest) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *DtoCreateSubscriptionRequest) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *DtoCreateSubscriptionRequest) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


