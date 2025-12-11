# DtoSubscriptionChangeExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to [**TypesSubscriptionChangeType**](TypesSubscriptionChangeType.md) |  | [optional] 
**CreditGrants** | Pointer to [**[]DtoCreditGrantResponse**](DtoCreditGrantResponse.md) | credit_grants contains any credit grants created for proration credits | [optional] 
**EffectiveDate** | Pointer to **string** | effective_date is when the change took effect | [optional] 
**Invoice** | Pointer to [**DtoInvoiceResponse**](DtoInvoiceResponse.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata from the request | [optional] 
**NewSubscription** | Pointer to [**DtoSubscriptionSummary**](DtoSubscriptionSummary.md) |  | [optional] 
**OldSubscription** | Pointer to [**DtoSubscriptionSummary**](DtoSubscriptionSummary.md) |  | [optional] 
**ProrationApplied** | Pointer to [**DtoProrationDetails**](DtoProrationDetails.md) |  | [optional] 

## Methods

### NewDtoSubscriptionChangeExecuteResponse

`func NewDtoSubscriptionChangeExecuteResponse() *DtoSubscriptionChangeExecuteResponse`

NewDtoSubscriptionChangeExecuteResponse instantiates a new DtoSubscriptionChangeExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionChangeExecuteResponseWithDefaults

`func NewDtoSubscriptionChangeExecuteResponseWithDefaults() *DtoSubscriptionChangeExecuteResponse`

NewDtoSubscriptionChangeExecuteResponseWithDefaults instantiates a new DtoSubscriptionChangeExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *DtoSubscriptionChangeExecuteResponse) GetChangeType() TypesSubscriptionChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetChangeTypeOk() (*TypesSubscriptionChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *DtoSubscriptionChangeExecuteResponse) SetChangeType(v TypesSubscriptionChangeType)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *DtoSubscriptionChangeExecuteResponse) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetCreditGrants

`func (o *DtoSubscriptionChangeExecuteResponse) GetCreditGrants() []DtoCreditGrantResponse`

GetCreditGrants returns the CreditGrants field if non-nil, zero value otherwise.

### GetCreditGrantsOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetCreditGrantsOk() (*[]DtoCreditGrantResponse, bool)`

GetCreditGrantsOk returns a tuple with the CreditGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrants

`func (o *DtoSubscriptionChangeExecuteResponse) SetCreditGrants(v []DtoCreditGrantResponse)`

SetCreditGrants sets CreditGrants field to given value.

### HasCreditGrants

`func (o *DtoSubscriptionChangeExecuteResponse) HasCreditGrants() bool`

HasCreditGrants returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *DtoSubscriptionChangeExecuteResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *DtoSubscriptionChangeExecuteResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *DtoSubscriptionChangeExecuteResponse) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetInvoice

`func (o *DtoSubscriptionChangeExecuteResponse) GetInvoice() DtoInvoiceResponse`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetInvoiceOk() (*DtoInvoiceResponse, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *DtoSubscriptionChangeExecuteResponse) SetInvoice(v DtoInvoiceResponse)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *DtoSubscriptionChangeExecuteResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionChangeExecuteResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionChangeExecuteResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionChangeExecuteResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNewSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) GetNewSubscription() DtoSubscriptionSummary`

GetNewSubscription returns the NewSubscription field if non-nil, zero value otherwise.

### GetNewSubscriptionOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetNewSubscriptionOk() (*DtoSubscriptionSummary, bool)`

GetNewSubscriptionOk returns a tuple with the NewSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) SetNewSubscription(v DtoSubscriptionSummary)`

SetNewSubscription sets NewSubscription field to given value.

### HasNewSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) HasNewSubscription() bool`

HasNewSubscription returns a boolean if a field has been set.

### GetOldSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) GetOldSubscription() DtoSubscriptionSummary`

GetOldSubscription returns the OldSubscription field if non-nil, zero value otherwise.

### GetOldSubscriptionOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetOldSubscriptionOk() (*DtoSubscriptionSummary, bool)`

GetOldSubscriptionOk returns a tuple with the OldSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) SetOldSubscription(v DtoSubscriptionSummary)`

SetOldSubscription sets OldSubscription field to given value.

### HasOldSubscription

`func (o *DtoSubscriptionChangeExecuteResponse) HasOldSubscription() bool`

HasOldSubscription returns a boolean if a field has been set.

### GetProrationApplied

`func (o *DtoSubscriptionChangeExecuteResponse) GetProrationApplied() DtoProrationDetails`

GetProrationApplied returns the ProrationApplied field if non-nil, zero value otherwise.

### GetProrationAppliedOk

`func (o *DtoSubscriptionChangeExecuteResponse) GetProrationAppliedOk() (*DtoProrationDetails, bool)`

GetProrationAppliedOk returns a tuple with the ProrationApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationApplied

`func (o *DtoSubscriptionChangeExecuteResponse) SetProrationApplied(v DtoProrationDetails)`

SetProrationApplied sets ProrationApplied field to given value.

### HasProrationApplied

`func (o *DtoSubscriptionChangeExecuteResponse) HasProrationApplied() bool`

HasProrationApplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


