# S3EncryptionType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.S3EncryptionTypeAes256

// Open enum: custom values can be created with a direct type cast
custom := types.S3EncryptionType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `S3EncryptionTypeAes256`     | AES256                       |
| `S3EncryptionTypeAwsKms`     | aws:kms                      |
| `S3EncryptionTypeAwsKmsDsse` | aws:kms:dsse                 |