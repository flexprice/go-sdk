# DtoUpdateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressCity** | Pointer to **string** | address_city is the updated city name with maximum 100 characters | [optional] 
**AddressCountry** | Pointer to **string** | address_country is the updated two-letter ISO 3166-1 alpha-2 country code | [optional] 
**AddressLine1** | Pointer to **string** | address_line1 is the updated primary address line with maximum 255 characters | [optional] 
**AddressLine2** | Pointer to **string** | address_line2 is the updated secondary address line with maximum 255 characters | [optional] 
**AddressPostalCode** | Pointer to **string** | address_postal_code is the updated postal code with maximum 20 characters | [optional] 
**AddressState** | Pointer to **string** | address_state is the updated state, province, or region name with maximum 100 characters | [optional] 
**Email** | Pointer to **string** | email is the updated email address and must be a valid email format if provided | [optional] 
**ExternalId** | Pointer to **string** | external_id is the updated external identifier for the customer | [optional] 
**IntegrationEntityMapping** | Pointer to [**[]DtoIntegrationEntityMapping**](DtoIntegrationEntityMapping.md) | integration_entity_mapping contains provider integration mappings for this customer | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata contains updated key-value pairs that will replace existing metadata | [optional] 
**Name** | Pointer to **string** | name is the updated name or company name for the customer | [optional] 
**ParentCustomerExternalId** | Pointer to **string** | parent_customer_external_id is the external ID of the parent customer from your system Exactly one of parent_customer_id or parent_customer_external_id may be provided If you provide the external ID, the parent customer value will be ignored | [optional] 
**ParentCustomerId** | Pointer to **string** | parent_customer_id is the internal FlexPrice ID of the parent customer | [optional] 

## Methods

### NewDtoUpdateCustomerRequest

`func NewDtoUpdateCustomerRequest() *DtoUpdateCustomerRequest`

NewDtoUpdateCustomerRequest instantiates a new DtoUpdateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateCustomerRequestWithDefaults

`func NewDtoUpdateCustomerRequestWithDefaults() *DtoUpdateCustomerRequest`

NewDtoUpdateCustomerRequestWithDefaults instantiates a new DtoUpdateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressCity

`func (o *DtoUpdateCustomerRequest) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *DtoUpdateCustomerRequest) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *DtoUpdateCustomerRequest) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *DtoUpdateCustomerRequest) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *DtoUpdateCustomerRequest) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *DtoUpdateCustomerRequest) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *DtoUpdateCustomerRequest) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *DtoUpdateCustomerRequest) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressLine1

`func (o *DtoUpdateCustomerRequest) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DtoUpdateCustomerRequest) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DtoUpdateCustomerRequest) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DtoUpdateCustomerRequest) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *DtoUpdateCustomerRequest) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DtoUpdateCustomerRequest) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DtoUpdateCustomerRequest) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DtoUpdateCustomerRequest) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressPostalCode

`func (o *DtoUpdateCustomerRequest) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *DtoUpdateCustomerRequest) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *DtoUpdateCustomerRequest) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *DtoUpdateCustomerRequest) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### GetAddressState

`func (o *DtoUpdateCustomerRequest) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *DtoUpdateCustomerRequest) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *DtoUpdateCustomerRequest) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *DtoUpdateCustomerRequest) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetEmail

`func (o *DtoUpdateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoUpdateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoUpdateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DtoUpdateCustomerRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalId

`func (o *DtoUpdateCustomerRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DtoUpdateCustomerRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DtoUpdateCustomerRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DtoUpdateCustomerRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIntegrationEntityMapping

`func (o *DtoUpdateCustomerRequest) GetIntegrationEntityMapping() []DtoIntegrationEntityMapping`

GetIntegrationEntityMapping returns the IntegrationEntityMapping field if non-nil, zero value otherwise.

### GetIntegrationEntityMappingOk

`func (o *DtoUpdateCustomerRequest) GetIntegrationEntityMappingOk() (*[]DtoIntegrationEntityMapping, bool)`

GetIntegrationEntityMappingOk returns a tuple with the IntegrationEntityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationEntityMapping

`func (o *DtoUpdateCustomerRequest) SetIntegrationEntityMapping(v []DtoIntegrationEntityMapping)`

SetIntegrationEntityMapping sets IntegrationEntityMapping field to given value.

### HasIntegrationEntityMapping

`func (o *DtoUpdateCustomerRequest) HasIntegrationEntityMapping() bool`

HasIntegrationEntityMapping returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateCustomerRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateCustomerRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateCustomerRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateCustomerRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateCustomerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateCustomerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateCustomerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateCustomerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentCustomerExternalId

`func (o *DtoUpdateCustomerRequest) GetParentCustomerExternalId() string`

GetParentCustomerExternalId returns the ParentCustomerExternalId field if non-nil, zero value otherwise.

### GetParentCustomerExternalIdOk

`func (o *DtoUpdateCustomerRequest) GetParentCustomerExternalIdOk() (*string, bool)`

GetParentCustomerExternalIdOk returns a tuple with the ParentCustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCustomerExternalId

`func (o *DtoUpdateCustomerRequest) SetParentCustomerExternalId(v string)`

SetParentCustomerExternalId sets ParentCustomerExternalId field to given value.

### HasParentCustomerExternalId

`func (o *DtoUpdateCustomerRequest) HasParentCustomerExternalId() bool`

HasParentCustomerExternalId returns a boolean if a field has been set.

### GetParentCustomerId

`func (o *DtoUpdateCustomerRequest) GetParentCustomerId() string`

GetParentCustomerId returns the ParentCustomerId field if non-nil, zero value otherwise.

### GetParentCustomerIdOk

`func (o *DtoUpdateCustomerRequest) GetParentCustomerIdOk() (*string, bool)`

GetParentCustomerIdOk returns a tuple with the ParentCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCustomerId

`func (o *DtoUpdateCustomerRequest) SetParentCustomerId(v string)`

SetParentCustomerId sets ParentCustomerId field to given value.

### HasParentCustomerId

`func (o *DtoUpdateCustomerRequest) HasParentCustomerId() bool`

HasParentCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


