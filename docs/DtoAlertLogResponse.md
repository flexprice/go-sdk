# DtoAlertLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertInfo** | Pointer to [**TypesAlertInfo**](TypesAlertInfo.md) |  | [optional] 
**AlertStatus** | Pointer to [**TypesAlertState**](TypesAlertState.md) |  | [optional] 
**AlertType** | Pointer to [**TypesAlertType**](TypesAlertType.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**DtoCustomerResponse**](DtoCustomerResponse.md) |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesAlertEntityType**](TypesAlertEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Feature** | Pointer to [**DtoFeatureResponse**](DtoFeatureResponse.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ParentEntityId** | Pointer to **string** |  | [optional] 
**ParentEntityType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Wallet** | Pointer to [**DtoWalletResponse**](DtoWalletResponse.md) |  | [optional] 

## Methods

### NewDtoAlertLogResponse

`func NewDtoAlertLogResponse() *DtoAlertLogResponse`

NewDtoAlertLogResponse instantiates a new DtoAlertLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoAlertLogResponseWithDefaults

`func NewDtoAlertLogResponseWithDefaults() *DtoAlertLogResponse`

NewDtoAlertLogResponseWithDefaults instantiates a new DtoAlertLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertInfo

`func (o *DtoAlertLogResponse) GetAlertInfo() TypesAlertInfo`

GetAlertInfo returns the AlertInfo field if non-nil, zero value otherwise.

### GetAlertInfoOk

`func (o *DtoAlertLogResponse) GetAlertInfoOk() (*TypesAlertInfo, bool)`

GetAlertInfoOk returns a tuple with the AlertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertInfo

`func (o *DtoAlertLogResponse) SetAlertInfo(v TypesAlertInfo)`

SetAlertInfo sets AlertInfo field to given value.

### HasAlertInfo

`func (o *DtoAlertLogResponse) HasAlertInfo() bool`

HasAlertInfo returns a boolean if a field has been set.

### GetAlertStatus

`func (o *DtoAlertLogResponse) GetAlertStatus() TypesAlertState`

GetAlertStatus returns the AlertStatus field if non-nil, zero value otherwise.

### GetAlertStatusOk

`func (o *DtoAlertLogResponse) GetAlertStatusOk() (*TypesAlertState, bool)`

GetAlertStatusOk returns a tuple with the AlertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertStatus

`func (o *DtoAlertLogResponse) SetAlertStatus(v TypesAlertState)`

SetAlertStatus sets AlertStatus field to given value.

### HasAlertStatus

`func (o *DtoAlertLogResponse) HasAlertStatus() bool`

HasAlertStatus returns a boolean if a field has been set.

### GetAlertType

`func (o *DtoAlertLogResponse) GetAlertType() TypesAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *DtoAlertLogResponse) GetAlertTypeOk() (*TypesAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *DtoAlertLogResponse) SetAlertType(v TypesAlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *DtoAlertLogResponse) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoAlertLogResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoAlertLogResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoAlertLogResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoAlertLogResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoAlertLogResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoAlertLogResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoAlertLogResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoAlertLogResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoAlertLogResponse) GetCustomer() DtoCustomerResponse`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoAlertLogResponse) GetCustomerOk() (*DtoCustomerResponse, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoAlertLogResponse) SetCustomer(v DtoCustomerResponse)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoAlertLogResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoAlertLogResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoAlertLogResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoAlertLogResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoAlertLogResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoAlertLogResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoAlertLogResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoAlertLogResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoAlertLogResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoAlertLogResponse) GetEntityType() TypesAlertEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoAlertLogResponse) GetEntityTypeOk() (*TypesAlertEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoAlertLogResponse) SetEntityType(v TypesAlertEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoAlertLogResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoAlertLogResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoAlertLogResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoAlertLogResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoAlertLogResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFeature

`func (o *DtoAlertLogResponse) GetFeature() DtoFeatureResponse`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DtoAlertLogResponse) GetFeatureOk() (*DtoFeatureResponse, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DtoAlertLogResponse) SetFeature(v DtoFeatureResponse)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DtoAlertLogResponse) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetId

`func (o *DtoAlertLogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoAlertLogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoAlertLogResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoAlertLogResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentEntityId

`func (o *DtoAlertLogResponse) GetParentEntityId() string`

GetParentEntityId returns the ParentEntityId field if non-nil, zero value otherwise.

### GetParentEntityIdOk

`func (o *DtoAlertLogResponse) GetParentEntityIdOk() (*string, bool)`

GetParentEntityIdOk returns a tuple with the ParentEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntityId

`func (o *DtoAlertLogResponse) SetParentEntityId(v string)`

SetParentEntityId sets ParentEntityId field to given value.

### HasParentEntityId

`func (o *DtoAlertLogResponse) HasParentEntityId() bool`

HasParentEntityId returns a boolean if a field has been set.

### GetParentEntityType

`func (o *DtoAlertLogResponse) GetParentEntityType() string`

GetParentEntityType returns the ParentEntityType field if non-nil, zero value otherwise.

### GetParentEntityTypeOk

`func (o *DtoAlertLogResponse) GetParentEntityTypeOk() (*string, bool)`

GetParentEntityTypeOk returns a tuple with the ParentEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntityType

`func (o *DtoAlertLogResponse) SetParentEntityType(v string)`

SetParentEntityType sets ParentEntityType field to given value.

### HasParentEntityType

`func (o *DtoAlertLogResponse) HasParentEntityType() bool`

HasParentEntityType returns a boolean if a field has been set.

### GetStatus

`func (o *DtoAlertLogResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoAlertLogResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoAlertLogResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoAlertLogResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoAlertLogResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoAlertLogResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoAlertLogResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoAlertLogResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoAlertLogResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoAlertLogResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoAlertLogResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoAlertLogResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoAlertLogResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoAlertLogResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoAlertLogResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoAlertLogResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWallet

`func (o *DtoAlertLogResponse) GetWallet() DtoWalletResponse`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *DtoAlertLogResponse) GetWalletOk() (*DtoWalletResponse, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *DtoAlertLogResponse) SetWallet(v DtoWalletResponse)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *DtoAlertLogResponse) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


