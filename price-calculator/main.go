package main

import (
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for taxIdx, taxRate := range taxRates {
		doneChans[taxIdx] = make(chan bool)
		errorChans[taxIdx] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// go routines do not return values
		// and this fails to generate files
		go priceJob.Process(doneChans[taxIdx], errorChans[taxIdx])
		//err := priceJob.Process()
		//if err != nil {
		//	fmt.Println("Could not process price job")
		//	fmt.Println(err)
		//}

		//priceJobcdm := prices.NewTaxIncludedPriceJob(cdm, taxRate)
		//priceJobcdm.Process()
	}

	for idx, val := range taxRates {
		// we wait for only one channel to emit a value and not care about the other case
		select {
		case err := <-errorChans[idx]:
			if err != nil {
				fmt.Printf("Error processing tax rate %v: %v\n", val, err)
			}
		case <-doneChans[idx]:
			fmt.Printf("Processed tax rate %v\n", val)
		}
	}

	//for _, errorChan := range errorChans {
	//	// will crash, because it's waiting forever
	//	<-errorChan
	//}
	//for _, doneChan := range doneChans {
	//	<-doneChan
	//}

}
