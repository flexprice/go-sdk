# SubModifyGroupedInvoicingParams


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Action`                                                                     | [types.GroupedInvoicingAction](../../models/types/groupedinvoicingaction.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `ChildSubscriptionIds`                                                       | []`string`                                                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `ParentSubscriptionID`                                                       | `*string`                                                                    | :heavy_minus_sign:                                                           | ParentSubscriptionID is required for action 'add'.                           |