# PaymentTermsType

DAYS (due payment_terms_days after the issue date), DUE_ON_RECEIPT, END_OF_MONTH or END_OF_NEXT_MONTH.

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentTermsTypeDays

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentTermsType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PaymentTermsTypeDays`           | DAYS                             |
| `PaymentTermsTypeDueOnReceipt`   | DUE_ON_RECEIPT                   |
| `PaymentTermsTypeEndOfMonth`     | END_OF_MONTH                     |
| `PaymentTermsTypeEndOfNextMonth` | END_OF_NEXT_MONTH                |