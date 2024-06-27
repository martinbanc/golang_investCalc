package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturn = 5.5
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Years:")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)
}
