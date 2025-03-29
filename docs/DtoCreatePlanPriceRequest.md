# DtoCreatePlanPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**BillingCadence** | [**TypesBillingCadence**](TypesBillingCadence.md) |  | 
**BillingModel** | [**TypesBillingModel**](TypesBillingModel.md) |  | 
**BillingPeriod** | [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | 
**BillingPeriodCount** | **int32** |  | 
**Currency** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FilterValues** | Pointer to **map[string][]string** |  | [optional] 
**InvoiceCadence** | [**TypesInvoiceCadence**](TypesInvoiceCadence.md) |  | 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**TierMode** | Pointer to [**TypesBillingTier**](TypesBillingTier.md) |  | [optional] 
**Tiers** | Pointer to [**[]DtoCreatePriceTier**](DtoCreatePriceTier.md) |  | [optional] 
**TransformQuantity** | Pointer to [**PriceTransformQuantity**](PriceTransformQuantity.md) |  | [optional] 
**TrialPeriod** | Pointer to **int32** |  | [optional] 
**Type** | [**TypesPriceType**](TypesPriceType.md) |  | 

## Methods

### NewDtoCreatePlanPriceRequest

`func NewDtoCreatePlanPriceRequest(billingCadence TypesBillingCadence, billingModel TypesBillingModel, billingPeriod TypesBillingPeriod, billingPeriodCount int32, currency string, invoiceCadence TypesInvoiceCadence, type_ TypesPriceType, ) *DtoCreatePlanPriceRequest`

NewDtoCreatePlanPriceRequest instantiates a new DtoCreatePlanPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreatePlanPriceRequestWithDefaults

`func NewDtoCreatePlanPriceRequestWithDefaults() *DtoCreatePlanPriceRequest`

NewDtoCreatePlanPriceRequestWithDefaults instantiates a new DtoCreatePlanPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoCreatePlanPriceRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreatePlanPriceRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreatePlanPriceRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoCreatePlanPriceRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoCreatePlanPriceRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoCreatePlanPriceRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoCreatePlanPriceRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingModel

`func (o *DtoCreatePlanPriceRequest) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *DtoCreatePlanPriceRequest) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *DtoCreatePlanPriceRequest) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.


### GetBillingPeriod

`func (o *DtoCreatePlanPriceRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoCreatePlanPriceRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoCreatePlanPriceRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoCreatePlanPriceRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoCreatePlanPriceRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoCreatePlanPriceRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.


### GetCurrency

`func (o *DtoCreatePlanPriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreatePlanPriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreatePlanPriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *DtoCreatePlanPriceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreatePlanPriceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreatePlanPriceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreatePlanPriceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterValues

`func (o *DtoCreatePlanPriceRequest) GetFilterValues() map[string][]string`

GetFilterValues returns the FilterValues field if non-nil, zero value otherwise.

### GetFilterValuesOk

`func (o *DtoCreatePlanPriceRequest) GetFilterValuesOk() (*map[string][]string, bool)`

GetFilterValuesOk returns a tuple with the FilterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValues

`func (o *DtoCreatePlanPriceRequest) SetFilterValues(v map[string][]string)`

SetFilterValues sets FilterValues field to given value.

### HasFilterValues

`func (o *DtoCreatePlanPriceRequest) HasFilterValues() bool`

HasFilterValues returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *DtoCreatePlanPriceRequest) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *DtoCreatePlanPriceRequest) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *DtoCreatePlanPriceRequest) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.


### GetLookupKey

`func (o *DtoCreatePlanPriceRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreatePlanPriceRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreatePlanPriceRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreatePlanPriceRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreatePlanPriceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreatePlanPriceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreatePlanPriceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreatePlanPriceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoCreatePlanPriceRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoCreatePlanPriceRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoCreatePlanPriceRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoCreatePlanPriceRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreatePlanPriceRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreatePlanPriceRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreatePlanPriceRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreatePlanPriceRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTierMode

`func (o *DtoCreatePlanPriceRequest) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *DtoCreatePlanPriceRequest) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *DtoCreatePlanPriceRequest) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *DtoCreatePlanPriceRequest) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *DtoCreatePlanPriceRequest) GetTiers() []DtoCreatePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *DtoCreatePlanPriceRequest) GetTiersOk() (*[]DtoCreatePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *DtoCreatePlanPriceRequest) SetTiers(v []DtoCreatePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *DtoCreatePlanPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *DtoCreatePlanPriceRequest) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *DtoCreatePlanPriceRequest) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *DtoCreatePlanPriceRequest) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *DtoCreatePlanPriceRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *DtoCreatePlanPriceRequest) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *DtoCreatePlanPriceRequest) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *DtoCreatePlanPriceRequest) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *DtoCreatePlanPriceRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetType

`func (o *DtoCreatePlanPriceRequest) GetType() TypesPriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreatePlanPriceRequest) GetTypeOk() (*TypesPriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreatePlanPriceRequest) SetType(v TypesPriceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


