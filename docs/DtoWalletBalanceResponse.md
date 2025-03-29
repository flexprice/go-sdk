# DtoWalletBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopupAmount** | Pointer to **float32** |  | [optional] 
**AutoTopupMinBalance** | Pointer to **float32** |  | [optional] 
**AutoTopupTrigger** | Pointer to [**TypesAutoTopupTrigger**](TypesAutoTopupTrigger.md) |  | [optional] 
**Balance** | Pointer to **float32** |  | [optional] 
**BalanceUpdatedAt** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**TypesWalletConfig**](TypesWalletConfig.md) |  | [optional] 
**ConversionRate** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditBalance** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrentPeriodUsage** | Pointer to **float32** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RealTimeBalance** | Pointer to **float32** |  | [optional] 
**RealTimeCreditBalance** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UnpaidInvoiceAmount** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**WalletStatus** | Pointer to [**TypesWalletStatus**](TypesWalletStatus.md) |  | [optional] 
**WalletType** | Pointer to [**TypesWalletType**](TypesWalletType.md) |  | [optional] 

## Methods

### NewDtoWalletBalanceResponse

`func NewDtoWalletBalanceResponse() *DtoWalletBalanceResponse`

NewDtoWalletBalanceResponse instantiates a new DtoWalletBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoWalletBalanceResponseWithDefaults

`func NewDtoWalletBalanceResponseWithDefaults() *DtoWalletBalanceResponse`

NewDtoWalletBalanceResponseWithDefaults instantiates a new DtoWalletBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTopupAmount

`func (o *DtoWalletBalanceResponse) GetAutoTopupAmount() float32`

GetAutoTopupAmount returns the AutoTopupAmount field if non-nil, zero value otherwise.

### GetAutoTopupAmountOk

`func (o *DtoWalletBalanceResponse) GetAutoTopupAmountOk() (*float32, bool)`

GetAutoTopupAmountOk returns a tuple with the AutoTopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupAmount

`func (o *DtoWalletBalanceResponse) SetAutoTopupAmount(v float32)`

SetAutoTopupAmount sets AutoTopupAmount field to given value.

### HasAutoTopupAmount

`func (o *DtoWalletBalanceResponse) HasAutoTopupAmount() bool`

HasAutoTopupAmount returns a boolean if a field has been set.

### GetAutoTopupMinBalance

`func (o *DtoWalletBalanceResponse) GetAutoTopupMinBalance() float32`

GetAutoTopupMinBalance returns the AutoTopupMinBalance field if non-nil, zero value otherwise.

### GetAutoTopupMinBalanceOk

`func (o *DtoWalletBalanceResponse) GetAutoTopupMinBalanceOk() (*float32, bool)`

GetAutoTopupMinBalanceOk returns a tuple with the AutoTopupMinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupMinBalance

`func (o *DtoWalletBalanceResponse) SetAutoTopupMinBalance(v float32)`

SetAutoTopupMinBalance sets AutoTopupMinBalance field to given value.

### HasAutoTopupMinBalance

`func (o *DtoWalletBalanceResponse) HasAutoTopupMinBalance() bool`

HasAutoTopupMinBalance returns a boolean if a field has been set.

### GetAutoTopupTrigger

`func (o *DtoWalletBalanceResponse) GetAutoTopupTrigger() TypesAutoTopupTrigger`

GetAutoTopupTrigger returns the AutoTopupTrigger field if non-nil, zero value otherwise.

### GetAutoTopupTriggerOk

`func (o *DtoWalletBalanceResponse) GetAutoTopupTriggerOk() (*TypesAutoTopupTrigger, bool)`

GetAutoTopupTriggerOk returns a tuple with the AutoTopupTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupTrigger

`func (o *DtoWalletBalanceResponse) SetAutoTopupTrigger(v TypesAutoTopupTrigger)`

SetAutoTopupTrigger sets AutoTopupTrigger field to given value.

