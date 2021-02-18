package exchangerates

import "github.com/go-playground/validator/v10"

// The validator instance.
var validation = validator.New()

// ConvertCurrency function gets the currency value from API.
func ConvertCurrency(params *RequestParameter) (float64, error) {
	validation.RegisterValidation("validCurrency", ValidateCurrency)
	var errValidation = validation.Struct(params)

	if errValidation != nil {
		return 0, errValidation
	}

	endpoint := "https://api.vatcomply.com/rates?base=" + params.From

	if !params.Date.IsZero() {
		date := params.Date.Format("2006-01-02")

		endpoint += "&date=" + date
	}

	currency, err := SendRequest(endpoint)

	if err != nil {
		return 0, err
	}

	targetRate := currency.Rates[params.To]

	if params.Value == 0 {
		params.Value = 1
	}

	return float64(params.Value) * targetRate, nil
}
