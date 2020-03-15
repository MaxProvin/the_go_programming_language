// Echo1 prints its command-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // os.Args first element is the name of the command itself
		sep = " "
	}
	fmt.Println(s)
}
