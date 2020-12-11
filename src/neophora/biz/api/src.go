package api

import (
	"fmt"
	"neophora/biz/data"
	"neophora/var/stderr"
	"net/url"
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
		Keys:   []string{args.Hash},
	}, ret)
}

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

// GetHashByHeightInHex ...
func (me *T) GetHashByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "hash",
		Index:  "height",
		Keys:   []string{fmt.Sprintf("%016x", args.Height)},
	}, ret)
}

// GetTransactionByHashInHex ...
func (me *T) GetTransactionByHashInHex(args struct {
	Hash string
}, ret *string) error {
	return me.Data.GetArgsHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "tx",
		Index:  "hash",
		Keys:   []string{args.Hash},
	}, ret)
}

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

// GetContractByHashHeightInHex ...
func (me *T) GetContractByHashHeightInHex(args struct {
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
		Target: "contract",
		Index:  "hash-height",
		Keys:   []string{args.Hash, "_"},
		Max:    args.Height,
		Min:    0,
	}, ret)
}

// GetCountInUInt64 ...
func (me *T) GetCountInUInt64(args struct{}, ret *uint64) error {
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
