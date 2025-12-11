# DtoUpdateSubscriptionLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is the new price amount that overrides the original price | [optional] 
**BillingModel** | Pointer to [**TypesBillingModel**](TypesBillingModel.md) |  | [optional] 
**EffectiveFrom** | Pointer to **string** | EffectiveFrom for the existing line item (if not provided, defaults to now) | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata for the new line item | [optional] 
**TierMode** | Pointer to [**TypesBillingTier**](TypesBillingTier.md) |  | [optional] 
**Tiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) | Tiers determines the pricing tiers for this line item | [optional] 
**TransformQuantity** | Pointer to [**PriceTransformQuantity**](PriceTransformQuantity.md) |  | [optional] 

## Methods

### NewDtoUpdateSubscriptionLineItemRequest

`func NewDtoUpdateSubscriptionLineItemRequest() *DtoUpdateSubscriptionLineItemRequest`

NewDtoUpdateSubscriptionLineItemRequest instantiates a new DtoUpdateSubscriptionLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateSubscriptionLineItemRequestWithDefaults

`func NewDtoUpdateSubscriptionLineItemRequestWithDefaults() *DtoUpdateSubscriptionLineItemRequest`

NewDtoUpdateSubscriptionLineItemRequestWithDefaults instantiates a new DtoUpdateSubscriptionLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoUpdateSubscriptionLineItemRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoUpdateSubscriptionLineItemRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoUpdateSubscriptionLineItemRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingModel

`func (o *DtoUpdateSubscriptionLineItemRequest) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *DtoUpdateSubscriptionLineItemRequest) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *DtoUpdateSubscriptionLineItemRequest) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *DtoUpdateSubscriptionLineItemRequest) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *DtoUpdateSubscriptionLineItemRequest) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *DtoUpdateSubscriptionLineItemRequest) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateSubscriptionLineItemRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateSubscriptionLineItemRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateSubscriptionLineItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTierMode

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *DtoUpdateSubscriptionLineItemRequest) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *DtoUpdateSubscriptionLineItemRequest) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTiers() []DtoCreatePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTiersOk() (*[]DtoCreatePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *DtoUpdateSubscriptionLineItemRequest) SetTiers(v []DtoCreatePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *DtoUpdateSubscriptionLineItemRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *DtoUpdateSubscriptionLineItemRequest) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *DtoUpdateSubscriptionLineItemRequest) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *DtoUpdateSubscriptionLineItemRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


