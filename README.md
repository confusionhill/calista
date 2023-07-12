# Calista

<p align="center"><img src="doc/heimdall-logo.png" width="360"></p>
<p align="center">
  <a href="https://travis-ci.com/gojek/heimdall"><img src="https://travis-ci.com/gojek/heimdall.svg?branch=master" alt="Build Status"></img></a>
  <a href="https://goreportcard.com/report/github.com/confusionhill/calista"><img src="https://goreportcard.com/badge/github.com/gojek/heimdall"></img></a>
  <a href="https://golangci.com"><img src="https://golangci.com/badges/github.com/gojek/heimdall.svg"></img></a>
  <a href="https://coveralls.io/github.com/confusionhill/calista?branch=master"><img src="https://coveralls.io/repos/github/gojek/heimdall/badge.svg?branch=master"></img></a>
</p>

* [Description](#description)
* [Installation](#installation)
* [Usage](#usage)
    + [Making a simple `GET` request](#making-a-simple-get-request)
    + [Creating a hystrix-like circuit breaker](#creating-a-hystrix-like-circuit-breaker)
    + [Creating a hystrix-like circuit breaker with fallbacks](#creating-a-hystrix-like-circuit-breaker-with-fallbacks)
    + [Creating an HTTP client with a retry mechanism](#creating-an-http-client-with-a-retry-mechanism)
    + [Custom retry mechanisms](#custom-retry-mechanisms)
    + [Custom HTTP clients](#custom-http-clients)
* [Plugins](#plugins)
* [Documentation](#documentation)
* [FAQ](#faq)
* [License](#license)

## Description

Calista is an easy to use HTTP client that helps you develope your application faster. With Calista, you can:
- Bind JSON object easily

All HTTP methods are exposed as a fluent interface.

## Installation
```
go get -u github.com/gojek/heimdall/v7
```

## Usage

### Importing the package

This package can be used by adding the following import statement to your `.go` files.

```go
import "github.com/gojek/heimdall/v7/httpclient" 
```

### Making a simple `GET` request
The below example will print the contents from `https://fakestoreapi.com/products/1`:

```go
// creates http client
httpCli, err := httpclient.NewClient()
if err != nil {
    log.Fatal(err)
}

// response struct
var resp struct {
    ID          int64   `json:"id"`
    Title       string  `json:"title"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
    Category    string  `json:"category"`
    Image       string  `json:"image"`
    Rating      struct {
        Rate  float64 `json:"rate"`
        Count int64   `json:"count"`
    } `json:"rating"`
}

// create a get request
err = httpCli.Get("https://fakestoreapi.com/products/1", nil, &resp)
if err != nil {
    log.Fatal(err)
}
fmt.Println(resp)
```
## License

```
lorem ipsum
```