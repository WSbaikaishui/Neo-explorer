package api

import (
	"fmt"
	"neophora/lib/trans"
)

// GetStorageByHashHeightLEInHex ...
func (me *T) GetStorageByHashHeightLEInHex(args struct {
	Hash   string
	Height uint64
}, ret *string) error {
	tr := &trans.T{V: args.Hash}
	if err := tr.HexReverse(); err != nil {
		return err
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h256-key-uint.hgt",
		Keys:   []string{tr.V.(string), fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
