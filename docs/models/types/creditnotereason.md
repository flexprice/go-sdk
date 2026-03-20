# CreditNoteReason

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CreditNoteReasonDuplicate

// Open enum: custom values can be created with a direct type cast
custom := types.CreditNoteReason("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreditNoteReasonDuplicate`                | DUPLICATE                                  |
| `CreditNoteReasonFraudulent`               | FRAUDULENT                                 |
| `CreditNoteReasonOrderChange`              | ORDER_CHANGE                               |
| `CreditNoteReasonUnsatisfactory`           | UNSATISFACTORY                             |
| `CreditNoteReasonServiceIssue`             | SERVICE_ISSUE                              |
| `CreditNoteReasonBillingError`             | BILLING_ERROR                              |
| `CreditNoteReasonSubscriptionCancellation` | SUBSCRIPTION_CANCELLATION                  |