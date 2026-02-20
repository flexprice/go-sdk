# DtoSubscriptionLineItemLookupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorsErrorResponse**](ErrorsErrorResponse.md) |  | [optional] 
**MatchedLineItems** | Pointer to [**[]DtoMatchedSubscriptionLineItem**](DtoMatchedSubscriptionLineItem.md) |  | [optional] 
**Status** | Pointer to [**TypesDebugTrackerStatus**](TypesDebugTrackerStatus.md) |  | [optional] 

## Methods

### NewDtoSubscriptionLineItemLookupResult

`func NewDtoSubscriptionLineItemLookupResult() *DtoSubscriptionLineItemLookupResult`

NewDtoSubscriptionLineItemLookupResult instantiates a new DtoSubscriptionLineItemLookupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionLineItemLookupResultWithDefaults

`func NewDtoSubscriptionLineItemLookupResultWithDefaults() *DtoSubscriptionLineItemLookupResult`

NewDtoSubscriptionLineItemLookupResultWithDefaults instantiates a new DtoSubscriptionLineItemLookupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DtoSubscriptionLineItemLookupResult) GetError() ErrorsErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DtoSubscriptionLineItemLookupResult) GetErrorOk() (*ErrorsErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DtoSubscriptionLineItemLookupResult) SetError(v ErrorsErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *DtoSubscriptionLineItemLookupResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMatchedLineItems

`func (o *DtoSubscriptionLineItemLookupResult) GetMatchedLineItems() []DtoMatchedSubscriptionLineItem`

GetMatchedLineItems returns the MatchedLineItems field if non-nil, zero value otherwise.

### GetMatchedLineItemsOk

`func (o *DtoSubscriptionLineItemLookupResult) GetMatchedLineItemsOk() (*[]DtoMatchedSubscriptionLineItem, bool)`

GetMatchedLineItemsOk returns a tuple with the MatchedLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedLineItems

`func (o *DtoSubscriptionLineItemLookupResult) SetMatchedLineItems(v []DtoMatchedSubscriptionLineItem)`

SetMatchedLineItems sets MatchedLineItems field to given value.

### HasMatchedLineItems

`func (o *DtoSubscriptionLineItemLookupResult) HasMatchedLineItems() bool`

HasMatchedLineItems returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionLineItemLookupResult) GetStatus() TypesDebugTrackerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionLineItemLookupResult) GetStatusOk() (*TypesDebugTrackerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionLineItemLookupResult) SetStatus(v TypesDebugTrackerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionLineItemLookupResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


