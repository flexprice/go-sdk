# DtoUpdateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingDetails** | Pointer to [**DtoTenantBillingDetails**](DtoTenantBillingDetails.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUpdateTenantRequest

`func NewDtoUpdateTenantRequest() *DtoUpdateTenantRequest`

NewDtoUpdateTenantRequest instantiates a new DtoUpdateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateTenantRequestWithDefaults

`func NewDtoUpdateTenantRequestWithDefaults() *DtoUpdateTenantRequest`

NewDtoUpdateTenantRequestWithDefaults instantiates a new DtoUpdateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingDetails

`func (o *DtoUpdateTenantRequest) GetBillingDetails() DtoTenantBillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *DtoUpdateTenantRequest) GetBillingDetailsOk() (*DtoTenantBillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *DtoUpdateTenantRequest) SetBillingDetails(v DtoTenantBillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *DtoUpdateTenantRequest) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateTenantRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateTenantRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateTenantRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateTenantRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateTenantRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateTenantRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


