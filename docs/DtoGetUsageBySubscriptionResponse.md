# DtoGetUsageBySubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Charges** | Pointer to [**[]DtoSubscriptionUsageByMetersResponse**](DtoSubscriptionUsageByMetersResponse.md) |  | [optional] 
**CommitmentAmount** | Pointer to **float32** |  | [optional] 
**CommitmentUtilized** | Pointer to **float32** | Amount of commitment used | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayAmount** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**HasOverage** | Pointer to **bool** | Whether any usage exceeded commitment | [optional] 
**OverageAmount** | Pointer to **float32** | Amount charged at overage rate | [optional] 
**OverageFactor** | Pointer to **float32** |  | [optional] 
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

### GetCommitmentAmount

`func (o *DtoGetUsageBySubscriptionResponse) GetCommitmentAmount() float32`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoGetUsageBySubscriptionResponse) GetCommitmentAmountOk() (*float32, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoGetUsageBySubscriptionResponse) SetCommitmentAmount(v float32)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoGetUsageBySubscriptionResponse) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentUtilized

`func (o *DtoGetUsageBySubscriptionResponse) GetCommitmentUtilized() float32`

GetCommitmentUtilized returns the CommitmentUtilized field if non-nil, zero value otherwise.

### GetCommitmentUtilizedOk

`func (o *DtoGetUsageBySubscriptionResponse) GetCommitmentUtilizedOk() (*float32, bool)`

GetCommitmentUtilizedOk returns a tuple with the CommitmentUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentUtilized

`func (o *DtoGetUsageBySubscriptionResponse) SetCommitmentUtilized(v float32)`

SetCommitmentUtilized sets CommitmentUtilized field to given value.

### HasCommitmentUtilized

`func (o *DtoGetUsageBySubscriptionResponse) HasCommitmentUtilized() bool`

HasCommitmentUtilized returns a boolean if a field has been set.

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

### GetHasOverage

`func (o *DtoGetUsageBySubscriptionResponse) GetHasOverage() bool`

GetHasOverage returns the HasOverage field if non-nil, zero value otherwise.

### GetHasOverageOk

`func (o *DtoGetUsageBySubscriptionResponse) GetHasOverageOk() (*bool, bool)`

GetHasOverageOk returns a tuple with the HasOverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverage

`func (o *DtoGetUsageBySubscriptionResponse) SetHasOverage(v bool)`

SetHasOverage sets HasOverage field to given value.

### HasHasOverage

`func (o *DtoGetUsageBySubscriptionResponse) HasHasOverage() bool`

HasHasOverage returns a boolean if a field has been set.

### GetOverageAmount

`func (o *DtoGetUsageBySubscriptionResponse) GetOverageAmount() float32`

GetOverageAmount returns the OverageAmount field if non-nil, zero value otherwise.

### GetOverageAmountOk

`func (o *DtoGetUsageBySubscriptionResponse) GetOverageAmountOk() (*float32, bool)`

GetOverageAmountOk returns a tuple with the OverageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageAmount

`func (o *DtoGetUsageBySubscriptionResponse) SetOverageAmount(v float32)`

SetOverageAmount sets OverageAmount field to given value.

### HasOverageAmount

`func (o *DtoGetUsageBySubscriptionResponse) HasOverageAmount() bool`

HasOverageAmount returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoGetUsageBySubscriptionResponse) GetOverageFactor() float32`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoGetUsageBySubscriptionResponse) GetOverageFactorOk() (*float32, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoGetUsageBySubscriptionResponse) SetOverageFactor(v float32)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoGetUsageBySubscriptionResponse) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

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


