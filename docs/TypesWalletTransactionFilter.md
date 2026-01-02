# TypesWalletTransactionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditsAvailableGt** | Pointer to **float32** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**ExpiryDateAfter** | Pointer to **string** |  | [optional] 
**ExpiryDateBefore** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | filters allows complex filtering based on multiple fields | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceType** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TransactionReason** | Pointer to [**TypesTransactionReason**](TypesTransactionReason.md) |  | [optional] 
**TransactionStatus** | Pointer to [**TypesTransactionStatus**](TypesTransactionStatus.md) |  | [optional] 
**Type** | Pointer to [**TypesTransactionType**](TypesTransactionType.md) |  | [optional] 

## Methods

### NewTypesWalletTransactionFilter

`func NewTypesWalletTransactionFilter() *TypesWalletTransactionFilter`

NewTypesWalletTransactionFilter instantiates a new TypesWalletTransactionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesWalletTransactionFilterWithDefaults

`func NewTypesWalletTransactionFilterWithDefaults() *TypesWalletTransactionFilter`

NewTypesWalletTransactionFilterWithDefaults instantiates a new TypesWalletTransactionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *TypesWalletTransactionFilter) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TypesWalletTransactionFilter) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TypesWalletTransactionFilter) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TypesWalletTransactionFilter) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditsAvailableGt

`func (o *TypesWalletTransactionFilter) GetCreditsAvailableGt() float32`

GetCreditsAvailableGt returns the CreditsAvailableGt field if non-nil, zero value otherwise.

### GetCreditsAvailableGtOk

`func (o *TypesWalletTransactionFilter) GetCreditsAvailableGtOk() (*float32, bool)`

GetCreditsAvailableGtOk returns a tuple with the CreditsAvailableGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAvailableGt

`func (o *TypesWalletTransactionFilter) SetCreditsAvailableGt(v float32)`

SetCreditsAvailableGt sets CreditsAvailableGt field to given value.

### HasCreditsAvailableGt

`func (o *TypesWalletTransactionFilter) HasCreditsAvailableGt() bool`

HasCreditsAvailableGt returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesWalletTransactionFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesWalletTransactionFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesWalletTransactionFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesWalletTransactionFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesWalletTransactionFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesWalletTransactionFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesWalletTransactionFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesWalletTransactionFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExpiryDateAfter

`func (o *TypesWalletTransactionFilter) GetExpiryDateAfter() string`

GetExpiryDateAfter returns the ExpiryDateAfter field if non-nil, zero value otherwise.

### GetExpiryDateAfterOk

`func (o *TypesWalletTransactionFilter) GetExpiryDateAfterOk() (*string, bool)`

GetExpiryDateAfterOk returns a tuple with the ExpiryDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateAfter

`func (o *TypesWalletTransactionFilter) SetExpiryDateAfter(v string)`

SetExpiryDateAfter sets ExpiryDateAfter field to given value.

### HasExpiryDateAfter

`func (o *TypesWalletTransactionFilter) HasExpiryDateAfter() bool`

HasExpiryDateAfter returns a boolean if a field has been set.

### GetExpiryDateBefore

`func (o *TypesWalletTransactionFilter) GetExpiryDateBefore() string`

GetExpiryDateBefore returns the ExpiryDateBefore field if non-nil, zero value otherwise.

### GetExpiryDateBeforeOk

`func (o *TypesWalletTransactionFilter) GetExpiryDateBeforeOk() (*string, bool)`

GetExpiryDateBeforeOk returns a tuple with the ExpiryDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateBefore

`func (o *TypesWalletTransactionFilter) SetExpiryDateBefore(v string)`

SetExpiryDateBefore sets ExpiryDateBefore field to given value.

### HasExpiryDateBefore

`func (o *TypesWalletTransactionFilter) HasExpiryDateBefore() bool`

HasExpiryDateBefore returns a boolean if a field has been set.

### GetFilters

`func (o *TypesWalletTransactionFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesWalletTransactionFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesWalletTransactionFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesWalletTransactionFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *TypesWalletTransactionFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesWalletTransactionFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesWalletTransactionFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TypesWalletTransactionFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimit

`func (o *TypesWalletTransactionFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesWalletTransactionFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesWalletTransactionFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesWalletTransactionFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesWalletTransactionFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesWalletTransactionFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesWalletTransactionFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesWalletTransactionFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesWalletTransactionFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesWalletTransactionFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesWalletTransactionFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesWalletTransactionFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPriority

`func (o *TypesWalletTransactionFilter) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TypesWalletTransactionFilter) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TypesWalletTransactionFilter) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TypesWalletTransactionFilter) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReferenceId

`func (o *TypesWalletTransactionFilter) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TypesWalletTransactionFilter) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TypesWalletTransactionFilter) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *TypesWalletTransactionFilter) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceType

`func (o *TypesWalletTransactionFilter) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *TypesWalletTransactionFilter) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *TypesWalletTransactionFilter) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *TypesWalletTransactionFilter) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetSort

`func (o *TypesWalletTransactionFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesWalletTransactionFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesWalletTransactionFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesWalletTransactionFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesWalletTransactionFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesWalletTransactionFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesWalletTransactionFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesWalletTransactionFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesWalletTransactionFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesWalletTransactionFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesWalletTransactionFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesWalletTransactionFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionReason

`func (o *TypesWalletTransactionFilter) GetTransactionReason() TypesTransactionReason`

GetTransactionReason returns the TransactionReason field if non-nil, zero value otherwise.

### GetTransactionReasonOk

`func (o *TypesWalletTransactionFilter) GetTransactionReasonOk() (*TypesTransactionReason, bool)`

GetTransactionReasonOk returns a tuple with the TransactionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReason

`func (o *TypesWalletTransactionFilter) SetTransactionReason(v TypesTransactionReason)`

SetTransactionReason sets TransactionReason field to given value.

### HasTransactionReason

`func (o *TypesWalletTransactionFilter) HasTransactionReason() bool`

HasTransactionReason returns a boolean if a field has been set.

### GetTransactionStatus

`func (o *TypesWalletTransactionFilter) GetTransactionStatus() TypesTransactionStatus`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *TypesWalletTransactionFilter) GetTransactionStatusOk() (*TypesTransactionStatus, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *TypesWalletTransactionFilter) SetTransactionStatus(v TypesTransactionStatus)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *TypesWalletTransactionFilter) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.

### GetType

`func (o *TypesWalletTransactionFilter) GetType() TypesTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypesWalletTransactionFilter) GetTypeOk() (*TypesTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypesWalletTransactionFilter) SetType(v TypesTransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *TypesWalletTransactionFilter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


