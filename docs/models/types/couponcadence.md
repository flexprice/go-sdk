# CouponCadence

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CouponCadenceOnce

// Open enum: custom values can be created with a direct type cast
custom := types.CouponCadence("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `CouponCadenceOnce`     | once                    |
| `CouponCadenceRepeated` | repeated                |
| `CouponCadenceForever`  | forever                 |