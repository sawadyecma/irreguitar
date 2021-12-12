
type lenType int

const (
	lenTypeMajor lenType = iota
	lenTypeMinor
	lenTypeNone
)

type parsed struct {
	thirdType   lenType
	seventhType lenType
}

func (r parsed) ChordType() string {
	switch r.thirdType {
	case lenTypeMajor:
		{
			switch r.seventhType {
			case lenTypeMajor:
				return "M7"
			case lenTypeMinor:
				return "7"
			case lenTypeNone:
				return ""
			}
		}
	case lenTypeMinor:
		{
			switch r.seventhType {
			case lenTypeMajor:
				return "mM7"
			case lenTypeMinor:
				return "m7"
			case lenTypeNone:
				return "m"
			}
		}
	case lenTypeNone: // Power Chord
		{
			return ""
		}
	}
	panic("invalid chord")
}
