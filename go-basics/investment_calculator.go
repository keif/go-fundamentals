package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value (adjusted for inflation): %.2f\n", futureRealValue)
	// prints with a new line
	//fmt.Println("Future Value:", futureValue)
	// prints with two decimals and new lines
	//	fmt.Printf(`Future Value: %.2f
	//
	//Future Value (adjusted for inflation): %.2f`, futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// go will return them since we declared them above - not very explicit, though.
	//return
}
