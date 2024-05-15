# Go Tastytrade Open API Wrapper

![Build](https://github.com/optionsvamp/tastytrade/actions/workflows/build.yaml/badge.svg)

Go API wrapper for the Tastytrade Open API

## Table of Contents
1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Testing](#testing)
5. [Contributing](#contributing)
6. [License](#license)

## Introduction

This project provides a Go wrapper for the Tastytrade Open API. It allows developers to interact with Tastytrade's financial data and services in a more Go-idiomatic way, abstracting away the details of direct HTTP requests and responses.

## Installation

To install this project, you can use `go get`:

```bash
go get github.com/optionsvamp/tastytrade
```

Then, import it in your Go code:

```
import "github.com/optionsvamp/tastytrade"
```

## Usage

Here's a basic example of how to use this wrapper to get account balances:

```
api := tastytrade.NewTastytradeAPI("your-api-key")
balances, err := api.GetAccountBalances("your-account-number")
if err != nil {
    log.Fatal(err)
}
fmt.Println(balances)
```

## Testing

To run the tests for this project, you can use go test:

```bash
go test ./...
```

## Contributing

Contributions to this project are welcome! Please submit a pull request or open an issue on GitHub.

## License

This project is released into the public domain under the Unlicense. For more information, please refer to the LICENSE file or visit https://unlicense.org.