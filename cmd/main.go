package main

import (
	"flag"
	"fmt"
	"log"
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

	log.Printf("Looks like you're searching for %s", *filenameToSearch)

	found, err := Query(".", *filenameToSearch)

	if err != nil {
		fmt.Printf("Couldn't find %s", *filenameToSearch)
		os.Exit(0)
	}
	filepath, _ := filepath.Abs(found[0].Name())

	fmt.Printf("Found %s at %s", found[0].Name(), filepath)
}
