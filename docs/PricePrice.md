# PricePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount stored in main currency units (e.g., dollars, not cents) For USD: 12.50 means $12.50 | [optional] 
**BillingCadence** | Pointer to [**TypesBillingCadence**](TypesBillingCadence.md) |  | [optional] 
**BillingModel** | Pointer to [**TypesBillingModel**](TypesBillingModel.md) |  | [optional] 
**BillingPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 
**BillingPeriodCount** | Pointer to **int32** | BillingPeriodCount is the count of the billing period ex 1, 3, 6, 12 | [optional] 
**ConversionRate** | Pointer to **float32** | ConversionRate is the rate of the price unit to the base currency For BTC: 1 BTC &#x3D; 100000000 USD | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | Currency 3 digit ISO currency code in lowercase ex usd, eur, gbp | [optional] 
**Description** | Pointer to **string** | Description of the price | [optional] 
**DisplayAmount** | Pointer to **string** | DisplayAmount is the formatted amount with currency symbol For USD: $12.50 | [optional] 
**DisplayPriceUnitAmount** | Pointer to **string** | DisplayPriceUnitAmount is the formatted amount with price unit symbol For BTC: 0.00000001 BTC | [optional] 
**EndDate** | Pointer to **string** | EndDate is the end date of the price | [optional] 
**EntityId** | Pointer to **string** | EntityID holds the value of the \&quot;entity_id\&quot; field. | [optional] 
**EntityType** | Pointer to [**TypesPriceEntityType**](TypesPriceEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** | EnvironmentID is the environment identifier for the price | [optional] 
**GroupId** | Pointer to **string** | GroupID references the group this price belongs to | [optional] 
**Id** | Pointer to **string** | ID uuid identifier for the price | [optional] 
**InvoiceCadence** | Pointer to [**TypesInvoiceCadence**](TypesInvoiceCadence.md) |  | [optional] 
**LookupKey** | Pointer to **string** | LookupKey used for looking up the price in the database | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterId** | Pointer to **string** | MeterID is the id of the meter for usage based pricing | [optional] 
**ParentPriceId** | Pointer to **string** | ParentPriceID references the root price (always set for price lineage tracking) | [optional] 
**PriceUnit** | Pointer to **string** | PriceUnit 3 digit ISO currency code in lowercase ex btc For BTC: btc | [optional] 
**PriceUnitAmount** | Pointer to **float32** | PriceUnitAmount is the amount stored in price unit For BTC: 0.00000001 means 0.00000001 BTC | [optional] 
**PriceUnitId** | Pointer to **string** | PriceUnitID is the id of the price unit | [optional] 
**PriceUnitTiers** | Pointer to [**[]PricePriceTier**](PricePriceTier.md) | PriceUnitTiers are the tiers for the price unit | [optional] 
**PriceUnitType** | Pointer to [**TypesPriceUnitType**](TypesPriceUnitType.md) |  | [optional] 
**StartDate** | Pointer to **string** | StartDate is the start date of the price | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TierMode** | Pointer to [**TypesBillingTier**](TypesBillingTier.md) |  | [optional] 
**Tiers** | Pointer to [**[]PricePriceTier**](PricePriceTier.md) |  | [optional] 
**TransformQuantity** | Pointer to [**PriceJSONBTransformQuantity**](PriceJSONBTransformQuantity.md) |  | [optional] 
**TrialPeriod** | Pointer to **int32** | TrialPeriod is the number of days for the trial period Note: This is only applicable for recurring prices (BILLING_CADENCE_RECURRING) | [optional] 
**Type** | Pointer to [**TypesPriceType**](TypesPriceType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewPricePrice

`func NewPricePrice() *PricePrice`

NewPricePrice instantiates a new PricePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricePriceWithDefaults

`func NewPricePriceWithDefaults() *PricePrice`

NewPricePriceWithDefaults instantiates a new PricePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PricePrice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PricePrice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PricePrice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PricePrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCadence

`func (o *PricePrice) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *PricePrice) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *PricePrice) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *PricePrice) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingModel

`func (o *PricePrice) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *PricePrice) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *PricePrice) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *PricePrice) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *PricePrice) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *PricePrice) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *PricePrice) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *PricePrice) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBillingPeriodCount

