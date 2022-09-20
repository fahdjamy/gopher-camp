package utils

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

func ReadJsonFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func AbsPathToProject(filePathName string) string {
	absPath, _ := filepath.Abs(filePathName)
	return path.Join(absPath)
}
