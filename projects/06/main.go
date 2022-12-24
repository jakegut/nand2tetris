package main

import (
	"fmt"
	"hack-assembler/assembler"
	"hack-assembler/parser"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("not enough arguments")
	}
	file := os.Args[1]

	f, err := os.Stat(file)
	if err != nil {
		log.Fatalf("getting file: %+v", err)
	}

	if f.IsDir() {
		log.Fatalf("path is dir")
	}

	fd, err := os.Open(file)
	if err != nil {
		log.Fatalf("opening file: %+v", err)
	}

	p := parser.New(fd)
	a := assembler.New(p)

	for _, s := range a.Assemble() {
		fmt.Println(s)
	}
}
