package lib

import "testing"

func TestNotes(t *testing.T) {
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
