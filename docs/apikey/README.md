# APIKey

## Overview

These methods allow you to create and manage API keys for your Neon account. For related information, see [Manage API keys](https://neon.tech/docs/manage/api-keys).

### Available Operations

* [CreateAPIKey](#createapikey) - Create an API key
* [ListAPIKeys](#listapikeys) - Get a list of API keys
* [RevokeAPIKey](#revokeapikey) - Revoke an API key

## CreateAPIKey

Creates an API key.
The `key_name` is a user-specified name for the key.
This method returns an `id` and `key`. The `key` is a randomly generated, 64-bit token required to access the Neon API.
API keys can also be managed in the Neon Console.
See [Manage API keys](https://neon.tech/docs/manage/api-keys/).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIKey.CreateAPIKey(ctx, shared.APIKeyCreateRequest{
        KeyName: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIKeyCreateResponse != nil {
        // handle response
    }
}
```

## ListAPIKeys

Retrieves the API keys for your Neon account.
The response does not include API key tokens. A token is only provided when creating an API key.
API keys can also be managed in the Neon Console.
For more information, see [Manage API keys](https://neon.tech/docs/manage/api-keys/).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIKey.ListAPIKeys(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIKeysListResponseItems != nil {
        // handle response
    }
}
```

## RevokeAPIKey

Revokes the specified API key.
An API key that is no longer needed can be revoked.
This action cannot be reversed.
You can obtain `key_id` values by listing the API keys for your Neon account.
API keys can also be managed in the Neon Console.
See [Manage API keys](https://neon.tech/docs/manage/api-keys/).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIKey.RevokeAPIKey(ctx, 715190)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIKeyRevokeResponse != nil {
        // handle response
    }
}
```
