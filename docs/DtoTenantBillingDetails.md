# DtoTenantBillingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**DtoAddress**](DtoAddress.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**HelpEmail** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTenantBillingDetails

`func NewDtoTenantBillingDetails() *DtoTenantBillingDetails`

NewDtoTenantBillingDetails instantiates a new DtoTenantBillingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTenantBillingDetailsWithDefaults

`func NewDtoTenantBillingDetailsWithDefaults() *DtoTenantBillingDetails`

NewDtoTenantBillingDetailsWithDefaults instantiates a new DtoTenantBillingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DtoTenantBillingDetails) GetAddress() DtoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DtoTenantBillingDetails) GetAddressOk() (*DtoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DtoTenantBillingDetails) SetAddress(v DtoAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DtoTenantBillingDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *DtoTenantBillingDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoTenantBillingDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoTenantBillingDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DtoTenantBillingDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHelpEmail

`func (o *DtoTenantBillingDetails) GetHelpEmail() string`

GetHelpEmail returns the HelpEmail field if non-nil, zero value otherwise.

### GetHelpEmailOk

`func (o *DtoTenantBillingDetails) GetHelpEmailOk() (*string, bool)`

GetHelpEmailOk returns a tuple with the HelpEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpEmail

`func (o *DtoTenantBillingDetails) SetHelpEmail(v string)`

SetHelpEmail sets HelpEmail field to given value.

### HasHelpEmail

`func (o *DtoTenantBillingDetails) HasHelpEmail() bool`

HasHelpEmail returns a boolean if a field has been set.

### GetPhone

`func (o *DtoTenantBillingDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DtoTenantBillingDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DtoTenantBillingDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DtoTenantBillingDetails) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


