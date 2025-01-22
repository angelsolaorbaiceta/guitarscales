# Guitar Scales

A CLI tool to show the notes of a scale on the fretboard.

## Installation

Clone the repo:

```bash
$ git clone https://github.com/angelsolaorbaiceta/guitarscales.git
```

Install the binary:

```bash
$ make install
```

You should now have the `guitars` binary in your path:

```
$ which guitars
```

## Usage

You can get the tool's help menu by passing it the `--help` flag:

```bash
$ guitars --help
```

This prints:

```
              .__  __
   ____  __ __|__|/  |______ _______  ______
  / ___\|  |  \  \   __\__  \\_  __ \/  ___/
 / /_/  >  |  /  ||  |  / __ \|  | \/\___ \
 \___  /|____/|__||__| (____  /__|  /____  >
/_____/                     \/           \/

Guitar Scales.
Show the positions of a scale at a given root note on the guitar.


USAGE:

  $ guitars [options] <scale> <root>

  Where <scale> is one of:
    - major/ioninan      The major scale or ionian mode
    - nminor/aeolian     The natural minor scale or aeolian mode
    - hminor             The harmonic minor scale
    - mminor             The melodic minor scale
    - phrygian/phr       The phrygian mode
    - lydian/lyd         The lydian mode
    - mixolydian/myx     The mixolidian mode
    - locrian/loc        The locrian mode
    - minpenta           The minor pentatonic scale
    - majpenta           The major pentatonic scale

  And <root> can be any root note: A, Asharp, Bflat, B, C, Csharp ...
  Note that only the first two letters of the note need to be defined for
  sharp or flat notes (e.g. As, Bf).


OPTIONS:

  --help      Prints the help menu and exits
  --frets     Then number of frets to show in the diagram (default 17)


EXAMPLES:

  $ guitars major A
  $ guitars ioninan A

  Prints the positions of the Major/Ionian scale with the root in A.

  $ guitars --frets 21 phr fsharp

  Prints the positions of the Phrygian mode in a fretboard with 21 frets.


AUTHOR:

  Written by Angel Sola Orbaiceta <https://github.com/angelsolaorbaiceta>.
```
