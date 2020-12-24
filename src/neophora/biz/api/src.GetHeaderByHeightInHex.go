package api

import "fmt"

// GetHeaderByHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
