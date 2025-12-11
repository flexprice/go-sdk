# DtoSubscriptionEntitlementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]DtoAggregatedFeature**](DtoAggregatedFeature.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoSubscriptionEntitlementsResponse

`func NewDtoSubscriptionEntitlementsResponse() *DtoSubscriptionEntitlementsResponse`

NewDtoSubscriptionEntitlementsResponse instantiates a new DtoSubscriptionEntitlementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionEntitlementsResponseWithDefaults

`func NewDtoSubscriptionEntitlementsResponseWithDefaults() *DtoSubscriptionEntitlementsResponse`

NewDtoSubscriptionEntitlementsResponseWithDefaults instantiates a new DtoSubscriptionEntitlementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *DtoSubscriptionEntitlementsResponse) GetFeatures() []DtoAggregatedFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DtoSubscriptionEntitlementsResponse) GetFeaturesOk() (*[]DtoAggregatedFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DtoSubscriptionEntitlementsResponse) SetFeatures(v []DtoAggregatedFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DtoSubscriptionEntitlementsResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoSubscriptionEntitlementsResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoSubscriptionEntitlementsResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoSubscriptionEntitlementsResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoSubscriptionEntitlementsResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoSubscriptionEntitlementsResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoSubscriptionEntitlementsResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoSubscriptionEntitlementsResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoSubscriptionEntitlementsResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


