# DtoRemoveAddonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonAssociationId** | **string** |  | 
**EffectiveFrom** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoRemoveAddonRequest

`func NewDtoRemoveAddonRequest(addonAssociationId string, ) *DtoRemoveAddonRequest`

NewDtoRemoveAddonRequest instantiates a new DtoRemoveAddonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoRemoveAddonRequestWithDefaults

`func NewDtoRemoveAddonRequestWithDefaults() *DtoRemoveAddonRequest`

NewDtoRemoveAddonRequestWithDefaults instantiates a new DtoRemoveAddonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonAssociationId

`func (o *DtoRemoveAddonRequest) GetAddonAssociationId() string`

GetAddonAssociationId returns the AddonAssociationId field if non-nil, zero value otherwise.

### GetAddonAssociationIdOk

`func (o *DtoRemoveAddonRequest) GetAddonAssociationIdOk() (*string, bool)`

GetAddonAssociationIdOk returns a tuple with the AddonAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonAssociationId

`func (o *DtoRemoveAddonRequest) SetAddonAssociationId(v string)`

SetAddonAssociationId sets AddonAssociationId field to given value.


### GetEffectiveFrom

`func (o *DtoRemoveAddonRequest) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *DtoRemoveAddonRequest) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *DtoRemoveAddonRequest) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *DtoRemoveAddonRequest) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetReason

`func (o *DtoRemoveAddonRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoRemoveAddonRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoRemoveAddonRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoRemoveAddonRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


