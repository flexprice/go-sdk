# DtoCouponApplicationResponse

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
**DiscountPercentage** | Pointer to **float32** |  | [optional] 
**DiscountType** | Pointer to [**TypesCouponType**](TypesCouponType.md) |  | [optional] 
**DiscountedAmount** | Pointer to **float32** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**FinalPrice** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceLineItemId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OriginalPrice** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCouponApplicationResponse

`func NewDtoCouponApplicationResponse() *DtoCouponApplicationResponse`

NewDtoCouponApplicationResponse instantiates a new DtoCouponApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCouponApplicationResponseWithDefaults

`func NewDtoCouponApplicationResponseWithDefaults() *DtoCouponApplicationResponse`

NewDtoCouponApplicationResponseWithDefaults instantiates a new DtoCouponApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAt

`func (o *DtoCouponApplicationResponse) GetAppliedAt() string`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *DtoCouponApplicationResponse) GetAppliedAtOk() (*string, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *DtoCouponApplicationResponse) SetAppliedAt(v string)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *DtoCouponApplicationResponse) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### GetCouponAssociationId

`func (o *DtoCouponApplicationResponse) GetCouponAssociationId() string`

GetCouponAssociationId returns the CouponAssociationId field if non-nil, zero value otherwise.

### GetCouponAssociationIdOk

`func (o *DtoCouponApplicationResponse) GetCouponAssociationIdOk() (*string, bool)`

GetCouponAssociationIdOk returns a tuple with the CouponAssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponAssociationId

`func (o *DtoCouponApplicationResponse) SetCouponAssociationId(v string)`

SetCouponAssociationId sets CouponAssociationId field to given value.

### HasCouponAssociationId

`func (o *DtoCouponApplicationResponse) HasCouponAssociationId() bool`

HasCouponAssociationId returns a boolean if a field has been set.

### GetCouponId

`func (o *DtoCouponApplicationResponse) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *DtoCouponApplicationResponse) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *DtoCouponApplicationResponse) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *DtoCouponApplicationResponse) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCouponSnapshot

`func (o *DtoCouponApplicationResponse) GetCouponSnapshot() map[string]interface{}`

GetCouponSnapshot returns the CouponSnapshot field if non-nil, zero value otherwise.

### GetCouponSnapshotOk

`func (o *DtoCouponApplicationResponse) GetCouponSnapshotOk() (*map[string]interface{}, bool)`

GetCouponSnapshotOk returns a tuple with the CouponSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSnapshot

`func (o *DtoCouponApplicationResponse) SetCouponSnapshot(v map[string]interface{})`

SetCouponSnapshot sets CouponSnapshot field to given value.

### HasCouponSnapshot

`func (o *DtoCouponApplicationResponse) HasCouponSnapshot() bool`

HasCouponSnapshot returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoCouponApplicationResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoCouponApplicationResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoCouponApplicationResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoCouponApplicationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoCouponApplicationResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoCouponApplicationResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoCouponApplicationResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoCouponApplicationResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCouponApplicationResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCouponApplicationResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCouponApplicationResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCouponApplicationResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *DtoCouponApplicationResponse) GetDiscountPercentage() float32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *DtoCouponApplicationResponse) GetDiscountPercentageOk() (*float32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *DtoCouponApplicationResponse) SetDiscountPercentage(v float32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *DtoCouponApplicationResponse) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *DtoCouponApplicationResponse) GetDiscountType() TypesCouponType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *DtoCouponApplicationResponse) GetDiscountTypeOk() (*TypesCouponType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *DtoCouponApplicationResponse) SetDiscountType(v TypesCouponType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *DtoCouponApplicationResponse) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDiscountedAmount

`func (o *DtoCouponApplicationResponse) GetDiscountedAmount() float32`

GetDiscountedAmount returns the DiscountedAmount field if non-nil, zero value otherwise.

### GetDiscountedAmountOk

`func (o *DtoCouponApplicationResponse) GetDiscountedAmountOk() (*float32, bool)`

GetDiscountedAmountOk returns a tuple with the DiscountedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedAmount

`func (o *DtoCouponApplicationResponse) SetDiscountedAmount(v float32)`

SetDiscountedAmount sets DiscountedAmount field to given value.

### HasDiscountedAmount

`func (o *DtoCouponApplicationResponse) HasDiscountedAmount() bool`

HasDiscountedAmount returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCouponApplicationResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCouponApplicationResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCouponApplicationResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCouponApplicationResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFinalPrice

`func (o *DtoCouponApplicationResponse) GetFinalPrice() float32`

GetFinalPrice returns the FinalPrice field if non-nil, zero value otherwise.

### GetFinalPriceOk

`func (o *DtoCouponApplicationResponse) GetFinalPriceOk() (*float32, bool)`

GetFinalPriceOk returns a tuple with the FinalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPrice

`func (o *DtoCouponApplicationResponse) SetFinalPrice(v float32)`

SetFinalPrice sets FinalPrice field to given value.

### HasFinalPrice

`func (o *DtoCouponApplicationResponse) HasFinalPrice() bool`

HasFinalPrice returns a boolean if a field has been set.

### GetId

`func (o *DtoCouponApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoCouponApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoCouponApplicationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoCouponApplicationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *DtoCouponApplicationResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DtoCouponApplicationResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DtoCouponApplicationResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *DtoCouponApplicationResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceLineItemId

`func (o *DtoCouponApplicationResponse) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *DtoCouponApplicationResponse) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *DtoCouponApplicationResponse) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.

### HasInvoiceLineItemId

`func (o *DtoCouponApplicationResponse) HasInvoiceLineItemId() bool`

HasInvoiceLineItemId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCouponApplicationResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCouponApplicationResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCouponApplicationResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCouponApplicationResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *DtoCouponApplicationResponse) GetOriginalPrice() float32`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *DtoCouponApplicationResponse) GetOriginalPriceOk() (*float32, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *DtoCouponApplicationResponse) SetOriginalPrice(v float32)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *DtoCouponApplicationResponse) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCouponApplicationResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCouponApplicationResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCouponApplicationResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCouponApplicationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCouponApplicationResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCouponApplicationResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCouponApplicationResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCouponApplicationResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoCouponApplicationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoCouponApplicationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoCouponApplicationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoCouponApplicationResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoCouponApplicationResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoCouponApplicationResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoCouponApplicationResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoCouponApplicationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoCouponApplicationResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoCouponApplicationResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoCouponApplicationResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoCouponApplicationResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


