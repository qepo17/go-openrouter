# go-openrouter

## Description

Unofficial Go client for the OpenRouter API.

## Installation

```bash
go get github.com/qepo17/go-openrouter
```

## Usage

```go
package main

import (
    "github.com/qepo17/go-openrouter"
)

func main() {
    openRouterSvc, err := openrouter.New(openRouterAPIKey, openrouter.ClientOptions{})
	if err != nil {
		panic(err)
	}

    ...
}
```
