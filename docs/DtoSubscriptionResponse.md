# DtoSubscriptionResponse

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
**CouponAssociations** | Pointer to [**[]DtoCouponAssociationResponse**](DtoCouponAssociationResponse.md) | CouponAssociations are the coupon associations for this subscription | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditGrants** | Pointer to [**[]DtoCreditGrantResponse**](DtoCreditGrantResponse.md) | Credit grants are the credit grants for this subscription | [optional] 
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
**LatestInvoice** | Pointer to [**DtoInvoiceResponse**](DtoInvoiceResponse.md) |  | [optional] 
**LineItems** | Pointer to [**[]SubscriptionSubscriptionLineItem**](SubscriptionSubscriptionLineItem.md) |  | [optional] 
**LookupKey** | Pointer to **string** | LookupKey is the key used to lookup the subscription in our system | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OverageFactor** | Pointer to **string** | OverageFactor is a multiplier applied to usage beyond the commitment amount | [optional] 
**PauseStatus** | Pointer to [**TypesPauseStatus**](TypesPauseStatus.md) |  | [optional] 
**Pauses** | Pointer to [**[]SubscriptionSubscriptionPause**](SubscriptionSubscriptionPause.md) |  | [optional] 
**PaymentBehavior** | Pointer to **string** | PaymentBehavior determines how subscription payments are handled | [optional] 
**Phases** | Pointer to [**[]DtoSubscriptionPhaseResponse**](DtoSubscriptionPhaseResponse.md) | Phases are the subscription phases for this subscription | [optional] 
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

### NewDtoSubscriptionResponse

`func NewDtoSubscriptionResponse() *DtoSubscriptionResponse`

NewDtoSubscriptionResponse instantiates a new DtoSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionResponseWithDefaults

`func NewDtoSubscriptionResponseWithDefaults() *DtoSubscriptionResponse`

NewDtoSubscriptionResponseWithDefaults instantiates a new DtoSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePauseId

`func (o *DtoSubscriptionResponse) GetActivePauseId() string`

GetActivePauseId returns the ActivePauseId field if non-nil, zero value otherwise.

### GetActivePauseIdOk

`func (o *DtoSubscriptionResponse) GetActivePauseIdOk() (*string, bool)`

GetActivePauseIdOk returns a tuple with the ActivePauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePauseId

`func (o *DtoSubscriptionResponse) SetActivePauseId(v string)`

SetActivePauseId sets ActivePauseId field to given value.

### HasActivePauseId

`func (o *DtoSubscriptionResponse) HasActivePauseId() bool`

HasActivePauseId returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *DtoSubscriptionResponse) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *DtoSubscriptionResponse) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *DtoSubscriptionResponse) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *DtoSubscriptionResponse) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetBillingCadence

`func (o *DtoSubscriptionResponse) GetBillingCadence() TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *DtoSubscriptionResponse) GetBillingCadenceOk() (*TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *DtoSubscriptionResponse) SetBillingCadence(v TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *DtoSubscriptionResponse) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingCycle

`func (o *DtoSubscriptionResponse) GetBillingCycle() TypesBillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DtoSubscriptionResponse) GetBillingCycleOk() (*TypesBillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DtoSubscriptionResponse) SetBillingCycle(v TypesBillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DtoSubscriptionResponse) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DtoSubscriptionResponse) GetBillingPeriod() TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoSubscriptionResponse) GetBillingPeriodOk() (*TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoSubscriptionResponse) SetBillingPeriod(v TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DtoSubscriptionResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBillingPeriodCount

`func (o *DtoSubscriptionResponse) GetBillingPeriodCount() int32`

GetBillingPeriodCount returns the BillingPeriodCount field if non-nil, zero value otherwise.

### GetBillingPeriodCountOk

`func (o *DtoSubscriptionResponse) GetBillingPeriodCountOk() (*int32, bool)`

GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodCount

`func (o *DtoSubscriptionResponse) SetBillingPeriodCount(v int32)`

SetBillingPeriodCount sets BillingPeriodCount field to given value.

### HasBillingPeriodCount

`func (o *DtoSubscriptionResponse) HasBillingPeriodCount() bool`

HasBillingPeriodCount returns a boolean if a field has been set.

### GetCancelAt

`func (o *DtoSubscriptionResponse) GetCancelAt() string`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *DtoSubscriptionResponse) GetCancelAtOk() (*string, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *DtoSubscriptionResponse) SetCancelAt(v string)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *DtoSubscriptionResponse) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *DtoSubscriptionResponse) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *DtoSubscriptionResponse) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *DtoSubscriptionResponse) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *DtoSubscriptionResponse) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCancelledAt

