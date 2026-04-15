# ChangedInvoiceAction

created (proration invoice) | wallet_credit (downgrade credit)

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ChangedInvoiceActionCreated

// Open enum: custom values can be created with a direct type cast
custom := types.ChangedInvoiceAction("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ChangedInvoiceActionCreated`      | created                            |
| `ChangedInvoiceActionWalletCredit` | wallet_credit                      |