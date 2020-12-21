package api

import "fmt"

// GetHashByHeightInHex ...
func (me *T) GetHashByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "hash",
		Index:  "height",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
