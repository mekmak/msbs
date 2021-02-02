package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mekmak/msbs/pkg/solver"
)

func main() {

	if len := len(os.Args); len != 4 {
		log.Fatal("Need to pass in three args -- file path, a string representing valid letters with the first letter being required, and the min word length")
	}

	validLetters := []rune(os.Args[2])

	minSize, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	words, err := solver.LoadWords(os.Args[1], validLetters[0], validLetters, minSize)
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range words {
		fmt.Println(w)
	}
}
