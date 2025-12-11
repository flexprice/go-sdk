# DtoCostAnalyticItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostByPeriod** | Pointer to [**[]DtoCostPoint**](DtoCostPoint.md) | Breakdown | [optional] 
**CostsheetId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | Metadata | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** |  | [optional] 
**Meter** | Pointer to [**MeterMeter**](MeterMeter.md) |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**MeterName** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**PricePrice**](PricePrice.md) |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **string** | Aggregated metrics | [optional] 
**TotalEvents** | Pointer to **int32** |  | [optional] 
**TotalQuantity** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCostAnalyticItem

`func NewDtoCostAnalyticItem() *DtoCostAnalyticItem`

NewDtoCostAnalyticItem instantiates a new DtoCostAnalyticItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCostAnalyticItemWithDefaults

`func NewDtoCostAnalyticItemWithDefaults() *DtoCostAnalyticItem`

NewDtoCostAnalyticItemWithDefaults instantiates a new DtoCostAnalyticItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostByPeriod

`func (o *DtoCostAnalyticItem) GetCostByPeriod() []DtoCostPoint`

GetCostByPeriod returns the CostByPeriod field if non-nil, zero value otherwise.

### GetCostByPeriodOk

`func (o *DtoCostAnalyticItem) GetCostByPeriodOk() (*[]DtoCostPoint, bool)`

GetCostByPeriodOk returns a tuple with the CostByPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostByPeriod

`func (o *DtoCostAnalyticItem) SetCostByPeriod(v []DtoCostPoint)`

SetCostByPeriod sets CostByPeriod field to given value.

### HasCostByPeriod

`func (o *DtoCostAnalyticItem) HasCostByPeriod() bool`

HasCostByPeriod returns a boolean if a field has been set.

### GetCostsheetId

`func (o *DtoCostAnalyticItem) GetCostsheetId() string`

GetCostsheetId returns the CostsheetId field if non-nil, zero value otherwise.

### GetCostsheetIdOk

`func (o *DtoCostAnalyticItem) GetCostsheetIdOk() (*string, bool)`

GetCostsheetIdOk returns a tuple with the CostsheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsheetId

`func (o *DtoCostAnalyticItem) SetCostsheetId(v string)`

SetCostsheetId sets CostsheetId field to given value.

### HasCostsheetId

`func (o *DtoCostAnalyticItem) HasCostsheetId() bool`

HasCostsheetId returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCostAnalyticItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCostAnalyticItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCostAnalyticItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCostAnalyticItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoCostAnalyticItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCostAnalyticItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCostAnalyticItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCostAnalyticItem) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *DtoCostAnalyticItem) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoCostAnalyticItem) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoCostAnalyticItem) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoCostAnalyticItem) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetMeter

`func (o *DtoCostAnalyticItem) GetMeter() MeterMeter`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *DtoCostAnalyticItem) GetMeterOk() (*MeterMeter, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *DtoCostAnalyticItem) SetMeter(v MeterMeter)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *DtoCostAnalyticItem) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoCostAnalyticItem) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoCostAnalyticItem) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoCostAnalyticItem) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoCostAnalyticItem) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetMeterName

`func (o *DtoCostAnalyticItem) GetMeterName() string`

GetMeterName returns the MeterName field if non-nil, zero value otherwise.

### GetMeterNameOk

`func (o *DtoCostAnalyticItem) GetMeterNameOk() (*string, bool)`

GetMeterNameOk returns a tuple with the MeterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterName

`func (o *DtoCostAnalyticItem) SetMeterName(v string)`

SetMeterName sets MeterName field to given value.

### HasMeterName

`func (o *DtoCostAnalyticItem) HasMeterName() bool`

HasMeterName returns a boolean if a field has been set.

### GetPrice

`func (o *DtoCostAnalyticItem) GetPrice() PricePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DtoCostAnalyticItem) GetPriceOk() (*PricePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DtoCostAnalyticItem) SetPrice(v PricePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DtoCostAnalyticItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoCostAnalyticItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoCostAnalyticItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoCostAnalyticItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoCostAnalyticItem) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetProperties

`func (o *DtoCostAnalyticItem) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DtoCostAnalyticItem) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DtoCostAnalyticItem) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DtoCostAnalyticItem) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSource

`func (o *DtoCostAnalyticItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DtoCostAnalyticItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DtoCostAnalyticItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DtoCostAnalyticItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotalCost

`func (o *DtoCostAnalyticItem) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *DtoCostAnalyticItem) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *DtoCostAnalyticItem) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *DtoCostAnalyticItem) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalEvents

`func (o *DtoCostAnalyticItem) GetTotalEvents() int32`

GetTotalEvents returns the TotalEvents field if non-nil, zero value otherwise.

### GetTotalEventsOk

`func (o *DtoCostAnalyticItem) GetTotalEventsOk() (*int32, bool)`

GetTotalEventsOk returns a tuple with the TotalEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEvents

`func (o *DtoCostAnalyticItem) SetTotalEvents(v int32)`

SetTotalEvents sets TotalEvents field to given value.

### HasTotalEvents

`func (o *DtoCostAnalyticItem) HasTotalEvents() bool`

HasTotalEvents returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *DtoCostAnalyticItem) GetTotalQuantity() string`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *DtoCostAnalyticItem) GetTotalQuantityOk() (*string, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *DtoCostAnalyticItem) SetTotalQuantity(v string)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *DtoCostAnalyticItem) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