`func (o *PricePrice) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *PricePrice) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *PricePrice) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *PricePrice) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetConversionRate

`func (o *PricePrice) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *PricePrice) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *PricePrice) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *PricePrice) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PricePrice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PricePrice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PricePrice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PricePrice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PricePrice) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PricePrice) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PricePrice) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PricePrice) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *PricePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PricePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PricePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PricePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *PricePrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricePrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricePrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricePrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAmount

`func (o *PricePrice) GetDisplayAmount() string`

GetDisplayAmount returns the DisplayAmount field if non-nil, zero value otherwise.

### GetDisplayAmountOk

`func (o *PricePrice) GetDisplayAmountOk() (*string, bool)`

GetDisplayAmountOk returns a tuple with the DisplayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAmount

`func (o *PricePrice) SetDisplayAmount(v string)`

SetDisplayAmount sets DisplayAmount field to given value.

### HasDisplayAmount

`func (o *PricePrice) HasDisplayAmount() bool`

HasDisplayAmount returns a boolean if a field has been set.

### GetDisplayPriceUnitAmount

`func (o *PricePrice) GetDisplayPriceUnitAmount() string`

GetDisplayPriceUnitAmount returns the DisplayPriceUnitAmount field if non-nil, zero value otherwise.

### GetDisplayPriceUnitAmountOk

`func (o *PricePrice) GetDisplayPriceUnitAmountOk() (*string, bool)`

GetDisplayPriceUnitAmountOk returns a tuple with the DisplayPriceUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPriceUnitAmount

`func (o *PricePrice) SetDisplayPriceUnitAmount(v string)`

SetDisplayPriceUnitAmount sets DisplayPriceUnitAmount field to given value.

### HasDisplayPriceUnitAmount

`func (o *PricePrice) HasDisplayPriceUnitAmount() bool`

HasDisplayPriceUnitAmount returns a boolean if a field has been set.

### GetEndDate

`func (o *PricePrice) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PricePrice) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PricePrice) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PricePrice) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntityId

`func (o *PricePrice) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *PricePrice) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *PricePrice) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *PricePrice) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *PricePrice) GetEntityType() TypesPriceEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PricePrice) GetEntityTypeOk() (*TypesPriceEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PricePrice) SetEntityType(v TypesPriceEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PricePrice) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *PricePrice) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PricePrice) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PricePrice) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *PricePrice) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetGroupId

`func (o *PricePrice) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PricePrice) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PricePrice) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PricePrice) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *PricePrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricePrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricePrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricePrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *PricePrice) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *PricePrice) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *PricePrice) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.

### HasInvoiceCadence

`func (o *PricePrice) HasInvoiceCadence() bool`

HasInvoiceCadence returns a boolean if a field has been set.

### GetLookupKey

`func (o *PricePrice) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *PricePrice) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *PricePrice) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *PricePrice) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *PricePrice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PricePrice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PricePrice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PricePrice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterId

`func (o *PricePrice) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *PricePrice) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *PricePrice) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *PricePrice) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetParentPriceId

`func (o *PricePrice) GetParentPriceId() string`

GetParentPriceId returns the ParentPriceId field if non-nil, zero value otherwise.

### GetParentPriceIdOk

`func (o *PricePrice) GetParentPriceIdOk() (*string, bool)`

GetParentPriceIdOk returns a tuple with the ParentPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPriceId

`func (o *PricePrice) SetParentPriceId(v string)`

SetParentPriceId sets ParentPriceId field to given value.

### HasParentPriceId

`func (o *PricePrice) HasParentPriceId() bool`

HasParentPriceId returns a boolean if a field has been set.

### GetPriceUnit

`func (o *PricePrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricePrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricePrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricePrice) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPriceUnitAmount

