package store

import (
	"os"
)

const jsonFilePath = "./ToDoList/todolist.json"

func readJsonFile() ([]byte, error) {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return data, err
	}
	return data, nil
}

func writeJsonFile(data []byte) {
	os.WriteFile(jsonFilePath, data, 0644)
}
