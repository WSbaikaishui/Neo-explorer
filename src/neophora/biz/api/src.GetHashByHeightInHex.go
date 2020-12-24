package api

import "fmt"

// GetHashByHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHashByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "h256.blk",
		Index:  "uint.hgt",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
