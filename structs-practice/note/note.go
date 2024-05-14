package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title   string
	content string
	created time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titlted \"%v\" has the following content:\n\n%v\n\n", note.title, note.content)
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("No user input")
	}

	return Note{
		title:   title,
		content: content,
		created: time.Now(),
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
