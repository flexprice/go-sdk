# TypesSortCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**TypesSortDirection**](TypesSortDirection.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesSortCondition

`func NewTypesSortCondition() *TypesSortCondition`

NewTypesSortCondition instantiates a new TypesSortCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesSortConditionWithDefaults

`func NewTypesSortConditionWithDefaults() *TypesSortCondition`

NewTypesSortConditionWithDefaults instantiates a new TypesSortCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *TypesSortCondition) GetDirection() TypesSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TypesSortCondition) GetDirectionOk() (*TypesSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TypesSortCondition) SetDirection(v TypesSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *TypesSortCondition) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetField

`func (o *TypesSortCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TypesSortCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TypesSortCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TypesSortCondition) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


