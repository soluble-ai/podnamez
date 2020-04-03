package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("no prefix argument provided")
	}
	f, err := os.Open("trails.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	g, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()
	scanner := bufio.NewScanner(g)
	for scanner.Scan() {
		fmt.Println(os.Args[1] + "-" + scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
