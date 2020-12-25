package api

import (
	"fmt"
	"neophora/var/stderr"
	"net/url"
)

// GetCurrentBlockHeightInUint64 ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"id":1,"jsonrpc":"2.0","method":"GetCurrentBlockHeightInUint64","params":{}}'
// {"id":1,"result":583121,"error":null}
// ```
func (me *T) GetCurrentBlockHeightInUint64(args struct{}, ret *uint64) error {
	var result url.URL
	if err := me.Data.GetLastestUint64KeyInURL(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "h256.blk",
		Index:  "uint.hgt",
		Keys:   []string{},
	}, &result); err != nil {
		return err
	}

	if n, err := fmt.Sscanf(result.Path, "/%x", ret); err != nil {
		return stderr.ErrNotFound
	} else if n != 1 {
		return stderr.ErrNotFound
	}

	return nil
}
