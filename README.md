# FlexPrice Go SDK

This is the Go client library for the FlexPrice API.

## Installation

```bash
go get github.com/flexprice/go-sdk
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    flexpriceclient "github.com/flexprice/go-sdk"
)

func main() {
    cfg := flexpriceclient.NewConfiguration()
    cfg.Host = "https://api.flexprice.io"
    
    // Add your API key to the default header
    cfg.DefaultHeader["X-API-Key"] = "YOUR_API_KEY"
    
    client := flexpriceclient.NewAPIClient(cfg)
    
    // Example: Get customer by ID
    customer, resp, err := client.CustomersAPI.CustomersIdGet(context.Background(), "customer-123").Execute()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Customer: %s\n", customer.GetName())
}
```

## Features

- Complete API coverage
- Type-safe client
- Detailed documentation
- Error handling

## Documentation

For detailed API documentation, refer to the code comments and the official FlexPrice API documentation. 

## Auto-Generated API Documentation


## Documentation for API Endpoints

All URIs are relative to */v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthAPI* | [**AuthLoginPost**](docs/AuthAPI.md#authloginpost) | **Post** /auth/login | Login
*AuthAPI* | [**AuthSignupPost**](docs/AuthAPI.md#authsignuppost) | **Post** /auth/signup | Sign up
*CustomersAPI* | [**CustomersGet**](docs/CustomersAPI.md#customersget) | **Get** /customers | Get customers
*CustomersAPI* | [**CustomersIdDelete**](docs/CustomersAPI.md#customersiddelete) | **Delete** /customers/{id} | Delete a customer
*CustomersAPI* | [**CustomersIdEntitlementsGet**](docs/CustomersAPI.md#customersidentitlementsget) | **Get** /customers/{id}/entitlements | Get customer entitlements
*CustomersAPI* | [**CustomersIdGet**](docs/CustomersAPI.md#customersidget) | **Get** /customers/{id} | Get a customer
*CustomersAPI* | [**CustomersIdPut**](docs/CustomersAPI.md#customersidput) | **Put** /customers/{id} | Update a customer
*CustomersAPI* | [**CustomersIdUsageGet**](docs/CustomersAPI.md#customersidusageget) | **Get** /customers/{id}/usage | Get customer usage summary
*CustomersAPI* | [**CustomersLookupLookupKeyGet**](docs/CustomersAPI.md#customerslookuplookupkeyget) | **Get** /customers/lookup/{lookup_key} | Get a customer by lookup key
*CustomersAPI* | [**CustomersPost**](docs/CustomersAPI.md#customerspost) | **Post** /customers | Create a customer
*EntitlementsAPI* | [**EntitlementsGet**](docs/EntitlementsAPI.md#entitlementsget) | **Get** /entitlements | Get entitlements
*EntitlementsAPI* | [**EntitlementsIdDelete**](docs/EntitlementsAPI.md#entitlementsiddelete) | **Delete** /entitlements/{id} | Delete an entitlement
*EntitlementsAPI* | [**EntitlementsIdGet**](docs/EntitlementsAPI.md#entitlementsidget) | **Get** /entitlements/{id} | Get an entitlement by ID
*EntitlementsAPI* | [**EntitlementsIdPut**](docs/EntitlementsAPI.md#entitlementsidput) | **Put** /entitlements/{id} | Update an entitlement
*EntitlementsAPI* | [**EntitlementsPost**](docs/EntitlementsAPI.md#entitlementspost) | **Post** /entitlements | Create a new entitlement
*EntitlementsAPI* | [**PlansIdEntitlementsGet**](docs/EntitlementsAPI.md#plansidentitlementsget) | **Get** /plans/{id}/entitlements | Get plan entitlements
*EnvironmentsAPI* | [**EnvironmentsGet**](docs/EnvironmentsAPI.md#environmentsget) | **Get** /environments | Get environments
*EnvironmentsAPI* | [**EnvironmentsIdGet**](docs/EnvironmentsAPI.md#environmentsidget) | **Get** /environments/{id} | Get an environment
*EnvironmentsAPI* | [**EnvironmentsIdPut**](docs/EnvironmentsAPI.md#environmentsidput) | **Put** /environments/{id} | Update an environment
*EnvironmentsAPI* | [**EnvironmentsPost**](docs/EnvironmentsAPI.md#environmentspost) | **Post** /environments | Create an environment
*EventsAPI* | [**EventsBulkPost**](docs/EventsAPI.md#eventsbulkpost) | **Post** /events/bulk | Bulk Ingest events
*EventsAPI* | [**EventsGet**](docs/EventsAPI.md#eventsget) | **Get** /events | Get raw events
*EventsAPI* | [**EventsPost**](docs/EventsAPI.md#eventspost) | **Post** /events | Ingest event
*EventsAPI* | [**EventsUsageMeterPost**](docs/EventsAPI.md#eventsusagemeterpost) | **Post** /events/usage/meter | Get usage by meter
*EventsAPI* | [**EventsUsagePost**](docs/EventsAPI.md#eventsusagepost) | **Post** /events/usage | Get usage statistics
*FeaturesAPI* | [**FeaturesGet**](docs/FeaturesAPI.md#featuresget) | **Get** /features | List features
*FeaturesAPI* | [**FeaturesIdDelete**](docs/FeaturesAPI.md#featuresiddelete) | **Delete** /features/{id} | Delete a feature
*FeaturesAPI* | [**FeaturesIdGet**](docs/FeaturesAPI.md#featuresidget) | **Get** /features/{id} | Get a feature by ID
*FeaturesAPI* | [**FeaturesIdPut**](docs/FeaturesAPI.md#featuresidput) | **Put** /features/{id} | Update a feature
*FeaturesAPI* | [**FeaturesPost**](docs/FeaturesAPI.md#featurespost) | **Post** /features | Create a new feature
*IntegrationsAPI* | [**SecretsIntegrationsIdDelete**](docs/IntegrationsAPI.md#secretsintegrationsiddelete) | **Delete** /secrets/integrations/{id} | Delete an integration
*IntegrationsAPI* | [**SecretsIntegrationsLinkedGet**](docs/IntegrationsAPI.md#secretsintegrationslinkedget) | **Get** /secrets/integrations/linked | List linked integrations
*IntegrationsAPI* | [**SecretsIntegrationsProviderGet**](docs/IntegrationsAPI.md#secretsintegrationsproviderget) | **Get** /secrets/integrations/{provider} | Get integration details
*IntegrationsAPI* | [**SecretsIntegrationsProviderPost**](docs/IntegrationsAPI.md#secretsintegrationsproviderpost) | **Post** /secrets/integrations/{provider} | Create or update an integration
*InvoicesAPI* | [**CustomersIdInvoicesSummaryGet**](docs/InvoicesAPI.md#customersidinvoicessummaryget) | **Get** /customers/{id}/invoices/summary | Get a customer invoice summary
*InvoicesAPI* | [**InvoicesGet**](docs/InvoicesAPI.md#invoicesget) | **Get** /invoices | List invoices
*InvoicesAPI* | [**InvoicesIdFinalizePost**](docs/InvoicesAPI.md#invoicesidfinalizepost) | **Post** /invoices/{id}/finalize | Finalize an invoice
*InvoicesAPI* | [**InvoicesIdGet**](docs/InvoicesAPI.md#invoicesidget) | **Get** /invoices/{id} | Get an invoice by ID
*InvoicesAPI* | [**InvoicesIdPaymentAttemptPost**](docs/InvoicesAPI.md#invoicesidpaymentattemptpost) | **Post** /invoices/{id}/payment/attempt | Attempt payment for an invoice
*InvoicesAPI* | [**InvoicesIdPaymentPut**](docs/InvoicesAPI.md#invoicesidpaymentput) | **Put** /invoices/{id}/payment | Update invoice payment status
*InvoicesAPI* | [**InvoicesIdPdfGet**](docs/InvoicesAPI.md#invoicesidpdfget) | **Get** /invoices/{id}/pdf | Get PDF for an invoice
*InvoicesAPI* | [**InvoicesIdVoidPost**](docs/InvoicesAPI.md#invoicesidvoidpost) | **Post** /invoices/{id}/void | Void an invoice
*InvoicesAPI* | [**InvoicesPost**](docs/InvoicesAPI.md#invoicespost) | **Post** /invoices | Create a new invoice
*InvoicesAPI* | [**InvoicesPreviewPost**](docs/InvoicesAPI.md#invoicespreviewpost) | **Post** /invoices/preview | Get a preview invoice
*MetersAPI* | [**MetersGet**](docs/MetersAPI.md#metersget) | **Get** /meters | List meters
*MetersAPI* | [**MetersIdDelete**](docs/MetersAPI.md#metersiddelete) | **Delete** /meters/{id} | Delete meter
*MetersAPI* | [**MetersIdDisablePost**](docs/MetersAPI.md#metersiddisablepost) | **Post** /meters/{id}/disable | Disable meter [TODO: Deprecate]
*MetersAPI* | [**MetersIdGet**](docs/MetersAPI.md#metersidget) | **Get** /meters/{id} | Get meter
*MetersAPI* | [**MetersIdPut**](docs/MetersAPI.md#metersidput) | **Put** /meters/{id} | Update meter
*MetersAPI* | [**MetersPost**](docs/MetersAPI.md#meterspost) | **Post** /meters | Create meter
*PaymentsAPI* | [**PaymentsGet**](docs/PaymentsAPI.md#paymentsget) | **Get** /payments | List payments
*PaymentsAPI* | [**PaymentsIdDelete**](docs/PaymentsAPI.md#paymentsiddelete) | **Delete** /payments/{id} | Delete a payment
*PaymentsAPI* | [**PaymentsIdGet**](docs/PaymentsAPI.md#paymentsidget) | **Get** /payments/{id} | Get a payment by ID
*PaymentsAPI* | [**PaymentsIdProcessPost**](docs/PaymentsAPI.md#paymentsidprocesspost) | **Post** /payments/{id}/process | Process a payment
*PaymentsAPI* | [**PaymentsIdPut**](docs/PaymentsAPI.md#paymentsidput) | **Put** /payments/{id} | Update a payment
*PaymentsAPI* | [**PaymentsPost**](docs/PaymentsAPI.md#paymentspost) | **Post** /payments | Create a new payment
*PlansAPI* | [**PlansGet**](docs/PlansAPI.md#plansget) | **Get** /plans | Get plans
*PlansAPI* | [**PlansIdDelete**](docs/PlansAPI.md#plansiddelete) | **Delete** /plans/{id} | Delete a plan
*PlansAPI* | [**PlansIdGet**](docs/PlansAPI.md#plansidget) | **Get** /plans/{id} | Get a plan
*PlansAPI* | [**PlansIdPut**](docs/PlansAPI.md#plansidput) | **Put** /plans/{id} | Update a plan
*PlansAPI* | [**PlansPost**](docs/PlansAPI.md#planspost) | **Post** /plans | Create a new plan
*PricesAPI* | [**PricesGet**](docs/PricesAPI.md#pricesget) | **Get** /prices | Get prices
*PricesAPI* | [**PricesIdDelete**](docs/PricesAPI.md#pricesiddelete) | **Delete** /prices/{id} | Delete a price
*PricesAPI* | [**PricesIdGet**](docs/PricesAPI.md#pricesidget) | **Get** /prices/{id} | Get a price by ID
*PricesAPI* | [**PricesIdPut**](docs/PricesAPI.md#pricesidput) | **Put** /prices/{id} | Update a price
*PricesAPI* | [**PricesPost**](docs/PricesAPI.md#pricespost) | **Post** /prices | Create a new price
*SecretsAPI* | [**SecretsApiKeysGet**](docs/SecretsAPI.md#secretsapikeysget) | **Get** /secrets/api/keys | List API keys
*SecretsAPI* | [**SecretsApiKeysIdDelete**](docs/SecretsAPI.md#secretsapikeysiddelete) | **Delete** /secrets/api/keys/{id} | Delete an API key
*SecretsAPI* | [**SecretsApiKeysPost**](docs/SecretsAPI.md#secretsapikeyspost) | **Post** /secrets/api/keys | Create a new API key
*SubscriptionsAPI* | [**SubscriptionsGet**](docs/SubscriptionsAPI.md#subscriptionsget) | **Get** /subscriptions | List subscriptions
*SubscriptionsAPI* | [**SubscriptionsIdCancelPost**](docs/SubscriptionsAPI.md#subscriptionsidcancelpost) | **Post** /subscriptions/{id}/cancel | Cancel subscription
*SubscriptionsAPI* | [**SubscriptionsIdGet**](docs/SubscriptionsAPI.md#subscriptionsidget) | **Get** /subscriptions/{id} | Get subscription
*SubscriptionsAPI* | [**SubscriptionsIdPausePost**](docs/SubscriptionsAPI.md#subscriptionsidpausepost) | **Post** /subscriptions/{id}/pause | Pause a subscription
*SubscriptionsAPI* | [**SubscriptionsIdPausesGet**](docs/SubscriptionsAPI.md#subscriptionsidpausesget) | **Get** /subscriptions/{id}/pauses | List all pauses for a subscription
*SubscriptionsAPI* | [**SubscriptionsIdResumePost**](docs/SubscriptionsAPI.md#subscriptionsidresumepost) | **Post** /subscriptions/{id}/resume | Resume a paused subscription
*SubscriptionsAPI* | [**SubscriptionsPost**](docs/SubscriptionsAPI.md#subscriptionspost) | **Post** /subscriptions | Create subscription
*SubscriptionsAPI* | [**SubscriptionsUsagePost**](docs/SubscriptionsAPI.md#subscriptionsusagepost) | **Post** /subscriptions/usage | Get usage by subscription
*TasksAPI* | [**TasksGet**](docs/TasksAPI.md#tasksget) | **Get** /tasks | List tasks
*TasksAPI* | [**TasksIdGet**](docs/TasksAPI.md#tasksidget) | **Get** /tasks/{id} | Get a task by ID
*TasksAPI* | [**TasksIdProcessPost**](docs/TasksAPI.md#tasksidprocesspost) | **Post** /tasks/{id}/process | Process a task
*TasksAPI* | [**TasksIdStatusPut**](docs/TasksAPI.md#tasksidstatusput) | **Put** /tasks/{id}/status | Update task status
*TasksAPI* | [**TasksPost**](docs/TasksAPI.md#taskspost) | **Post** /tasks | Create a new task
*TenantsAPI* | [**TenantBillingGet**](docs/TenantsAPI.md#tenantbillingget) | **Get** /tenant/billing | Get billing usage for the current tenant
*TenantsAPI* | [**TenantsIdGet**](docs/TenantsAPI.md#tenantsidget) | **Get** /tenants/{id} | Get tenant by ID
*TenantsAPI* | [**TenantsPost**](docs/TenantsAPI.md#tenantspost) | **Post** /tenants | Create a new tenant
*TenantsAPI* | [**TenantsUpdatePut**](docs/TenantsAPI.md#tenantsupdateput) | **Put** /tenants/update | Update a tenant
*UsersAPI* | [**UsersMeGet**](docs/UsersAPI.md#usersmeget) | **Get** /users/me | Get user info
*WalletsAPI* | [**CustomersIdWalletsGet**](docs/WalletsAPI.md#customersidwalletsget) | **Get** /customers/{id}/wallets | Get wallets by customer ID
*WalletsAPI* | [**WalletsIdBalanceRealTimeGet**](docs/WalletsAPI.md#walletsidbalancerealtimeget) | **Get** /wallets/{id}/balance/real-time | Get wallet balance
*WalletsAPI* | [**WalletsIdGet**](docs/WalletsAPI.md#walletsidget) | **Get** /wallets/{id} | Get wallet by ID
*WalletsAPI* | [**WalletsIdPut**](docs/WalletsAPI.md#walletsidput) | **Put** /wallets/{id} | Update a wallet
*WalletsAPI* | [**WalletsIdTerminatePost**](docs/WalletsAPI.md#walletsidterminatepost) | **Post** /wallets/{id}/terminate | Terminate a wallet
*WalletsAPI* | [**WalletsIdTopUpPost**](docs/WalletsAPI.md#walletsidtopuppost) | **Post** /wallets/{id}/top-up | Top up wallet
*WalletsAPI* | [**WalletsIdTransactionsGet**](docs/WalletsAPI.md#walletsidtransactionsget) | **Get** /wallets/{id}/transactions | Get wallet transactions
*WalletsAPI* | [**WalletsPost**](docs/WalletsAPI.md#walletspost) | **Post** /wallets | Create a new wallet


## Documentation For Models

 - [DtoAddress](docs/DtoAddress.md)
 - [DtoAggregatedEntitlement](docs/DtoAggregatedEntitlement.md)
 - [DtoAggregatedFeature](docs/DtoAggregatedFeature.md)
 - [DtoAuthResponse](docs/DtoAuthResponse.md)
 - [DtoBillingPeriodInfo](docs/DtoBillingPeriodInfo.md)
 - [DtoBulkIngestEventRequest](docs/DtoBulkIngestEventRequest.md)
 - [DtoCreateAPIKeyRequest](docs/DtoCreateAPIKeyRequest.md)
 - [DtoCreateAPIKeyResponse](docs/DtoCreateAPIKeyResponse.md)
 - [DtoCreateCustomerRequest](docs/DtoCreateCustomerRequest.md)
 - [DtoCreateEntitlementRequest](docs/DtoCreateEntitlementRequest.md)
 - [DtoCreateEnvironmentRequest](docs/DtoCreateEnvironmentRequest.md)
 - [DtoCreateFeatureRequest](docs/DtoCreateFeatureRequest.md)
 - [DtoCreateIntegrationRequest](docs/DtoCreateIntegrationRequest.md)
 - [DtoCreateInvoiceLineItemRequest](docs/DtoCreateInvoiceLineItemRequest.md)
 - [DtoCreateInvoiceRequest](docs/DtoCreateInvoiceRequest.md)
 - [DtoCreateMeterRequest](docs/DtoCreateMeterRequest.md)
 - [DtoCreatePaymentRequest](docs/DtoCreatePaymentRequest.md)
 - [DtoCreatePlanEntitlementRequest](docs/DtoCreatePlanEntitlementRequest.md)
 - [DtoCreatePlanPriceRequest](docs/DtoCreatePlanPriceRequest.md)
 - [DtoCreatePlanRequest](docs/DtoCreatePlanRequest.md)
 - [DtoCreatePriceRequest](docs/DtoCreatePriceRequest.md)
 - [DtoCreatePriceTier](docs/DtoCreatePriceTier.md)
 - [DtoCreateSubscriptionRequest](docs/DtoCreateSubscriptionRequest.md)
 - [DtoCreateTaskRequest](docs/DtoCreateTaskRequest.md)
 - [DtoCreateTenantRequest](docs/DtoCreateTenantRequest.md)
 - [DtoCreateWalletRequest](docs/DtoCreateWalletRequest.md)
 - [DtoCustomerEntitlementsResponse](docs/DtoCustomerEntitlementsResponse.md)
 - [DtoCustomerInvoiceSummary](docs/DtoCustomerInvoiceSummary.md)
 - [DtoCustomerMultiCurrencyInvoiceSummary](docs/DtoCustomerMultiCurrencyInvoiceSummary.md)
 - [DtoCustomerResponse](docs/DtoCustomerResponse.md)
 - [DtoCustomerUsageSummaryResponse](docs/DtoCustomerUsageSummaryResponse.md)
 - [DtoEntitlementResponse](docs/DtoEntitlementResponse.md)
 - [DtoEntitlementSource](docs/DtoEntitlementSource.md)
 - [DtoEnvironmentResponse](docs/DtoEnvironmentResponse.md)
 - [DtoEvent](docs/DtoEvent.md)
 - [DtoFeatureResponse](docs/DtoFeatureResponse.md)
 - [DtoFeatureUsageSummary](docs/DtoFeatureUsageSummary.md)
 - [DtoGetEventsResponse](docs/DtoGetEventsResponse.md)
 - [DtoGetPreviewInvoiceRequest](docs/DtoGetPreviewInvoiceRequest.md)
 - [DtoGetUsageByMeterRequest](docs/DtoGetUsageByMeterRequest.md)
 - [DtoGetUsageBySubscriptionRequest](docs/DtoGetUsageBySubscriptionRequest.md)
 - [DtoGetUsageBySubscriptionResponse](docs/DtoGetUsageBySubscriptionResponse.md)
 - [DtoGetUsageRequest](docs/DtoGetUsageRequest.md)
 - [DtoGetUsageResponse](docs/DtoGetUsageResponse.md)
 - [DtoIngestEventRequest](docs/DtoIngestEventRequest.md)
 - [DtoInvoiceLineItemResponse](docs/DtoInvoiceLineItemResponse.md)
 - [DtoInvoiceResponse](docs/DtoInvoiceResponse.md)
 - [DtoLinkedIntegrationsResponse](docs/DtoLinkedIntegrationsResponse.md)
 - [DtoListCustomersResponse](docs/DtoListCustomersResponse.md)
 - [DtoListEntitlementsResponse](docs/DtoListEntitlementsResponse.md)
 - [DtoListEnvironmentsResponse](docs/DtoListEnvironmentsResponse.md)
 - [DtoListFeaturesResponse](docs/DtoListFeaturesResponse.md)
 - [DtoListInvoicesResponse](docs/DtoListInvoicesResponse.md)
 - [DtoListPaymentsResponse](docs/DtoListPaymentsResponse.md)
 - [DtoListPlansResponse](docs/DtoListPlansResponse.md)
 - [DtoListPricesResponse](docs/DtoListPricesResponse.md)
 - [DtoListSecretsResponse](docs/DtoListSecretsResponse.md)
 - [DtoListSubscriptionPausesResponse](docs/DtoListSubscriptionPausesResponse.md)
 - [DtoListSubscriptionsResponse](docs/DtoListSubscriptionsResponse.md)
 - [DtoListTasksResponse](docs/DtoListTasksResponse.md)
 - [DtoListWalletTransactionsResponse](docs/DtoListWalletTransactionsResponse.md)
 - [DtoLoginRequest](docs/DtoLoginRequest.md)
 - [DtoMeterResponse](docs/DtoMeterResponse.md)
 - [DtoPauseSubscriptionRequest](docs/DtoPauseSubscriptionRequest.md)
 - [DtoPaymentAttemptResponse](docs/DtoPaymentAttemptResponse.md)
 - [DtoPaymentResponse](docs/DtoPaymentResponse.md)
 - [DtoPlanResponse](docs/DtoPlanResponse.md)
 - [DtoPriceResponse](docs/DtoPriceResponse.md)
 - [DtoResumeSubscriptionRequest](docs/DtoResumeSubscriptionRequest.md)
 - [DtoSecretResponse](docs/DtoSecretResponse.md)
 - [DtoSignUpRequest](docs/DtoSignUpRequest.md)
 - [DtoSubscriptionPauseResponse](docs/DtoSubscriptionPauseResponse.md)
 - [DtoSubscriptionResponse](docs/DtoSubscriptionResponse.md)
 - [DtoSubscriptionUsageByMetersResponse](docs/DtoSubscriptionUsageByMetersResponse.md)
 - [DtoTaskResponse](docs/DtoTaskResponse.md)
 - [DtoTenantBillingDetails](docs/DtoTenantBillingDetails.md)
 - [DtoTenantBillingUsage](docs/DtoTenantBillingUsage.md)
 - [DtoTenantResponse](docs/DtoTenantResponse.md)
 - [DtoTopUpWalletRequest](docs/DtoTopUpWalletRequest.md)
 - [DtoUpdateCustomerRequest](docs/DtoUpdateCustomerRequest.md)
 - [DtoUpdateEntitlementRequest](docs/DtoUpdateEntitlementRequest.md)
 - [DtoUpdateEnvironmentRequest](docs/DtoUpdateEnvironmentRequest.md)
 - [DtoUpdateFeatureRequest](docs/DtoUpdateFeatureRequest.md)
 - [DtoUpdateMeterRequest](docs/DtoUpdateMeterRequest.md)
 - [DtoUpdatePaymentRequest](docs/DtoUpdatePaymentRequest.md)
 - [DtoUpdatePaymentStatusRequest](docs/DtoUpdatePaymentStatusRequest.md)
 - [DtoUpdatePlanEntitlementRequest](docs/DtoUpdatePlanEntitlementRequest.md)
 - [DtoUpdatePlanPriceRequest](docs/DtoUpdatePlanPriceRequest.md)
 - [DtoUpdatePlanRequest](docs/DtoUpdatePlanRequest.md)
 - [DtoUpdatePriceRequest](docs/DtoUpdatePriceRequest.md)
 - [DtoUpdateTaskStatusRequest](docs/DtoUpdateTaskStatusRequest.md)
 - [DtoUpdateTenantRequest](docs/DtoUpdateTenantRequest.md)
 - [DtoUpdateWalletRequest](docs/DtoUpdateWalletRequest.md)
 - [DtoUsageResult](docs/DtoUsageResult.md)
 - [DtoUserResponse](docs/DtoUserResponse.md)
 - [DtoWalletBalanceResponse](docs/DtoWalletBalanceResponse.md)
 - [DtoWalletResponse](docs/DtoWalletResponse.md)
 - [DtoWalletTransactionResponse](docs/DtoWalletTransactionResponse.md)
 - [ErrorsErrorDetail](docs/ErrorsErrorDetail.md)
 - [ErrorsErrorResponse](docs/ErrorsErrorResponse.md)
 - [MeterAggregation](docs/MeterAggregation.md)
 - [MeterFilter](docs/MeterFilter.md)
 - [PriceJSONBTransformQuantity](docs/PriceJSONBTransformQuantity.md)
 - [PricePrice](docs/PricePrice.md)
 - [PricePriceTier](docs/PricePriceTier.md)
 - [PriceTransformQuantity](docs/PriceTransformQuantity.md)
 - [SubscriptionSubscriptionLineItem](docs/SubscriptionSubscriptionLineItem.md)
 - [SubscriptionSubscriptionPause](docs/SubscriptionSubscriptionPause.md)
 - [TypesAggregationType](docs/TypesAggregationType.md)
 - [TypesAutoTopupTrigger](docs/TypesAutoTopupTrigger.md)
 - [TypesBillingCadence](docs/TypesBillingCadence.md)
 - [TypesBillingModel](docs/TypesBillingModel.md)
 - [TypesBillingPeriod](docs/TypesBillingPeriod.md)
 - [TypesBillingTier](docs/TypesBillingTier.md)
 - [TypesEntityType](docs/TypesEntityType.md)
 - [TypesFeatureType](docs/TypesFeatureType.md)
 - [TypesFileType](docs/TypesFileType.md)
 - [TypesInvoiceBillingReason](docs/TypesInvoiceBillingReason.md)
 - [TypesInvoiceCadence](docs/TypesInvoiceCadence.md)
 - [TypesInvoiceStatus](docs/TypesInvoiceStatus.md)
 - [TypesInvoiceType](docs/TypesInvoiceType.md)
 - [TypesPaginationResponse](docs/TypesPaginationResponse.md)
 - [TypesPauseMode](docs/TypesPauseMode.md)
 - [TypesPauseStatus](docs/TypesPauseStatus.md)
 - [TypesPaymentDestinationType](docs/TypesPaymentDestinationType.md)
 - [TypesPaymentMethodType](docs/TypesPaymentMethodType.md)
 - [TypesPaymentStatus](docs/TypesPaymentStatus.md)
 - [TypesPriceType](docs/TypesPriceType.md)
 - [TypesResetUsage](docs/TypesResetUsage.md)
 - [TypesResumeMode](docs/TypesResumeMode.md)
 - [TypesSecretProvider](docs/TypesSecretProvider.md)
 - [TypesSecretType](docs/TypesSecretType.md)
 - [TypesStatus](docs/TypesStatus.md)
 - [TypesSubscriptionStatus](docs/TypesSubscriptionStatus.md)
 - [TypesTaskStatus](docs/TypesTaskStatus.md)
 - [TypesTaskType](docs/TypesTaskType.md)
 - [TypesTransactionReason](docs/TypesTransactionReason.md)
 - [TypesTransactionStatus](docs/TypesTransactionStatus.md)
 - [TypesTransactionType](docs/TypesTransactionType.md)
 - [TypesWalletConfig](docs/TypesWalletConfig.md)
 - [TypesWalletConfigPriceType](docs/TypesWalletConfigPriceType.md)
 - [TypesWalletStatus](docs/TypesWalletStatus.md)
 - [TypesWalletTxReferenceType](docs/TypesWalletTxReferenceType.md)
 - [TypesWalletType](docs/TypesWalletType.md)
 - [TypesWindowSize](docs/TypesWindowSize.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ApiKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		flexpriceclient.ContextAPIKeys,
		map[string]flexpriceclient.APIKey{
			"ApiKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



