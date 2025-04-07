# DtoCreateInvoiceLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterDisplayName** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PlanDisplayName** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**Quantity** | **float32** |  | 

## Methods

### NewDtoCreateInvoiceLineItemRequest

`func NewDtoCreateInvoiceLineItemRequest(amount float32, quantity float32, ) *DtoCreateInvoiceLineItemRequest`

NewDtoCreateInvoiceLineItemRequest instantiates a new DtoCreateInvoiceLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateInvoiceLineItemRequestWithDefaults

`func NewDtoCreateInvoiceLineItemRequestWithDefaults() *DtoCreateInvoiceLineItemRequest`

NewDtoCreateInvoiceLineItemRequestWithDefaults instantiates a new DtoCreateInvoiceLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoCreateInvoiceLineItemRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreateInvoiceLineItemRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreateInvoiceLineItemRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DtoCreateInvoiceLineItemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateInvoiceLineItemRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateInvoiceLineItemRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateInvoiceLineItemRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateInvoiceLineItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *DtoCreateInvoiceLineItemRequest) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoCreateInvoiceLineItemRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoCreateInvoiceLineItemRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoCreateInvoiceLineItemRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoCreateInvoiceLineItemRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DtoCreateInvoiceLineItemRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DtoCreateInvoiceLineItemRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DtoCreateInvoiceLineItemRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DtoCreateInvoiceLineItemRequest) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DtoCreateInvoiceLineItemRequest) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DtoCreateInvoiceLineItemRequest) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlanDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) GetPlanDisplayName() string`

GetPlanDisplayName returns the PlanDisplayName field if non-nil, zero value otherwise.

### GetPlanDisplayNameOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPlanDisplayNameOk() (*string, bool)`

GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) SetPlanDisplayName(v string)`

SetPlanDisplayName sets PlanDisplayName field to given value.

### HasPlanDisplayName

`func (o *DtoCreateInvoiceLineItemRequest) HasPlanDisplayName() bool`

HasPlanDisplayName returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreateInvoiceLineItemRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreateInvoiceLineItemRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreateInvoiceLineItemRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoCreateInvoiceLineItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoCreateInvoiceLineItemRequest) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPriceType

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *DtoCreateInvoiceLineItemRequest) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *DtoCreateInvoiceLineItemRequest) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoCreateInvoiceLineItemRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoCreateInvoiceLineItemRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoCreateInvoiceLineItemRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


