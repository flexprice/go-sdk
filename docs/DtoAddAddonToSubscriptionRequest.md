# DtoAddAddonToSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **string** |  | 
**EndDate** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoAddAddonToSubscriptionRequest

`func NewDtoAddAddonToSubscriptionRequest(addonId string, ) *DtoAddAddonToSubscriptionRequest`

NewDtoAddAddonToSubscriptionRequest instantiates a new DtoAddAddonToSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoAddAddonToSubscriptionRequestWithDefaults

`func NewDtoAddAddonToSubscriptionRequestWithDefaults() *DtoAddAddonToSubscriptionRequest`

NewDtoAddAddonToSubscriptionRequestWithDefaults instantiates a new DtoAddAddonToSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *DtoAddAddonToSubscriptionRequest) GetAddonId() string`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *DtoAddAddonToSubscriptionRequest) GetAddonIdOk() (*string, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *DtoAddAddonToSubscriptionRequest) SetAddonId(v string)`

SetAddonId sets AddonId field to given value.


### GetEndDate

`func (o *DtoAddAddonToSubscriptionRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoAddAddonToSubscriptionRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoAddAddonToSubscriptionRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoAddAddonToSubscriptionRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoAddAddonToSubscriptionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoAddAddonToSubscriptionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoAddAddonToSubscriptionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoAddAddonToSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoAddAddonToSubscriptionRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoAddAddonToSubscriptionRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoAddAddonToSubscriptionRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoAddAddonToSubscriptionRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


