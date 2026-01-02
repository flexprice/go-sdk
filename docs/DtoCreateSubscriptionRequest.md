# DtoCreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]DtoAddAddonToSubscriptionRequest**](DtoAddAddonToSubscriptionRequest.md) | Addons represents addons to be added to the subscription during creation | [optional] 
**BillingCadence** | [**TypesBillingCadence**](TypesBillingCadence.md) |  | 
**BillingCycle** | Pointer to [**TypesBillingCycle**](TypesBillingCycle.md) |  | [optional] 
**BillingPeriod** | [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | 
**BillingPeriodCount** | Pointer to **int32** |  | [optional] 
**CollectionMethod** | Pointer to [**TypesCollectionMethod**](TypesCollectionMethod.md) |  | [optional] 
**CommitmentAmount** | Pointer to **string** | CommitmentAmount is the minimum amount a customer commits to paying for a billing period | [optional] 
**Coupons** | Pointer to **[]string** |  | [optional] 
**CreditGrants** | Pointer to [**[]DtoCreateCreditGrantRequest**](DtoCreateCreditGrantRequest.md) | Credit grants to be applied when subscription is created | [optional] 
**Currency** | **string** |  | 
**CustomerId** | Pointer to **string** | customer_id is the flexprice customer id and it is prioritized over external_customer_id in case both are provided. | [optional] 
**CustomerTimezone** | Pointer to **string** | Timezone of the customer. If not set, the default value is UTC. | [optional] 
**EnableTrueUp** | Pointer to **bool** | Enable Commitment True Up Fee | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** | external_customer_id is the customer id in your DB and must be same as what you provided as external_id while creating the customer in flexprice. | [optional] 
**GatewayPaymentMethodId** | Pointer to **string** |  | [optional] 
**InvoiceBilling** | Pointer to [**TypesInvoiceBilling**](TypesInvoiceBilling.md) |  | [optional] 
**LineItemCommitments** | Pointer to [**map[string]DtoLineItemCommitmentConfig**](DtoLineItemCommitmentConfig.md) | LineItemCommitments allows setting commitment configuration per line item (keyed by price_id) | [optional] 
**LineItemCoupons** | Pointer to **map[string][]string** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OverageFactor** | Pointer to **string** | OverageFactor is a multiplier applied to usage beyond the commitment amount | [optional] 
**OverrideEntitlements** | Pointer to [**[]DtoOverrideEntitlementRequest**](DtoOverrideEntitlementRequest.md) | OverrideEntitlements allows customizing specific entitlements for this subscription | [optional] 
**OverrideLineItems** | Pointer to [**[]DtoOverrideLineItemRequest**](DtoOverrideLineItemRequest.md) | OverrideLineItems allows customizing specific prices for this subscription | [optional] 
**PaymentBehavior** | Pointer to [**TypesPaymentBehavior**](TypesPaymentBehavior.md) |  | [optional] 
**Phases** | Pointer to [**[]DtoSubscriptionPhaseCreateRequest**](DtoSubscriptionPhaseCreateRequest.md) | Phases represents subscription phases to be created with the subscription | [optional] 
**PlanId** | **string** |  | 
**ProrationBehavior** | Pointer to [**TypesProrationBehavior**](TypesProrationBehavior.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**SubscriptionStatus** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 
**TaxRateOverrides** | Pointer to [**[]DtoTaxRateOverride**](DtoTaxRateOverride.md) | tax_rate_overrides is the tax rate overrides to be applied to the subscription | [optional] 
**TrialEnd** | Pointer to **string** |  | [optional] 
**TrialStart** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateSubscriptionRequest

`func NewDtoCreateSubscriptionRequest(billingCadence TypesBillingCadence, billingPeriod TypesBillingPeriod, currency string, planId string, ) *DtoCreateSubscriptionRequest`

NewDtoCreateSubscriptionRequest instantiates a new DtoCreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateSubscriptionRequestWithDefaults

`func NewDtoCreateSubscriptionRequestWithDefaults() *DtoCreateSubscriptionRequest`

NewDtoCreateSubscriptionRequestWithDefaults instantiates a new DtoCreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *DtoCreateSubscriptionRequest) GetAddons() []DtoAddAddonToSubscriptionRequest`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *DtoCreateSubscriptionRequest) GetAddonsOk() (*[]DtoAddAddonToSubscriptionRequest, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *DtoCreateSubscriptionRequest) SetAddons(v []DtoAddAddonToSubscriptionRequest)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *DtoCreateSubscriptionRequest) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoCreateSubscriptionRequest) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoCreateSubscriptionRequest) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoCreateSubscriptionRequest) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.


### GetBillingCycle

`func (o *DtoCreateSubscriptionRequest) GetBillingCycle() TypesBillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DtoCreateSubscriptionRequest) GetBillingCycleOk() (*TypesBillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DtoCreateSubscriptionRequest) SetBillingCycle(v TypesBillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DtoCreateSubscriptionRequest) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoCreateSubscriptionRequest) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodCount

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoCreateSubscriptionRequest) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoCreateSubscriptionRequest) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *DtoCreateSubscriptionRequest) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetCollectionMethod

