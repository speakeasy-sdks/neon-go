# DeleteProjectResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `GeneralError`                                                    | [*shared.GeneralError](../../models/shared/generalerror.md)       | :heavy_minus_sign:                                                | General Error                                                     |
| `ProjectResponse`                                                 | [*shared.ProjectResponse](../../models/shared/projectresponse.md) | :heavy_minus_sign:                                                | Deleted the specified project                                     |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | N/A                                                               |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | N/A                                                               |