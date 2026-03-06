# EventProcessingStatusType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.EventProcessingStatusTypeProcessed

// Open enum: custom values can be created with a direct type cast
custom := types.EventProcessingStatusType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `EventProcessingStatusTypeProcessed`  | processed                             |
| `EventProcessingStatusTypeProcessing` | processing                            |
| `EventProcessingStatusTypeFailed`     | failed                                |