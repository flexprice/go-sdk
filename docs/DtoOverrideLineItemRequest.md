# DtoOverrideLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is the new price amount that overrides the original price (optional) | [optional] 
**BillingModel** | Pointer to [**TypesBillingModel**](TypesBillingModel.md) |  | [optional] 
**PriceId** | **string** | PriceID references the plan price to override | 
**PriceUnitAmount** | Pointer to **string** | PriceUnitAmount is the amount of the price unit (for CUSTOM type, FLAT_FEE/PACKAGE billing models) | [optional] 
**PriceUnitTiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) | PriceUnitTiers are the tiers for the price unit (for CUSTOM type, TIERED billing model) | [optional] 
**Quantity** | Pointer to **string** | Quantity for this line item (optional) | [optional] 
**TierMode** | Pointer to [**TypesBillingTier**](TypesBillingTier.md) |  | [optional] 
**Tiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) | Tiers determines the pricing tiers for this line item | [optional] 
**TransformQuantity** | Pointer to [**PriceTransformQuantity**](PriceTransformQuantity.md) |  | [optional] 

## Methods

### NewDtoOverrideLineItemRequest

`func NewDtoOverrideLineItemRequest(priceId string, ) *DtoOverrideLineItemRequest`

NewDtoOverrideLineItemRequest instantiates a new DtoOverrideLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoOverrideLineItemRequestWithDefaults

`func NewDtoOverrideLineItemRequestWithDefaults() *DtoOverrideLineItemRequest`

NewDtoOverrideLineItemRequestWithDefaults instantiates a new DtoOverrideLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoOverrideLineItemRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoOverrideLineItemRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoOverrideLineItemRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoOverrideLineItemRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingModel

`func (o *DtoOverrideLineItemRequest) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *DtoOverrideLineItemRequest) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *DtoOverrideLineItemRequest) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *DtoOverrideLineItemRequest) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoOverrideLineItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoOverrideLineItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoOverrideLineItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetPriceUnitAmount

`func (o *DtoOverrideLineItemRequest) GetPriceUnitAmount() string`

GetPriceUnitAmount returns the PriceUnitAmount field if non-nil, zero value otherwise.

### GetPriceUnitAmountOk

`func (o *DtoOverrideLineItemRequest) GetPriceUnitAmountOk() (*string, bool)`

GetPriceUnitAmountOk returns a tuple with the PriceUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitAmount

`func (o *DtoOverrideLineItemRequest) SetPriceUnitAmount(v string)`

SetPriceUnitAmount sets PriceUnitAmount field to given value.

### HasPriceUnitAmount

`func (o *DtoOverrideLineItemRequest) HasPriceUnitAmount() bool`

HasPriceUnitAmount returns a boolean if a field has been set.

### GetPriceUnitTiers

`func (o *DtoOverrideLineItemRequest) GetPriceUnitTiers() []DtoCreatePriceTier`

GetPriceUnitTiers returns the PriceUnitTiers field if non-nil, zero value otherwise.

### GetPriceUnitTiersOk

`func (o *DtoOverrideLineItemRequest) GetPriceUnitTiersOk() (*[]DtoCreatePriceTier, bool)`

GetPriceUnitTiersOk returns a tuple with the PriceUnitTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitTiers

`func (o *DtoOverrideLineItemRequest) SetPriceUnitTiers(v []DtoCreatePriceTier)`

SetPriceUnitTiers sets PriceUnitTiers field to given value.

### HasPriceUnitTiers

`func (o *DtoOverrideLineItemRequest) HasPriceUnitTiers() bool`

HasPriceUnitTiers returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoOverrideLineItemRequest) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoOverrideLineItemRequest) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoOverrideLineItemRequest) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoOverrideLineItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTierMode

`func (o *DtoOverrideLineItemRequest) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *DtoOverrideLineItemRequest) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *DtoOverrideLineItemRequest) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *DtoOverrideLineItemRequest) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *DtoOverrideLineItemRequest) GetTiers() []DtoCreatePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *DtoOverrideLineItemRequest) GetTiersOk() (*[]DtoCreatePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *DtoOverrideLineItemRequest) SetTiers(v []DtoCreatePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *DtoOverrideLineItemRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *DtoOverrideLineItemRequest) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *DtoOverrideLineItemRequest) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *DtoOverrideLineItemRequest) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *DtoOverrideLineItemRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


