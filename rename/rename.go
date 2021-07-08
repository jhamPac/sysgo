package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	oFlag := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()

	flags := flag.Args()
	if len(flags) < 2 {
		fmt.Println("Please provide two arguments.")
		os.Exit(1)
	}

	source := flags[0]
	destination := flags[1]

	fileInfo, err := os.Stat(source)
	if err == nil {
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			fmt.Println("Sorry, only regular files are supported as the source.")
			os.Exit(1)
		} else {
			fmt.Println("Error reading:", source)
			os.Exit(1)
		}
	}

	newDestination := destination
	destInfo, err := os.Stat(destination)
	if err == nil {
		mode := destInfo.Mode()
		if mode.IsDir() {
			name := filepath.Base(source)
			newDestination = destination + "/" + name
		}
	}

	destination = newDestination
	_, err = os.Stat(destination)
	if err == nil {
		if !*oFlag {
			fmt.Println("Destination file already exists.")
			os.Exit(1)
		}
	}

	err = os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
