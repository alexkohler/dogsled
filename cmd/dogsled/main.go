package main

import (
	"flag"
	"go/build"
	"log"
	"os"

	"github.com/alexkohler/dogsled"
)

func init() {
	build.Default.UseAllFiles = false
}

func usage() {
	log.Printf("Usage of %s:\n", os.Args[0])
	log.Printf("\ndogsled[flags] # runs on package in current directory\n")
	log.Printf("\ndogsled [flags] [packages]\n")
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func main() {

	// Remove log timestamp
	log.SetFlags(0)

	includeTests := flag.Bool("tests", true, "include test (*_test.go) files")
	maxBlankIdentifiers := flag.Int("n", 2, "maximum number of blank identifiers allowed in an assignment statement")
	setExitStatus := flag.Bool("set_exit_status", false, "Set exit status to 1 if any issues are found")
	flag.Usage = usage
	flag.Parse()

	flags := dogsled.Flags{
		IncludeTests:             *includeTests,
		BlankIdentifierThreshold: *maxBlankIdentifiers,
		SetExitStatus:            *setExitStatus,
	}

	if err := dogsled.CheckForDogSledding(flag.Args(), flags); err != nil {
		log.Println(err)
	}
}
