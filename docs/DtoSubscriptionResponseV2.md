# DtoSubscriptionResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePauseId** | Pointer to **string** | ActivePauseID references the current active pause configuration This will be null if no pause is active or scheduled | [optional] 
**BillingAnchor** | Pointer to **string** | BillingAnchor is the reference point that aligns future billing cycle dates. It sets the day of week for week intervals, the day of month for month and year intervals, and the month of year for year intervals. The timestamp is in UTC format. | [optional] 
**BillingCadence** | Pointer to [**TypesBillingCadence**](TypesBillingCadence.md) |  | [optional] 
**BillingCycle** | Pointer to [**TypesBillingCycle**](TypesBillingCycle.md) |  | [optional] 
**BillingPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 
**BillingPeriodCount** | Pointer to **int32** | BillingPeriodCount is the total number units of the billing period. | [optional] 
**CancelAt** | Pointer to **string** | CancelAt is the date the subscription will be canceled | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** | CancelAtPeriodEnd is whether the subscription was canceled at the end of the current period | [optional] 
**CancelledAt** | Pointer to **string** | CanceledAt is the date the subscription was canceled | [optional] 
**CollectionMethod** | Pointer to **string** | CollectionMethod determines how invoices are collected | [optional] 
**CommitmentAmount** | Pointer to **string** | CommitmentAmount is the minimum amount a customer commits to paying for a billing period | [optional] 
**CommitmentDuration** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 
**CouponAssociations** | Pointer to [**[]DtoCouponAssociationResponse**](DtoCouponAssociationResponse.md) | CouponAssociations are included when \&quot;coupon_associations\&quot; is in expand parameter | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditGrants** | Pointer to [**[]DtoCreditGrantResponse**](DtoCreditGrantResponse.md) | CreditGrants are included when \&quot;credit_grants\&quot; is in expand parameter | [optional] 
**Currency** | Pointer to **string** | Currency is the currency of the subscription in lowercase 3 digit ISO codes | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | CurrentPeriodEnd is the end of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created. | [optional] 
**CurrentPeriodStart** | Pointer to **string** | CurrentPeriodStart is the end of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created. | [optional] 
**Customer** | Pointer to [**DtoCustomerResponse**](DtoCustomerResponse.md) |  | [optional] 
**CustomerId** | Pointer to **string** | CustomerID is the identifier for the customer in our system | [optional] 
**CustomerTimezone** | Pointer to **string** |  | [optional] 
**EnableTrueUp** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **string** | EndDate is the end date of the subscription | [optional] 
**EnvironmentId** | Pointer to **string** | EnvironmentID is the environment identifier for the subscription | [optional] 
**GatewayPaymentMethodId** | Pointer to **string** | GatewayPaymentMethodID is the gateway payment method ID for this subscription | [optional] 
**Id** | Pointer to **string** | ID is the unique identifier for the subscription | [optional] 
**InvoicingCustomerId** | Pointer to **string** | InvoicingCustomerID is the customer ID to use for invoicing This can differ from the subscription customer (e.g., parent company invoicing for child company) | [optional] 
**LineItems** | Pointer to [**[]DtoSubscriptionLineItemResponse**](DtoSubscriptionLineItemResponse.md) | LineItems is expanded only if \&quot;subscription_line_items\&quot; is in expand parameter Each line item can optionally include expanded price data | [optional] 
**LookupKey** | Pointer to **string** | LookupKey is the key used to lookup the subscription in our system | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OverageFactor** | Pointer to **string** | OverageFactor is a multiplier applied to usage beyond the commitment amount | [optional] 
**ParentSubscriptionId** | Pointer to **string** | ParentSubscriptionID is the parent subscription ID for hierarchy (e.g. child subscription under a parent) | [optional] 
**PauseStatus** | Pointer to [**TypesPauseStatus**](TypesPauseStatus.md) |  | [optional] 
**Pauses** | Pointer to [**[]SubscriptionSubscriptionPause**](SubscriptionSubscriptionPause.md) | Pauses are included when subscription has pause status | [optional] 
**PaymentBehavior** | Pointer to **string** | PaymentBehavior determines how subscription payments are handled | [optional] 
**PaymentTerms** | Pointer to [**TypesPaymentTerms**](TypesPaymentTerms.md) |  | [optional] 
**Phases** | Pointer to [**[]DtoSubscriptionPhaseResponse**](DtoSubscriptionPhaseResponse.md) | Phases are included when \&quot;phases\&quot; is in expand parameter | [optional] 
**Plan** | Pointer to [**DtoPlanResponse**](DtoPlanResponse.md) |  | [optional] 
**PlanId** | Pointer to **string** | PlanID is the identifier for the plan in our system | [optional] 
**ProrationBehavior** | Pointer to [**TypesProrationBehavior**](TypesProrationBehavior.md) |  | [optional] 
**StartDate** | Pointer to **string** | StartDate is the start date of the subscription | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionStatus** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TrialEnd** | Pointer to **string** | TrialEnd is the end date of the trial period | [optional] 
**TrialStart** | Pointer to **string** | TrialStart is the start date of the trial period | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** | Version is used for optimistic locking | [optional] 

