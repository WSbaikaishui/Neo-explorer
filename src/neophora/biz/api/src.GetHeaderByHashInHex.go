package api

// GetHeaderByHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetHeaderByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "h256.blk",
		Keys:   []string{args.Hash},
	}, ret)
}
