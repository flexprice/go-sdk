# DtoUpdatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditGrants** | Pointer to [**[]DtoUpdatePlanCreditGrantRequest**](DtoUpdatePlanCreditGrantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Entitlements** | Pointer to [**[]DtoUpdatePlanEntitlementRequest**](DtoUpdatePlanEntitlementRequest.md) |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]DtoUpdatePlanPriceRequest**](DtoUpdatePlanPriceRequest.md) |  | [optional] 

## Methods

### NewDtoUpdatePlanRequest

`func NewDtoUpdatePlanRequest() *DtoUpdatePlanRequest`

NewDtoUpdatePlanRequest instantiates a new DtoUpdatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePlanRequestWithDefaults

`func NewDtoUpdatePlanRequestWithDefaults() *DtoUpdatePlanRequest`

NewDtoUpdatePlanRequestWithDefaults instantiates a new DtoUpdatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditGrants

`func (o *DtoUpdatePlanRequest) GetCreditGrants() []DtoUpdatePlanCreditGrantRequest`

GetCreditGrants returns the CreditGrants field if non-nil, zero value otherwise.

### GetCreditGrantsOk

`func (o *DtoUpdatePlanRequest) GetCreditGrantsOk() (*[]DtoUpdatePlanCreditGrantRequest, bool)`

GetCreditGrantsOk returns a tuple with the CreditGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrants

`func (o *DtoUpdatePlanRequest) SetCreditGrants(v []DtoUpdatePlanCreditGrantRequest)`

SetCreditGrants sets CreditGrants field to given value.

### HasCreditGrants

`func (o *DtoUpdatePlanRequest) HasCreditGrants() bool`

HasCreditGrants returns a boolean if a field has been set.

### GetDescription

`func (o *DtoUpdatePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoUpdatePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoUpdatePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoUpdatePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *DtoUpdatePlanRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *DtoUpdatePlanRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *DtoUpdatePlanRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *DtoUpdatePlanRequest) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetEntitlements

`func (o *DtoUpdatePlanRequest) GetEntitlements() []DtoUpdatePlanEntitlementRequest`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *DtoUpdatePlanRequest) GetEntitlementsOk() (*[]DtoUpdatePlanEntitlementRequest, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *DtoUpdatePlanRequest) SetEntitlements(v []DtoUpdatePlanEntitlementRequest)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *DtoUpdatePlanRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoUpdatePlanRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoUpdatePlanRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoUpdatePlanRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoUpdatePlanRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdatePlanRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdatePlanRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdatePlanRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdatePlanRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdatePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdatePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdatePlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdatePlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrices

`func (o *DtoUpdatePlanRequest) GetPrices() []DtoUpdatePlanPriceRequest`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DtoUpdatePlanRequest) GetPricesOk() (*[]DtoUpdatePlanPriceRequest, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DtoUpdatePlanRequest) SetPrices(v []DtoUpdatePlanPriceRequest)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *DtoUpdatePlanRequest) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


