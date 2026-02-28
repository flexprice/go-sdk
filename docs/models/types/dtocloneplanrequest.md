# DtoClonePlanRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Description`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | Description overrides the source plan's description when provided    |
| `DisplayOrder`                                                       | **int64*                                                             | :heavy_minus_sign:                                                   | DisplayOrder overrides the source plan's display order when provided |
| `LookupKey`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | LookupKey is required and must be unique across published plans      |
| `Metadata`                                                           | map[string]*string*                                                  | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Name`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Name is required and must be different from the source plan's name   |