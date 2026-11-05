# UpdateCommitmentRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `EndDate`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | EndDate can't be changed yet; terminate the commitment to end it. |
| `InvoiceDescription`                                              | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `InvoiceLabel`                                                    | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Metadata`                                                        | map[string]`string`                                               | :heavy_minus_sign:                                                | N/A                                                               |
| `Name`                                                            | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |