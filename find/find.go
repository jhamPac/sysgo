package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, e := os.Stat(path)
	if e != nil {
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsDir() || mode.IsRegular() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("More arguments needed.")
		os.Exit(1)
	}

	path := flags[0]

	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
