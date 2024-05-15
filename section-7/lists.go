package main

import "fmt"

func main() {
	prices := []float64{10.99, 18.99} // dynamic array! since we didn't declare a length
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	//prices[2] = 11.99 // out of range of the slice, do not add this way

	// add this way
	//updatedPrices := append(prices, 5.99) // adds a new value to the slice, and CREATES a new slice
	//fmt.Println(updatedPrices)
	//fmt.Println(prices)

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{5.99, 2.99, 1.99}
	// ... spreads
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}

//func main() {
//	// indices 1,2,3 are nil
//	var productNames [4]string = [4]string{"A Book"}
//	// no nil values
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//
//	// indices 1,3 are still nil
//	productNames[2] = "A Carpet"
//	fmt.Println(productNames)
//	fmt.Println(prices)
//
//	fmt.Println(prices[0])
//
//	// a slice of prices array
//	//featuredPrices := prices[1:3]
//	// starts with 0 then to 3
//	//featuredPrices := prices[:3]
//	// starts at 1 and continues to the end
//	featuredPrices := prices[1:]
//	// negatives are not valid, and not out of bounds
//	// equivalent to the above, since 4 is the length of the array
//	//featuredPrices := prices[1:4]
//	featuredPrices[0] = 199.99 // as a slice is a window into the array
//	fmt.Println(featuredPrices)
//
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(highlightedPrices)
//	fmt.Println(prices)
//	fmt.Println(len(featuredPrices), cap(featuredPrices))
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1, 3 because the original slice has more capacity.
//	// you can always select more at the end, not at the start - you can't go back, so it only counts towards the original of the array
//
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(highlightedPrices)
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 3, 3 because we sliced up to the length - go knows there was more capacity to the right
//
//}
