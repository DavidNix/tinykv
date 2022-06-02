package tinykv

type Tuple struct {
	Key     string
	Value   string
	Deleted bool
}

type Store struct {
	db    map[string]*Tuple
	Index Index
}

func NewStore() *Store {
	return &Store{
		db:    make(map[string]*Tuple),
		Index: make(Index),
	}
}

func (store *Store) Set(k, v string) *Tuple {
	existing, ok := store.db[k]
	if !ok {
		tup := &Tuple{Key: k, Value: v}
		store.db[k] = tup
		store.Index.Add(tup)
		return tup
	}

	if existing.Value != v {
		existing.Deleted = true
		tup := &Tuple{Key: k, Value: v}
		store.db[k] = tup
		store.Index.Add(tup)
		return tup
	}

	existing.Value = v
	existing.Deleted = false
	store.Index.Add(existing)
	return existing
}

func (store *Store) Get(k string) (*Tuple, bool) {
	found, ok := store.db[k]
	if !ok {
		return nil, false
	}
	return found, true
}

func (store *Store) Delete(k string) (*Tuple, bool) {
	found, ok := store.db[k]
	if !ok {
		return nil, false
	}
	found.Deleted = true
	return found, true
}

func (store *Store) Count(val string) int {
	return store.Index.Count(val)
}
