# DtoUsageAnalyticItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOnId** | Pointer to **string** |  | [optional] 
**Addon** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainAddonAddon**](GithubComFlexpriceFlexpriceInternalDomainAddonAddon.md) |  | [optional] 
**AggregationType** | Pointer to [**TypesAggregationType**](TypesAggregationType.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EventCount** | Pointer to **int32** | Number of events that contributed to this aggregation | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**Feature** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainFeatureFeature**](GithubComFlexpriceFlexpriceInternalDomainFeatureFeature.md) |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**Meter** | Pointer to [**MeterMeter**](MeterMeter.md) |  | [optional] 
**MeterId** | Pointer to **string** | Meter ID | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainPlanPlan**](GithubComFlexpriceFlexpriceInternalDomainPlanPlan.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Points** | Pointer to [**[]DtoUsageAnalyticPoint**](DtoUsageAnalyticPoint.md) |  | [optional] 
**Price** | Pointer to [**DtoPriceResponse**](DtoPriceResponse.md) |  | [optional] 
**PriceId** | Pointer to **string** | Price ID used for this usage | [optional] 
**Properties** | Pointer to **map[string]string** | Stores property values for flexible grouping (e.g., org_id -&gt; \&quot;org123\&quot;) | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SubLineItemId** | Pointer to **string** | Subscription line item ID | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription ID | [optional] 
**SubscriptionLineItem** | Pointer to [**SubscriptionSubscriptionLineItem**](SubscriptionSubscriptionLineItem.md) |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**TotalUsage** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UnitPlural** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUsageAnalyticItem

`func NewDtoUsageAnalyticItem() *DtoUsageAnalyticItem`

NewDtoUsageAnalyticItem instantiates a new DtoUsageAnalyticItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUsageAnalyticItemWithDefaults

`func NewDtoUsageAnalyticItemWithDefaults() *DtoUsageAnalyticItem`

NewDtoUsageAnalyticItemWithDefaults instantiates a new DtoUsageAnalyticItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOnId

`func (o *DtoUsageAnalyticItem) GetAddOnId() string`

GetAddOnId returns the AddOnId field if non-nil, zero value otherwise.

### GetAddOnIdOk

`func (o *DtoUsageAnalyticItem) GetAddOnIdOk() (*string, bool)`

GetAddOnIdOk returns a tuple with the AddOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnId

`func (o *DtoUsageAnalyticItem) SetAddOnId(v string)`

SetAddOnId sets AddOnId field to given value.

### HasAddOnId

`func (o *DtoUsageAnalyticItem) HasAddOnId() bool`

HasAddOnId returns a boolean if a field has been set.

### GetAddon

`func (o *DtoUsageAnalyticItem) GetAddon() GithubComFlexpriceFlexpriceInternalDomainAddonAddon`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *DtoUsageAnalyticItem) GetAddonOk() (*GithubComFlexpriceFlexpriceInternalDomainAddonAddon, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *DtoUsageAnalyticItem) SetAddon(v GithubComFlexpriceFlexpriceInternalDomainAddonAddon)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *DtoUsageAnalyticItem) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAggregationType

`func (o *DtoUsageAnalyticItem) GetAggregationType() TypesAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *DtoUsageAnalyticItem) GetAggregationTypeOk() (*TypesAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *DtoUsageAnalyticItem) SetAggregationType(v TypesAggregationType)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *DtoUsageAnalyticItem) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoUsageAnalyticItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoUsageAnalyticItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoUsageAnalyticItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoUsageAnalyticItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEventCount

`func (o *DtoUsageAnalyticItem) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *DtoUsageAnalyticItem) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *DtoUsageAnalyticItem) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.

### HasEventCount

`func (o *DtoUsageAnalyticItem) HasEventCount() bool`

HasEventCount returns a boolean if a field has been set.

### GetEventName

`func (o *DtoUsageAnalyticItem) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoUsageAnalyticItem) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoUsageAnalyticItem) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DtoUsageAnalyticItem) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetFeature

`func (o *DtoUsageAnalyticItem) GetFeature() GithubComFlexpriceFlexpriceInternalDomainFeatureFeature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DtoUsageAnalyticItem) GetFeatureOk() (*GithubComFlexpriceFlexpriceInternalDomainFeatureFeature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DtoUsageAnalyticItem) SetFeature(v GithubComFlexpriceFlexpriceInternalDomainFeatureFeature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DtoUsageAnalyticItem) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureId

