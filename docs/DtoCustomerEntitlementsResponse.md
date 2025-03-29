# DtoCustomerEntitlementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]DtoAggregatedFeature**](DtoAggregatedFeature.md) |  | [optional] 

## Methods

### NewDtoCustomerEntitlementsResponse

`func NewDtoCustomerEntitlementsResponse() *DtoCustomerEntitlementsResponse`

NewDtoCustomerEntitlementsResponse instantiates a new DtoCustomerEntitlementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCustomerEntitlementsResponseWithDefaults

`func NewDtoCustomerEntitlementsResponseWithDefaults() *DtoCustomerEntitlementsResponse`

NewDtoCustomerEntitlementsResponseWithDefaults instantiates a new DtoCustomerEntitlementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DtoCustomerEntitlementsResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCustomerEntitlementsResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCustomerEntitlementsResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCustomerEntitlementsResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFeatures

`func (o *DtoCustomerEntitlementsResponse) GetFeatures() []DtoAggregatedFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DtoCustomerEntitlementsResponse) GetFeaturesOk() (*[]DtoAggregatedFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DtoCustomerEntitlementsResponse) SetFeatures(v []DtoAggregatedFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DtoCustomerEntitlementsResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


