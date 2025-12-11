# DtoEntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | Pointer to [**DtoAddonResponse**](DtoAddonResponse.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesEntitlementEntityType**](TypesEntitlementEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Feature** | Pointer to [**DtoFeatureResponse**](DtoFeatureResponse.md) |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**FeatureType** | Pointer to [**TypesFeatureType**](TypesFeatureType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**ParentEntitlementId** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**DtoPlanResponse**](DtoPlanResponse.md) |  | [optional] 
**PlanId** | Pointer to **string** | TODO: Remove this once we have a proper entitlement entity type | [optional] 
**StaticValue** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesEntitlementUsageResetPeriod**](TypesEntitlementUsageResetPeriod.md) |  | [optional] 

## Methods

### NewDtoEntitlementResponse

`func NewDtoEntitlementResponse() *DtoEntitlementResponse`

NewDtoEntitlementResponse instantiates a new DtoEntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoEntitlementResponseWithDefaults

`func NewDtoEntitlementResponseWithDefaults() *DtoEntitlementResponse`

NewDtoEntitlementResponseWithDefaults instantiates a new DtoEntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *DtoEntitlementResponse) GetAddon() DtoAddonResponse`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *DtoEntitlementResponse) GetAddonOk() (*DtoAddonResponse, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *DtoEntitlementResponse) SetAddon(v DtoAddonResponse)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *DtoEntitlementResponse) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoEntitlementResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoEntitlementResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoEntitlementResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoEntitlementResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoEntitlementResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoEntitlementResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoEntitlementResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoEntitlementResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *DtoEntitlementResponse) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *DtoEntitlementResponse) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *DtoEntitlementResponse) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *DtoEntitlementResponse) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoEntitlementResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoEntitlementResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoEntitlementResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoEntitlementResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoEntitlementResponse) GetEntityType() TypesEntitlementEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoEntitlementResponse) GetEntityTypeOk() (*TypesEntitlementEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoEntitlementResponse) SetEntityType(v TypesEntitlementEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoEntitlementResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoEntitlementResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoEntitlementResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoEntitlementResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoEntitlementResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFeature

`func (o *DtoEntitlementResponse) GetFeature() DtoFeatureResponse`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DtoEntitlementResponse) GetFeatureOk() (*DtoFeatureResponse, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DtoEntitlementResponse) SetFeature(v DtoFeatureResponse)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DtoEntitlementResponse) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureId

`func (o *DtoEntitlementResponse) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *DtoEntitlementResponse) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *DtoEntitlementResponse) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *DtoEntitlementResponse) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureType

`func (o *DtoEntitlementResponse) GetFeatureType() TypesFeatureType`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *DtoEntitlementResponse) GetFeatureTypeOk() (*TypesFeatureType, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *DtoEntitlementResponse) SetFeatureType(v TypesFeatureType)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *DtoEntitlementResponse) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetId

`func (o *DtoEntitlementResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoEntitlementResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoEntitlementResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoEntitlementResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DtoEntitlementResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoEntitlementResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoEntitlementResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoEntitlementResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoEntitlementResponse) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoEntitlementResponse) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoEntitlementResponse) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoEntitlementResponse) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetParentEntitlementId

`func (o *DtoEntitlementResponse) GetParentEntitlementId() string`

GetParentEntitlementId returns the ParentEntitlementId field if non-nil, zero value otherwise.

### GetParentEntitlementIdOk

`func (o *DtoEntitlementResponse) GetParentEntitlementIdOk() (*string, bool)`

GetParentEntitlementIdOk returns a tuple with the ParentEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntitlementId

`func (o *DtoEntitlementResponse) SetParentEntitlementId(v string)`

SetParentEntitlementId sets ParentEntitlementId field to given value.

### HasParentEntitlementId

`func (o *DtoEntitlementResponse) HasParentEntitlementId() bool`

HasParentEntitlementId returns a boolean if a field has been set.

### GetPlan

`func (o *DtoEntitlementResponse) GetPlan() DtoPlanResponse`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *DtoEntitlementResponse) GetPlanOk() (*DtoPlanResponse, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *DtoEntitlementResponse) SetPlan(v DtoPlanResponse)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *DtoEntitlementResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoEntitlementResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoEntitlementResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoEntitlementResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoEntitlementResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoEntitlementResponse) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoEntitlementResponse) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoEntitlementResponse) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoEntitlementResponse) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetStatus

`func (o *DtoEntitlementResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoEntitlementResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoEntitlementResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoEntitlementResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoEntitlementResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoEntitlementResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoEntitlementResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoEntitlementResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoEntitlementResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoEntitlementResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoEntitlementResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoEntitlementResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoEntitlementResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoEntitlementResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoEntitlementResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoEntitlementResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoEntitlementResponse) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoEntitlementResponse) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoEntitlementResponse) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoEntitlementResponse) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoEntitlementResponse) GetUsageResetPeriod() TypesEntitlementUsageResetPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoEntitlementResponse) GetUsageResetPeriodOk() (*TypesEntitlementUsageResetPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoEntitlementResponse) SetUsageResetPeriod(v TypesEntitlementUsageResetPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoEntitlementResponse) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


