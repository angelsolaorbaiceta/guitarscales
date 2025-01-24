package lib

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	t.Run("Notes of the major scale in C", func(t *testing.T) {
		notes := ScaleMajor.NotesWithRoot(NoteC)
		want := []Note{NoteC, NoteD, NoteE, NoteF, NoteG, NoteA, NoteB}

		if !reflect.DeepEqual(notes, want) {
			t.Errorf("Want %v, got %v", want, notes)
		}
	})

	t.Run("Chords of the major scale in C", func(t *testing.T) {
		chords := ScaleMajor.ChordsWithRoot(NoteC)
		want := "C Dm Em F G Am Bdim "

		if got := chordsToString(chords); got != want {
			t.Errorf("Want '%s', got '%s'", want, got)
		}
	})
}
