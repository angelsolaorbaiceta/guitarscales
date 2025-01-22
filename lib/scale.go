package lib

import (
	"fmt"
	"strings"
)

type Scale struct {
	Name      string
	Intervals []Interval
}

var (
	ScaleMajor = Scale{
		Name:      "Major",
		Intervals: []Interval{IntP1, IntM2, IntM3, IntP4, IntP5, IntM6, IntM7},
	}
	ScaleMinor = Scale{
		Name:      "Natural Minor",
		Intervals: []Interval{IntP1, IntM2, Intm3, IntP4, IntP5, Intm6, Intm7},
	}
	ScaleHarmonicMinor = Scale{
		Name:      "Harmonic Minor",
		Intervals: []Interval{IntP1, IntM2, Intm3, IntP4, IntP5, Intm6, IntM7},
	}
	ScaleMelodicMinor = Scale{
		Name:      "Melodic Minor",
		Intervals: []Interval{IntP1, IntM2, Intm3, IntP4, IntP5, IntM6, IntM7},
	}
	ScaleMajorPentatonic = Scale{
		Name:      "Major Pentatonic",
		Intervals: []Interval{IntP1, IntM2, IntM3, IntP5, IntM6},
	}
	ScaleMinorPentatonic = Scale{
		Name:      "Minor Pentatonic",
		Intervals: []Interval{IntP1, Intm3, IntP4, IntP5, Intm7},
	}
	ScaleDorian = Scale{
		Name:      "Dorian",
		Intervals: []Interval{IntP1, IntM2, Intm3, IntP4, IntP5, IntM6, Intm7},
	}
	ScalePhrygian = Scale{
		Name:      "Phrygian",
		Intervals: []Interval{IntP1, Intm2, Intm3, IntP4, IntP5, Intm6, Intm7},
	}
	ScaleLydian = Scale{
		Name:      "Lydian",
		Intervals: []Interval{IntP1, IntM2, IntM3, IntTT, IntP5, IntM6, IntM7},
	}
	ScaleMixolydian = Scale{
		Name:      "Mixolydian",
		Intervals: []Interval{IntP1, IntM2, IntM3, IntP4, IntP5, IntM6, Intm7},
	}
	ScaleLocrian = Scale{
		Name:      "Locrian",
		Intervals: []Interval{IntP1, Intm2, Intm3, IntP4, IntTT, Intm6, Intm7},
	}
)

var scalesByName = map[string]Scale{
	"major":  ScaleMajor,
	"ionian": ScaleMajor,

	"minor":   ScaleMinor,
	"nminor":  ScaleMinor,
	"aeolian": ScaleMinor,
	"hminor":  ScaleHarmonicMinor,
	"mminor":  ScaleMelodicMinor,

	"majpenta": ScaleMajorPentatonic,
	"minpenta": ScaleMinorPentatonic,

	"dorian": ScaleDorian,
	"dor":    ScaleDorian,

	"phrygian": ScalePhrygian,
	"phr":      ScalePhrygian,

	"lydian": ScaleLydian,
	"lyd":    ScaleLydian,

	"mixolydian": ScaleMixolydian,
	"mix":        ScaleMixolydian,

	"locrian": ScaleLocrian,
	"loc":     ScaleLocrian,
}

// NotesWithRoot returns the slice of all notes in the scale, starting
// at the given root note.
func (s Scale) NotesWithRoot(root Note) []Note {
	notes := make([]Note, len(s.Intervals))
	for i, interval := range s.Intervals {
		notes[i] = root.PlusInterval(interval)
	}

	return notes
}

func ScaleWithName(name string) (Scale, error) {
	scale, ok := scalesByName[strings.ToLower(name)]
	if !ok {
		return Scale{}, fmt.Errorf("Uknown scale: '%s'", name)
	}

	return scale, nil
}
