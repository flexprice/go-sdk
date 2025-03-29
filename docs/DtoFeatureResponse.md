# DtoFeatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Meter** | Pointer to [**DtoMeterResponse**](DtoMeterResponse.md) |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TypesFeatureType**](TypesFeatureType.md) |  | [optional] 
**UnitPlural** | Pointer to **string** |  | [optional] 
**UnitSingular** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoFeatureResponse

`func NewDtoFeatureResponse() *DtoFeatureResponse`

NewDtoFeatureResponse instantiates a new DtoFeatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoFeatureResponseWithDefaults

`func NewDtoFeatureResponseWithDefaults() *DtoFeatureResponse`

NewDtoFeatureResponseWithDefaults instantiates a new DtoFeatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DtoFeatureResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoFeatureResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoFeatureResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoFeatureResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoFeatureResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoFeatureResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoFeatureResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoFeatureResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *DtoFeatureResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoFeatureResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoFeatureResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoFeatureResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoFeatureResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoFeatureResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoFeatureResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoFeatureResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoFeatureResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoFeatureResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoFeatureResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoFeatureResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoFeatureResponse) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoFeatureResponse) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoFeatureResponse) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoFeatureResponse) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoFeatureResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoFeatureResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoFeatureResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoFeatureResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeter

`func (o *DtoFeatureResponse) GetMeter() DtoMeterResponse`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *DtoFeatureResponse) GetMeterOk() (*DtoMeterResponse, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *DtoFeatureResponse) SetMeter(v DtoMeterResponse)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *DtoFeatureResponse) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoFeatureResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoFeatureResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoFeatureResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoFeatureResponse) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetName

`func (o *DtoFeatureResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoFeatureResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoFeatureResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoFeatureResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DtoFeatureResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoFeatureResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoFeatureResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoFeatureResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoFeatureResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoFeatureResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoFeatureResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoFeatureResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetType

`func (o *DtoFeatureResponse) GetType() TypesFeatureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoFeatureResponse) GetTypeOk() (*TypesFeatureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoFeatureResponse) SetType(v TypesFeatureType)`

SetType sets Type field to given value.

### HasType

`func (o *DtoFeatureResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnitPlural

`func (o *DtoFeatureResponse) GetUnitPlural() string`

GetUnitPlural returns the UnitPlural field if non-nil, zero value otherwise.

### GetUnitPluralOk

`func (o *DtoFeatureResponse) GetUnitPluralOk() (*string, bool)`

GetUnitPluralOk returns a tuple with the UnitPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPlural

`func (o *DtoFeatureResponse) SetUnitPlural(v string)`

SetUnitPlural sets UnitPlural field to given value.

### HasUnitPlural

`func (o *DtoFeatureResponse) HasUnitPlural() bool`

HasUnitPlural returns a boolean if a field has been set.

### GetUnitSingular

`func (o *DtoFeatureResponse) GetUnitSingular() string`

GetUnitSingular returns the UnitSingular field if non-nil, zero value otherwise.

### GetUnitSingularOk

`func (o *DtoFeatureResponse) GetUnitSingularOk() (*string, bool)`

GetUnitSingularOk returns a tuple with the UnitSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSingular

`func (o *DtoFeatureResponse) SetUnitSingular(v string)`

SetUnitSingular sets UnitSingular field to given value.

### HasUnitSingular

`func (o *DtoFeatureResponse) HasUnitSingular() bool`

HasUnitSingular returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoFeatureResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoFeatureResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoFeatureResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoFeatureResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoFeatureResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoFeatureResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoFeatureResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoFeatureResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


