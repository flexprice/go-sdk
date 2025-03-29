# DtoCustomerInvoiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**OverdueInvoiceCount** | Pointer to **int32** |  | [optional] 
**TotalInvoiceCount** | Pointer to **int32** |  | [optional] 
**TotalOverdueAmount** | Pointer to **float32** |  | [optional] 
**TotalRevenueAmount** | Pointer to **float32** |  | [optional] 
**TotalUnpaidAmount** | Pointer to **float32** |  | [optional] 
**UnpaidFixedCharges** | Pointer to **float32** |  | [optional] 
**UnpaidInvoiceCount** | Pointer to **int32** |  | [optional] 
**UnpaidUsageCharges** | Pointer to **float32** |  | [optional] 

## Methods

### NewDtoCustomerInvoiceSummary

`func NewDtoCustomerInvoiceSummary() *DtoCustomerInvoiceSummary`

NewDtoCustomerInvoiceSummary instantiates a new DtoCustomerInvoiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCustomerInvoiceSummaryWithDefaults

`func NewDtoCustomerInvoiceSummaryWithDefaults() *DtoCustomerInvoiceSummary`

NewDtoCustomerInvoiceSummaryWithDefaults instantiates a new DtoCustomerInvoiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *DtoCustomerInvoiceSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCustomerInvoiceSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCustomerInvoiceSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCustomerInvoiceSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoCustomerInvoiceSummary) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCustomerInvoiceSummary) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCustomerInvoiceSummary) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCustomerInvoiceSummary) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetOverdueInvoiceCount

`func (o *DtoCustomerInvoiceSummary) GetOverdueInvoiceCount() int32`

GetOverdueInvoiceCount returns the OverdueInvoiceCount field if non-nil, zero value otherwise.

### GetOverdueInvoiceCountOk

`func (o *DtoCustomerInvoiceSummary) GetOverdueInvoiceCountOk() (*int32, bool)`

GetOverdueInvoiceCountOk returns a tuple with the OverdueInvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueInvoiceCount

`func (o *DtoCustomerInvoiceSummary) SetOverdueInvoiceCount(v int32)`

SetOverdueInvoiceCount sets OverdueInvoiceCount field to given value.

### HasOverdueInvoiceCount

`func (o *DtoCustomerInvoiceSummary) HasOverdueInvoiceCount() bool`

HasOverdueInvoiceCount returns a boolean if a field has been set.

### GetTotalInvoiceCount

`func (o *DtoCustomerInvoiceSummary) GetTotalInvoiceCount() int32`

GetTotalInvoiceCount returns the TotalInvoiceCount field if non-nil, zero value otherwise.

### GetTotalInvoiceCountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalInvoiceCountOk() (*int32, bool)`

GetTotalInvoiceCountOk returns a tuple with the TotalInvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoiceCount

`func (o *DtoCustomerInvoiceSummary) SetTotalInvoiceCount(v int32)`

SetTotalInvoiceCount sets TotalInvoiceCount field to given value.

### HasTotalInvoiceCount

`func (o *DtoCustomerInvoiceSummary) HasTotalInvoiceCount() bool`

HasTotalInvoiceCount returns a boolean if a field has been set.

### GetTotalOverdueAmount

`func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmount() float32`

GetTotalOverdueAmount returns the TotalOverdueAmount field if non-nil, zero value otherwise.

### GetTotalOverdueAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmountOk() (*float32, bool)`

GetTotalOverdueAmountOk returns a tuple with the TotalOverdueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOverdueAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalOverdueAmount(v float32)`

SetTotalOverdueAmount sets TotalOverdueAmount field to given value.

### HasTotalOverdueAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalOverdueAmount() bool`

HasTotalOverdueAmount returns a boolean if a field has been set.

### GetTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmount() float32`

GetTotalRevenueAmount returns the TotalRevenueAmount field if non-nil, zero value otherwise.

### GetTotalRevenueAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmountOk() (*float32, bool)`

GetTotalRevenueAmountOk returns a tuple with the TotalRevenueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalRevenueAmount(v float32)`

SetTotalRevenueAmount sets TotalRevenueAmount field to given value.

### HasTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalRevenueAmount() bool`

HasTotalRevenueAmount returns a boolean if a field has been set.

### GetTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmount() float32`

GetTotalUnpaidAmount returns the TotalUnpaidAmount field if non-nil, zero value otherwise.

### GetTotalUnpaidAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmountOk() (*float32, bool)`

GetTotalUnpaidAmountOk returns a tuple with the TotalUnpaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalUnpaidAmount(v float32)`

SetTotalUnpaidAmount sets TotalUnpaidAmount field to given value.

### HasTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalUnpaidAmount() bool`

HasTotalUnpaidAmount returns a boolean if a field has been set.

### GetUnpaidFixedCharges

`func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedCharges() float32`

GetUnpaidFixedCharges returns the UnpaidFixedCharges field if non-nil, zero value otherwise.

### GetUnpaidFixedChargesOk

`func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedChargesOk() (*float32, bool)`

GetUnpaidFixedChargesOk returns a tuple with the UnpaidFixedCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidFixedCharges

`func (o *DtoCustomerInvoiceSummary) SetUnpaidFixedCharges(v float32)`

SetUnpaidFixedCharges sets UnpaidFixedCharges field to given value.

### HasUnpaidFixedCharges

`func (o *DtoCustomerInvoiceSummary) HasUnpaidFixedCharges() bool`

HasUnpaidFixedCharges returns a boolean if a field has been set.

### GetUnpaidInvoiceCount

`func (o *DtoCustomerInvoiceSummary) GetUnpaidInvoiceCount() int32`

GetUnpaidInvoiceCount returns the UnpaidInvoiceCount field if non-nil, zero value otherwise.

### GetUnpaidInvoiceCountOk

`func (o *DtoCustomerInvoiceSummary) GetUnpaidInvoiceCountOk() (*int32, bool)`

GetUnpaidInvoiceCountOk returns a tuple with the UnpaidInvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidInvoiceCount

`func (o *DtoCustomerInvoiceSummary) SetUnpaidInvoiceCount(v int32)`

SetUnpaidInvoiceCount sets UnpaidInvoiceCount field to given value.

### HasUnpaidInvoiceCount

`func (o *DtoCustomerInvoiceSummary) HasUnpaidInvoiceCount() bool`

HasUnpaidInvoiceCount returns a boolean if a field has been set.

### GetUnpaidUsageCharges

`func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageCharges() float32`

GetUnpaidUsageCharges returns the UnpaidUsageCharges field if non-nil, zero value otherwise.

### GetUnpaidUsageChargesOk

`func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageChargesOk() (*float32, bool)`

GetUnpaidUsageChargesOk returns a tuple with the UnpaidUsageCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidUsageCharges

`func (o *DtoCustomerInvoiceSummary) SetUnpaidUsageCharges(v float32)`

SetUnpaidUsageCharges sets UnpaidUsageCharges field to given value.

### HasUnpaidUsageCharges

`func (o *DtoCustomerInvoiceSummary) HasUnpaidUsageCharges() bool`

HasUnpaidUsageCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


