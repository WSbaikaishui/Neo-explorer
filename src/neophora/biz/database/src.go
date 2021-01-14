package database

import (
	"log"

	"github.com/tecbot/gorocksdb"
)

// T ...
type T struct {
	DB *gorocksdb.DB
	WO *gorocksdb.WriteOptions
	RO *gorocksdb.ReadOptions
}

// Put ...
func (me *T) Put(args map[string][]byte, ret *bool) error {
	log.Println("[RPC][Put]", len(args))
	batch := gorocksdb.NewWriteBatch()
	defer batch.Destroy()
	for k, v := range args {
		batch.Put([]byte(k), v)
	}
	if err := me.DB.Write(me.WO, batch); err != nil {
		log.Println("[RPC][Put]", err)
		return err
	}
	*ret = true
	return nil
}

// Get ...
func (me *T) Get(args struct {
	Key []byte
}, ret *[]byte) error {
	result, err := me.DB.GetBytes(me.RO, args.Key)
	if err != nil {
		log.Println("[RPC][Get]", err)
		return err
	}
	*ret = result
	return nil
}

// GetLastVal ...
func (me *T) GetLastVal(args struct {
	Key    []byte
	Prefix uint
}, ret *[]byte) error {
	klen := uint(len(args.Key))
	if args.Prefix > klen {
		args.Prefix = klen
	}
	it := me.DB.NewIterator(me.RO)
	defer it.Close()
	it.SeekForPrev(args.Key)
	if it.ValidForPrefix(args.Key[:args.Prefix]) == false {
		return nil
	}
	result := it.Value().Data()
	*ret = make([]byte, len(result))
	copy(*ret, result)
	return nil
}

// ListKey ...
func (me *T) ListKey(args struct {
	Key    []byte
	Prefix uint
	C      uint
}, ret *[][]byte) error {
	klen := uint(len(args.Key))
	if args.Prefix > klen {
		args.Prefix = klen
	}
	if args.C > 0x10000 {
		args.C = 0x10000
	}
	it := me.DB.NewIterator(me.RO)
	defer it.Close()
	it.Seek(args.Key)
	results := make([][]byte, 0, args.C)
	for i := uint(0); i < args.C; i++ {
		if it.ValidForPrefix(args.Key[:args.Prefix]) == false {
			return nil
		}
		data := it.Key().Data()
		result := make([]byte, len(data))
		copy(result, data)
		results = append(results, result)
	}
	return nil
}

// ListVal ...
func (me *T) ListVal(args struct {
	Key    []byte
	Prefix uint
	C      uint
}, ret *[][]byte) error {
	klen := uint(len(args.Key))
	if args.Prefix > klen {
		args.Prefix = klen
	}
	if args.C > 0x10000 {
		args.C = 0x10000
	}
	it := me.DB.NewIterator(me.RO)
	defer it.Close()
	it.Seek(args.Key)
	results := make([][]byte, 0, args.C)
	for i := uint(0); i < args.C; i++ {
		if it.ValidForPrefix(args.Key[:args.Prefix]) == false {
			return nil
		}
		data := it.Value().Data()
		result := make([]byte, len(data))
		copy(result, data)
		results = append(results, result)
	}
	return nil
}

// GetLastKey ...
func (me *T) GetLastKey(args struct {
	Key    []byte
	Prefix uint
}, ret *[]byte) error {
	klen := uint(len(args.Key))
	if args.Prefix > klen {
		args.Prefix = klen
	}
	it := me.DB.NewIterator(me.RO)
	defer it.Close()
	it.SeekForPrev(args.Key)
	if it.ValidForPrefix(args.Key[:args.Prefix]) == false {
		return nil
	}
	result := it.Key().Data()
	*ret = make([]byte, len(result))
	copy(*ret, result)
	return nil
}

// DeleteRange ...
func (me *T) DeleteRange(args struct {
	Key []byte
}, ret *bool) error {
	it := me.DB.NewIterator(me.RO)
	defer it.Close()
	it.Seek(args.Key)
	if it.ValidForPrefix(args.Key) {
		if err := me.DB.Delete(me.WO, it.Key().Data()); err != nil {
			return err
		}
	}
	*ret = true
	return nil
}
