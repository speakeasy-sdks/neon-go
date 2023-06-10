# CreateProjectBranchDatabaseResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `DatabaseOperations`                                                    | [*shared.DatabaseOperations](../../models/shared/databaseoperations.md) | :heavy_minus_sign:                                                      | Created a database in the specified branch                              |
| `GeneralError`                                                          | [*shared.GeneralError](../../models/shared/generalerror.md)             | :heavy_minus_sign:                                                      | General Error                                                           |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |