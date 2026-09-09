# ListRefundsRequest


## Fields

| Field                          | Type                           | Required                       | Description                    |
| ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ |
| `InvoiceIds`                   | []`string`                     | :heavy_minus_sign:             | Filter by invoice IDs          |
| `PaymentIds`                   | []`string`                     | :heavy_minus_sign:             | Filter by payment IDs          |
| `CreditNoteIds`                | []`string`                     | :heavy_minus_sign:             | Filter by credit note IDs      |
| `RefundStatuses`               | []`string`                     | :heavy_minus_sign:             | Filter by refund status        |
| `RefundDestinations`           | []`string`                     | :heavy_minus_sign:             | Filter by refund destination   |
| `Gateway`                      | `*string`                      | :heavy_minus_sign:             | Filter by payment gateway      |
| `OnlySettled`                  | `*bool`                        | :heavy_minus_sign:             | Only refunds that have settled |
| `Limit`                        | `*int64`                       | :heavy_minus_sign:             | Limit                          |
| `Offset`                       | `*int64`                       | :heavy_minus_sign:             | Offset                         |