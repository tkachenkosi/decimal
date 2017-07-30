// изучение библиотекии decimal
package main

import (
	"fmt"
	// "strings"
	// "unicode/utf8"
	// "encoding/json"
	// "html/template"
	// "log"
	// "net/http"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("Проверка работы с типом decimal")

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromFloat(3)
	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)
	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))
	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	println("****")

	x1 := decimal.NewFromFloat(31.4)
	x2 := x1.Div(decimal.NewFromFloat(3))
	x3 := x1.DivRound(decimal.NewFromFloat(3), 4)
	println("x2", x2.String())
	println("x2", x2.StringFixed(3))
	println("x3", x3.String())

	y1 := decimal.New(3505, -2) // создал занчени 35.05
	println("y1", y1.String())

	println("проверка округления")
	b0 := decimal.New(0, 0)     // создал занчени 0.0
	b1 := decimal.New(3515, -2) // создал занчени 35.15
	b2 := decimal.New(303, -2)  // создал занчени 3.05
	b3 := b1.Div(b2)
	b4 := b1.DivRound(b2, 4)

	println("b0", b0.String(), "b1", b1.String(), "b2", b2.String())
	println("b3", b3.String(), "b4", b4.String())

	b5 := b1.DivRound(decimal.New(303, -2), 3)
	b6 := b1.DivRound(decimal.New(30311, -4), 3)

	println("влияние точночти")
	println("b5", b5.String(), "b6", b6.String())

	b3 = b1.Div(b4)
	b4 = b1.DivRound(b4, 4).Truncate(2)
	// b4 = b1.DivRound(b4, 4).Truncate(2)

	println("b3", b3.String(), "b4", b4.String())

}
