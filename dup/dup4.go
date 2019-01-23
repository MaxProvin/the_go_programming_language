// Supplied arguments are read and interpreted as files, prints filename of files with duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		return
	}
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		if countLines2(f) > 1 {
			fmt.Println(f.Name())
		}
		f.Close()
	}
}

func countLines2(f *os.File) (count int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count++
	}
	return count
	// NOTE: ignoring potential errors from input.Err()
}
