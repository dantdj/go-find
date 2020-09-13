package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Query(baseDirectory string, filename string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		log.Printf("Failed to read directory, encountered error: %s", err)
		return nil, err
	}
	var matchingFiles []os.FileInfo
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matchingFiles = append(matchingFiles, file)
		}
	}

	return matchingFiles, nil
}
