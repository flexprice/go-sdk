# BillingPeriod

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.BillingPeriodMonthly

// Open enum: custom values can be created with a direct type cast
custom := types.BillingPeriod("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `BillingPeriodMonthly`    | MONTHLY                   |
| `BillingPeriodAnnual`     | ANNUAL                    |
| `BillingPeriodWeekly`     | WEEKLY                    |
| `BillingPeriodDaily`      | DAILY                     |
| `BillingPeriodQuarterly`  | QUARTERLY                 |
| `BillingPeriodHalfYearly` | HALF_YEARLY               |
| `BillingPeriodOnetime`    | ONETIME                   |