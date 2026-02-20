# TypesFailurePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorsErrorResponse**](ErrorsErrorResponse.md) |  | [optional] 
**FailurePointType** | Pointer to [**TypesFailurePointType**](TypesFailurePointType.md) |  | [optional] 

## Methods

### NewTypesFailurePoint

`func NewTypesFailurePoint() *TypesFailurePoint`

NewTypesFailurePoint instantiates a new TypesFailurePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesFailurePointWithDefaults

`func NewTypesFailurePointWithDefaults() *TypesFailurePoint`

NewTypesFailurePointWithDefaults instantiates a new TypesFailurePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TypesFailurePoint) GetError() ErrorsErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TypesFailurePoint) GetErrorOk() (*ErrorsErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TypesFailurePoint) SetError(v ErrorsErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *TypesFailurePoint) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFailurePointType

`func (o *TypesFailurePoint) GetFailurePointType() TypesFailurePointType`

GetFailurePointType returns the FailurePointType field if non-nil, zero value otherwise.

### GetFailurePointTypeOk

`func (o *TypesFailurePoint) GetFailurePointTypeOk() (*TypesFailurePointType, bool)`

GetFailurePointTypeOk returns a tuple with the FailurePointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePointType

`func (o *TypesFailurePoint) SetFailurePointType(v TypesFailurePointType)`

SetFailurePointType sets FailurePointType field to given value.

### HasFailurePointType

`func (o *TypesFailurePoint) HasFailurePointType() bool`

HasFailurePointType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


