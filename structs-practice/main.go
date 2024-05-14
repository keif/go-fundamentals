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

type displayer interface {
	Display()
}

// no longer -er
//type outputtable interface {
//	Save() error
//	Display()
//}

type outputtable interface {
	saver
	displayer // this could be just the function method if we wanted
	//Display()
	//DoSomething(int) string
}

func main() {
	printSomething(1)
	printSomething("string")
	printSomething(5.5)

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

	err = outputData(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	printSomething(todo)
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// this will accept any type of value - interface{}
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	//if !ok {
	//	fmt.Println("Error: value is not int")
	//}
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}
	//switch value := value.(type) {
	//case bool:
	//	fmt.Println("Bool: ", value)
	//case float64:
	//	fmt.Println("Float:", value)
	//case int:
	//	fmt.Println("Int: ", value)
	//case string:
	//	fmt.Println("String: ", value)
	//}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
