package tinykv

type Tuple struct {
	Value   string
	Deleted bool
}

type Store struct {
	db    map[string]*Tuple
	index Index
}

func NewStore() *Store {
	return &Store{
		db:    make(map[string]*Tuple),
		index: make(Index),
	}
}

func (store *Store) Set(k, v string) *Tuple {
	existing, ok := store.db[k]
	if !ok {
		tup := &Tuple{Value: v}
		store.db[k] = tup
		store.index.Add(tup)
		return tup
	}
	store.index.Remove(existing)
	existing.Value = v
	existing.Deleted = false
	store.index.Add(existing)
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

func (store *Store) Count(k string) int {
	return store.index.Count(k)
}
