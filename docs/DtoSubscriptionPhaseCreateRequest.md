# DtoSubscriptionPhaseCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coupons** | Pointer to **[]string** | Coupons represents subscription-level coupons to be applied to this phase | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**LineItemCoupons** | Pointer to **map[string][]string** | LineItemCoupons represents line item-level coupons (map of line_item_id to coupon IDs) | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OverrideLineItems** | Pointer to [**[]DtoOverrideLineItemRequest**](DtoOverrideLineItemRequest.md) | OverrideLineItems allows customizing specific prices for this phase If not provided, phase will use the same line items as the subscription (plan prices) | [optional] 
**StartDate** | **string** |  | 

## Methods

### NewDtoSubscriptionPhaseCreateRequest

`func NewDtoSubscriptionPhaseCreateRequest(startDate string, ) *DtoSubscriptionPhaseCreateRequest`

NewDtoSubscriptionPhaseCreateRequest instantiates a new DtoSubscriptionPhaseCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionPhaseCreateRequestWithDefaults

`func NewDtoSubscriptionPhaseCreateRequestWithDefaults() *DtoSubscriptionPhaseCreateRequest`

NewDtoSubscriptionPhaseCreateRequestWithDefaults instantiates a new DtoSubscriptionPhaseCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoSubscriptionPhaseCreateRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoSubscriptionPhaseCreateRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoSubscriptionPhaseCreateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLineItemCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) GetLineItemCoupons() map[string][]string`

GetLineItemCoupons returns the LineItemCoupons field if non-nil, zero value otherwise.

### GetLineItemCouponsOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetLineItemCouponsOk() (*map[string][]string, bool)`

GetLineItemCouponsOk returns a tuple with the LineItemCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) SetLineItemCoupons(v map[string][]string)`

SetLineItemCoupons sets LineItemCoupons field to given value.

### HasLineItemCoupons

`func (o *DtoSubscriptionPhaseCreateRequest) HasLineItemCoupons() bool`

HasLineItemCoupons returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionPhaseCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionPhaseCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionPhaseCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverrideLineItems

`func (o *DtoSubscriptionPhaseCreateRequest) GetOverrideLineItems() []DtoOverrideLineItemRequest`

GetOverrideLineItems returns the OverrideLineItems field if non-nil, zero value otherwise.

### GetOverrideLineItemsOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetOverrideLineItemsOk() (*[]DtoOverrideLineItemRequest, bool)`

GetOverrideLineItemsOk returns a tuple with the OverrideLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLineItems

`func (o *DtoSubscriptionPhaseCreateRequest) SetOverrideLineItems(v []DtoOverrideLineItemRequest)`

SetOverrideLineItems sets OverrideLineItems field to given value.

### HasOverrideLineItems

`func (o *DtoSubscriptionPhaseCreateRequest) HasOverrideLineItems() bool`

HasOverrideLineItems returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoSubscriptionPhaseCreateRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoSubscriptionPhaseCreateRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoSubscriptionPhaseCreateRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


