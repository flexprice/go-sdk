# DtoProrationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | Pointer to **string** | charge_amount is the charge amount for the new subscription | [optional] 
**ChargeDescription** | Pointer to **string** | charge_description describes what the charge is for | [optional] 
**CreditAmount** | Pointer to **string** | credit_amount is the credit amount from the old subscription | [optional] 
**CreditDescription** | Pointer to **string** | credit_description describes what the credit is for | [optional] 
**Currency** | Pointer to **string** | currency is the currency for all amounts | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | current_period_end is the end of the current billing period | [optional] 
**CurrentPeriodStart** | Pointer to **string** | current_period_start is the start of the current billing period | [optional] 
**DaysRemaining** | Pointer to **int32** | days_remaining is the number of days remaining in the current period | [optional] 
**DaysUsed** | Pointer to **int32** | days_used is the number of days used in the current period | [optional] 
**NetAmount** | Pointer to **string** | net_amount is the net amount (charge - credit) | [optional] 
**ProrationDate** | Pointer to **string** | proration_date is the date used for proration calculations | [optional] 

## Methods

### NewDtoProrationDetails

`func NewDtoProrationDetails() *DtoProrationDetails`

NewDtoProrationDetails instantiates a new DtoProrationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoProrationDetailsWithDefaults

`func NewDtoProrationDetailsWithDefaults() *DtoProrationDetails`

NewDtoProrationDetailsWithDefaults instantiates a new DtoProrationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeAmount

`func (o *DtoProrationDetails) GetChargeAmount() string`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *DtoProrationDetails) GetChargeAmountOk() (*string, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *DtoProrationDetails) SetChargeAmount(v string)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *DtoProrationDetails) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetChargeDescription

`func (o *DtoProrationDetails) GetChargeDescription() string`

GetChargeDescription returns the ChargeDescription field if non-nil, zero value otherwise.

### GetChargeDescriptionOk

`func (o *DtoProrationDetails) GetChargeDescriptionOk() (*string, bool)`

GetChargeDescriptionOk returns a tuple with the ChargeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDescription

`func (o *DtoProrationDetails) SetChargeDescription(v string)`

SetChargeDescription sets ChargeDescription field to given value.

### HasChargeDescription

`func (o *DtoProrationDetails) HasChargeDescription() bool`

HasChargeDescription returns a boolean if a field has been set.

### GetCreditAmount

`func (o *DtoProrationDetails) GetCreditAmount() string`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *DtoProrationDetails) GetCreditAmountOk() (*string, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *DtoProrationDetails) SetCreditAmount(v string)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *DtoProrationDetails) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCreditDescription

`func (o *DtoProrationDetails) GetCreditDescription() string`

GetCreditDescription returns the CreditDescription field if non-nil, zero value otherwise.

### GetCreditDescriptionOk

`func (o *DtoProrationDetails) GetCreditDescriptionOk() (*string, bool)`

GetCreditDescriptionOk returns a tuple with the CreditDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDescription

`func (o *DtoProrationDetails) SetCreditDescription(v string)`

SetCreditDescription sets CreditDescription field to given value.

### HasCreditDescription

`func (o *DtoProrationDetails) HasCreditDescription() bool`

HasCreditDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoProrationDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoProrationDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoProrationDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoProrationDetails) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *DtoProrationDetails) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *DtoProrationDetails) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *DtoProrationDetails) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *DtoProrationDetails) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *DtoProrationDetails) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *DtoProrationDetails) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *DtoProrationDetails) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *DtoProrationDetails) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetDaysRemaining

`func (o *DtoProrationDetails) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *DtoProrationDetails) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *DtoProrationDetails) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.

### HasDaysRemaining

`func (o *DtoProrationDetails) HasDaysRemaining() bool`

HasDaysRemaining returns a boolean if a field has been set.

### GetDaysUsed

`func (o *DtoProrationDetails) GetDaysUsed() int32`

GetDaysUsed returns the DaysUsed field if non-nil, zero value otherwise.

### GetDaysUsedOk

`func (o *DtoProrationDetails) GetDaysUsedOk() (*int32, bool)`

GetDaysUsedOk returns a tuple with the DaysUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUsed

`func (o *DtoProrationDetails) SetDaysUsed(v int32)`

SetDaysUsed sets DaysUsed field to given value.

### HasDaysUsed

`func (o *DtoProrationDetails) HasDaysUsed() bool`

HasDaysUsed returns a boolean if a field has been set.

### GetNetAmount

`func (o *DtoProrationDetails) GetNetAmount() string`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *DtoProrationDetails) GetNetAmountOk() (*string, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *DtoProrationDetails) SetNetAmount(v string)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *DtoProrationDetails) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetProrationDate

`func (o *DtoProrationDetails) GetProrationDate() string`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *DtoProrationDetails) GetProrationDateOk() (*string, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *DtoProrationDetails) SetProrationDate(v string)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *DtoProrationDetails) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


