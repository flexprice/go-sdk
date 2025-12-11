# DtoUpdatePlanCreditGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | [**TypesCreditGrantCadence**](TypesCreditGrantCadence.md) |  | 
**Credits** | **string** |  | 
**ExpirationDuration** | Pointer to **int32** |  | [optional] 
**ExpirationDurationUnit** | Pointer to [**TypesCreditGrantExpiryDurationUnit**](TypesCreditGrantExpiryDurationUnit.md) |  | [optional] 
**ExpirationType** | Pointer to [**TypesCreditGrantExpiryType**](TypesCreditGrantExpiryType.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the credit grant to update (present if the credit grant is being updated) | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**Period** | Pointer to [**TypesCreditGrantPeriod**](TypesCreditGrantPeriod.md) |  | [optional] 
**PeriodCount** | Pointer to **int32** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Scope** | [**TypesCreditGrantScope**](TypesCreditGrantScope.md) |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUpdatePlanCreditGrantRequest

`func NewDtoUpdatePlanCreditGrantRequest(cadence TypesCreditGrantCadence, credits string, name string, scope TypesCreditGrantScope, ) *DtoUpdatePlanCreditGrantRequest`

NewDtoUpdatePlanCreditGrantRequest instantiates a new DtoUpdatePlanCreditGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePlanCreditGrantRequestWithDefaults

`func NewDtoUpdatePlanCreditGrantRequestWithDefaults() *DtoUpdatePlanCreditGrantRequest`

NewDtoUpdatePlanCreditGrantRequestWithDefaults instantiates a new DtoUpdatePlanCreditGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *DtoUpdatePlanCreditGrantRequest) GetCadence() TypesCreditGrantCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetCadenceOk() (*TypesCreditGrantCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DtoUpdatePlanCreditGrantRequest) SetCadence(v TypesCreditGrantCadence)`

SetCadence sets Cadence field to given value.


### GetCredits

`func (o *DtoUpdatePlanCreditGrantRequest) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *DtoUpdatePlanCreditGrantRequest) SetCredits(v string)`

SetCredits sets Credits field to given value.


### GetExpirationDuration

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationDuration() int32`

GetExpirationDuration returns the ExpirationDuration field if non-nil, zero value otherwise.

### GetExpirationDurationOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationDurationOk() (*int32, bool)`

GetExpirationDurationOk returns a tuple with the ExpirationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDuration

`func (o *DtoUpdatePlanCreditGrantRequest) SetExpirationDuration(v int32)`

SetExpirationDuration sets ExpirationDuration field to given value.

### HasExpirationDuration

`func (o *DtoUpdatePlanCreditGrantRequest) HasExpirationDuration() bool`

HasExpirationDuration returns a boolean if a field has been set.

### GetExpirationDurationUnit

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationDurationUnit() TypesCreditGrantExpiryDurationUnit`

GetExpirationDurationUnit returns the ExpirationDurationUnit field if non-nil, zero value otherwise.

### GetExpirationDurationUnitOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationDurationUnitOk() (*TypesCreditGrantExpiryDurationUnit, bool)`

GetExpirationDurationUnitOk returns a tuple with the ExpirationDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDurationUnit

`func (o *DtoUpdatePlanCreditGrantRequest) SetExpirationDurationUnit(v TypesCreditGrantExpiryDurationUnit)`

SetExpirationDurationUnit sets ExpirationDurationUnit field to given value.

### HasExpirationDurationUnit

`func (o *DtoUpdatePlanCreditGrantRequest) HasExpirationDurationUnit() bool`

HasExpirationDurationUnit returns a boolean if a field has been set.

### GetExpirationType

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationType() TypesCreditGrantExpiryType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetExpirationTypeOk() (*TypesCreditGrantExpiryType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *DtoUpdatePlanCreditGrantRequest) SetExpirationType(v TypesCreditGrantExpiryType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *DtoUpdatePlanCreditGrantRequest) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetId

`func (o *DtoUpdatePlanCreditGrantRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoUpdatePlanCreditGrantRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoUpdatePlanCreditGrantRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdatePlanCreditGrantRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdatePlanCreditGrantRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdatePlanCreditGrantRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdatePlanCreditGrantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdatePlanCreditGrantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPeriod

`func (o *DtoUpdatePlanCreditGrantRequest) GetPeriod() TypesCreditGrantPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetPeriodOk() (*TypesCreditGrantPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DtoUpdatePlanCreditGrantRequest) SetPeriod(v TypesCreditGrantPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DtoUpdatePlanCreditGrantRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodCount

`func (o *DtoUpdatePlanCreditGrantRequest) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *DtoUpdatePlanCreditGrantRequest) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.

### HasPeriodCount

`func (o *DtoUpdatePlanCreditGrantRequest) HasPeriodCount() bool`

HasPeriodCount returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoUpdatePlanCreditGrantRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoUpdatePlanCreditGrantRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoUpdatePlanCreditGrantRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriority

`func (o *DtoUpdatePlanCreditGrantRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoUpdatePlanCreditGrantRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoUpdatePlanCreditGrantRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetScope

`func (o *DtoUpdatePlanCreditGrantRequest) GetScope() TypesCreditGrantScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetScopeOk() (*TypesCreditGrantScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DtoUpdatePlanCreditGrantRequest) SetScope(v TypesCreditGrantScope)`

SetScope sets Scope field to given value.


### GetSubscriptionId

`func (o *DtoUpdatePlanCreditGrantRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoUpdatePlanCreditGrantRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoUpdatePlanCreditGrantRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoUpdatePlanCreditGrantRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


