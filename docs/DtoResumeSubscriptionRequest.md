# DtoResumeSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | Whether to perform a dry run @Description If true, validates the request and shows impact without actually resuming the subscription @Example false | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata as key-value pairs @Description Optional metadata for storing additional information about the resume operation @Example {\&quot;resumed_by\&quot;: \&quot;admin\&quot;, \&quot;reason\&quot;: \&quot;issue_resolved\&quot;} | [optional] 
**ResumeMode** | [**TypesResumeMode**](TypesResumeMode.md) |  | 

## Methods

### NewDtoResumeSubscriptionRequest

`func NewDtoResumeSubscriptionRequest(resumeMode TypesResumeMode, ) *DtoResumeSubscriptionRequest`

NewDtoResumeSubscriptionRequest instantiates a new DtoResumeSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoResumeSubscriptionRequestWithDefaults

`func NewDtoResumeSubscriptionRequestWithDefaults() *DtoResumeSubscriptionRequest`

NewDtoResumeSubscriptionRequestWithDefaults instantiates a new DtoResumeSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *DtoResumeSubscriptionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DtoResumeSubscriptionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DtoResumeSubscriptionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DtoResumeSubscriptionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoResumeSubscriptionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoResumeSubscriptionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoResumeSubscriptionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoResumeSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResumeMode

`func (o *DtoResumeSubscriptionRequest) GetResumeMode() TypesResumeMode`

GetResumeMode returns the ResumeMode field if non-nil, zero value otherwise.

### GetResumeModeOk

`func (o *DtoResumeSubscriptionRequest) GetResumeModeOk() (*TypesResumeMode, bool)`

GetResumeModeOk returns a tuple with the ResumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeMode

`func (o *DtoResumeSubscriptionRequest) SetResumeMode(v TypesResumeMode)`

SetResumeMode sets ResumeMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


