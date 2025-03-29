# DtoCreatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to [**[]DtoCreatePlanEntitlementRequest**](DtoCreatePlanEntitlementRequest.md) |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Prices** | Pointer to [**[]DtoCreatePlanPriceRequest**](DtoCreatePlanPriceRequest.md) |  | [optional] 

## Methods

### NewDtoCreatePlanRequest

`func NewDtoCreatePlanRequest(name string, ) *DtoCreatePlanRequest`

NewDtoCreatePlanRequest instantiates a new DtoCreatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreatePlanRequestWithDefaults

`func NewDtoCreatePlanRequestWithDefaults() *DtoCreatePlanRequest`

NewDtoCreatePlanRequestWithDefaults instantiates a new DtoCreatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DtoCreatePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreatePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreatePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreatePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntitlements

`func (o *DtoCreatePlanRequest) GetEntitlements() []DtoCreatePlanEntitlementRequest`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *DtoCreatePlanRequest) GetEntitlementsOk() (*[]DtoCreatePlanEntitlementRequest, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *DtoCreatePlanRequest) SetEntitlements(v []DtoCreatePlanEntitlementRequest)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *DtoCreatePlanRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoCreatePlanRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreatePlanRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreatePlanRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreatePlanRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetName

`func (o *DtoCreatePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreatePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreatePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPrices

`func (o *DtoCreatePlanRequest) GetPrices() []DtoCreatePlanPriceRequest`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DtoCreatePlanRequest) GetPricesOk() (*[]DtoCreatePlanPriceRequest, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DtoCreatePlanRequest) SetPrices(v []DtoCreatePlanPriceRequest)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *DtoCreatePlanRequest) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


