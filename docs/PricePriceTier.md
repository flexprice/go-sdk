# PricePriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlatAmount** | Pointer to **float32** | flat_amount is the flat amount for the given tier (optional) Applied on top of unit_amount*quantity. Useful for cases like \&quot;2.7$ + 5c\&quot; | [optional] 
**UnitAmount** | Pointer to **float32** | unit_amount is the amount per unit for the given tier | [optional] 
**UpTo** | Pointer to **int32** | up_to is the quantity up to which this tier applies. It is null for the last tier. IMPORTANT: Tier boundaries are INCLUSIVE. - If up_to is 1000, then quantity less than or equal to 1000 belongs to this tier - This behavior is consistent across both VOLUME and SLAB tier modes | [optional] 

## Methods

### NewPricePriceTier

`func NewPricePriceTier() *PricePriceTier`

NewPricePriceTier instantiates a new PricePriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricePriceTierWithDefaults

`func NewPricePriceTierWithDefaults() *PricePriceTier`

NewPricePriceTierWithDefaults instantiates a new PricePriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlatAmount

`func (o *PricePriceTier) GetFlatAmount() float32`

GetFlatAmount returns the FlatAmount field if non-nil, zero value otherwise.

### GetFlatAmountOk

`func (o *PricePriceTier) GetFlatAmountOk() (*float32, bool)`

GetFlatAmountOk returns a tuple with the FlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmount

`func (o *PricePriceTier) SetFlatAmount(v float32)`

SetFlatAmount sets FlatAmount field to given value.

### HasFlatAmount

`func (o *PricePriceTier) HasFlatAmount() bool`

HasFlatAmount returns a boolean if a field has been set.

### GetUnitAmount

`func (o *PricePriceTier) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PricePriceTier) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PricePriceTier) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *PricePriceTier) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetUpTo

`func (o *PricePriceTier) GetUpTo() int32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *PricePriceTier) GetUpToOk() (*int32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *PricePriceTier) SetUpTo(v int32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *PricePriceTier) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


