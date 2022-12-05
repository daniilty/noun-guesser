package main

import (
	"bufio"
	"fmt"
	"log"
	"noun-guesser/internal/tree"
	"os"
)

func run() error {
	f, err := os.Open("source.txt")
	if err != nil {
		return err
	}

	t := tree.NewWord()

	s := bufio.NewScanner(f)
	for s.Scan() {
		t.Insert(s.Text())
	}

	letters := ""
	lettersGuessed := ""
	ignored := []rune{}
	notInOrder := []rune{}

	for _, r := range letters {
		ignored = append(ignored, r)
	}

	for _, r := range lettersGuessed {
		notInOrder = append(notInOrder, r)
	}

	fmt.Println(t.Find("*****", ignored, notInOrder))

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
