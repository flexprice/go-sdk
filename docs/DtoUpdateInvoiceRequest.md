# DtoUpdateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **string** |  | [optional] 
**InvoicePdfUrl** | Pointer to **string** | invoice_pdf_url is the URL where customers can download the PDF version of this invoice | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDtoUpdateInvoiceRequest

`func NewDtoUpdateInvoiceRequest() *DtoUpdateInvoiceRequest`

NewDtoUpdateInvoiceRequest instantiates a new DtoUpdateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateInvoiceRequestWithDefaults

`func NewDtoUpdateInvoiceRequestWithDefaults() *DtoUpdateInvoiceRequest`

NewDtoUpdateInvoiceRequestWithDefaults instantiates a new DtoUpdateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *DtoUpdateInvoiceRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DtoUpdateInvoiceRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DtoUpdateInvoiceRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DtoUpdateInvoiceRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInvoicePdfUrl

`func (o *DtoUpdateInvoiceRequest) GetInvoicePdfUrl() string`

GetInvoicePdfUrl returns the InvoicePdfUrl field if non-nil, zero value otherwise.

### GetInvoicePdfUrlOk

`func (o *DtoUpdateInvoiceRequest) GetInvoicePdfUrlOk() (*string, bool)`

GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfUrl

`func (o *DtoUpdateInvoiceRequest) SetInvoicePdfUrl(v string)`

SetInvoicePdfUrl sets InvoicePdfUrl field to given value.

### HasInvoicePdfUrl

`func (o *DtoUpdateInvoiceRequest) HasInvoicePdfUrl() bool`

HasInvoicePdfUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateInvoiceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateInvoiceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateInvoiceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateInvoiceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


