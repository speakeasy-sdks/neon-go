# GetProjectEndpointResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `EndpointResponse`                                                  | [*shared.EndpointResponse](../../models/shared/endpointresponse.md) | :heavy_minus_sign:                                                  | Returned information about the specified endpoint                   |
| `GeneralError`                                                      | [*shared.GeneralError](../../models/shared/generalerror.md)         | :heavy_minus_sign:                                                  | General Error                                                       |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |