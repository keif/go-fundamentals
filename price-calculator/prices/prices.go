package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
}

// constructor
// return pointer so we only share it once in memory
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (t TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range t.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + t.TaxRate)
	}

	fmt.Println(result)
}
