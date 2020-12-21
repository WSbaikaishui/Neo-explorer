package api

import (
	"encoding/binary"
	"neophora/lib/trans"
	"neophora/var/stderr"
)

// GetTransactionSystemFeeByHashLEInUint64 ...
func (me *T) GetTransactionSystemFeeByHashLEInUint64(args struct {
	Hash string
}, ret *uint64) error {
	var result []byte
	tr := &trans.T{V: args.Hash}
	if err := tr.HexReverse(); err != nil {
		return err
	}
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.fos",
		Index:  "h256.trx",
		Keys:   []string{tr.V.(string)},
	}, &result); err != nil {
		return err
	}
	if len(result) != 8 {
		return stderr.ErrNotFound
	}
	*ret = binary.BigEndian.Uint64(result)
	return nil
}
