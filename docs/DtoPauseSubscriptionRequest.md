# DtoPauseSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PauseDays** | Pointer to **int32** |  | [optional] 
**PauseEnd** | Pointer to **string** |  | [optional] 
**PauseMode** | [**TypesPauseMode**](TypesPauseMode.md) |  | 
**PauseStart** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoPauseSubscriptionRequest

`func NewDtoPauseSubscriptionRequest(pauseMode TypesPauseMode, ) *DtoPauseSubscriptionRequest`

NewDtoPauseSubscriptionRequest instantiates a new DtoPauseSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPauseSubscriptionRequestWithDefaults

`func NewDtoPauseSubscriptionRequestWithDefaults() *DtoPauseSubscriptionRequest`

NewDtoPauseSubscriptionRequestWithDefaults instantiates a new DtoPauseSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DtoPauseSubscriptionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DtoPauseSubscriptionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DtoPauseSubscriptionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DtoPauseSubscriptionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoPauseSubscriptionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoPauseSubscriptionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoPauseSubscriptionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoPauseSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPauseDays

`func (o *DtoPauseSubscriptionRequest) GetPauseDays() int32`

GetPauseDays returns the PauseDays field if non-nil, zero value otherwise.

### GetPauseDaysOk

`func (o *DtoPauseSubscriptionRequest) GetPauseDaysOk() (*int32, bool)`

GetPauseDaysOk returns a tuple with the PauseDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDays

`func (o *DtoPauseSubscriptionRequest) SetPauseDays(v int32)`

SetPauseDays sets PauseDays field to given value.

### HasPauseDays

`func (o *DtoPauseSubscriptionRequest) HasPauseDays() bool`

HasPauseDays returns a boolean if a field has been set.

### GetPauseEnd

`func (o *DtoPauseSubscriptionRequest) GetPauseEnd() string`

GetPauseEnd returns the PauseEnd field if non-nil, zero value otherwise.

### GetPauseEndOk

`func (o *DtoPauseSubscriptionRequest) GetPauseEndOk() (*string, bool)`

GetPauseEndOk returns a tuple with the PauseEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseEnd

`func (o *DtoPauseSubscriptionRequest) SetPauseEnd(v string)`

SetPauseEnd sets PauseEnd field to given value.

### HasPauseEnd

`func (o *DtoPauseSubscriptionRequest) HasPauseEnd() bool`

HasPauseEnd returns a boolean if a field has been set.

### GetPauseMode

`func (o *DtoPauseSubscriptionRequest) GetPauseMode() TypesPauseMode`

GetPauseMode returns the PauseMode field if non-nil, zero value otherwise.

### GetPauseModeOk

`func (o *DtoPauseSubscriptionRequest) GetPauseModeOk() (*TypesPauseMode, bool)`

GetPauseModeOk returns a tuple with the PauseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseMode

`func (o *DtoPauseSubscriptionRequest) SetPauseMode(v TypesPauseMode)`

SetPauseMode sets PauseMode field to given value.


### GetPauseStart

`func (o *DtoPauseSubscriptionRequest) GetPauseStart() string`

GetPauseStart returns the PauseStart field if non-nil, zero value otherwise.

### GetPauseStartOk

`func (o *DtoPauseSubscriptionRequest) GetPauseStartOk() (*string, bool)`

GetPauseStartOk returns a tuple with the PauseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseStart

`func (o *DtoPauseSubscriptionRequest) SetPauseStart(v string)`

SetPauseStart sets PauseStart field to given value.

### HasPauseStart

`func (o *DtoPauseSubscriptionRequest) HasPauseStart() bool`

HasPauseStart returns a boolean if a field has been set.

### GetReason

`func (o *DtoPauseSubscriptionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoPauseSubscriptionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoPauseSubscriptionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoPauseSubscriptionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


