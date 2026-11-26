# PaymentAttemptMethod

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentAttemptMethodCard

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentAttemptMethod("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PaymentAttemptMethodCard`         | CARD                               |
| `PaymentAttemptMethodUpi`          | UPI                                |
| `PaymentAttemptMethodNetbanking`   | NETBANKING                         |
| `PaymentAttemptMethodWallet`       | WALLET                             |
| `PaymentAttemptMethodEmi`          | EMI                                |
| `PaymentAttemptMethodPaylater`     | PAYLATER                           |
| `PaymentAttemptMethodBankTransfer` | BANK_TRANSFER                      |
| `PaymentAttemptMethodOther`        | OTHER                              |