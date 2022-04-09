# Fasapay XML API SDK GO (Unofficial)
[![Build Status](https://app.travis-ci.com/Kachit/fasapay-sdk-go.svg?branch=master)](https://app.travis-ci.com/github/Kachit/fasapay-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/fasapay-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/fasapay-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/fasapay-sdk-go)](https://goreportcard.com/report/github.com/kachit/fasapay-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/fasapay-sdk-go.svg)](https://github.com/Kachit/fasapay-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/fasapay-sdk-go/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/kachit/fasapay-sdk-go)](https://pkg.go.dev/github.com/kachit/fasapay-sdk-go)

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
    cfg := fasapay.NewConfig("Your API key", "Your API secret word")
    client, err := fasapay.NewClientFromConfig(cfg, nil)
    if err != nil {
        fmt.Printf("config parameter error " + err.Error())
        panic(err)
    }
}
```
### Get balances list
```go
ctx := context.Background()
currencies := []fasapay.CurrencyCode{fasapay.CurrencyCodeIDR, fasapay.CurrencyCodeUSD}
result, resp, err := client.Accounts().GetBalances(currencies, ctx, nil)

if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(response)

//Dump result
fmt.Println(result.Balances.IDR)
fmt.Println(result.Balances.USD)
```
### Get accounts list
```go
ctx := context.Background()
accounts := []string{"FP0000001", "FP0000002"}
result, resp, err := client.Accounts().GetAccounts(accounts, ctx, nil)

if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(response)

//Dump result
fmt.Println(result.Accounts[0].FullName)
fmt.Println(result.Accounts[0].Account)
fmt.Println(result.Accounts[0].Status)
```

### Create transfer
```go
ctx := context.Background()
transfer1 := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}

transfer2 := &CreateTransferRequestParams{
		Id:       "1234",
		To:       "FP89681",
		Amount:   1001.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
result, resp, err := client.Transfers().CreateTransfer([]*CreateTransferRequestParams{transfer, transfer2}, ctx, nil)

if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(response)

//Dump result
fmt.Println(result.Transfers[0].BatchNumber)
fmt.Println(result.Transfers[0].Datetime)
fmt.Println(result.Transfers[0].From)
fmt.Println(result.Transfers[0].To)
fmt.Println(result.Transfers[0].Amount)
fmt.Println(result.Transfers[0].Note)

fmt.Println(result.Transfers[1].BatchNumber)
fmt.Println(result.Transfers[1].Datetime)
fmt.Println(result.Transfers[1].From)
fmt.Println(result.Transfers[1].To)
fmt.Println(result.Transfers[1].Amount)
fmt.Println(result.Transfers[1].Note)
```

### Get transfers history
```go
ctx := context.Background()
history := &fasapay.GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
result, resp, err := client.Transfers().GetHistory(history, ctx, nil)

if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(response)

//Dump result
fmt.Println(result.History.Page.TotalItem)
fmt.Println(result.History.Page.PageCount)
fmt.Println(result.History.Page.CurrentPage)

fmt.Println(result.History.Details[0].BatchNumber)
fmt.Println(result.History.Details[0].Datetime)
fmt.Println(result.History.Details[0].From)
fmt.Println(result.History.Details[0].To)
fmt.Println(result.History.Details[0].Amount)
fmt.Println(result.History.Details[0].Note)
```

### Get transfers details
```go
ctx := context.Background()
var detail1 fasapay.GetDetailsRequestDetailParamsString = "TR0000000001"
var detail2 fasapay.GetDetailsRequestDetailParamsString = "TR0000000002"
details := []fasapay.GetDetailsDetailParamsInterface{&detail1, &detail2}
result, resp, err := client.Transfers().GetDetails(details, ctx, nil)

if err != nil {
    fmt.Printf("Wrong API request " + err.Error())
    panic(err)
}

//Dump raw response
fmt.Println(response)

//Dump result
fmt.Println(result.Details[0].BatchNumber)
fmt.Println(result.Details[0].Datetime)
fmt.Println(result.Details[0].From)
fmt.Println(result.Details[0].To)
fmt.Println(result.Details[0].Amount)
fmt.Println(result.Details[0].Note)

fmt.Println(result.Details[1].BatchNumber)
fmt.Println(result.Details[1].Datetime)
fmt.Println(result.Details[1].From)
fmt.Println(result.Details[1].To)
fmt.Println(result.Details[1].Amount)
fmt.Println(result.Details[1].Note)
```