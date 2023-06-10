# UpdateProjectEndpointResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `EndpointOperations`                                                    | [*shared.EndpointOperations](../../models/shared/endpointoperations.md) | :heavy_minus_sign:                                                      | Updated the specified endpoint                                          |
| `GeneralError`                                                          | [*shared.GeneralError](../../models/shared/generalerror.md)             | :heavy_minus_sign:                                                      | General Error                                                           |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |