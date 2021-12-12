package sound

type chordGenerator struct {
	parser ChordParser
}

func NewChordGenerator(parser ChordParser) ChordGenerator {
	g := new(chordGenerator)

	g.parser = parser

	return g
}

func (g chordGenerator) Generate(
	rootNote Absnote,
	notes []Absnote,
) Chord {
	return NewChord(
		rootNote,
		notes,
		g.parser.Parse(rootNote, notes),
	)
}
