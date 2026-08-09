package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"todolist/types"
)

func ReadDB() *ItemStore {
	data, err := readJsonFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) == false {
			fmt.Printf("Unexpected error occured %s", err)
		}
		return nil
	}

	var todoList types.TodoList
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return &ItemStore{items: todoList}
}

func WriteDB(itemStore *ItemStore) {
	b, err := json.Marshal(itemStore.GetEntries())
	if err != nil {
		fmt.Println("Error marshalling list:", err)
		return
	}
	writeJsonFile(b)
}
