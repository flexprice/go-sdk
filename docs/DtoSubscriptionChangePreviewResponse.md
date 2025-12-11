# DtoSubscriptionChangePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to [**TypesSubscriptionChangeType**](TypesSubscriptionChangeType.md) |  | [optional] 
**CurrentPlan** | Pointer to [**DtoPlanSummary**](DtoPlanSummary.md) |  | [optional] 
**EffectiveDate** | Pointer to **string** | effective_date is when the change would take effect | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata from the request | [optional] 
**NewBillingCycle** | Pointer to [**DtoBillingCycleInfo**](DtoBillingCycleInfo.md) |  | [optional] 
**NextInvoicePreview** | Pointer to [**DtoInvoicePreview**](DtoInvoicePreview.md) |  | [optional] 
**ProrationDetails** | Pointer to [**DtoProrationDetails**](DtoProrationDetails.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the ID of the subscription being changed | [optional] 
**TargetPlan** | Pointer to [**DtoPlanSummary**](DtoPlanSummary.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | warnings contains any warnings about the change | [optional] 

## Methods

### NewDtoSubscriptionChangePreviewResponse

`func NewDtoSubscriptionChangePreviewResponse() *DtoSubscriptionChangePreviewResponse`

NewDtoSubscriptionChangePreviewResponse instantiates a new DtoSubscriptionChangePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionChangePreviewResponseWithDefaults

`func NewDtoSubscriptionChangePreviewResponseWithDefaults() *DtoSubscriptionChangePreviewResponse`

NewDtoSubscriptionChangePreviewResponseWithDefaults instantiates a new DtoSubscriptionChangePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *DtoSubscriptionChangePreviewResponse) GetChangeType() TypesSubscriptionChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *DtoSubscriptionChangePreviewResponse) GetChangeTypeOk() (*TypesSubscriptionChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *DtoSubscriptionChangePreviewResponse) SetChangeType(v TypesSubscriptionChangeType)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *DtoSubscriptionChangePreviewResponse) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetCurrentPlan

`func (o *DtoSubscriptionChangePreviewResponse) GetCurrentPlan() DtoPlanSummary`

GetCurrentPlan returns the CurrentPlan field if non-nil, zero value otherwise.

### GetCurrentPlanOk

`func (o *DtoSubscriptionChangePreviewResponse) GetCurrentPlanOk() (*DtoPlanSummary, bool)`

GetCurrentPlanOk returns a tuple with the CurrentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlan

`func (o *DtoSubscriptionChangePreviewResponse) SetCurrentPlan(v DtoPlanSummary)`

SetCurrentPlan sets CurrentPlan field to given value.

### HasCurrentPlan

`func (o *DtoSubscriptionChangePreviewResponse) HasCurrentPlan() bool`

HasCurrentPlan returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *DtoSubscriptionChangePreviewResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *DtoSubscriptionChangePreviewResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *DtoSubscriptionChangePreviewResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *DtoSubscriptionChangePreviewResponse) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionChangePreviewResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionChangePreviewResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionChangePreviewResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionChangePreviewResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNewBillingCycle

`func (o *DtoSubscriptionChangePreviewResponse) GetNewBillingCycle() DtoBillingCycleInfo`

GetNewBillingCycle returns the NewBillingCycle field if non-nil, zero value otherwise.

### GetNewBillingCycleOk

`func (o *DtoSubscriptionChangePreviewResponse) GetNewBillingCycleOk() (*DtoBillingCycleInfo, bool)`

GetNewBillingCycleOk returns a tuple with the NewBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewBillingCycle

`func (o *DtoSubscriptionChangePreviewResponse) SetNewBillingCycle(v DtoBillingCycleInfo)`

SetNewBillingCycle sets NewBillingCycle field to given value.

### HasNewBillingCycle

`func (o *DtoSubscriptionChangePreviewResponse) HasNewBillingCycle() bool`

HasNewBillingCycle returns a boolean if a field has been set.

### GetNextInvoicePreview

`func (o *DtoSubscriptionChangePreviewResponse) GetNextInvoicePreview() DtoInvoicePreview`

GetNextInvoicePreview returns the NextInvoicePreview field if non-nil, zero value otherwise.

### GetNextInvoicePreviewOk

`func (o *DtoSubscriptionChangePreviewResponse) GetNextInvoicePreviewOk() (*DtoInvoicePreview, bool)`

GetNextInvoicePreviewOk returns a tuple with the NextInvoicePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoicePreview

`func (o *DtoSubscriptionChangePreviewResponse) SetNextInvoicePreview(v DtoInvoicePreview)`

SetNextInvoicePreview sets NextInvoicePreview field to given value.

### HasNextInvoicePreview

`func (o *DtoSubscriptionChangePreviewResponse) HasNextInvoicePreview() bool`

HasNextInvoicePreview returns a boolean if a field has been set.

### GetProrationDetails

`func (o *DtoSubscriptionChangePreviewResponse) GetProrationDetails() DtoProrationDetails`

GetProrationDetails returns the ProrationDetails field if non-nil, zero value otherwise.

### GetProrationDetailsOk

`func (o *DtoSubscriptionChangePreviewResponse) GetProrationDetailsOk() (*DtoProrationDetails, bool)`

GetProrationDetailsOk returns a tuple with the ProrationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDetails

`func (o *DtoSubscriptionChangePreviewResponse) SetProrationDetails(v DtoProrationDetails)`

SetProrationDetails sets ProrationDetails field to given value.

### HasProrationDetails

`func (o *DtoSubscriptionChangePreviewResponse) HasProrationDetails() bool`

HasProrationDetails returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoSubscriptionChangePreviewResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoSubscriptionChangePreviewResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoSubscriptionChangePreviewResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoSubscriptionChangePreviewResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTargetPlan

`func (o *DtoSubscriptionChangePreviewResponse) GetTargetPlan() DtoPlanSummary`

GetTargetPlan returns the TargetPlan field if non-nil, zero value otherwise.

### GetTargetPlanOk

`func (o *DtoSubscriptionChangePreviewResponse) GetTargetPlanOk() (*DtoPlanSummary, bool)`

GetTargetPlanOk returns a tuple with the TargetPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlan

`func (o *DtoSubscriptionChangePreviewResponse) SetTargetPlan(v DtoPlanSummary)`

SetTargetPlan sets TargetPlan field to given value.

### HasTargetPlan

`func (o *DtoSubscriptionChangePreviewResponse) HasTargetPlan() bool`

HasTargetPlan returns a boolean if a field has been set.

### GetWarnings

`func (o *DtoSubscriptionChangePreviewResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DtoSubscriptionChangePreviewResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DtoSubscriptionChangePreviewResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DtoSubscriptionChangePreviewResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


