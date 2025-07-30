package main

import (
	"errors"
	"fmt"
)

const USDToEuro float64 = 0.85
const USDToRub float64 = 79.40

func main() {

	var err error
	var sourceValuta string
	var targetValuta string
	var amount float64

	for {
		fmt.Print("Введите исходную валюту (USD, RUB, EUR): ")
		sourceValuta, err = getValuta()

		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	for {
		fmt.Print("Введите сумму: ")
		amount, err = getAmount()

		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	for {
		fmt.Print("Введите целевую валюту (USD, RUB, EUR): ")
		targetValuta, err = getValuta()

		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	convert(amount, sourceValuta, targetValuta)

	// eurosToRubles := amount / USDToEuro * USDToRub
}

func getAmount() (float64, error) {
	var amount float64

	fmt.Scan(&amount)

	if amount <= 0 {
		return 0, errors.New("Ошибка ввода суммы, некорректное значение")
	}

	return amount, nil
}

func getValuta() (string, error) {
	var valuta string

	fmt.Scan(&valuta)

	if valuta == "RUB" || valuta == "USD" || valuta == "EUR" {
		return valuta, nil
	} else {
		return "", errors.New("Ошибка ввода валюты, некорректное значение")
	}
}

func convert(amount float64, sourceValuta string, targetValuta string) {

	switch targetValuta {
	case "USD":
		if sourceValuta == "RUB" {
			fmt.Printf("Ваше перевод рублей в доллары: \n%.2f", amount/USDToRub)
		} else {
			fmt.Printf("Ваше перевод евро в доллары: \n%.2f", amount/USDToEuro)
		}
	case "EUR":
		if sourceValuta == "RUB" {
			fmt.Printf("Ваше перевод рублей в евро: \n%.2f", amount/USDToRub*USDToEuro)
		} else {
			fmt.Printf("Ваше перевод долларов в евро: \n%.2f", amount*USDToEuro)
		}
	case "RUB":
		if sourceValuta == "USD" {
			fmt.Printf("Ваше перевод долларов в рубли: \n%.2f", amount*USDToRub)
		} else {
			fmt.Printf("Ваше перевод евро в рубли: \n%.2f", amount*USDToEuro*USDToRub)
		}
	}
}
