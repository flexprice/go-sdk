# InvoiceInvoiceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CommitmentInfo** | Pointer to [**TypesCommitmentInfo**](TypesCommitmentInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceLevelDiscount** | Pointer to **float32** | invoice_level_discount is the discount amount in invoice currency applied to all line items on the invoice. | [optional] 
**LineItemDiscount** | Pointer to **float32** | line_item_discount is the discount amount in invoice currency applied directly to this line item. | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterDisplayName** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**PlanDisplayName** | Pointer to **string** |  | [optional] 
**PrepaidCreditsApplied** | Pointer to **float32** | prepaid_credits_applied is the amount in invoice currency reduced from this line item due to prepaid credits application. | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**PriceUnitAmount** | Pointer to **float32** |  | [optional] 
**PriceUnitId** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceInvoiceLineItem

`func NewInvoiceInvoiceLineItem() *InvoiceInvoiceLineItem`

NewInvoiceInvoiceLineItem instantiates a new InvoiceInvoiceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInvoiceLineItemWithDefaults

`func NewInvoiceInvoiceLineItemWithDefaults() *InvoiceInvoiceLineItem`

NewInvoiceInvoiceLineItemWithDefaults instantiates a new InvoiceInvoiceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceInvoiceLineItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceInvoiceLineItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceInvoiceLineItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceInvoiceLineItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCommitmentInfo

`func (o *InvoiceInvoiceLineItem) GetCommitmentInfo() TypesCommitmentInfo`

GetCommitmentInfo returns the CommitmentInfo field if non-nil, zero value otherwise.

### GetCommitmentInfoOk

`func (o *InvoiceInvoiceLineItem) GetCommitmentInfoOk() (*TypesCommitmentInfo, bool)`

GetCommitmentInfoOk returns a tuple with the CommitmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentInfo

`func (o *InvoiceInvoiceLineItem) SetCommitmentInfo(v TypesCommitmentInfo)`

SetCommitmentInfo sets CommitmentInfo field to given value.

### HasCommitmentInfo

`func (o *InvoiceInvoiceLineItem) HasCommitmentInfo() bool`

HasCommitmentInfo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceInvoiceLineItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceInvoiceLineItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceInvoiceLineItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceInvoiceLineItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InvoiceInvoiceLineItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InvoiceInvoiceLineItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InvoiceInvoiceLineItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InvoiceInvoiceLineItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceInvoiceLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceInvoiceLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceInvoiceLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceInvoiceLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *InvoiceInvoiceLineItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceInvoiceLineItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceInvoiceLineItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InvoiceInvoiceLineItem) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *InvoiceInvoiceLineItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InvoiceInvoiceLineItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InvoiceInvoiceLineItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InvoiceInvoiceLineItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntityId

`func (o *InvoiceInvoiceLineItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InvoiceInvoiceLineItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InvoiceInvoiceLineItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InvoiceInvoiceLineItem) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *InvoiceInvoiceLineItem) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *InvoiceInvoiceLineItem) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *InvoiceInvoiceLineItem) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *InvoiceInvoiceLineItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *InvoiceInvoiceLineItem) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *InvoiceInvoiceLineItem) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *InvoiceInvoiceLineItem) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *InvoiceInvoiceLineItem) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *InvoiceInvoiceLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceInvoiceLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceInvoiceLineItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceInvoiceLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *InvoiceInvoiceLineItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceInvoiceLineItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceInvoiceLineItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceInvoiceLineItem) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceLevelDiscount

`func (o *InvoiceInvoiceLineItem) GetInvoiceLevelDiscount() float32`

GetInvoiceLevelDiscount returns the InvoiceLevelDiscount field if non-nil, zero value otherwise.

### GetInvoiceLevelDiscountOk

`func (o *InvoiceInvoiceLineItem) GetInvoiceLevelDiscountOk() (*float32, bool)`

GetInvoiceLevelDiscountOk returns a tuple with the InvoiceLevelDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLevelDiscount

`func (o *InvoiceInvoiceLineItem) SetInvoiceLevelDiscount(v float32)`

SetInvoiceLevelDiscount sets InvoiceLevelDiscount field to given value.

### HasInvoiceLevelDiscount

`func (o *InvoiceInvoiceLineItem) HasInvoiceLevelDiscount() bool`

HasInvoiceLevelDiscount returns a boolean if a field has been set.

### GetLineItemDiscount

`func (o *InvoiceInvoiceLineItem) GetLineItemDiscount() float32`

GetLineItemDiscount returns the LineItemDiscount field if non-nil, zero value otherwise.

### GetLineItemDiscountOk

`func (o *InvoiceInvoiceLineItem) GetLineItemDiscountOk() (*float32, bool)`

GetLineItemDiscountOk returns a tuple with the LineItemDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemDiscount

`func (o *InvoiceInvoiceLineItem) SetLineItemDiscount(v float32)`

SetLineItemDiscount sets LineItemDiscount field to given value.

### HasLineItemDiscount

`func (o *InvoiceInvoiceLineItem) HasLineItemDiscount() bool`

HasLineItemDiscount returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceInvoiceLineItem) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceInvoiceLineItem) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceInvoiceLineItem) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceInvoiceLineItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *InvoiceInvoiceLineItem) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *InvoiceInvoiceLineItem) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *InvoiceInvoiceLineItem) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *InvoiceInvoiceLineItem) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *InvoiceInvoiceLineItem) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *InvoiceInvoiceLineItem) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *InvoiceInvoiceLineItem) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *InvoiceInvoiceLineItem) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *InvoiceInvoiceLineItem) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *InvoiceInvoiceLineItem) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *InvoiceInvoiceLineItem) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *InvoiceInvoiceLineItem) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *InvoiceInvoiceLineItem) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *InvoiceInvoiceLineItem) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *InvoiceInvoiceLineItem) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *InvoiceInvoiceLineItem) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlanDisplayName

`func (o *InvoiceInvoiceLineItem) GetPlanDisplayName() string`

GetPlanDisplayName returns the PlanDisplayName field if non-nil, zero value otherwise.

### GetPlanDisplayNameOk

`func (o *InvoiceInvoiceLineItem) GetPlanDisplayNameOk() (*string, bool)`

GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDisplayName

`func (o *InvoiceInvoiceLineItem) SetPlanDisplayName(v string)`

SetPlanDisplayName sets PlanDisplayName field to given value.

### HasPlanDisplayName

`func (o *InvoiceInvoiceLineItem) HasPlanDisplayName() bool`

HasPlanDisplayName returns a boolean if a field has been set.

### GetPrepaidCreditsApplied

`func (o *InvoiceInvoiceLineItem) GetPrepaidCreditsApplied() float32`

GetPrepaidCreditsApplied returns the PrepaidCreditsApplied field if non-nil, zero value otherwise.

### GetPrepaidCreditsAppliedOk

`func (o *InvoiceInvoiceLineItem) GetPrepaidCreditsAppliedOk() (*float32, bool)`

GetPrepaidCreditsAppliedOk returns a tuple with the PrepaidCreditsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidCreditsApplied

`func (o *InvoiceInvoiceLineItem) SetPrepaidCreditsApplied(v float32)`

SetPrepaidCreditsApplied sets PrepaidCreditsApplied field to given value.

### HasPrepaidCreditsApplied

`func (o *InvoiceInvoiceLineItem) HasPrepaidCreditsApplied() bool`

HasPrepaidCreditsApplied returns a boolean if a field has been set.

### GetPriceId

`func (o *InvoiceInvoiceLineItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *InvoiceInvoiceLineItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *InvoiceInvoiceLineItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *InvoiceInvoiceLineItem) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPriceType

`func (o *InvoiceInvoiceLineItem) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *InvoiceInvoiceLineItem) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *InvoiceInvoiceLineItem) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *InvoiceInvoiceLineItem) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *InvoiceInvoiceLineItem) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *InvoiceInvoiceLineItem) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *InvoiceInvoiceLineItem) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *InvoiceInvoiceLineItem) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPriceUnitAmount

`func (o *InvoiceInvoiceLineItem) GetPriceUnitAmount() float32`

GetPriceUnitAmount returns the PriceUnitAmount field if non-nil, zero value otherwise.

### GetPriceUnitAmountOk

`func (o *InvoiceInvoiceLineItem) GetPriceUnitAmountOk() (*float32, bool)`

GetPriceUnitAmountOk returns a tuple with the PriceUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitAmount

`func (o *InvoiceInvoiceLineItem) SetPriceUnitAmount(v float32)`

SetPriceUnitAmount sets PriceUnitAmount field to given value.

### HasPriceUnitAmount

`func (o *InvoiceInvoiceLineItem) HasPriceUnitAmount() bool`

HasPriceUnitAmount returns a boolean if a field has been set.

### GetPriceUnitId

`func (o *InvoiceInvoiceLineItem) GetPriceUnitId() string`

GetPriceUnitId returns the PriceUnitId field if non-nil, zero value otherwise.

### GetPriceUnitIdOk

`func (o *InvoiceInvoiceLineItem) GetPriceUnitIdOk() (*string, bool)`

GetPriceUnitIdOk returns a tuple with the PriceUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitId

`func (o *InvoiceInvoiceLineItem) SetPriceUnitId(v string)`

SetPriceUnitId sets PriceUnitId field to given value.

### HasPriceUnitId

`func (o *InvoiceInvoiceLineItem) HasPriceUnitId() bool`

HasPriceUnitId returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceInvoiceLineItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceInvoiceLineItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceInvoiceLineItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceInvoiceLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceInvoiceLineItem) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceInvoiceLineItem) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceInvoiceLineItem) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceInvoiceLineItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *InvoiceInvoiceLineItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *InvoiceInvoiceLineItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *InvoiceInvoiceLineItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *InvoiceInvoiceLineItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *InvoiceInvoiceLineItem) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceInvoiceLineItem) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceInvoiceLineItem) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceInvoiceLineItem) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceInvoiceLineItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceInvoiceLineItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceInvoiceLineItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceInvoiceLineItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *InvoiceInvoiceLineItem) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *InvoiceInvoiceLineItem) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *InvoiceInvoiceLineItem) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *InvoiceInvoiceLineItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


