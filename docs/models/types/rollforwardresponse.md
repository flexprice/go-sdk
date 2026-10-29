# RollforwardResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Items`                                                               | [][types.RollforwardRow](../../models/types/rollforwardrow.md)        | :heavy_minus_sign:                                                    | per period, account and group; empty without group_by                 |
| `Pagination`                                                          | [*types.PaginationResponse](../../models/types/paginationresponse.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Totals`                                                              | [][types.RollforwardRow](../../models/types/rollforwardrow.md)        | :heavy_minus_sign:                                                    | per period and account, all groups                                    |