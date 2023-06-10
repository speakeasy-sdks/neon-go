# ListProjectBranchRolesResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `GeneralError`                                                | [*shared.GeneralError](../../models/shared/generalerror.md)   | :heavy_minus_sign:                                            | General Error                                                 |
| `RolesResponse`                                               | [*shared.RolesResponse](../../models/shared/rolesresponse.md) | :heavy_minus_sign:                                            | Returned a list of roles from the specified branch.           |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |