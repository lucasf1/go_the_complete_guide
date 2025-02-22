package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-pratice/note"
	"example.com/structs-pratice/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello, world!")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Todo saved successfully!")

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// default:
	// 	fmt.Println("Unknown type:", value)
	// }

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
	}

	fmt.Println(value, ok)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return err
	}

	fmt.Println("Note saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

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
