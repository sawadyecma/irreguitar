package chord

import "github.com/sawadyecma/irreguitar/sound"

type consistNote struct {
	note     sound.Absnote
	interval int
}

type consistNotes []consistNote

func (r consistNotes) filter(judge func(ele consistNote) bool) consistNotes {
	ret := make([]consistNote, 0, len(r))
	for i := range r {
		if judge(r[i]) {
			ret = append(ret, r[i])
		}
	}
	return ret
}

func (r consistNotes) lowersByRoot(root sound.Absnote) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note) < 0
	})
}

func (r consistNotes) highersByRoot(root sound.Absnote) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note) > 0
	})
}

func (r consistNotes) matchIntervals(root sound.Absnote, interval int) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note)%12 == interval
	})
}

func (r consistNotes) getThirdType(root sound.Absnote) lenType {
	containMinor := len(r.matchIntervals(root, 3)) != 0
	containMajor := len(r.matchIntervals(root, 4)) != 0

	return judgeLenType(containMinor, containMajor)
}

func (r consistNotes) getSeventhType(root sound.Absnote) lenType {
	containMinor := len(r.matchIntervals(root, 10)) != 0
	containMajor := len(r.matchIntervals(root, 11)) != 0

	return judgeLenType(containMinor, containMajor)
}

func judgeLenType(containMinor bool, containMajor bool) lenType {
	if containMajor {
		if containMinor {
			return lenTypeInvalid
		}
		return lenTypeMajor
	}

	if containMinor {
		return lenTypeMinor
	}

	return lenTypeNone
}
