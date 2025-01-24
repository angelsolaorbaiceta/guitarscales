package lib

import (
	"fmt"
	"testing"
)

func TestNotes(t *testing.T) {
	for i, note := range notes {
		t.Run(fmt.Sprintf("Intervals from C to %s", note), func(t *testing.T) {
			got := NoteC.IntervalTo(note)
			want := intervals[i]

			if got != want {
				t.Errorf("Want %s, got %s", want, got)
			}
		})

		t.Run(fmt.Sprintf("Intervals from B to %s", note), func(t *testing.T) {
			got := NoteB.IntervalTo(note)
			want := intervals[(i+1)%len(notes)]

			if got != want {
				t.Errorf("Want %s, got %s", want, got)
			}
		})
	}

	t.Run("Note plus unison interval", func(t *testing.T) {
		got := NoteC.PlusInterval(IntP1)
		want := NoteC

		if got != want {
			t.Errorf("Want %s, got %s", want, got)
		}
	})

	t.Run("Note plus minor second interval", func(t *testing.T) {
		got := NoteC.PlusInterval(Intm2)
		want := NoteCSharp

		if got != want {
			t.Errorf("Want %s, got %s", want, got)
		}
	})
}
