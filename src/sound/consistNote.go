package sound

type consistNote struct {
	note     Absnote
	interval int
}

type consistNotes []consistNote

func (r consistNotes) lowersByRoot(root Absnote) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note) < 0
	})
}

func (r consistNotes) highersByRoot(root Absnote) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note) > 0
	})
}

func (r consistNotes) matchIntervals(root Absnote, interval int) consistNotes {
	return r.filter(func(ele consistNote) bool {
		return root.Diff(ele.note)%12 == interval
	})
}

func (r consistNotes) filter(judge func(ele consistNote) bool) consistNotes {
	ret := make([]consistNote, 0, len(r))
	for i := range r {
		if judge(r[i]) {
			ret = append(ret, r[i])
		}
	}
	return ret
}
