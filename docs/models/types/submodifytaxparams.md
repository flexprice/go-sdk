# SubModifyTaxParams


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Action`                                                                | [types.SubModifyTaxAction](../../models/types/submodifytaxaction.md)    | :heavy_check_mark:                                                      | N/A                                                                     |
| `AssociationID`                                                         | `*string`                                                               | :heavy_minus_sign:                                                      | Required when action="remove". ID of the TaxAssociation to soft-delete. |
| `EffectiveDate`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | Optional. When to apply the change; defaults to now if omitted.         |
| `TaxRateID`                                                             | `*string`                                                               | :heavy_minus_sign:                                                      | Required when action="add". ID of the active tax rate to attach.        |