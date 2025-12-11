# TypesEntitySyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | Pointer to **bool** | Inbound from external provider to FlexPrice | [optional] 
**Outbound** | Pointer to **bool** | Outbound from FlexPrice to external provider | [optional] 

## Methods

### NewTypesEntitySyncConfig

`func NewTypesEntitySyncConfig() *TypesEntitySyncConfig`

NewTypesEntitySyncConfig instantiates a new TypesEntitySyncConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesEntitySyncConfigWithDefaults

`func NewTypesEntitySyncConfigWithDefaults() *TypesEntitySyncConfig`

NewTypesEntitySyncConfigWithDefaults instantiates a new TypesEntitySyncConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *TypesEntitySyncConfig) GetInbound() bool`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *TypesEntitySyncConfig) GetInboundOk() (*bool, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *TypesEntitySyncConfig) SetInbound(v bool)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *TypesEntitySyncConfig) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *TypesEntitySyncConfig) GetOutbound() bool`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *TypesEntitySyncConfig) GetOutboundOk() (*bool, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *TypesEntitySyncConfig) SetOutbound(v bool)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *TypesEntitySyncConfig) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


