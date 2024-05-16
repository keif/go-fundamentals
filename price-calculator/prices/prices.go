package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` // exclude IOManager
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// constructor
// return pointer so we only share it once in memory
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

// error return type removed as it's a go routine
func (t *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := t.LoadData()
	if err != nil {
		//return err
	}

	result := make(map[string]string)

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	t.TaxIncludedPrices = result
	t.IOManager.WriteResult(t)
	doneChan <- true
}

// must be a pointer so we update the values, not copy them
func (t *TaxIncludedPriceJob) LoadData() error {
	lines, err := t.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	t.InputPrices = prices

	return nil
}
