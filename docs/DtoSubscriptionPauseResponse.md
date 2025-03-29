# DtoSubscriptionPauseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** | EnvironmentID is the environment identifier for the pause | [optional] 
**Id** | Pointer to **string** | ID is the unique identifier for the subscription pause | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OriginalPeriodEnd** | Pointer to **string** | OriginalPeriodEnd is the end of the billing period when the pause was created | [optional] 
**OriginalPeriodStart** | Pointer to **string** | OriginalPeriodStart is the start of the billing period when the pause was created | [optional] 
**PauseEnd** | Pointer to **string** | PauseEnd is when the pause will end (null for indefinite) | [optional] 
**PauseMode** | Pointer to [**TypesPauseMode**](TypesPauseMode.md) |  | [optional] 
**PauseStart** | Pointer to **string** | PauseStart is when the pause actually started | [optional] 
**PauseStatus** | Pointer to [**TypesPauseStatus**](TypesPauseStatus.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason is the reason for pausing | [optional] 
**ResumeMode** | Pointer to [**TypesResumeMode**](TypesResumeMode.md) |  | [optional] 
**ResumedAt** | Pointer to **string** | ResumedAt is when the pause was actually ended (if manually resumed) | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionID is the identifier for the subscription | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoSubscriptionPauseResponse

`func NewDtoSubscriptionPauseResponse() *DtoSubscriptionPauseResponse`

NewDtoSubscriptionPauseResponse instantiates a new DtoSubscriptionPauseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionPauseResponseWithDefaults

`func NewDtoSubscriptionPauseResponseWithDefaults() *DtoSubscriptionPauseResponse`

NewDtoSubscriptionPauseResponseWithDefaults instantiates a new DtoSubscriptionPauseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DtoSubscriptionPauseResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionPauseResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionPauseResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionPauseResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoSubscriptionPauseResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoSubscriptionPauseResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoSubscriptionPauseResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoSubscriptionPauseResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoSubscriptionPauseResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoSubscriptionPauseResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoSubscriptionPauseResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoSubscriptionPauseResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionPauseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionPauseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionPauseResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionPauseResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionPauseResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionPauseResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionPauseResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionPauseResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginalPeriodEnd

`func (o *DtoSubscriptionPauseResponse) GetOriginalPeriodEnd() string`

GetOriginalPeriodEnd returns the OriginalPeriodEnd field if non-nil, zero value otherwise.

### GetOriginalPeriodEndOk

`func (o *DtoSubscriptionPauseResponse) GetOriginalPeriodEndOk() (*string, bool)`

GetOriginalPeriodEndOk returns a tuple with the OriginalPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPeriodEnd

`func (o *DtoSubscriptionPauseResponse) SetOriginalPeriodEnd(v string)`

SetOriginalPeriodEnd sets OriginalPeriodEnd field to given value.

### HasOriginalPeriodEnd

`func (o *DtoSubscriptionPauseResponse) HasOriginalPeriodEnd() bool`

HasOriginalPeriodEnd returns a boolean if a field has been set.

### GetOriginalPeriodStart

`func (o *DtoSubscriptionPauseResponse) GetOriginalPeriodStart() string`

GetOriginalPeriodStart returns the OriginalPeriodStart field if non-nil, zero value otherwise.

### GetOriginalPeriodStartOk

`func (o *DtoSubscriptionPauseResponse) GetOriginalPeriodStartOk() (*string, bool)`

GetOriginalPeriodStartOk returns a tuple with the OriginalPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPeriodStart

`func (o *DtoSubscriptionPauseResponse) SetOriginalPeriodStart(v string)`

SetOriginalPeriodStart sets OriginalPeriodStart field to given value.

### HasOriginalPeriodStart

`func (o *DtoSubscriptionPauseResponse) HasOriginalPeriodStart() bool`

HasOriginalPeriodStart returns a boolean if a field has been set.

### GetPauseEnd

`func (o *DtoSubscriptionPauseResponse) GetPauseEnd() string`

GetPauseEnd returns the PauseEnd field if non-nil, zero value otherwise.

### GetPauseEndOk

`func (o *DtoSubscriptionPauseResponse) GetPauseEndOk() (*string, bool)`

GetPauseEndOk returns a tuple with the PauseEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseEnd

`func (o *DtoSubscriptionPauseResponse) SetPauseEnd(v string)`

SetPauseEnd sets PauseEnd field to given value.

### HasPauseEnd

`func (o *DtoSubscriptionPauseResponse) HasPauseEnd() bool`

HasPauseEnd returns a boolean if a field has been set.

### GetPauseMode

`func (o *DtoSubscriptionPauseResponse) GetPauseMode() TypesPauseMode`

GetPauseMode returns the PauseMode field if non-nil, zero value otherwise.

### GetPauseModeOk

`func (o *DtoSubscriptionPauseResponse) GetPauseModeOk() (*TypesPauseMode, bool)`

GetPauseModeOk returns a tuple with the PauseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseMode

`func (o *DtoSubscriptionPauseResponse) SetPauseMode(v TypesPauseMode)`

SetPauseMode sets PauseMode field to given value.

### HasPauseMode

`func (o *DtoSubscriptionPauseResponse) HasPauseMode() bool`

HasPauseMode returns a boolean if a field has been set.

### GetPauseStart

`func (o *DtoSubscriptionPauseResponse) GetPauseStart() string`

GetPauseStart returns the PauseStart field if non-nil, zero value otherwise.

### GetPauseStartOk

`func (o *DtoSubscriptionPauseResponse) GetPauseStartOk() (*string, bool)`

GetPauseStartOk returns a tuple with the PauseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStart

`func (o *DtoSubscriptionPauseResponse) SetPauseStart(v string)`

SetPauseStart sets PauseStart field to given value.

### HasPauseStart

`func (o *DtoSubscriptionPauseResponse) HasPauseStart() bool`

HasPauseStart returns a boolean if a field has been set.

### GetPauseStatus

`func (o *DtoSubscriptionPauseResponse) GetPauseStatus() TypesPauseStatus`

GetPauseStatus returns the PauseStatus field if non-nil, zero value otherwise.

### GetPauseStatusOk

`func (o *DtoSubscriptionPauseResponse) GetPauseStatusOk() (*TypesPauseStatus, bool)`

GetPauseStatusOk returns a tuple with the PauseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStatus

`func (o *DtoSubscriptionPauseResponse) SetPauseStatus(v TypesPauseStatus)`

SetPauseStatus sets PauseStatus field to given value.

### HasPauseStatus

`func (o *DtoSubscriptionPauseResponse) HasPauseStatus() bool`

HasPauseStatus returns a boolean if a field has been set.

### GetReason

`func (o *DtoSubscriptionPauseResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoSubscriptionPauseResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoSubscriptionPauseResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoSubscriptionPauseResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResumeMode

`func (o *DtoSubscriptionPauseResponse) GetResumeMode() TypesResumeMode`

GetResumeMode returns the ResumeMode field if non-nil, zero value otherwise.

### GetResumeModeOk

`func (o *DtoSubscriptionPauseResponse) GetResumeModeOk() (*TypesResumeMode, bool)`

GetResumeModeOk returns a tuple with the ResumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeMode

`func (o *DtoSubscriptionPauseResponse) SetResumeMode(v TypesResumeMode)`

SetResumeMode sets ResumeMode field to given value.

### HasResumeMode

`func (o *DtoSubscriptionPauseResponse) HasResumeMode() bool`

HasResumeMode returns a boolean if a field has been set.

### GetResumedAt

`func (o *DtoSubscriptionPauseResponse) GetResumedAt() string`

GetResumedAt returns the ResumedAt field if non-nil, zero value otherwise.

### GetResumedAtOk

`func (o *DtoSubscriptionPauseResponse) GetResumedAtOk() (*string, bool)`

GetResumedAtOk returns a tuple with the ResumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumedAt

`func (o *DtoSubscriptionPauseResponse) SetResumedAt(v string)`

SetResumedAt sets ResumedAt field to given value.

### HasResumedAt

`func (o *DtoSubscriptionPauseResponse) HasResumedAt() bool`

HasResumedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionPauseResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionPauseResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionPauseResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionPauseResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoSubscriptionPauseResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoSubscriptionPauseResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoSubscriptionPauseResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoSubscriptionPauseResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoSubscriptionPauseResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoSubscriptionPauseResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoSubscriptionPauseResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoSubscriptionPauseResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSubscriptionPauseResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSubscriptionPauseResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSubscriptionPauseResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSubscriptionPauseResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoSubscriptionPauseResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoSubscriptionPauseResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoSubscriptionPauseResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoSubscriptionPauseResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


