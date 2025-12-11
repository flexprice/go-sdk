# TypesInvoiceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDueGt** | Pointer to **float32** | amount_due_gt filters invoices with a total amount due greater than the specified value Useful for finding invoices above a certain threshold or identifying high-value invoices | [optional] 
**AmountRemainingGt** | Pointer to **float32** | amount_remaining_gt filters invoices with an outstanding balance greater than the specified value Useful for finding invoices that still have significant unpaid amounts | [optional] 
**CustomerId** | Pointer to **string** | customer_id filters invoices for a specific customer using FlexPrice&#39;s internal customer ID This is the ID returned by FlexPrice when creating or retrieving customers | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** | external_customer_id filters invoices for a customer using your system&#39;s customer identifier This is the ID you provided when creating the customer in FlexPrice | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) |  | [optional] 
**InvoiceIds** | Pointer to **[]string** | invoice_ids restricts results to invoices with the specified IDs Use this to retrieve specific invoices when you know their exact identifiers | [optional] 
**InvoiceStatus** | Pointer to [**[]TypesInvoiceStatus**](TypesInvoiceStatus.md) | invoice_status filters by the current state of invoices in their lifecycle Multiple statuses can be specified to include invoices in any of the listed states | [optional] 
**InvoiceType** | Pointer to [**TypesInvoiceType**](TypesInvoiceType.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to [**[]TypesPaymentStatus**](TypesPaymentStatus.md) | payment_status filters by the payment state of invoices Multiple statuses can be specified to include invoices with any of the listed payment states | [optional] 
**SkipLineItems** | Pointer to **bool** | SkipLineItems if true, will not include line items in the response | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id filters invoices generated for a specific subscription Only returns invoices that were created as part of the specified subscription&#39;s billing | [optional] 

## Methods

### NewTypesInvoiceFilter

`func NewTypesInvoiceFilter() *TypesInvoiceFilter`

NewTypesInvoiceFilter instantiates a new TypesInvoiceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesInvoiceFilterWithDefaults

`func NewTypesInvoiceFilterWithDefaults() *TypesInvoiceFilter`

NewTypesInvoiceFilterWithDefaults instantiates a new TypesInvoiceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDueGt

`func (o *TypesInvoiceFilter) GetAmountDueGt() float32`

GetAmountDueGt returns the AmountDueGt field if non-nil, zero value otherwise.

### GetAmountDueGtOk

`func (o *TypesInvoiceFilter) GetAmountDueGtOk() (*float32, bool)`

GetAmountDueGtOk returns a tuple with the AmountDueGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueGt

`func (o *TypesInvoiceFilter) SetAmountDueGt(v float32)`

SetAmountDueGt sets AmountDueGt field to given value.

### HasAmountDueGt

`func (o *TypesInvoiceFilter) HasAmountDueGt() bool`

HasAmountDueGt returns a boolean if a field has been set.

### GetAmountRemainingGt

`func (o *TypesInvoiceFilter) GetAmountRemainingGt() float32`

GetAmountRemainingGt returns the AmountRemainingGt field if non-nil, zero value otherwise.

### GetAmountRemainingGtOk

`func (o *TypesInvoiceFilter) GetAmountRemainingGtOk() (*float32, bool)`

GetAmountRemainingGtOk returns a tuple with the AmountRemainingGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemainingGt

`func (o *TypesInvoiceFilter) SetAmountRemainingGt(v float32)`

SetAmountRemainingGt sets AmountRemainingGt field to given value.

### HasAmountRemainingGt

`func (o *TypesInvoiceFilter) HasAmountRemainingGt() bool`

HasAmountRemainingGt returns a boolean if a field has been set.

### GetCustomerId

`func (o *TypesInvoiceFilter) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TypesInvoiceFilter) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TypesInvoiceFilter) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TypesInvoiceFilter) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesInvoiceFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesInvoiceFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesInvoiceFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesInvoiceFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesInvoiceFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesInvoiceFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesInvoiceFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesInvoiceFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *TypesInvoiceFilter) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *TypesInvoiceFilter) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *TypesInvoiceFilter) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *TypesInvoiceFilter) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetFilters

`func (o *TypesInvoiceFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesInvoiceFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesInvoiceFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesInvoiceFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInvoiceIds

`func (o *TypesInvoiceFilter) GetInvoiceIds() []string`

GetInvoiceIds returns the InvoiceIds field if non-nil, zero value otherwise.

### GetInvoiceIdsOk

`func (o *TypesInvoiceFilter) GetInvoiceIdsOk() (*[]string, bool)`

GetInvoiceIdsOk returns a tuple with the InvoiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIds

`func (o *TypesInvoiceFilter) SetInvoiceIds(v []string)`

SetInvoiceIds sets InvoiceIds field to given value.

### HasInvoiceIds

`func (o *TypesInvoiceFilter) HasInvoiceIds() bool`

HasInvoiceIds returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *TypesInvoiceFilter) GetInvoiceStatus() []TypesInvoiceStatus`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *TypesInvoiceFilter) GetInvoiceStatusOk() (*[]TypesInvoiceStatus, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *TypesInvoiceFilter) SetInvoiceStatus(v []TypesInvoiceStatus)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *TypesInvoiceFilter) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceType

`func (o *TypesInvoiceFilter) GetInvoiceType() TypesInvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *TypesInvoiceFilter) GetInvoiceTypeOk() (*TypesInvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *TypesInvoiceFilter) SetInvoiceType(v TypesInvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *TypesInvoiceFilter) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetLimit

`func (o *TypesInvoiceFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesInvoiceFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesInvoiceFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesInvoiceFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesInvoiceFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesInvoiceFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesInvoiceFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesInvoiceFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesInvoiceFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesInvoiceFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesInvoiceFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesInvoiceFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *TypesInvoiceFilter) GetPaymentStatus() []TypesPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *TypesInvoiceFilter) GetPaymentStatusOk() (*[]TypesPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *TypesInvoiceFilter) SetPaymentStatus(v []TypesPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *TypesInvoiceFilter) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetSkipLineItems

`func (o *TypesInvoiceFilter) GetSkipLineItems() bool`

GetSkipLineItems returns the SkipLineItems field if non-nil, zero value otherwise.

### GetSkipLineItemsOk

`func (o *TypesInvoiceFilter) GetSkipLineItemsOk() (*bool, bool)`

GetSkipLineItemsOk returns a tuple with the SkipLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipLineItems

`func (o *TypesInvoiceFilter) SetSkipLineItems(v bool)`

SetSkipLineItems sets SkipLineItems field to given value.

### HasSkipLineItems

`func (o *TypesInvoiceFilter) HasSkipLineItems() bool`

HasSkipLineItems returns a boolean if a field has been set.

### GetSort

`func (o *TypesInvoiceFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesInvoiceFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesInvoiceFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesInvoiceFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesInvoiceFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesInvoiceFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesInvoiceFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesInvoiceFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesInvoiceFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesInvoiceFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesInvoiceFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesInvoiceFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *TypesInvoiceFilter) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TypesInvoiceFilter) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TypesInvoiceFilter) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TypesInvoiceFilter) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


