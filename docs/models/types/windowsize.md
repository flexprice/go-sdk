# WindowSize

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.WindowSizeMonth

// Open enum: custom values can be created with a direct type cast
custom := types.WindowSize("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `WindowSizeMonth`      | MONTH                  |
| `WindowSizeMinute`     | MINUTE                 |
| `WindowSizeFifteenMin` | 15MIN                  |
| `WindowSizeThirtyMin`  | 30MIN                  |
| `WindowSizeHour`       | HOUR                   |
| `WindowSizeThreeHour`  | 3HOUR                  |
| `WindowSizeSixHour`    | 6HOUR                  |
| `WindowSizeTwelveHour` | 12HOUR                 |
| `WindowSizeDay`        | DAY                    |
| `WindowSizeWeek`       | WEEK                   |