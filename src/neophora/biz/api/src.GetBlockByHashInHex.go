package api

// GetBlockByHashInHex ...
func (me *T) GetBlockByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{args.Hash},
	}, ret)
}
