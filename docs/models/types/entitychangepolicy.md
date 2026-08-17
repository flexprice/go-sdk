# EntityChangePolicy


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `DefaultBehaviour`                                                                    | [*types.EntityChangeBehaviour](../../models/types/entitychangebehaviour.md)           | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `Overrides`                                                                           | map[string][types.EntityChangeBehaviour](../../models/types/entitychangebehaviour.md) | :heavy_minus_sign:                                                                    | Overrides is keyed by addon_associations.id (instance), not catalogue addon_id.       |