## Methods

### NewDtoSubscriptionResponseV2

`func NewDtoSubscriptionResponseV2() *DtoSubscriptionResponseV2`

NewDtoSubscriptionResponseV2 instantiates a new DtoSubscriptionResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionResponseV2WithDefaults

`func NewDtoSubscriptionResponseV2WithDefaults() *DtoSubscriptionResponseV2`

NewDtoSubscriptionResponseV2WithDefaults instantiates a new DtoSubscriptionResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePauseId

`func (o *DtoSubscriptionResponseV2) GetActivePauseId() string`

GetActivePauseId returns the ActivePauseId field if non-nil, zero value otherwise.

### GetActivePauseIdOk

`func (o *DtoSubscriptionResponseV2) GetActivePauseIdOk() (*string, bool)`

GetActivePauseIdOk returns a tuple with the ActivePauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePauseId

`func (o *DtoSubscriptionResponseV2) SetActivePauseId(v string)`

SetActivePauseId sets ActivePauseId field to given value.

### HasActivePauseId

`func (o *DtoSubscriptionResponseV2) HasActivePauseId() bool`

HasActivePauseId returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *DtoSubscriptionResponseV2) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *DtoSubscriptionResponseV2) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *DtoSubscriptionResponseV2) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *DtoSubscriptionResponseV2) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoSubscriptionResponseV2) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoSubscriptionResponseV2) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoSubscriptionResponseV2) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *DtoSubscriptionResponseV2) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingCycle

`func (o *DtoSubscriptionResponseV2) GetBillingCycle() TypesBillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DtoSubscriptionResponseV2) GetBillingCycleOk() (*TypesBillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DtoSubscriptionResponseV2) SetBillingCycle(v TypesBillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DtoSubscriptionResponseV2) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DtoSubscriptionResponseV2) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoSubscriptionResponseV2) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoSubscriptionResponseV2) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DtoSubscriptionResponseV2) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBillingPeriodCount

`func (o *DtoSubscriptionResponseV2) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoSubscriptionResponseV2) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoSubscriptionResponseV2) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *DtoSubscriptionResponseV2) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetCancelAt

`func (o *DtoSubscriptionResponseV2) GetCancelAt() string`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *DtoSubscriptionResponseV2) GetCancelAtOk() (*string, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *DtoSubscriptionResponseV2) SetCancelAt(v string)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *DtoSubscriptionResponseV2) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *DtoSubscriptionResponseV2) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *DtoSubscriptionResponseV2) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *DtoSubscriptionResponseV2) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *DtoSubscriptionResponseV2) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCancelledAt

`func (o *DtoSubscriptionResponseV2) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DtoSubscriptionResponseV2) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DtoSubscriptionResponseV2) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *DtoSubscriptionResponseV2) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetCollectionMethod

`func (o *DtoSubscriptionResponseV2) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *DtoSubscriptionResponseV2) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *DtoSubscriptionResponseV2) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *DtoSubscriptionResponseV2) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetCommitmentAmount

`func (o *DtoSubscriptionResponseV2) GetCommitmentAmount() string`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoSubscriptionResponseV2) GetCommitmentAmountOk() (*string, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoSubscriptionResponseV2) SetCommitmentAmount(v string)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoSubscriptionResponseV2) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentDuration

`func (o *DtoSubscriptionResponseV2) GetCommitmentDuration() TypesBillingPeriod`

GetCommitmentDuration returns the CommitmentDuration field if non-nil, zero value otherwise.

### GetCommitmentDurationOk

`func (o *DtoSubscriptionResponseV2) GetCommitmentDurationOk() (*TypesBillingPeriod, bool)`

GetCommitmentDurationOk returns a tuple with the CommitmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentDuration

