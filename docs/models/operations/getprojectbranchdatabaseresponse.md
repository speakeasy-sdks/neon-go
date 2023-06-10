# GetProjectBranchDatabaseResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `DatabaseResponse`                                                  | [*shared.DatabaseResponse](../../models/shared/databaseresponse.md) | :heavy_minus_sign:                                                  | Returned the database details                                       |
| `GeneralError`                                                      | [*shared.GeneralError](../../models/shared/generalerror.md)         | :heavy_minus_sign:                                                  | General Error                                                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |