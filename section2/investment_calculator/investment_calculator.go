package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmont float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmont)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmont, expectedReturnRate, years)

	// var futureValue = investmentAmont * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Value after Inflation: %.2f\n", futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmont, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmont * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow((1+inflationRate/100), years)
	// return fv, rfv

	return
}
