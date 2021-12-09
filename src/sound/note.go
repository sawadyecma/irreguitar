package sound

// TODO
// そのそもNoteが必要か

type note struct{}

func NewNote() Note {
	return note{}
}
func (note) String() string {
	return "A1"
}
