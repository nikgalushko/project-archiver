package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/valenok-husky/pa/internal/pa"
)

func main() {
	archiveCommand := flag.NewFlagSet("archive", flag.ExitOnError)
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)

	archiveType := archiveCommand.String("to", "local", "choose archive destination")
	showItem := showCommand.String("all", "", "choose all of archived projects")

	switch os.Args[1] {
	case "archive":
		archiveCommand.Parse(os.Args[2:])
	case "show":
		showCommand.Parse(os.Args[2:])
	}

	archiver, err := pa.New()
	if err != nil {
		log.Fatal(err)
	}

	if archiveCommand.Parsed() {
		fmt.Println(*archiveType)
		log.Fatal(archiver.Archive("", pa.Local))
	}

	if showCommand.Parsed() {
		fmt.Println(*showItem)
	}
}
