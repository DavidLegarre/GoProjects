package main

import "os"

const jsonFilePath = "./ToDoList/todolist.json"

func readJsonFile() {
}

func writeJsonFile(data []byte) {
	os.WriteFile(jsonFilePath, data, 0644)
}
