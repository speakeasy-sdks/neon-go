# GetProjectOperationResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `GeneralError`                                                        | [*shared.GeneralError](../../models/shared/generalerror.md)           | :heavy_minus_sign:                                                    | General Error                                                         |
| `OperationResponse`                                                   | [*shared.OperationResponse](../../models/shared/operationresponse.md) | :heavy_minus_sign:                                                    | Returned details for the specified operation                          |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |