# Exchange Rates Go

Library for convert currencies from applications written with Go. Required Go v1.14 or newer.

- [Installation](#installation)
- [Usage Examples](#usage-examples)
  - [Convert Currency](#convert-currency)
  - [Convert Currency Based on Date](#convert-currency-based-on-date)
- [Test](#test)
- [Contributing](#contributing)

---

## Installation

Install exchange-rates-go using Go Module by following command:

```bash
go get github.com/yusufthedragon/exchange-rates-go
```

Then you import it by following code:

```go
import exchangerates "github.com/yusufthedragon/exchange-rates-go"
```

## Usage Examples

### Convert Currency

```go
var currency, err = exchangerates.ConvertCurrency(&exchangerates.RequestParameter{
    From:  "USD",
    To:    "IDR",
    Value: 2,
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Converted Currency: %+v\n", currency)
```

### Convert Currency Based on Date

```go
var currencyLastYear, err = exchangerates.ConvertCurrency(&exchangerates.RequestParameter{
    Date:  time.Now().AddDate(-1, 0, 0),
    From:  "USD",
    To:    "IDR",
    Value: 2,
})

if err != nil {
    panic(err.Error())
}

fmt.Printf("Converted Currency Last Year: %+v\n", currencyLastYear)
```

## Test

You can run the test by following command:

```bash
go test -v
```

## Contributing

For any requests, bugs, or comments, please open an [issue](github.com/yusufthedragon/exchange-rates-go/issues) or [submit a pull request](github.com/yusufthedragon/exchange-rates-go/pulls).
