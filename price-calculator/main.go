package main

import (
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for taxIdx, taxRate := range taxRates {
		doneChans[taxIdx] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// go routines do not return values
		// and this fails to generate files
		go priceJob.Process(doneChans[taxIdx])
		//err := priceJob.Process()
		//if err != nil {
		//	fmt.Println("Could not process price job")
		//	fmt.Println(err)
		//}

		//priceJobcdm := prices.NewTaxIncludedPriceJob(cdm, taxRate)
		//priceJobcdm.Process()
	}
	for _, doneChan := range doneChans {
		<-doneChan
	}

}
