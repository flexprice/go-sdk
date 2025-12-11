# DtoTaxRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**FixedValue** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PercentageValue** | Pointer to **float32** |  | [optional] 
**Scope** | Pointer to [**TypesTaxRateScope**](TypesTaxRateScope.md) |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TaxRateStatus** | Pointer to [**TypesTaxRateStatus**](TypesTaxRateStatus.md) |  | [optional] 
**TaxRateType** | Pointer to [**TypesTaxRateType**](TypesTaxRateType.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTaxRateResponse

`func NewDtoTaxRateResponse() *DtoTaxRateResponse`

NewDtoTaxRateResponse instantiates a new DtoTaxRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaxRateResponseWithDefaults

`func NewDtoTaxRateResponseWithDefaults() *DtoTaxRateResponse`

NewDtoTaxRateResponseWithDefaults instantiates a new DtoTaxRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DtoTaxRateResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DtoTaxRateResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DtoTaxRateResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DtoTaxRateResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTaxRateResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTaxRateResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTaxRateResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTaxRateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoTaxRateResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTaxRateResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTaxRateResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTaxRateResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *DtoTaxRateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoTaxRateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoTaxRateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoTaxRateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoTaxRateResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoTaxRateResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoTaxRateResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoTaxRateResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFixedValue

`func (o *DtoTaxRateResponse) GetFixedValue() float32`

GetFixedValue returns the FixedValue field if non-nil, zero value otherwise.

### GetFixedValueOk

`func (o *DtoTaxRateResponse) GetFixedValueOk() (*float32, bool)`

GetFixedValueOk returns a tuple with the FixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedValue

`func (o *DtoTaxRateResponse) SetFixedValue(v float32)`

SetFixedValue sets FixedValue field to given value.

### HasFixedValue

`func (o *DtoTaxRateResponse) HasFixedValue() bool`

HasFixedValue returns a boolean if a field has been set.

### GetId

`func (o *DtoTaxRateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTaxRateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTaxRateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTaxRateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTaxRateResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaxRateResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaxRateResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaxRateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoTaxRateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoTaxRateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoTaxRateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoTaxRateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPercentageValue

`func (o *DtoTaxRateResponse) GetPercentageValue() float32`

GetPercentageValue returns the PercentageValue field if non-nil, zero value otherwise.

### GetPercentageValueOk

`func (o *DtoTaxRateResponse) GetPercentageValueOk() (*float32, bool)`

GetPercentageValueOk returns a tuple with the PercentageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageValue

`func (o *DtoTaxRateResponse) SetPercentageValue(v float32)`

SetPercentageValue sets PercentageValue field to given value.

### HasPercentageValue

`func (o *DtoTaxRateResponse) HasPercentageValue() bool`

HasPercentageValue returns a boolean if a field has been set.

### GetScope

`func (o *DtoTaxRateResponse) GetScope() TypesTaxRateScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DtoTaxRateResponse) GetScopeOk() (*TypesTaxRateScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DtoTaxRateResponse) SetScope(v TypesTaxRateScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DtoTaxRateResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *DtoTaxRateResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoTaxRateResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoTaxRateResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoTaxRateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxRateStatus

`func (o *DtoTaxRateResponse) GetTaxRateStatus() TypesTaxRateStatus`

GetTaxRateStatus returns the TaxRateStatus field if non-nil, zero value otherwise.

### GetTaxRateStatusOk

`func (o *DtoTaxRateResponse) GetTaxRateStatusOk() (*TypesTaxRateStatus, bool)`

GetTaxRateStatusOk returns a tuple with the TaxRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateStatus

`func (o *DtoTaxRateResponse) SetTaxRateStatus(v TypesTaxRateStatus)`

SetTaxRateStatus sets TaxRateStatus field to given value.

### HasTaxRateStatus

`func (o *DtoTaxRateResponse) HasTaxRateStatus() bool`

HasTaxRateStatus returns a boolean if a field has been set.

### GetTaxRateType

`func (o *DtoTaxRateResponse) GetTaxRateType() TypesTaxRateType`

GetTaxRateType returns the TaxRateType field if non-nil, zero value otherwise.

### GetTaxRateTypeOk

`func (o *DtoTaxRateResponse) GetTaxRateTypeOk() (*TypesTaxRateType, bool)`

GetTaxRateTypeOk returns a tuple with the TaxRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateType

`func (o *DtoTaxRateResponse) SetTaxRateType(v TypesTaxRateType)`

SetTaxRateType sets TaxRateType field to given value.

### HasTaxRateType

`func (o *DtoTaxRateResponse) HasTaxRateType() bool`

HasTaxRateType returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoTaxRateResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoTaxRateResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoTaxRateResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoTaxRateResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoTaxRateResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTaxRateResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTaxRateResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTaxRateResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoTaxRateResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTaxRateResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTaxRateResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTaxRateResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


