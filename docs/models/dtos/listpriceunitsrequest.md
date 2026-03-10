# ListPriceUnitsRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Status`                                                              | `*string`                                                             | :heavy_minus_sign:                                                    | Filter by status                                                      |
| `Limit`                                                               | `*int64`                                                              | :heavy_minus_sign:                                                    | Limit number of results                                               |
| `Offset`                                                              | `*int64`                                                              | :heavy_minus_sign:                                                    | Offset for pagination                                                 |
| `Sort`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Sort field                                                            |
| `Order`                                                               | [*dtos.ListPriceUnitsOrder](../../models/dtos/listpriceunitsorder.md) | :heavy_minus_sign:                                                    | Sort order                                                            |