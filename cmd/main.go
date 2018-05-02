package main

import (
	"flag"
	"fmt"
	"gocode"
	"os"
	"log"
)

func main() {

	path := flag.String("f", "missing origin file or directory", "use -f file|directory")
	remotePath := flag.String("t", "missing target file or directory", "use -t file|directory")

	flag.Parse()
	if gocode.ValidatePath(*path) {
		repo := gocode.New(*path, *remotePath)
		log.Println("Sync start...")
		repo.Sync()
		gocode.WatchAnyEvent(repo)
	} else {
		fmt.Println("ERROR: Not a git repository.")
		os.Exit(1)
	}
}
