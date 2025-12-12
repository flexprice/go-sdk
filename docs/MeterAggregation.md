# MeterAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 
**Field** | Pointer to **string** | Field is the key in $event.properties on which the aggregation is to be applied For ex if the aggregation type is sum for API usage, the field could be \&quot;duration_ms\&quot; | [optional] 
**Multiplier** | Pointer to **string** | Multiplier is the multiplier for the aggregation For ex if the aggregation type is sum_with_multiplier for API usage, the multiplier could be 1000 to scale up by a factor of 1000. If not provided, it will be null. | [optional] 
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

### GetBucketSize

`func (o *MeterAggregation) GetBucketSize() TypesWindowSize`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *MeterAggregation) GetBucketSizeOk() (*TypesWindowSize, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *MeterAggregation) SetBucketSize(v TypesWindowSize)`

SetBucketSize sets BucketSize field to given value.

### HasBucketSize

`func (o *MeterAggregation) HasBucketSize() bool`

HasBucketSize returns a boolean if a field has been set.

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

### GetMultiplier

`func (o *MeterAggregation) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *MeterAggregation) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *MeterAggregation) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *MeterAggregation) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

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


