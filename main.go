package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/angelsolaorbaiceta/guitars/lib"
)

//go:embed usage.txt
var usageFS embed.FS

const (
	defaultFretsToShow = 17
)

func main() {
	var (
		helpFlag  = flag.Bool("help", false, "Show help")
		fretsFlag = flag.Int("frets", defaultFretsToShow, "Number of frets to show")
	)

	flag.Parse()
	args := flag.Args()

	if *helpFlag || len(args) < 2 {
		showHelp()
		return
	}

	scale, err := lib.ScaleWithName(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	rootNote, err := lib.NoteWithName(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	notes := scale.NotesWithRoot(rootNote)

	lib.PrintScale(os.Stdout, scale, rootNote)
	lib.PrintGuitarNotes(os.Stdout, notes, lib.TuningStd, *fretsFlag)
}

func showHelp() {
	usageFile, err := usageFS.Open("usage.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer usageFile.Close()

	_, err = io.Copy(os.Stdout, usageFile)
	if err != nil {
		log.Fatal(err)
	}
}
