package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

func Execute() {
	args := os.Args

	pwd, err := os.Getwd()
	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("Error:", err)
	}

	if len(args) == 1 {
		return
	}

	if args[1] != "-P" {
		return
	}

	fileinfo, _ := os.Lstat(pwd)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}
