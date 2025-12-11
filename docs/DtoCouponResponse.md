# DtoCouponResponse

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

### NewDtoCouponResponse

`func NewDtoCouponResponse() *DtoCouponResponse`

NewDtoCouponResponse instantiates a new DtoCouponResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCouponResponseWithDefaults

`func NewDtoCouponResponseWithDefaults() *DtoCouponResponse`

NewDtoCouponResponseWithDefaults instantiates a new DtoCouponResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOff

`func (o *DtoCouponResponse) GetAmountOff() float32`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *DtoCouponResponse) GetAmountOffOk() (*float32, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *DtoCouponResponse) SetAmountOff(v float32)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *DtoCouponResponse) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetCadence

`func (o *DtoCouponResponse) GetCadence() TypesCouponCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DtoCouponResponse) GetCadenceOk() (*TypesCouponCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DtoCouponResponse) SetCadence(v TypesCouponCadence)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *DtoCouponResponse) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoCouponResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoCouponResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoCouponResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoCouponResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoCouponResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoCouponResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoCouponResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoCouponResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCouponResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCouponResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCouponResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCouponResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDurationInPeriods

`func (o *DtoCouponResponse) GetDurationInPeriods() int32`

GetDurationInPeriods returns the DurationInPeriods field if non-nil, zero value otherwise.

### GetDurationInPeriodsOk

`func (o *DtoCouponResponse) GetDurationInPeriodsOk() (*int32, bool)`

GetDurationInPeriodsOk returns a tuple with the DurationInPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInPeriods

`func (o *DtoCouponResponse) SetDurationInPeriods(v int32)`

SetDurationInPeriods sets DurationInPeriods field to given value.

### HasDurationInPeriods

`func (o *DtoCouponResponse) HasDurationInPeriods() bool`

HasDurationInPeriods returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCouponResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCouponResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCouponResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCouponResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoCouponResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoCouponResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoCouponResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoCouponResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *DtoCouponResponse) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *DtoCouponResponse) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *DtoCouponResponse) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *DtoCouponResponse) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCouponResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCouponResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCouponResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCouponResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCouponResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCouponResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCouponResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoCouponResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPercentageOff

`func (o *DtoCouponResponse) GetPercentageOff() float32`

GetPercentageOff returns the PercentageOff field if non-nil, zero value otherwise.

### GetPercentageOffOk

`func (o *DtoCouponResponse) GetPercentageOffOk() (*float32, bool)`

GetPercentageOffOk returns a tuple with the PercentageOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOff

`func (o *DtoCouponResponse) SetPercentageOff(v float32)`

SetPercentageOff sets PercentageOff field to given value.

### HasPercentageOff

`func (o *DtoCouponResponse) HasPercentageOff() bool`

HasPercentageOff returns a boolean if a field has been set.

### GetRedeemAfter

`func (o *DtoCouponResponse) GetRedeemAfter() string`

GetRedeemAfter returns the RedeemAfter field if non-nil, zero value otherwise.

### GetRedeemAfterOk

`func (o *DtoCouponResponse) GetRedeemAfterOk() (*string, bool)`

GetRedeemAfterOk returns a tuple with the RedeemAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAfter

`func (o *DtoCouponResponse) SetRedeemAfter(v string)`

SetRedeemAfter sets RedeemAfter field to given value.

### HasRedeemAfter

`func (o *DtoCouponResponse) HasRedeemAfter() bool`

HasRedeemAfter returns a boolean if a field has been set.

### GetRedeemBefore

`func (o *DtoCouponResponse) GetRedeemBefore() string`

GetRedeemBefore returns the RedeemBefore field if non-nil, zero value otherwise.

### GetRedeemBeforeOk

`func (o *DtoCouponResponse) GetRedeemBeforeOk() (*string, bool)`

GetRedeemBeforeOk returns a tuple with the RedeemBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBefore

`func (o *DtoCouponResponse) SetRedeemBefore(v string)`

SetRedeemBefore sets RedeemBefore field to given value.

### HasRedeemBefore

`func (o *DtoCouponResponse) HasRedeemBefore() bool`

HasRedeemBefore returns a boolean if a field has been set.

### GetRules

`func (o *DtoCouponResponse) GetRules() map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DtoCouponResponse) GetRulesOk() (*map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DtoCouponResponse) SetRules(v map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *DtoCouponResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCouponResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCouponResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCouponResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCouponResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoCouponResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoCouponResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoCouponResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoCouponResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTotalRedemptions

`func (o *DtoCouponResponse) GetTotalRedemptions() int32`

GetTotalRedemptions returns the TotalRedemptions field if non-nil, zero value otherwise.

### GetTotalRedemptionsOk

`func (o *DtoCouponResponse) GetTotalRedemptionsOk() (*int32, bool)`

GetTotalRedemptionsOk returns a tuple with the TotalRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRedemptions

`func (o *DtoCouponResponse) SetTotalRedemptions(v int32)`

SetTotalRedemptions sets TotalRedemptions field to given value.

### HasTotalRedemptions

`func (o *DtoCouponResponse) HasTotalRedemptions() bool`

HasTotalRedemptions returns a boolean if a field has been set.

### GetType

`func (o *DtoCouponResponse) GetType() TypesCouponType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCouponResponse) GetTypeOk() (*TypesCouponType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCouponResponse) SetType(v TypesCouponType)`

SetType sets Type field to given value.

### HasType

`func (o *DtoCouponResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoCouponResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoCouponResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoCouponResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoCouponResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoCouponResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoCouponResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoCouponResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoCouponResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


