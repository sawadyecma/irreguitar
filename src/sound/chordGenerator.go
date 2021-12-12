package sound

type chordGenerator struct {
	parser ChordParser
}

type CgConfig struct {
	Parser ChordParser
}

func NewChordGenerator(conf CgConfig) ChordGenerator {
	g := new(chordGenerator)

	if conf.Parser == nil {
		g.parser = chordParser{}
	} else {
		g.parser = conf.Parser
	}

	return g
}

func (g *chordGenerator) SetParser(
	p ChordParser,
) {
	g.parser = p
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
