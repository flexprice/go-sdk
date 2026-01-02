# DtoLineItemCommitmentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitmentAmount** | Pointer to **float32** | CommitmentAmount is the minimum amount committed for this line item | [optional] 
**CommitmentQuantity** | Pointer to **float32** | CommitmentQuantity is the minimum quantity committed for this line item | [optional] 
**CommitmentType** | Pointer to [**TypesCommitmentType**](TypesCommitmentType.md) |  | [optional] 
**EnableTrueUp** | Pointer to **bool** | EnableTrueUp determines if true-up fee should be applied when usage is below commitment | [optional] 
**IsWindowCommitment** | Pointer to **bool** | IsWindowCommitment determines if commitment is applied per window (e.g., per day) rather than per billing period | [optional] 
**OverageFactor** | Pointer to **float32** | OverageFactor is a multiplier applied to usage beyond the commitment | [optional] 

## Methods

### NewDtoLineItemCommitmentConfig

`func NewDtoLineItemCommitmentConfig() *DtoLineItemCommitmentConfig`

NewDtoLineItemCommitmentConfig instantiates a new DtoLineItemCommitmentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoLineItemCommitmentConfigWithDefaults

`func NewDtoLineItemCommitmentConfigWithDefaults() *DtoLineItemCommitmentConfig`

NewDtoLineItemCommitmentConfigWithDefaults instantiates a new DtoLineItemCommitmentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitmentAmount

`func (o *DtoLineItemCommitmentConfig) GetCommitmentAmount() float32`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoLineItemCommitmentConfig) GetCommitmentAmountOk() (*float32, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoLineItemCommitmentConfig) SetCommitmentAmount(v float32)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoLineItemCommitmentConfig) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentQuantity

`func (o *DtoLineItemCommitmentConfig) GetCommitmentQuantity() float32`

GetCommitmentQuantity returns the CommitmentQuantity field if non-nil, zero value otherwise.

### GetCommitmentQuantityOk

`func (o *DtoLineItemCommitmentConfig) GetCommitmentQuantityOk() (*float32, bool)`

GetCommitmentQuantityOk returns a tuple with the CommitmentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentQuantity

`func (o *DtoLineItemCommitmentConfig) SetCommitmentQuantity(v float32)`

SetCommitmentQuantity sets CommitmentQuantity field to given value.

### HasCommitmentQuantity

`func (o *DtoLineItemCommitmentConfig) HasCommitmentQuantity() bool`

HasCommitmentQuantity returns a boolean if a field has been set.

### GetCommitmentType

`func (o *DtoLineItemCommitmentConfig) GetCommitmentType() TypesCommitmentType`

GetCommitmentType returns the CommitmentType field if non-nil, zero value otherwise.

### GetCommitmentTypeOk

`func (o *DtoLineItemCommitmentConfig) GetCommitmentTypeOk() (*TypesCommitmentType, bool)`

GetCommitmentTypeOk returns a tuple with the CommitmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentType

`func (o *DtoLineItemCommitmentConfig) SetCommitmentType(v TypesCommitmentType)`

SetCommitmentType sets CommitmentType field to given value.

### HasCommitmentType

`func (o *DtoLineItemCommitmentConfig) HasCommitmentType() bool`

HasCommitmentType returns a boolean if a field has been set.

### GetEnableTrueUp

`func (o *DtoLineItemCommitmentConfig) GetEnableTrueUp() bool`

GetEnableTrueUp returns the EnableTrueUp field if non-nil, zero value otherwise.

### GetEnableTrueUpOk

`func (o *DtoLineItemCommitmentConfig) GetEnableTrueUpOk() (*bool, bool)`

GetEnableTrueUpOk returns a tuple with the EnableTrueUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrueUp

`func (o *DtoLineItemCommitmentConfig) SetEnableTrueUp(v bool)`

SetEnableTrueUp sets EnableTrueUp field to given value.

### HasEnableTrueUp

`func (o *DtoLineItemCommitmentConfig) HasEnableTrueUp() bool`

HasEnableTrueUp returns a boolean if a field has been set.

### GetIsWindowCommitment

`func (o *DtoLineItemCommitmentConfig) GetIsWindowCommitment() bool`

GetIsWindowCommitment returns the IsWindowCommitment field if non-nil, zero value otherwise.

### GetIsWindowCommitmentOk

`func (o *DtoLineItemCommitmentConfig) GetIsWindowCommitmentOk() (*bool, bool)`

GetIsWindowCommitmentOk returns a tuple with the IsWindowCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindowCommitment

`func (o *DtoLineItemCommitmentConfig) SetIsWindowCommitment(v bool)`

SetIsWindowCommitment sets IsWindowCommitment field to given value.

### HasIsWindowCommitment

`func (o *DtoLineItemCommitmentConfig) HasIsWindowCommitment() bool`

HasIsWindowCommitment returns a boolean if a field has been set.

### GetOverageFactor

`func (o *DtoLineItemCommitmentConfig) GetOverageFactor() float32`

GetOverageFactor returns the OverageFactor field if non-nil, zero value otherwise.

### GetOverageFactorOk

`func (o *DtoLineItemCommitmentConfig) GetOverageFactorOk() (*float32, bool)`

GetOverageFactorOk returns a tuple with the OverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageFactor

`func (o *DtoLineItemCommitmentConfig) SetOverageFactor(v float32)`

SetOverageFactor sets OverageFactor field to given value.

### HasOverageFactor

`func (o *DtoLineItemCommitmentConfig) HasOverageFactor() bool`

HasOverageFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


