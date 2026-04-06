# Code

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CodeNotFound

// Open enum: custom values can be created with a direct type cast
custom := types.Code("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `CodeNotFound`           | not_found                |
| `CodeAlreadyExists`      | already_exists           |
| `CodeVersionConflict`    | version_conflict         |
| `CodeValidationError`    | validation_error         |
| `CodeInvalidOperation`   | invalid_operation        |
| `CodePermissionDenied`   | permission_denied        |
| `CodeHTTPClientError`    | http_client_error        |
| `CodeDatabaseError`      | database_error           |
| `CodeSystemError`        | system_error             |
| `CodeInternalError`      | internal_error           |
| `CodeServiceUnavailable` | service_unavailable      |