package guitar

import "github.com/sawadyecma/irreguitar/sound"

type Turning interface {
	OpenStrings() String
}

type String interface {
	StringNumber() int
	OpenNote() sound.Note
}
