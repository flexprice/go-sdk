# TriggerForceRunRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `CustomerExternalIds`                                                         | []`string`                                                                    | :heavy_check_mark:                                                            | Limits a wallet_ledger export to these customers; empty means every customer. |
| `EndTime`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `StartTime`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |