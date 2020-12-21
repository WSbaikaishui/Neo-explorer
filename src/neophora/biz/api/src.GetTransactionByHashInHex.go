package api

// GetTransactionByHashInHex ...
func (me *T) GetTransactionByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "tx",
		Index:  "hash",
		Keys:   []string{args.Hash},
	}, ret)
}
