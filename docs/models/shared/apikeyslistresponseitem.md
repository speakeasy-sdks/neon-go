# APIKeysListResponseItem


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `CreatedAt`                                         | [time.Time](https://pkg.go.dev/time#Time)           | :heavy_check_mark:                                  | A timestamp indicating when the API key was created |
| `ID`                                                | *int64*                                             | :heavy_check_mark:                                  | The API key ID                                      |
| `LastUsedAt`                                        | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | A timestamp indicating when the API was last used   |
| `LastUsedFromAddr`                                  | *string*                                            | :heavy_check_mark:                                  | The IP address from which the API key was last used |
| `Name`                                              | *string*                                            | :heavy_check_mark:                                  | The user-specified API key name                     |