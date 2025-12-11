# DtoListConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]DtoConnectionResponse**](DtoConnectionResponse.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewDtoListConnectionsResponse

`func NewDtoListConnectionsResponse() *DtoListConnectionsResponse`

NewDtoListConnectionsResponse instantiates a new DtoListConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoListConnectionsResponseWithDefaults

`func NewDtoListConnectionsResponseWithDefaults() *DtoListConnectionsResponse`

NewDtoListConnectionsResponseWithDefaults instantiates a new DtoListConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *DtoListConnectionsResponse) GetConnections() []DtoConnectionResponse`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *DtoListConnectionsResponse) GetConnectionsOk() (*[]DtoConnectionResponse, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *DtoListConnectionsResponse) SetConnections(v []DtoConnectionResponse)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *DtoListConnectionsResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetLimit

`func (o *DtoListConnectionsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DtoListConnectionsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DtoListConnectionsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DtoListConnectionsResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DtoListConnectionsResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DtoListConnectionsResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DtoListConnectionsResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DtoListConnectionsResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *DtoListConnectionsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoListConnectionsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoListConnectionsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoListConnectionsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


