package main

import "os"

import "flag"

import "fmt"

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

	if archiveCommand.Parsed() {
		fmt.Println(*archiveType)
	}

	if showCommand.Parsed() {
		fmt.Println(*showItem)
	}
}
