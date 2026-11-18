# EmailDeliveryStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.EmailDeliveryStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := types.EmailDeliveryStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `EmailDeliveryStatusQueued`          | queued                               |
| `EmailDeliveryStatusSent`            | sent                                 |
| `EmailDeliveryStatusDeliveryDelayed` | delivery_delayed                     |
| `EmailDeliveryStatusDelivered`       | delivered                            |
| `EmailDeliveryStatusBounced`         | bounced                              |
| `EmailDeliveryStatusComplained`      | complained                           |
| `EmailDeliveryStatusFailed`          | failed                               |
| `EmailDeliveryStatusSuppressed`      | suppressed                           |
| `EmailDeliveryStatusSkipped`         | skipped                              |