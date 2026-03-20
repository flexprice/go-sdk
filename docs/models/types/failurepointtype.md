# FailurePointType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.FailurePointTypeCustomerLookup

// Open enum: custom values can be created with a direct type cast
custom := types.FailurePointType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `FailurePointTypeCustomerLookup`             | customer_lookup                              |
| `FailurePointTypeMeterLookup`                | meter_lookup                                 |
| `FailurePointTypePriceLookup`                | price_lookup                                 |
| `FailurePointTypeSubscriptionLineItemLookup` | subscription_line_item_lookup                |