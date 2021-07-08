package which

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Execute() {
	optA := flag.Bool("a", false, "a")
	optS := flag.Bool("s", false, "s")

	flag.Parse()

	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Please provide arguments")
		os.Exit(1)
	}

	file := flags[0]
	found := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, dir := range pathSlice {
		fullPath := dir + "/" + file
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					found = true
					if *optS == true {
						os.Exit(0)
					}

					if *optA == true {
						fmt.Println(fullPath)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}

	if found == false {
		os.Exit(1)
	}

}
