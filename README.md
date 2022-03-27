# Fasapay XML API SDK GO (Unofficial)
[![Build Status](https://travis-ci.org/Kachit/dusupay-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/fasapay-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/dusupay-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/fasapay-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/dusupay-sdk-go)](https://goreportcard.com/report/github.com/kachit/fasapay-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/dusupay-sdk-go.svg)](https://github.com/Kachit/fasapay-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/fasapay-sdk-go/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/kachit/dusupay-sdk-go)](https://pkg.go.dev/github.com/kachit/fasapay-sdk-go)

## Description
Unofficial Fasapay payment gateway XML API Client for Go

## API documentation
https://www.fasapay.com/en/apiguide/index

## Installation
```shell
go get -u github.com/kachit/fasapay-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "context"
    fasapay "github.com/kachit/fasapay-sdk-go"
)

func main(){
    // Create a client instance
    cfg := fasapay.NewConfig("Your public key", "Your secret key")
    client, err := fasapay.NewClientFromConfig(cfg, nil)
    if err != nil {
        fmt.Printf("config parameter error " + err.Error())
        panic(err)
    }
}
```

