package main

import (
	"flag"
	"fmt"
	"os"
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
		if mode.IsRegular() == false {
			fmt.Println("Sorry, only regular files are supported as the source.")
			os.Exit(1)
		} else {
			fmt.Println("Error reading:", source)
			os.Exit(1)
		}
	}

}
