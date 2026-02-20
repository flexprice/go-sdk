# DtoMatchedSubscriptionLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **string** |  | [optional] 
**IsActiveForEvent** | Pointer to **bool** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**SubLineItemId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SubscriptionLineItem** | Pointer to [**SubscriptionSubscriptionLineItem**](SubscriptionSubscriptionLineItem.md) |  | [optional] 
**TimestampWithinRange** | Pointer to **bool** |  | [optional] 

## Methods

### NewDtoMatchedSubscriptionLineItem

`func NewDtoMatchedSubscriptionLineItem() *DtoMatchedSubscriptionLineItem`

NewDtoMatchedSubscriptionLineItem instantiates a new DtoMatchedSubscriptionLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoMatchedSubscriptionLineItemWithDefaults

`func NewDtoMatchedSubscriptionLineItemWithDefaults() *DtoMatchedSubscriptionLineItem`

NewDtoMatchedSubscriptionLineItemWithDefaults instantiates a new DtoMatchedSubscriptionLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *DtoMatchedSubscriptionLineItem) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoMatchedSubscriptionLineItem) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoMatchedSubscriptionLineItem) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoMatchedSubscriptionLineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActiveForEvent

`func (o *DtoMatchedSubscriptionLineItem) GetIsActiveForEvent() bool`

GetIsActiveForEvent returns the IsActiveForEvent field if non-nil, zero value otherwise.

### GetIsActiveForEventOk

`func (o *DtoMatchedSubscriptionLineItem) GetIsActiveForEventOk() (*bool, bool)`

GetIsActiveForEventOk returns a tuple with the IsActiveForEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveForEvent

`func (o *DtoMatchedSubscriptionLineItem) SetIsActiveForEvent(v bool)`

SetIsActiveForEvent sets IsActiveForEvent field to given value.

### HasIsActiveForEvent

`func (o *DtoMatchedSubscriptionLineItem) HasIsActiveForEvent() bool`

HasIsActiveForEvent returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoMatchedSubscriptionLineItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoMatchedSubscriptionLineItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoMatchedSubscriptionLineItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoMatchedSubscriptionLineItem) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoMatchedSubscriptionLineItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoMatchedSubscriptionLineItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoMatchedSubscriptionLineItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoMatchedSubscriptionLineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubLineItemId

`func (o *DtoMatchedSubscriptionLineItem) GetSubLineItemId() string`

GetSubLineItemId returns the SubLineItemId field if non-nil, zero value otherwise.

### GetSubLineItemIdOk

`func (o *DtoMatchedSubscriptionLineItem) GetSubLineItemIdOk() (*string, bool)`

GetSubLineItemIdOk returns a tuple with the SubLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLineItemId

`func (o *DtoMatchedSubscriptionLineItem) SetSubLineItemId(v string)`

SetSubLineItemId sets SubLineItemId field to given value.

### HasSubLineItemId

`func (o *DtoMatchedSubscriptionLineItem) HasSubLineItemId() bool`

HasSubLineItemId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoMatchedSubscriptionLineItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoMatchedSubscriptionLineItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoMatchedSubscriptionLineItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoMatchedSubscriptionLineItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionLineItem

`func (o *DtoMatchedSubscriptionLineItem) GetSubscriptionLineItem() SubscriptionSubscriptionLineItem`

GetSubscriptionLineItem returns the SubscriptionLineItem field if non-nil, zero value otherwise.

### GetSubscriptionLineItemOk

`func (o *DtoMatchedSubscriptionLineItem) GetSubscriptionLineItemOk() (*SubscriptionSubscriptionLineItem, bool)`

GetSubscriptionLineItemOk returns a tuple with the SubscriptionLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLineItem

`func (o *DtoMatchedSubscriptionLineItem) SetSubscriptionLineItem(v SubscriptionSubscriptionLineItem)`

SetSubscriptionLineItem sets SubscriptionLineItem field to given value.

### HasSubscriptionLineItem

`func (o *DtoMatchedSubscriptionLineItem) HasSubscriptionLineItem() bool`

HasSubscriptionLineItem returns a boolean if a field has been set.

### GetTimestampWithinRange

`func (o *DtoMatchedSubscriptionLineItem) GetTimestampWithinRange() bool`

GetTimestampWithinRange returns the TimestampWithinRange field if non-nil, zero value otherwise.

### GetTimestampWithinRangeOk

`func (o *DtoMatchedSubscriptionLineItem) GetTimestampWithinRangeOk() (*bool, bool)`

GetTimestampWithinRangeOk returns a tuple with the TimestampWithinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampWithinRange

`func (o *DtoMatchedSubscriptionLineItem) SetTimestampWithinRange(v bool)`

SetTimestampWithinRange sets TimestampWithinRange field to given value.

### HasTimestampWithinRange

`func (o *DtoMatchedSubscriptionLineItem) HasTimestampWithinRange() bool`

HasTimestampWithinRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