### HasAutoTopupTrigger

`func (o *DtoWalletBalanceResponse) HasAutoTopupTrigger() bool`

HasAutoTopupTrigger returns a boolean if a field has been set.

### GetBalance

`func (o *DtoWalletBalanceResponse) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DtoWalletBalanceResponse) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DtoWalletBalanceResponse) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *DtoWalletBalanceResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBalanceUpdatedAt

`func (o *DtoWalletBalanceResponse) GetBalanceUpdatedAt() string`

GetBalanceUpdatedAt returns the BalanceUpdatedAt field if non-nil, zero value otherwise.

### GetBalanceUpdatedAtOk

`func (o *DtoWalletBalanceResponse) GetBalanceUpdatedAtOk() (*string, bool)`

GetBalanceUpdatedAtOk returns a tuple with the BalanceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUpdatedAt

`func (o *DtoWalletBalanceResponse) SetBalanceUpdatedAt(v string)`

SetBalanceUpdatedAt sets BalanceUpdatedAt field to given value.

### HasBalanceUpdatedAt

`func (o *DtoWalletBalanceResponse) HasBalanceUpdatedAt() bool`

HasBalanceUpdatedAt returns a boolean if a field has been set.

### GetConfig

`func (o *DtoWalletBalanceResponse) GetConfig() TypesWalletConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DtoWalletBalanceResponse) GetConfigOk() (*TypesWalletConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DtoWalletBalanceResponse) SetConfig(v TypesWalletConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DtoWalletBalanceResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConversionRate

`func (o *DtoWalletBalanceResponse) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *DtoWalletBalanceResponse) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *DtoWalletBalanceResponse) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *DtoWalletBalanceResponse) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoWalletBalanceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoWalletBalanceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoWalletBalanceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoWalletBalanceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoWalletBalanceResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoWalletBalanceResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoWalletBalanceResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoWalletBalanceResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditBalance

`func (o *DtoWalletBalanceResponse) GetCreditBalance() float32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *DtoWalletBalanceResponse) GetCreditBalanceOk() (*float32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *DtoWalletBalanceResponse) SetCreditBalance(v float32)`

SetCreditBalance sets CreditBalance field to given value.

### HasCreditBalance

`func (o *DtoWalletBalanceResponse) HasCreditBalance() bool`

HasCreditBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoWalletBalanceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoWalletBalanceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoWalletBalanceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoWalletBalanceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodUsage

`func (o *DtoWalletBalanceResponse) GetCurrentPeriodUsage() float32`

GetCurrentPeriodUsage returns the CurrentPeriodUsage field if non-nil, zero value otherwise.

### GetCurrentPeriodUsageOk

`func (o *DtoWalletBalanceResponse) GetCurrentPeriodUsageOk() (*float32, bool)`

GetCurrentPeriodUsageOk returns a tuple with the CurrentPeriodUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodUsage

`func (o *DtoWalletBalanceResponse) SetCurrentPeriodUsage(v float32)`

SetCurrentPeriodUsage sets CurrentPeriodUsage field to given value.

### HasCurrentPeriodUsage

`func (o *DtoWalletBalanceResponse) HasCurrentPeriodUsage() bool`

HasCurrentPeriodUsage returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoWalletBalanceResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoWalletBalanceResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoWalletBalanceResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoWalletBalanceResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *DtoWalletBalanceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoWalletBalanceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoWalletBalanceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoWalletBalanceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoWalletBalanceResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoWalletBalanceResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoWalletBalanceResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoWalletBalanceResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoWalletBalanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoWalletBalanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoWalletBalanceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoWalletBalanceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoWalletBalanceResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoWalletBalanceResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoWalletBalanceResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoWalletBalanceResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoWalletBalanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoWalletBalanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoWalletBalanceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoWalletBalanceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealTimeBalance

`func (o *DtoWalletBalanceResponse) GetRealTimeBalance() float32`

GetRealTimeBalance returns the RealTimeBalance field if non-nil, zero value otherwise.

### GetRealTimeBalanceOk

`func (o *DtoWalletBalanceResponse) GetRealTimeBalanceOk() (*float32, bool)`

GetRealTimeBalanceOk returns a tuple with the RealTimeBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimeBalance

`func (o *DtoWalletBalanceResponse) SetRealTimeBalance(v float32)`

SetRealTimeBalance sets RealTimeBalance field to given value.

### HasRealTimeBalance

`func (o *DtoWalletBalanceResponse) HasRealTimeBalance() bool`

HasRealTimeBalance returns a boolean if a field has been set.

### GetRealTimeCreditBalance

`func (o *DtoWalletBalanceResponse) GetRealTimeCreditBalance() float32`

GetRealTimeCreditBalance returns the RealTimeCreditBalance field if non-nil, zero value otherwise.

### GetRealTimeCreditBalanceOk

`func (o *DtoWalletBalanceResponse) GetRealTimeCreditBalanceOk() (*float32, bool)`

GetRealTimeCreditBalanceOk returns a tuple with the RealTimeCreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimeCreditBalance

`func (o *DtoWalletBalanceResponse) SetRealTimeCreditBalance(v float32)`

SetRealTimeCreditBalance sets RealTimeCreditBalance field to given value.

### HasRealTimeCreditBalance

`func (o *DtoWalletBalanceResponse) HasRealTimeCreditBalance() bool`

HasRealTimeCreditBalance returns a boolean if a field has been set.

### GetStatus

`func (o *DtoWalletBalanceResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoWalletBalanceResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoWalletBalanceResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoWalletBalanceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoWalletBalanceResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoWalletBalanceResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoWalletBalanceResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoWalletBalanceResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUnpaidInvoiceAmount

`func (o *DtoWalletBalanceResponse) GetUnpaidInvoiceAmount() float32`

GetUnpaidInvoiceAmount returns the UnpaidInvoiceAmount field if non-nil, zero value otherwise.

### GetUnpaidInvoiceAmountOk

`func (o *DtoWalletBalanceResponse) GetUnpaidInvoiceAmountOk() (*float32, bool)`

GetUnpaidInvoiceAmountOk returns a tuple with the UnpaidInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidInvoiceAmount

`func (o *DtoWalletBalanceResponse) SetUnpaidInvoiceAmount(v float32)`

SetUnpaidInvoiceAmount sets UnpaidInvoiceAmount field to given value.

### HasUnpaidInvoiceAmount

`func (o *DtoWalletBalanceResponse) HasUnpaidInvoiceAmount() bool`

HasUnpaidInvoiceAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoWalletBalanceResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoWalletBalanceResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoWalletBalanceResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoWalletBalanceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoWalletBalanceResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoWalletBalanceResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoWalletBalanceResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoWalletBalanceResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWalletStatus

`func (o *DtoWalletBalanceResponse) GetWalletStatus() TypesWalletStatus`

GetWalletStatus returns the WalletStatus field if non-nil, zero value otherwise.

### GetWalletStatusOk

`func (o *DtoWalletBalanceResponse) GetWalletStatusOk() (*TypesWalletStatus, bool)`

GetWalletStatusOk returns a tuple with the WalletStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletStatus

`func (o *DtoWalletBalanceResponse) SetWalletStatus(v TypesWalletStatus)`

SetWalletStatus sets WalletStatus field to given value.

### HasWalletStatus

`func (o *DtoWalletBalanceResponse) HasWalletStatus() bool`

HasWalletStatus returns a boolean if a field has been set.

### GetWalletType

`func (o *DtoWalletBalanceResponse) GetWalletType() TypesWalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *DtoWalletBalanceResponse) GetWalletTypeOk() (*TypesWalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *DtoWalletBalanceResponse) SetWalletType(v TypesWalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *DtoWalletBalanceResponse) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


