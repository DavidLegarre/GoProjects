package store

import (
	"fmt"
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

func writeJsonFile(data []byte) error {
	err := os.WriteFile(jsonFilePath, data, 0644)
	if err != nil {
		err = fmt.Errorf("Error occured while writing to JsonFile %s: %w", jsonFilePath, err)
	}
	return err
}
