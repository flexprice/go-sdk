# DtoCreatePriceRequest

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

### NewDtoCreatePriceRequest

`func NewDtoCreatePriceRequest(billingCadence TypesBillingCadence, billingModel TypesBillingModel, billingPeriod TypesBillingPeriod, billingPeriodCount int32, currency string, invoiceCadence TypesInvoiceCadence, type_ TypesPriceType, ) *DtoCreatePriceRequest`

NewDtoCreatePriceRequest instantiates a new DtoCreatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreatePriceRequestWithDefaults

`func NewDtoCreatePriceRequestWithDefaults() *DtoCreatePriceRequest`

NewDtoCreatePriceRequestWithDefaults instantiates a new DtoCreatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoCreatePriceRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreatePriceRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreatePriceRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoCreatePriceRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoCreatePriceRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoCreatePriceRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoCreatePriceRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingModel

`func (o *DtoCreatePriceRequest) GetBillingModel() TypesBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *DtoCreatePriceRequest) GetBillingModelOk() (*TypesBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *DtoCreatePriceRequest) SetBillingModel(v TypesBillingModel)`

SetBillingModel sets BillingModel field to given value.


### GetBillingPeriod

`func (o *DtoCreatePriceRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoCreatePriceRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoCreatePriceRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoCreatePriceRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoCreatePriceRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoCreatePriceRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.


### GetCurrency

`func (o *DtoCreatePriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreatePriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreatePriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *DtoCreatePriceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreatePriceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreatePriceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreatePriceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterValues

`func (o *DtoCreatePriceRequest) GetFilterValues() map[string][]string`

GetFilterValues returns the FilterValues field if non-nil, zero value otherwise.

### GetFilterValuesOk

`func (o *DtoCreatePriceRequest) GetFilterValuesOk() (*map[string][]string, bool)`

GetFilterValuesOk returns a tuple with the FilterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValues

`func (o *DtoCreatePriceRequest) SetFilterValues(v map[string][]string)`

SetFilterValues sets FilterValues field to given value.

### HasFilterValues

`func (o *DtoCreatePriceRequest) HasFilterValues() bool`

HasFilterValues returns a boolean if a field has been set.

### GetInvoiceCadence

`func (o *DtoCreatePriceRequest) GetInvoiceCadence() TypesInvoiceCadence`

GetInvoiceCadence returns the InvoiceCadence field if non-nil, zero value otherwise.

### GetInvoiceCadenceOk

`func (o *DtoCreatePriceRequest) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool)`

GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCadence

`func (o *DtoCreatePriceRequest) SetInvoiceCadence(v TypesInvoiceCadence)`

SetInvoiceCadence sets InvoiceCadence field to given value.


### GetLookupKey

`func (o *DtoCreatePriceRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreatePriceRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreatePriceRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreatePriceRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreatePriceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreatePriceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreatePriceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreatePriceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoCreatePriceRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoCreatePriceRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoCreatePriceRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoCreatePriceRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreatePriceRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreatePriceRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreatePriceRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreatePriceRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTierMode

`func (o *DtoCreatePriceRequest) GetTierMode() TypesBillingTier`

GetTierMode returns the TierMode field if non-nil, zero value otherwise.

### GetTierModeOk

`func (o *DtoCreatePriceRequest) GetTierModeOk() (*TypesBillingTier, bool)`

GetTierModeOk returns a tuple with the TierMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMode

`func (o *DtoCreatePriceRequest) SetTierMode(v TypesBillingTier)`

SetTierMode sets TierMode field to given value.

### HasTierMode

`func (o *DtoCreatePriceRequest) HasTierMode() bool`

HasTierMode returns a boolean if a field has been set.

### GetTiers

`func (o *DtoCreatePriceRequest) GetTiers() []DtoCreatePriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *DtoCreatePriceRequest) GetTiersOk() (*[]DtoCreatePriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *DtoCreatePriceRequest) SetTiers(v []DtoCreatePriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *DtoCreatePriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *DtoCreatePriceRequest) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *DtoCreatePriceRequest) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *DtoCreatePriceRequest) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *DtoCreatePriceRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *DtoCreatePriceRequest) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *DtoCreatePriceRequest) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *DtoCreatePriceRequest) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *DtoCreatePriceRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetType

`func (o *DtoCreatePriceRequest) GetType() TypesPriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreatePriceRequest) GetTypeOk() (*TypesPriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreatePriceRequest) SetType(v TypesPriceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


