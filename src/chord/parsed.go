package chord

type lenType int

const (
	lenTypeNone lenType = iota
	lenTypeMajor
	lenTypeMinor
	lenTypeInvalid
)

type parsed struct {
	thirdType   lenType
	seventhType lenType
}

func (r parsed) ChordType() string {
	switch r.thirdType {
	case lenTypeNone: // Power Chord
		{
			return ""
		}
	case lenTypeMajor:
		{
			switch r.seventhType {
			case lenTypeNone:
				return ""
			case lenTypeMajor:
				return "M7"
			case lenTypeMinor:
				return "7"
			}
		}
	case lenTypeMinor:
		{
			switch r.seventhType {
			case lenTypeNone:
				return "m"
			case lenTypeMajor:
				return "mM7"
			case lenTypeMinor:
				return "m7"
			}
		}
	}
	panic("invalid chord")
}
