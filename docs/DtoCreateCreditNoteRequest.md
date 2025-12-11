# DtoCreateCreditNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNoteNumber** | Pointer to **string** | credit_note_number is an optional human-readable identifier for the credit note | [optional] 
**IdempotencyKey** | Pointer to **string** | idempotency_key is an optional key used to prevent duplicate credit note creation | [optional] 
**InvoiceId** | **string** | invoice_id is the unique identifier of the invoice this credit note is applied to | 
**LineItems** | Pointer to [**[]DtoCreateCreditNoteLineItemRequest**](DtoCreateCreditNoteLineItemRequest.md) | line_items contains the individual line items that make up this credit note (minimum 1 required) | [optional] 
**Memo** | Pointer to **string** | memo is an optional free-text field for additional notes about the credit note | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**ProcessCreditNote** | Pointer to **bool** | process_credit_note is a flag to process the credit note after creation | [optional] [default to true]
**Reason** | [**TypesCreditNoteReason**](TypesCreditNoteReason.md) |  | 

## Methods

### NewDtoCreateCreditNoteRequest

`func NewDtoCreateCreditNoteRequest(invoiceId string, reason TypesCreditNoteReason, ) *DtoCreateCreditNoteRequest`

NewDtoCreateCreditNoteRequest instantiates a new DtoCreateCreditNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateCreditNoteRequestWithDefaults

`func NewDtoCreateCreditNoteRequestWithDefaults() *DtoCreateCreditNoteRequest`

NewDtoCreateCreditNoteRequestWithDefaults instantiates a new DtoCreateCreditNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNoteNumber

`func (o *DtoCreateCreditNoteRequest) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *DtoCreateCreditNoteRequest) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *DtoCreateCreditNoteRequest) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumber

`func (o *DtoCreateCreditNoteRequest) HasCreditNoteNumber() bool`

HasCreditNoteNumber returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoCreateCreditNoteRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoCreateCreditNoteRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoCreateCreditNoteRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoCreateCreditNoteRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInvoiceId

`func (o *DtoCreateCreditNoteRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DtoCreateCreditNoteRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DtoCreateCreditNoteRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetLineItems

`func (o *DtoCreateCreditNoteRequest) GetLineItems() []DtoCreateCreditNoteLineItemRequest`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoCreateCreditNoteRequest) GetLineItemsOk() (*[]DtoCreateCreditNoteLineItemRequest, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoCreateCreditNoteRequest) SetLineItems(v []DtoCreateCreditNoteLineItemRequest)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoCreateCreditNoteRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMemo

`func (o *DtoCreateCreditNoteRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DtoCreateCreditNoteRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DtoCreateCreditNoteRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DtoCreateCreditNoteRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateCreditNoteRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateCreditNoteRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateCreditNoteRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateCreditNoteRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessCreditNote

`func (o *DtoCreateCreditNoteRequest) GetProcessCreditNote() bool`

GetProcessCreditNote returns the ProcessCreditNote field if non-nil, zero value otherwise.

### GetProcessCreditNoteOk

`func (o *DtoCreateCreditNoteRequest) GetProcessCreditNoteOk() (*bool, bool)`

GetProcessCreditNoteOk returns a tuple with the ProcessCreditNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessCreditNote

`func (o *DtoCreateCreditNoteRequest) SetProcessCreditNote(v bool)`

SetProcessCreditNote sets ProcessCreditNote field to given value.

### HasProcessCreditNote

`func (o *DtoCreateCreditNoteRequest) HasProcessCreditNote() bool`

HasProcessCreditNote returns a boolean if a field has been set.

### GetReason

`func (o *DtoCreateCreditNoteRequest) GetReason() TypesCreditNoteReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoCreateCreditNoteRequest) GetReasonOk() (*TypesCreditNoteReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoCreateCreditNoteRequest) SetReason(v TypesCreditNoteReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


