package tinykv

type Database struct {
	committed *Store
	txs       []*Store
}

func NewDatabase() *Database {
	return &Database{
		committed: NewStore(),
		txs:       nil,
	}
}

func (db *Database) Get(k string) string {
	return db.committed.Get(k)
}

func (db *Database) Set(k, v string) {
	db.committed.Set(k, v)
}

func (db *Database) Delete(k string) {
	db.committed.Delete(k)
}

func (db *Database) Count(val string) int {
	return db.committed.Count(val)
}
