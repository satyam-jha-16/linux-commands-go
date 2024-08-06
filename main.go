package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpMess)
	}
	log.SetFlags(0)

	flag.Parse()

	err := execCommand()
	if err != nil {
		log.Fatal(err)
	}

}
