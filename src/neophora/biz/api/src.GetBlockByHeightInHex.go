package api

import "fmt"

// GetBlockByHeightInHex ...
func (me *T) GetBlockByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "block",
		Index:  "height",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
