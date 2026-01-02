# DtoSubscriptionLineItemResponse

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
**Price** | Pointer to [**DtoPriceResponse**](DtoPriceResponse.md) |  | [optional] 
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

### NewDtoSubscriptionLineItemResponse

`func NewDtoSubscriptionLineItemResponse() *DtoSubscriptionLineItemResponse`

NewDtoSubscriptionLineItemResponse instantiates a new DtoSubscriptionLineItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionLineItemResponseWithDefaults

`func NewDtoSubscriptionLineItemResponseWithDefaults() *DtoSubscriptionLineItemResponse`

NewDtoSubscriptionLineItemResponseWithDefaults instantiates a new DtoSubscriptionLineItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingPeriod

`func (o *DtoSubscriptionLineItemResponse) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoSubscriptionLineItemResponse) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoSubscriptionLineItemResponse) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DtoSubscriptionLineItemResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetCommitmentAmount

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentAmount() float32`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentAmountOk() (*float32, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentAmount(v float32)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentOverageFactor

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentOverageFactor() float32`

GetCommitmentOverageFactor returns the CommitmentOverageFactor field if non-nil, zero value otherwise.

### GetCommitmentOverageFactorOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentOverageFactorOk() (*float32, bool)`

GetCommitmentOverageFactorOk returns a tuple with the CommitmentOverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentOverageFactor

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentOverageFactor(v float32)`

SetCommitmentOverageFactor sets CommitmentOverageFactor field to given value.

### HasCommitmentOverageFactor

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentOverageFactor() bool`

HasCommitmentOverageFactor returns a boolean if a field has been set.

### GetCommitmentQuantity

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentQuantity() float32`

GetCommitmentQuantity returns the CommitmentQuantity field if non-nil, zero value otherwise.

### GetCommitmentQuantityOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentQuantityOk() (*float32, bool)`

GetCommitmentQuantityOk returns a tuple with the CommitmentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentQuantity

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentQuantity(v float32)`

SetCommitmentQuantity sets CommitmentQuantity field to given value.

### HasCommitmentQuantity

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentQuantity() bool`

HasCommitmentQuantity returns a boolean if a field has been set.

### GetCommitmentTrueUpEnabled

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentTrueUpEnabled() bool`

GetCommitmentTrueUpEnabled returns the CommitmentTrueUpEnabled field if non-nil, zero value otherwise.

### GetCommitmentTrueUpEnabledOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentTrueUpEnabledOk() (*bool, bool)`

GetCommitmentTrueUpEnabledOk returns a tuple with the CommitmentTrueUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentTrueUpEnabled

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentTrueUpEnabled(v bool)`

SetCommitmentTrueUpEnabled sets CommitmentTrueUpEnabled field to given value.

### HasCommitmentTrueUpEnabled

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentTrueUpEnabled() bool`

HasCommitmentTrueUpEnabled returns a boolean if a field has been set.

### GetCommitmentType

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentType() TypesCommitmentType`

GetCommitmentType returns the CommitmentType field if non-nil, zero value otherwise.

### GetCommitmentTypeOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentTypeOk() (*TypesCommitmentType, bool)`

GetCommitmentTypeOk returns a tuple with the CommitmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentType

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentType(v TypesCommitmentType)`

SetCommitmentType sets CommitmentType field to given value.

### HasCommitmentType

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentType() bool`

HasCommitmentType returns a boolean if a field has been set.

### GetCommitmentWindowed

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentWindowed() bool`

GetCommitmentWindowed returns the CommitmentWindowed field if non-nil, zero value otherwise.

### GetCommitmentWindowedOk

`func (o *DtoSubscriptionLineItemResponse) GetCommitmentWindowedOk() (*bool, bool)`

GetCommitmentWindowedOk returns a tuple with the CommitmentWindowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentWindowed

`func (o *DtoSubscriptionLineItemResponse) SetCommitmentWindowed(v bool)`

SetCommitmentWindowed sets CommitmentWindowed field to given value.

### HasCommitmentWindowed

`func (o *DtoSubscriptionLineItemResponse) HasCommitmentWindowed() bool`

HasCommitmentWindowed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoSubscriptionLineItemResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionLineItemResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionLineItemResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionLineItemResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoSubscriptionLineItemResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoSubscriptionLineItemResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoSubscriptionLineItemResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoSubscriptionLineItemResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoSubscriptionLineItemResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoSubscriptionLineItemResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoSubscriptionLineItemResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoSubscriptionLineItemResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoSubscriptionLineItemResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoSubscriptionLineItemResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoSubscriptionLineItemResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoSubscriptionLineItemResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *DtoSubscriptionLineItemResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DtoSubscriptionLineItemResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DtoSubscriptionLineItemResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DtoSubscriptionLineItemResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoSubscriptionLineItemResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoSubscriptionLineItemResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoSubscriptionLineItemResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoSubscriptionLineItemResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoSubscriptionLineItemResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoSubscriptionLineItemResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoSubscriptionLineItemResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoSubscriptionLineItemResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoSubscriptionLineItemResponse) GetEntityType() TypesSubscriptionLineItemEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoSubscriptionLineItemResponse) GetEntityTypeOk() (*TypesSubscriptionLineItemEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoSubscriptionLineItemResponse) SetEntityType(v TypesSubscriptionLineItemEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoSubscriptionLineItemResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoSubscriptionLineItemResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoSubscriptionLineItemResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoSubscriptionLineItemResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoSubscriptionLineItemResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionLineItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionLineItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionLineItemResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionLineItemResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *DtoSubscriptionLineItemResponse) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *DtoSubscriptionLineItemResponse) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *DtoSubscriptionLineItemResponse) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.

### HasInvoiceCadence

`func (o *DtoSubscriptionLineItemResponse) HasInvoiceCadence() bool`

HasInvoiceCadence returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionLineItemResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionLineItemResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionLineItemResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionLineItemResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterDisplayName

`func (o *DtoSubscriptionLineItemResponse) GetMeterDisplayName() string`

GetMeterDisplayName returns the MeterDisplayName field if non-nil, zero value otherwise.

### GetMeterDisplayNameOk

`func (o *DtoSubscriptionLineItemResponse) GetMeterDisplayNameOk() (*string, bool)`

GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDisplayName

`func (o *DtoSubscriptionLineItemResponse) SetMeterDisplayName(v string)`

SetMeterDisplayName sets MeterDisplayName field to given value.

### HasMeterDisplayName

`func (o *DtoSubscriptionLineItemResponse) HasMeterDisplayName() bool`

HasMeterDisplayName returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoSubscriptionLineItemResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoSubscriptionLineItemResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoSubscriptionLineItemResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoSubscriptionLineItemResponse) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPlanDisplayName

`func (o *DtoSubscriptionLineItemResponse) GetPlanDisplayName() string`

GetPlanDisplayName returns the PlanDisplayName field if non-nil, zero value otherwise.

### GetPlanDisplayNameOk

`func (o *DtoSubscriptionLineItemResponse) GetPlanDisplayNameOk() (*string, bool)`

GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDisplayName

`func (o *DtoSubscriptionLineItemResponse) SetPlanDisplayName(v string)`

SetPlanDisplayName sets PlanDisplayName field to given value.

### HasPlanDisplayName

`func (o *DtoSubscriptionLineItemResponse) HasPlanDisplayName() bool`

HasPlanDisplayName returns a boolean if a field has been set.

### GetPrice

`func (o *DtoSubscriptionLineItemResponse) GetPrice() DtoPriceResponse`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DtoSubscriptionLineItemResponse) GetPriceOk() (*DtoPriceResponse, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DtoSubscriptionLineItemResponse) SetPrice(v DtoPriceResponse)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DtoSubscriptionLineItemResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoSubscriptionLineItemResponse) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoSubscriptionLineItemResponse) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoSubscriptionLineItemResponse) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoSubscriptionLineItemResponse) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPriceType

`func (o *DtoSubscriptionLineItemResponse) GetPriceType() TypesPriceType`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *DtoSubscriptionLineItemResponse) GetPriceTypeOk() (*TypesPriceType, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *DtoSubscriptionLineItemResponse) SetPriceType(v TypesPriceType)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *DtoSubscriptionLineItemResponse) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *DtoSubscriptionLineItemResponse) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *DtoSubscriptionLineItemResponse) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *DtoSubscriptionLineItemResponse) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *DtoSubscriptionLineItemResponse) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPriceUnitId

`func (o *DtoSubscriptionLineItemResponse) GetPriceUnitId() string`

GetPriceUnitId returns the PriceUnitId field if non-nil, zero value otherwise.

### GetPriceUnitIdOk

`func (o *DtoSubscriptionLineItemResponse) GetPriceUnitIdOk() (*string, bool)`

GetPriceUnitIdOk returns a tuple with the PriceUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitId

`func (o *DtoSubscriptionLineItemResponse) SetPriceUnitId(v string)`

SetPriceUnitId sets PriceUnitId field to given value.

### HasPriceUnitId

`func (o *DtoSubscriptionLineItemResponse) HasPriceUnitId() bool`

HasPriceUnitId returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoSubscriptionLineItemResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoSubscriptionLineItemResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoSubscriptionLineItemResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoSubscriptionLineItemResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoSubscriptionLineItemResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoSubscriptionLineItemResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoSubscriptionLineItemResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoSubscriptionLineItemResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionLineItemResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionLineItemResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionLineItemResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionLineItemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoSubscriptionLineItemResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoSubscriptionLineItemResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoSubscriptionLineItemResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoSubscriptionLineItemResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionPhaseId

`func (o *DtoSubscriptionLineItemResponse) GetSubscriptionPhaseId() string`

GetSubscriptionPhaseId returns the SubscriptionPhaseId field if non-nil, zero value otherwise.

### GetSubscriptionPhaseIdOk

`func (o *DtoSubscriptionLineItemResponse) GetSubscriptionPhaseIdOk() (*string, bool)`

GetSubscriptionPhaseIdOk returns a tuple with the SubscriptionPhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPhaseId

`func (o *DtoSubscriptionLineItemResponse) SetSubscriptionPhaseId(v string)`

SetSubscriptionPhaseId sets SubscriptionPhaseId field to given value.

### HasSubscriptionPhaseId

`func (o *DtoSubscriptionLineItemResponse) HasSubscriptionPhaseId() bool`

HasSubscriptionPhaseId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoSubscriptionLineItemResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoSubscriptionLineItemResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoSubscriptionLineItemResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoSubscriptionLineItemResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *DtoSubscriptionLineItemResponse) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *DtoSubscriptionLineItemResponse) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *DtoSubscriptionLineItemResponse) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *DtoSubscriptionLineItemResponse) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSubscriptionLineItemResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSubscriptionLineItemResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSubscriptionLineItemResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSubscriptionLineItemResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoSubscriptionLineItemResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoSubscriptionLineItemResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoSubscriptionLineItemResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoSubscriptionLineItemResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


