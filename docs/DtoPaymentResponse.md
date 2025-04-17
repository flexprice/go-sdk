# DtoPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Attempts** | Pointer to [**[]DtoPaymentAttemptResponse**](DtoPaymentAttemptResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to [**TypesPaymentDestinationType**](TypesPaymentDestinationType.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FailedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**PaymentMethodType** | Pointer to [**TypesPaymentMethodType**](TypesPaymentMethodType.md) |  | [optional] 
**PaymentStatus** | Pointer to [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | [optional] 
**RefundedAt** | Pointer to **string** |  | [optional] 
**SucceededAt** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TrackAttempts** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoPaymentResponse

`func NewDtoPaymentResponse() *DtoPaymentResponse`

NewDtoPaymentResponse instantiates a new DtoPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPaymentResponseWithDefaults

`func NewDtoPaymentResponseWithDefaults() *DtoPaymentResponse`

NewDtoPaymentResponseWithDefaults instantiates a new DtoPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoPaymentResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoPaymentResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoPaymentResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoPaymentResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAttempts

`func (o *DtoPaymentResponse) GetAttempts() []DtoPaymentAttemptResponse`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *DtoPaymentResponse) GetAttemptsOk() (*[]DtoPaymentAttemptResponse, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *DtoPaymentResponse) SetAttempts(v []DtoPaymentAttemptResponse)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *DtoPaymentResponse) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoPaymentResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoPaymentResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoPaymentResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoPaymentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoPaymentResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoPaymentResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoPaymentResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoPaymentResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoPaymentResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoPaymentResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoPaymentResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoPaymentResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDestinationId

`func (o *DtoPaymentResponse) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DtoPaymentResponse) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DtoPaymentResponse) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *DtoPaymentResponse) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationType

`func (o *DtoPaymentResponse) GetDestinationType() TypesPaymentDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DtoPaymentResponse) GetDestinationTypeOk() (*TypesPaymentDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DtoPaymentResponse) SetDestinationType(v TypesPaymentDestinationType)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *DtoPaymentResponse) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *DtoPaymentResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DtoPaymentResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DtoPaymentResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DtoPaymentResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFailedAt

`func (o *DtoPaymentResponse) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *DtoPaymentResponse) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *DtoPaymentResponse) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *DtoPaymentResponse) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetId

`func (o *DtoPaymentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoPaymentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoPaymentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoPaymentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoPaymentResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoPaymentResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoPaymentResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoPaymentResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoPaymentResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoPaymentResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoPaymentResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoPaymentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *DtoPaymentResponse) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *DtoPaymentResponse) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *DtoPaymentResponse) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *DtoPaymentResponse) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *DtoPaymentResponse) GetPaymentMethodType() TypesPaymentMethodType`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *DtoPaymentResponse) GetPaymentMethodTypeOk() (*TypesPaymentMethodType, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *DtoPaymentResponse) SetPaymentMethodType(v TypesPaymentMethodType)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *DtoPaymentResponse) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *DtoPaymentResponse) GetPaymentStatus() TypesPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *DtoPaymentResponse) GetPaymentStatusOk() (*TypesPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *DtoPaymentResponse) SetPaymentStatus(v TypesPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *DtoPaymentResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetRefundedAt

`func (o *DtoPaymentResponse) GetRefundedAt() string`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *DtoPaymentResponse) GetRefundedAtOk() (*string, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *DtoPaymentResponse) SetRefundedAt(v string)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *DtoPaymentResponse) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### GetSucceededAt

`func (o *DtoPaymentResponse) GetSucceededAt() string`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *DtoPaymentResponse) GetSucceededAtOk() (*string, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *DtoPaymentResponse) SetSucceededAt(v string)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *DtoPaymentResponse) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoPaymentResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoPaymentResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoPaymentResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoPaymentResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrackAttempts

`func (o *DtoPaymentResponse) GetTrackAttempts() bool`

GetTrackAttempts returns the TrackAttempts field if non-nil, zero value otherwise.

### GetTrackAttemptsOk

`func (o *DtoPaymentResponse) GetTrackAttemptsOk() (*bool, bool)`

GetTrackAttemptsOk returns a tuple with the TrackAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAttempts

`func (o *DtoPaymentResponse) SetTrackAttempts(v bool)`

SetTrackAttempts sets TrackAttempts field to given value.

### HasTrackAttempts

`func (o *DtoPaymentResponse) HasTrackAttempts() bool`

HasTrackAttempts returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoPaymentResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoPaymentResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoPaymentResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoPaymentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoPaymentResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoPaymentResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoPaymentResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoPaymentResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


