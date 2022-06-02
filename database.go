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
	for _, tx := range append(db.txs, db.committed) {
		if found, ok := tx.Get(k); ok {
			if found.Deleted {
				return null
			}
			return found.Value
		}
	}
	return null
}

func (db *Database) Set(k, v string) {
	if tx := db.currentTx(); tx != nil {
		tx.Set(k, v)
		return
	}
	db.committed.Set(k, v)
}

func (db *Database) Delete(k string) {
	var foundVal string
	reverse := append([]*Store{db.committed}, db.txs...)
	for i := len(reverse) - 1; i < len(reverse); i-- {
		v, ok := reverse[i].Get(k)
		if ok {
			foundVal = v.Value
			break
		}
	}
	if tx := db.currentTx(); tx != nil {
		if len(foundVal) > 0 {
			tx.Set(k, foundVal)
		}
		tx.Delete(k)
		return
	}
	db.committed.Delete(k)
}

func (db *Database) Count(val string) int {
	return 0
}

func (db *Database) Begin() {
	db.txs = append(db.txs, NewStore())
}

func (db *Database) Rollback() {
	if len(db.txs) == 0 {
		return
	}
	db.txs = db.txs[:len(db.txs)-1]
}

func (db *Database) currentTx() *Store {
	if len(db.txs) == 0 {
		return nil
	}
	return db.txs[len(db.txs)-1]
}
