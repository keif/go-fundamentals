package main

import "fmt"

func main() {
	// indices 1,2,3 are nil
	var productNames [4]string = [4]string{"A Book"}
	// no nil values
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	// indices 1,3 are still nil
	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println(prices)

	fmt.Println(prices[0])

	// a slice of prices array
	//featuredPrices := prices[1:3]
	// starts with 0 then to 3
	//featuredPrices := prices[:3]
	// starts at 1 and continues to the end
	featuredPrices := prices[1:]
	// negatives are not valid, and not out of bounds
	// equivalent to the above, since 4 is the length of the array
	//featuredPrices := prices[1:4]
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
}
