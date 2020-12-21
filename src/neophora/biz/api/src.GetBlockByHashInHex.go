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
		Target: "block",
		Index:  "hash",
		Keys:   []string{args.Hash},
	}, ret)
}
