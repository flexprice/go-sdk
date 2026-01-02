# TypesCommitmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**ComputedCommitmentUtilizedAmount** | Pointer to **string** |  | [optional] 
**ComputedOverageAmount** | Pointer to **string** |  | [optional] 
**ComputedTrueUpAmount** | Pointer to **string** | total_cost &#x3D; computed_commitment_utilized_amount + computed_overage_amount + computed_true_up_amount | [optional] 
**IsWindowed** | Pointer to **bool** |  | [optional] 
**OverageFactor** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** | Only used for quantity-based commitments | [optional] 
**TrueUpEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**TypesCommitmentType**](TypesCommitmentType.md) |  | [optional] 

## Methods

### NewTypesCommitmentInfo

`func NewTypesCommitmentInfo() *TypesCommitmentInfo`

NewTypesCommitmentInfo instantiates a new TypesCommitmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesCommitmentInfoWithDefaults

`func NewTypesCommitmentInfoWithDefaults() *TypesCommitmentInfo`

NewTypesCommitmentInfoWithDefaults instantiates a new TypesCommitmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TypesCommitmentInfo) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TypesCommitmentInfo) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TypesCommitmentInfo) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TypesCommitmentInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetComputedCommitmentUtilizedAmount

`func (o *TypesCommitmentInfo) GetComputedCommitmentUtilizedAmount() string`

GetComputedCommitmentUtilizedAmount returns the ComputedCommitmentUtilizedAmount field if non-nil, zero value otherwise.

### GetComputedCommitmentUtilizedAmountOk

`func (o *TypesCommitmentInfo) GetComputedCommitmentUtilizedAmountOk() (*string, bool)`

GetComputedCommitmentUtilizedAmountOk returns a tuple with the ComputedCommitmentUtilizedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedCommitmentUtilizedAmount

`func (o *TypesCommitmentInfo) SetComputedCommitmentUtilizedAmount(v string)`

SetComputedCommitmentUtilizedAmount sets ComputedCommitmentUtilizedAmount field to given value.

### HasComputedCommitmentUtilizedAmount

`func (o *TypesCommitmentInfo) HasComputedCommitmentUtilizedAmount() bool`

HasComputedCommitmentUtilizedAmount returns a boolean if a field has been set.

### GetComputedOverageAmount

`func (o *TypesCommitmentInfo) GetComputedOverageAmount() string`

GetComputedOverageAmount returns the ComputedOverageAmount field if non-nil, zero value otherwise.

### GetComputedOverageAmountOk

`func (o *TypesCommitmentInfo) GetComputedOverageAmountOk() (*string, bool)`

GetComputedOverageAmountOk returns a tuple with the ComputedOverageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedOverageAmount

`func (o *TypesCommitmentInfo) SetComputedOverageAmount(v string)`

SetComputedOverageAmount sets ComputedOverageAmount field to given value.

### HasComputedOverageAmount

`func (o *TypesCommitmentInfo) HasComputedOverageAmount() bool`

HasComputedOverageAmount returns a boolean if a field has been set.

### GetComputedTrueUpAmount

`func (o *TypesCommitmentInfo) GetComputedTrueUpAmount() string`

GetComputedTrueUpAmount returns the ComputedTrueUpAmount field if non-nil, zero value otherwise.

### GetComputedTrueUpAmountOk

`func (o *TypesCommitmentInfo) GetComputedTrueUpAmountOk() (*string, bool)`

GetComputedTrueUpAmountOk returns a tuple with the ComputedTrueUpAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedTrueUpAmount

`func (o *TypesCommitmentInfo) SetComputedTrueUpAmount(v string)`

SetComputedTrueUpAmount sets ComputedTrueUpAmount field to given value.

### HasComputedTrueUpAmount

`func (o *TypesCommitmentInfo) HasComputedTrueUpAmount() bool`

HasComputedTrueUpAmount returns a boolean if a field has been set.

### GetIsWindowed

`func (o *TypesCommitmentInfo) GetIsWindowed() bool`

GetIsWindowed returns the IsWindowed field if non-nil, zero value otherwise.

### GetIsWindowedOk

`func (o *TypesCommitmentInfo) GetIsWindowedOk() (*bool, bool)`

GetIsWindowedOk returns a tuple with the IsWindowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowed

`func (o *TypesCommitmentInfo) SetIsWindowed(v bool)`

SetIsWindowed sets IsWindowed field to given value.

### HasIsWindowed

`func (o *TypesCommitmentInfo) HasIsWindowed() bool`

HasIsWindowed returns a boolean if a field has been set.

### GetOverageFactor

`func (o *TypesCommitmentInfo) GetOverageFactor() string`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *TypesCommitmentInfo) GetOverageFactorOk() (*string, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *TypesCommitmentInfo) SetOverageFactor(v string)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *TypesCommitmentInfo) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.

### GetQuantity

`func (o *TypesCommitmentInfo) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TypesCommitmentInfo) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TypesCommitmentInfo) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TypesCommitmentInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTrueUpEnabled

`func (o *TypesCommitmentInfo) GetTrueUpEnabled() bool`

GetTrueUpEnabled returns the TrueUpEnabled field if non-nil, zero value otherwise.

### GetTrueUpEnabledOk

`func (o *TypesCommitmentInfo) GetTrueUpEnabledOk() (*bool, bool)`

GetTrueUpEnabledOk returns a tuple with the TrueUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueUpEnabled

`func (o *TypesCommitmentInfo) SetTrueUpEnabled(v bool)`

SetTrueUpEnabled sets TrueUpEnabled field to given value.

### HasTrueUpEnabled

`func (o *TypesCommitmentInfo) HasTrueUpEnabled() bool`

HasTrueUpEnabled returns a boolean if a field has been set.

### GetType

`func (o *TypesCommitmentInfo) GetType() TypesCommitmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypesCommitmentInfo) GetTypeOk() (*TypesCommitmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypesCommitmentInfo) SetType(v TypesCommitmentType)`

SetType sets Type field to given value.

### HasType

`func (o *TypesCommitmentInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


