package api

// GetStorageByHashHeightInHex ...
func (me *T) GetStorageByHashHeightInHex(args struct {
	Hash   string
	Height uint64
}, ret *string) error {
	return me.Data.GetLastExHex(struct {
		Target string
		Index  string
		Keys   []string
		Max    uint64
		Min    uint64
	}{
		Target: "storage",
		Index:  "hash-height",
		Keys:   []string{args.Hash, "_"},
		Max:    args.Height,
		Min:    0,
	}, ret)
}
