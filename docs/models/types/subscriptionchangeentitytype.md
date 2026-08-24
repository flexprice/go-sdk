# SubscriptionChangeEntityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SubscriptionChangeEntityTypePlan

// Open enum: custom values can be created with a direct type cast
custom := types.SubscriptionChangeEntityType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `SubscriptionChangeEntityTypePlan`             | plan                                           |
| `SubscriptionChangeEntityTypeAddon`            | addon                                          |
| `SubscriptionChangeEntityTypeCreditGrant`      | credit_grant                                   |
| `SubscriptionChangeEntityTypeEntitlement`      | entitlement                                    |
| `SubscriptionChangeEntityTypeEntitlementGrant` | entitlement_grant                              |