# DtoPriceLookupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorsErrorResponse**](ErrorsErrorResponse.md) |  | [optional] 
**MatchedPrices** | Pointer to [**[]DtoMatchedPrice**](DtoMatchedPrice.md) |  | [optional] 
**Status** | Pointer to [**TypesDebugTrackerStatus**](TypesDebugTrackerStatus.md) |  | [optional] 

## Methods

### NewDtoPriceLookupResult

`func NewDtoPriceLookupResult() *DtoPriceLookupResult`

NewDtoPriceLookupResult instantiates a new DtoPriceLookupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPriceLookupResultWithDefaults

`func NewDtoPriceLookupResultWithDefaults() *DtoPriceLookupResult`

NewDtoPriceLookupResultWithDefaults instantiates a new DtoPriceLookupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DtoPriceLookupResult) GetError() ErrorsErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DtoPriceLookupResult) GetErrorOk() (*ErrorsErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DtoPriceLookupResult) SetError(v ErrorsErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *DtoPriceLookupResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMatchedPrices

`func (o *DtoPriceLookupResult) GetMatchedPrices() []DtoMatchedPrice`

GetMatchedPrices returns the MatchedPrices field if non-nil, zero value otherwise.

### GetMatchedPricesOk

`func (o *DtoPriceLookupResult) GetMatchedPricesOk() (*[]DtoMatchedPrice, bool)`

GetMatchedPricesOk returns a tuple with the MatchedPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPrices

`func (o *DtoPriceLookupResult) SetMatchedPrices(v []DtoMatchedPrice)`

SetMatchedPrices sets MatchedPrices field to given value.

### HasMatchedPrices

`func (o *DtoPriceLookupResult) HasMatchedPrices() bool`

HasMatchedPrices returns a boolean if a field has been set.

### GetStatus

`func (o *DtoPriceLookupResult) GetStatus() TypesDebugTrackerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoPriceLookupResult) GetStatusOk() (*TypesDebugTrackerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoPriceLookupResult) SetStatus(v TypesDebugTrackerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoPriceLookupResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