`func (o *DtoCreateSubscriptionRequest) GetCollectionMethod() TypesCollectionMethod`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *DtoCreateSubscriptionRequest) GetCollectionMethodOk() (*TypesCollectionMethod, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *DtoCreateSubscriptionRequest) SetCollectionMethod(v TypesCollectionMethod)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *DtoCreateSubscriptionRequest) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetCommitmentAmount

`func (o *DtoCreateSubscriptionRequest) GetCommitmentAmount() string`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoCreateSubscriptionRequest) GetCommitmentAmountOk() (*string, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoCreateSubscriptionRequest) SetCommitmentAmount(v string)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoCreateSubscriptionRequest) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCoupons

`func (o *DtoCreateSubscriptionRequest) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *DtoCreateSubscriptionRequest) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *DtoCreateSubscriptionRequest) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *DtoCreateSubscriptionRequest) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetCreditGrants

`func (o *DtoCreateSubscriptionRequest) GetCreditGrants() []DtoCreateCreditGrantRequest`

GetCreditGrants returns the CreditGrants field if non-nil, zero value otherwise.

### GetCreditGrantsOk

`func (o *DtoCreateSubscriptionRequest) GetCreditGrantsOk() (*[]DtoCreateCreditGrantRequest, bool)`

GetCreditGrantsOk returns a tuple with the CreditGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrants

`func (o *DtoCreateSubscriptionRequest) SetCreditGrants(v []DtoCreateCreditGrantRequest)`

SetCreditGrants sets CreditGrants field to given value.

### HasCreditGrants

`func (o *DtoCreateSubscriptionRequest) HasCreditGrants() bool`

HasCreditGrants returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreateSubscriptionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateSubscriptionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateSubscriptionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *DtoCreateSubscriptionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCreateSubscriptionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCreateSubscriptionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCreateSubscriptionRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerTimezone

`func (o *DtoCreateSubscriptionRequest) GetCustomerTimezone() string`

GetCustomerTimezone returns the CustomerTimezone field if non-nil, zero value otherwise.

### GetCustomerTimezoneOk

`func (o *DtoCreateSubscriptionRequest) GetCustomerTimezoneOk() (*string, bool)`

GetCustomerTimezoneOk returns a tuple with the CustomerTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTimezone

`func (o *DtoCreateSubscriptionRequest) SetCustomerTimezone(v string)`

SetCustomerTimezone sets CustomerTimezone field to given value.

### HasCustomerTimezone

`func (o *DtoCreateSubscriptionRequest) HasCustomerTimezone() bool`

HasCustomerTimezone returns a boolean if a field has been set.

### GetEnableTrueUp

`func (o *DtoCreateSubscriptionRequest) GetEnableTrueUp() bool`

GetEnableTrueUp returns the EnableTrueUp field if non-nil, zero value otherwise.

### GetEnableTrueUpOk

`func (o *DtoCreateSubscriptionRequest) GetEnableTrueUpOk() (*bool, bool)`

GetEnableTrueUpOk returns a tuple with the EnableTrueUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrueUp

`func (o *DtoCreateSubscriptionRequest) SetEnableTrueUp(v bool)`

SetEnableTrueUp sets EnableTrueUp field to given value.

### HasEnableTrueUp

`func (o *DtoCreateSubscriptionRequest) HasEnableTrueUp() bool`

HasEnableTrueUp returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoCreateSubscriptionRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoCreateSubscriptionRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoCreateSubscriptionRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoCreateSubscriptionRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *DtoCreateSubscriptionRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoCreateSubscriptionRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoCreateSubscriptionRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoCreateSubscriptionRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetGatewayPaymentMethodId

`func (o *DtoCreateSubscriptionRequest) GetGatewayPaymentMethodId() string`

GetGatewayPaymentMethodId returns the GatewayPaymentMethodId field if non-nil, zero value otherwise.

### GetGatewayPaymentMethodIdOk

`func (o *DtoCreateSubscriptionRequest) GetGatewayPaymentMethodIdOk() (*string, bool)`

GetGatewayPaymentMethodIdOk returns a tuple with the GatewayPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentMethodId

`func (o *DtoCreateSubscriptionRequest) SetGatewayPaymentMethodId(v string)`

SetGatewayPaymentMethodId sets GatewayPaymentMethodId field to given value.

### HasGatewayPaymentMethodId

`func (o *DtoCreateSubscriptionRequest) HasGatewayPaymentMethodId() bool`

HasGatewayPaymentMethodId returns a boolean if a field has been set.

### GetInvoiceBilling

`func (o *DtoCreateSubscriptionRequest) GetInvoiceBilling() TypesInvoiceBilling`

GetInvoiceBilling returns the InvoiceBilling field if non-nil, zero value otherwise.

### GetInvoiceBillingOk

`func (o *DtoCreateSubscriptionRequest) GetInvoiceBillingOk() (*TypesInvoiceBilling, bool)`

GetInvoiceBillingOk returns a tuple with the InvoiceBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceBilling

`func (o *DtoCreateSubscriptionRequest) SetInvoiceBilling(v TypesInvoiceBilling)`

SetInvoiceBilling sets InvoiceBilling field to given value.

### HasInvoiceBilling

`func (o *DtoCreateSubscriptionRequest) HasInvoiceBilling() bool`

HasInvoiceBilling returns a boolean if a field has been set.

### GetLineItemCommitments

`func (o *DtoCreateSubscriptionRequest) GetLineItemCommitments() map[string]DtoLineItemCommitmentConfig`

GetLineItemCommitments returns the LineItemCommitments field if non-nil, zero value otherwise.

### GetLineItemCommitmentsOk

`func (o *DtoCreateSubscriptionRequest) GetLineItemCommitmentsOk() (*map[string]DtoLineItemCommitmentConfig, bool)`

GetLineItemCommitmentsOk returns a tuple with the LineItemCommitments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCommitments

`func (o *DtoCreateSubscriptionRequest) SetLineItemCommitments(v map[string]DtoLineItemCommitmentConfig)`

SetLineItemCommitments sets LineItemCommitments field to given value.

### HasLineItemCommitments

`func (o *DtoCreateSubscriptionRequest) HasLineItemCommitments() bool`

HasLineItemCommitments returns a boolean if a field has been set.

### GetLineItemCoupons

`func (o *DtoCreateSubscriptionRequest) GetLineItemCoupons() map[string][]string`

GetLineItemCoupons returns the LineItemCoupons field if non-nil, zero value otherwise.

### GetLineItemCouponsOk

`func (o *DtoCreateSubscriptionRequest) GetLineItemCouponsOk() (*map[string][]string, bool)`

GetLineItemCouponsOk returns a tuple with the LineItemCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCoupons

`func (o *DtoCreateSubscriptionRequest) SetLineItemCoupons(v map[string][]string)`

SetLineItemCoupons sets LineItemCoupons field to given value.

### HasLineItemCoupons

`func (o *DtoCreateSubscriptionRequest) HasLineItemCoupons() bool`

HasLineItemCoupons returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoCreateSubscriptionRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreateSubscriptionRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreateSubscriptionRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreateSubscriptionRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateSubscriptionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateSubscriptionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateSubscriptionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoCreateSubscriptionRequest) GetOverageFactor() string`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoCreateSubscriptionRequest) GetOverageFactorOk() (*string, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoCreateSubscriptionRequest) SetOverageFactor(v string)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoCreateSubscriptionRequest) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

