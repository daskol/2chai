package main

import (
	"flag"
	"log"
)

func printList(lst *Threads) {
	log.Println(lst)

	for i, thread := range lst.Threads {
		log.Printf("[%03d] %s\n", i, thread)
	}
}

func main() {
	board := flag.String("board", "abu", "Board to mine.")

	flag.Parse()

	if lst, err := ListThreads(*board, 0); err != nil {
		log.Fatal(err)
	} else {
		printList(lst)
	}

	if lst, err := ListThreadCatalog(*board); err != nil {
		log.Fatal(err)
	} else {
		printList(lst)
	}
}
