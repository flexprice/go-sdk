# DtoCreateCreditNoteLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | amount is the monetary amount to be credited for this line item | 
**DisplayName** | Pointer to **string** | display_name is an optional human-readable name for this credit note line item | [optional] 
**InvoiceLineItemId** | **string** | invoice_line_item_id is the unique identifier of the invoice line item being credited | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDtoCreateCreditNoteLineItemRequest

`func NewDtoCreateCreditNoteLineItemRequest(amount string, invoiceLineItemId string, ) *DtoCreateCreditNoteLineItemRequest`

NewDtoCreateCreditNoteLineItemRequest instantiates a new DtoCreateCreditNoteLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateCreditNoteLineItemRequestWithDefaults

`func NewDtoCreateCreditNoteLineItemRequestWithDefaults() *DtoCreateCreditNoteLineItemRequest`

NewDtoCreateCreditNoteLineItemRequestWithDefaults instantiates a new DtoCreateCreditNoteLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoCreateCreditNoteLineItemRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreateCreditNoteLineItemRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreateCreditNoteLineItemRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDisplayName

`func (o *DtoCreateCreditNoteLineItemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DtoCreateCreditNoteLineItemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DtoCreateCreditNoteLineItemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DtoCreateCreditNoteLineItemRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInvoiceLineItemId

`func (o *DtoCreateCreditNoteLineItemRequest) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *DtoCreateCreditNoteLineItemRequest) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *DtoCreateCreditNoteLineItemRequest) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.


### GetMetadata

`func (o *DtoCreateCreditNoteLineItemRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateCreditNoteLineItemRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateCreditNoteLineItemRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateCreditNoteLineItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


