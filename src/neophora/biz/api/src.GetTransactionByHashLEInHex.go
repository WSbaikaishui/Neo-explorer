package api

import "neophora/lib/trans"

// GetTransactionByHashLEInHex ...
func (me *T) GetTransactionByHashLEInHex(args struct {
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
		Target: "bins.trx",
		Index:  "h256.trx",
		Keys:   []string{tr.V.(string)},
	}, ret)
}
