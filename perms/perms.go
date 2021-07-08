package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument.")
		os.Exit(1)
	}

	file := args[1]
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	mode := info.Mode()
	fmt.Print(file, ": ", mode, "\n")
}
