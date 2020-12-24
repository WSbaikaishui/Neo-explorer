package api

import "neophora/lib/trans"

// GetHeaderByHashLEInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByHashLEInHex(args struct {
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
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{tr.V.(string)},
	}, ret)
}
