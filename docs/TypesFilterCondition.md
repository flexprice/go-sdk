# TypesFilterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to [**TypesDataType**](TypesDataType.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to [**TypesFilterOperatorType**](TypesFilterOperatorType.md) |  | [optional] 
**Value** | Pointer to [**GithubComFlexpriceFlexpriceInternalTypesValue**](GithubComFlexpriceFlexpriceInternalTypesValue.md) |  | [optional] 

## Methods

### NewTypesFilterCondition

`func NewTypesFilterCondition() *TypesFilterCondition`

NewTypesFilterCondition instantiates a new TypesFilterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesFilterConditionWithDefaults

`func NewTypesFilterConditionWithDefaults() *TypesFilterCondition`

NewTypesFilterConditionWithDefaults instantiates a new TypesFilterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TypesFilterCondition) GetDataType() TypesDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TypesFilterCondition) GetDataTypeOk() (*TypesDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TypesFilterCondition) SetDataType(v TypesDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *TypesFilterCondition) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetField

`func (o *TypesFilterCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TypesFilterCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TypesFilterCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *TypesFilterCondition) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOperator

`func (o *TypesFilterCondition) GetOperator() TypesFilterOperatorType`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TypesFilterCondition) GetOperatorOk() (*TypesFilterOperatorType, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TypesFilterCondition) SetOperator(v TypesFilterOperatorType)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *TypesFilterCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *TypesFilterCondition) GetValue() GithubComFlexpriceFlexpriceInternalTypesValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TypesFilterCondition) GetValueOk() (*GithubComFlexpriceFlexpriceInternalTypesValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TypesFilterCondition) SetValue(v GithubComFlexpriceFlexpriceInternalTypesValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *TypesFilterCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


