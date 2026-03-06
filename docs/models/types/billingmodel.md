# BillingModel

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.BillingModelFlatFee

// Open enum: custom values can be created with a direct type cast
custom := types.BillingModel("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `BillingModelFlatFee` | FLAT_FEE              |
| `BillingModelPackage` | PACKAGE               |
| `BillingModelTiered`  | TIERED                |