# PaymentDestinationType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentDestinationTypeInvoice

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentDestinationType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PaymentDestinationTypeInvoice`  | INVOICE                          |
| `PaymentDestinationTypeCustomer` | CUSTOMER                         |