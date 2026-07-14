# RejectedEventReason

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RejectedEventReasonNoMeterForEventName

// Open enum: custom values can be created with a direct type cast
custom := types.RejectedEventReason("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RejectedEventReasonNoMeterForEventName` | no_meter_for_event_name                  |
| `RejectedEventReasonNoMatchingMeter`     | no_matching_meter                        |