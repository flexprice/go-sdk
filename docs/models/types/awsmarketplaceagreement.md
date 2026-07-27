# AWSMarketplaceAgreement


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ConcurrentAgreements`                                                 | `*bool`                                                                | :heavy_minus_sign:                                                     | if true, ProductCode is omitted when reporting                         |
| `CustomerAwsAccountID`                                                 | `string`                                                               | :heavy_check_mark:                                                     | -> BatchMeterUsage's CustomerAWSAccountId                              |
| `Dimension`                                                            | `string`                                                               | :heavy_check_mark:                                                     | -> BatchMeterUsage's Dimension (always "usage_fee" in the cents model) |
| `LicenseArn`                                                           | `string`                                                               | :heavy_check_mark:                                                     | -> BatchMeterUsage's LicenseArn; identifies the buyer's agreement      |
| `ProductCode`                                                          | `string`                                                               | :heavy_check_mark:                                                     | -> BatchMeterUsage's ProductCode (omitted when ConcurrentAgreements)   |