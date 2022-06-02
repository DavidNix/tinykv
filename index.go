package tinykv

import "errors"

type Index map[string]map[*Tuple]struct{}

func (idx Index) Add(tup *Tuple) {
	idx.panicIfNil(tup)
	val := tup.Value
	set := idx[val]
	if set == nil {
		set = make(map[*Tuple]struct{})
	}
	set[tup] = struct{}{}
	idx[val] = set
}

func (idx Index) Remove(tup *Tuple) {
	idx.panicIfNil(tup)
	set := idx[tup.Value]
	if set == nil {
		return
	}
	delete(set, tup)
}

func (idx Index) Count(val string) int {
	set := idx[val]
	var count int
	for tup := range set {
		if tup.Deleted {
			continue
		}
		count++
	}
	return count
}

func (idx Index) panicIfNil(tup *Tuple) {
	if tup == nil {
		panic(errors.New("nil tuple"))
	}
}
