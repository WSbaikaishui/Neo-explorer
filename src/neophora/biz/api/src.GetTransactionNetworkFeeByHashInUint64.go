package api

import (
	"encoding/binary"
	"neophora/var/stderr"
)

// GetTransactionNetworkFeeByHashInUint64 ...
func (me *T) GetTransactionNetworkFeeByHashInUint64(args struct {
	Hash string
}, ret *uint64) error {
	var result []byte
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "uint.fon",
		Index:  "h256.trx",
		Keys:   []string{args.Hash},
	}, &result); err != nil {
		return err
	}
	if len(result) != 8 {
		return stderr.ErrNotFound
	}
	*ret = binary.BigEndian.Uint64(result)
	return nil
}
