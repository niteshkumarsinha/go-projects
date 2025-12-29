package main


func main(){
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := CmdFlags{}
	cmdFlags.Execute(&todos)
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