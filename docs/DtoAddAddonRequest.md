# DtoAddAddonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**SubscriptionId** | **string** |  | 

## Methods

### NewDtoAddAddonRequest

`func NewDtoAddAddonRequest(addonId string, subscriptionId string, ) *DtoAddAddonRequest`

NewDtoAddAddonRequest instantiates a new DtoAddAddonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoAddAddonRequestWithDefaults

`func NewDtoAddAddonRequestWithDefaults() *DtoAddAddonRequest`

NewDtoAddAddonRequestWithDefaults instantiates a new DtoAddAddonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *DtoAddAddonRequest) GetAddonId() string`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *DtoAddAddonRequest) GetAddonIdOk() (*string, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *DtoAddAddonRequest) SetAddonId(v string)`

SetAddonId sets AddonId field to given value.


### GetMetadata

`func (o *DtoAddAddonRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoAddAddonRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoAddAddonRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoAddAddonRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoAddAddonRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoAddAddonRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoAddAddonRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoAddAddonRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoAddAddonRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoAddAddonRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoAddAddonRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


