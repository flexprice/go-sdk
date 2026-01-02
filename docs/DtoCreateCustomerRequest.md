# DtoCreateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressCity** | Pointer to **string** | address_city is the city name with maximum 100 characters | [optional] 
**AddressCountry** | Pointer to **string** | address_country is the two-letter ISO 3166-1 alpha-2 country code | [optional] 
**AddressLine1** | Pointer to **string** | address_line1 is the primary address line with maximum 255 characters | [optional] 
**AddressLine2** | Pointer to **string** | address_line2 is the secondary address line with maximum 255 characters | [optional] 
**AddressPostalCode** | Pointer to **string** | address_postal_code is the ZIP code or postal code with maximum 20 characters | [optional] 
**AddressState** | Pointer to **string** | address_state is the state, province, or region name with maximum 100 characters | [optional] 
**Email** | Pointer to **string** | email is the customer&#39;s email address and must be a valid email format if provided | [optional] 
**ExternalId** | **string** | external_id is the unique identifier from your system to reference this customer (required) | 
**IntegrationEntityMapping** | Pointer to [**[]DtoIntegrationEntityMapping**](DtoIntegrationEntityMapping.md) | integration_entity_mapping contains provider integration mappings for this customer | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata contains additional key-value pairs for storing extra information | [optional] 
**Name** | Pointer to **string** | name is the full name or company name of the customer | [optional] 
**ParentCustomerExternalId** | Pointer to **string** | parent_customer_external_id is the external ID of the parent customer from your system Exactly one of parent_customer_id or parent_customer_external_id may be provided | [optional] 
**ParentCustomerId** | Pointer to **string** | parent_customer_id is the internal FlexPrice ID of the parent customer | [optional] 
**SkipOnboardingWorkflow** | Pointer to **bool** | skip_onboarding_workflow when true, prevents the customer onboarding workflow from being triggered This is used internally when a customer is created via a workflow to prevent infinite loops Default: false | [optional] 
**TaxRateOverrides** | Pointer to [**[]DtoTaxRateOverride**](DtoTaxRateOverride.md) | tax_rate_overrides contains tax rate configurations to be linked to this customer | [optional] 

## Methods

### NewDtoCreateCustomerRequest

`func NewDtoCreateCustomerRequest(externalId string, ) *DtoCreateCustomerRequest`

NewDtoCreateCustomerRequest instantiates a new DtoCreateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateCustomerRequestWithDefaults

`func NewDtoCreateCustomerRequestWithDefaults() *DtoCreateCustomerRequest`

NewDtoCreateCustomerRequestWithDefaults instantiates a new DtoCreateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressCity

`func (o *DtoCreateCustomerRequest) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *DtoCreateCustomerRequest) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *DtoCreateCustomerRequest) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *DtoCreateCustomerRequest) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *DtoCreateCustomerRequest) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *DtoCreateCustomerRequest) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *DtoCreateCustomerRequest) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *DtoCreateCustomerRequest) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressLine1

`func (o *DtoCreateCustomerRequest) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DtoCreateCustomerRequest) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DtoCreateCustomerRequest) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *DtoCreateCustomerRequest) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *DtoCreateCustomerRequest) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DtoCreateCustomerRequest) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DtoCreateCustomerRequest) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DtoCreateCustomerRequest) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressPostalCode

`func (o *DtoCreateCustomerRequest) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *DtoCreateCustomerRequest) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *DtoCreateCustomerRequest) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *DtoCreateCustomerRequest) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### GetAddressState

`func (o *DtoCreateCustomerRequest) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *DtoCreateCustomerRequest) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *DtoCreateCustomerRequest) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *DtoCreateCustomerRequest) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetEmail

