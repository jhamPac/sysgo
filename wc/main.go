package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func count(filename string) (numberOfLines int, numberOfWords int, numberOfCharacters int) {
	var err error

	numberOfLines = 0
	numberOfCharacters = 0
	numberOfWords = 0

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		numberOfLines++
		reg := regexp.MustCompile(`[^\\s]+`)
		for range reg.FindAllString(line, -1) {
			numberOfWords++
		}
		numberOfCharacters += len(line)
	}
	return numberOfLines, numberOfWords, numberOfCharacters
}
