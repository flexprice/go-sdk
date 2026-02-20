# DtoCustomerLookupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer**](GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer.md) |  | [optional] 
**Error** | Pointer to [**ErrorsErrorResponse**](ErrorsErrorResponse.md) |  | [optional] 
**Status** | Pointer to [**TypesDebugTrackerStatus**](TypesDebugTrackerStatus.md) |  | [optional] 

## Methods

### NewDtoCustomerLookupResult

`func NewDtoCustomerLookupResult() *DtoCustomerLookupResult`

NewDtoCustomerLookupResult instantiates a new DtoCustomerLookupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCustomerLookupResultWithDefaults

`func NewDtoCustomerLookupResultWithDefaults() *DtoCustomerLookupResult`

NewDtoCustomerLookupResultWithDefaults instantiates a new DtoCustomerLookupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *DtoCustomerLookupResult) GetCustomer() GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoCustomerLookupResult) GetCustomerOk() (*GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoCustomerLookupResult) SetCustomer(v GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoCustomerLookupResult) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetError

`func (o *DtoCustomerLookupResult) GetError() ErrorsErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DtoCustomerLookupResult) GetErrorOk() (*ErrorsErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DtoCustomerLookupResult) SetError(v ErrorsErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *DtoCustomerLookupResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCustomerLookupResult) GetStatus() TypesDebugTrackerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCustomerLookupResult) GetStatusOk() (*TypesDebugTrackerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCustomerLookupResult) SetStatus(v TypesDebugTrackerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCustomerLookupResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


