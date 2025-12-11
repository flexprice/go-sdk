# GithubComFlexpriceFlexpriceInternalDomainCouponCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOff** | Pointer to **float32** |  | [optional] 
**Cadence** | Pointer to [**TypesCouponCadence**](TypesCouponCadence.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DurationInPeriods** | Pointer to **int32** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxRedemptions** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PercentageOff** | Pointer to **float32** |  | [optional] 
**RedeemAfter** | Pointer to **string** |  | [optional] 
**RedeemBefore** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TotalRedemptions** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**TypesCouponType**](TypesCouponType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComFlexpriceFlexpriceInternalDomainCouponCoupon

`func NewGithubComFlexpriceFlexpriceInternalDomainCouponCoupon() *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon`

NewGithubComFlexpriceFlexpriceInternalDomainCouponCoupon instantiates a new GithubComFlexpriceFlexpriceInternalDomainCouponCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComFlexpriceFlexpriceInternalDomainCouponCouponWithDefaults

`func NewGithubComFlexpriceFlexpriceInternalDomainCouponCouponWithDefaults() *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon`

NewGithubComFlexpriceFlexpriceInternalDomainCouponCouponWithDefaults instantiates a new GithubComFlexpriceFlexpriceInternalDomainCouponCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetAmountOff() float32`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetAmountOffOk() (*float32, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetAmountOff(v float32)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetCadence

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCadence() TypesCouponCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCadenceOk() (*TypesCouponCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetCadence(v TypesCouponCadence)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDurationInPeriods

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetDurationInPeriods() int32`

GetDurationInPeriods returns the DurationInPeriods field if non-nil, zero value otherwise.

### GetDurationInPeriodsOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetDurationInPeriodsOk() (*int32, bool)`

GetDurationInPeriodsOk returns a tuple with the DurationInPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInPeriods

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetDurationInPeriods(v int32)`

SetDurationInPeriods sets DurationInPeriods field to given value.

### HasDurationInPeriods

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasDurationInPeriods() bool`

HasDurationInPeriods returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetMetadata

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPercentageOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetPercentageOff() float32`

GetPercentageOff returns the PercentageOff field if non-nil, zero value otherwise.

### GetPercentageOffOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetPercentageOffOk() (*float32, bool)`

GetPercentageOffOk returns a tuple with the PercentageOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetPercentageOff(v float32)`

SetPercentageOff sets PercentageOff field to given value.

### HasPercentageOff

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasPercentageOff() bool`

HasPercentageOff returns a boolean if a field has been set.

### GetRedeemAfter

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRedeemAfter() string`

GetRedeemAfter returns the RedeemAfter field if non-nil, zero value otherwise.

### GetRedeemAfterOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRedeemAfterOk() (*string, bool)`

GetRedeemAfterOk returns a tuple with the RedeemAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAfter

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetRedeemAfter(v string)`

SetRedeemAfter sets RedeemAfter field to given value.

### HasRedeemAfter

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasRedeemAfter() bool`

HasRedeemAfter returns a boolean if a field has been set.

### GetRedeemBefore

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRedeemBefore() string`

GetRedeemBefore returns the RedeemBefore field if non-nil, zero value otherwise.

### GetRedeemBeforeOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRedeemBeforeOk() (*string, bool)`

GetRedeemBeforeOk returns a tuple with the RedeemBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBefore

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetRedeemBefore(v string)`

SetRedeemBefore sets RedeemBefore field to given value.

### HasRedeemBefore

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasRedeemBefore() bool`

HasRedeemBefore returns a boolean if a field has been set.

### GetRules

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRules() map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetRulesOk() (*map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetRules(v map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTotalRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetTotalRedemptions() int32`

GetTotalRedemptions returns the TotalRedemptions field if non-nil, zero value otherwise.

### GetTotalRedemptionsOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetTotalRedemptionsOk() (*int32, bool)`

GetTotalRedemptionsOk returns a tuple with the TotalRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetTotalRedemptions(v int32)`

SetTotalRedemptions sets TotalRedemptions field to given value.

### HasTotalRedemptions

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasTotalRedemptions() bool`

HasTotalRedemptions returns a boolean if a field has been set.

### GetType

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetType() TypesCouponType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetTypeOk() (*TypesCouponType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetType(v TypesCouponType)`

SetType sets Type field to given value.

### HasType

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GithubComFlexpriceFlexpriceInternalDomainCouponCoupon) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


