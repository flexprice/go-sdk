# TypesAlertThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**TypesAlertCondition**](TypesAlertCondition.md) |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 

## Methods

### NewTypesAlertThreshold

`func NewTypesAlertThreshold() *TypesAlertThreshold`

NewTypesAlertThreshold instantiates a new TypesAlertThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAlertThresholdWithDefaults

`func NewTypesAlertThresholdWithDefaults() *TypesAlertThreshold`

NewTypesAlertThresholdWithDefaults instantiates a new TypesAlertThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *TypesAlertThreshold) GetCondition() TypesAlertCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TypesAlertThreshold) GetConditionOk() (*TypesAlertCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TypesAlertThreshold) SetCondition(v TypesAlertCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TypesAlertThreshold) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetThreshold

`func (o *TypesAlertThreshold) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TypesAlertThreshold) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TypesAlertThreshold) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *TypesAlertThreshold) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


