package main


func main(){
	todos := Todos{}
	todos.Add("Buy Milk")
	todos.Add("Buy Bread")
	todos.Toggle(0)
	todos.List()
	todos.Toggle(1)
	todos.List()
	todos.Toggle(0)
	todos.Toggle(1)
	todos.List()
}