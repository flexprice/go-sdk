# CloneFeatureRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Description`                                                         | `*string`                                                             | :heavy_minus_sign:                                                    | Description overrides the source feature's description when provided  |
| `LookupKey`                                                           | `*string`                                                             | :heavy_minus_sign:                                                    | LookupKey is required and must be unique across published features    |
| `Metadata`                                                            | map[string]`string`                                                   | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Name`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Name is required and must be different from the source feature's name |