### GetOverrideEntitlements

`func (o *DtoCreateSubscriptionRequest) GetOverrideEntitlements() []DtoOverrideEntitlementRequest`

GetOverrideEntitlements returns the OverrideEntitlements field if non-nil, zero value otherwise.

### GetOverrideEntitlementsOk

`func (o *DtoCreateSubscriptionRequest) GetOverrideEntitlementsOk() (*[]DtoOverrideEntitlementRequest, bool)`

GetOverrideEntitlementsOk returns a tuple with the OverrideEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideEntitlements

`func (o *DtoCreateSubscriptionRequest) SetOverrideEntitlements(v []DtoOverrideEntitlementRequest)`

SetOverrideEntitlements sets OverrideEntitlements field to given value.

### HasOverrideEntitlements

`func (o *DtoCreateSubscriptionRequest) HasOverrideEntitlements() bool`

HasOverrideEntitlements returns a boolean if a field has been set.

### GetOverrideLineItems

`func (o *DtoCreateSubscriptionRequest) GetOverrideLineItems() []DtoOverrideLineItemRequest`

GetOverrideLineItems returns the OverrideLineItems field if non-nil, zero value otherwise.

### GetOverrideLineItemsOk

`func (o *DtoCreateSubscriptionRequest) GetOverrideLineItemsOk() (*[]DtoOverrideLineItemRequest, bool)`

