package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("No user input")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
