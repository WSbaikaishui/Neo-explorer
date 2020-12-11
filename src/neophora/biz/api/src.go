package api

import (
	"fmt"
	"neophora/biz/data"
)

// T ...
type T struct {
	Data *data.T
}

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

// GetBlockByHashInHex ...
func (me *T) GetBlockByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "block",
		Index:  "hash",
		Keys:   []string{fmt.Sprintf("%016x", args.Hash)},
	}, ret)
}