GetOverrideLineItemsOk returns a tuple with the OverrideLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLineItems

`func (o *DtoCreateSubscriptionRequest) SetOverrideLineItems(v []DtoOverrideLineItemRequest)`

SetOverrideLineItems sets OverrideLineItems field to given value.

### HasOverrideLineItems

`func (o *DtoCreateSubscriptionRequest) HasOverrideLineItems() bool`

HasOverrideLineItems returns a boolean if a field has been set.

### GetPaymentBehavior

`func (o *DtoCreateSubscriptionRequest) GetPaymentBehavior() TypesPaymentBehavior`

GetPaymentBehavior returns the PaymentBehavior field if non-nil, zero value otherwise.

### GetPaymentBehaviorOk

`func (o *DtoCreateSubscriptionRequest) GetPaymentBehaviorOk() (*TypesPaymentBehavior, bool)`

GetPaymentBehaviorOk returns a tuple with the PaymentBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBehavior

`func (o *DtoCreateSubscriptionRequest) SetPaymentBehavior(v TypesPaymentBehavior)`

SetPaymentBehavior sets PaymentBehavior field to given value.

### HasPaymentBehavior

`func (o *DtoCreateSubscriptionRequest) HasPaymentBehavior() bool`

HasPaymentBehavior returns a boolean if a field has been set.

### GetPhases

`func (o *DtoCreateSubscriptionRequest) GetPhases() []DtoSubscriptionPhaseCreateRequest`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *DtoCreateSubscriptionRequest) GetPhasesOk() (*[]DtoSubscriptionPhaseCreateRequest, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *DtoCreateSubscriptionRequest) SetPhases(v []DtoSubscriptionPhaseCreateRequest)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *DtoCreateSubscriptionRequest) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreateSubscriptionRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreateSubscriptionRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreateSubscriptionRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetProrationBehavior

`func (o *DtoCreateSubscriptionRequest) GetProrationBehavior() TypesProrationBehavior`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DtoCreateSubscriptionRequest) GetProrationBehaviorOk() (*TypesProrationBehavior, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DtoCreateSubscriptionRequest) SetProrationBehavior(v TypesProrationBehavior)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *DtoCreateSubscriptionRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoCreateSubscriptionRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoCreateSubscriptionRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoCreateSubscriptionRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoCreateSubscriptionRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *DtoCreateSubscriptionRequest) GetSubscriptionStatus() TypesSubscriptionStatus`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *DtoCreateSubscriptionRequest) GetSubscriptionStatusOk() (*TypesSubscriptionStatus, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *DtoCreateSubscriptionRequest) SetSubscriptionStatus(v TypesSubscriptionStatus)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *DtoCreateSubscriptionRequest) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTaxRateOverrides

`func (o *DtoCreateSubscriptionRequest) GetTaxRateOverrides() []DtoTaxRateOverride`

GetTaxRateOverrides returns the TaxRateOverrides field if non-nil, zero value otherwise.

### GetTaxRateOverridesOk

`func (o *DtoCreateSubscriptionRequest) GetTaxRateOverridesOk() (*[]DtoTaxRateOverride, bool)`

GetTaxRateOverridesOk returns a tuple with the TaxRateOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateOverrides

`func (o *DtoCreateSubscriptionRequest) SetTaxRateOverrides(v []DtoTaxRateOverride)`

SetTaxRateOverrides sets TaxRateOverrides field to given value.

### HasTaxRateOverrides

`func (o *DtoCreateSubscriptionRequest) HasTaxRateOverrides() bool`

HasTaxRateOverrides returns a boolean if a field has been set.

### GetTrialEnd

`func (o *DtoCreateSubscriptionRequest) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *DtoCreateSubscriptionRequest) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *DtoCreateSubscriptionRequest) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *DtoCreateSubscriptionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *DtoCreateSubscriptionRequest) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *DtoCreateSubscriptionRequest) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *DtoCreateSubscriptionRequest) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *DtoCreateSubscriptionRequest) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


