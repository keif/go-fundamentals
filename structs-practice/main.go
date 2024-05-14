package main

import (
	"bufio"
	"example.com/note/note"
	"example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

// common convention:
// when your interface has one method, the name is the method + "er" - ergo, saver
// you don't have to associate the interface explicity
// it checks that what's passed in has a Save method with an error return type
type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		fmt.Println("Saving the todo failed:", err)
		return
	}

	fmt.Println("Saved todo:", title)

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return
	}

	fmt.Println("Saved note:", title)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed:", err)
		return err
	}
	fmt.Println("Saved data:", data)
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
