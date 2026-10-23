# TenantInternalStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.TenantInternalStatusTrialing

// Open enum: custom values can be created with a direct type cast
custom := types.TenantInternalStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `TenantInternalStatusTrialing`  | trialing                        |
| `TenantInternalStatusActive`    | active                          |
| `TenantInternalStatusSuspended` | suspended                       |