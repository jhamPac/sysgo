package main

import (
	"bufio"
	"flag"
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

func main() {
	cFlag := flag.Bool("c", false, "Characters")
	wFlag := flag.Bool("w", false, "Words")
	lFlag := flag.Bool("l", false, "Lines")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Print("usage: wc <file> [<file2> [... <fileN]\n")
		os.Exit(1)
	}

	totalLines := 0
	totalWords := 0
	totalChars := 0
	printAll := false

	for _, filename := range flag.Args() {
		nLines, nWords, nChars := count(filename)

		totalLines = totalLines + nLines
		totalWords = totalWords + nWords
		totalChars = totalChars + nChars

		if (*cFlag && *wFlag && *lFlag) || !(*cFlag || *wFlag || *lFlag) {
			fmt.Printf("%d", nLines)
			fmt.Printf("\t%d", nWords)
			fmt.Printf("\t%d", nChars)
			printAll = true
			continue
		}

		if *lFlag {
			fmt.Printf("%d", nLines)
		}

		if *wFlag {
			fmt.Printf("%d", nWords)
		}

		if *cFlag {
			fmt.Printf("%d", nChars)
		}

		fmt.Printf("\t%s\n", filename)
	}
}
