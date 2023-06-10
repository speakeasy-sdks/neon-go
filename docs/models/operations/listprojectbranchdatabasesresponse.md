# ListProjectBranchDatabasesResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `DatabasesResponse`                                                   | [*shared.DatabasesResponse](../../models/shared/databasesresponse.md) | :heavy_minus_sign:                                                    | Returned a list of databases of the specified branch                  |
| `GeneralError`                                                        | [*shared.GeneralError](../../models/shared/generalerror.md)           | :heavy_minus_sign:                                                    | General Error                                                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |