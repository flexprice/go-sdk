# SubscriptionSubscriptionLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 
**CommitmentAmount** | Pointer to **float32** | Commitment fields | [optional] 
**CommitmentOverageFactor** | Pointer to **float32** |  | [optional] 
**CommitmentQuantity** | Pointer to **float32** |  | [optional] 
**CommitmentTrueUpEnabled** | Pointer to **bool** |  | [optional] 
**CommitmentType** | Pointer to [**TypesCommitmentType**](TypesCommitmentType.md) |  | [optional] 
**CommitmentWindowed** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesSubscriptionLineItemEntityType**](TypesSubscriptionLineItemEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceCadence** | Pointer to [**TypesInvoiceCadence**](TypesInvoiceCadence.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterDisplayName** | Pointer to **string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PlanDisplayName** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**PricePrice**](PricePrice.md) |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PriceType** | Pointer to [**TypesPriceType**](TypesPriceType.md) |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**PriceUnitId** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SubscriptionPhaseId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TrialPeriod** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionSubscriptionLineItem

`func NewSubscriptionSubscriptionLineItem() *SubscriptionSubscriptionLineItem`

NewSubscriptionSubscriptionLineItem instantiates a new SubscriptionSubscriptionLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionSubscriptionLineItemWithDefaults

`func NewSubscriptionSubscriptionLineItemWithDefaults() *SubscriptionSubscriptionLineItem`

NewSubscriptionSubscriptionLineItemWithDefaults instantiates a new SubscriptionSubscriptionLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingPeriod

`func (o *SubscriptionSubscriptionLineItem) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *SubscriptionSubscriptionLineItem) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *SubscriptionSubscriptionLineItem) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *SubscriptionSubscriptionLineItem) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetCommitmentAmount

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentAmount() float32`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentAmountOk() (*float32, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentAmount(v float32)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentOverageFactor

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentOverageFactor() float32`

GetCommitmentOverageFactor returns the CommitmentOverageFactor field if non-nil, zero value otherwise.

### GetCommitmentOverageFactorOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentOverageFactorOk() (*float32, bool)`

GetCommitmentOverageFactorOk returns a tuple with the CommitmentOverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentOverageFactor

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentOverageFactor(v float32)`

SetCommitmentOverageFactor sets CommitmentOverageFactor field to given value.

### HasCommitmentOverageFactor

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentOverageFactor() bool`

HasCommitmentOverageFactor returns a boolean if a field has been set.

### GetCommitmentQuantity

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentQuantity() float32`

GetCommitmentQuantity returns the CommitmentQuantity field if non-nil, zero value otherwise.

### GetCommitmentQuantityOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentQuantityOk() (*float32, bool)`

GetCommitmentQuantityOk returns a tuple with the CommitmentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentQuantity

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentQuantity(v float32)`

SetCommitmentQuantity sets CommitmentQuantity field to given value.

### HasCommitmentQuantity

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentQuantity() bool`

HasCommitmentQuantity returns a boolean if a field has been set.

### GetCommitmentTrueUpEnabled

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentTrueUpEnabled() bool`

GetCommitmentTrueUpEnabled returns the CommitmentTrueUpEnabled field if non-nil, zero value otherwise.

### GetCommitmentTrueUpEnabledOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentTrueUpEnabledOk() (*bool, bool)`

GetCommitmentTrueUpEnabledOk returns a tuple with the CommitmentTrueUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentTrueUpEnabled

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentTrueUpEnabled(v bool)`

SetCommitmentTrueUpEnabled sets CommitmentTrueUpEnabled field to given value.

### HasCommitmentTrueUpEnabled

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentTrueUpEnabled() bool`

HasCommitmentTrueUpEnabled returns a boolean if a field has been set.

### GetCommitmentType

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentType() TypesCommitmentType`

GetCommitmentType returns the CommitmentType field if non-nil, zero value otherwise.

### GetCommitmentTypeOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentTypeOk() (*TypesCommitmentType, bool)`

GetCommitmentTypeOk returns a tuple with the CommitmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentType

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentType(v TypesCommitmentType)`

SetCommitmentType sets CommitmentType field to given value.

### HasCommitmentType

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentType() bool`

HasCommitmentType returns a boolean if a field has been set.

### GetCommitmentWindowed

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentWindowed() bool`

GetCommitmentWindowed returns the CommitmentWindowed field if non-nil, zero value otherwise.

### GetCommitmentWindowedOk

`func (o *SubscriptionSubscriptionLineItem) GetCommitmentWindowedOk() (*bool, bool)`

GetCommitmentWindowedOk returns a tuple with the CommitmentWindowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentWindowed

`func (o *SubscriptionSubscriptionLineItem) SetCommitmentWindowed(v bool)`

SetCommitmentWindowed sets CommitmentWindowed field to given value.

### HasCommitmentWindowed

`func (o *SubscriptionSubscriptionLineItem) HasCommitmentWindowed() bool`

HasCommitmentWindowed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionSubscriptionLineItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionSubscriptionLineItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionSubscriptionLineItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionSubscriptionLineItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SubscriptionSubscriptionLineItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubscriptionSubscriptionLineItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubscriptionSubscriptionLineItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SubscriptionSubscriptionLineItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionSubscriptionLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionSubscriptionLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionSubscriptionLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionSubscriptionLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionSubscriptionLineItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionSubscriptionLineItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionSubscriptionLineItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionSubscriptionLineItem) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionSubscriptionLineItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionSubscriptionLineItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionSubscriptionLineItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionSubscriptionLineItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndDate

`func (o *SubscriptionSubscriptionLineItem) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionSubscriptionLineItem) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionSubscriptionLineItem) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SubscriptionSubscriptionLineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntityId

`func (o *SubscriptionSubscriptionLineItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SubscriptionSubscriptionLineItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SubscriptionSubscriptionLineItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SubscriptionSubscriptionLineItem) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *SubscriptionSubscriptionLineItem) GetEntityType() TypesSubscriptionLineItemEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SubscriptionSubscriptionLineItem) GetEntityTypeOk() (*TypesSubscriptionLineItemEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SubscriptionSubscriptionLineItem) SetEntityType(v TypesSubscriptionLineItemEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *SubscriptionSubscriptionLineItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SubscriptionSubscriptionLineItem) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SubscriptionSubscriptionLineItem) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SubscriptionSubscriptionLineItem) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SubscriptionSubscriptionLineItem) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionSubscriptionLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionSubscriptionLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionSubscriptionLineItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionSubscriptionLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *SubscriptionSubscriptionLineItem) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *SubscriptionSubscriptionLineItem) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *SubscriptionSubscriptionLineItem) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.

### HasInvoiceCadence

`func (o *SubscriptionSubscriptionLineItem) HasInvoiceCadence() bool`

HasInvoiceCadence returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionSubscriptionLineItem) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionSubscriptionLineItem) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionSubscriptionLineItem) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionSubscriptionLineItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *SubscriptionSubscriptionLineItem) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *SubscriptionSubscriptionLineItem) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *SubscriptionSubscriptionLineItem) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *SubscriptionSubscriptionLineItem) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *SubscriptionSubscriptionLineItem) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *SubscriptionSubscriptionLineItem) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *SubscriptionSubscriptionLineItem) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *SubscriptionSubscriptionLineItem) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPlanDisplayName

`func (o *SubscriptionSubscriptionLineItem) GetPlanDisplayName() string`

GetPlanDisplayName returns the PlanDisplayName field if non-nil, zero value otherwise.

### GetPlanDisplayNameOk

`func (o *SubscriptionSubscriptionLineItem) GetPlanDisplayNameOk() (*string, bool)`

GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDisplayName

`func (o *SubscriptionSubscriptionLineItem) SetPlanDisplayName(v string)`

SetPlanDisplayName sets PlanDisplayName field to given value.

### HasPlanDisplayName

`func (o *SubscriptionSubscriptionLineItem) HasPlanDisplayName() bool`

HasPlanDisplayName returns a boolean if a field has been set.

### GetPrice

`func (o *SubscriptionSubscriptionLineItem) GetPrice() PricePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionSubscriptionLineItem) GetPriceOk() (*PricePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionSubscriptionLineItem) SetPrice(v PricePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SubscriptionSubscriptionLineItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *SubscriptionSubscriptionLineItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionSubscriptionLineItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionSubscriptionLineItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *SubscriptionSubscriptionLineItem) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPriceType

`func (o *SubscriptionSubscriptionLineItem) GetPriceType() TypesPriceType`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *SubscriptionSubscriptionLineItem) GetPriceTypeOk() (*TypesPriceType, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *SubscriptionSubscriptionLineItem) SetPriceType(v TypesPriceType)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *SubscriptionSubscriptionLineItem) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *SubscriptionSubscriptionLineItem) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *SubscriptionSubscriptionLineItem) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *SubscriptionSubscriptionLineItem) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *SubscriptionSubscriptionLineItem) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPriceUnitId

`func (o *SubscriptionSubscriptionLineItem) GetPriceUnitId() string`

GetPriceUnitId returns the PriceUnitId field if non-nil, zero value otherwise.

### GetPriceUnitIdOk

`func (o *SubscriptionSubscriptionLineItem) GetPriceUnitIdOk() (*string, bool)`

GetPriceUnitIdOk returns a tuple with the PriceUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitId

`func (o *SubscriptionSubscriptionLineItem) SetPriceUnitId(v string)`

SetPriceUnitId sets PriceUnitId field to given value.

### HasPriceUnitId

`func (o *SubscriptionSubscriptionLineItem) HasPriceUnitId() bool`

HasPriceUnitId returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionSubscriptionLineItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionSubscriptionLineItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionSubscriptionLineItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionSubscriptionLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *SubscriptionSubscriptionLineItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionSubscriptionLineItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionSubscriptionLineItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionSubscriptionLineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionSubscriptionLineItem) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionSubscriptionLineItem) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionSubscriptionLineItem) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionSubscriptionLineItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionSubscriptionLineItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionSubscriptionLineItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionSubscriptionLineItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionSubscriptionLineItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionPhaseId

`func (o *SubscriptionSubscriptionLineItem) GetSubscriptionPhaseId() string`

GetSubscriptionPhaseId returns the SubscriptionPhaseId field if non-nil, zero value otherwise.

### GetSubscriptionPhaseIdOk

`func (o *SubscriptionSubscriptionLineItem) GetSubscriptionPhaseIdOk() (*string, bool)`

GetSubscriptionPhaseIdOk returns a tuple with the SubscriptionPhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPhaseId

`func (o *SubscriptionSubscriptionLineItem) SetSubscriptionPhaseId(v string)`

SetSubscriptionPhaseId sets SubscriptionPhaseId field to given value.

### HasSubscriptionPhaseId

`func (o *SubscriptionSubscriptionLineItem) HasSubscriptionPhaseId() bool`

HasSubscriptionPhaseId returns a boolean if a field has been set.

### GetTenantId

`func (o *SubscriptionSubscriptionLineItem) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SubscriptionSubscriptionLineItem) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SubscriptionSubscriptionLineItem) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SubscriptionSubscriptionLineItem) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *SubscriptionSubscriptionLineItem) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *SubscriptionSubscriptionLineItem) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *SubscriptionSubscriptionLineItem) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *SubscriptionSubscriptionLineItem) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionSubscriptionLineItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionSubscriptionLineItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionSubscriptionLineItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionSubscriptionLineItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SubscriptionSubscriptionLineItem) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SubscriptionSubscriptionLineItem) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SubscriptionSubscriptionLineItem) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SubscriptionSubscriptionLineItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


