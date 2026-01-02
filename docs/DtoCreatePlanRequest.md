# DtoCreatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 

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

### GetDisplayOrder

`func (o *DtoCreatePlanRequest) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *DtoCreatePlanRequest) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *DtoCreatePlanRequest) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *DtoCreatePlanRequest) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

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

### GetMetadata

`func (o *DtoCreatePlanRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreatePlanRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreatePlanRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreatePlanRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


