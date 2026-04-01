# WebhookEventName

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.WebhookEventNameSubscriptionCreated

// Open enum: custom values can be created with a direct type cast
custom := types.WebhookEventName("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `WebhookEventNameSubscriptionCreated`           | subscription.created                            |
| `WebhookEventNameSubscriptionDraftCreated`      | subscription.draft.created                      |
| `WebhookEventNameSubscriptionActivated`         | subscription.activated                          |
| `WebhookEventNameSubscriptionUpdated`           | subscription.updated                            |
| `WebhookEventNameSubscriptionPaused`            | subscription.paused                             |
| `WebhookEventNameSubscriptionCancelled`         | subscription.cancelled                          |
| `WebhookEventNameSubscriptionResumed`           | subscription.resumed                            |
| `WebhookEventNameSubscriptionPhaseCreated`      | subscription.phase.created                      |
| `WebhookEventNameSubscriptionPhaseUpdated`      | subscription.phase.updated                      |
| `WebhookEventNameSubscriptionPhaseDeleted`      | subscription.phase.deleted                      |
| `WebhookEventNameFeatureCreated`                | feature.created                                 |
| `WebhookEventNameFeatureUpdated`                | feature.updated                                 |
| `WebhookEventNameFeatureDeleted`                | feature.deleted                                 |
| `WebhookEventNameFeatureWalletBalanceAlert`     | feature.wallet_balance.alert                    |
| `WebhookEventNameEntitlementCreated`            | entitlement.created                             |
| `WebhookEventNameEntitlementUpdated`            | entitlement.updated                             |
| `WebhookEventNameEntitlementDeleted`            | entitlement.deleted                             |
| `WebhookEventNameWalletCreated`                 | wallet.created                                  |
| `WebhookEventNameWalletUpdated`                 | wallet.updated                                  |
| `WebhookEventNameWalletTerminated`              | wallet.terminated                               |
| `WebhookEventNameWalletTransactionCreated`      | wallet.transaction.created                      |
| `WebhookEventNamePaymentCreated`                | payment.created                                 |
| `WebhookEventNamePaymentUpdated`                | payment.updated                                 |
| `WebhookEventNamePaymentFailed`                 | payment.failed                                  |
| `WebhookEventNamePaymentSuccess`                | payment.success                                 |
| `WebhookEventNamePaymentPending`                | payment.pending                                 |
| `WebhookEventNameCustomerCreated`               | customer.created                                |
| `WebhookEventNameCustomerUpdated`               | customer.updated                                |
| `WebhookEventNameCustomerDeleted`               | customer.deleted                                |
| `WebhookEventNameInvoiceUpdateFinalized`        | invoice.update.finalized                        |
| `WebhookEventNameInvoiceUpdatePayment`          | invoice.update.payment                          |
| `WebhookEventNameInvoiceUpdateVoided`           | invoice.update.voided                           |
| `WebhookEventNameInvoiceUpdate`                 | invoice.update                                  |
| `WebhookEventNameInvoicePaymentOverdue`         | invoice.payment.overdue                         |
| `WebhookEventNameWalletCreditBalanceDropped`    | wallet.credit_balance.dropped                   |
| `WebhookEventNameWalletCreditBalanceRecovered`  | wallet.credit_balance.recovered                 |
| `WebhookEventNameWalletOngoingBalanceDropped`   | wallet.ongoing_balance.dropped                  |
| `WebhookEventNameWalletOngoingBalanceRecovered` | wallet.ongoing_balance.recovered                |
| `WebhookEventNameSubscriptionRenewalDue`        | subscription.renewal.due                        |
| `WebhookEventNameInvoiceCommunicationTriggered` | invoice.communication.triggered                 |
| `WebhookEventNameCreditNoteCreated`             | credit_note.created                             |
| `WebhookEventNameCreditNoteUpdated`             | credit_note.updated                             |