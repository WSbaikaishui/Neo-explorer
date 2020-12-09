package biz

import (
	"encoding/hex"
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// Data ...
type Data struct {
	Client *cli.T
}

// Ping ...
func (me *Data) Ping(arg struct {
}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// GetDataInHex ...
func (me *Data) GetDataInHex(args struct {
	Key string
}, ret *string) error {
	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(args.Key), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetDataInString ...
func (me *Data) GetDataInString(args struct {
	Key string
}, ret *string) error {
	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(args.Key), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

	return nil
}

// GetBlockByHeightInHex ...
func (me *Data) GetBlockByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "block",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetBlockByHashInHex ...
func (me *Data) GetBlockByHashInHex(args struct {
	Hash string
}, ret *string) error {
	uri := &url.URL{
		Scheme: "block",
		Host:   "hash",
		Path:   fmt.Sprintf("/%s", args.Hash),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetHeaderByHeightInHex ...
func (me *Data) GetHeaderByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "header",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetHeaderByHashInHex ...
func (me *Data) GetHeaderByHashInHex(args struct {
	Hash string
}, ret *string) error {
	uri := &url.URL{
		Scheme: "header",
		Host:   "hash",
		Path:   fmt.Sprintf("/%s", args.Hash),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetStorageByDBKeyHeightInHex ...
func (me *Data) GetStorageByDBKeyHeightInHex(args struct {
	DBKey  string
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "storage",
		Host:   "dbkey-height",
		Path:   fmt.Sprintf("/%s/%016x", args.DBKey, args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetCountInUInt64 ...
func (me *Data) GetCountInUInt64(args struct {
}, ret *uint64) error {
	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/ffffffffffffffff"

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLastKey", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	key, err := url.Parse(string(result))
	if err != nil {
		return stderr.ErrUnknown
	}

	var count uint64

	fmt.Sscanf(key.Path, "/%x", &count)

	*ret = count
	return nil
}

// GetContractByHashHeightInHex ...
func (me *Data) GetContractByHashHeightInHex(args struct {
	Hash   string
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "contract",
		Host:   "hash-height",
		Path:   fmt.Sprintf("/%s/%016x", args.Hash, args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetHashByHeightInHex ...
func (me *Data) GetHashByHeightInHex(args struct {
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "hash",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetTransactionByHashInHex ...
func (me *Data) GetTransactionByHashInHex(args struct {
	Hash uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "tx",
		Host:   "hash",
		Path:   fmt.Sprintf("/%016x", args.Hash),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}
