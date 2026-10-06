# CreateFXRateRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `EndDate`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `FromCurrency`                                            | `string`                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `Metadata`                                                | map[string]`string`                                       | :heavy_minus_sign:                                        | N/A                                                       |
| `Rate`                                                    | `*string`                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Scope`                                                   | [types.FXRateScope](../../models/types/fxratescope.md)    | :heavy_check_mark:                                        | N/A                                                       |
| `ScopeID`                                                 | `*string`                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Source`                                                  | [*types.FXRateSource](../../models/types/fxratesource.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `StartDate`                                               | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `ToCurrency`                                              | `string`                                                  | :heavy_check_mark:                                        | N/A                                                       |