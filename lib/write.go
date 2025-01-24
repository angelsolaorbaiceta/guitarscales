package lib

import (
	"fmt"
	"io"
	"strings"
)

var fretsToNumber = map[int]struct{}{
	1:  {},
	3:  {},
	5:  {},
	7:  {},
	9:  {},
	12: {},
	15: {},
	17: {},
	19: {},
	21: {},
	24: {},
	27: {},
}

func PrintScale(w io.Writer, scale Scale, root Note) {
	w.Write([]byte(fmt.Sprintf("The %s Scale (root=%s)\n", scale.Name, root)))
	w.Write([]byte("\tNotes: "))
	for _, note := range scale.NotesWithRoot(root) {
		w.Write([]byte(fmt.Sprintf("%s ", note)))
	}
	w.Write([]byte("\n\tIntervals: "))
	for _, interval := range scale.Intervals {
		w.Write([]byte(fmt.Sprintf("%s ", interval)))
	}
	w.Write([]byte("\n\n"))
}

func printNotesInString(w io.Writer, stringNote Note, notes []Note, frets int) {
	var (
		rootNote    = notes[0]
		notePos     = make(map[Note]int)
		currentNote = stringNote
	)
	for i, note := range notes {
		notePos[note] = i + 1
	}

	w.Write([]byte(
		fmt.Sprintf("%2s|", stringNote),
	))
	for fret := 1; fret <= frets; fret++ {
		currentNote = currentNote.PlusInterval(Intm2)

		if currentNote == rootNote {
			w.Write([]byte("-R-|"))
		} else if _, ok := notePos[currentNote]; ok {
			w.Write([]byte("-O-|"))
		} else {
			w.Write([]byte("---|"))
		}
	}
	w.Write([]byte("\n"))
}

func PrintGuitarNotes(w io.Writer, notes []Note, tuning Tuning, frets int) {
	for _, note := range tuning {
		printNotesInString(w, note, notes, frets)
	}

	// Write fret numbers.
	w.Write([]byte("   "))

	for fret := 1; fret <= frets; fret++ {
		if _, ok := fretsToNumber[fret]; ok {
			w.Write([]byte(fmt.Sprintf("%2d |", fret)))
		} else {
			w.Write([]byte("   |"))
		}
	}

	w.Write([]byte("\n"))
}

func chordsToString(chords []Chord) string {
	var builder strings.Builder

	for _, chord := range chords {
		builder.WriteString(fmt.Sprintf("%s ", chord))
	}

	return builder.String()
}
