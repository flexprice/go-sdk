# GCPMarketplaceAgreement


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AccountID`                                                         | `string`                                                            | :heavy_check_mark:                                                  | writes the customer mapping; not read in the report payload         |
| `MetricName`                                                        | `string`                                                            | :heavy_check_mark:                                                  | -> services.report's metricName (always "{service_name}/usage_fee") |
| `ServiceName`                                                       | `string`                                                            | :heavy_check_mark:                                                  | -> services.report URL's service_name; identifies the product       |
| `UsageReportingID`                                                  | `string`                                                            | :heavy_check_mark:                                                  | -> services.report's consumerId; identifies the buyer               |