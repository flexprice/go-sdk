# CommitmentStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CommitmentStatusActive

// Open enum: custom values can be created with a direct type cast
custom := types.CommitmentStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CommitmentStatusActive`     | active                       |
| `CommitmentStatusEnded`      | ended                        |
| `CommitmentStatusTerminated` | terminated                   |