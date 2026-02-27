# DtoCreateMeterRequest


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Aggregation`                                          | [types.MeterAggregation](../types/meteraggregation.md) | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `EventName`                                            | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | api_request                                            |
| `Filters`                                              | [][types.MeterFilter](../types/meterfilter.md)         | :heavy_minus_sign:                                     | N/A                                                    |                                                        |
| `Name`                                                 | *string*                                               | :heavy_check_mark:                                     | N/A                                                    | API Usage Meter                                        |
| `ResetUsage`                                           | [types.ResetUsage](../types/resetusage.md)             | :heavy_check_mark:                                     | N/A                                                    |                                                        |