# DtoCreateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingDetails** | Pointer to [**DtoTenantBillingDetails**](DtoTenantBillingDetails.md) |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewDtoCreateTenantRequest

`func NewDtoCreateTenantRequest(name string, ) *DtoCreateTenantRequest`

NewDtoCreateTenantRequest instantiates a new DtoCreateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateTenantRequestWithDefaults

`func NewDtoCreateTenantRequestWithDefaults() *DtoCreateTenantRequest`

NewDtoCreateTenantRequestWithDefaults instantiates a new DtoCreateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingDetails

`func (o *DtoCreateTenantRequest) GetBillingDetails() DtoTenantBillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *DtoCreateTenantRequest) GetBillingDetailsOk() (*DtoTenantBillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *DtoCreateTenantRequest) SetBillingDetails(v DtoTenantBillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *DtoCreateTenantRequest) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateTenantRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


