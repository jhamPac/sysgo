package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

	walkFn := func(path string, info os.FileInfo, e error) error {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if printAll {
			fmt.Println(path)
			return nil
		}

		mode := fileInfo.Mode()
		if mode.IsRegular() && *fFlag {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *dFlag {
			fmt.Println(path)
			return nil
		}

		fileInfo, _ = os.Lstat(path)

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *lkFlag {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *pFlag {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *sFlag {
				fmt.Println(path)
				return nil
			}
		}

		return nil
	}

	err := filepath.Walk(path, walkFn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
