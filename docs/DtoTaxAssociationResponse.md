# DtoTaxAssociationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApply** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesTaxRateEntityType**](TypesTaxRateEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaxRate** | Pointer to [**DtoTaxRateResponse**](DtoTaxRateResponse.md) |  | [optional] 
**TaxRateId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**ValidFrom** | Pointer to **string** |  | [optional] 
**ValidTo** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTaxAssociationResponse

`func NewDtoTaxAssociationResponse() *DtoTaxAssociationResponse`

NewDtoTaxAssociationResponse instantiates a new DtoTaxAssociationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaxAssociationResponseWithDefaults

`func NewDtoTaxAssociationResponseWithDefaults() *DtoTaxAssociationResponse`

NewDtoTaxAssociationResponseWithDefaults instantiates a new DtoTaxAssociationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApply

`func (o *DtoTaxAssociationResponse) GetAutoApply() bool`

GetAutoApply returns the AutoApply field if non-nil, zero value otherwise.

### GetAutoApplyOk

`func (o *DtoTaxAssociationResponse) GetAutoApplyOk() (*bool, bool)`

GetAutoApplyOk returns a tuple with the AutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApply

`func (o *DtoTaxAssociationResponse) SetAutoApply(v bool)`

SetAutoApply sets AutoApply field to given value.

### HasAutoApply

`func (o *DtoTaxAssociationResponse) HasAutoApply() bool`

HasAutoApply returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTaxAssociationResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTaxAssociationResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTaxAssociationResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTaxAssociationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoTaxAssociationResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTaxAssociationResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTaxAssociationResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTaxAssociationResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoTaxAssociationResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoTaxAssociationResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoTaxAssociationResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoTaxAssociationResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoTaxAssociationResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoTaxAssociationResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoTaxAssociationResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoTaxAssociationResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoTaxAssociationResponse) GetEntityType() TypesTaxRateEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoTaxAssociationResponse) GetEntityTypeOk() (*TypesTaxRateEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoTaxAssociationResponse) SetEntityType(v TypesTaxRateEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoTaxAssociationResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoTaxAssociationResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoTaxAssociationResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoTaxAssociationResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoTaxAssociationResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoTaxAssociationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTaxAssociationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTaxAssociationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTaxAssociationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTaxAssociationResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaxAssociationResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaxAssociationResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaxAssociationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPriority

`func (o *DtoTaxAssociationResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoTaxAssociationResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoTaxAssociationResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoTaxAssociationResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *DtoTaxAssociationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoTaxAssociationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoTaxAssociationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoTaxAssociationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxRate

`func (o *DtoTaxAssociationResponse) GetTaxRate() DtoTaxRateResponse`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *DtoTaxAssociationResponse) GetTaxRateOk() (*DtoTaxRateResponse, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *DtoTaxAssociationResponse) SetTaxRate(v DtoTaxRateResponse)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *DtoTaxAssociationResponse) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetTaxRateId

`func (o *DtoTaxAssociationResponse) GetTaxRateId() string`

GetTaxRateId returns the TaxRateId field if non-nil, zero value otherwise.

### GetTaxRateIdOk

`func (o *DtoTaxAssociationResponse) GetTaxRateIdOk() (*string, bool)`

GetTaxRateIdOk returns a tuple with the TaxRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateId

`func (o *DtoTaxAssociationResponse) SetTaxRateId(v string)`

SetTaxRateId sets TaxRateId field to given value.

### HasTaxRateId

`func (o *DtoTaxAssociationResponse) HasTaxRateId() bool`

HasTaxRateId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoTaxAssociationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoTaxAssociationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoTaxAssociationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoTaxAssociationResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoTaxAssociationResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTaxAssociationResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTaxAssociationResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTaxAssociationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoTaxAssociationResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTaxAssociationResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTaxAssociationResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTaxAssociationResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValidFrom

`func (o *DtoTaxAssociationResponse) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *DtoTaxAssociationResponse) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *DtoTaxAssociationResponse) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *DtoTaxAssociationResponse) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *DtoTaxAssociationResponse) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *DtoTaxAssociationResponse) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *DtoTaxAssociationResponse) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *DtoTaxAssociationResponse) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


