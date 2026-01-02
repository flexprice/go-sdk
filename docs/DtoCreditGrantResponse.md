# DtoCreditGrantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | Pointer to [**TypesCreditGrantCadence**](TypesCreditGrantCadence.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditGrantAnchor** | Pointer to **string** |  | [optional] 
**Credits** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**ExpirationDuration** | Pointer to **int32** |  | [optional] 
**ExpirationDurationUnit** | Pointer to [**TypesCreditGrantExpiryDurationUnit**](TypesCreditGrantExpiryDurationUnit.md) |  | [optional] 
**ExpirationType** | Pointer to [**TypesCreditGrantExpiryType**](TypesCreditGrantExpiryType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Period** | Pointer to [**TypesCreditGrantPeriod**](TypesCreditGrantPeriod.md) |  | [optional] 
**PeriodCount** | Pointer to **int32** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Scope** | Pointer to [**TypesCreditGrantScope**](TypesCreditGrantScope.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreditGrantResponse

`func NewDtoCreditGrantResponse() *DtoCreditGrantResponse`

NewDtoCreditGrantResponse instantiates a new DtoCreditGrantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreditGrantResponseWithDefaults

`func NewDtoCreditGrantResponseWithDefaults() *DtoCreditGrantResponse`

NewDtoCreditGrantResponseWithDefaults instantiates a new DtoCreditGrantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *DtoCreditGrantResponse) GetCadence() TypesCreditGrantCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DtoCreditGrantResponse) GetCadenceOk() (*TypesCreditGrantCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DtoCreditGrantResponse) SetCadence(v TypesCreditGrantCadence)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *DtoCreditGrantResponse) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoCreditGrantResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoCreditGrantResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoCreditGrantResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoCreditGrantResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoCreditGrantResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoCreditGrantResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoCreditGrantResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoCreditGrantResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditGrantAnchor

`func (o *DtoCreditGrantResponse) GetCreditGrantAnchor() string`

GetCreditGrantAnchor returns the CreditGrantAnchor field if non-nil, zero value otherwise.

### GetCreditGrantAnchorOk

`func (o *DtoCreditGrantResponse) GetCreditGrantAnchorOk() (*string, bool)`

GetCreditGrantAnchorOk returns a tuple with the CreditGrantAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrantAnchor

`func (o *DtoCreditGrantResponse) SetCreditGrantAnchor(v string)`

SetCreditGrantAnchor sets CreditGrantAnchor field to given value.

### HasCreditGrantAnchor

`func (o *DtoCreditGrantResponse) HasCreditGrantAnchor() bool`

HasCreditGrantAnchor returns a boolean if a field has been set.

### GetCredits

`func (o *DtoCreditGrantResponse) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *DtoCreditGrantResponse) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *DtoCreditGrantResponse) SetCredits(v string)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *DtoCreditGrantResponse) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoCreditGrantResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoCreditGrantResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoCreditGrantResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoCreditGrantResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCreditGrantResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCreditGrantResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCreditGrantResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCreditGrantResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetExpirationDuration

`func (o *DtoCreditGrantResponse) GetExpirationDuration() int32`

GetExpirationDuration returns the ExpirationDuration field if non-nil, zero value otherwise.

### GetExpirationDurationOk

`func (o *DtoCreditGrantResponse) GetExpirationDurationOk() (*int32, bool)`

GetExpirationDurationOk returns a tuple with the ExpirationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDuration

`func (o *DtoCreditGrantResponse) SetExpirationDuration(v int32)`

SetExpirationDuration sets ExpirationDuration field to given value.

### HasExpirationDuration

`func (o *DtoCreditGrantResponse) HasExpirationDuration() bool`

HasExpirationDuration returns a boolean if a field has been set.

### GetExpirationDurationUnit

`func (o *DtoCreditGrantResponse) GetExpirationDurationUnit() TypesCreditGrantExpiryDurationUnit`

GetExpirationDurationUnit returns the ExpirationDurationUnit field if non-nil, zero value otherwise.

### GetExpirationDurationUnitOk

`func (o *DtoCreditGrantResponse) GetExpirationDurationUnitOk() (*TypesCreditGrantExpiryDurationUnit, bool)`

GetExpirationDurationUnitOk returns a tuple with the ExpirationDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDurationUnit

`func (o *DtoCreditGrantResponse) SetExpirationDurationUnit(v TypesCreditGrantExpiryDurationUnit)`

SetExpirationDurationUnit sets ExpirationDurationUnit field to given value.

### HasExpirationDurationUnit

`func (o *DtoCreditGrantResponse) HasExpirationDurationUnit() bool`

HasExpirationDurationUnit returns a boolean if a field has been set.

### GetExpirationType

`func (o *DtoCreditGrantResponse) GetExpirationType() TypesCreditGrantExpiryType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *DtoCreditGrantResponse) GetExpirationTypeOk() (*TypesCreditGrantExpiryType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *DtoCreditGrantResponse) SetExpirationType(v TypesCreditGrantExpiryType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *DtoCreditGrantResponse) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetId

`func (o *DtoCreditGrantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoCreditGrantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoCreditGrantResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoCreditGrantResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreditGrantResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreditGrantResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreditGrantResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreditGrantResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreditGrantResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreditGrantResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreditGrantResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoCreditGrantResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeriod

`func (o *DtoCreditGrantResponse) GetPeriod() TypesCreditGrantPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DtoCreditGrantResponse) GetPeriodOk() (*TypesCreditGrantPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DtoCreditGrantResponse) SetPeriod(v TypesCreditGrantPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DtoCreditGrantResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodCount

`func (o *DtoCreditGrantResponse) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *DtoCreditGrantResponse) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *DtoCreditGrantResponse) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.

### HasPeriodCount

`func (o *DtoCreditGrantResponse) HasPeriodCount() bool`

HasPeriodCount returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreditGrantResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreditGrantResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreditGrantResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreditGrantResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriority

`func (o *DtoCreditGrantResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoCreditGrantResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoCreditGrantResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoCreditGrantResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetScope

`func (o *DtoCreditGrantResponse) GetScope() TypesCreditGrantScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DtoCreditGrantResponse) GetScopeOk() (*TypesCreditGrantScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DtoCreditGrantResponse) SetScope(v TypesCreditGrantScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DtoCreditGrantResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoCreditGrantResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoCreditGrantResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoCreditGrantResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoCreditGrantResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCreditGrantResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCreditGrantResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCreditGrantResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCreditGrantResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCreditGrantResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCreditGrantResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCreditGrantResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCreditGrantResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoCreditGrantResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoCreditGrantResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoCreditGrantResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoCreditGrantResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoCreditGrantResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoCreditGrantResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoCreditGrantResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoCreditGrantResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoCreditGrantResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoCreditGrantResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoCreditGrantResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoCreditGrantResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


