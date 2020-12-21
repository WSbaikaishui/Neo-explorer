package api

import "neophora/lib/trans"

// GetBlockByHashLEInHex ...
func (me *T) GetBlockByHashLEInHex(args struct {
	Hash string
}, ret *string) error {
	var tr trans.T
	tr.V = args.Hash
	if err := tr.HexToBytes(); err != nil {
		return err
	}
	if err := tr.BytesReverse(); err != nil {
		return err
	}
	if err := tr.BytesToHex(); err != nil {
		return err
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "block",
		Index:  "hash",
		Keys:   []string{tr.V.(string)},
	}, ret)
}
