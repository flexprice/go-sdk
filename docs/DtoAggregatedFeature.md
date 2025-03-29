# DtoAggregatedFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlement** | Pointer to [**DtoAggregatedEntitlement**](DtoAggregatedEntitlement.md) |  | [optional] 
**Feature** | Pointer to [**DtoFeatureResponse**](DtoFeatureResponse.md) |  | [optional] 
**Sources** | Pointer to [**[]DtoEntitlementSource**](DtoEntitlementSource.md) |  | [optional] 

## Methods

### NewDtoAggregatedFeature

`func NewDtoAggregatedFeature() *DtoAggregatedFeature`

NewDtoAggregatedFeature instantiates a new DtoAggregatedFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoAggregatedFeatureWithDefaults

`func NewDtoAggregatedFeatureWithDefaults() *DtoAggregatedFeature`

NewDtoAggregatedFeatureWithDefaults instantiates a new DtoAggregatedFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlement

`func (o *DtoAggregatedFeature) GetEntitlement() DtoAggregatedEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *DtoAggregatedFeature) GetEntitlementOk() (*DtoAggregatedEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *DtoAggregatedFeature) SetEntitlement(v DtoAggregatedEntitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *DtoAggregatedFeature) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetFeature

`func (o *DtoAggregatedFeature) GetFeature() DtoFeatureResponse`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DtoAggregatedFeature) GetFeatureOk() (*DtoFeatureResponse, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DtoAggregatedFeature) SetFeature(v DtoFeatureResponse)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DtoAggregatedFeature) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetSources

`func (o *DtoAggregatedFeature) GetSources() []DtoEntitlementSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtoAggregatedFeature) GetSourcesOk() (*[]DtoEntitlementSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtoAggregatedFeature) SetSources(v []DtoEntitlementSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtoAggregatedFeature) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


