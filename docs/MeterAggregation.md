# MeterAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Field is the key in $event.properties on which the aggregation is to be applied For ex if the aggregation type is sum for API usage, the field could be \&quot;duration_ms\&quot; | [optional] 
**Type** | Pointer to [**TypesAggregationType**](TypesAggregationType.md) |  | [optional] 

## Methods

### NewMeterAggregation

`func NewMeterAggregation() *MeterAggregation`

NewMeterAggregation instantiates a new MeterAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterAggregationWithDefaults

`func NewMeterAggregationWithDefaults() *MeterAggregation`

NewMeterAggregationWithDefaults instantiates a new MeterAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MeterAggregation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MeterAggregation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MeterAggregation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MeterAggregation) HasField() bool`

HasField returns a boolean if a field has been set.

### GetType

`func (o *MeterAggregation) GetType() TypesAggregationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MeterAggregation) GetTypeOk() (*TypesAggregationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MeterAggregation) SetType(v TypesAggregationType)`

SetType sets Type field to given value.

### HasType

`func (o *MeterAggregation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


