# InstalmentStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.InstalmentStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := types.InstalmentStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `InstalmentStatusScheduled` | scheduled                   |
| `InstalmentStatusInvoiced`  | invoiced                    |
| `InstalmentStatusCancelled` | cancelled                   |