`func (o *DtoSubscriptionResponse) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DtoSubscriptionResponse) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DtoSubscriptionResponse) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *DtoSubscriptionResponse) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetCollectionMethod

`func (o *DtoSubscriptionResponse) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *DtoSubscriptionResponse) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *DtoSubscriptionResponse) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *DtoSubscriptionResponse) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetCommitmentAmount

`func (o *DtoSubscriptionResponse) GetCommitmentAmount() string`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoSubscriptionResponse) GetCommitmentAmountOk() (*string, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoSubscriptionResponse) SetCommitmentAmount(v string)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoSubscriptionResponse) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCouponAssociations

`func (o *DtoSubscriptionResponse) GetCouponAssociations() []DtoCouponAssociationResponse`

GetCouponAssociations returns the CouponAssociations field if non-nil, zero value otherwise.

### GetCouponAssociationsOk

`func (o *DtoSubscriptionResponse) GetCouponAssociationsOk() (*[]DtoCouponAssociationResponse, bool)`

GetCouponAssociationsOk returns a tuple with the CouponAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAssociations

`func (o *DtoSubscriptionResponse) SetCouponAssociations(v []DtoCouponAssociationResponse)`

SetCouponAssociations sets CouponAssociations field to given value.

### HasCouponAssociations

`func (o *DtoSubscriptionResponse) HasCouponAssociations() bool`

HasCouponAssociations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoSubscriptionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoSubscriptionResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoSubscriptionResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoSubscriptionResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoSubscriptionResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditGrants

`func (o *DtoSubscriptionResponse) GetCreditGrants() []DtoCreditGrantResponse`

GetCreditGrants returns the CreditGrants field if non-nil, zero value otherwise.

### GetCreditGrantsOk

`func (o *DtoSubscriptionResponse) GetCreditGrantsOk() (*[]DtoCreditGrantResponse, bool)`

GetCreditGrantsOk returns a tuple with the CreditGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrants

`func (o *DtoSubscriptionResponse) SetCreditGrants(v []DtoCreditGrantResponse)`

SetCreditGrants sets CreditGrants field to given value.

### HasCreditGrants

`func (o *DtoSubscriptionResponse) HasCreditGrants() bool`

