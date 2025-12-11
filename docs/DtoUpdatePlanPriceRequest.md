# DtoUpdatePlanPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**BillingCadence** | [**TypesBillingCadence**](TypesBillingCadence.md) |  | 
**BillingModel** | [**TypesBillingModel**](TypesBillingModel.md) |  | 
**BillingPeriod** | [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | 
**BillingPeriodCount** | Pointer to **int32** |  | [optional] 
**Currency** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** | TODO: this will be required in the future as we will not allow prices to be created without an entity id | [optional] 
**EntityType** | Pointer to [**TypesPriceEntityType**](TypesPriceEntityType.md) |  | [optional] 
**FilterValues** | Pointer to **map[string][]string** |  | [optional] 
**GroupId** | Pointer to **string** | GroupID is the id of the group to add the price to | [optional] 
**Id** | Pointer to **string** | The ID of the price to update (present if the price is being updated) | [optional] 
**InvoiceCadence** | [**TypesInvoiceCadence**](TypesInvoiceCadence.md) |  | 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** | TODO: This is deprecated and will be removed in the future | [optional] 
**PriceUnitConfig** | Pointer to [**DtoPriceUnitConfig**](DtoPriceUnitConfig.md) |  | [optional] 
**PriceUnitType** | [**TypesPriceUnitType**](TypesPriceUnitType.md) |  | 
**StartDate** | Pointer to **string** |  | [optional] 
**TierMode** | Pointer to [**TypesBillingTier**](TypesBillingTier.md) |  | [optional] 
**Tiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) |  | [optional] 
**TransformQuantity** | Pointer to [**PriceTransformQuantity**](PriceTransformQuantity.md) |  | [optional] 
**TrialPeriod** | Pointer to **int32** |  | [optional] 
**Type** | [**TypesPriceType**](TypesPriceType.md) |  | 

## Methods

### NewDtoUpdatePlanPriceRequest

`func NewDtoUpdatePlanPriceRequest(billingCadence TypesBillingCadence, billingModel TypesBillingModel, billingPeriod TypesBillingPeriod, currency string, invoiceCadence TypesInvoiceCadence, priceUnitType TypesPriceUnitType, type_ TypesPriceType, ) *DtoUpdatePlanPriceRequest`

NewDtoUpdatePlanPriceRequest instantiates a new DtoUpdatePlanPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePlanPriceRequestWithDefaults

`func NewDtoUpdatePlanPriceRequestWithDefaults() *DtoUpdatePlanPriceRequest`

NewDtoUpdatePlanPriceRequestWithDefaults instantiates a new DtoUpdatePlanPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoUpdatePlanPriceRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoUpdatePlanPriceRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoUpdatePlanPriceRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoUpdatePlanPriceRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoUpdatePlanPriceRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoUpdatePlanPriceRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoUpdatePlanPriceRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingModel

`func (o *DtoUpdatePlanPriceRequest) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *DtoUpdatePlanPriceRequest) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *DtoUpdatePlanPriceRequest) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.


### GetBillingPeriod

`func (o *DtoUpdatePlanPriceRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoUpdatePlanPriceRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoUpdatePlanPriceRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoUpdatePlanPriceRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoUpdatePlanPriceRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoUpdatePlanPriceRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *DtoUpdatePlanPriceRequest) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoUpdatePlanPriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoUpdatePlanPriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoUpdatePlanPriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *DtoUpdatePlanPriceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoUpdatePlanPriceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoUpdatePlanPriceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoUpdatePlanPriceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoUpdatePlanPriceRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoUpdatePlanPriceRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoUpdatePlanPriceRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoUpdatePlanPriceRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoUpdatePlanPriceRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoUpdatePlanPriceRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoUpdatePlanPriceRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoUpdatePlanPriceRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoUpdatePlanPriceRequest) GetEntityType() TypesPriceEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoUpdatePlanPriceRequest) GetEntityTypeOk() (*TypesPriceEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoUpdatePlanPriceRequest) SetEntityType(v TypesPriceEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoUpdatePlanPriceRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFilterValues

`func (o *DtoUpdatePlanPriceRequest) GetFilterValues() map[string][]string`

GetFilterValues returns the FilterValues field if non-nil, zero value otherwise.

### GetFilterValuesOk

`func (o *DtoUpdatePlanPriceRequest) GetFilterValuesOk() (*map[string][]string, bool)`

GetFilterValuesOk returns a tuple with the FilterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValues

`func (o *DtoUpdatePlanPriceRequest) SetFilterValues(v map[string][]string)`

SetFilterValues sets FilterValues field to given value.

### HasFilterValues

`func (o *DtoUpdatePlanPriceRequest) HasFilterValues() bool`

HasFilterValues returns a boolean if a field has been set.

### GetGroupId

`func (o *DtoUpdatePlanPriceRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DtoUpdatePlanPriceRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DtoUpdatePlanPriceRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DtoUpdatePlanPriceRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *DtoUpdatePlanPriceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoUpdatePlanPriceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoUpdatePlanPriceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoUpdatePlanPriceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *DtoUpdatePlanPriceRequest) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *DtoUpdatePlanPriceRequest) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *DtoUpdatePlanPriceRequest) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.


### GetLookupKey

`func (o *DtoUpdatePlanPriceRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoUpdatePlanPriceRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoUpdatePlanPriceRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoUpdatePlanPriceRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdatePlanPriceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdatePlanPriceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdatePlanPriceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdatePlanPriceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoUpdatePlanPriceRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoUpdatePlanPriceRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoUpdatePlanPriceRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoUpdatePlanPriceRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoUpdatePlanPriceRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoUpdatePlanPriceRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoUpdatePlanPriceRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoUpdatePlanPriceRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriceUnitConfig

`func (o *DtoUpdatePlanPriceRequest) GetPriceUnitConfig() DtoPriceUnitConfig`

GetPriceUnitConfig returns the PriceUnitConfig field if non-nil, zero value otherwise.

### GetPriceUnitConfigOk

`func (o *DtoUpdatePlanPriceRequest) GetPriceUnitConfigOk() (*DtoPriceUnitConfig, bool)`

GetPriceUnitConfigOk returns a tuple with the PriceUnitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitConfig

`func (o *DtoUpdatePlanPriceRequest) SetPriceUnitConfig(v DtoPriceUnitConfig)`

SetPriceUnitConfig sets PriceUnitConfig field to given value.

### HasPriceUnitConfig

`func (o *DtoUpdatePlanPriceRequest) HasPriceUnitConfig() bool`

HasPriceUnitConfig returns a boolean if a field has been set.

### GetPriceUnitType

`func (o *DtoUpdatePlanPriceRequest) GetPriceUnitType() TypesPriceUnitType`

GetPriceUnitType returns the PriceUnitType field if non-nil, zero value otherwise.

### GetPriceUnitTypeOk

`func (o *DtoUpdatePlanPriceRequest) GetPriceUnitTypeOk() (*TypesPriceUnitType, bool)`

GetPriceUnitTypeOk returns a tuple with the PriceUnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitType

`func (o *DtoUpdatePlanPriceRequest) SetPriceUnitType(v TypesPriceUnitType)`

SetPriceUnitType sets PriceUnitType field to given value.


### GetStartDate

`func (o *DtoUpdatePlanPriceRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoUpdatePlanPriceRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoUpdatePlanPriceRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoUpdatePlanPriceRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTierMode

`func (o *DtoUpdatePlanPriceRequest) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *DtoUpdatePlanPriceRequest) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *DtoUpdatePlanPriceRequest) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *DtoUpdatePlanPriceRequest) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *DtoUpdatePlanPriceRequest) GetTiers() []DtoCreatePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *DtoUpdatePlanPriceRequest) GetTiersOk() (*[]DtoCreatePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *DtoUpdatePlanPriceRequest) SetTiers(v []DtoCreatePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *DtoUpdatePlanPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *DtoUpdatePlanPriceRequest) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *DtoUpdatePlanPriceRequest) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *DtoUpdatePlanPriceRequest) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *DtoUpdatePlanPriceRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *DtoUpdatePlanPriceRequest) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *DtoUpdatePlanPriceRequest) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *DtoUpdatePlanPriceRequest) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *DtoUpdatePlanPriceRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetType

`func (o *DtoUpdatePlanPriceRequest) GetType() TypesPriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoUpdatePlanPriceRequest) GetTypeOk() (*TypesPriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoUpdatePlanPriceRequest) SetType(v TypesPriceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


