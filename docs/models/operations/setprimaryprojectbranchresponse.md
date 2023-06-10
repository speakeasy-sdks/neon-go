# SetPrimaryProjectBranchResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BranchOperations`                                                  | [*shared.BranchOperations](../../models/shared/branchoperations.md) | :heavy_minus_sign:                                                  | Updated the specified branch                                        |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `GeneralError`                                                      | [*shared.GeneralError](../../models/shared/generalerror.md)         | :heavy_minus_sign:                                                  | General Error                                                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |