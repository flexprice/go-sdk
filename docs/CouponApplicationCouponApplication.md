# CouponApplicationCouponApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAt** | Pointer to **string** |  | [optional] 
**CouponAssociationId** | Pointer to **string** |  | [optional] 
**CouponId** | Pointer to **string** |  | [optional] 
**CouponSnapshot** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DiscountPercentage** | Pointer to **string** |  | [optional] 
**DiscountType** | Pointer to [**TypesCouponType**](TypesCouponType.md) |  | [optional] 
**DiscountedAmount** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**FinalPrice** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceLineItemId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OriginalPrice** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCouponApplicationCouponApplication

`func NewCouponApplicationCouponApplication() *CouponApplicationCouponApplication`

NewCouponApplicationCouponApplication instantiates a new CouponApplicationCouponApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponApplicationCouponApplicationWithDefaults

`func NewCouponApplicationCouponApplicationWithDefaults() *CouponApplicationCouponApplication`

NewCouponApplicationCouponApplicationWithDefaults instantiates a new CouponApplicationCouponApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAt

`func (o *CouponApplicationCouponApplication) GetAppliedAt() string`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *CouponApplicationCouponApplication) GetAppliedAtOk() (*string, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *CouponApplicationCouponApplication) SetAppliedAt(v string)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *CouponApplicationCouponApplication) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### GetCouponAssociationId

`func (o *CouponApplicationCouponApplication) GetCouponAssociationId() string`

GetCouponAssociationId returns the CouponAssociationId field if non-nil, zero value otherwise.

### GetCouponAssociationIdOk

`func (o *CouponApplicationCouponApplication) GetCouponAssociationIdOk() (*string, bool)`

GetCouponAssociationIdOk returns a tuple with the CouponAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAssociationId

`func (o *CouponApplicationCouponApplication) SetCouponAssociationId(v string)`

SetCouponAssociationId sets CouponAssociationId field to given value.

### HasCouponAssociationId

`func (o *CouponApplicationCouponApplication) HasCouponAssociationId() bool`

HasCouponAssociationId returns a boolean if a field has been set.

### GetCouponId

`func (o *CouponApplicationCouponApplication) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CouponApplicationCouponApplication) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CouponApplicationCouponApplication) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CouponApplicationCouponApplication) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCouponSnapshot

`func (o *CouponApplicationCouponApplication) GetCouponSnapshot() map[string]interface{}`

GetCouponSnapshot returns the CouponSnapshot field if non-nil, zero value otherwise.

### GetCouponSnapshotOk

`func (o *CouponApplicationCouponApplication) GetCouponSnapshotOk() (*map[string]interface{}, bool)`

GetCouponSnapshotOk returns a tuple with the CouponSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSnapshot

`func (o *CouponApplicationCouponApplication) SetCouponSnapshot(v map[string]interface{})`

SetCouponSnapshot sets CouponSnapshot field to given value.

### HasCouponSnapshot

`func (o *CouponApplicationCouponApplication) HasCouponSnapshot() bool`

HasCouponSnapshot returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CouponApplicationCouponApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponApplicationCouponApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponApplicationCouponApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CouponApplicationCouponApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CouponApplicationCouponApplication) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponApplicationCouponApplication) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CouponApplicationCouponApplication) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CouponApplicationCouponApplication) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *CouponApplicationCouponApplication) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CouponApplicationCouponApplication) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CouponApplicationCouponApplication) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CouponApplicationCouponApplication) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *CouponApplicationCouponApplication) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *CouponApplicationCouponApplication) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *CouponApplicationCouponApplication) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *CouponApplicationCouponApplication) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *CouponApplicationCouponApplication) GetDiscountType() TypesCouponType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *CouponApplicationCouponApplication) GetDiscountTypeOk() (*TypesCouponType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *CouponApplicationCouponApplication) SetDiscountType(v TypesCouponType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *CouponApplicationCouponApplication) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDiscountedAmount

`func (o *CouponApplicationCouponApplication) GetDiscountedAmount() string`

GetDiscountedAmount returns the DiscountedAmount field if non-nil, zero value otherwise.

### GetDiscountedAmountOk

`func (o *CouponApplicationCouponApplication) GetDiscountedAmountOk() (*string, bool)`

GetDiscountedAmountOk returns a tuple with the DiscountedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedAmount

`func (o *CouponApplicationCouponApplication) SetDiscountedAmount(v string)`

SetDiscountedAmount sets DiscountedAmount field to given value.

### HasDiscountedAmount

`func (o *CouponApplicationCouponApplication) HasDiscountedAmount() bool`

HasDiscountedAmount returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *CouponApplicationCouponApplication) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CouponApplicationCouponApplication) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CouponApplicationCouponApplication) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CouponApplicationCouponApplication) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFinalPrice

`func (o *CouponApplicationCouponApplication) GetFinalPrice() string`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *CouponApplicationCouponApplication) GetFinalPriceOk() (*string, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *CouponApplicationCouponApplication) SetFinalPrice(v string)`

SetFinalPrice sets FinalPrice field to given value.

### HasFinalPrice

`func (o *CouponApplicationCouponApplication) HasFinalPrice() bool`

HasFinalPrice returns a boolean if a field has been set.

### GetId

`func (o *CouponApplicationCouponApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponApplicationCouponApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponApplicationCouponApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CouponApplicationCouponApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CouponApplicationCouponApplication) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CouponApplicationCouponApplication) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CouponApplicationCouponApplication) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CouponApplicationCouponApplication) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceLineItemId

`func (o *CouponApplicationCouponApplication) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *CouponApplicationCouponApplication) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *CouponApplicationCouponApplication) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.

### HasInvoiceLineItemId

`func (o *CouponApplicationCouponApplication) HasInvoiceLineItemId() bool`

HasInvoiceLineItemId returns a boolean if a field has been set.

### GetMetadata

`func (o *CouponApplicationCouponApplication) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CouponApplicationCouponApplication) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CouponApplicationCouponApplication) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CouponApplicationCouponApplication) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *CouponApplicationCouponApplication) GetOriginalPrice() string`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *CouponApplicationCouponApplication) GetOriginalPriceOk() (*string, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *CouponApplicationCouponApplication) SetOriginalPrice(v string)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *CouponApplicationCouponApplication) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetStatus

`func (o *CouponApplicationCouponApplication) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponApplicationCouponApplication) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponApplicationCouponApplication) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CouponApplicationCouponApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CouponApplicationCouponApplication) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CouponApplicationCouponApplication) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CouponApplicationCouponApplication) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CouponApplicationCouponApplication) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *CouponApplicationCouponApplication) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CouponApplicationCouponApplication) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CouponApplicationCouponApplication) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CouponApplicationCouponApplication) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CouponApplicationCouponApplication) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponApplicationCouponApplication) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponApplicationCouponApplication) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CouponApplicationCouponApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CouponApplicationCouponApplication) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CouponApplicationCouponApplication) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CouponApplicationCouponApplication) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CouponApplicationCouponApplication) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


