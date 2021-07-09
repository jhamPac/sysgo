package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	sFlag := flag.Bool("s", false, "Sockets")
	pFlag := flag.Bool("p", false, "Pipes")
	lkFlag := flag.Bool("lk", false, "Symbolic Links")
	dFlag := flag.Bool("d", false, "Directories")
	fFlag := flag.Bool("f", false, "Files")

	flag.Parse()
	flags := flag.Args()

	printAll := false
	if *sFlag && *pFlag && *lkFlag && *dFlag && *fFlag {
		printAll = true
	}

	// all flags begin as flase
	if !(*sFlag || *pFlag || *lkFlag || *dFlag || *fFlag) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments.")
		os.Exit(1)
	}

	path := flags[0]

}
