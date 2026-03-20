# WalletTxReferenceType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.WalletTxReferenceTypePayment

// Open enum: custom values can be created with a direct type cast
custom := types.WalletTxReferenceType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `WalletTxReferenceTypePayment`  | PAYMENT                         |
| `WalletTxReferenceTypeExternal` | EXTERNAL                        |
| `WalletTxReferenceTypeRequest`  | REQUEST                         |
| `WalletTxReferenceTypeInvoice`  | INVOICE                         |