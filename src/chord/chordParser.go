package chord

import (
	"fmt"

	"github.com/sawadyecma/irreguitar/sound"
)

type chordParser struct {
}

func (chordParser) Parse(root sound.Absnote, notes []sound.Absnote) string {
	rootName := root.Name()
	if notes == nil || len(notes) == 0 {
		return rootName
	}

	consists := make(consistNotes, len(notes))
	for i := range notes {
		interval := root.Diff(notes[i])
		consists[i] = consistNote{
			note:     notes[i],
			interval: interval,
		}
	}

	highers := consists.highersByRoot(root)

	pd := parsed{
		thirdType:   highers.getThirdType(root),
		seventhType: highers.getSeventhType(root),
	}

	lowers := consists.lowersByRoot(root)
	if len(lowers) == 1 {
		return fmt.Sprintf(
			"%s/%s",
			rootName+pd.ChordType(),
			lowers[0].note.Name(),
		)
	}
	if len(lowers) >= 2 {
		panic("invalid on chord")
	}

	return rootName + pd.ChordType()
}