HasCreditGrants returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoSubscriptionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoSubscriptionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoSubscriptionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoSubscriptionResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *DtoSubscriptionResponse) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *DtoSubscriptionResponse) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *DtoSubscriptionResponse) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *DtoSubscriptionResponse) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *DtoSubscriptionResponse) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *DtoSubscriptionResponse) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *DtoSubscriptionResponse) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *DtoSubscriptionResponse) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoSubscriptionResponse) GetCustomer() DtoCustomerResponse`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoSubscriptionResponse) GetCustomerOk() (*DtoCustomerResponse, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoSubscriptionResponse) SetCustomer(v DtoCustomerResponse)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoSubscriptionResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoSubscriptionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoSubscriptionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoSubscriptionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoSubscriptionResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerTimezone

`func (o *DtoSubscriptionResponse) GetCustomerTimezone() string`

GetCustomerTimezone returns the CustomerTimezone field if non-nil, zero value otherwise.

### GetCustomerTimezoneOk

`func (o *DtoSubscriptionResponse) GetCustomerTimezoneOk() (*string, bool)`

GetCustomerTimezoneOk returns a tuple with the CustomerTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTimezone

`func (o *DtoSubscriptionResponse) SetCustomerTimezone(v string)`

SetCustomerTimezone sets CustomerTimezone field to given value.

### HasCustomerTimezone

`func (o *DtoSubscriptionResponse) HasCustomerTimezone() bool`

HasCustomerTimezone returns a boolean if a field has been set.

### GetEnableTrueUp

`func (o *DtoSubscriptionResponse) GetEnableTrueUp() bool`

GetEnableTrueUp returns the EnableTrueUp field if non-nil, zero value otherwise.

### GetEnableTrueUpOk

`func (o *DtoSubscriptionResponse) GetEnableTrueUpOk() (*bool, bool)`

GetEnableTrueUpOk returns a tuple with the EnableTrueUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrueUp

`func (o *DtoSubscriptionResponse) SetEnableTrueUp(v bool)`

SetEnableTrueUp sets EnableTrueUp field to given value.

### HasEnableTrueUp

`func (o *DtoSubscriptionResponse) HasEnableTrueUp() bool`

HasEnableTrueUp returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoSubscriptionResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoSubscriptionResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoSubscriptionResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoSubscriptionResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoSubscriptionResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoSubscriptionResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoSubscriptionResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoSubscriptionResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetGatewayPaymentMethodId

`func (o *DtoSubscriptionResponse) GetGatewayPaymentMethodId() string`

GetGatewayPaymentMethodId returns the GatewayPaymentMethodId field if non-nil, zero value otherwise.

### GetGatewayPaymentMethodIdOk

`func (o *DtoSubscriptionResponse) GetGatewayPaymentMethodIdOk() (*string, bool)`

GetGatewayPaymentMethodIdOk returns a tuple with the GatewayPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentMethodId

`func (o *DtoSubscriptionResponse) SetGatewayPaymentMethodId(v string)`

SetGatewayPaymentMethodId sets GatewayPaymentMethodId field to given value.

### HasGatewayPaymentMethodId

`func (o *DtoSubscriptionResponse) HasGatewayPaymentMethodId() bool`

HasGatewayPaymentMethodId returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoicingCustomerId

`func (o *DtoSubscriptionResponse) GetInvoicingCustomerId() string`

GetInvoicingCustomerId returns the InvoicingCustomerId field if non-nil, zero value otherwise.

### GetInvoicingCustomerIdOk

`func (o *DtoSubscriptionResponse) GetInvoicingCustomerIdOk() (*string, bool)`

GetInvoicingCustomerIdOk returns a tuple with the InvoicingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCustomerId

`func (o *DtoSubscriptionResponse) SetInvoicingCustomerId(v string)`

SetInvoicingCustomerId sets InvoicingCustomerId field to given value.

### HasInvoicingCustomerId

`func (o *DtoSubscriptionResponse) HasInvoicingCustomerId() bool`

HasInvoicingCustomerId returns a boolean if a field has been set.

### GetLatestInvoice

`func (o *DtoSubscriptionResponse) GetLatestInvoice() DtoInvoiceResponse`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *DtoSubscriptionResponse) GetLatestInvoiceOk() (*DtoInvoiceResponse, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *DtoSubscriptionResponse) SetLatestInvoice(v DtoInvoiceResponse)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *DtoSubscriptionResponse) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoSubscriptionResponse) GetLineItems() []SubscriptionSubscriptionLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoSubscriptionResponse) GetLineItemsOk() (*[]SubscriptionSubscriptionLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoSubscriptionResponse) SetLineItems(v []SubscriptionSubscriptionLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoSubscriptionResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoSubscriptionResponse) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoSubscriptionResponse) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoSubscriptionResponse) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoSubscriptionResponse) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoSubscriptionResponse) GetOverageFactor() string`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoSubscriptionResponse) GetOverageFactorOk() (*string, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoSubscriptionResponse) SetOverageFactor(v string)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoSubscriptionResponse) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

### GetPauseStatus

`func (o *DtoSubscriptionResponse) GetPauseStatus() TypesPauseStatus`

GetPauseStatus returns the PauseStatus field if non-nil, zero value otherwise.

### GetPauseStatusOk

`func (o *DtoSubscriptionResponse) GetPauseStatusOk() (*TypesPauseStatus, bool)`

GetPauseStatusOk returns a tuple with the PauseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStatus

