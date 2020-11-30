package dat

import (
	"github.com/tecbot/gorocksdb"
)

// T ...
type T struct {
	db *gorocksdb.DB
	wo *gorocksdb.WriteOptions
	ro *gorocksdb.ReadOptions
}

// Init ...
func (me *T) Init(path string) error {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	defer opts.Destroy()
	db, err := gorocksdb.OpenDb(opts, path)
	if err != nil {
		return err
	}
	me.db = db
	me.wo = gorocksdb.NewDefaultWriteOptions()
	me.ro = gorocksdb.NewDefaultReadOptions()
	return nil
}

// Put ...
func (me *T) Put(k []byte, v []byte) error {
	return me.db.Put(me.wo, k, v)
}

// // Puts ...
// func (me *T) Puts(batch [][2][]byte) error {
// 	wb := gorocksdb.NewWriteBatch()
// 	defer wb.Destroy()
// 	for _, v := range batch {
// 		wb.Put(v[0], v[1])
// 	}
// 	return me.db.Write(me.wo, wb)
// }

// Get ...
func (me *T) Get(k []byte) ([]byte, error) {
	return me.db.GetBytes(me.ro, k)
}

// GetLast ...
func (me *T) GetLast(k []byte, p int) []byte {
	iter := me.db.NewIterator(me.ro)
	defer iter.Close()
	iter.SeekForPrev(k)
	if iter.ValidForPrefix(k[:p]) == false {
		return nil
	}

	val := iter.Value()
	defer val.Free()
	data := val.Data()
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret
}

// GetLastKey ...
func (me *T) GetLastKey(k []byte, p int) []byte {
	iter := me.db.NewIterator(me.ro)
	defer iter.Close()
	iter.SeekForPrev(k)
	if iter.ValidForPrefix(k[:p]) == false {
		return nil
	}

	val := iter.Key()
	defer val.Free()
	data := val.Data()
	ret := make([]byte, len(data))
	copy(ret, data)
	return ret
}

// Close ...
func (me *T) Close() {
	me.wo.Destroy()
	me.ro.Destroy()
	me.db.Close()
}
