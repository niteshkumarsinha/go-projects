package main

import (
	"fmt"
	"os"
)

func main(){
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCommandFlags()
	if err := cmdFlags.Execute(&todos); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// todos.Add("Buy Milk")
	// todos.Add("Buy Bread")
	// todos.Toggle(0)
	// todos.List()
	// todos.Toggle(1)
	// todos.List()
	// todos.Toggle(0)
	// todos.Toggle(1)
	// todos.List()
	todos.List()
	storage.Save(todos)
}
