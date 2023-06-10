# GetProjectBranchRoleResponse


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ContentType`                                               | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `GeneralError`                                              | [*shared.GeneralError](../../models/shared/generalerror.md) | :heavy_minus_sign:                                          | General Error                                               |
| `RoleResponse`                                              | [*shared.RoleResponse](../../models/shared/roleresponse.md) | :heavy_minus_sign:                                          | Successfully returned details for the specified role        |
| `StatusCode`                                                | *int*                                                       | :heavy_check_mark:                                          | N/A                                                         |
| `RawResponse`                                               | [*http.Response](https://pkg.go.dev/net/http#Response)      | :heavy_minus_sign:                                          | N/A                                                         |