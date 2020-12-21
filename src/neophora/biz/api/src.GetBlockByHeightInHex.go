package api

import "fmt"

// GetBlockByHeightInHex ...
func (me *T) GetBlockByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
