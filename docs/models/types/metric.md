# Metric

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.MetricUsageQuantity

// Open enum: custom values can be created with a direct type cast
custom := types.Metric("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `MetricUsageQuantity`      | usage_quantity             |
| `MetricEventCount`         | event_count                |
| `MetricEventsVolume`       | events_volume              |
| `MetricInvoiceTotal`       | invoice_total              |
| `MetricTotalReceivables`   | total_receivables          |
| `MetricReceivablesAging`   | receivables_aging          |
| `MetricDaysToPay`          | days_to_pay                |
| `MetricArpu`               | arpu                       |
| `MetricArpuPaid`           | arpu_paid                  |
| `MetricCreditConsumption`  | credit_consumption         |
| `MetricCreditTopups`       | credit_topups              |
| `MetricNetBilling`         | net_billing                |
| `MetricInvoiceTotalIssued` | invoice_total_issued       |
| `MetricNetBillingIssued`   | net_billing_issued         |
| `MetricActiveCustomers`    | active_customers           |