# ListProjectEndpointsResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `EndpointsResponse`                                                   | [*shared.EndpointsResponse](../../models/shared/endpointsresponse.md) | :heavy_minus_sign:                                                    | Returned a list of endpoints for the specified project                |
| `GeneralError`                                                        | [*shared.GeneralError](../../models/shared/generalerror.md)           | :heavy_minus_sign:                                                    | General Error                                                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |