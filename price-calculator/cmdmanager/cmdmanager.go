package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

type CMDManager struct {
}

func (cm *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices, confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cm *CMDManager) WriteResult(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	fmt.Println(encoder.Encode(data))
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
