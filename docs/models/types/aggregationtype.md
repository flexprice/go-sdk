# AggregationType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.AggregationTypeCount

// Open enum: custom values can be created with a direct type cast
custom := types.AggregationType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AggregationTypeCount`             | COUNT                              |
| `AggregationTypeSum`               | SUM                                |
| `AggregationTypeAvg`               | AVG                                |
| `AggregationTypeCountUnique`       | COUNT_UNIQUE                       |
| `AggregationTypeLatest`            | LATEST                             |
| `AggregationTypeSumWithMultiplier` | SUM_WITH_MULTIPLIER                |
| `AggregationTypeMax`               | MAX                                |
| `AggregationTypeWeightedSum`       | WEIGHTED_SUM                       |