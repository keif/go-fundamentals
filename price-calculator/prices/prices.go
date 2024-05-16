package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"` // exclude IOManager
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

// constructor
// return pointer so we only share it once in memory
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (t *TaxIncludedPriceJob) Process() {
	t.LoadData()

	result := make(map[string]string)

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	t.TaxIncludedPrices = result
	t.IOManager.WriteResult(t)
}

// must be a pointer so we update the values, not copy them
func (t *TaxIncludedPriceJob) LoadData() {
	lines, err := t.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("There was an error converting the prices to float.")
		fmt.Println(err)
		return
	}

	t.InputPrices = prices
}
