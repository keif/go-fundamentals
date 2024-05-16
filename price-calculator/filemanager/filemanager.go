package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not open file: %s", err))
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New(fmt.Sprintf("There was an error processing the file: %s", err))
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not create file: %s", err))
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New(fmt.Sprintf("There was an error processing the file: %s", err))
	}

	file.Close()
	return nil
}
