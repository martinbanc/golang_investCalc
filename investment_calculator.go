package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturn = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Println("Future value is", futureValue)

}
