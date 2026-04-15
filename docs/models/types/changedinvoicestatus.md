# ChangedInvoiceStatus

preview | issued | INITIATED | PENDING | PROCESSING | SUCCEEDED | OVERPAID | FAILED | REFUNDED | PARTIALLY_REFUNDED

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ChangedInvoiceStatusPreview

// Open enum: custom values can be created with a direct type cast
custom := types.ChangedInvoiceStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ChangedInvoiceStatusPreview` | preview                       |
| `ChangedInvoiceStatusIssued`  | issued                        |