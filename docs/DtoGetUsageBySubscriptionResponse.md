# DtoGetUsageBySubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Charges** | Pointer to [**[]DtoSubscriptionUsageByMetersResponse**](DtoSubscriptionUsageByMetersResponse.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayAmount** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoGetUsageBySubscriptionResponse

`func NewDtoGetUsageBySubscriptionResponse() *DtoGetUsageBySubscriptionResponse`

NewDtoGetUsageBySubscriptionResponse instantiates a new DtoGetUsageBySubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageBySubscriptionResponseWithDefaults

`func NewDtoGetUsageBySubscriptionResponseWithDefaults() *DtoGetUsageBySubscriptionResponse`

NewDtoGetUsageBySubscriptionResponseWithDefaults instantiates a new DtoGetUsageBySubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoGetUsageBySubscriptionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoGetUsageBySubscriptionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoGetUsageBySubscriptionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoGetUsageBySubscriptionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCharges

`func (o *DtoGetUsageBySubscriptionResponse) GetCharges() []DtoSubscriptionUsageByMetersResponse`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *DtoGetUsageBySubscriptionResponse) GetChargesOk() (*[]DtoSubscriptionUsageByMetersResponse, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *DtoGetUsageBySubscriptionResponse) SetCharges(v []DtoSubscriptionUsageByMetersResponse)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *DtoGetUsageBySubscriptionResponse) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoGetUsageBySubscriptionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoGetUsageBySubscriptionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoGetUsageBySubscriptionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoGetUsageBySubscriptionResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayAmount

`func (o *DtoGetUsageBySubscriptionResponse) GetDisplayAmount() string`

GetDisplayAmount returns the DisplayAmount field if non-nil, zero value otherwise.

### GetDisplayAmountOk

`func (o *DtoGetUsageBySubscriptionResponse) GetDisplayAmountOk() (*string, bool)`

GetDisplayAmountOk returns a tuple with the DisplayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAmount

`func (o *DtoGetUsageBySubscriptionResponse) SetDisplayAmount(v string)`

SetDisplayAmount sets DisplayAmount field to given value.

### HasDisplayAmount

`func (o *DtoGetUsageBySubscriptionResponse) HasDisplayAmount() bool`

HasDisplayAmount returns a boolean if a field has been set.

### GetEndTime

`func (o *DtoGetUsageBySubscriptionResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetUsageBySubscriptionResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetUsageBySubscriptionResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetUsageBySubscriptionResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetUsageBySubscriptionResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetUsageBySubscriptionResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetUsageBySubscriptionResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetUsageBySubscriptionResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


