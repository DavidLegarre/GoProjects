package store

import (
	"fmt"
	"os"
)

const rootPath = "./ToDoList/"
const jsonFilePath = rootPath + "todolist.json"

func readJsonFile() ([]byte, error) {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return data, err
	}
	return data, nil
}

func writeJsonFile(data []byte) error {
	tmpFile, err := os.CreateTemp(rootPath, "tmp.*.json")
	if err != nil {
		err = fmt.Errorf("Error occurred while writing to JsonFile %s: %w", jsonFilePath, err)
		return err
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)
	defer tmpFile.Close()

	if err := tmpFile.Chmod(0644); err != nil {
		return fmt.Errorf("set temp file mode: %w", err)
	}
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := tmpFile.Sync(); err != nil {
		return fmt.Errorf("sync temp file: %w", err)
	}
	if err := os.Rename(tmpPath, jsonFilePath); err != nil {
		return fmt.Errorf("rename to %s: %w", jsonFilePath, err)
	}
	return nil
}
