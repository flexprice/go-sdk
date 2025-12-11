# DtoCustomerInvoiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | currency is the three-letter ISO currency code for this summary | [optional] 
**CustomerId** | Pointer to **string** | customer_id is the unique identifier of the customer | [optional] 
**OverdueInvoiceCount** | Pointer to **int32** | overdue_invoice_count is the number of overdue invoices for this customer in this currency | [optional] 
**TotalInvoiceCount** | Pointer to **int32** | total_invoice_count is the total number of invoices for this customer in this currency | [optional] 
**TotalOverdueAmount** | Pointer to **string** | total_overdue_amount is the total amount of overdue invoices in this currency | [optional] 
**TotalRevenueAmount** | Pointer to **string** | total_revenue_amount is the total revenue generated from this customer in this currency | [optional] 
**TotalUnpaidAmount** | Pointer to **string** | total_unpaid_amount is the total amount of unpaid invoices in this currency | [optional] 
**UnpaidFixedCharges** | Pointer to **string** | unpaid_fixed_charges is the total amount of unpaid fixed charges in this currency | [optional] 
**UnpaidInvoiceCount** | Pointer to **int32** | unpaid_invoice_count is the number of unpaid invoices for this customer in this currency | [optional] 
**UnpaidUsageCharges** | Pointer to **string** | unpaid_usage_charges is the total amount of unpaid usage-based charges in this currency | [optional] 

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

`func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmount() string`

GetTotalOverdueAmount returns the TotalOverdueAmount field if non-nil, zero value otherwise.

### GetTotalOverdueAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmountOk() (*string, bool)`

GetTotalOverdueAmountOk returns a tuple with the TotalOverdueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOverdueAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalOverdueAmount(v string)`

SetTotalOverdueAmount sets TotalOverdueAmount field to given value.

### HasTotalOverdueAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalOverdueAmount() bool`

HasTotalOverdueAmount returns a boolean if a field has been set.

### GetTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmount() string`

GetTotalRevenueAmount returns the TotalRevenueAmount field if non-nil, zero value otherwise.

### GetTotalRevenueAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmountOk() (*string, bool)`

GetTotalRevenueAmountOk returns a tuple with the TotalRevenueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalRevenueAmount(v string)`

SetTotalRevenueAmount sets TotalRevenueAmount field to given value.

### HasTotalRevenueAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalRevenueAmount() bool`

HasTotalRevenueAmount returns a boolean if a field has been set.

### GetTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmount() string`

GetTotalUnpaidAmount returns the TotalUnpaidAmount field if non-nil, zero value otherwise.

### GetTotalUnpaidAmountOk

`func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmountOk() (*string, bool)`

GetTotalUnpaidAmountOk returns a tuple with the TotalUnpaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) SetTotalUnpaidAmount(v string)`

SetTotalUnpaidAmount sets TotalUnpaidAmount field to given value.

### HasTotalUnpaidAmount

`func (o *DtoCustomerInvoiceSummary) HasTotalUnpaidAmount() bool`

HasTotalUnpaidAmount returns a boolean if a field has been set.

### GetUnpaidFixedCharges

`func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedCharges() string`

GetUnpaidFixedCharges returns the UnpaidFixedCharges field if non-nil, zero value otherwise.

### GetUnpaidFixedChargesOk

`func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedChargesOk() (*string, bool)`

GetUnpaidFixedChargesOk returns a tuple with the UnpaidFixedCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidFixedCharges

`func (o *DtoCustomerInvoiceSummary) SetUnpaidFixedCharges(v string)`

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

`func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageCharges() string`

GetUnpaidUsageCharges returns the UnpaidUsageCharges field if non-nil, zero value otherwise.

### GetUnpaidUsageChargesOk

`func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageChargesOk() (*string, bool)`

GetUnpaidUsageChargesOk returns a tuple with the UnpaidUsageCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidUsageCharges

`func (o *DtoCustomerInvoiceSummary) SetUnpaidUsageCharges(v string)`

SetUnpaidUsageCharges sets UnpaidUsageCharges field to given value.

### HasUnpaidUsageCharges

`func (o *DtoCustomerInvoiceSummary) HasUnpaidUsageCharges() bool`

HasUnpaidUsageCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


