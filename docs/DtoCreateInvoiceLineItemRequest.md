# DtoCreateInvoiceLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | amount is the monetary amount for this line item | 
**DisplayName** | Pointer to **string** | display_name is the optional human-readable name for this line item | [optional] 
**EntityId** | Pointer to **string** | entity_id is the optional unique identifier of the entity associated with this line item | [optional] 
**EntityType** | Pointer to **string** | entity_type is the optional type of the entity associated with this line item | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterDisplayName** | Pointer to **string** | meter_display_name is the optional human-readable name of the meter | [optional] 
**MeterId** | Pointer to **string** | meter_id is the optional unique identifier of the meter used for usage tracking | [optional] 
**PeriodEnd** | Pointer to **string** | period_end is the optional end date of the period this line item covers | [optional] 
**PeriodStart** | Pointer to **string** | period_start is the optional start date of the period this line item covers | [optional] 
**PlanDisplayName** | Pointer to **string** | plan_display_name is the optional human-readable name of the plan | [optional] 
**PlanId** | Pointer to **string** | TODO: !REMOVE after migration plan_id is the optional unique identifier of the plan associated with this line item | [optional] 
**PriceId** | Pointer to **string** | price_id is the optional unique identifier of the price associated with this line item | [optional] 
**PriceType** | Pointer to **string** | price_type indicates the type of pricing (fixed, usage, tiered, etc.) | [optional] 
**PriceUnit** | Pointer to **string** | price_unit is the optional 3-digit ISO code of the price unit associated with this line item | [optional] 
**PriceUnitAmount** | Pointer to **string** | price_unit_amount is the optional amount converted to the price unit currency | [optional] 
**Quantity** | **string** | quantity is the quantity of units for this line item | 

## Methods

### NewDtoCreateInvoiceLineItemRequest

`func NewDtoCreateInvoiceLineItemRequest(amount string, quantity string, ) *DtoCreateInvoiceLineItemRequest`

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

`func (o *DtoCreateInvoiceLineItemRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreateInvoiceLineItemRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreateInvoiceLineItemRequest) SetAmount(v string)`

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

### GetEntityId

`func (o *DtoCreateInvoiceLineItemRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoCreateInvoiceLineItemRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoCreateInvoiceLineItemRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoCreateInvoiceLineItemRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoCreateInvoiceLineItemRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoCreateInvoiceLineItemRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoCreateInvoiceLineItemRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoCreateInvoiceLineItemRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

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

### GetPriceUnit

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *DtoCreateInvoiceLineItemRequest) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *DtoCreateInvoiceLineItemRequest) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPriceUnitAmount

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceUnitAmount() string`

GetPriceUnitAmount returns the PriceUnitAmount field if non-nil, zero value otherwise.

### GetPriceUnitAmountOk

`func (o *DtoCreateInvoiceLineItemRequest) GetPriceUnitAmountOk() (*string, bool)`

GetPriceUnitAmountOk returns a tuple with the PriceUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitAmount

`func (o *DtoCreateInvoiceLineItemRequest) SetPriceUnitAmount(v string)`

SetPriceUnitAmount sets PriceUnitAmount field to given value.

### HasPriceUnitAmount

`func (o *DtoCreateInvoiceLineItemRequest) HasPriceUnitAmount() bool`

HasPriceUnitAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoCreateInvoiceLineItemRequest) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoCreateInvoiceLineItemRequest) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoCreateInvoiceLineItemRequest) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


