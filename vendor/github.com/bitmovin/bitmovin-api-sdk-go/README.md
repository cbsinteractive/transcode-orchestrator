<p align="center">
  <a href="https://www.bitmovin.com">
    <img alt="Bitmovin Go API SDK Header" src="https://cdn.bitmovin.com/frontend/encoding/openapi-clients/readme-headers/ReadmeHeader_Go.png" >
  </a>

  <h4 align="center">
    Go API SDK which enables you to seamlessly integrate the Bitmovin API into your projects.
  </h4>

  <p align="center">
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></img></a>
  </p>
</p>

***Important! This is an alpha version, there may be breaking changes.***

Using this API client requires an active account.

> Don't have an account yet? [Sign up for a free Bitmovin trial plan](https://dashboard.bitmovin.com/signup)!

For full documentation of all available API endpoints, see the [Bitmovin API reference](https://bitmovin.com/docs).

## Installation:

```bash
go get github.com/bitmovin/bitmovin-api-sdk-go
```

## Initializiation

```go
package main

import (
    "github.com/bitmovin/bitmovin-api-sdk-go"
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

func main() {
    bitmovinApi, err := bitmovin.NewBitmovinApi(func(apiClient *common.ApiClient) {
        apiClient.ApiKey = "<YOUR_API_KEY>"
    })
    if err != nil {
        panic(err)
    }

   // ...
}
```
