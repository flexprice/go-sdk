# DtoTaxAppliedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesTaxRateEntityType**](TypesTaxRateEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TaxAmount** | Pointer to **string** |  | [optional] 
**TaxAssociationId** | Pointer to **string** |  | [optional] 
**TaxRate** | Pointer to [**DtoTaxRateResponse**](DtoTaxRateResponse.md) |  | [optional] 
**TaxRateId** | Pointer to **string** |  | [optional] 
**TaxableAmount** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTaxAppliedResponse

`func NewDtoTaxAppliedResponse() *DtoTaxAppliedResponse`

NewDtoTaxAppliedResponse instantiates a new DtoTaxAppliedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaxAppliedResponseWithDefaults

`func NewDtoTaxAppliedResponseWithDefaults() *DtoTaxAppliedResponse`

NewDtoTaxAppliedResponseWithDefaults instantiates a new DtoTaxAppliedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAt

`func (o *DtoTaxAppliedResponse) GetAppliedAt() string`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *DtoTaxAppliedResponse) GetAppliedAtOk() (*string, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *DtoTaxAppliedResponse) SetAppliedAt(v string)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *DtoTaxAppliedResponse) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTaxAppliedResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTaxAppliedResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTaxAppliedResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTaxAppliedResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoTaxAppliedResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTaxAppliedResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTaxAppliedResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTaxAppliedResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoTaxAppliedResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoTaxAppliedResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoTaxAppliedResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoTaxAppliedResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoTaxAppliedResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoTaxAppliedResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoTaxAppliedResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoTaxAppliedResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoTaxAppliedResponse) GetEntityType() TypesTaxRateEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoTaxAppliedResponse) GetEntityTypeOk() (*TypesTaxRateEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoTaxAppliedResponse) SetEntityType(v TypesTaxRateEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoTaxAppliedResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoTaxAppliedResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoTaxAppliedResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoTaxAppliedResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoTaxAppliedResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoTaxAppliedResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTaxAppliedResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTaxAppliedResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTaxAppliedResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoTaxAppliedResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoTaxAppliedResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoTaxAppliedResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoTaxAppliedResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTaxAppliedResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaxAppliedResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaxAppliedResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaxAppliedResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *DtoTaxAppliedResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoTaxAppliedResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoTaxAppliedResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoTaxAppliedResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAmount

`func (o *DtoTaxAppliedResponse) GetTaxAmount() string`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *DtoTaxAppliedResponse) GetTaxAmountOk() (*string, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *DtoTaxAppliedResponse) SetTaxAmount(v string)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *DtoTaxAppliedResponse) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxAssociationId

`func (o *DtoTaxAppliedResponse) GetTaxAssociationId() string`

GetTaxAssociationId returns the TaxAssociationId field if non-nil, zero value otherwise.

### GetTaxAssociationIdOk

`func (o *DtoTaxAppliedResponse) GetTaxAssociationIdOk() (*string, bool)`

GetTaxAssociationIdOk returns a tuple with the TaxAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAssociationId

`func (o *DtoTaxAppliedResponse) SetTaxAssociationId(v string)`

SetTaxAssociationId sets TaxAssociationId field to given value.

### HasTaxAssociationId

`func (o *DtoTaxAppliedResponse) HasTaxAssociationId() bool`

HasTaxAssociationId returns a boolean if a field has been set.

### GetTaxRate

`func (o *DtoTaxAppliedResponse) GetTaxRate() DtoTaxRateResponse`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *DtoTaxAppliedResponse) GetTaxRateOk() (*DtoTaxRateResponse, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *DtoTaxAppliedResponse) SetTaxRate(v DtoTaxRateResponse)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *DtoTaxAppliedResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxRateId

`func (o *DtoTaxAppliedResponse) GetTaxRateId() string`

GetTaxRateId returns the TaxRateId field if non-nil, zero value otherwise.

### GetTaxRateIdOk

`func (o *DtoTaxAppliedResponse) GetTaxRateIdOk() (*string, bool)`

GetTaxRateIdOk returns a tuple with the TaxRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateId

`func (o *DtoTaxAppliedResponse) SetTaxRateId(v string)`

SetTaxRateId sets TaxRateId field to given value.

### HasTaxRateId

`func (o *DtoTaxAppliedResponse) HasTaxRateId() bool`

HasTaxRateId returns a boolean if a field has been set.

### GetTaxableAmount

`func (o *DtoTaxAppliedResponse) GetTaxableAmount() string`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *DtoTaxAppliedResponse) GetTaxableAmountOk() (*string, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *DtoTaxAppliedResponse) SetTaxableAmount(v string)`

SetTaxableAmount sets TaxableAmount field to given value.

### HasTaxableAmount

`func (o *DtoTaxAppliedResponse) HasTaxableAmount() bool`

HasTaxableAmount returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoTaxAppliedResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoTaxAppliedResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoTaxAppliedResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoTaxAppliedResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoTaxAppliedResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTaxAppliedResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTaxAppliedResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTaxAppliedResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoTaxAppliedResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTaxAppliedResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTaxAppliedResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTaxAppliedResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


