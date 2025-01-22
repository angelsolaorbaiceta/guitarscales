package lib

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	t.Run("Major scale in C", func(t *testing.T) {
		notes := ScaleMajor.NotesWithRoot(NoteC)
		want := []Note{NoteC, NoteD, NoteE, NoteF, NoteG, NoteA, NoteB}

		if !reflect.DeepEqual(notes, want) {
			t.Errorf("Want %v, got %v", want, notes)
		}

	})
}
