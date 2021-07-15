package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need one file name")
		os.Exit(-1)
	}

	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("file %s already exists", filename)
		os.Exit(1)
	}

	output, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer output.Close()
}
