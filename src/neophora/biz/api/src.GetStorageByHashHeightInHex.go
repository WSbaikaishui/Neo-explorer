package api

import "fmt"

// GetStorageByHashHeightInHex ...
func (me *T) GetStorageByHashHeightInHex(args struct {
	Hash   string
	Height uint64
}, ret *string) error {
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h256-key-uint.hgt",
		Keys:   []string{args.Hash, fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
