# RefundDestination

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RefundDestinationGateway

// Open enum: custom values can be created with a direct type cast
custom := types.RefundDestination("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `RefundDestinationGateway`   | GATEWAY                      |
| `RefundDestinationWallet`    | WALLET                       |
| `RefundDestinationOutOfBand` | OUT_OF_BAND                  |