package store

import (
	"encoding/json"
	"todolist/types"
)

func ReadDB(path string) (*ItemStore, error) {
	data, err := readJsonFile(path)
	if err != nil {
		return nil, err
	}

	var todoList types.TodoList
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		return nil, err
	}

	return &ItemStore{items: todoList, path: path}, nil
}

func WriteDB(itemStore *ItemStore) error {
	b, err := json.Marshal(itemStore.GetEntries())
	if err != nil {
		return err
	}
	return writeJsonFile(itemStore.path, b)
}
