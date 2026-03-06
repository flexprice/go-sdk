# IntegrationEntityType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.IntegrationEntityTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := types.IntegrationEntityType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `IntegrationEntityTypeCustomer`     | customer                            |
| `IntegrationEntityTypePlan`         | plan                                |
| `IntegrationEntityTypeInvoice`      | invoice                             |
| `IntegrationEntityTypeSubscription` | subscription                        |
| `IntegrationEntityTypePayment`      | payment                             |
| `IntegrationEntityTypeCreditNote`   | credit_note                         |
| `IntegrationEntityTypeAddon`        | addon                               |
| `IntegrationEntityTypeItem`         | item                                |
| `IntegrationEntityTypeItemPrice`    | item_price                          |
| `IntegrationEntityTypePrice`        | price                               |