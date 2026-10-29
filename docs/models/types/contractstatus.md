# ContractStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ContractStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := types.ContractStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ContractStatusDraft`     | draft                     |
| `ContractStatusActive`    | active                    |
| `ContractStatusEnded`     | ended                     |
| `ContractStatusCancelled` | cancelled                 |