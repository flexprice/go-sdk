# TaxRateEntityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.TaxRateEntityTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := types.TaxRateEntityType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `TaxRateEntityTypeCustomer`     | customer                        |
| `TaxRateEntityTypeSubscription` | subscription                    |
| `TaxRateEntityTypeInvoice`      | invoice                         |
| `TaxRateEntityTypeTenant`       | tenant                          |