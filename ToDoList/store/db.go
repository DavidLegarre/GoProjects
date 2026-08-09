package store

import (
	"encoding/json"
	"fmt"
	"todolist/types"
)

func ReadDB() *ItemStore {
	data := readJsonFile()
	if data == nil {
		fmt.Println("Error reading JSON file.")
		return nil
	}

	var todoList types.TodoList
	err := json.Unmarshal(data, &todoList)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return &ItemStore{items: todoList}
}

func WriteDB() {
	b, err := json.Marshal(Items)
	if err != nil {
		fmt.Println("Error marshalling list:", err)
		return
	}
	writeJsonFile(b)
}

func DeleteDB() {
	Items = types.TodoList{Entries: []types.TodoEntry{}}
	WriteDB()
}
