# ListProjectBranchesResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `BranchesResponse`                                                  | [*shared.BranchesResponse](../../models/shared/branchesresponse.md) | :heavy_minus_sign:                                                  | Returned a list of branches for the specified project               |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `GeneralError`                                                      | [*shared.GeneralError](../../models/shared/generalerror.md)         | :heavy_minus_sign:                                                  | General Error                                                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |