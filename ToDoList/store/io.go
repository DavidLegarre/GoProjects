package store

import "os"

const jsonFilePath = "./ToDoList/todolist.json"

func readJsonFile() []byte {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil
	}
	return data
}

func writeJsonFile(data []byte) {
	os.WriteFile(jsonFilePath, data, 0644)
}
