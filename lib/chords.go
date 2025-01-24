package lib

import "fmt"

type ChordQuality string

const (
	ChordQualityUnknown    ChordQuality = "?"
	ChordQualityMajor      ChordQuality = ""
	ChordQualityMinor      ChordQuality = "m"
	ChordQualityDiminished ChordQuality = "dim"
	ChordQualityAugmented  ChordQuality = "aug"
	ChordQualitySuspended2 ChordQuality = "sus2"
	ChordQualitySuspended4 ChordQuality = "sus4"
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
	// An augmented quality is derived from two major 3rd intervals
	augKey = intervalsToQualityKey(IntM3, IntM3)
	// A suspended second quality is derived from a major 2nd and perfect 4th intervals
	sus2Key = intervalsToQualityKey(IntM2, IntP4)
	// A suspended fourth quality is derived from a perfect 4th and major 2nd intervals
	sus4Key = intervalsToQualityKey(IntP4, IntM2)
)
var intervalsToQualities = map[string]ChordQuality{
	majKey:  ChordQualityMajor,
	minKey:  ChordQualityMinor,
	dimKey:  ChordQualityDiminished,
	augKey:  ChordQualityAugmented,
	sus2Key: ChordQualitySuspended2,
	sus4Key: ChordQualitySuspended4,
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
