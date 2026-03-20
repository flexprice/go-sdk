# InvoiceStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.InvoiceStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := types.InvoiceStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `InvoiceStatusDraft`     | DRAFT                    |
| `InvoiceStatusFinalized` | FINALIZED                |
| `InvoiceStatusVoided`    | VOIDED                   |