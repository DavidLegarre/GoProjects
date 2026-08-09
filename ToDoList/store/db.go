package store

import (
	"encoding/json"
	"fmt"
	"todolist/types"
)

func ReadDB() (*ItemStore, error) {
	data, err := readJsonFile()
	if err != nil {
		return nil, err
	}

	var todoList types.TodoList
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		return nil, err
	}

	return &ItemStore{items: todoList}, nil
}

func WriteDB(itemStore *ItemStore) error {
	b, err := json.Marshal(itemStore.GetEntries())
	if err != nil {
		fmt.Println("Error marshalling list:", err)
		return err
	}
	return writeJsonFile(b)
}
