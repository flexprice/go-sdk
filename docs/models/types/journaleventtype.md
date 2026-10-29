# JournalEventType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.JournalEventTypeInvoiceFinalized

// Open enum: custom values can be created with a direct type cast
custom := types.JournalEventType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `JournalEventTypeInvoiceFinalized`    | invoice_finalized                     |
| `JournalEventTypeInvoiceVoided`       | invoice_voided                        |
| `JournalEventTypeCreditNoteFinalized` | credit_note_finalized                 |
| `JournalEventTypeCreditNoteVoided`    | credit_note_voided                    |
| `JournalEventTypePaymentSucceeded`    | payment_succeeded                     |
| `JournalEventTypePaymentReversed`     | payment_reversed                      |
| `JournalEventTypeRefundSucceeded`     | refund_succeeded                      |
| `JournalEventTypeRevenueEarned`       | revenue_earned                        |
| `JournalEventTypeRevenueOffset`       | revenue_offset                        |