`func (o *DtoSubscriptionResponse) SetPauseStatus(v TypesPauseStatus)`

SetPauseStatus sets PauseStatus field to given value.

### HasPauseStatus

`func (o *DtoSubscriptionResponse) HasPauseStatus() bool`

HasPauseStatus returns a boolean if a field has been set.

### GetPauses

`func (o *DtoSubscriptionResponse) GetPauses() []SubscriptionSubscriptionPause`

GetPauses returns the Pauses field if non-nil, zero value otherwise.

### GetPausesOk

`func (o *DtoSubscriptionResponse) GetPausesOk() (*[]SubscriptionSubscriptionPause, bool)`

GetPausesOk returns a tuple with the Pauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauses

`func (o *DtoSubscriptionResponse) SetPauses(v []SubscriptionSubscriptionPause)`

SetPauses sets Pauses field to given value.

### HasPauses

`func (o *DtoSubscriptionResponse) HasPauses() bool`

HasPauses returns a boolean if a field has been set.

### GetPaymentBehavior

`func (o *DtoSubscriptionResponse) GetPaymentBehavior() string`

GetPaymentBehavior returns the PaymentBehavior field if non-nil, zero value otherwise.

### GetPaymentBehaviorOk

`func (o *DtoSubscriptionResponse) GetPaymentBehaviorOk() (*string, bool)`

GetPaymentBehaviorOk returns a tuple with the PaymentBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBehavior

`func (o *DtoSubscriptionResponse) SetPaymentBehavior(v string)`

SetPaymentBehavior sets PaymentBehavior field to given value.

### HasPaymentBehavior

`func (o *DtoSubscriptionResponse) HasPaymentBehavior() bool`

HasPaymentBehavior returns a boolean if a field has been set.

### GetPhases

`func (o *DtoSubscriptionResponse) GetPhases() []DtoSubscriptionPhaseResponse`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *DtoSubscriptionResponse) GetPhasesOk() (*[]DtoSubscriptionPhaseResponse, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *DtoSubscriptionResponse) SetPhases(v []DtoSubscriptionPhaseResponse)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *DtoSubscriptionResponse) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetPlan

`func (o *DtoSubscriptionResponse) GetPlan() DtoPlanResponse`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *DtoSubscriptionResponse) GetPlanOk() (*DtoPlanResponse, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *DtoSubscriptionResponse) SetPlan(v DtoPlanResponse)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *DtoSubscriptionResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoSubscriptionResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoSubscriptionResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoSubscriptionResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoSubscriptionResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *DtoSubscriptionResponse) GetProrationBehavior() TypesProrationBehavior`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *DtoSubscriptionResponse) GetProrationBehaviorOk() (*TypesProrationBehavior, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *DtoSubscriptionResponse) SetProrationBehavior(v TypesProrationBehavior)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *DtoSubscriptionResponse) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoSubscriptionResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoSubscriptionResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoSubscriptionResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoSubscriptionResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *DtoSubscriptionResponse) GetSubscriptionStatus() TypesSubscriptionStatus`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *DtoSubscriptionResponse) GetSubscriptionStatusOk() (*TypesSubscriptionStatus, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *DtoSubscriptionResponse) SetSubscriptionStatus(v TypesSubscriptionStatus)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *DtoSubscriptionResponse) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoSubscriptionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoSubscriptionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoSubscriptionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoSubscriptionResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrialEnd

`func (o *DtoSubscriptionResponse) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *DtoSubscriptionResponse) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *DtoSubscriptionResponse) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *DtoSubscriptionResponse) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetTrialStart

`func (o *DtoSubscriptionResponse) GetTrialStart() string`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *DtoSubscriptionResponse) GetTrialStartOk() (*string, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *DtoSubscriptionResponse) SetTrialStart(v string)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *DtoSubscriptionResponse) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSubscriptionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSubscriptionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSubscriptionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSubscriptionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoSubscriptionResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoSubscriptionResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoSubscriptionResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoSubscriptionResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *DtoSubscriptionResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DtoSubscriptionResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DtoSubscriptionResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DtoSubscriptionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


