package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetMainPath() string {
	path, _ := os.Getwd()
	for !isMainDir(path) {
		path = filepath.Dir(path)
	}
	return path
}

func isMainDir(path string) bool {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.Name() == "main.go" {
			return true
		}
	}
	return false
}
