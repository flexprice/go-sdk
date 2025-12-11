# DtoPriceUnitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**PriceUnit** | **string** |  | 
**PriceUnitTiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) |  | [optional] 

## Methods

### NewDtoPriceUnitConfig

`func NewDtoPriceUnitConfig(priceUnit string, ) *DtoPriceUnitConfig`

NewDtoPriceUnitConfig instantiates a new DtoPriceUnitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPriceUnitConfigWithDefaults

`func NewDtoPriceUnitConfigWithDefaults() *DtoPriceUnitConfig`

NewDtoPriceUnitConfigWithDefaults instantiates a new DtoPriceUnitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoPriceUnitConfig) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoPriceUnitConfig) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoPriceUnitConfig) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoPriceUnitConfig) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPriceUnit

`func (o *DtoPriceUnitConfig) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *DtoPriceUnitConfig) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *DtoPriceUnitConfig) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.


### GetPriceUnitTiers

`func (o *DtoPriceUnitConfig) GetPriceUnitTiers() []DtoCreatePriceTier`

GetPriceUnitTiers returns the PriceUnitTiers field if non-nil, zero value otherwise.

### GetPriceUnitTiersOk

`func (o *DtoPriceUnitConfig) GetPriceUnitTiersOk() (*[]DtoCreatePriceTier, bool)`

GetPriceUnitTiersOk returns a tuple with the PriceUnitTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitTiers

`func (o *DtoPriceUnitConfig) SetPriceUnitTiers(v []DtoCreatePriceTier)`

SetPriceUnitTiers sets PriceUnitTiers field to given value.

### HasPriceUnitTiers

`func (o *DtoPriceUnitConfig) HasPriceUnitTiers() bool`

HasPriceUnitTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


