# MeterFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the key for the filter from $event.properties Currently we support only first level keys in the properties and not nested keys | [optional] 
**Values** | Pointer to **[]string** | Values are the possible values for the filter to be considered for the meter For ex \&quot;model_name\&quot; could have values \&quot;o1-mini\&quot;, \&quot;gpt-4o\&quot; etc | [optional] 

## Methods

### NewMeterFilter

`func NewMeterFilter() *MeterFilter`

NewMeterFilter instantiates a new MeterFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterFilterWithDefaults

`func NewMeterFilterWithDefaults() *MeterFilter`

NewMeterFilterWithDefaults instantiates a new MeterFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MeterFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MeterFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MeterFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MeterFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValues

`func (o *MeterFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MeterFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MeterFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *MeterFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