`func (o *DtoCreateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoCreateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoCreateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DtoCreateCustomerRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalId

`func (o *DtoCreateCustomerRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DtoCreateCustomerRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DtoCreateCustomerRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetIntegrationEntityMapping

`func (o *DtoCreateCustomerRequest) GetIntegrationEntityMapping() []DtoIntegrationEntityMapping`

GetIntegrationEntityMapping returns the IntegrationEntityMapping field if non-nil, zero value otherwise.

### GetIntegrationEntityMappingOk

`func (o *DtoCreateCustomerRequest) GetIntegrationEntityMappingOk() (*[]DtoIntegrationEntityMapping, bool)`

GetIntegrationEntityMappingOk returns a tuple with the IntegrationEntityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationEntityMapping

`func (o *DtoCreateCustomerRequest) SetIntegrationEntityMapping(v []DtoIntegrationEntityMapping)`

SetIntegrationEntityMapping sets IntegrationEntityMapping field to given value.

### HasIntegrationEntityMapping

`func (o *DtoCreateCustomerRequest) HasIntegrationEntityMapping() bool`

HasIntegrationEntityMapping returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateCustomerRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateCustomerRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateCustomerRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateCustomerRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateCustomerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateCustomerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateCustomerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoCreateCustomerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentCustomerExternalId

`func (o *DtoCreateCustomerRequest) GetParentCustomerExternalId() string`

GetParentCustomerExternalId returns the ParentCustomerExternalId field if non-nil, zero value otherwise.

### GetParentCustomerExternalIdOk

`func (o *DtoCreateCustomerRequest) GetParentCustomerExternalIdOk() (*string, bool)`

GetParentCustomerExternalIdOk returns a tuple with the ParentCustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCustomerExternalId

`func (o *DtoCreateCustomerRequest) SetParentCustomerExternalId(v string)`

SetParentCustomerExternalId sets ParentCustomerExternalId field to given value.

### HasParentCustomerExternalId

`func (o *DtoCreateCustomerRequest) HasParentCustomerExternalId() bool`

HasParentCustomerExternalId returns a boolean if a field has been set.

### GetParentCustomerId

`func (o *DtoCreateCustomerRequest) GetParentCustomerId() string`

GetParentCustomerId returns the ParentCustomerId field if non-nil, zero value otherwise.

### GetParentCustomerIdOk

`func (o *DtoCreateCustomerRequest) GetParentCustomerIdOk() (*string, bool)`

GetParentCustomerIdOk returns a tuple with the ParentCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCustomerId

`func (o *DtoCreateCustomerRequest) SetParentCustomerId(v string)`

SetParentCustomerId sets ParentCustomerId field to given value.

### HasParentCustomerId

`func (o *DtoCreateCustomerRequest) HasParentCustomerId() bool`

HasParentCustomerId returns a boolean if a field has been set.

### GetSkipOnboardingWorkflow

`func (o *DtoCreateCustomerRequest) GetSkipOnboardingWorkflow() bool`

GetSkipOnboardingWorkflow returns the SkipOnboardingWorkflow field if non-nil, zero value otherwise.

### GetSkipOnboardingWorkflowOk

`func (o *DtoCreateCustomerRequest) GetSkipOnboardingWorkflowOk() (*bool, bool)`

GetSkipOnboardingWorkflowOk returns a tuple with the SkipOnboardingWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOnboardingWorkflow

`func (o *DtoCreateCustomerRequest) SetSkipOnboardingWorkflow(v bool)`

SetSkipOnboardingWorkflow sets SkipOnboardingWorkflow field to given value.

### HasSkipOnboardingWorkflow

`func (o *DtoCreateCustomerRequest) HasSkipOnboardingWorkflow() bool`

HasSkipOnboardingWorkflow returns a boolean if a field has been set.

### GetTaxRateOverrides

`func (o *DtoCreateCustomerRequest) GetTaxRateOverrides() []DtoTaxRateOverride`

GetTaxRateOverrides returns the TaxRateOverrides field if non-nil, zero value otherwise.

### GetTaxRateOverridesOk

`func (o *DtoCreateCustomerRequest) GetTaxRateOverridesOk() (*[]DtoTaxRateOverride, bool)`

GetTaxRateOverridesOk returns a tuple with the TaxRateOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateOverrides

`func (o *DtoCreateCustomerRequest) SetTaxRateOverrides(v []DtoTaxRateOverride)`

SetTaxRateOverrides sets TaxRateOverrides field to given value.

### HasTaxRateOverrides

`func (o *DtoCreateCustomerRequest) HasTaxRateOverrides() bool`

HasTaxRateOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


