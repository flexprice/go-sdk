# DtoUpdateFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]MeterFilter**](MeterFilter.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UnitPlural** | Pointer to **string** |  | [optional] 
**UnitSingular** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUpdateFeatureRequest

`func NewDtoUpdateFeatureRequest() *DtoUpdateFeatureRequest`

NewDtoUpdateFeatureRequest instantiates a new DtoUpdateFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateFeatureRequestWithDefaults

`func NewDtoUpdateFeatureRequestWithDefaults() *DtoUpdateFeatureRequest`

NewDtoUpdateFeatureRequestWithDefaults instantiates a new DtoUpdateFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DtoUpdateFeatureRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoUpdateFeatureRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoUpdateFeatureRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoUpdateFeatureRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *DtoUpdateFeatureRequest) GetFilters() []MeterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DtoUpdateFeatureRequest) GetFiltersOk() (*[]MeterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DtoUpdateFeatureRequest) SetFilters(v []MeterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DtoUpdateFeatureRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateFeatureRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateFeatureRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateFeatureRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateFeatureRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateFeatureRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateFeatureRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateFeatureRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateFeatureRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnitPlural

`func (o *DtoUpdateFeatureRequest) GetUnitPlural() string`

GetUnitPlural returns the UnitPlural field if non-nil, zero value otherwise.

### GetUnitPluralOk

`func (o *DtoUpdateFeatureRequest) GetUnitPluralOk() (*string, bool)`

GetUnitPluralOk returns a tuple with the UnitPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPlural

`func (o *DtoUpdateFeatureRequest) SetUnitPlural(v string)`

SetUnitPlural sets UnitPlural field to given value.

### HasUnitPlural

`func (o *DtoUpdateFeatureRequest) HasUnitPlural() bool`

HasUnitPlural returns a boolean if a field has been set.

### GetUnitSingular

`func (o *DtoUpdateFeatureRequest) GetUnitSingular() string`

GetUnitSingular returns the UnitSingular field if non-nil, zero value otherwise.

### GetUnitSingularOk

`func (o *DtoUpdateFeatureRequest) GetUnitSingularOk() (*string, bool)`

GetUnitSingularOk returns a tuple with the UnitSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSingular

`func (o *DtoUpdateFeatureRequest) SetUnitSingular(v string)`

SetUnitSingular sets UnitSingular field to given value.

### HasUnitSingular

`func (o *DtoUpdateFeatureRequest) HasUnitSingular() bool`

HasUnitSingular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