`func (o *DtoSubscriptionResponseV2) SetCommitmentDuration(v TypesBillingPeriod)`

SetCommitmentDuration sets CommitmentDuration field to given value.

### HasCommitmentDuration

`func (o *DtoSubscriptionResponseV2) HasCommitmentDuration() bool`

HasCommitmentDuration returns a boolean if a field has been set.

### GetCouponAssociations

`func (o *DtoSubscriptionResponseV2) GetCouponAssociations() []DtoCouponAssociationResponse`

GetCouponAssociations returns the CouponAssociations field if non-nil, zero value otherwise.

### GetCouponAssociationsOk

`func (o *DtoSubscriptionResponseV2) GetCouponAssociationsOk() (*[]DtoCouponAssociationResponse, bool)`

GetCouponAssociationsOk returns a tuple with the CouponAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAssociations

`func (o *DtoSubscriptionResponseV2) SetCouponAssociations(v []DtoCouponAssociationResponse)`

SetCouponAssociations sets CouponAssociations field to given value.

### HasCouponAssociations

`func (o *DtoSubscriptionResponseV2) HasCouponAssociations() bool`

HasCouponAssociations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoSubscriptionResponseV2) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionResponseV2) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionResponseV2) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionResponseV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoSubscriptionResponseV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoSubscriptionResponseV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoSubscriptionResponseV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoSubscriptionResponseV2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditGrants

`func (o *DtoSubscriptionResponseV2) GetCreditGrants() []DtoCreditGrantResponse`

GetCreditGrants returns the CreditGrants field if non-nil, zero value otherwise.

### GetCreditGrantsOk

`func (o *DtoSubscriptionResponseV2) GetCreditGrantsOk() (*[]DtoCreditGrantResponse, bool)`

GetCreditGrantsOk returns a tuple with the CreditGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrants

`func (o *DtoSubscriptionResponseV2) SetCreditGrants(v []DtoCreditGrantResponse)`

SetCreditGrants sets CreditGrants field to given value.

### HasCreditGrants

`func (o *DtoSubscriptionResponseV2) HasCreditGrants() bool`