`func (o *PricePrice) GetPriceUnitAmount() float32`

GetPriceUnitAmount returns the PriceUnitAmount field if non-nil, zero value otherwise.

### GetPriceUnitAmountOk

`func (o *PricePrice) GetPriceUnitAmountOk() (*float32, bool)`

GetPriceUnitAmountOk returns a tuple with the PriceUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitAmount

`func (o *PricePrice) SetPriceUnitAmount(v float32)`

SetPriceUnitAmount sets PriceUnitAmount field to given value.

### HasPriceUnitAmount

`func (o *PricePrice) HasPriceUnitAmount() bool`

HasPriceUnitAmount returns a boolean if a field has been set.

### GetPriceUnitId

`func (o *PricePrice) GetPriceUnitId() string`

GetPriceUnitId returns the PriceUnitId field if non-nil, zero value otherwise.

### GetPriceUnitIdOk

`func (o *PricePrice) GetPriceUnitIdOk() (*string, bool)`

GetPriceUnitIdOk returns a tuple with the PriceUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitId

`func (o *PricePrice) SetPriceUnitId(v string)`

SetPriceUnitId sets PriceUnitId field to given value.

### HasPriceUnitId

`func (o *PricePrice) HasPriceUnitId() bool`

HasPriceUnitId returns a boolean if a field has been set.

### GetPriceUnitTiers

`func (o *PricePrice) GetPriceUnitTiers() []PricePriceTier`

GetPriceUnitTiers returns the PriceUnitTiers field if non-nil, zero value otherwise.

### GetPriceUnitTiersOk

`func (o *PricePrice) GetPriceUnitTiersOk() (*[]PricePriceTier, bool)`

GetPriceUnitTiersOk returns a tuple with the PriceUnitTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitTiers

`func (o *PricePrice) SetPriceUnitTiers(v []PricePriceTier)`

SetPriceUnitTiers sets PriceUnitTiers field to given value.

### HasPriceUnitTiers

`func (o *PricePrice) HasPriceUnitTiers() bool`

HasPriceUnitTiers returns a boolean if a field has been set.

### GetPriceUnitType

`func (o *PricePrice) GetPriceUnitType() TypesPriceUnitType`

GetPriceUnitType returns the PriceUnitType field if non-nil, zero value otherwise.

### GetPriceUnitTypeOk

`func (o *PricePrice) GetPriceUnitTypeOk() (*TypesPriceUnitType, bool)`

GetPriceUnitTypeOk returns a tuple with the PriceUnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitType

`func (o *PricePrice) SetPriceUnitType(v TypesPriceUnitType)`

SetPriceUnitType sets PriceUnitType field to given value.

### HasPriceUnitType

`func (o *PricePrice) HasPriceUnitType() bool`

HasPriceUnitType returns a boolean if a field has been set.

### GetStartDate

`func (o *PricePrice) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PricePrice) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PricePrice) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PricePrice) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *PricePrice) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PricePrice) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PricePrice) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PricePrice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *PricePrice) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PricePrice) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PricePrice) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PricePrice) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTierMode

`func (o *PricePrice) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *PricePrice) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *PricePrice) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *PricePrice) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *PricePrice) GetTiers() []PricePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PricePrice) GetTiersOk() (*[]PricePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PricePrice) SetTiers(v []PricePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PricePrice) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *PricePrice) GetTransformQuantity() PriceJSONBTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *PricePrice) GetTransformQuantityOk() (*PriceJSONBTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *PricePrice) SetTransformQuantity(v PriceJSONBTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *PricePrice) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *PricePrice) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *PricePrice) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *PricePrice) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *PricePrice) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetType

`func (o *PricePrice) GetType() TypesPriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PricePrice) GetTypeOk() (*TypesPriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PricePrice) SetType(v TypesPriceType)`

SetType sets Type field to given value.

### HasType

`func (o *PricePrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PricePrice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PricePrice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PricePrice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PricePrice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *PricePrice) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *PricePrice) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *PricePrice) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *PricePrice) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


