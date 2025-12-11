# DtoIntegrationEntityMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id is the external entity ID from the provider | 
**Provider** | **string** | provider is the integration provider name (e.g., \&quot;stripe\&quot;, \&quot;razorpay\&quot;) | 

## Methods

### NewDtoIntegrationEntityMapping

`func NewDtoIntegrationEntityMapping(id string, provider string, ) *DtoIntegrationEntityMapping`

NewDtoIntegrationEntityMapping instantiates a new DtoIntegrationEntityMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoIntegrationEntityMappingWithDefaults

`func NewDtoIntegrationEntityMappingWithDefaults() *DtoIntegrationEntityMapping`

NewDtoIntegrationEntityMappingWithDefaults instantiates a new DtoIntegrationEntityMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DtoIntegrationEntityMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoIntegrationEntityMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoIntegrationEntityMapping) SetId(v string)`

SetId sets Id field to given value.


### GetProvider

`func (o *DtoIntegrationEntityMapping) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DtoIntegrationEntityMapping) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DtoIntegrationEntityMapping) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


