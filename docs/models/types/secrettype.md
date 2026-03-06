# SecretType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.SecretTypePrivateKey

// Open enum: custom values can be created with a direct type cast
custom := types.SecretType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SecretTypePrivateKey`     | private_key                |
| `SecretTypePublishableKey` | publishable_key            |
| `SecretTypeIntegration`    | integration                |