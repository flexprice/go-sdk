# DtoTenantBillingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]DtoSubscriptionResponse**](DtoSubscriptionResponse.md) |  | [optional] 
**Usage** | Pointer to [**DtoCustomerUsageSummaryResponse**](DtoCustomerUsageSummaryResponse.md) |  | [optional] 

## Methods

### NewDtoTenantBillingUsage

`func NewDtoTenantBillingUsage() *DtoTenantBillingUsage`

NewDtoTenantBillingUsage instantiates a new DtoTenantBillingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTenantBillingUsageWithDefaults

`func NewDtoTenantBillingUsageWithDefaults() *DtoTenantBillingUsage`

NewDtoTenantBillingUsageWithDefaults instantiates a new DtoTenantBillingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *DtoTenantBillingUsage) GetSubscriptions() []DtoSubscriptionResponse`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DtoTenantBillingUsage) GetSubscriptionsOk() (*[]DtoSubscriptionResponse, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DtoTenantBillingUsage) SetSubscriptions(v []DtoSubscriptionResponse)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DtoTenantBillingUsage) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUsage

`func (o *DtoTenantBillingUsage) GetUsage() DtoCustomerUsageSummaryResponse`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DtoTenantBillingUsage) GetUsageOk() (*DtoCustomerUsageSummaryResponse, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DtoTenantBillingUsage) SetUsage(v DtoCustomerUsageSummaryResponse)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DtoTenantBillingUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


