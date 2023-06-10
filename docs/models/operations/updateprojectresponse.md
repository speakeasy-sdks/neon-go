# UpdateProjectResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `GeneralError`                                                        | [*shared.GeneralError](../../models/shared/generalerror.md)           | :heavy_minus_sign:                                                    | General Error                                                         |
| `ProjectOperations`                                                   | [*shared.ProjectOperations](../../models/shared/projectoperations.md) | :heavy_minus_sign:                                                    | Updated the specified project                                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |