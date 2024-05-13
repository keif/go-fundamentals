package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title   string
	content string
	created time.Time
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

func (note Note) Display() {
	fmt.Printf("Your note titlted %v has the following content:\n\n%v\n\n", note)
}
