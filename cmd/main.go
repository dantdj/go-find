package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filenameToSearch := flag.String("filename", "", "Filename to search for. (Required)")
	flag.Parse()

	if *filenameToSearch == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Searching for %s...\n", *filenameToSearch)

	foundFiles, err := Query(".", *filenameToSearch)

	if err != nil {
		fmt.Printf("Couldn't find %s\n", *filenameToSearch)
		os.Exit(0)
	}

	fmt.Println("Found the following list of matches:")
	for _, file := range foundFiles {
		filepath, err := filepath.Abs(file.Name())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(0)
		}
		fmt.Printf("%s\n", filepath)
	}
}
