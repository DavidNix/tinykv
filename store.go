package tinykv

type data struct {
	Value string
}

type Store struct {
	db map[string]string
}

func NewStore() *Store {
	return &Store{
		db: make(map[string]string),
	}
}

func (store *Store) Set(k, v string) {
	store.db[k] = v
}

func (store *Store) Get(k string) string {
	found, ok := store.db[k]
	if !ok {
		return "NULL"
	}
	return found
}
