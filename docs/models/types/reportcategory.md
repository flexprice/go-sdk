# ReportCategory

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ReportCategoryBilling

// Open enum: custom values can be created with a direct type cast
custom := types.ReportCategory("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ReportCategoryBilling`     | billing                     |
| `ReportCategoryReceivables` | receivables                 |
| `ReportCategoryCustomers`   | customers                   |
| `ReportCategoryCredits`     | credits                     |
| `ReportCategoryUsage`       | usage                       |