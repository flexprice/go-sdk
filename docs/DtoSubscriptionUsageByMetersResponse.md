# DtoSubscriptionUsageByMetersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayAmount** | Pointer to **string** |  | [optional] 
**FilterValues** | Pointer to **map[string][]string** |  | [optional] 
**IsOverage** | Pointer to **bool** | Whether this charge is at overage rate | [optional] 
**MeterDisplayName** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**OverageFactor** | Pointer to **float32** | Factor applied to this charge if in overage | [optional] 
**Price** | Pointer to [**PricePrice**](PricePrice.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 

## Methods

### NewDtoSubscriptionUsageByMetersResponse

`func NewDtoSubscriptionUsageByMetersResponse() *DtoSubscriptionUsageByMetersResponse`

NewDtoSubscriptionUsageByMetersResponse instantiates a new DtoSubscriptionUsageByMetersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionUsageByMetersResponseWithDefaults

`func NewDtoSubscriptionUsageByMetersResponseWithDefaults() *DtoSubscriptionUsageByMetersResponse`

NewDtoSubscriptionUsageByMetersResponseWithDefaults instantiates a new DtoSubscriptionUsageByMetersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoSubscriptionUsageByMetersResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoSubscriptionUsageByMetersResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoSubscriptionUsageByMetersResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoSubscriptionUsageByMetersResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoSubscriptionUsageByMetersResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoSubscriptionUsageByMetersResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayAmount

`func (o *DtoSubscriptionUsageByMetersResponse) GetDisplayAmount() string`

GetDisplayAmount returns the DisplayAmount field if non-nil, zero value otherwise.

### GetDisplayAmountOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetDisplayAmountOk() (*string, bool)`

GetDisplayAmountOk returns a tuple with the DisplayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAmount

`func (o *DtoSubscriptionUsageByMetersResponse) SetDisplayAmount(v string)`

SetDisplayAmount sets DisplayAmount field to given value.

### HasDisplayAmount

`func (o *DtoSubscriptionUsageByMetersResponse) HasDisplayAmount() bool`

HasDisplayAmount returns a boolean if a field has been set.

### GetFilterValues

`func (o *DtoSubscriptionUsageByMetersResponse) GetFilterValues() map[string][]string`

GetFilterValues returns the FilterValues field if non-nil, zero value otherwise.

### GetFilterValuesOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetFilterValuesOk() (*map[string][]string, bool)`

GetFilterValuesOk returns a tuple with the FilterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValues

`func (o *DtoSubscriptionUsageByMetersResponse) SetFilterValues(v map[string][]string)`

SetFilterValues sets FilterValues field to given value.

### HasFilterValues

`func (o *DtoSubscriptionUsageByMetersResponse) HasFilterValues() bool`

HasFilterValues returns a boolean if a field has been set.

### GetIsOverage

`func (o *DtoSubscriptionUsageByMetersResponse) GetIsOverage() bool`

GetIsOverage returns the IsOverage field if non-nil, zero value otherwise.

### GetIsOverageOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetIsOverageOk() (*bool, bool)`

GetIsOverageOk returns a tuple with the IsOverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverage

`func (o *DtoSubscriptionUsageByMetersResponse) SetIsOverage(v bool)`

SetIsOverage sets IsOverage field to given value.

### HasIsOverage

`func (o *DtoSubscriptionUsageByMetersResponse) HasIsOverage() bool`

HasIsOverage returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *DtoSubscriptionUsageByMetersResponse) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *DtoSubscriptionUsageByMetersResponse) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *DtoSubscriptionUsageByMetersResponse) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoSubscriptionUsageByMetersResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoSubscriptionUsageByMetersResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoSubscriptionUsageByMetersResponse) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoSubscriptionUsageByMetersResponse) GetOverageFactor() float32`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetOverageFactorOk() (*float32, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoSubscriptionUsageByMetersResponse) SetOverageFactor(v float32)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoSubscriptionUsageByMetersResponse) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

### GetPrice

`func (o *DtoSubscriptionUsageByMetersResponse) GetPrice() PricePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetPriceOk() (*PricePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DtoSubscriptionUsageByMetersResponse) SetPrice(v PricePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DtoSubscriptionUsageByMetersResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoSubscriptionUsageByMetersResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoSubscriptionUsageByMetersResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoSubscriptionUsageByMetersResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoSubscriptionUsageByMetersResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


