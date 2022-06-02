package tinykv

type tuple struct {
	Value   string
	Deleted bool
}

type Store struct {
	db    map[string]tuple
	index map[string]int
}

func NewStore() *Store {
	return &Store{
		db:    make(map[string]tuple),
		index: make(map[string]int),
	}
}

func (store *Store) Set(k, v string) {
	old, ok := store.db[k]
	if !ok {
		store.index[v]++
	} else if v != old.Value {
		store.index[old.Value]--
	}
	store.db[k] = tuple{Value: v}
}

func (store *Store) Get(k string) (entry string, tupleExists bool) {
	const null = "NULL"
	found, ok := store.db[k]
	if !ok {
		return null, false
	}
	if found.Deleted {
		return null, true
	}
	return found.Value, true
}

func (store *Store) Delete(k string) {
	found := store.db[k]
	if !found.Deleted {
		store.index[found.Value]--
	}
	found.Deleted = true
	store.db[k] = found
}

func (store *Store) Count(val string) int {
	return store.index[val]
}

func (store *Store) Clone() *Store {
	clone := NewStore()
	for k, v := range store.db {
		clone.db[k] = v
	}
	for k, v := range store.index {
		clone.index[k] = v
	}
	return clone
}
