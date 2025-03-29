# DtoCustomerMultiCurrencyInvoiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**DefaultCurrency** | Pointer to **string** |  | [optional] 
**Summaries** | Pointer to [**[]DtoCustomerInvoiceSummary**](DtoCustomerInvoiceSummary.md) |  | [optional] 

## Methods

### NewDtoCustomerMultiCurrencyInvoiceSummary

`func NewDtoCustomerMultiCurrencyInvoiceSummary() *DtoCustomerMultiCurrencyInvoiceSummary`

NewDtoCustomerMultiCurrencyInvoiceSummary instantiates a new DtoCustomerMultiCurrencyInvoiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCustomerMultiCurrencyInvoiceSummaryWithDefaults

`func NewDtoCustomerMultiCurrencyInvoiceSummaryWithDefaults() *DtoCustomerMultiCurrencyInvoiceSummary`

NewDtoCustomerMultiCurrencyInvoiceSummaryWithDefaults instantiates a new DtoCustomerMultiCurrencyInvoiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetDefaultCurrency() string`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetDefaultCurrencyOk() (*string, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) SetDefaultCurrency(v string)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetSummaries

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetSummaries() []DtoCustomerInvoiceSummary`

GetSummaries returns the Summaries field if non-nil, zero value otherwise.

### GetSummariesOk

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) GetSummariesOk() (*[]DtoCustomerInvoiceSummary, bool)`

GetSummariesOk returns a tuple with the Summaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaries

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) SetSummaries(v []DtoCustomerInvoiceSummary)`

SetSummaries sets Summaries field to given value.

### HasSummaries

`func (o *DtoCustomerMultiCurrencyInvoiceSummary) HasSummaries() bool`

HasSummaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


