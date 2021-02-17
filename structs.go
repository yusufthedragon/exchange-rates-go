package exchangerates

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// RequestParameter struct contains parameter for get currency.
type RequestParameter struct {
	Date  time.Time
	From  string `validate:"required,validCurrency"`
	To    string `validate:"required,validCurrency"`
	Value float64
}

// ResponseParameter struct contains response from get currency.
type ResponseParameter struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// listCurrencies contains all of currencies in the world.
var listCurrencies = []string{
	"AUD",
	"BGN",
	"BRL",
	"CAD",
	"CHF",
	"CNY",
	"CZK",
	"DKK",
	"EUR",
	"GBP",
	"HKD",
	"HRK",
	"HUF",
	"IDR",
	"ILS",
	"INR",
	"ISK",
	"JPY",
	"KRW",
	"MXN",
	"MYR",
	"NOK",
	"NZD",
	"PHP",
	"PLN",
	"RON",
	"RUB",
	"SEK",
	"SGD",
	"THB",
	"TRY",
	"USD",
	"ZAR",
}

// ValidateCurrency function checks if the selected currency is valid.
func ValidateCurrency(fl validator.FieldLevel) bool {
	result := false

	for _, val := range listCurrencies {
		if val == fl.Field().String() {
			result = true
		}
	}

	return result
}