HasCreditGrants returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoSubscriptionResponseV2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoSubscriptionResponseV2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoSubscriptionResponseV2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoSubscriptionResponseV2) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *DtoSubscriptionResponseV2) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *DtoSubscriptionResponseV2) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *DtoSubscriptionResponseV2) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *DtoSubscriptionResponseV2) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *DtoSubscriptionResponseV2) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *DtoSubscriptionResponseV2) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *DtoSubscriptionResponseV2) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *DtoSubscriptionResponseV2) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoSubscriptionResponseV2) GetCustomer() DtoCustomerResponse`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoSubscriptionResponseV2) GetCustomerOk() (*DtoCustomerResponse, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoSubscriptionResponseV2) SetCustomer(v DtoCustomerResponse)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoSubscriptionResponseV2) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoSubscriptionResponseV2) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoSubscriptionResponseV2) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoSubscriptionResponseV2) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoSubscriptionResponseV2) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerTimezone

`func (o *DtoSubscriptionResponseV2) GetCustomerTimezone() string`

GetCustomerTimezone returns the CustomerTimezone field if non-nil, zero value otherwise.

### GetCustomerTimezoneOk

`func (o *DtoSubscriptionResponseV2) GetCustomerTimezoneOk() (*string, bool)`

GetCustomerTimezoneOk returns a tuple with the CustomerTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTimezone

`func (o *DtoSubscriptionResponseV2) SetCustomerTimezone(v string)`

SetCustomerTimezone sets CustomerTimezone field to given value.

### HasCustomerTimezone

`func (o *DtoSubscriptionResponseV2) HasCustomerTimezone() bool`

HasCustomerTimezone returns a boolean if a field has been set.

### GetEnableTrueUp

`func (o *DtoSubscriptionResponseV2) GetEnableTrueUp() bool`

GetEnableTrueUp returns the EnableTrueUp field if non-nil, zero value otherwise.

### GetEnableTrueUpOk

`func (o *DtoSubscriptionResponseV2) GetEnableTrueUpOk() (*bool, bool)`

GetEnableTrueUpOk returns a tuple with the EnableTrueUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrueUp

`func (o *DtoSubscriptionResponseV2) SetEnableTrueUp(v bool)`

SetEnableTrueUp sets EnableTrueUp field to given value.

### HasEnableTrueUp

`func (o *DtoSubscriptionResponseV2) HasEnableTrueUp() bool`

HasEnableTrueUp returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoSubscriptionResponseV2) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoSubscriptionResponseV2) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoSubscriptionResponseV2) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoSubscriptionResponseV2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoSubscriptionResponseV2) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoSubscriptionResponseV2) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoSubscriptionResponseV2) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoSubscriptionResponseV2) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetGatewayPaymentMethodId

`func (o *DtoSubscriptionResponseV2) GetGatewayPaymentMethodId() string`

GetGatewayPaymentMethodId returns the GatewayPaymentMethodId field if non-nil, zero value otherwise.

### GetGatewayPaymentMethodIdOk

`func (o *DtoSubscriptionResponseV2) GetGatewayPaymentMethodIdOk() (*string, bool)`

GetGatewayPaymentMethodIdOk returns a tuple with the GatewayPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentMethodId

`func (o *DtoSubscriptionResponseV2) SetGatewayPaymentMethodId(v string)`

SetGatewayPaymentMethodId sets GatewayPaymentMethodId field to given value.

### HasGatewayPaymentMethodId

`func (o *DtoSubscriptionResponseV2) HasGatewayPaymentMethodId() bool`

HasGatewayPaymentMethodId returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionResponseV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionResponseV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionResponseV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionResponseV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoicingCustomerId

`func (o *DtoSubscriptionResponseV2) GetInvoicingCustomerId() string`

GetInvoicingCustomerId returns the InvoicingCustomerId field if non-nil, zero value otherwise.

### GetInvoicingCustomerIdOk

`func (o *DtoSubscriptionResponseV2) GetInvoicingCustomerIdOk() (*string, bool)`

GetInvoicingCustomerIdOk returns a tuple with the InvoicingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCustomerId

`func (o *DtoSubscriptionResponseV2) SetInvoicingCustomerId(v string)`

SetInvoicingCustomerId sets InvoicingCustomerId field to given value.

### HasInvoicingCustomerId

`func (o *DtoSubscriptionResponseV2) HasInvoicingCustomerId() bool`

HasInvoicingCustomerId returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoSubscriptionResponseV2) GetLineItems() []DtoSubscriptionLineItemResponse`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoSubscriptionResponseV2) GetLineItemsOk() (*[]DtoSubscriptionLineItemResponse, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoSubscriptionResponseV2) SetLineItems(v []DtoSubscriptionLineItemResponse)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoSubscriptionResponseV2) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoSubscriptionResponseV2) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoSubscriptionResponseV2) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoSubscriptionResponseV2) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoSubscriptionResponseV2) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionResponseV2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionResponseV2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionResponseV2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionResponseV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoSubscriptionResponseV2) GetOverageFactor() string`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoSubscriptionResponseV2) GetOverageFactorOk() (*string, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoSubscriptionResponseV2) SetOverageFactor(v string)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoSubscriptionResponseV2) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

### GetParentSubscriptionId

`func (o *DtoSubscriptionResponseV2) GetParentSubscriptionId() string`

GetParentSubscriptionId returns the ParentSubscriptionId field if non-nil, zero value otherwise.

### GetParentSubscriptionIdOk

`func (o *DtoSubscriptionResponseV2) GetParentSubscriptionIdOk() (*string, bool)`

GetParentSubscriptionIdOk returns a tuple with the ParentSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSubscriptionId

`func (o *DtoSubscriptionResponseV2) SetParentSubscriptionId(v string)`

SetParentSubscriptionId sets ParentSubscriptionId field to given value.

### HasParentSubscriptionId

`func (o *DtoSubscriptionResponseV2) HasParentSubscriptionId() bool`

HasParentSubscriptionId returns a boolean if a field has been set.

### GetPauseStatus

`func (o *DtoSubscriptionResponseV2) GetPauseStatus() TypesPauseStatus`

GetPauseStatus returns the PauseStatus field if non-nil, zero value otherwise.

### GetPauseStatusOk

`func (o *DtoSubscriptionResponseV2) GetPauseStatusOk() (*TypesPauseStatus, bool)`

GetPauseStatusOk returns a tuple with the PauseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStatus

`func (o *DtoSubscriptionResponseV2) SetPauseStatus(v TypesPauseStatus)`

SetPauseStatus sets PauseStatus field to given value.

### HasPauseStatus

`func (o *DtoSubscriptionResponseV2) HasPauseStatus() bool`

HasPauseStatus returns a boolean if a field has been set.

### GetPauses

`func (o *DtoSubscriptionResponseV2) GetPauses() []SubscriptionSubscriptionPause`

GetPauses returns the Pauses field if non-nil, zero value otherwise.

### GetPausesOk

`func (o *DtoSubscriptionResponseV2) GetPausesOk() (*[]SubscriptionSubscriptionPause, bool)`

GetPausesOk returns a tuple with the Pauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauses

`func (o *DtoSubscriptionResponseV2) SetPauses(v []SubscriptionSubscriptionPause)`

SetPauses sets Pauses field to given value.

### HasPauses

`func (o *DtoSubscriptionResponseV2) HasPauses() bool`

HasPauses returns a boolean if a field has been set.

### GetPaymentBehavior

`func (o *DtoSubscriptionResponseV2) GetPaymentBehavior() string`

GetPaymentBehavior returns the PaymentBehavior field if non-nil, zero value otherwise.

### GetPaymentBehaviorOk

`func (o *DtoSubscriptionResponseV2) GetPaymentBehaviorOk() (*string, bool)`

GetPaymentBehaviorOk returns a tuple with the PaymentBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBehavior

`func (o *DtoSubscriptionResponseV2) SetPaymentBehavior(v string)`

SetPaymentBehavior sets PaymentBehavior field to given value.

### HasPaymentBehavior

`func (o *DtoSubscriptionResponseV2) HasPaymentBehavior() bool`

HasPaymentBehavior returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *DtoSubscriptionResponseV2) GetPaymentTerms() TypesPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *DtoSubscriptionResponseV2) GetPaymentTermsOk() (*TypesPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *DtoSubscriptionResponseV2) SetPaymentTerms(v TypesPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *DtoSubscriptionResponseV2) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetPhases

`func (o *DtoSubscriptionResponseV2) GetPhases() []DtoSubscriptionPhaseResponse`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *DtoSubscriptionResponseV2) GetPhasesOk() (*[]DtoSubscriptionPhaseResponse, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *DtoSubscriptionResponseV2) SetPhases(v []DtoSubscriptionPhaseResponse)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *DtoSubscriptionResponseV2) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetPlan

`func (o *DtoSubscriptionResponseV2) GetPlan() DtoPlanResponse`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *DtoSubscriptionResponseV2) GetPlanOk() (*DtoPlanResponse, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *DtoSubscriptionResponseV2) SetPlan(v DtoPlanResponse)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *DtoSubscriptionResponseV2) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoSubscriptionResponseV2) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoSubscriptionResponseV2) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoSubscriptionResponseV2) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoSubscriptionResponseV2) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *DtoSubscriptionResponseV2) GetProrationBehavior() TypesProrationBehavior`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DtoSubscriptionResponseV2) GetProrationBehaviorOk() (*TypesProrationBehavior, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DtoSubscriptionResponseV2) SetProrationBehavior(v TypesProrationBehavior)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *DtoSubscriptionResponseV2) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoSubscriptionResponseV2) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoSubscriptionResponseV2) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoSubscriptionResponseV2) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoSubscriptionResponseV2) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionResponseV2) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionResponseV2) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionResponseV2) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionResponseV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *DtoSubscriptionResponseV2) GetSubscriptionStatus() TypesSubscriptionStatus`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *DtoSubscriptionResponseV2) GetSubscriptionStatusOk() (*TypesSubscriptionStatus, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *DtoSubscriptionResponseV2) SetSubscriptionStatus(v TypesSubscriptionStatus)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *DtoSubscriptionResponseV2) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoSubscriptionResponseV2) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoSubscriptionResponseV2) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoSubscriptionResponseV2) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoSubscriptionResponseV2) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrialEnd

`func (o *DtoSubscriptionResponseV2) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *DtoSubscriptionResponseV2) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *DtoSubscriptionResponseV2) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *DtoSubscriptionResponseV2) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *DtoSubscriptionResponseV2) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *DtoSubscriptionResponseV2) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *DtoSubscriptionResponseV2) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *DtoSubscriptionResponseV2) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSubscriptionResponseV2) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSubscriptionResponseV2) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSubscriptionResponseV2) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSubscriptionResponseV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoSubscriptionResponseV2) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoSubscriptionResponseV2) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoSubscriptionResponseV2) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoSubscriptionResponseV2) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *DtoSubscriptionResponseV2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DtoSubscriptionResponseV2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DtoSubscriptionResponseV2) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DtoSubscriptionResponseV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


