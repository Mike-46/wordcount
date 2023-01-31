package main

import (
	"fmt"
	"os"
	"strings"
	// "github.com/Mike-46/wordcount"
)

func main() {
	// new comment #1

	s := os.Args[1]
	if s == "" {
		fmt.Println(0)
	} else {
		st := strings.Split(s, " ")
		fmt.Println(len(st))
	}
}
