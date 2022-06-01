package tinykv

type data struct {
	Value   string
	Deleted bool
}

type Store struct {
	db map[string]data
}

func NewStore() *Store {
	return &Store{
		db: make(map[string]data),
	}
}

func (store *Store) Set(k, v string) {
	store.db[k] = data{Value: v}
}

func (store *Store) Get(k string) string {
	const null = "NULL"
	found, ok := store.db[k]
	if !ok {
		return null
	}
	if found.Deleted {
		return null
	}
	return found.Value
}

func (store *Store) Delete(k string) {
	found, ok := store.db[k]
	if !ok {
		return
	}
	found.Deleted = true
	store.db[k] = found
}
