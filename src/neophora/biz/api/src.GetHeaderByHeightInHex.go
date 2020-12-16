package api

import "fmt"

// GetHeaderByHeightInHex ...
func (me *T) GetHeaderByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "header",
		Index:  "height",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}
