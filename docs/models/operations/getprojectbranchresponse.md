# GetProjectBranchResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `BranchResponse`                                                | [*shared.BranchResponse](../../models/shared/branchresponse.md) | :heavy_minus_sign:                                              | Returned information about the specified branch                 |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `GeneralError`                                                  | [*shared.GeneralError](../../models/shared/generalerror.md)     | :heavy_minus_sign:                                              | General Error                                                   |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |