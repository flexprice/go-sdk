# DtoInvoiceLineItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterDisplayName** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PlanDisplayName** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoInvoiceLineItemResponse

`func NewDtoInvoiceLineItemResponse() *DtoInvoiceLineItemResponse`

NewDtoInvoiceLineItemResponse instantiates a new DtoInvoiceLineItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoInvoiceLineItemResponseWithDefaults

`func NewDtoInvoiceLineItemResponseWithDefaults() *DtoInvoiceLineItemResponse`

NewDtoInvoiceLineItemResponseWithDefaults instantiates a new DtoInvoiceLineItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoInvoiceLineItemResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoInvoiceLineItemResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoInvoiceLineItemResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoInvoiceLineItemResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoInvoiceLineItemResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoInvoiceLineItemResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoInvoiceLineItemResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoInvoiceLineItemResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoInvoiceLineItemResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoInvoiceLineItemResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoInvoiceLineItemResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoInvoiceLineItemResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoInvoiceLineItemResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoInvoiceLineItemResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoInvoiceLineItemResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoInvoiceLineItemResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoInvoiceLineItemResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoInvoiceLineItemResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoInvoiceLineItemResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoInvoiceLineItemResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *DtoInvoiceLineItemResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DtoInvoiceLineItemResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DtoInvoiceLineItemResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DtoInvoiceLineItemResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *DtoInvoiceLineItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoInvoiceLineItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoInvoiceLineItemResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoInvoiceLineItemResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *DtoInvoiceLineItemResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DtoInvoiceLineItemResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DtoInvoiceLineItemResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *DtoInvoiceLineItemResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoInvoiceLineItemResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoInvoiceLineItemResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoInvoiceLineItemResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoInvoiceLineItemResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *DtoInvoiceLineItemResponse) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *DtoInvoiceLineItemResponse) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *DtoInvoiceLineItemResponse) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *DtoInvoiceLineItemResponse) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoInvoiceLineItemResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoInvoiceLineItemResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoInvoiceLineItemResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoInvoiceLineItemResponse) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DtoInvoiceLineItemResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DtoInvoiceLineItemResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DtoInvoiceLineItemResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DtoInvoiceLineItemResponse) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DtoInvoiceLineItemResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DtoInvoiceLineItemResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DtoInvoiceLineItemResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DtoInvoiceLineItemResponse) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlanDisplayName

`func (o *DtoInvoiceLineItemResponse) GetPlanDisplayName() string`

GetPlanDisplayName returns the PlanDisplayName field if non-nil, zero value otherwise.

### GetPlanDisplayNameOk

`func (o *DtoInvoiceLineItemResponse) GetPlanDisplayNameOk() (*string, bool)`

GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDisplayName

`func (o *DtoInvoiceLineItemResponse) SetPlanDisplayName(v string)`

SetPlanDisplayName sets PlanDisplayName field to given value.

### HasPlanDisplayName

`func (o *DtoInvoiceLineItemResponse) HasPlanDisplayName() bool`

HasPlanDisplayName returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoInvoiceLineItemResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoInvoiceLineItemResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoInvoiceLineItemResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoInvoiceLineItemResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoInvoiceLineItemResponse) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoInvoiceLineItemResponse) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoInvoiceLineItemResponse) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoInvoiceLineItemResponse) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPriceType

`func (o *DtoInvoiceLineItemResponse) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *DtoInvoiceLineItemResponse) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *DtoInvoiceLineItemResponse) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *DtoInvoiceLineItemResponse) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoInvoiceLineItemResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoInvoiceLineItemResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoInvoiceLineItemResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoInvoiceLineItemResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *DtoInvoiceLineItemResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoInvoiceLineItemResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoInvoiceLineItemResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoInvoiceLineItemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoInvoiceLineItemResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoInvoiceLineItemResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoInvoiceLineItemResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoInvoiceLineItemResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoInvoiceLineItemResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoInvoiceLineItemResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoInvoiceLineItemResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoInvoiceLineItemResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoInvoiceLineItemResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoInvoiceLineItemResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoInvoiceLineItemResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoInvoiceLineItemResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoInvoiceLineItemResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoInvoiceLineItemResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoInvoiceLineItemResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoInvoiceLineItemResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


