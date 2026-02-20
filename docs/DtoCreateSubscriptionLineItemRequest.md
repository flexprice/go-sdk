# DtoCreateSubscriptionLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitmentAmount** | Pointer to **float32** | Commitment fields | [optional] 
**CommitmentDuration** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 
**CommitmentOverageFactor** | Pointer to **float32** |  | [optional] 
**CommitmentQuantity** | Pointer to **float32** |  | [optional] 
**CommitmentTrueUpEnabled** | Pointer to **bool** |  | [optional] 
**CommitmentType** | Pointer to [**TypesCommitmentType**](TypesCommitmentType.md) |  | [optional] 
**CommitmentWindowed** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Price** | Pointer to [**DtoSubscriptionPriceCreateRequest**](DtoSubscriptionPriceCreateRequest.md) |  | [optional] 
**PriceId** | Pointer to **string** | PriceID references an existing price (plan, addon, or subscription-scoped). Exactly one of price_id or price must be set. | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**SubscriptionPhaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateSubscriptionLineItemRequest

`func NewDtoCreateSubscriptionLineItemRequest() *DtoCreateSubscriptionLineItemRequest`

NewDtoCreateSubscriptionLineItemRequest instantiates a new DtoCreateSubscriptionLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateSubscriptionLineItemRequestWithDefaults

`func NewDtoCreateSubscriptionLineItemRequestWithDefaults() *DtoCreateSubscriptionLineItemRequest`

NewDtoCreateSubscriptionLineItemRequestWithDefaults instantiates a new DtoCreateSubscriptionLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitmentAmount

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentAmount() float32`

GetCommitmentAmount returns the CommitmentAmount field if non-nil, zero value otherwise.

### GetCommitmentAmountOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentAmountOk() (*float32, bool)`

GetCommitmentAmountOk returns a tuple with the CommitmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentAmount

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentAmount(v float32)`

SetCommitmentAmount sets CommitmentAmount field to given value.

### HasCommitmentAmount

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentAmount() bool`

HasCommitmentAmount returns a boolean if a field has been set.

### GetCommitmentDuration

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentDuration() TypesBillingPeriod`

GetCommitmentDuration returns the CommitmentDuration field if non-nil, zero value otherwise.

### GetCommitmentDurationOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentDurationOk() (*TypesBillingPeriod, bool)`

GetCommitmentDurationOk returns a tuple with the CommitmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentDuration

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentDuration(v TypesBillingPeriod)`

SetCommitmentDuration sets CommitmentDuration field to given value.

### HasCommitmentDuration

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentDuration() bool`

HasCommitmentDuration returns a boolean if a field has been set.

### GetCommitmentOverageFactor

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentOverageFactor() float32`

GetCommitmentOverageFactor returns the CommitmentOverageFactor field if non-nil, zero value otherwise.

### GetCommitmentOverageFactorOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentOverageFactorOk() (*float32, bool)`

GetCommitmentOverageFactorOk returns a tuple with the CommitmentOverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentOverageFactor

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentOverageFactor(v float32)`

SetCommitmentOverageFactor sets CommitmentOverageFactor field to given value.

### HasCommitmentOverageFactor

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentOverageFactor() bool`

HasCommitmentOverageFactor returns a boolean if a field has been set.

### GetCommitmentQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentQuantity() float32`

GetCommitmentQuantity returns the CommitmentQuantity field if non-nil, zero value otherwise.

### GetCommitmentQuantityOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentQuantityOk() (*float32, bool)`

GetCommitmentQuantityOk returns a tuple with the CommitmentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentQuantity(v float32)`

SetCommitmentQuantity sets CommitmentQuantity field to given value.

### HasCommitmentQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentQuantity() bool`

HasCommitmentQuantity returns a boolean if a field has been set.

### GetCommitmentTrueUpEnabled

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentTrueUpEnabled() bool`

GetCommitmentTrueUpEnabled returns the CommitmentTrueUpEnabled field if non-nil, zero value otherwise.

### GetCommitmentTrueUpEnabledOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentTrueUpEnabledOk() (*bool, bool)`

GetCommitmentTrueUpEnabledOk returns a tuple with the CommitmentTrueUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentTrueUpEnabled

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentTrueUpEnabled(v bool)`

SetCommitmentTrueUpEnabled sets CommitmentTrueUpEnabled field to given value.

### HasCommitmentTrueUpEnabled

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentTrueUpEnabled() bool`

HasCommitmentTrueUpEnabled returns a boolean if a field has been set.

### GetCommitmentType

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentType() TypesCommitmentType`

GetCommitmentType returns the CommitmentType field if non-nil, zero value otherwise.

### GetCommitmentTypeOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentTypeOk() (*TypesCommitmentType, bool)`

GetCommitmentTypeOk returns a tuple with the CommitmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentType

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentType(v TypesCommitmentType)`

SetCommitmentType sets CommitmentType field to given value.

### HasCommitmentType

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentType() bool`

HasCommitmentType returns a boolean if a field has been set.

### GetCommitmentWindowed

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentWindowed() bool`

GetCommitmentWindowed returns the CommitmentWindowed field if non-nil, zero value otherwise.

### GetCommitmentWindowedOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetCommitmentWindowedOk() (*bool, bool)`

GetCommitmentWindowedOk returns a tuple with the CommitmentWindowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentWindowed

`func (o *DtoCreateSubscriptionLineItemRequest) SetCommitmentWindowed(v bool)`

SetCommitmentWindowed sets CommitmentWindowed field to given value.

### HasCommitmentWindowed

`func (o *DtoCreateSubscriptionLineItemRequest) HasCommitmentWindowed() bool`

HasCommitmentWindowed returns a boolean if a field has been set.

### GetDisplayName

`func (o *DtoCreateSubscriptionLineItemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DtoCreateSubscriptionLineItemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DtoCreateSubscriptionLineItemRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndDate

`func (o *DtoCreateSubscriptionLineItemRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DtoCreateSubscriptionLineItemRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DtoCreateSubscriptionLineItemRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateSubscriptionLineItemRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateSubscriptionLineItemRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateSubscriptionLineItemRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrice

`func (o *DtoCreateSubscriptionLineItemRequest) GetPrice() DtoSubscriptionPriceCreateRequest`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetPriceOk() (*DtoSubscriptionPriceCreateRequest, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DtoCreateSubscriptionLineItemRequest) SetPrice(v DtoSubscriptionPriceCreateRequest)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DtoCreateSubscriptionLineItemRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *DtoCreateSubscriptionLineItemRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *DtoCreateSubscriptionLineItemRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *DtoCreateSubscriptionLineItemRequest) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoCreateSubscriptionLineItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *DtoCreateSubscriptionLineItemRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DtoCreateSubscriptionLineItemRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DtoCreateSubscriptionLineItemRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionPhaseId

`func (o *DtoCreateSubscriptionLineItemRequest) GetSubscriptionPhaseId() string`

GetSubscriptionPhaseId returns the SubscriptionPhaseId field if non-nil, zero value otherwise.

### GetSubscriptionPhaseIdOk

`func (o *DtoCreateSubscriptionLineItemRequest) GetSubscriptionPhaseIdOk() (*string, bool)`

GetSubscriptionPhaseIdOk returns a tuple with the SubscriptionPhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPhaseId

`func (o *DtoCreateSubscriptionLineItemRequest) SetSubscriptionPhaseId(v string)`

SetSubscriptionPhaseId sets SubscriptionPhaseId field to given value.

### HasSubscriptionPhaseId

`func (o *DtoCreateSubscriptionLineItemRequest) HasSubscriptionPhaseId() bool`

HasSubscriptionPhaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


