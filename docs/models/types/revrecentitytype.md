# RevRecEntityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RevRecEntityTypeSubscription

// Open enum: custom values can be created with a direct type cast
custom := types.RevRecEntityType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `RevRecEntityTypeSubscription`         | subscription                           |
| `RevRecEntityTypeSubscriptionLineItem` | subscription_line_item                 |
| `RevRecEntityTypeInvoice`              | invoice                                |
| `RevRecEntityTypeCreditNote`           | credit_note                            |
| `RevRecEntityTypePayment`              | payment                                |
| `RevRecEntityTypeRefund`               | refund                                 |
| `RevRecEntityTypePob`                  | pob                                    |