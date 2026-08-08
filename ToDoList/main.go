package main

import (
	"fmt"
	"todolist/types"
)

func main() {
	fmt.Println("Hello, World!")
	todo := types.TodoEntry{Name: "Buy groceries"}
	fmt.Println(todo)
	Run()
}
