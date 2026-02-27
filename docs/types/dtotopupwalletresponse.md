# DtoTopUpWalletResponse


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `InvoiceID`                                                                     | **string*                                                                       | :heavy_minus_sign:                                                              | Invoice ID if an invoice was created (only for PURCHASED_CREDIT_INVOICED)       |
| `Wallet`                                                                        | [*types.DtoWalletResponse](../types/dtowalletresponse.md)                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `WalletTransaction`                                                             | [*types.DtoWalletTransactionResponse](../types/dtowallettransactionresponse.md) | :heavy_minus_sign:                                                              | N/A                                                                             |