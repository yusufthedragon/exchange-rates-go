package exchangerates

import (
	"fmt"
	"testing"
	"time"
)

func TestUSDToIDR(t *testing.T) {
	currency, err := ConvertCurrency(&RequestParameter{
		From:  "USD",
		To:    "IDR",
		Value: 2,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Result: ", currency)
}

func TestUSDToIDRLastYear(t *testing.T) {
	currency, err := ConvertCurrency(&RequestParameter{
		Date:  time.Now().AddDate(-1, 0, 0),
		From:  "USD",
		To:    "IDR",
		Value: 2,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Result: ", currency)
}

func TestIDRToUSD(t *testing.T) {
	currency, err := ConvertCurrency(&RequestParameter{
		From:  "IDR",
		To:    "USD",
		Value: 160000,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Result: ", currency)
}

func TestIDRToUSDLastYear(t *testing.T) {
	currency, err := ConvertCurrency(&RequestParameter{
		Date:  time.Now().AddDate(-1, 0, 0),
		From:  "IDR",
		To:    "USD",
		Value: 160000,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("Result: ", currency)
}
