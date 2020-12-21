package api

import "fmt"

// GetContractByHashHeightInHex ...
func (me *T) GetContractByHashHeightInHex(args struct {
	Hash   string
	Height uint64
}, ret *string) error {
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.ctr",
		Index:  "h160.ctr-uint.hgt",
		Keys:   []string{args.Hash, fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
