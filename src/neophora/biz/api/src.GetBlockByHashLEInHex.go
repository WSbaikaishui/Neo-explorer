package api

import "neophora/lib/trans"

// GetBlockByHashLEInHex ...
func (me *T) GetBlockByHashLEInHex(args struct {
	Hash string
}, ret *string) error {
	tr := &trans.T{V: args.Hash}
	if err := tr.HexReverse(); err != nil {
		return err
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{tr.V.(string)},
	}, ret)
}
