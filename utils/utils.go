package utils

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func ValueExists(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func ProcessPath(path string) string {
	var finalPath string
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		folder := usr.HomeDir
		finalPath = folder + path[1:]
	} else {
		finalPath = path
	}
	return finalPath
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFileWithPath(path string) error {
	// Extract the directory from the path
	dir := filepath.Dir(path)

	// Create the directory including any necessary parents
	err := os.MkdirAll(dir, 0755) // 0755 is the permission (rwxr-xr-x)
	if err != nil {
		return fmt.Errorf("failed to create directories: %s", err)
	}

	if exists, _ := fileExists(path); !exists {
		// Create the file
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create file: %s", err)
		}

		file.Close()
	}
	return err
}
