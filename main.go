package main

import (
	"fmt"
)

type currencyAlias = map[string]float64

func main() {
	var amount float64
	var sourceCurrency string
	var targetCurrency string
	var result float64
	var err error

	currency := map[string]float64{
		"USD": 1,
		"EUR": 0.8563,
		"RUB": 80.19,
	}
	currencyPointer := &currency

	for {
		sourceCurrency, err = getSourceCurrency(currencyPointer)
		if err != nil {
			fmt.Printf("ошибка конвертации: %v\n", err)
			continue
		}
		break
	}

	for {
		targetCurrency, err = getTargetCurrency(currencyPointer)
		if err != nil {
			fmt.Printf("ошибка конвертации: %v\n", err)
			continue
		}
		break
	}

	for {
		amount, err = getAmount()
		if err != nil {
			fmt.Printf("ошибка конвертации: %v\n", err)
			continue
		}
		break
	}

	result = convert(amount, sourceCurrency, targetCurrency, currencyPointer)
	fmt.Printf("Результат конвертации %s в %s = %v\n", sourceCurrency, targetCurrency, result)

}

func getSourceCurrency(currency *currencyAlias) (string, error) {
	if currency == nil {
		return "", fmt.Errorf("передан nil вместо map")
	}

	fmt.Print("Введите исходную валюту (USD, RUB, EUR): ")
	var sourceCurrency string
	_, err := fmt.Scan(&sourceCurrency)
	if err != nil {
		return "", fmt.Errorf("ошибка ввода: %v", err)
	}

	_, ok := (*currency)[sourceCurrency]
	if !ok {
		return "", fmt.Errorf("валюта %s не поддерживается", sourceCurrency)
	}

	return sourceCurrency, nil
}

func getTargetCurrency(currency *currencyAlias) (string, error) {
	if currency == nil {
		return "", fmt.Errorf("передан nil вместо map")
	}

	fmt.Print("Введите целевую валюту (USD, RUB, EUR): ")
	var targetCurrency string
	_, err := fmt.Scan(&targetCurrency)
	if err != nil {
		return "", fmt.Errorf("ошибка ввода: %v", err)
	}

	_, ok := (*currency)[targetCurrency]
	if !ok {
		return "", fmt.Errorf("валюта %s не поддерживается", targetCurrency)
	}

	return targetCurrency, nil
}

func getAmount() (float64, error) {
	fmt.Print("Введите сумму: ")
	var amount float64
	_, err := fmt.Scan(&amount)
	if err != nil || amount < 0 {
		return 0, fmt.Errorf("введите положительное число")
	}

	return amount, nil
}

func convert(
	amount float64,
	source string,
	target string,
	currency *currencyAlias,
) float64 {
	var result float64

	sourceRate := (*currency)[source]
	targetRate := (*currency)[target]

	result = (amount / sourceRate) * targetRate
	return result
}
