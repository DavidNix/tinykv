package tinykv

type Database struct {
	committed *Store
	txs       []*Store
}

func NewDatabase() *Database {
	return &Database{
		committed: NewStore(),
	}
}

func (db *Database) Get(k string) string {
	const null = "NULL"
	reverse := append([]*Store{db.committed}, db.txs...)
	for i := len(reverse) - 1; i >= 0; i-- {
		if found, ok := reverse[i].Get(k); ok {
			if found.Deleted {
				return null
			}
			return found.Value
		}
	}
	return null
}

func (db *Database) Set(k, v string) {
	db.currentStore().Set(k, v)
}

func (db *Database) Delete(k string) {
	var foundVal string
	reverse := append([]*Store{db.committed}, db.txs...)
	for i := len(reverse) - 1; i >= 0; i-- {
		v, ok := reverse[i].Get(k)
		if ok {
			foundVal = v.Value
			break
		}
	}
	store := db.currentStore()
	if len(foundVal) > 0 {
		store.Set(k, foundVal)
	}
	store.Delete(k)
}

func (db *Database) Count(val string) int {
	return db.currentStore().Count(val)
}

func (db *Database) Begin() {
	store := NewStore()
	store.Index = db.currentStore().Index.Clone()
	db.txs = append(db.txs, store)
}

func (db *Database) Rollback() bool {
	if len(db.txs) == 0 {
		return false
	}
	db.txs = db.txs[:len(db.txs)-1]
	return true
}

func (db *Database) currentStore() *Store {
	if len(db.txs) == 0 {
		return db.committed
	}
	return db.txs[len(db.txs)-1]
}
