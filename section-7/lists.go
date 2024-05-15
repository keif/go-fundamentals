package main

import "fmt"

func main() {
	// indices 1,2,3 are nil
	var productNames [4]string = [4]string{"A Book"}
	// no nil values
	prices := [4]float64{10.99, 9.99, 11.99, 12.99}

	// indices 1,3 are still nil
	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println(prices)

	fmt.Println(prices[0])

	// a slice of prices array
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
