package lib

import (
	"fmt"
	"strings"
)

type Note struct {
	Name, AltName string
	idx           int
}

// A Tuning is the notes each string in the guitar is tuned to.
// Tunings respect the string natural ordering, starting from the
// first string, up to the last.
type Tuning []Note

var (
	// TuningStd is the standard tuning for the guitar.
	TuningStd = []Note{NoteE, NoteB, NoteG, NoteD, NoteA, NoteE}
)

var (
	NoteC      = Note{Name: "C", AltName: "", idx: 0}
	NoteCSharp = Note{Name: "C#", AltName: "Db", idx: 1}
	NoteD      = Note{Name: "D", AltName: "", idx: 2}
	NoteDSharp = Note{Name: "D#", AltName: "Eb", idx: 3}
	NoteE      = Note{Name: "E", AltName: "", idx: 4}
	NoteF      = Note{Name: "F", AltName: "", idx: 5}
	NoteFSharp = Note{Name: "F#", AltName: "Gb", idx: 6}
	NoteG      = Note{Name: "G", AltName: "", idx: 7}
	NoteGSharp = Note{Name: "G#", AltName: "Ab", idx: 8}
	NoteA      = Note{Name: "A", AltName: "", idx: 9}
	NoteASharp = Note{Name: "A#", AltName: "Bb", idx: 10}
	NoteB      = Note{Name: "B", AltName: "", idx: 11}
)
var notes = []Note{
	NoteC, NoteCSharp,
	NoteD, NoteDSharp,
	NoteE,
	NoteF, NoteFSharp,
	NoteG, NoteGSharp,
	NoteA, NoteASharp,
	NoteB,
}
var notesByName = map[string]Note{
	"a":  NoteA,
	"as": NoteASharp,
	"bb": NoteASharp,
	"b":  NoteB,
	"c":  NoteC,
	"cs": NoteCSharp,
	"db": NoteCSharp,
	"d":  NoteD,
	"ds": NoteDSharp,
	"eb": NoteDSharp,
	"e":  NoteE,
	"f":  NoteF,
	"fs": NoteFSharp,
	"gb": NoteFSharp,
	"g":  NoteG,
	"gs": NoteGSharp,
	"ab": NoteGSharp,
}

func (n Note) String() string {
	return n.Name
}

func (i Interval) String() string {
	return i.ShortName
}

// IntervalBetween returns the iterval that separates this note and another one.
func (n Note) IntervalTo(other Note) Interval {
	var semitones = n.idx - other.idx
	if semitones < 0 {
		semitones = -semitones
	}

	return intervals[semitones]
}

// PlusInterval returns the note that result from adding an interval to this note.
func (n Note) PlusInterval(interval Interval) Note {
	idx := (n.idx + interval.Semitones) % len(notes)
	return notes[idx]
}

func NoteWithName(name string) (Note, error) {
	// Just keep at most two letters for each note
	if len(name) > 2 {
		name = name[:2]
	}

	note, ok := notesByName[strings.ToLower(name)]
	if !ok {
		return Note{}, fmt.Errorf("Unknown note: '%s'", name)
	}

	return note, nil
}
