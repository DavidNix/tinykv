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
	if tx := db.currentTx(); tx != nil {
		if found, ok := tx.Get(k); ok {
			return found
		}
	}
	found, _ := db.committed.Get(k)
	return found
}

func (db *Database) Set(k, v string) {
	if tx := db.currentTx(); tx != nil {
		tx.Set(k, v)
		return
	}
	db.committed.Set(k, v)
}

func (db *Database) Delete(k string) {
	if tx := db.currentTx(); tx != nil {
		tx.Delete(k)
		return
	}
	db.committed.Delete(k)
}

func (db *Database) Count(val string) int {
	return db.committed.Count(val)
}

func (db *Database) Begin() {
	var newTx *Store
	if tx := db.currentTx(); tx != nil {
		newTx = tx.Clone()
	} else {
		newTx = NewStore()
	}
	db.txs = append(db.txs, newTx)
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
