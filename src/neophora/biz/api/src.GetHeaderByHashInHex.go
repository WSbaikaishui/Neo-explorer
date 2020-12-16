package api

// GetHeaderByHashInHex ...
func (me *T) GetHeaderByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "header",
		Index:  "hash",
		Keys:   []string{args.Hash},
	}, ret)
}
