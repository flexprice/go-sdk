# DtoCreateCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOff** | Pointer to **string** |  | [optional] 
**Cadence** | [**TypesCouponCadence**](TypesCouponCadence.md) |  | 
**Currency** | Pointer to **string** |  | [optional] 
**DurationInPeriods** | Pointer to **int32** |  | [optional] 
**MaxRedemptions** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**PercentageOff** | Pointer to **string** |  | [optional] 
**RedeemAfter** | Pointer to **string** |  | [optional] 
**RedeemBefore** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | [**TypesCouponType**](TypesCouponType.md) |  | 

## Methods

### NewDtoCreateCouponRequest

`func NewDtoCreateCouponRequest(cadence TypesCouponCadence, name string, type_ TypesCouponType, ) *DtoCreateCouponRequest`

NewDtoCreateCouponRequest instantiates a new DtoCreateCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateCouponRequestWithDefaults

`func NewDtoCreateCouponRequestWithDefaults() *DtoCreateCouponRequest`

NewDtoCreateCouponRequestWithDefaults instantiates a new DtoCreateCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOff

`func (o *DtoCreateCouponRequest) GetAmountOff() string`

GetAmountOff returns the AmountOff field if non-nil, zero value otherwise.

### GetAmountOffOk

`func (o *DtoCreateCouponRequest) GetAmountOffOk() (*string, bool)`

GetAmountOffOk returns a tuple with the AmountOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOff

`func (o *DtoCreateCouponRequest) SetAmountOff(v string)`

SetAmountOff sets AmountOff field to given value.

### HasAmountOff

`func (o *DtoCreateCouponRequest) HasAmountOff() bool`

HasAmountOff returns a boolean if a field has been set.

### GetCadence

`func (o *DtoCreateCouponRequest) GetCadence() TypesCouponCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *DtoCreateCouponRequest) GetCadenceOk() (*TypesCouponCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *DtoCreateCouponRequest) SetCadence(v TypesCouponCadence)`

SetCadence sets Cadence field to given value.


### GetCurrency

`func (o *DtoCreateCouponRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateCouponRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateCouponRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCreateCouponRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDurationInPeriods

`func (o *DtoCreateCouponRequest) GetDurationInPeriods() int32`

GetDurationInPeriods returns the DurationInPeriods field if non-nil, zero value otherwise.

### GetDurationInPeriodsOk

`func (o *DtoCreateCouponRequest) GetDurationInPeriodsOk() (*int32, bool)`

GetDurationInPeriodsOk returns a tuple with the DurationInPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInPeriods

`func (o *DtoCreateCouponRequest) SetDurationInPeriods(v int32)`

SetDurationInPeriods sets DurationInPeriods field to given value.

### HasDurationInPeriods

`func (o *DtoCreateCouponRequest) HasDurationInPeriods() bool`

HasDurationInPeriods returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *DtoCreateCouponRequest) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *DtoCreateCouponRequest) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *DtoCreateCouponRequest) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *DtoCreateCouponRequest) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateCouponRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateCouponRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateCouponRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateCouponRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateCouponRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateCouponRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateCouponRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPercentageOff

`func (o *DtoCreateCouponRequest) GetPercentageOff() string`

GetPercentageOff returns the PercentageOff field if non-nil, zero value otherwise.

### GetPercentageOffOk

`func (o *DtoCreateCouponRequest) GetPercentageOffOk() (*string, bool)`

GetPercentageOffOk returns a tuple with the PercentageOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOff

`func (o *DtoCreateCouponRequest) SetPercentageOff(v string)`

SetPercentageOff sets PercentageOff field to given value.

### HasPercentageOff

`func (o *DtoCreateCouponRequest) HasPercentageOff() bool`

HasPercentageOff returns a boolean if a field has been set.

### GetRedeemAfter

`func (o *DtoCreateCouponRequest) GetRedeemAfter() string`

GetRedeemAfter returns the RedeemAfter field if non-nil, zero value otherwise.

### GetRedeemAfterOk

`func (o *DtoCreateCouponRequest) GetRedeemAfterOk() (*string, bool)`

GetRedeemAfterOk returns a tuple with the RedeemAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAfter

`func (o *DtoCreateCouponRequest) SetRedeemAfter(v string)`

SetRedeemAfter sets RedeemAfter field to given value.

### HasRedeemAfter

`func (o *DtoCreateCouponRequest) HasRedeemAfter() bool`

HasRedeemAfter returns a boolean if a field has been set.

### GetRedeemBefore

`func (o *DtoCreateCouponRequest) GetRedeemBefore() string`

GetRedeemBefore returns the RedeemBefore field if non-nil, zero value otherwise.

### GetRedeemBeforeOk

`func (o *DtoCreateCouponRequest) GetRedeemBeforeOk() (*string, bool)`

GetRedeemBeforeOk returns a tuple with the RedeemBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemBefore

`func (o *DtoCreateCouponRequest) SetRedeemBefore(v string)`

SetRedeemBefore sets RedeemBefore field to given value.

### HasRedeemBefore

`func (o *DtoCreateCouponRequest) HasRedeemBefore() bool`

HasRedeemBefore returns a boolean if a field has been set.

### GetRules

`func (o *DtoCreateCouponRequest) GetRules() map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DtoCreateCouponRequest) GetRulesOk() (*map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DtoCreateCouponRequest) SetRules(v map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *DtoCreateCouponRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetType

`func (o *DtoCreateCouponRequest) GetType() TypesCouponType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreateCouponRequest) GetTypeOk() (*TypesCouponType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreateCouponRequest) SetType(v TypesCouponType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


