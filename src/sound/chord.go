package sound

import (
	"fmt"
	"reflect"
)

type chord struct {
	// not lowest note
	rootNote absnote
	notes    []consistNote
}

type consistNote struct {
	absnote
	interval int
}

func NewChord(rootNote Absnote, notes []Absnote) Chord {

	noteVals := make([]consistNote, len(notes))
	for i := range notes {
		an := reflect.ValueOf(notes[i]).Interface().(absnote)
		noteVals[i] = consistNote{
			absnote:  an,
			interval: rootNote.Diff(an),
		}
	}
	return chord{
		rootNote: reflect.ValueOf(rootNote).Interface().(absnote),
		notes:    noteVals,
	}
}

func (c chord) Name() string {
	// example: C Cm CM7 Cm7 C7 C9 Cadd9 Cm9 C11 C13 C7/E

	// 分数コード
	// ルートより低い音が1つあれば、分数コード
	// ルートより低い音が2つ以上あったら命名しない
	lowerRoots := c.getLowerNoteByRoot()
	if len(lowerRoots) == 1 {
		return fmt.Sprintf("%s/%s", c.rootNote, lowerRoots[0].absnote)
	}

	if len(lowerRoots) == 2 {
		panic("invalid on chord")
	}

	return fmt.Sprintf("%s", c.rootNote)
}

func (c chord) getLowerNoteByRoot() []consistNote {
	ret := make([]consistNote, 0, len(c.notes))
	for i := range c.notes {
		if c.rootNote.Diff(c.notes[i].absnote) < 0 {
			ret = append(ret, c.notes[i])
		}
	}
	return ret
}
