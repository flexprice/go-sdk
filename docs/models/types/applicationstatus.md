# ApplicationStatus

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.ApplicationStatusApplied

// Open enum: custom values can be created with a direct type cast
custom := types.ApplicationStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ApplicationStatusApplied`   | applied                      |
| `ApplicationStatusFailed`    | failed                       |
| `ApplicationStatusPending`   | pending                      |
| `ApplicationStatusSkipped`   | skipped                      |
| `ApplicationStatusCancelled` | cancelled                    |