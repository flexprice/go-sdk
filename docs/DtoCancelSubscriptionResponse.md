# DtoCancelSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationType** | Pointer to [**TypesCancellationType**](TypesCancellationType.md) |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** | Response metadata | [optional] 
**ProcessedAt** | Pointer to **string** |  | [optional] 
**ProrationDetails** | Pointer to [**[]DtoProrationDetail**](DtoProrationDetail.md) |  | [optional] 
**ProrationInvoice** | Pointer to [**DtoInvoiceResponse**](DtoInvoiceResponse.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | Basic cancellation info | [optional] 
**TotalCreditAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCancelSubscriptionResponse

`func NewDtoCancelSubscriptionResponse() *DtoCancelSubscriptionResponse`

NewDtoCancelSubscriptionResponse instantiates a new DtoCancelSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCancelSubscriptionResponseWithDefaults

`func NewDtoCancelSubscriptionResponseWithDefaults() *DtoCancelSubscriptionResponse`

NewDtoCancelSubscriptionResponseWithDefaults instantiates a new DtoCancelSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationType

`func (o *DtoCancelSubscriptionResponse) GetCancellationType() TypesCancellationType`

GetCancellationType returns the CancellationType field if non-nil, zero value otherwise.

### GetCancellationTypeOk

`func (o *DtoCancelSubscriptionResponse) GetCancellationTypeOk() (*TypesCancellationType, bool)`

GetCancellationTypeOk returns a tuple with the CancellationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationType

`func (o *DtoCancelSubscriptionResponse) SetCancellationType(v TypesCancellationType)`

SetCancellationType sets CancellationType field to given value.

### HasCancellationType

`func (o *DtoCancelSubscriptionResponse) HasCancellationType() bool`

HasCancellationType returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *DtoCancelSubscriptionResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *DtoCancelSubscriptionResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *DtoCancelSubscriptionResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *DtoCancelSubscriptionResponse) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetMessage

`func (o *DtoCancelSubscriptionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DtoCancelSubscriptionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DtoCancelSubscriptionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DtoCancelSubscriptionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProcessedAt

`func (o *DtoCancelSubscriptionResponse) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *DtoCancelSubscriptionResponse) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *DtoCancelSubscriptionResponse) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *DtoCancelSubscriptionResponse) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetProrationDetails

`func (o *DtoCancelSubscriptionResponse) GetProrationDetails() []DtoProrationDetail`

GetProrationDetails returns the ProrationDetails field if non-nil, zero value otherwise.

### GetProrationDetailsOk

`func (o *DtoCancelSubscriptionResponse) GetProrationDetailsOk() (*[]DtoProrationDetail, bool)`

GetProrationDetailsOk returns a tuple with the ProrationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDetails

`func (o *DtoCancelSubscriptionResponse) SetProrationDetails(v []DtoProrationDetail)`

SetProrationDetails sets ProrationDetails field to given value.

### HasProrationDetails

`func (o *DtoCancelSubscriptionResponse) HasProrationDetails() bool`

HasProrationDetails returns a boolean if a field has been set.

### GetProrationInvoice

`func (o *DtoCancelSubscriptionResponse) GetProrationInvoice() DtoInvoiceResponse`

GetProrationInvoice returns the ProrationInvoice field if non-nil, zero value otherwise.

### GetProrationInvoiceOk

`func (o *DtoCancelSubscriptionResponse) GetProrationInvoiceOk() (*DtoInvoiceResponse, bool)`

GetProrationInvoiceOk returns a tuple with the ProrationInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationInvoice

`func (o *DtoCancelSubscriptionResponse) SetProrationInvoice(v DtoInvoiceResponse)`

SetProrationInvoice sets ProrationInvoice field to given value.

### HasProrationInvoice

`func (o *DtoCancelSubscriptionResponse) HasProrationInvoice() bool`

HasProrationInvoice returns a boolean if a field has been set.

### GetReason

`func (o *DtoCancelSubscriptionResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoCancelSubscriptionResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoCancelSubscriptionResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoCancelSubscriptionResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCancelSubscriptionResponse) GetStatus() TypesSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCancelSubscriptionResponse) GetStatusOk() (*TypesSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCancelSubscriptionResponse) SetStatus(v TypesSubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCancelSubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCancelSubscriptionResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCancelSubscriptionResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCancelSubscriptionResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCancelSubscriptionResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTotalCreditAmount

`func (o *DtoCancelSubscriptionResponse) GetTotalCreditAmount() string`

GetTotalCreditAmount returns the TotalCreditAmount field if non-nil, zero value otherwise.

### GetTotalCreditAmountOk

`func (o *DtoCancelSubscriptionResponse) GetTotalCreditAmountOk() (*string, bool)`

GetTotalCreditAmountOk returns a tuple with the TotalCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditAmount

`func (o *DtoCancelSubscriptionResponse) SetTotalCreditAmount(v string)`

SetTotalCreditAmount sets TotalCreditAmount field to given value.

### HasTotalCreditAmount

`func (o *DtoCancelSubscriptionResponse) HasTotalCreditAmount() bool`

HasTotalCreditAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


