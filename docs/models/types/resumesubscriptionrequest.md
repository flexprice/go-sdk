# ResumeSubscriptionRequest

Resumes a paused subscription now. The billing anchor is kept, so the first period runs from the resume time to the next billing date.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Metadata`                                                          | map[string]`string`                                                 | :heavy_minus_sign:                                                  | Metadata merged into the pause<br/>@Example {"resumed_by": "admin"} |
| `ProrationBehavior`                                                 | [*types.ProrationBehavior](../../models/types/prorationbehavior.md) | :heavy_minus_sign:                                                  | N/A                                                                 |