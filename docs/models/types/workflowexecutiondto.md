# WorkflowExecutionDTO


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CloseTime`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `CreatedBy`                                | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `DurationMs`                               | `*int64`                                   | :heavy_minus_sign:                         | N/A                                        |
| `Entity`                                   | `*string`                                  | :heavy_minus_sign:                         | e.g. plan, invoice, subscription           |
| `EntityID`                                 | `*string`                                  | :heavy_minus_sign:                         | e.g. plan ID, invoice ID                   |
| `RunID`                                    | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `StartTime`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Status`                                   | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `TaskQueue`                                | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `TotalDuration`                            | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `WorkflowID`                               | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `WorkflowType`                             | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |