package store

import (
	"fmt"
	"os"
	"path/filepath"
)

func readJsonFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func writeJsonFile(path string, data []byte) error {
	tmpFile, err := os.CreateTemp(filepath.Dir(path), "tmp.*.json")
	dataFile := filepath.Base(path)
	if err != nil {
		err = fmt.Errorf("Error occurred while writing to JsonFile %s: %w", dataFile, err)
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
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("rename to %s: %w", dataFile, err)
	}
	return nil
}
