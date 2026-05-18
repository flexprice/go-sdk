# ErrorCode

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ErrorCodeHTTPClientError

// Open enum: custom values can be created with a direct type cast
custom := types.ErrorCode("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ErrorCodeHTTPClientError`    | http_client_error             |
| `ErrorCodeSystemError`        | system_error                  |
| `ErrorCodeInternalError`      | internal_error                |
| `ErrorCodeNotFound`           | not_found                     |
| `ErrorCodeAlreadyExists`      | already_exists                |
| `ErrorCodeVersionConflict`    | version_conflict              |
| `ErrorCodeValidationError`    | validation_error              |
| `ErrorCodeInvalidOperation`   | invalid_operation             |
| `ErrorCodePermissionDenied`   | permission_denied             |
| `ErrorCodeDatabaseError`      | database_error                |
| `ErrorCodeServiceUnavailable` | service_unavailable           |
| `ErrorCodeTooManyRequests`    | too_many_requests             |