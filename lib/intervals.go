package lib

type Interval struct {
	Name      string
	ShortName string
	Semitones int
}

var (
	IntP1 = Interval{Name: "Unison", ShortName: "P1", Semitones: 0}
	Intm2 = Interval{Name: "Minor Second", ShortName: "m2", Semitones: 1}
	IntM2 = Interval{Name: "Major Second", ShortName: "M2", Semitones: 2}
	Intm3 = Interval{Name: "Minor Third", ShortName: "m3", Semitones: 3}
	IntM3 = Interval{Name: "Major Third", ShortName: "M3", Semitones: 4}
	IntP4 = Interval{Name: "Perfect Fourth", ShortName: "P4", Semitones: 5}
	IntTT = Interval{Name: "Tritone", ShortName: "TT", Semitones: 6} // Augmented fourth or diminished fifth
	IntP5 = Interval{Name: "Perfect Fifth", ShortName: "P5", Semitones: 7}
	Intm6 = Interval{Name: "Minor Sixth", ShortName: "m6", Semitones: 8}
	IntM6 = Interval{Name: "Major Sixth", ShortName: "M6", Semitones: 9}
	Intm7 = Interval{Name: "Minor Seventh", ShortName: "m7", Semitones: 10}
	IntM7 = Interval{Name: "Major Seventh", ShortName: "M7", Semitones: 11}
	IntP8 = Interval{Name: "Octave", ShortName: "P8", Semitones: 12}
)
var intervals = []Interval{
	IntP1,
	Intm2, IntM2,
	Intm3, IntM3,
	IntP4, IntTT,
	IntP5,
	Intm6, IntM6,
	Intm7, IntM7,
	IntP8,
}
