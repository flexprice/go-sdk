# CouponAssociationCouponAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coupon** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainCouponCoupon**](GithubComFlexpriceFlexpriceInternalDomainCouponCoupon.md) |  | [optional] 
**CouponId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** | Optional | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | Mandatory | [optional] 
**SubscriptionLineItemId** | Pointer to **string** | Optional | [optional] 
**SubscriptionPhaseId** | Pointer to **string** | Optional | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCouponAssociationCouponAssociation

`func NewCouponAssociationCouponAssociation() *CouponAssociationCouponAssociation`

NewCouponAssociationCouponAssociation instantiates a new CouponAssociationCouponAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponAssociationCouponAssociationWithDefaults

`func NewCouponAssociationCouponAssociationWithDefaults() *CouponAssociationCouponAssociation`

NewCouponAssociationCouponAssociationWithDefaults instantiates a new CouponAssociationCouponAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoupon

`func (o *CouponAssociationCouponAssociation) GetCoupon() GithubComFlexpriceFlexpriceInternalDomainCouponCoupon`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *CouponAssociationCouponAssociation) GetCouponOk() (*GithubComFlexpriceFlexpriceInternalDomainCouponCoupon, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *CouponAssociationCouponAssociation) SetCoupon(v GithubComFlexpriceFlexpriceInternalDomainCouponCoupon)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *CouponAssociationCouponAssociation) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetCouponId

`func (o *CouponAssociationCouponAssociation) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CouponAssociationCouponAssociation) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CouponAssociationCouponAssociation) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *CouponAssociationCouponAssociation) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CouponAssociationCouponAssociation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponAssociationCouponAssociation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponAssociationCouponAssociation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CouponAssociationCouponAssociation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CouponAssociationCouponAssociation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponAssociationCouponAssociation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CouponAssociationCouponAssociation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CouponAssociationCouponAssociation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEndDate

`func (o *CouponAssociationCouponAssociation) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CouponAssociationCouponAssociation) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CouponAssociationCouponAssociation) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CouponAssociationCouponAssociation) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *CouponAssociationCouponAssociation) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CouponAssociationCouponAssociation) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CouponAssociationCouponAssociation) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CouponAssociationCouponAssociation) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *CouponAssociationCouponAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponAssociationCouponAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponAssociationCouponAssociation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CouponAssociationCouponAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CouponAssociationCouponAssociation) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CouponAssociationCouponAssociation) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CouponAssociationCouponAssociation) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CouponAssociationCouponAssociation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStartDate

`func (o *CouponAssociationCouponAssociation) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CouponAssociationCouponAssociation) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CouponAssociationCouponAssociation) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CouponAssociationCouponAssociation) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *CouponAssociationCouponAssociation) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponAssociationCouponAssociation) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponAssociationCouponAssociation) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CouponAssociationCouponAssociation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CouponAssociationCouponAssociation) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CouponAssociationCouponAssociation) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CouponAssociationCouponAssociation) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CouponAssociationCouponAssociation) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionLineItemId

`func (o *CouponAssociationCouponAssociation) GetSubscriptionLineItemId() string`

GetSubscriptionLineItemId returns the SubscriptionLineItemId field if non-nil, zero value otherwise.

### GetSubscriptionLineItemIdOk

`func (o *CouponAssociationCouponAssociation) GetSubscriptionLineItemIdOk() (*string, bool)`

GetSubscriptionLineItemIdOk returns a tuple with the SubscriptionLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLineItemId

`func (o *CouponAssociationCouponAssociation) SetSubscriptionLineItemId(v string)`

SetSubscriptionLineItemId sets SubscriptionLineItemId field to given value.

### HasSubscriptionLineItemId

`func (o *CouponAssociationCouponAssociation) HasSubscriptionLineItemId() bool`

HasSubscriptionLineItemId returns a boolean if a field has been set.

### GetSubscriptionPhaseId

`func (o *CouponAssociationCouponAssociation) GetSubscriptionPhaseId() string`

GetSubscriptionPhaseId returns the SubscriptionPhaseId field if non-nil, zero value otherwise.

### GetSubscriptionPhaseIdOk

`func (o *CouponAssociationCouponAssociation) GetSubscriptionPhaseIdOk() (*string, bool)`

GetSubscriptionPhaseIdOk returns a tuple with the SubscriptionPhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPhaseId

`func (o *CouponAssociationCouponAssociation) SetSubscriptionPhaseId(v string)`

SetSubscriptionPhaseId sets SubscriptionPhaseId field to given value.

### HasSubscriptionPhaseId

`func (o *CouponAssociationCouponAssociation) HasSubscriptionPhaseId() bool`

HasSubscriptionPhaseId returns a boolean if a field has been set.

### GetTenantId

`func (o *CouponAssociationCouponAssociation) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CouponAssociationCouponAssociation) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CouponAssociationCouponAssociation) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CouponAssociationCouponAssociation) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CouponAssociationCouponAssociation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponAssociationCouponAssociation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponAssociationCouponAssociation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CouponAssociationCouponAssociation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CouponAssociationCouponAssociation) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CouponAssociationCouponAssociation) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CouponAssociationCouponAssociation) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CouponAssociationCouponAssociation) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


