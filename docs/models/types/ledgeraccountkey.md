# LedgerAccountKey

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.LedgerAccountKeyCash

// Open enum: custom values can be created with a direct type cast
custom := types.LedgerAccountKey("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `LedgerAccountKeyCash`               | cash                                 |
| `LedgerAccountKeyAccountsReceivable` | accounts_receivable                  |
| `LedgerAccountKeyUnbilledReceivable` | unbilled_receivable                  |
| `LedgerAccountKeyDeferredRevenue`    | deferred_revenue                     |
| `LedgerAccountKeyCustomerCredit`     | customer_credit                      |
| `LedgerAccountKeyTaxPayable`         | tax_payable                          |
| `LedgerAccountKeyRevenue`            | revenue                              |
| `LedgerAccountKeyContraRevenue`      | contra_revenue                       |