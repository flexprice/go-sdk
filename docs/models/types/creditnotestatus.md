# CreditNoteStatus

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.CreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := types.CreditNoteStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `CreditNoteStatusDraft`     | DRAFT                       |
| `CreditNoteStatusFinalized` | FINALIZED                   |
| `CreditNoteStatusVoided`    | VOIDED                      |