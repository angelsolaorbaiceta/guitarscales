package lib

import "fmt"

type ChordQuality string

const (
	ChordQualityUnknown    ChordQuality = "?"
	ChordQualityMajor      ChordQuality = ""
	ChordQualityMinor      ChordQuality = "m"
	ChordQualityDiminished ChordQuality = "dim"
	ChordQualityAugmented  ChordQuality = "aug"
)

func intervalsToQualityKey(first, second Interval) string {
	return fmt.Sprintf("%s_%s", first, second)
}

var (
	// A major quality = Major 3rd + Perfect 5th
	majKey = intervalsToQualityKey(IntM3, IntP5)
	// A minor quality = Minor 3rd + Perfect 5th
	minKey = intervalsToQualityKey(Intm3, IntP5)
	// A diminished quality = Minor 3rd + Tritone
	dimKey = intervalsToQualityKey(Intm3, IntTT)
	// An augmented quality = Major 3rd + Minor 6th
	augKey = intervalsToQualityKey(IntM3, Intm6)
)
var intervalsToQualities = map[string]ChordQuality{
	majKey: ChordQualityMajor,
	minKey: ChordQualityMinor,
	dimKey: ChordQualityDiminished,
	augKey: ChordQualityAugmented,
}

type Chord struct {
	RootNote  Note
	intervals []Interval
}

func (c Chord) Quality() ChordQuality {
	key := intervalsToQualityKey(c.intervals[0], c.intervals[1])
	if quality, ok := intervalsToQualities[key]; !ok {
		return ChordQualityUnknown
	} else {
		return quality
	}
}

func (c Chord) String() string {
	return fmt.Sprintf("%s%s", c.RootNote, c.Quality())
}
