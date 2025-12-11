# DtoCreditGrantApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationReason** | Pointer to [**TypesCreditGrantApplicationReason**](TypesCreditGrantApplicationReason.md) |  | [optional] 
**ApplicationStatus** | Pointer to [**TypesApplicationStatus**](TypesApplicationStatus.md) |  | [optional] 
**AppliedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditGrantId** | Pointer to **string** |  | [optional] 
**Credits** | Pointer to **float32** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**ScheduledFor** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SubscriptionStatusAtApplication** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreditGrantApplicationResponse

`func NewDtoCreditGrantApplicationResponse() *DtoCreditGrantApplicationResponse`

NewDtoCreditGrantApplicationResponse instantiates a new DtoCreditGrantApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreditGrantApplicationResponseWithDefaults

`func NewDtoCreditGrantApplicationResponseWithDefaults() *DtoCreditGrantApplicationResponse`

NewDtoCreditGrantApplicationResponseWithDefaults instantiates a new DtoCreditGrantApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationReason

`func (o *DtoCreditGrantApplicationResponse) GetApplicationReason() TypesCreditGrantApplicationReason`

GetApplicationReason returns the ApplicationReason field if non-nil, zero value otherwise.

### GetApplicationReasonOk

`func (o *DtoCreditGrantApplicationResponse) GetApplicationReasonOk() (*TypesCreditGrantApplicationReason, bool)`

GetApplicationReasonOk returns a tuple with the ApplicationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationReason

`func (o *DtoCreditGrantApplicationResponse) SetApplicationReason(v TypesCreditGrantApplicationReason)`

SetApplicationReason sets ApplicationReason field to given value.

### HasApplicationReason

`func (o *DtoCreditGrantApplicationResponse) HasApplicationReason() bool`

HasApplicationReason returns a boolean if a field has been set.

### GetApplicationStatus

`func (o *DtoCreditGrantApplicationResponse) GetApplicationStatus() TypesApplicationStatus`

GetApplicationStatus returns the ApplicationStatus field if non-nil, zero value otherwise.

### GetApplicationStatusOk

`func (o *DtoCreditGrantApplicationResponse) GetApplicationStatusOk() (*TypesApplicationStatus, bool)`

GetApplicationStatusOk returns a tuple with the ApplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationStatus

`func (o *DtoCreditGrantApplicationResponse) SetApplicationStatus(v TypesApplicationStatus)`

SetApplicationStatus sets ApplicationStatus field to given value.

### HasApplicationStatus

`func (o *DtoCreditGrantApplicationResponse) HasApplicationStatus() bool`

HasApplicationStatus returns a boolean if a field has been set.

### GetAppliedAt

`func (o *DtoCreditGrantApplicationResponse) GetAppliedAt() string`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *DtoCreditGrantApplicationResponse) GetAppliedAtOk() (*string, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *DtoCreditGrantApplicationResponse) SetAppliedAt(v string)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *DtoCreditGrantApplicationResponse) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoCreditGrantApplicationResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoCreditGrantApplicationResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoCreditGrantApplicationResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoCreditGrantApplicationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoCreditGrantApplicationResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoCreditGrantApplicationResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoCreditGrantApplicationResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoCreditGrantApplicationResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditGrantId

`func (o *DtoCreditGrantApplicationResponse) GetCreditGrantId() string`

GetCreditGrantId returns the CreditGrantId field if non-nil, zero value otherwise.

### GetCreditGrantIdOk

`func (o *DtoCreditGrantApplicationResponse) GetCreditGrantIdOk() (*string, bool)`

GetCreditGrantIdOk returns a tuple with the CreditGrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrantId

`func (o *DtoCreditGrantApplicationResponse) SetCreditGrantId(v string)`

SetCreditGrantId sets CreditGrantId field to given value.

### HasCreditGrantId

`func (o *DtoCreditGrantApplicationResponse) HasCreditGrantId() bool`

HasCreditGrantId returns a boolean if a field has been set.

### GetCredits

`func (o *DtoCreditGrantApplicationResponse) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *DtoCreditGrantApplicationResponse) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *DtoCreditGrantApplicationResponse) SetCredits(v float32)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *DtoCreditGrantApplicationResponse) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCreditGrantApplicationResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCreditGrantApplicationResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCreditGrantApplicationResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCreditGrantApplicationResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFailureReason

`func (o *DtoCreditGrantApplicationResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DtoCreditGrantApplicationResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DtoCreditGrantApplicationResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DtoCreditGrantApplicationResponse) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetId

`func (o *DtoCreditGrantApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoCreditGrantApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoCreditGrantApplicationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoCreditGrantApplicationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoCreditGrantApplicationResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoCreditGrantApplicationResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoCreditGrantApplicationResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoCreditGrantApplicationResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreditGrantApplicationResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreditGrantApplicationResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreditGrantApplicationResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreditGrantApplicationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DtoCreditGrantApplicationResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DtoCreditGrantApplicationResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DtoCreditGrantApplicationResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DtoCreditGrantApplicationResponse) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DtoCreditGrantApplicationResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DtoCreditGrantApplicationResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DtoCreditGrantApplicationResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DtoCreditGrantApplicationResponse) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetRetryCount

`func (o *DtoCreditGrantApplicationResponse) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *DtoCreditGrantApplicationResponse) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *DtoCreditGrantApplicationResponse) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *DtoCreditGrantApplicationResponse) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetScheduledFor

`func (o *DtoCreditGrantApplicationResponse) GetScheduledFor() string`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *DtoCreditGrantApplicationResponse) GetScheduledForOk() (*string, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *DtoCreditGrantApplicationResponse) SetScheduledFor(v string)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *DtoCreditGrantApplicationResponse) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCreditGrantApplicationResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCreditGrantApplicationResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCreditGrantApplicationResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCreditGrantApplicationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCreditGrantApplicationResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCreditGrantApplicationResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCreditGrantApplicationResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCreditGrantApplicationResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionStatusAtApplication

`func (o *DtoCreditGrantApplicationResponse) GetSubscriptionStatusAtApplication() TypesSubscriptionStatus`

GetSubscriptionStatusAtApplication returns the SubscriptionStatusAtApplication field if non-nil, zero value otherwise.

### GetSubscriptionStatusAtApplicationOk

`func (o *DtoCreditGrantApplicationResponse) GetSubscriptionStatusAtApplicationOk() (*TypesSubscriptionStatus, bool)`

GetSubscriptionStatusAtApplicationOk returns a tuple with the SubscriptionStatusAtApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusAtApplication

`func (o *DtoCreditGrantApplicationResponse) SetSubscriptionStatusAtApplication(v TypesSubscriptionStatus)`

SetSubscriptionStatusAtApplication sets SubscriptionStatusAtApplication field to given value.

### HasSubscriptionStatusAtApplication

`func (o *DtoCreditGrantApplicationResponse) HasSubscriptionStatusAtApplication() bool`

HasSubscriptionStatusAtApplication returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoCreditGrantApplicationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoCreditGrantApplicationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoCreditGrantApplicationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoCreditGrantApplicationResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoCreditGrantApplicationResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoCreditGrantApplicationResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoCreditGrantApplicationResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoCreditGrantApplicationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoCreditGrantApplicationResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoCreditGrantApplicationResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoCreditGrantApplicationResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoCreditGrantApplicationResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


