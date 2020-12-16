package api

import (
	"fmt"
	"neophora/var/stderr"
	"net/url"
)

// GetHeightOfTopInUint64 ...
func (me *T) GetHeightOfTopInUint64(args struct{}, ret *uint64) error {
	var result []byte
	if err := me.Data.GetLastKey(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "hash",
		Index:  "height",
		Keys:   []string{"_"},
	}, &result); err != nil {
		return err
	}

	key, err := url.Parse(string(result))
	if err != nil {
		return stderr.ErrNotFound
	}

	if n, err := fmt.Sscanf(key.Path, "/%x", ret); err != nil {
		return stderr.ErrNotFound
	} else if n != 1 {
		return stderr.ErrNotFound
	}

	return nil
}
