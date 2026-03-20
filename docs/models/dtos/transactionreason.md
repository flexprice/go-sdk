# TransactionReason

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/dtos"
)

value := dtos.TransactionReasonInvoicePayment
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `TransactionReasonInvoicePayment`          | INVOICE_PAYMENT                            |
| `TransactionReasonFreeCreditGrant`         | FREE_CREDIT_GRANT                          |
| `TransactionReasonSubscriptionCreditGrant` | SUBSCRIPTION_CREDIT_GRANT                  |
| `TransactionReasonPurchasedCreditInvoiced` | PURCHASED_CREDIT_INVOICED                  |
| `TransactionReasonPurchasedCreditDirect`   | PURCHASED_CREDIT_DIRECT                    |
| `TransactionReasonCreditNote`              | CREDIT_NOTE                                |
| `TransactionReasonCreditExpired`           | CREDIT_EXPIRED                             |
| `TransactionReasonWalletTermination`       | WALLET_TERMINATION                         |
| `TransactionReasonManualBalanceDebit`      | MANUAL_BALANCE_DEBIT                       |
| `TransactionReasonCreditAdjustment`        | CREDIT_ADJUSTMENT                          |
| `TransactionReasonInvoiceVoidRefund`       | INVOICE_VOID_REFUND                        |