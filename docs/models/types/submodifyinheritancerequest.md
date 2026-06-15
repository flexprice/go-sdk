# SubModifyInheritanceRequest


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Action`                                                            | [*types.InheritanceAction](../../models/types/inheritanceaction.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `ExternalCustomerIdsToInheritSubscription`                          | []`string`                                                          | :heavy_minus_sign:                                                  | ExternalCustomerIDsToInheritSubscription is used for action="add".  |
| `ExternalCustomerIdsToRemove`                                       | []`string`                                                          | :heavy_minus_sign:                                                  | ExternalCustomerIDsToRemove is used for action="remove".            |