`func (o *DtoUsageAnalyticItem) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *DtoUsageAnalyticItem) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *DtoUsageAnalyticItem) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *DtoUsageAnalyticItem) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMeter

`func (o *DtoUsageAnalyticItem) GetMeter() MeterMeter`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *DtoUsageAnalyticItem) GetMeterOk() (*MeterMeter, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *DtoUsageAnalyticItem) SetMeter(v MeterMeter)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *DtoUsageAnalyticItem) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoUsageAnalyticItem) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoUsageAnalyticItem) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoUsageAnalyticItem) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoUsageAnalyticItem) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetName

`func (o *DtoUsageAnalyticItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUsageAnalyticItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUsageAnalyticItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUsageAnalyticItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *DtoUsageAnalyticItem) GetPlan() GithubComFlexpriceFlexpriceInternalDomainPlanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *DtoUsageAnalyticItem) GetPlanOk() (*GithubComFlexpriceFlexpriceInternalDomainPlanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *DtoUsageAnalyticItem) SetPlan(v GithubComFlexpriceFlexpriceInternalDomainPlanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *DtoUsageAnalyticItem) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoUsageAnalyticItem) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoUsageAnalyticItem) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoUsageAnalyticItem) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoUsageAnalyticItem) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPoints

`func (o *DtoUsageAnalyticItem) GetPoints() []DtoUsageAnalyticPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DtoUsageAnalyticItem) GetPointsOk() (*[]DtoUsageAnalyticPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *DtoUsageAnalyticItem) SetPoints(v []DtoUsageAnalyticPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *DtoUsageAnalyticItem) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPrice

`func (o *DtoUsageAnalyticItem) GetPrice() DtoPriceResponse`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DtoUsageAnalyticItem) GetPriceOk() (*DtoPriceResponse, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DtoUsageAnalyticItem) SetPrice(v DtoPriceResponse)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DtoUsageAnalyticItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoUsageAnalyticItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoUsageAnalyticItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoUsageAnalyticItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoUsageAnalyticItem) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetProperties

`func (o *DtoUsageAnalyticItem) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DtoUsageAnalyticItem) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DtoUsageAnalyticItem) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DtoUsageAnalyticItem) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSource

`func (o *DtoUsageAnalyticItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DtoUsageAnalyticItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DtoUsageAnalyticItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DtoUsageAnalyticItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubLineItemId

`func (o *DtoUsageAnalyticItem) GetSubLineItemId() string`

GetSubLineItemId returns the SubLineItemId field if non-nil, zero value otherwise.

### GetSubLineItemIdOk

`func (o *DtoUsageAnalyticItem) GetSubLineItemIdOk() (*string, bool)`

GetSubLineItemIdOk returns a tuple with the SubLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLineItemId

`func (o *DtoUsageAnalyticItem) SetSubLineItemId(v string)`

SetSubLineItemId sets SubLineItemId field to given value.

### HasSubLineItemId

`func (o *DtoUsageAnalyticItem) HasSubLineItemId() bool`

HasSubLineItemId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoUsageAnalyticItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoUsageAnalyticItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoUsageAnalyticItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoUsageAnalyticItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionLineItem

`func (o *DtoUsageAnalyticItem) GetSubscriptionLineItem() SubscriptionSubscriptionLineItem`

GetSubscriptionLineItem returns the SubscriptionLineItem field if non-nil, zero value otherwise.

### GetSubscriptionLineItemOk

`func (o *DtoUsageAnalyticItem) GetSubscriptionLineItemOk() (*SubscriptionSubscriptionLineItem, bool)`

GetSubscriptionLineItemOk returns a tuple with the SubscriptionLineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLineItem

`func (o *DtoUsageAnalyticItem) SetSubscriptionLineItem(v SubscriptionSubscriptionLineItem)`

SetSubscriptionLineItem sets SubscriptionLineItem field to given value.

### HasSubscriptionLineItem

`func (o *DtoUsageAnalyticItem) HasSubscriptionLineItem() bool`

HasSubscriptionLineItem returns a boolean if a field has been set.

### GetTotalCost

`func (o *DtoUsageAnalyticItem) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *DtoUsageAnalyticItem) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *DtoUsageAnalyticItem) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *DtoUsageAnalyticItem) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalUsage

`func (o *DtoUsageAnalyticItem) GetTotalUsage() string`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *DtoUsageAnalyticItem) GetTotalUsageOk() (*string, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *DtoUsageAnalyticItem) SetTotalUsage(v string)`

SetTotalUsage sets TotalUsage field to given value.

### HasTotalUsage

`func (o *DtoUsageAnalyticItem) HasTotalUsage() bool`

HasTotalUsage returns a boolean if a field has been set.

### GetUnit

`func (o *DtoUsageAnalyticItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DtoUsageAnalyticItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DtoUsageAnalyticItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DtoUsageAnalyticItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPlural

`func (o *DtoUsageAnalyticItem) GetUnitPlural() string`

GetUnitPlural returns the UnitPlural field if non-nil, zero value otherwise.

### GetUnitPluralOk

`func (o *DtoUsageAnalyticItem) GetUnitPluralOk() (*string, bool)`

GetUnitPluralOk returns a tuple with the UnitPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPlural

`func (o *DtoUsageAnalyticItem) SetUnitPlural(v string)`

SetUnitPlural sets UnitPlural field to given value.

### HasUnitPlural

`func (o *DtoUsageAnalyticItem) HasUnitPlural() bool`

HasUnitPlural returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


