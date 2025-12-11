# DtoInvoicePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | currency is the currency for all amounts | [optional] 
**DueDate** | Pointer to **string** | due_date is when the invoice would be due | [optional] 
**LineItems** | Pointer to [**[]DtoInvoiceLineItemPreview**](DtoInvoiceLineItemPreview.md) | line_items contains preview of line items | [optional] 
**Subtotal** | Pointer to **string** | subtotal is the subtotal amount before taxes | [optional] 
**TaxAmount** | Pointer to **string** | tax_amount is the total tax amount | [optional] 
**Total** | Pointer to **string** | total is the total amount including taxes | [optional] 

## Methods

### NewDtoInvoicePreview

`func NewDtoInvoicePreview() *DtoInvoicePreview`

NewDtoInvoicePreview instantiates a new DtoInvoicePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoInvoicePreviewWithDefaults

`func NewDtoInvoicePreviewWithDefaults() *DtoInvoicePreview`

NewDtoInvoicePreviewWithDefaults instantiates a new DtoInvoicePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *DtoInvoicePreview) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoInvoicePreview) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoInvoicePreview) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoInvoicePreview) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueDate

`func (o *DtoInvoicePreview) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DtoInvoicePreview) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DtoInvoicePreview) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DtoInvoicePreview) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoInvoicePreview) GetLineItems() []DtoInvoiceLineItemPreview`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoInvoicePreview) GetLineItemsOk() (*[]DtoInvoiceLineItemPreview, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoInvoicePreview) SetLineItems(v []DtoInvoiceLineItemPreview)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoInvoicePreview) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetSubtotal

`func (o *DtoInvoicePreview) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *DtoInvoicePreview) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *DtoInvoicePreview) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *DtoInvoicePreview) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTaxAmount

`func (o *DtoInvoicePreview) GetTaxAmount() string`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *DtoInvoicePreview) GetTaxAmountOk() (*string, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *DtoInvoicePreview) SetTaxAmount(v string)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *DtoInvoicePreview) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotal

`func (o *DtoInvoicePreview) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoInvoicePreview) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoInvoicePreview) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoInvoicePreview) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


