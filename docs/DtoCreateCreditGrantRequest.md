# DtoCreateCreditGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | [**TypesCreditGrantCadence**](TypesCreditGrantCadence.md) |  | 
**Credits** | **string** |  | 
**ExpirationDuration** | Pointer to **int32** |  | [optional] 
**ExpirationDurationUnit** | Pointer to [**TypesCreditGrantExpiryDurationUnit**](TypesCreditGrantExpiryDurationUnit.md) |  | [optional] 
**ExpirationType** | Pointer to [**TypesCreditGrantExpiryType**](TypesCreditGrantExpiryType.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**Period** | Pointer to [**TypesCreditGrantPeriod**](TypesCreditGrantPeriod.md) |  | [optional] 
**PeriodCount** | Pointer to **int32** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Scope** | [**TypesCreditGrantScope**](TypesCreditGrantScope.md) |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateCreditGrantRequest

`func NewDtoCreateCreditGrantRequest(cadence TypesCreditGrantCadence, credits string, name string, scope TypesCreditGrantScope, ) *DtoCreateCreditGrantRequest`

NewDtoCreateCreditGrantRequest instantiates a new DtoCreateCreditGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateCreditGrantRequestWithDefaults

`func NewDtoCreateCreditGrantRequestWithDefaults() *DtoCreateCreditGrantRequest`

NewDtoCreateCreditGrantRequestWithDefaults instantiates a new DtoCreateCreditGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *DtoCreateCreditGrantRequest) GetCadence() TypesCreditGrantCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DtoCreateCreditGrantRequest) GetCadenceOk() (*TypesCreditGrantCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DtoCreateCreditGrantRequest) SetCadence(v TypesCreditGrantCadence)`

SetCadence sets Cadence field to given value.


### GetCredits

`func (o *DtoCreateCreditGrantRequest) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *DtoCreateCreditGrantRequest) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *DtoCreateCreditGrantRequest) SetCredits(v string)`

SetCredits sets Credits field to given value.


### GetExpirationDuration

`func (o *DtoCreateCreditGrantRequest) GetExpirationDuration() int32`

GetExpirationDuration returns the ExpirationDuration field if non-nil, zero value otherwise.

### GetExpirationDurationOk

`func (o *DtoCreateCreditGrantRequest) GetExpirationDurationOk() (*int32, bool)`

GetExpirationDurationOk returns a tuple with the ExpirationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDuration

`func (o *DtoCreateCreditGrantRequest) SetExpirationDuration(v int32)`

SetExpirationDuration sets ExpirationDuration field to given value.

### HasExpirationDuration

`func (o *DtoCreateCreditGrantRequest) HasExpirationDuration() bool`

HasExpirationDuration returns a boolean if a field has been set.

### GetExpirationDurationUnit

`func (o *DtoCreateCreditGrantRequest) GetExpirationDurationUnit() TypesCreditGrantExpiryDurationUnit`

GetExpirationDurationUnit returns the ExpirationDurationUnit field if non-nil, zero value otherwise.

### GetExpirationDurationUnitOk

`func (o *DtoCreateCreditGrantRequest) GetExpirationDurationUnitOk() (*TypesCreditGrantExpiryDurationUnit, bool)`

GetExpirationDurationUnitOk returns a tuple with the ExpirationDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDurationUnit

`func (o *DtoCreateCreditGrantRequest) SetExpirationDurationUnit(v TypesCreditGrantExpiryDurationUnit)`

SetExpirationDurationUnit sets ExpirationDurationUnit field to given value.

### HasExpirationDurationUnit

`func (o *DtoCreateCreditGrantRequest) HasExpirationDurationUnit() bool`

HasExpirationDurationUnit returns a boolean if a field has been set.

### GetExpirationType

`func (o *DtoCreateCreditGrantRequest) GetExpirationType() TypesCreditGrantExpiryType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *DtoCreateCreditGrantRequest) GetExpirationTypeOk() (*TypesCreditGrantExpiryType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *DtoCreateCreditGrantRequest) SetExpirationType(v TypesCreditGrantExpiryType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *DtoCreateCreditGrantRequest) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateCreditGrantRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateCreditGrantRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateCreditGrantRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateCreditGrantRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateCreditGrantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateCreditGrantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateCreditGrantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPeriod

`func (o *DtoCreateCreditGrantRequest) GetPeriod() TypesCreditGrantPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DtoCreateCreditGrantRequest) GetPeriodOk() (*TypesCreditGrantPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DtoCreateCreditGrantRequest) SetPeriod(v TypesCreditGrantPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DtoCreateCreditGrantRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodCount

`func (o *DtoCreateCreditGrantRequest) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *DtoCreateCreditGrantRequest) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *DtoCreateCreditGrantRequest) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.

### HasPeriodCount

`func (o *DtoCreateCreditGrantRequest) HasPeriodCount() bool`

HasPeriodCount returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreateCreditGrantRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreateCreditGrantRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreateCreditGrantRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreateCreditGrantRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriority

`func (o *DtoCreateCreditGrantRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoCreateCreditGrantRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoCreateCreditGrantRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoCreateCreditGrantRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetScope

`func (o *DtoCreateCreditGrantRequest) GetScope() TypesCreditGrantScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DtoCreateCreditGrantRequest) GetScopeOk() (*TypesCreditGrantScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DtoCreateCreditGrantRequest) SetScope(v TypesCreditGrantScope)`

SetScope sets Scope field to given value.


### GetSubscriptionId

`func (o *DtoCreateCreditGrantRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCreateCreditGrantRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCreateCreditGrantRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCreateCreditGrantRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


